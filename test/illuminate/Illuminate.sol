// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import "./Interfaces.sol";
import "./ZcToken.sol";
import "./Safe.sol";

contract Illuminate {
  /// @notice for any given illuminate deployment the available principles are contained here
  enum Principals {
    Illuminate,
    Swivel,
    Yield,
    Element,
    Pendle,
    Tempus,
    Sense,
    Apwine
  }

  /// markets are defined by a market pair which point to a fixed length array of principal token addresses.
  /// the principal tokens those addresses represent correspond to their Principals enum value, thus the
  /// array should be ordered in that way
  mapping (address => mapping (uint256 => address[8])) public markets;

  address public admin;
  /// @notice address of the deployed lender contract
  address public lender; // TODO add authorized methods to set these
  /// @notice address of the deployed redeemer contract
  address public redeemer;

  event CreateMarket(address indexed underlying, uint256 indexed maturity);

  // TODO move to lender.sol
  // event Mint(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount);

  constructor() {
    admin = msg.sender;
  }

  /// @param u underlying
  /// @param m maturity
  /// @param t principal token addresses for this market minus the illuminate principal (which is added here)
  /// @param n name for the illuminate token
  /// @param s symbol for the illuminate token
  /// @param d decimals for the illuminate token
  function createMarket(address u, uint256 m, address[7] memory t, string calldata n, string calldata s, uint8 d) external authorized(admin) returns (bool) {
    require(markets[u][m][uint256(Principals.Illuminate)] == address(0), 'market exists'); 

    // deploy an illuminate token with this new market NOTE: ATM is using name as symbol args
    address iToken = address(new ZcToken(u, m, n, s, d));

    // the market will have the illuminate principal as its zeroth item, thus t should have Principals[1] as [0]
    // TODO we could choose to put illuminate last in
    address[8] memory market = [iToken, t[0], t[1], t[2], t[3], t[4], t[5], t[6]];

    uint256 max = 2**256 - 1;
    // approve the underlying for max per given principal
    for (uint8 i; i < 8; i++) {
      Safe.approve(IErc20(t[i]), redeemer, max);
    }

    markets[u][m] = market;

    // TODO "args" for this event?
    emit CreateMarket(u, m);
    // ...
    return true;
  }

  /// @dev mint is uniform across all principals, thus there is no need for a 'minter'
  // @param p value of a specific principal according to the MarketPlace Principals Enum.
  // @param u address of an underlying asset
  // @param m maturity (timestamp) of the market
  // @param a amount being minted
  // TODO move to lender.sol
  // function mint(uint8 p, address u, uint256 m, uint256 a) public returns (bool) {
    // address mPlace = marketPlace(marketPlace);
    // use market interface to fetch the market for the given market pair
    // address[8] market = mPlace.markets(u, m);
    // use safe transfer lib and ERC interface...
    // Safe.transferFrom(Erc20(market[p]), msg.sender, address(this), a);
    // use zctoken interface...
    // ZcToken(market[mPlace.Principals.Illuminate]).mint(msg.sender, a);

    // emit Mint(p, u, m, a);

    // return true;
  // }

  modifier authorized(address a) {
    require(msg.sender == a, 'sender must be authorized');
    _;
  }
}
