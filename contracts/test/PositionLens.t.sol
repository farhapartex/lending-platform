// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {BaseTest} from "./Base.t.sol";
import {PositionLens} from "../src/PositionLens.sol";
import {AccountData, MarketData} from "../src/interfaces/IPositionLens.sol";
import {ICollateralVault} from "../src/interfaces/ICollateralVault.sol";
import {ILendingController} from "../src/interfaces/ILendingController.sol";
import {ILendingPool} from "../src/interfaces/ILendingPool.sol";
import {IPriceOracle} from "../src/interfaces/IPriceOracle.sol";
import {Errors} from "../src/libraries/Errors.sol";

contract PositionLensTest is BaseTest {
    function test_constructor_rejects_zero_dependencies() public {
        vm.expectRevert(Errors.ZeroAddress.selector);
        new PositionLens(
            ILendingPool(address(0)),
            ICollateralVault(address(vault)),
            IPriceOracle(address(oracle)),
            ILendingController(address(controller))
        );
    }

    function test_empty_account_reads_cleanly() public view {
        AccountData memory data = lens.accountData(stranger);

        assertEq(data.supplyShares, 0);
        assertEq(data.supplyAssets, 0);
        assertEq(data.collateralAmount, 0);
        assertEq(data.debtAmount, 0);
        assertEq(data.healthFactorBps, type(uint256).max);
        assertFalse(data.isLiquidatable);
        assertFalse(data.priceStale);
    }

    function test_accountData_matches_controller() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        AccountData memory data = lens.accountData(alice);

        assertEq(data.collateralAmount, 3.2e18);
        assertEq(data.collateralValue, controller.collateralValueOf(alice));
        assertEq(data.debtAmount, controller.debtOf(alice));
        assertEq(data.debtValue, controller.debtValueOf(alice));
        assertEq(data.healthFactorBps, controller.healthFactorBps(alice));
        assertEq(data.maxBorrowable, controller.maxBorrowable(alice));
        assertEq(data.maxWithdrawableCollateral, controller.maxWithdrawableCollateral(alice));
        assertEq(data.isLiquidatable, manager.isLiquidatable(alice));
        assertEq(data.collateralPrice, uint256(ETH_PRICE));
    }

    function test_lender_view_matches_pool() public {
        _deposit(lender, 50_000e6);

        AccountData memory data = lens.accountData(lender);

        assertEq(data.supplyShares, pool.sharesOf(lender));
        assertEq(data.supplyAssets, pool.balanceOfAssets(lender));
    }

    function test_accountData_projects_pending_interest() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);
        _skip(YEAR);

        AccountData memory data = lens.accountData(alice);

        assertEq(data.debtAmount, 7046842350);
        assertEq(lens.accountData(lender).supplyAssets, 50132158115);

        pool.accrueInterest();

        assertEq(pool.debtOf(alice), data.debtAmount);
    }

    function test_stale_price_flags_instead_of_reverting() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        vm.warp(block.timestamp + MAX_PRICE_AGE + 1);

        AccountData memory data = lens.accountData(alice);

        assertTrue(data.priceStale);
        assertEq(data.collateralValue, 0);
        assertEq(data.debtValue, 0);
        assertEq(data.healthFactorBps, 0);
        assertEq(data.maxBorrowable, 0);
        assertFalse(data.isLiquidatable);

        assertGt(data.debtAmount, 0);
        assertEq(data.collateralAmount, 3.2e18);

        vm.expectRevert();
        controller.healthFactorBps(alice);
    }

    function test_broken_feed_flags_instead_of_reverting() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        wethFeed.setFeedDown(true);

        AccountData memory data = lens.accountData(alice);

        assertTrue(data.priceStale);
        assertEq(data.collateralPrice, 0);
    }

    function test_debt_free_user_sees_full_collateral_during_outage() public {
        _addCollateral(alice, 3.2e18);

        wethFeed.setFeedDown(true);

        AccountData memory data = lens.accountData(alice);

        assertTrue(data.priceStale);
        assertEq(data.debtAmount, 0);
        assertEq(data.healthFactorBps, type(uint256).max);
        assertEq(data.maxWithdrawableCollateral, 3.2e18);
    }

    function test_marketData_reports_the_whole_market() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);

        MarketData memory market = lens.marketData();

        assertEq(market.totalSupplied, 50_000e6);
        assertEq(market.totalBorrowed, 6_900e6);
        assertEq(market.availableLiquidity, 43_100e6);
        assertEq(market.utilizationBps, 1380);
        assertEq(market.supplyIndex, 1e18);
        assertEq(market.borrowIndex, 1e18);
        assertEq(market.maxLtvBps, MAX_LTV_BPS);
        assertEq(market.liquidationThresholdBps, LIQUIDATION_THRESHOLD_BPS);
        assertEq(market.liquidationBonusBps, LIQUIDATION_BONUS_BPS);
        assertEq(market.kinkUtilizationBps, KINK_BPS);
        assertEq(market.reserveFactorBps, RESERVE_FACTOR_BPS);
        assertEq(market.minDeposit, MIN_DEPOSIT);
        assertEq(market.accruedReserves, 0);
        assertFalse(market.depositsPaused);
        assertFalse(market.borrowPaused);
        assertGt(market.borrowAprBps, market.supplyAprBps);
    }

    function test_marketData_reflects_pause_flags() public {
        vm.startPrank(admin);
        pool.setDepositsPaused(true);
        controller.setBorrowPaused(true);
        vm.stopPrank();

        MarketData memory market = lens.marketData();

        assertTrue(market.depositsPaused);
        assertTrue(market.borrowPaused);
    }

    function test_marketData_survives_a_broken_feed() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);

        wethFeed.setFeedDown(true);

        MarketData memory market = lens.marketData();

        assertEq(market.totalSupplied, 50_000e6);
        assertEq(market.totalBorrowed, 6_900e6);
    }

    function test_marketData_projects_reserves() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);
        _skip(YEAR);

        MarketData memory market = lens.marketData();

        assertEq(market.accruedReserves, 14684235);
        assertEq(market.totalBorrowed, 7046842350);

        pool.accrueInterest();

        assertEq(pool.accruedReserves(), market.accruedReserves);
    }

    function testFuzz_lens_always_agrees_with_controller(uint256 collateral, uint256 portionBps) public {
        collateral = bound(collateral, 1e17, 100e18);
        portionBps = bound(portionBps, 1, 10_000);

        _deposit(lender, 500_000e6);
        _addCollateral(alice, collateral);

        uint256 room = controller.maxBorrowable(alice);
        uint256 amount = (room * portionBps) / 10_000;

        if (amount > 0) {
            _borrow(alice, amount);
        }

        AccountData memory data = lens.accountData(alice);

        assertEq(data.debtAmount, controller.debtOf(alice));
        assertEq(data.healthFactorBps, controller.healthFactorBps(alice));
        assertEq(data.maxBorrowable, controller.maxBorrowable(alice));
        assertEq(data.maxWithdrawableCollateral, controller.maxWithdrawableCollateral(alice));
        assertEq(data.isLiquidatable, manager.isLiquidatable(alice));
    }

    function testFuzz_accountData_never_reverts(uint256 elapsed, bool feedDown, int256 price) public {
        elapsed = bound(elapsed, 0, 5 * YEAR);
        price = bound(price, type(int64).min, type(int64).max);

        _deposit(lender, 100_000e6);
        _openPosition(alice, 10e18, 5_000e6);

        vm.warp(block.timestamp + elapsed);
        wethFeed.setPrice(price);
        wethFeed.setFeedDown(feedDown);

        AccountData memory data = lens.accountData(alice);
        MarketData memory market = lens.marketData();

        assertGe(data.collateralAmount, 0);
        assertGe(market.totalSupplied, 0);
    }
}
