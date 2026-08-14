// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

library Errors {
    error ZeroAmount();

    error ZeroAddress();

    error AlreadyInitialized();

    error NotAuthorized(address caller);

    error MarketPaused();

    error BelowMinimumDeposit(uint256 provided, uint256 minimum);

    error ExceedsSupplyBalance(uint256 requested, uint256 available);

    error NotEnoughLiquidity(uint256 requested, uint256 available);

    error ExceedsBorrowLimit(uint256 requested, uint256 allowed);

    error ExceedsDebt(uint256 requested, uint256 owed);

    error ExceedsReserves(uint256 requested, uint256 available);

    error ExceedsCollateralBalance(uint256 requested, uint256 held);

    error WouldBreakBorrowLimit(uint256 requested, uint256 safeAmount);

    error PositionIsHealthy(uint256 healthFactorBps);

    error PriceIsStale(address asset, uint256 updatedAt, uint256 maxAge);

    error PriceIsInvalid(address asset);

    error UnsupportedFeedDecimals(address aggregator, uint8 provided, uint8 expected);

    error InvalidRiskSettings();
}
