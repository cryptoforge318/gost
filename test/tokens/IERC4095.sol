// SPDX-License-Identifier: MIT
// Code adapted from https://github.com/OpenZeppelin/openzeppelin-contracts/pull/2237/
pragma solidity ^0.8.0;


interface IERC4095 {

    event Redeem(address indexed from, address indexed to, uint256 amount);
    event Mature(uint256 timestamp);

    function maturity() external view returns (uint256);
    function underlying() external view returns (address);
    function maturityRate() external view returns (uint256);
    function convertToUnderlying(uint256 principalAmount) external view returns (uint256);
    function convertToPrincipal(uint256 underlyingAmount) external view returns (uint256);
    function maxRedeem(address owner) external view returns (uint256);
    function previewRedeem(uint256 principalAmount) external view returns (uint256);
    function maxWithdraw(address owner) external view returns (uint256);
    function previewWithdraw(uint256 principalAmount) external view returns (uint256);  

    function withdraw(uint256 underlyingAmount, address owner, address receiver) external returns (uint256);
    function redeem(uint256 principalAmount, address owner, address receiver) external returns (uint256);
}