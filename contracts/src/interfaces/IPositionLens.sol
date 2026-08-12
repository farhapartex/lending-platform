// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

struct AccountData {
    uint256 supplyShares;
    uint256 supplyAssets;
    uint256 collateralAmount;
    uint256 collateralValue;
    uint256 debtAmount;
    uint256 debtValue;
    uint256 healthFactorBps;
    uint256 maxBorrowable;
    uint256 maxWithdrawableCollateral;
    uint256 collateralPrice;
    uint256 priceUpdatedAt;
    bool isLiquidatable;
    bool priceStale;
}

struct MarketData {
    uint256 totalSupplied;
    uint256 totalBorrowed;
    uint256 availableLiquidity;
    uint256 utilizationBps;
    uint256 supplyRatePerSecond;
    uint256 borrowRatePerSecond;
    uint256 supplyAprBps;
    uint256 borrowAprBps;
    uint256 supplyIndex;
    uint256 borrowIndex;
    uint256 maxLtvBps;
    uint256 liquidationThresholdBps;
    uint256 liquidationBonusBps;
    uint256 kinkUtilizationBps;
    uint256 reserveFactorBps;
    uint256 minDeposit;
    uint256 accruedReserves;
    bool depositsPaused;
    bool borrowPaused;
}

interface IPositionLens {
    function accountData(address account) external view returns (AccountData memory);

    function marketData() external view returns (MarketData memory);
}
