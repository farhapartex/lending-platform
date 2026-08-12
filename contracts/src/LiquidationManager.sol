// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

import {ICollateralVault} from "./interfaces/ICollateralVault.sol";
import {ILendingController} from "./interfaces/ILendingController.sol";
import {ILendingPool} from "./interfaces/ILendingPool.sol";
import {ILiquidationManager} from "./interfaces/ILiquidationManager.sol";
import {IPriceOracle} from "./interfaces/IPriceOracle.sol";
import {Errors} from "./libraries/Errors.sol";
import {HealthMath} from "./libraries/HealthMath.sol";
import {LiquidationMath} from "./libraries/LiquidationMath.sol";

contract LiquidationManager is ILiquidationManager, ReentrancyGuard {
    event LiquidationExecuted(
        address indexed borrower,
        address indexed liquidator,
        uint256 debtRepaid,
        uint256 collateralSeized,
        uint256 bonusValue,
        uint256 healthFactorBeforeBps,
        uint256 collateralPrice,
        uint8 priceDecimals,
        uint256 shortfall
    );

    struct LiquidationPlan {
        uint256 debtToRepay;
        uint256 collateralToSeize;
        uint256 bonusValue;
        uint256 shortfall;
        uint256 healthFactorBps;
        uint256 collateralPrice;
        uint8 priceDecimals;
    }

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

    function isLiquidatable(address borrower) external view returns (bool) {
        return HealthMath.isLiquidatable(_planFor(borrower).healthFactorBps);
    }

    function previewLiquidation(address borrower)
        external
        view
        returns (uint256 debtToRepay, uint256 collateralToSeize, uint256 bonusValue, uint256 shortfall)
    {
        LiquidationPlan memory plan = _planFor(borrower);

        return (plan.debtToRepay, plan.collateralToSeize, plan.bonusValue, plan.shortfall);
    }

    function liquidate(address borrower) external nonReentrant {
        pool.accrueInterest();

        LiquidationPlan memory plan = _planFor(borrower);

        if (!HealthMath.isLiquidatable(plan.healthFactorBps)) {
            revert Errors.PositionIsHealthy(plan.healthFactorBps);
        }

        uint256 amountRepaid = pool.repayAllFor(borrower, msg.sender);

        vault.seize(borrower, msg.sender, plan.collateralToSeize);

        emit LiquidationExecuted(
            borrower,
            msg.sender,
            amountRepaid,
            plan.collateralToSeize,
            plan.bonusValue,
            plan.healthFactorBps,
            plan.collateralPrice,
            plan.priceDecimals,
            plan.shortfall
        );
    }

    function _planFor(address borrower) private view returns (LiquidationPlan memory plan) {
        (uint256 collateralPrice, uint8 priceDecimals) = oracle.getPrice(collateralAsset);
        (uint256 debtPrice,) = oracle.getPrice(debtAsset);

        uint256 owed = pool.debtOf(borrower);
        uint256 held = vault.collateralOf(borrower);

        uint256 debtValue = HealthMath.valueOfAmount(owed, debtDecimals, debtPrice);
        uint256 collateralValue = HealthMath.valueOfAmount(held, collateralDecimals, collateralPrice);

        LiquidationMath.SeizePlan memory seizePlan = LiquidationMath.planSeize(
            debtValue,
            collateralValue,
            held,
            collateralDecimals,
            collateralPrice,
            controller.liquidationBonusBps()
        );

        plan.debtToRepay = owed;
        plan.collateralToSeize = seizePlan.collateralToSeize;
        plan.bonusValue = seizePlan.bonusValue;
        plan.shortfall = seizePlan.shortfall;
        plan.healthFactorBps =
            HealthMath.healthFactor(collateralValue, debtValue, controller.liquidationThresholdBps());
        plan.collateralPrice = collateralPrice;
        plan.priceDecimals = priceDecimals;

        return plan;
    }
}
