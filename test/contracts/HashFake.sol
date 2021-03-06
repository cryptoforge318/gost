// SPDX-License-Identifier: UNLICENSED

/**
  @dev HashFake.sol is written specfically to test the functions which exist in our Hash.sol "embedded" library
*/

pragma solidity 0.8.0;

import './Hash.sol';

contract HashFake {
  /// @dev convenience method to get the domain type hash
  function domainTypeHash() public pure returns (bytes32) {
    return keccak256(abi.encodePacked(
      'EIP712Domain(',
      'string name,',
      'string version,',
      'uint256 chainId,',
      'address verifyingContract',
      ')'
    ));
  }

  function domainTest(string memory n, string memory version, uint256 c, address verifier) public pure returns (bytes32) {
    return Hash.domain(n, version, c, verifier);  
  }

  function messageTest(bytes32 d, bytes32 h) public pure returns (bytes32) {
    return Hash.message(d, h);
  }

  /// @dev convenience method to get the order type hash
  function orderTypeHash() public pure returns (bytes32) {
    return keccak256(abi.encodePacked(
      'Order(',
      'bytes32 key,',
      'address maker,',
      'address underlying,',
      'bool floating,',
      'uint256 principal,',
      'uint256 interest,',
      'uint256 duration,',
      'uint256 expiry',
      ')'
    ));
  }

  function orderTest(Hash.Order calldata o) external pure returns (bytes32) {
    return Hash.order(o);
  }
}
