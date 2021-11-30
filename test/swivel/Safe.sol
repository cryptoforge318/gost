// SPDX-License-Identifier: UNLICENSED
// Adapted from: https://github.com/Rari-Capital/solmate/blob/main/src/utils/SafeTransferLib.sol
pragma solidity 0.8.4;

import {Erc20} from "./Interfaces.sol";

/// @notice Safe ETH and ERC20 transfer library that gracefully handles missing return values.
/// @author Modified from Gnosis (https://github.com/gnosis/gp-v2-contracts/blob/main/src/contracts/libraries/GPv2SafeERC20.sol)
/// @dev Use with caution! Some functions in this library knowingly create dirty bits at the destination of the free memory pointer.
library Safe {
  function transferFrom(
    Erc20 token,
    address from,
    address to,
    uint256 amount
  ) internal {
    bool callStatus;

    assembly {
      // Get a pointer to some free memory.
      let freeMemoryPointer := mload(0x40)

      // Write the abi-encoded calldata to memory piece by piece:
      mstore(freeMemoryPointer, 0x23b872dd00000000000000000000000000000000000000000000000000000000) // Begin with the function selector.
      mstore(add(freeMemoryPointer, 4), and(from, 0xffffffffffffffffffffffffffffffffffffffff)) // Mask and append the "from" argument.
      mstore(add(freeMemoryPointer, 36), and(to, 0xffffffffffffffffffffffffffffffffffffffff)) // Mask and append the "to" argument.
      mstore(add(freeMemoryPointer, 68), amount) // Finally append the "amount" argument. No mask as it's a full 32 byte value.

      // Call the token and store if it succeeded or not.
      // We use 100 because the calldata length is 4 + 32 * 3.
      callStatus := call(gas(), token, 0, freeMemoryPointer, 100, 0, 0)
    }

    require(callSuccess(callStatus), "transfer from failed");
  }

  function transfer(
    Erc20 token,
    address to,
    uint256 amount
  ) internal {
    bool callStatus;

    assembly {
      // Get a pointer to some free memory.
      let freeMemoryPointer := mload(0x40)

      // Write the abi-encoded calldata to memory piece by piece:
      mstore(freeMemoryPointer, 0xa9059cbb00000000000000000000000000000000000000000000000000000000) // Begin with the function selector.
      mstore(add(freeMemoryPointer, 4), and(to, 0xffffffffffffffffffffffffffffffffffffffff)) // Mask and append the "to" argument.
      mstore(add(freeMemoryPointer, 36), amount) // Finally append the "amount" argument. No mask as it's a full 32 byte value.

      // Call the token and store if it succeeded or not.
      // We use 68 because the calldata length is 4 + 32 * 2.
      callStatus := call(gas(), token, 0, freeMemoryPointer, 68, 0, 0)
    }

    require(callSuccess(callStatus), "transfer failed");
  }

  function approve(
    Erc20 token,
    address to,
    uint256 amount
  ) internal {
    bool callStatus;

    assembly {
      // Get a pointer to some free memory.
      let freeMemoryPointer := mload(0x40)

      // Write the abi-encoded calldata to memory piece by piece:
      mstore(freeMemoryPointer, 0x095ea7b300000000000000000000000000000000000000000000000000000000) // Begin with the function selector.
      mstore(add(freeMemoryPointer, 4), and(to, 0xffffffffffffffffffffffffffffffffffffffff)) // Mask and append the "to" argument.
      mstore(add(freeMemoryPointer, 36), amount) // Finally append the "amount" argument. No mask as it's a full 32 byte value.

      // Call the token and store if it succeeded or not.
      // We use 68 because the calldata length is 4 + 32 * 2.
      callStatus := call(gas(), token, 0, freeMemoryPointer, 68, 0, 0)
    }

    require(callSuccess(callStatus), "approve failed");
  }

  function callSuccess(bool callStatus) private pure returns (bool success) {
    assembly {
      // Get how many bytes the call returned.
      let returnDataSize := returndatasize()

      // If the call reverted:
      if iszero(callStatus) {
        // Copy the revert message into memory.
        returndatacopy(0, 0, returnDataSize)

        // Revert with the same message.
        revert(0, returnDataSize)
      }

      switch returnDataSize
      case 32 {
        // Copy the return data into memory.
        returndatacopy(0, 0, returnDataSize)

        // Set success to whether it returned true.
        success := iszero(iszero(mload(0)))
      }
      case 0 {
        // There was no return data.
        success := 1
      }
      default {
        // It returned some malformed input.
        success := 0
      }
    }
  }
}
