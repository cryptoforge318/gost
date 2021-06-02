// SPDX-License-Identifier: UNLICENSED

/**
  ZcToken is a mock which records arguments passed to its methods as well as
  provides setters allowing us to dictate method return values
*/

pragma solidity 0.8.4;

contract ZcToken {
  // a struct to hold the arguments passed to transferFrom
  struct TransferFromArgs {
    address to;
    uint256 amount;
  }

  // don't feel the need to setup return/returns for the strings...
  string name;
  string symbol;

  address private underlyingReturn;
  uint256 private maturityReturn;

  // mapping of arguments sent to burn. key is the passed in address.
  mapping (address => uint256) public burnCalled;
  // a boolean flag which allows us to dictate the return of burn().
  bool private burnReturn;

  // mapping of arguments sent to mint. key is the passed in address.
  mapping (address => uint256) public mintCalled;
  // a boolean flag which allows us to dictate the return of mint().
  bool private mintReturn;

  // mapping of arguments sent to transferFrom. key is passed from address.
  mapping (address => TransferFromArgs) public transferFromCalled;
  // a boolean flag which allows us to dictate the return of transferFrom().
  bool private transferFromReturn;

  /// @param u Underlying
  /// @param m Maturity
  /// @param n Name
  /// @param s Symbol
  constructor(address u, uint256 m, string memory n, string memory s) {
    // we can set the privates in the constructor as well...
    underlyingReturn = u;
    maturityReturn = m;

    name = n;
    symbol = s;
  }

  function burnReturns(bool b) public {
    burnReturn = b;
  }

  function burn(address f, uint256 a) public returns(bool) {
    burnCalled[f] = a;
    return burnReturn;
  }

  function mintReturns(bool b) public {
    mintReturn = b;
  }

  function mint(address f, uint256 a) public returns(bool) {
    mintCalled[f] = a;
    return mintReturn;
  }

  function underlyingReturns(address u) public {
    underlyingReturn = u;
  }
  
  // override what would be the autogenerated getter...
  function underlying() public view returns (address) {
    return underlyingReturn;
  }

  function maturityReturns(uint256 n) public {
    maturityReturn = n;
  }
  
  // override what would be the autogenerated getter...
  function maturity() public view returns (uint256) {
    return maturityReturn;
  }

  function transferFrom(address f, address t, uint256 a) public returns (bool) {
    TransferFromArgs memory args;
    args.to = t;
    args.amount = a;
    transferFromCalled[f] = args;
    return transferFromReturn;
  }

  function transferFromReturns(bool b) public {
    transferFromReturn = b;
  }
}
