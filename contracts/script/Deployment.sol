// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

import {CollateralVault} from "../src/CollateralVault.sol";
import {InterestRateModel} from "../src/InterestRateModel.sol";
import {LendingController} from "../src/LendingController.sol";
import {LendingPool} from "../src/LendingPool.sol";
import {LiquidationManager} from "../src/LiquidationManager.sol";
import {PositionLens} from "../src/PositionLens.sol";
import {PriceOracleAdapter} from "../src/PriceOracleAdapter.sol";
import {IAggregatorV3} from "../src/interfaces/IAggregatorV3.sol";
import {ICollateralVault} from "../src/interfaces/ICollateralVault.sol";
import {IInterestRateModel, RateCurve} from "../src/interfaces/IInterestRateModel.sol";
import {ILendingController} from "../src/interfaces/ILendingController.sol";
import {ILendingPool} from "../src/interfaces/ILendingPool.sol";
import {IPriceOracle} from "../src/interfaces/IPriceOracle.sol";
import {WadMath} from "../src/libraries/WadMath.sol";

struct DeployConfig {
    address finalOwner;
    address collateralToken;
    address debtToken;
    address collateralFeed;
    address debtFeed;
    uint32 maxPriceAge;
    uint256 minDeposit;
    uint16 reserveFactorBps;
    uint16 maxLtvBps;
    uint16 liquidationThresholdBps;
    uint16 liquidationBonusBps;
    uint16 baseRateAprBps;
    uint16 slopeBelowKinkAprBps;
    uint32 slopeAboveKinkAprBps;
    uint16 kinkUtilizationBps;
}

struct Deployment {
    PriceOracleAdapter oracle;
    InterestRateModel rateModel;
    LendingPool pool;
    CollateralVault vault;
    LendingController controller;
    LiquidationManager manager;
    PositionLens lens;
}

