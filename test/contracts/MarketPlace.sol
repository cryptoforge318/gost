// SPDX-License-Identifier: UNLICENSED

// TODO update to 0.8.4 (or whatever latest is...)

// NOTE the pattern [underlying, maturity*, cToken, ...]

pragma solidity 0.8.0;

// TODO ZcToken
// TODO VaultTracker

// TODO what is the relationship with Swivel? Does swivel call here? Does here call Swivel?
contract MarketPlace {
  address public admin = msg.sender;

  struct Market {
    address cTokenAddr;
    address zcTokenAddr;
    address vaultAddr; // TODO explore the relationship between MarketPlace.Market <-> VaultTracker.Vault (and the relationship of Swivel to both)
  }

  mapping (address => mapping (uint256 => Market)) public markets;
  mapping (address => mapping (uint256 => bool)) public mature;
  mapping (address => mapping (uint256 => uint256)) public maturityRate;

  // TODO what should be indexed here?
  event Create(address underlying, uint256 maturity, address cToken, address zcToken); 
  event Mature(address underlying, uint256 maturity, uint256 maturityRate, uint256 maturedOn);

  // TODO constructor - likely passing the Swivel contract address...

    /// @notice Allows the owner to create new markets
    /// @param n Name of the new zcToken market
    /// @param s Symbol of the new zcToken market
    /// @param u Underlying token address associated with the new market
    /// @param m Maturity timestamp of the new market
    /// @param c cToken address associated with underlying for the new market 
  function createMarket(
    string memory n,
    string memory s,
    address u,
    uint256 m,
    address c
  ) public restricted(admin) returns (bool) {
    // TODO can we live with the factory pattern here both bytecode size wise and CREATE opcode cost wise?
    // address zctAddr = address(new ZcToken(n, s, u, m));
    // address vAddr = address(new VaultTracker(u, m, c));
    // markets[u][m] = Market(c, zctAddr, vAddr);

    // emit Create(u, m, c, zctAddr);

    return true;
  }

  function matureMarket(address u, uint256 m) public returns (bool) {
    // require mature...
    // window...
    // CErc20...
    // zcToken...
    // Vault...

    // require zcToken.maturity...

    // maturity rate stuff...

    // why does vault have matureMarket ?

    // set mature...

    // emit mature

    return true;
  }

  function redeemZcToken(address u, uint256 m, uint256 a) public returns (bool) {
    // ...  

    return true;
  }

  function reddemVaultInterest(address u, uint256 m) public returns (bool) {

    return true;
  }

  function calculateReturn(address u, uint256 m, uint256 a) internal returns (uint256) {
    uint256 total;

    return total;
  }

  modifier restricted(address a) {
    require(msg.sender == a, 'sender must be admin');
    _;
  }
}
