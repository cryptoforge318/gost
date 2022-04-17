// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.8.4;

import "./LibFuse.sol";

contract RariAdapter {

    function exchangeRateCurrent(address a) external view returns (uint256){
        return(LibFuse.viewExchangeRate(CERC20(a)));
    }
}