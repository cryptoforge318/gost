// SPDX-License-Identifier: BUSL-1.1

pragma solidity >=0.8.4;

import './Interfaces.sol';
import './ZcToken.sol';
import './VaultTracker.sol';
import './IAdapter.sol';

contract MarketPlace {
  struct Market {
    address cToken;
    address zcToken;
    address vault;
    address adapter;
    uint256 maturityRate;
  }
  // Protocol enum => Underlying address => Maturity uint256 => Market
  mapping (uint8 => mapping (address => mapping (uint256 => Market))) public markets;

  address public admin;
  address public swivel;
  bool public paused;

  event Create(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address cToken, address zcToken, address vaultTracker, address adapter);
  event Mature(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, uint256 maturityRate, uint256 matured);
  event RedeemZcToken(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address sender, uint256 amount);
  event RedeemVaultInterest(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address sender);
  event CustodialInitiate(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address zcTarget, address nTarget, uint256 amount);
  event CustodialExit(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address zcTarget, address nTarget, uint256 amount);
  event P2pZcTokenExchange(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount);
  event P2pVaultExchange(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount);
  event TransferVaultNotional(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount);

  error Initialize();
  error Maturity();
  error Paused();
  error Unauthorized();
  error Market_Exists();

  constructor() {
    admin = msg.sender;
  }

  /// @param s Address of the deployed swivel contract
  /// @notice We only allow this to be set once
  function setSwivelAddress(address s) external authorized(admin) returns (bool) {
    if (swivel != address(0)) {
      revert Initialize();
    }

    swivel = s;
    return true;
  }

  /// @param a Address of a new admin
  function transferAdmin(address a) external authorized(admin) returns (bool) {
    admin = a;
    return true;
  }

  /// @notice Allows the owner to create new markets
  /// @param m Maturity timestamp of the new market
  /// @param c cToken address associated with underlying for the new market
  /// @param a Adapter address associated with the given lending market
  /// @param n Name of the new zcToken market
  /// @param s Symbol of the new zcToken market
  function createMarket(
    uint256 m,
    address c,
    address a,
    string memory n,
    string memory s
  ) external authorized(admin) unpaused() returns (bool) {
    if (swivel == address(0)) {
      revert Initialize();
    }
    uint8 protocol = IAdapter(a).protocol();
    address under = CErc20(c).underlying();
    
    if (markets[protocol][under][m].vault != address(0)) {
      revert Market_Exists();
    }
    
    address zcToken = address(new ZcToken(under, m, n, s, Erc20(under).decimals()));
    address vault = address(new VaultTracker(m, c, a, swivel));
    markets[protocol][under][m] = Market(c, zcToken, vault, a, 0);

    emit Create(protocol, under, m, c, zcToken, vault, a);

    return true;
  }

  /// @notice Can be called after maturity, allowing all of the zcTokens to earn floating interest on Compound until they release their funds
  /// @param p Protocol of the given market (Enum)
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  function matureMarket(uint8 p, address u, uint256 m) public unpaused() returns (bool) {
    Market memory mkt = markets[p][u][m];

    if (mkt.maturityRate != 0 || block.timestamp < m) {
      revert Maturity();
    }

    // set the base maturity cToken exchange rate at maturity to the current cToken exchange rate
    uint256 currentExchangeRate = IAdapter(mkt.adapter).exchangeRateCurrent(mkt.cToken);
    markets[p][u][m].maturityRate = currentExchangeRate;

    VaultTracker(mkt.vault).matureVault(currentExchangeRate);

    emit Mature(p, u, m, currentExchangeRate, block.timestamp);

    return true;
  }

  /// @notice Allows Swivel caller to deposit their underlying, in the process splitting it - minting both zcTokens and vault notional.
  /// @param p Protocol of the given market (Enum)
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Address of the depositing user
  /// @param a Amount of notional being added
  function mintZcTokenAddingNotional(uint8 p, address u, uint256 m, address t, uint256 a) external authorized(swivel) unpaused() returns (bool) {
    Market memory mkt = markets[p][u][m];
    ZcToken(mkt.zcToken).mint(t, a);
    VaultTracker(mkt.vault).addNotional(t, a);
    
    return true;
  }

  /// @notice Allows Swivel caller to deposit/burn both zcTokens + vault notional. This process is "combining" the two and redeeming underlying.
  /// @param p Protocol of the given market (Enum)
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Address of the combining/redeeming user
  /// @param a Amount of zcTokens being burned
  function burnZcTokenRemovingNotional(uint8 p, address u, uint256 m, address t, uint256 a) external authorized(swivel) unpaused() returns(bool) {
    Market memory mkt = markets[p][u][m];
    ZcToken(mkt.zcToken).burn(t, a);
    VaultTracker(mkt.vault).removeNotional(t, a);
    
    return true;
  }

  /// @notice Allows (via swivel) zcToken holders to redeem their tokens for underlying tokens after maturity has been reached.
  /// @param p Protocol of the given market (Enum)
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Address of the redeeming user
  /// @param a Amount of zcTokens being redeemed
  function redeemZcToken(uint8 p, address u, uint256 m, address t, uint256 a) external authorized(swivel) unpaused() returns (uint256) {
    Market memory mkt = markets[p][u][m];

    // if the market has not matured, mature it and redeem exactly the amount
    if (mkt.maturityRate == 0) {
      matureMarket(p, u, m);
    }

    // burn user's zcTokens
    ZcToken(mkt.zcToken).burn(t, a);

    emit RedeemZcToken(p, u, m, t, a);

    if (mkt.maturityRate == 0) {
      return a;
    } else { 
      // if the market was already mature the return should include the amount + marginal floating interest generated on Compound since maturity
      return calculateReturn(p, u, m, a);
    }
  }

  /// @notice Allows Vault owners (via Swivel) to redeem any currently accrued interest
  /// @param p Protocol of the given market (Enum)
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Address of the redeeming user
  function redeemVaultInterest(uint8 p, address u, uint256 m, address t) external authorized(swivel) unpaused() returns (uint256) {
    // call to the floating market contract to release the position and calculate the interest generated
    uint256 interest = VaultTracker(markets[p][u][m].vault).redeemInterest(t);

    emit RedeemVaultInterest(p, u, m, t);

    return interest;
  }

  /// @notice Calculates the total amount of underlying returned including interest generated since the `matureMarket` function has been called
  /// @param p Protocol of the given market (Enum)
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param a Amount of zcTokens being redeemed
  function calculateReturn(uint8 p, address u, uint256 m, uint256 a) internal view returns (uint256) {
    Market memory mkt = markets[p][u][m];
    uint256 rate = IAdapter(mkt.adapter).exchangeRateCurrent(mkt.cToken);

    return (a * rate) / mkt.maturityRate;
  }

  /// @notice Return the ctoken address for a given market
  /// @param p Protocol of the given market (Enum)
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  function cTokenAddress(uint8 p, address u, uint256 m) external view returns (address) {
    return markets[p][u][m].cToken;
  }

  /// @notice Called by swivel IVFZI && IZFVI
  /// @dev Call with underlying, maturity, mint-target, add-notional-target and an amount
  /// @param p Protocol of the given market (Enum)
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param z Recipient of the minted zcToken
  /// @param n Recipient of the added notional
  /// @param a Amount of zcToken minted and notional added
  function custodialInitiate(uint8 p, address u, uint256 m, address z, address n, uint256 a) external authorized(swivel) unpaused() returns (bool) {
    Market memory mkt = markets[p][u][m];
    ZcToken(mkt.zcToken).mint(z, a);
    VaultTracker(mkt.vault).addNotional(n, a);
    emit CustodialInitiate(p, u, m, z, n, a);
    return true;
  }

  /// @notice Called by swivel EVFZE FF EZFVE
  /// @dev Call with underlying, maturity, burn-target, remove-notional-target and an amount
  /// @param p Protocol of the given market (Enum)
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param z Owner of the zcToken to be burned
  /// @param n Target to remove notional from
  /// @param a Amount of zcToken burned and notional removed
  function custodialExit(uint8 p, address u, uint256 m, address z, address n, uint256 a) external authorized(swivel) unpaused() returns (bool) {
    Market memory mkt = markets[p][u][m];
    ZcToken(mkt.zcToken).burn(z, a);
    VaultTracker(mkt.vault).removeNotional(n, a);
    emit CustodialExit(p, u, m, z, n, a);
    return true;
  }

  /// @notice Called by swivel IZFZE, EZFZI
  /// @dev Call with underlying, maturity, transfer-from, transfer-to, amount
  /// @param p Protocol of the given market (Enum)
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param f Owner of the zcToken to be burned
  /// @param t Target to be minted to
  /// @param a Amount of zcToken transfer
  function p2pZcTokenExchange(uint8 p, address u, uint256 m, address f, address t, uint256 a) external authorized(swivel) unpaused() returns (bool) {
    Market memory mkt = markets[p][u][m];
    ZcToken(mkt.zcToken).burn(f, a);
    ZcToken(mkt.zcToken).mint(t, a);
    emit P2pZcTokenExchange(p, u, m, f, t, a);
    return true;
  }

  /// @notice Called by swivel IVFVE, EVFVI
  /// @dev Call with underlying, maturity, remove-from, add-to, amount
  /// @param p Protocol of the given market (Enum)
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param f Owner of the notional to be transferred
  /// @param t Target to be transferred to
  /// @param a Amount of notional transfer
  function p2pVaultExchange(uint8 p, address u, uint256 m, address f, address t, uint256 a) external authorized(swivel) unpaused() returns (bool) {
    VaultTracker(markets[p][u][m].vault).transferNotionalFrom(f, t, a);
    emit P2pVaultExchange(p, u, m, f, t, a);
    return true;
  }

  /// @notice External method giving access to this functionality within a given vault
  /// @dev Note that this method calculates yield and interest as well
  /// @param p Protocol of the given market (Enum)
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Target to be transferred to
  /// @param a Amount of notional to be transferred
  function transferVaultNotional(uint8 p, address u, uint256 m, address t, uint256 a) external unpaused() returns (bool) {
    VaultTracker(markets[p][u][m].vault).transferNotionalFrom(msg.sender, t, a);
    emit TransferVaultNotional(p, u, m, msg.sender, t, a);
    return true;
  }

  /// @notice Transfers notional fee to the Swivel contract without recalculating marginal interest for from
  /// @param p Protocol of the given market (Enum)
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param f Owner of the amount
  /// @param a Amount to transfer
  function transferVaultNotionalFee(uint8 p, address u, uint256 m, address f, uint256 a) external authorized(swivel) returns (bool) {
    VaultTracker(markets[p][u][m].vault).transferNotionalFee(f, a);
    return true;
  }

  /// @notice Called by admin at any point to pause / unpause market transactions
  /// @param b Boolean which indicates the markets paused status
  function pause(bool b) external authorized(admin) returns (bool) {
    paused = b;
    return true;
  }

  modifier authorized(address a) {
     if (msg.sender != a) {
       revert Unauthorized();
     }
    _;
  }

  modifier unpaused() {
    if (paused) {
       revert Paused();
    }
    _;
  }
}