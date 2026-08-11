// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface ILendingController {
    function borrow(uint256 amount) external;

    function repay(uint256 amount) external;

    function repayAll() external returns (uint256 amountPaid);

    function debtOf(address borrower) external view returns (uint256);

    function collateralValueOf(address borrower) external view returns (uint256);

    function debtValueOf(address borrower) external view returns (uint256);

    function healthFactorBps(address borrower) external view returns (uint256);

    function maxBorrowable(address borrower) external view returns (uint256);

    function maxWithdrawableCollateral(address borrower) external view returns (uint256);

    function isLiquidatable(address borrower) external view returns (bool);

    function maxLtvBps() external view returns (uint16);

    function liquidationThresholdBps() external view returns (uint16);

    function liquidationBonusBps() external view returns (uint16);
}
