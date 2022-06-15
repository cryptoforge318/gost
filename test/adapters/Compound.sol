// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.13;

import './Interfaces.sol';

contract Compound is ICompounding {
  // TODO this could be public with an authorized setter if we envision ourselves changing them?
  uint8 constant public PROTOCOL = 0;
  
  /// @param c Compounding token address
  function exchangeRate(address c) external override returns (uint256) {
    // in this case it _is_ a compound token...
    return ICompoundToken(c).exchangeRateCurrent();
  }

  /// @param c Compounding token address
  function underlying(address c) external override returns (address) {
    return ICompoundToken(c).underlying();
  }
}
