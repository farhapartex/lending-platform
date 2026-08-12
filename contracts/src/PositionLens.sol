// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

import {ICollateralVault} from "./interfaces/ICollateralVault.sol";
import {IInterestRateModel} from "./interfaces/IInterestRateModel.sol";
import {ILendingController} from "./interfaces/ILendingController.sol";
import {ILendingPool} from "./interfaces/ILendingPool.sol";
import {IPositionLens, AccountData, MarketData} from "./interfaces/IPositionLens.sol";
import {IPriceOracle, PriceData} from "./interfaces/IPriceOracle.sol";
import {Errors} from "./libraries/Errors.sol";
import {HealthMath} from "./libraries/HealthMath.sol";
import {ShareMath} from "./libraries/ShareMath.sol";
import {WadMath} from "./libraries/WadMath.sol";

contract PositionLens is IPositionLens {
    ILendingPool public immutable pool;
    ICollateralVault public immutable vault;
    IPriceOracle public immutable oracle;
    ILendingController public immutable controller;

    address public immutable collateralAsset;
    address public immutable debtAsset;

    uint8 public immutable collateralDecimals;
    uint8 public immutable debtDecimals;

    constructor(
        ILendingPool lendingPool,
        ICollateralVault collateralVault,
        IPriceOracle priceOracle,
        ILendingController lendingController
    ) {
        if (
            address(lendingPool) == address(0) || address(collateralVault) == address(0)
                || address(priceOracle) == address(0) || address(lendingController) == address(0)
        ) {
            revert Errors.ZeroAddress();
        }

        pool = lendingPool;
        vault = collateralVault;
        oracle = priceOracle;
        controller = lendingController;

        collateralAsset = collateralVault.collateralAsset();
        debtAsset = lendingPool.asset();

        collateralDecimals = IERC20Metadata(collateralAsset).decimals();
        debtDecimals = IERC20Metadata(debtAsset).decimals();
    }

    function accountData(address account) external view returns (AccountData memory data) {
        (uint256 supplyIndexNow, uint256 borrowIndexNow,) = pool.previewAccrual();

        data.supplyShares = pool.sharesOf(account);
        data.supplyAssets = ShareMath.assetsFromSharesDown(data.supplyShares, supplyIndexNow);
        data.collateralAmount = vault.collateralOf(account);
        data.debtAmount = ShareMath.assetsFromSharesUp(pool.debtSharesOf(account), borrowIndexNow);

        PriceData memory collateralFeed = oracle.readPrice(collateralAsset);
        PriceData memory debtFeed = oracle.readPrice(debtAsset);

        data.collateralPrice = collateralFeed.price;
        data.priceUpdatedAt = collateralFeed.updatedAt;
        data.priceStale = !collateralFeed.isValid || !debtFeed.isValid;

        if (data.debtAmount == 0) {
            data.healthFactorBps = HealthMath.NO_DEBT_HEALTH_FACTOR;
            data.maxWithdrawableCollateral = data.collateralAmount;
        }

        if (data.priceStale) {
            return data;
        }

        _fillPricedFields(data, collateralFeed.price, debtFeed.price);

        data.maxBorrowable =
            WadMath.smaller(data.maxBorrowable, _availableLiquidity(supplyIndexNow, borrowIndexNow));

        return data;
    }

    function marketData() external view returns (MarketData memory data) {
        (uint256 supplyIndexNow, uint256 borrowIndexNow, uint256 reservesToAdd) = pool.previewAccrual();

        data.supplyIndex = supplyIndexNow;
        data.borrowIndex = borrowIndexNow;
        data.totalSupplied = ShareMath.assetsFromSharesDown(pool.totalSupplyShares(), supplyIndexNow);
        data.totalBorrowed = ShareMath.assetsFromSharesUp(pool.totalDebtShares(), borrowIndexNow);
        data.availableLiquidity = WadMath.subtractOrZero(data.totalSupplied, data.totalBorrowed);
        data.accruedReserves = pool.accruedReserves() + reservesToAdd;

        _fillRateFields(data);
        _fillRiskFields(data);

        return data;
    }

    function _fillPricedFields(AccountData memory data, uint256 collateralPrice, uint256 debtPrice)
        private
        view
    {
        data.collateralValue =
            HealthMath.valueOfAmount(data.collateralAmount, collateralDecimals, collateralPrice);
        data.debtValue = HealthMath.valueOfAmount(data.debtAmount, debtDecimals, debtPrice);

        uint256 maxLtvBps = controller.maxLtvBps();

        data.healthFactorBps = HealthMath.healthFactor(
            data.collateralValue, data.debtValue, controller.liquidationThresholdBps()
        );
        data.isLiquidatable = HealthMath.isLiquidatable(data.healthFactorBps);

        data.maxBorrowable = HealthMath.remainingBorrowAmount(
            data.collateralValue, data.debtValue, maxLtvBps, debtDecimals, debtPrice
        );

        data.maxWithdrawableCollateral = HealthMath.freeCollateralAmount(
            data.collateralAmount, collateralDecimals, collateralPrice, data.debtValue, maxLtvBps
        );
    }

    function _fillRateFields(MarketData memory data) private view {
        IInterestRateModel model = pool.rateModel();
        uint256 reserveFactorBps = pool.reserveFactorBps();

        data.utilizationBps = model.utilizationBps(data.totalSupplied, data.totalBorrowed);
        data.borrowRatePerSecond = model.borrowRatePerSecond(data.utilizationBps);
        data.supplyRatePerSecond = model.supplyRatePerSecond(data.utilizationBps, reserveFactorBps);
        data.borrowAprBps = model.borrowAprBps(data.utilizationBps);
        data.supplyAprBps = model.supplyAprBps(data.utilizationBps, reserveFactorBps);
        data.kinkUtilizationBps = model.curve().kinkUtilizationBps;
        data.reserveFactorBps = reserveFactorBps;
    }

    function _fillRiskFields(MarketData memory data) private view {
        data.maxLtvBps = controller.maxLtvBps();
        data.liquidationThresholdBps = controller.liquidationThresholdBps();
        data.liquidationBonusBps = controller.liquidationBonusBps();

        data.minDeposit = pool.minDeposit();
        data.depositsPaused = pool.depositsPaused();
        data.borrowPaused = controller.borrowPaused();
    }

    function _availableLiquidity(uint256 supplyIndexNow, uint256 borrowIndexNow)
        private
        view
        returns (uint256)
    {
        return WadMath.subtractOrZero(
            ShareMath.assetsFromSharesDown(pool.totalSupplyShares(), supplyIndexNow),
            ShareMath.assetsFromSharesUp(pool.totalDebtShares(), borrowIndexNow)
        );
    }
}