library ProtocolDeployer {
    uint8 internal constant EXPECTED_FEED_DECIMALS = 8;

    error FeedDecimalsNotSupported(address feed, uint8 decimals);
    error WiringCheckFailed(string what);

    function aprBpsToPerSecond(uint256 aprBps) internal pure returns (uint64) {
        return
            SafeCast.toUint64((aprBps * WadMath.WAD) / (WadMath.FULL_PERCENT_BPS * WadMath.SECONDS_PER_YEAR));
    }

    function buildCurve(DeployConfig memory config) internal pure returns (RateCurve memory) {
        return RateCurve({
            baseRatePerSecond: aprBpsToPerSecond(config.baseRateAprBps),
            slopeBelowKinkPerSecond: aprBpsToPerSecond(config.slopeBelowKinkAprBps),
            slopeAboveKinkPerSecond: aprBpsToPerSecond(config.slopeAboveKinkAprBps),
            kinkUtilizationBps: config.kinkUtilizationBps
        });
    }

    function requireEightDecimalFeeds(DeployConfig memory config) internal view {
        uint8 collateralFeedDecimals = IAggregatorV3(config.collateralFeed).decimals();
        if (collateralFeedDecimals != EXPECTED_FEED_DECIMALS) {
            revert FeedDecimalsNotSupported(config.collateralFeed, collateralFeedDecimals);
        }

        uint8 debtFeedDecimals = IAggregatorV3(config.debtFeed).decimals();
        if (debtFeedDecimals != EXPECTED_FEED_DECIMALS) {
            revert FeedDecimalsNotSupported(config.debtFeed, debtFeedDecimals);
        }
    }

    function deployAll(DeployConfig memory config, address initialOwner)
        internal
        returns (Deployment memory deployed)
    {
        deployed.oracle = new PriceOracleAdapter(initialOwner, config.maxPriceAge);
        deployed.rateModel = new InterestRateModel(initialOwner, buildCurve(config));

        deployed.pool = new LendingPool(
            initialOwner,
            IERC20(config.debtToken),
            IInterestRateModel(address(deployed.rateModel)),
            config.minDeposit,
            config.reserveFactorBps
        );

        deployed.vault = new CollateralVault(initialOwner, IERC20(config.collateralToken));

        deployed.controller = new LendingController(
            initialOwner,
            ILendingPool(address(deployed.pool)),
            ICollateralVault(address(deployed.vault)),
            IPriceOracle(address(deployed.oracle)),
            config.maxLtvBps,
            config.liquidationThresholdBps,
            config.liquidationBonusBps
        );

        deployed.manager = new LiquidationManager(
            ILendingPool(address(deployed.pool)),
            ICollateralVault(address(deployed.vault)),
            IPriceOracle(address(deployed.oracle)),
            ILendingController(address(deployed.controller))
        );

        deployed.lens = new PositionLens(
            ILendingPool(address(deployed.pool)),
            ICollateralVault(address(deployed.vault)),
            IPriceOracle(address(deployed.oracle)),
            ILendingController(address(deployed.controller))
        );

        return deployed;
    }

    function wireAll(Deployment memory deployed, DeployConfig memory config) internal {
        deployed.oracle.setFeed(config.collateralToken, config.collateralFeed);
        deployed.oracle.setFeed(config.debtToken, config.debtFeed);

        deployed.pool.linkController(address(deployed.controller));
        deployed.pool.linkLiquidationManager(address(deployed.manager));

        deployed.vault.linkController(ILendingController(address(deployed.controller)));
        deployed.vault.linkLiquidationManager(address(deployed.manager));
    }

    function handOverOwnership(Deployment memory deployed, address newOwner) internal {
        deployed.oracle.transferOwnership(newOwner);
        deployed.rateModel.transferOwnership(newOwner);
        deployed.pool.transferOwnership(newOwner);
        deployed.vault.transferOwnership(newOwner);
        deployed.controller.transferOwnership(newOwner);
    }

    function verifyWiring(Deployment memory deployed, DeployConfig memory config) internal view {
        _expect(deployed.oracle.feedOf(config.collateralToken) == config.collateralFeed, "collateral feed");
        _expect(deployed.oracle.feedOf(config.debtToken) == config.debtFeed, "debt feed");

        _expect(deployed.pool.controller() == address(deployed.controller), "pool controller");
        _expect(deployed.pool.liquidationManager() == address(deployed.manager), "pool liquidation manager");

        _expect(address(deployed.vault.controller()) == address(deployed.controller), "vault controller");
        _expect(deployed.vault.liquidationManager() == address(deployed.manager), "vault liquidation manager");

        _expect(deployed.pool.asset() == config.debtToken, "pool asset");
        _expect(deployed.vault.collateralAsset() == config.collateralToken, "vault asset");

        _expect(deployed.controller.maxLtvBps() == config.maxLtvBps, "max ltv");
        _expect(
            deployed.controller.liquidationThresholdBps() == config.liquidationThresholdBps,
            "liquidation threshold"
        );
        _expect(deployed.controller.liquidationBonusBps() == config.liquidationBonusBps, "liquidation bonus");

        _expect(deployed.lens.collateralAsset() == config.collateralToken, "lens collateral asset");
        _expect(deployed.lens.debtAsset() == config.debtToken, "lens debt asset");

        _expect(deployed.manager.collateralAsset() == config.collateralToken, "manager collateral asset");
        _expect(deployed.manager.debtAsset() == config.debtToken, "manager debt asset");
    }

    function verifyOwnership(Deployment memory deployed, address expectedOwner) internal view {
        _expect(deployed.oracle.owner() == expectedOwner, "oracle owner");
        _expect(deployed.rateModel.owner() == expectedOwner, "rate model owner");
        _expect(deployed.pool.owner() == expectedOwner, "pool owner");
        _expect(deployed.vault.owner() == expectedOwner, "vault owner");
        _expect(deployed.controller.owner() == expectedOwner, "controller owner");
    }

    function _expect(bool condition, string memory what) private pure {
        if (!condition) {
            revert WiringCheckFailed(what);
        }
    }
}
