// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {CollateralVault} from "../src/CollateralVault.sol";
import {InterestRateModel} from "../src/InterestRateModel.sol";
import {LendingController} from "../src/LendingController.sol";
import {LendingPool} from "../src/LendingPool.sol";
import {LiquidationManager} from "../src/LiquidationManager.sol";
import {PositionLens} from "../src/PositionLens.sol";
import {PriceOracleAdapter} from "../src/PriceOracleAdapter.sol";
import {MockAggregator} from "../src/mocks/MockAggregator.sol";
import {MockERC20} from "../src/mocks/MockERC20.sol";
import {ICollateralVault} from "../src/interfaces/ICollateralVault.sol";
import {IInterestRateModel, RateCurve} from "../src/interfaces/IInterestRateModel.sol";
import {ILendingController} from "../src/interfaces/ILendingController.sol";
import {ILendingPool} from "../src/interfaces/ILendingPool.sol";
import {IPriceOracle} from "../src/interfaces/IPriceOracle.sol";

abstract contract BaseTest is Test {
    uint256 internal constant YEAR = 365 days;

    uint16 internal constant MAX_LTV_BPS = 7500;
    uint16 internal constant LIQUIDATION_THRESHOLD_BPS = 8000;
    uint16 internal constant LIQUIDATION_BONUS_BPS = 500;
    uint16 internal constant RESERVE_FACTOR_BPS = 1000;
    uint16 internal constant KINK_BPS = 8000;
    uint32 internal constant MAX_PRICE_AGE = 3600;
    uint256 internal constant MIN_DEPOSIT = 1e6;

    int256 internal constant ETH_PRICE = 341258000000;
    int256 internal constant USDC_PRICE = 100000000;

    MockERC20 internal weth;
    MockERC20 internal usdc;
    MockAggregator internal wethFeed;
    MockAggregator internal usdcFeed;

    PriceOracleAdapter internal oracle;
    InterestRateModel internal rateModel;
    LendingPool internal pool;
    CollateralVault internal vault;
    LendingController internal controller;
    LiquidationManager internal manager;
    PositionLens internal lens;

    address internal admin = makeAddr("admin");
    address internal treasury = makeAddr("treasury");
    address internal lender = makeAddr("lender");
    address internal lenderTwo = makeAddr("lenderTwo");
    address internal alice = makeAddr("alice");
    address internal bob = makeAddr("bob");
    address internal liquidator = makeAddr("liquidator");
    address internal stranger = makeAddr("stranger");

    function setUp() public virtual {
        vm.warp(1_700_000_000);

        weth = new MockERC20("Wrapped Ether", "WETH", 18);
        usdc = new MockERC20("USD Coin", "USDC", 6);
        wethFeed = new MockAggregator("ETH / USD", 8, ETH_PRICE);
        usdcFeed = new MockAggregator("USDC / USD", 8, USDC_PRICE);

        oracle = new PriceOracleAdapter(admin, MAX_PRICE_AGE);
        rateModel = new InterestRateModel(admin, _defaultCurve());

        pool = new LendingPool(
            admin,
            IERC20(address(usdc)),
            IInterestRateModel(address(rateModel)),
            MIN_DEPOSIT,
            RESERVE_FACTOR_BPS
        );

        vault = new CollateralVault(admin, IERC20(address(weth)));

        controller = new LendingController(
            admin,
            ILendingPool(address(pool)),
            ICollateralVault(address(vault)),
            IPriceOracle(address(oracle)),
            MAX_LTV_BPS,
            LIQUIDATION_THRESHOLD_BPS,
            LIQUIDATION_BONUS_BPS
        );

        manager = new LiquidationManager(
            ILendingPool(address(pool)),
            ICollateralVault(address(vault)),
            IPriceOracle(address(oracle)),
            ILendingController(address(controller))
        );

        lens = new PositionLens(
            ILendingPool(address(pool)),
            ICollateralVault(address(vault)),
            IPriceOracle(address(oracle)),
            ILendingController(address(controller))
        );

        vm.startPrank(admin);
        oracle.setFeed(address(weth), address(wethFeed));
        oracle.setFeed(address(usdc), address(usdcFeed));
        pool.linkController(address(controller));
        pool.linkLiquidationManager(address(manager));
        vault.linkController(ILendingController(address(controller)));
        vault.linkLiquidationManager(address(manager));
        vm.stopPrank();

        _fund(lender, 0, 1_000_000e6);
        _fund(lenderTwo, 0, 1_000_000e6);
        _fund(alice, 1_000e18, 100_000e6);
        _fund(bob, 1_000e18, 100_000e6);
        _fund(liquidator, 0, 1_000_000e6);
    }

    function _defaultCurve() internal pure returns (RateCurve memory) {
        return RateCurve({
            baseRatePerSecond: uint64(uint256(1e16) / YEAR),
            slopeBelowKinkPerSecond: uint64(uint256(654e14) / YEAR),
            slopeAboveKinkPerSecond: uint64(uint256(2e18) / YEAR),
            kinkUtilizationBps: KINK_BPS
        });
    }

    function _fund(address who, uint256 wethAmount, uint256 usdcAmount) internal {
        if (wethAmount > 0) {
            weth.mint(who, wethAmount);
        }

        if (usdcAmount > 0) {
            usdc.mint(who, usdcAmount);
        }

        vm.startPrank(who);
        weth.approve(address(vault), type(uint256).max);
        usdc.approve(address(pool), type(uint256).max);
        vm.stopPrank();
    }

    function _deposit(address who, uint256 amount) internal returns (uint256 shares) {
        vm.prank(who);
        return pool.deposit(amount);
    }

    function _addCollateral(address who, uint256 amount) internal {
        vm.prank(who);
        vault.depositCollateral(amount);
    }

    function _borrow(address who, uint256 amount) internal {
        vm.prank(who);
        controller.borrow(amount);
    }

    function _openPosition(address who, uint256 collateral, uint256 debt) internal {
        _addCollateral(who, collateral);

        if (debt > 0) {
            _borrow(who, debt);
        }
    }

    function _setEthPrice(int256 price) internal {
        wethFeed.setPrice(price);
    }

    function _refreshFeeds() internal {
        (, int256 ethAnswer,,,) = wethFeed.latestRoundData();
        (, int256 usdcAnswer,,,) = usdcFeed.latestRoundData();

        wethFeed.setPrice(ethAnswer);
        usdcFeed.setPrice(usdcAnswer);
    }

    function _skip(uint256 secondsForward) internal {
        vm.warp(block.timestamp + secondsForward);
        _refreshFeeds();
    }

    function _poolCash() internal view returns (uint256) {
        return usdc.balanceOf(address(pool));
    }
}
