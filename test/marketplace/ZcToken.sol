// SPDX-License-Identifier: UNLICENSED

/**
  ZcToken is a mock which records arguments passed to its methods as well as
  provides setters allowing us to dictate method return values
*/

pragma solidity 0.8.13;

contract ZcToken {
  // mapping for arguments passed to approve
  mapping (address => uint256) public approveCalled;

  address private underlyingReturn;
  uint256 private maturityReturn;
  bool private approveReturn;
  string public name;
  string public symbol;
  uint8 public decimals;
  
  mapping (address => uint256) private balanceOfReturn;

  /// @param u Underlying
  /// @param m Maturity
  /// @param n Name
  /// @param s Symbol
  /// @param d Decimals
  constructor(address u, uint256 m, string memory n, string memory s, uint8 d) {
    // we can set the privates in the constructor as well...
    underlyingReturn = u;
    maturityReturn = m;
    // set approve to return true by default so that createMarket won't revert
    approveReturn = true;

    name = n;
    symbol = s;
    decimals = d;
  }

  function approve(address s, uint256 a) public returns (bool) {
    approveCalled[s] = a;
    return approveReturn;
  }

  function approveReturns(bool b) public {
    approveReturn = b;
  }

  function balanceOf(address a) public view returns (uint256) {
    return balanceOfReturn[a];
  }

  function balanceOfReturns(address u, uint256 a) public {
    balanceOfReturn[u] = a;
  }
}
