// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Interfaces.sol';
import './Compounding.sol';
import './ZcToken.sol';
import './VaultTracker.sol';

contract MarketPlace {
  /// @dev A single custom error capable of indicating a wide range of detected errors by providing
  /// an error code value whose string representation is documented <here>, and any possible other values
  /// that are pertinent to the error.
  error Exception(uint8, uint256, uint256, address, address);

  struct Market {
    address cTokenAddr;
    address zcToken;
    address vaultTracker;
    uint256 maturityRate;
  }

  mapping (uint8 => mapping (address => mapping (uint256 => Market))) public markets;

  address public admin;
  address public swivel;
  bool public paused;

  event Create(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address cToken, address zcToken, address vaultTracker);
  event Mature(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, uint256 maturityRate, uint256 matured);
  event RedeemZcToken(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address sender, uint256 amount);
  event RedeemVaultInterest(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address sender);
  event CustodialInitiate(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address zcTarget, address nTarget, uint256 amount);
  event CustodialExit(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address zcTarget, address nTarget, uint256 amount);
  event P2pZcTokenExchange(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount);
  event P2pVaultExchange(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount);
  event TransferVaultNotional(uint8 indexed protocol, address indexed underlying, uint256 indexed maturity, address from, address to, uint256 amount);

  constructor() {
    admin = msg.sender;
  }

  /// @param s Address of the deployed swivel contract
  /// @notice We only allow this to be set once
  function setSwivel(address s) external authorized(admin) returns (bool) {
    // TODO remove on confirm
    // require(swivel == address(0), 'swivel contract address already set');
    if (swivel != address(0)) {
      revert Exception(20, 0, 0, swivel, address(0)); 
    }

    swivel = s;
    return true;
  }

  /// @param a Address of a new admin
  function setAdmin(address a) external authorized(admin) returns (bool) {
    admin = a;
    return true;
  }

  /// @notice Allows the owner to create new markets
  /// @param p Protocol associated with the new market
  /// @param m Maturity timestamp of the new market
  /// @param c Compounding Token address associated with the new market
  /// @param n Name of the new market zcToken
  /// @param s Symbol of the new market zcToken
  function createMarket(
    uint8 p,
    uint256 m,
    address c,
    string memory n,
    string memory s
  ) external authorized(admin) unpaused() returns (bool) {
    // TODO remove on confirm
    // require(swivel != address(0), 'swivel contract address not set');
    if (swivel == address(0)) {
      revert Exception(21, 0, 0, address(0), address(0));
    }

    address underAddr = Compounding.underlying(p, c);

    // TODO remove on confirm
    // require(markets[p][underAddr][m].vaultTracker == address(0), 'market already exists');
    if (markets[p][underAddr][m].vaultTracker != address(0)) {
      // NOTE: not saving and publishing that found tracker addr as stack limitations...
      revert Exception(22, 0, 0, address(0), address(0));
    }

    address zct;
    address tracker;

    {
      uint8 decimals = IErc20(underAddr).decimals();
      zct = address(new ZcToken(underAddr, m, n, s, decimals));
      tracker = address(new VaultTracker(p, m, c, swivel));
    }

    markets[p][underAddr][m] = Market(c, zct, tracker, 0);

    emit Create(p, underAddr, m, c, zct, tracker);

    return true;
  }

  /// @notice Can be called after maturity, allowing all of the zcTokens to earn floating interest on Compound until they release their funds
  /// @param p Protocol Enum value associated with the market being matured
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  function matureMarket(uint8 p, address u, uint256 m) public unpaused() returns (bool) {
    Market memory market = markets[p][u][m];

    // TODO remove on confirm
    // require(market.maturityRate == 0, 'market already matured');
    if (market.maturityRate != 0) {
      revert Exception(23, market.maturityRate, 0, address(0), address(0));
    }

    // TODO remove on confirm
    // require(block.timestamp >= m, 'maturity not reached');
    if (block.timestamp < m) {
      revert Exception(24, block.timestamp, m, address(0), address(0));
    }

    // set the base maturity cToken exchange rate at maturity to the current cToken exchange rate
    uint256 exchangeRate = Compounding.exchangeRate(p, market.cTokenAddr);
    markets[p][u][m].maturityRate = exchangeRate;

    // TODO remove the require here? why?
    // TODO remove on confirm
    // require(VaultTracker(market.vaultTracker).matureVault(exchangeRate), 'mature vault failed');
    if (!VaultTracker(market.vaultTracker).matureVault(exchangeRate)) {
      revert Exception(25, 0, 0, address(0), address(0));
    }

    emit Mature(p, u, m, exchangeRate, block.timestamp);

    return true;
  }

  /// @notice Allows Swivel caller to deposit their underlying, in the process splitting it - minting both zcTokens and vault notional.
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Address of the depositing user
  /// @param a Amount of notional being added
  function mintZcTokenAddingNotional(uint8 p, address u, uint256 m, address t, uint256 a) external authorized(swivel) unpaused() returns (bool) {
    Market memory market = markets[p][u][m];

    // TODO remove on confirm
    // require(ZcToken(market.zcToken).mint(t, a), 'mint zcToken failed');
    if (!ZcToken(market.zcToken).mint(t, a)) {
      revert Exception(29, 0, 0, address(0), address(0));
    }

    // TODO remove on confirm
    // require(VaultTracker(market.vaultTracker).addNotional(t, a), 'add notional failed');
    if (!VaultTracker(market.vaultTracker).addNotional(t, a)) {
      revert Exception(26, 0, 0, address(0), address(0));
    }
    
    return true;
  }

  /// @notice Allows Swivel caller to deposit/burn both zcTokens + vault notional. This process is "combining" the two and redeeming underlying.
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Address of the combining/redeeming user
  /// @param a Amount of zcTokens being burned
  function burnZcTokenRemovingNotional(uint8 p, address u, uint256 m, address t, uint256 a) external authorized(swivel) unpaused() returns(bool) {
    Market memory market = markets[p][u][m];

    // TODO remove on confirm
    // require(ZcToken(market.zcToken).burn(t, a), 'burn failed');
    if (!ZcToken(market.zcToken).burn(t, a)) {
      revert Exception(30, 0, 0, address(0), address(0));
    }

    // TODO remove on confirm
    // require(VaultTracker(market.vaultTracker).removeNotional(t, a), 'remove notional failed');
    if (!VaultTracker(market.vaultTracker).removeNotional(t, a)) {
      revert Exception(27, 0, 0, address(0), address(0));
    }
    
    return true;
  }

  /// @notice Allows (via swivel) zcToken holders to redeem their tokens for underlying tokens after maturity has been reached.
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Address of the redeeming user
  /// @param a Amount of zcTokens being redeemed
  function redeemZcToken(uint8 p, address u, uint256 m, address t, uint256 a) external authorized(swivel) unpaused() returns (uint256) {
    Market memory market = markets[p][u][m];

    // if the market has not matured, mature it and redeem exactly the amount
    if (market.maturityRate == 0) {
      // TODO remove on confirm
      // require(matureMarket(p, u, m), 'failed to mature the market');
      if (!matureMarket(p, u, m)) {
        revert Exception(31, 0, 0, address(0), address(0));
      }
    }

    // burn user's zcTokens
    // TODO remove on confirm
    // require(ZcToken(market.zcToken).burn(t, a), 'could not burn');
    if (!ZcToken(market.zcToken).burn(t, a)) {
      revert Exception(30, 0, 0, address(0), address(0));
    }

    emit RedeemZcToken(p, u, m, t, a);

    if (market.maturityRate == 0) {
      return a;
    } else { 
      // if the market was already mature the return should include the amount + marginal floating interest generated on Compound since maturity
      return calculateReturn(p, u, m, a);
    }
  }

  /// @notice Allows Vault owners (via Swivel) to redeem any currently accrued interest
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Address of the redeeming user
  function redeemVaultInterest(uint8 p, address u, uint256 m, address t) external authorized(swivel) unpaused() returns (uint256) {
    // call to the floating market contract to release the position and calculate the interest generated
    uint256 interest = VaultTracker(markets[p][u][m].vaultTracker).redeemInterest(t);

    emit RedeemVaultInterest(p, u, m, t);

    return interest;
  }

  /// @notice Calculates the total amount of underlying returned including interest generated since the `matureMarket` function has been called
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param a Amount of zcTokens being redeemed
  function calculateReturn(uint8 p, address u, uint256 m, uint256 a) internal view returns (uint256) {
    Market memory market = markets[p][u][m];

    uint256 exchangeRate = Compounding.exchangeRate(p, market.cTokenAddr);

    return (a * exchangeRate) / market.maturityRate;
  }

  /// @notice Return the compounding token address for a given market
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  function cTokenAddress(uint8 p, address u, uint256 m) external view returns (address) {
    Market memory market = markets[p][u][m];
    return market.cTokenAddr;
  }

  /// @notice Called by swivel IVFZI && IZFVI
  /// @dev Call with protocol, underlying, maturity, mint-target, add-notional-target and an amount
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param z Recipient of the minted zcToken
  /// @param n Recipient of the added notional
  /// @param a Amount of zcToken minted and notional added
  function custodialInitiate(uint8 p, address u, uint256 m, address z, address n, uint256 a) external authorized(swivel) unpaused() returns (bool) {
    Market memory market = markets[p][u][m];
    // TODO remove on confirm
    // require(ZcToken(market.zcToken).mint(z, a), 'mint failed');
    if (!ZcToken(market.zcToken).mint(z, a)) {
      revert Exception(29, 0, 0, address(0), address(0));
    }

    // TODO remove on confirm
    // require(VaultTracker(market.vaultTracker).addNotional(n, a), 'add notional failed');
    if (!VaultTracker(market.vaultTracker).addNotional(n, a)) {
      revert Exception(26, 0, 0, address(0), address(0));
    }

    emit CustodialInitiate(p, u, m, z, n, a);
    return true;
  }

  /// @notice Called by swivel EVFZE FF EZFVE
  /// @dev Call with protocol, underlying, maturity, burn-target, remove-notional-target and an amount
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param z Owner of the zcToken to be burned
  /// @param n Target to remove notional from
  /// @param a Amount of zcToken burned and notional removed
  function custodialExit(uint8 p, address u, uint256 m, address z, address n, uint256 a) external authorized(swivel) unpaused() returns (bool) {
    Market memory market = markets[p][u][m];
    // TODO remove on confirm
    // require(ZcToken(market.zcToken).burn(z, a), 'burn failed');
    if (!ZcToken(market.zcToken).burn(z, a)) {
      revert Exception(30, 0, 0, address(0), address(0));
    }

    // TODO remove on confirm
    // require(VaultTracker(market.vaultTracker).removeNotional(n, a), 'remove notional failed');
    if (!VaultTracker(market.vaultTracker).removeNotional(n, a)) {
      revert Exception(27, 0, 0, address(0), address(0));
    }

    emit CustodialExit(p, u, m, z, n, a);
    return true;
  }

  /// @notice Called by swivel IZFZE, EZFZI
  /// @dev Call with underlying, maturity, transfer-from, transfer-to, amount
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param f Owner of the zcToken to be burned
  /// @param t Target to be minted to
  /// @param a Amount of zcToken transfer
  function p2pZcTokenExchange(uint8 p, address u, uint256 m, address f, address t, uint256 a) external authorized(swivel) unpaused() returns (bool) {
    Market memory market = markets[p][u][m];
    // TODO remove on confirm
    // require(ZcToken(market.zcToken).burn(f, a), 'zcToken burn failed');
    if (!ZcToken(market.zcToken).burn(f, a)) {
      revert Exception(30, 0, 0, address(0), address(0));
    }

    // TODO remove on confirm
    // require(ZcToken(market.zcToken).mint(t, a), 'zcToken mint failed');
    if (!ZcToken(market.zcToken).mint(t, a)) {
      revert Exception(29, 0, 0, address(0), address(0));
    }

    emit P2pZcTokenExchange(p, u, m, f, t, a);
    return true;
  }

  /// @notice Called by swivel IVFVE, EVFVI
  /// @dev Call with protocol, underlying, maturity, remove-from, add-to, amount
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param f Owner of the notional to be transferred
  /// @param t Target to be transferred to
  /// @param a Amount of notional transfer
  function p2pVaultExchange(uint8 p, address u, uint256 m, address f, address t, uint256 a) external authorized(swivel) unpaused() returns (bool) {
    // TODO remove on confirm
    // require(VaultTracker(markets[p][u][m].vaultTracker).transferNotionalFrom(f, t, a), 'transfer notional failed');
    if (!VaultTracker(markets[p][u][m].vaultTracker).transferNotionalFrom(f, t, a)) {
      revert Exception(28, 0, 0, address(0), address(0));
    }

    emit P2pVaultExchange(p, u, m, f, t, a);
    return true;
  }

  /// @notice External method giving access to this functionality within a given vault
  /// @dev Note that this method calculates yield and interest as well
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param t Target to be transferred to
  /// @param a Amount of notional to be transferred
  function transferVaultNotional(uint8 p, address u, uint256 m, address t, uint256 a) external unpaused() returns (bool) {
    // TODO remove on confirm
    // require(VaultTracker(markets[p][u][m].vaultTracker).transferNotionalFrom(msg.sender, t, a), 'vault transfer failed');
    if (!VaultTracker(markets[p][u][m].vaultTracker).transferNotionalFrom(msg.sender, t, a)) {
      revert Exception(28, 0, 0, address(0), address(0));
    }

    emit TransferVaultNotional(p, u, m, msg.sender, t, a);
    return true;
  }

  /// @notice Transfers notional fee to the Swivel contract without recalculating marginal interest for from
  /// @param p Protocol Enum value associated with this market
  /// @param u Underlying token address associated with the market
  /// @param m Maturity timestamp of the market
  /// @param f Owner of the amount
  /// @param a Amount to transfer
  function transferVaultNotionalFee(uint8 p, address u, uint256 m, address f, uint256 a) external authorized(swivel) returns (bool) {
    VaultTracker(markets[p][u][m].vaultTracker).transferNotionalFee(f, a);
    return true;
  }

  // TODO this becomes per-protocol
  /// @notice Called by admin at any point to pause / unpause market transactions
  /// @param b Boolean which indicates the markets paused status
  function pause(bool b) external authorized(admin) returns (bool) {
    paused = b;
    return true;
  }

  modifier authorized(address a) {
    if(msg.sender != a) {
      revert Exception(0, 0, 0, msg.sender, a);
    }
    _;
  }

  modifier unpaused() {
    if(paused) {
      revert Exception(1, 0, 0, address(0), address(0));
    }
    _;
  }
}
