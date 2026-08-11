// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface ILendingPool {
    function asset() external view returns (address);

    function deposit(uint256 assets) external returns (uint256 shares);

    function withdraw(uint256 assets) external returns (uint256 shares);

    function redeemShares(uint256 shares) external returns (uint256 assets);

    function supplyIndex() external view returns (uint256);

    function totalSupplied() external view returns (uint256);

    function totalBorrowed() external view returns (uint256);

    function borrowIndex() external view returns (uint256);

    function accrueInterest() external;

    function debtOf(address borrower) external view returns (uint256);

    function borrowFor(address borrower, address recipient, uint256 amount) external;

    function repayFor(address borrower, address payer, uint256 amount) external;

    function repayAllFor(address borrower, address payer) external returns (uint256 amountPaid);

    function availableLiquidity() external view returns (uint256);

    function sharesOf(address lender) external view returns (uint256);

    function balanceOfAssets(address lender) external view returns (uint256);

    function maxWithdrawable(address lender) external view returns (uint256);

    function minDeposit() external view returns (uint256);
}
