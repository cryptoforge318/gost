// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.0;

/// @dev MarketPlace is a mock whose bindings are imported by unit tests in any pkg/*testing that needs it
contract MarketPlace {
  // the address that will be returned when marketCTokenAddress is called
  address private cTokenAddr;
  // the arguments that were passed to marketCTokenAddress when it was called
  // TODO: we should likely standardize on using the `Called` suffix for these mappings in the mocks.
  //       the token mocks use a differnt pattern, change them to match this...
  mapping (address => uint256) public cTokenAddressCalled;
  bool private initiateFillingInitiateReturn;
  bool private exitFillingExitReturn;
  bool private transferFromZcTokenReturn;
  bool private transferFromNotionalReturn;

  struct MethodArgs {
    uint256 maturity;
    address one; // is sender or maker depending on method
    address two; // same as above
    uint256 amount;
  }

  mapping (address => MethodArgs) public initiateFillingInitiateCalled;
  mapping (address => MethodArgs) public exitFillingExitCalled;
  mapping (address => MethodArgs) public transferFromZcTokenCalled;
  mapping (address => MethodArgs) public transferFromNotionalCalled;

  function cTokenAddressReturns(address a) external {
    cTokenAddr = a;
  }

  function cTokenAddress(address u, uint256 m) external returns (address) {
    cTokenAddressCalled[u] = m;
    return cTokenAddr;
  }

  function initiateFillingInitiateReturns(bool b) external {
    initiateFillingInitiateReturn = b;
  }

  // called by swivel IVFZI && IZFVI 
  // call with underlying, maturity, mint-target, add-notional-target and an amount
  function initiateFillingInitiate(address u, uint256 m, address o, address t, uint256 a) external returns (bool) {
    MethodArgs memory args; 
    args.maturity = m;
    args.one = o; // will be the recipient of minted zctoken
    args.two = t; // will be the recipient of added notional
    args.amount = a; // the amount of minted zctoken and notional added
    initiateFillingInitiateCalled[u] = args;

    return initiateFillingInitiateReturn;
  }

  function exitFillingExitReturns(bool b) external {
    exitFillingExitReturn = b;
  }

  // called by swivel EVFZE FF EZFVE
  // call with underlying, maturity, burn-target, remove-notional-target and an amount
  function exitFillingExit(address u, uint256 m, address o, address t, uint256 a) external returns (bool) {
    MethodArgs memory args; 
    args.maturity = m;
    args.one = o; // will be the burn-from target
    args.two = t; // will be the remove-notional target
    args.amount = a; // zctoken burned, notional removed
    exitFillingExitCalled[u] = args;

    return exitFillingExitReturn;
  }

  function transferFromZcTokenReturns(bool b) external {
    transferFromZcTokenReturn = b;
  }

  function transferFromZcToken(address u, uint256 m, address o, address t, uint256 a) external returns (bool) {
    MethodArgs memory args;
    args.maturity = m;
    args.one = o;
    args.two = t;
    args.amount = a;
    transferFromZcTokenCalled[u] = args;

    return transferFromZcTokenReturn;
  }

  function transferFromNotionalReturns(bool b) external {
    transferFromNotionalReturn = b;
  }

  function transferFromNotional(address u, uint256 m, address o, address t, uint256 a) external returns (bool) {
    MethodArgs memory args;
    args.maturity = m;
    args.one = o;
    args.two = t;
    args.amount = a;
    transferFromNotionalCalled[u] = args;

    return transferFromNotionalReturn;
  }
}
