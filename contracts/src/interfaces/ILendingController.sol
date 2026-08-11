// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface ILendingController {
    function maxWithdrawableCollateral(address borrower) external view returns (uint256);
}
