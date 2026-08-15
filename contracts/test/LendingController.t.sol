// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

import {BaseTest} from "./Base.t.sol";
import {LendingController} from "../src/LendingController.sol";
import {ICollateralVault} from "../src/interfaces/ICollateralVault.sol";
import {ILendingPool} from "../src/interfaces/ILendingPool.sol";
import {IPriceOracle} from "../src/interfaces/IPriceOracle.sol";
import {Errors} from "../src/libraries/Errors.sol";

contract LendingControllerTest is BaseTest {
    function test_constructor_rejects_zero_dependencies() public {
        vm.expectRevert(Errors.ZeroAddress.selector);
        new LendingController(
            admin,
            ILendingPool(address(0)),
            ICollateralVault(address(vault)),
            IPriceOracle(address(oracle)),
            7500,
            8000,
            500
        );
    }

    function test_constructor_enforces_ltv_below_threshold() public {
        vm.expectRevert(Errors.InvalidRiskSettings.selector);
        new LendingController(
            admin,
            ILendingPool(address(pool)),
            ICollateralVault(address(vault)),
            IPriceOracle(address(oracle)),
            8000,
            8000,
            500
        );

        vm.expectRevert(Errors.InvalidRiskSettings.selector);
        new LendingController(
            admin,
            ILendingPool(address(pool)),
            ICollateralVault(address(vault)),
            IPriceOracle(address(oracle)),
            0,
            8000,
            500
        );
    }

    function test_constructor_discovers_assets_and_decimals() public view {
        assertEq(controller.collateralAsset(), address(weth));
        assertEq(controller.debtAsset(), address(usdc));
        assertEq(controller.collateralDecimals(), 18);
        assertEq(controller.debtDecimals(), 6);
    }

    function test_setRiskSettings_validations() public {
        vm.startPrank(admin);

        vm.expectRevert(Errors.InvalidRiskSettings.selector);
        controller.setRiskSettings(8000, 8000, 500);

        vm.expectRevert(Errors.InvalidRiskSettings.selector);
        controller.setRiskSettings(7500, 10_001, 500);

        vm.expectRevert(Errors.InvalidRiskSettings.selector);
        controller.setRiskSettings(7500, 8000, 10_001);

        controller.setRiskSettings(7000, 7500, 400);
        vm.stopPrank();

        assertEq(controller.maxLtvBps(), 7000);
        assertEq(controller.liquidationThresholdBps(), 7500);
        assertEq(controller.liquidationBonusBps(), 400);
    }

    function test_admin_functions_are_owner_only() public {
        vm.startPrank(stranger);

        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, stranger));
        controller.setRiskSettings(7000, 7500, 400);

        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, stranger));
        controller.setBorrowPaused(true);

        vm.stopPrank();
    }

    function test_borrow_happy_path_matches_expected_figures() public {
        _deposit(lender, 50_000e6);
        _addCollateral(alice, 3.2e18);

        assertEq(controller.collateralValueOf(alice), 1092025600000);
        assertEq(controller.maxBorrowable(alice), 8190192000);

        _borrow(alice, 6_900e6);

        assertEq(controller.debtOf(alice), 6_900e6);
        assertEq(controller.healthFactorBps(alice), 12661);
        assertEq(controller.maxBorrowable(alice), 1290192000);
        assertEq(controller.maxWithdrawableCollateral(alice), 504092504791096472);
        assertEq(usdc.balanceOf(alice), 100_000e6 + 6_900e6);
    }

    function test_borrow_rejects_zero_and_pause() public {
        _deposit(lender, 50_000e6);
        _addCollateral(alice, 3.2e18);

        vm.prank(alice);
        vm.expectRevert(Errors.ZeroAmount.selector);
        controller.borrow(0);

        vm.prank(admin);
        controller.setBorrowPaused(true);

        vm.prank(alice);
        vm.expectRevert(Errors.MarketPaused.selector);
        controller.borrow(1e6);

        vm.prank(admin);
        controller.setBorrowPaused(false);

        _borrow(alice, 1e6);
        assertEq(controller.debtOf(alice), 1e6);
    }

    function test_collateral_shortfall_and_liquidity_shortfall_are_distinct() public {
        _deposit(lender, 50_000e6);
        _addCollateral(alice, 3.2e18);

        uint256 room = controller.maxBorrowable(alice);

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSelector(Errors.ExceedsBorrowLimit.selector, room + 1, room));
        controller.borrow(room + 1);

        _addCollateral(bob, 500e18);
        _borrow(bob, 45_000e6);

        uint256 liquidity = pool.availableLiquidity();

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSelector(Errors.NotEnoughLiquidity.selector, liquidity + 1, liquidity));
        controller.borrow(liquidity + 1);
    }

    function test_borrowing_the_maximum_is_never_liquidatable() public {
        _deposit(lender, 500_000e6);
        _addCollateral(alice, 10e18);

        _borrow(alice, controller.maxBorrowable(alice));

        assertFalse(manager.isLiquidatable(alice));
        assertGe(controller.healthFactorBps(alice), 10_000);
    }

    function test_repay_reduces_debt() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        vm.prank(alice);
        controller.repay(1_000e6);

        assertEq(controller.debtOf(alice), 5_900e6);
    }

    function test_repay_rejects_zero_and_overpayment() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        vm.prank(alice);
        vm.expectRevert(Errors.ZeroAmount.selector);
        controller.repay(0);

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSelector(Errors.ExceedsDebt.selector, 7_000e6, 6_900e6));
        controller.repay(7_000e6);
    }

    function test_partial_repay_leaves_dust_but_repayAll_clears_it() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);
        _skip(YEAR);

        pool.accrueInterest();
        assertEq(controller.debtOf(alice), 7046842350);

        vm.prank(alice);
        controller.repay(1_000e6);

        assertEq(controller.debtOf(alice), 6046842351);

        vm.prank(alice);
        uint256 paid = controller.repayAll();

        assertEq(paid, 6046842351);
        assertEq(controller.debtOf(alice), 0);
        assertEq(pool.debtSharesOf(alice), 0);
    }

    function test_repayAll_rejects_when_there_is_no_debt() public {
        vm.prank(alice);
        vm.expectRevert(Errors.ZeroAmount.selector);
        controller.repayAll();
    }

    function test_repay_works_while_the_oracle_is_down() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        wethFeed.setFeedDown(true);

        vm.expectRevert();
        controller.healthFactorBps(alice);

        vm.prank(alice);
        controller.repay(1_000e6);

        vm.prank(alice);
        controller.repayAll();

        assertEq(controller.debtOf(alice), 0);
    }

    function test_repay_works_while_the_price_is_stale() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        vm.warp(block.timestamp + MAX_PRICE_AGE + 1);

        vm.prank(alice);
        controller.repayAll();

        assertEq(controller.debtOf(alice), 0);
    }

    function test_borrow_is_blocked_while_the_price_is_stale() public {
        _deposit(lender, 50_000e6);
        _addCollateral(alice, 3.2e18);

        vm.warp(block.timestamp + MAX_PRICE_AGE + 1);

        vm.prank(alice);
        vm.expectRevert();
        controller.borrow(1_000e6);
    }

    function test_views_project_pending_interest() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);

        uint256 healthBefore = controller.healthFactorBps(alice);

        _skip(YEAR);

        assertEq(controller.debtOf(alice), 7046842350);
        assertLt(controller.healthFactorBps(alice), healthBefore);
        assertGt(pool.totalBorrowed(), 0);
    }

    function test_debt_free_user_can_withdraw_everything() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        vm.prank(alice);
        controller.repayAll();

        assertEq(controller.maxWithdrawableCollateral(alice), 3.2e18);

        vm.prank(alice);
        vault.withdrawCollateral(3.2e18);

        assertEq(vault.collateralOf(alice), 0);
    }

    function testFuzz_borrow_within_limit_always_leaves_position_safe(uint256 collateral, uint256 portionBps)
        public
    {
        collateral = bound(collateral, 1e17, 100e18);
        portionBps = bound(portionBps, 1, 10_000);

        _deposit(lender, 500_000e6);
        _addCollateral(alice, collateral);

        uint256 room = controller.maxBorrowable(alice);
        uint256 amount = (room * portionBps) / 10_000;

        if (amount == 0) {
            return;
        }

        _borrow(alice, amount);

        assertFalse(manager.isLiquidatable(alice));
    }

    function testFuzz_borrow_beyond_limit_always_rejected(uint256 collateral, uint256 excess) public {
        collateral = bound(collateral, 1e17, 100e18);
        excess = bound(excess, 1, 1_000e6);

        _deposit(lender, 500_000e6);
        _addCollateral(alice, collateral);

        uint256 room = controller.maxBorrowable(alice);

        vm.prank(alice);
        vm.expectRevert();
        controller.borrow(room + excess);
    }

    function testFuzz_repay_then_borrow_again_is_consistent(uint256 first, uint256 repayPortionBps) public {
        _deposit(lender, 500_000e6);
        _addCollateral(alice, 50e18);

        uint256 room = controller.maxBorrowable(alice);
        first = bound(first, 1e6, room);
        repayPortionBps = bound(repayPortionBps, 1, 10_000);

        _borrow(alice, first);

        uint256 repayAmount = (first * repayPortionBps) / 10_000;
        if (repayAmount == 0) {
            return;
        }

        vm.prank(alice);
        controller.repay(repayAmount);

        assertEq(controller.debtOf(alice), first - repayAmount);
    }

    function testFuzz_repayAll_always_clears_debt(uint256 borrowAmount, uint256 elapsed) public {
        _deposit(lender, 500_000e6);
        _addCollateral(alice, 50e18);

        uint256 room = controller.maxBorrowable(alice);
        borrowAmount = bound(borrowAmount, 1e6, room);
        elapsed = bound(elapsed, 0, 2 * YEAR);

        _borrow(alice, borrowAmount);
        _skip(elapsed);

        vm.prank(alice);
        controller.repayAll();

        assertEq(controller.debtOf(alice), 0);
        assertEq(pool.debtSharesOf(alice), 0);
    }
}
