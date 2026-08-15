// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {BaseTest} from "./Base.t.sol";
import {CollateralVault} from "../src/CollateralVault.sol";
import {ILendingController} from "../src/interfaces/ILendingController.sol";
import {Errors} from "../src/libraries/Errors.sol";

contract CollateralVaultTest is BaseTest {
    function test_constructor_rejects_zero_asset() public {
        vm.expectRevert(Errors.ZeroAddress.selector);
        new CollateralVault(admin, IERC20(address(0)));
    }

    function test_deposit_records_and_pulls_tokens() public {
        _addCollateral(alice, 3.2e18);

        assertEq(vault.collateralOf(alice), 3.2e18);
        assertEq(vault.totalCollateral(), 3.2e18);
        assertEq(weth.balanceOf(address(vault)), 3.2e18);
    }

    function test_deposit_rejects_zero() public {
        vm.prank(alice);
        vm.expectRevert(Errors.ZeroAmount.selector);
        vault.depositCollateral(0);
    }

    function test_deposit_works_while_the_oracle_is_down() public {
        wethFeed.setFeedDown(true);

        _addCollateral(alice, 1e18);

        assertEq(vault.collateralOf(alice), 1e18);
    }

    function test_deposit_works_while_the_price_is_stale() public {
        vm.warp(block.timestamp + MAX_PRICE_AGE + 1);

        _addCollateral(alice, 1e18);

        assertEq(vault.collateralOf(alice), 1e18);
    }

    function test_deposits_accumulate() public {
        _addCollateral(alice, 1e18);
        _addCollateral(alice, 2e18);

        assertEq(vault.collateralOf(alice), 3e18);
        assertEq(vault.totalCollateral(), 3e18);
    }

    function test_withdraw_without_debt_releases_everything() public {
        _addCollateral(alice, 3.2e18);

        vm.prank(alice);
        vault.withdrawCollateral(3.2e18);

        assertEq(vault.collateralOf(alice), 0);
        assertEq(vault.totalCollateral(), 0);
    }

    function test_withdraw_rejects_zero() public {
        _addCollateral(alice, 1e18);

        vm.prank(alice);
        vm.expectRevert(Errors.ZeroAmount.selector);
        vault.withdrawCollateral(0);
    }

    function test_withdraw_beyond_balance_is_its_own_error() public {
        _addCollateral(alice, 1e18);

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSelector(Errors.ExceedsCollateralBalance.selector, 2e18, 1e18));
        vault.withdrawCollateral(2e18);
    }

    function test_withdraw_blocked_when_it_would_break_the_borrow_limit() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        uint256 safeAmount = controller.maxWithdrawableCollateral(alice);

        vm.prank(alice);
        vm.expectRevert(
            abi.encodeWithSelector(Errors.WouldBreakBorrowLimit.selector, safeAmount + 1, safeAmount)
        );
        vault.withdrawCollateral(safeAmount + 1);

        vm.prank(alice);
        vault.withdrawCollateral(safeAmount);

        assertEq(vault.collateralOf(alice), 3.2e18 - safeAmount);
    }

    function test_withdrawing_the_safe_amount_keeps_position_healthy() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        uint256 safeAmount = controller.maxWithdrawableCollateral(alice);

        vm.prank(alice);
        vault.withdrawCollateral(safeAmount);

        assertFalse(manager.isLiquidatable(alice));
        assertGe(controller.healthFactorBps(alice), 10_000);
    }

    function test_withdraw_guard_tightens_as_interest_accrues() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        uint256 before = controller.maxWithdrawableCollateral(alice);

        _skip(YEAR);

        uint256 later = controller.maxWithdrawableCollateral(alice);

        assertLt(later, before);

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSelector(Errors.WouldBreakBorrowLimit.selector, before, later));
        vault.withdrawCollateral(before);
    }

    function test_withdraw_fails_closed_without_a_controller() public {
        CollateralVault fresh = new CollateralVault(admin, IERC20(address(weth)));

        vm.startPrank(alice);
        weth.approve(address(fresh), type(uint256).max);
        fresh.depositCollateral(1e18);

        vm.expectRevert(abi.encodeWithSelector(Errors.NotAuthorized.selector, address(0)));
        fresh.withdrawCollateral(1e18);
        vm.stopPrank();
    }

    function test_seize_is_liquidation_manager_only() public {
        _addCollateral(alice, 1e18);

        vm.prank(stranger);
        vm.expectRevert(abi.encodeWithSelector(Errors.NotAuthorized.selector, stranger));
        vault.seize(alice, stranger, 1e18);

        vm.prank(admin);
        vm.expectRevert(abi.encodeWithSelector(Errors.NotAuthorized.selector, admin));
        vault.seize(alice, admin, 1e18);

        vm.prank(address(controller));
        vm.expectRevert(abi.encodeWithSelector(Errors.NotAuthorized.selector, address(controller)));
        vault.seize(alice, address(controller), 1e18);
    }

    function test_seize_validations() public {
        _addCollateral(alice, 1e18);

        vm.startPrank(address(manager));

        vm.expectRevert(Errors.ZeroAddress.selector);
        vault.seize(alice, address(0), 1e18);

        vm.expectRevert(abi.encodeWithSelector(Errors.ExceedsCollateralBalance.selector, 2e18, 1e18));
        vault.seize(alice, liquidator, 2e18);

        vault.seize(alice, liquidator, 1e18);
        vm.stopPrank();

        assertEq(vault.collateralOf(alice), 0);
        assertEq(weth.balanceOf(liquidator), 1e18);
    }

    function test_links_are_one_time_and_owner_only() public {
        vm.startPrank(admin);

        vm.expectRevert(Errors.AlreadyInitialized.selector);
        vault.linkController(ILendingController(stranger));

        vm.expectRevert(Errors.AlreadyInitialized.selector);
        vault.linkLiquidationManager(stranger);

        vm.stopPrank();

        CollateralVault fresh = new CollateralVault(admin, IERC20(address(weth)));

        vm.prank(stranger);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, stranger));
        fresh.linkController(ILendingController(address(controller)));

        vm.startPrank(admin);
        vm.expectRevert(Errors.ZeroAddress.selector);
        fresh.linkController(ILendingController(address(0)));

        vm.expectRevert(Errors.ZeroAddress.selector);
        fresh.linkLiquidationManager(address(0));
        vm.stopPrank();
    }

    function test_owner_cannot_touch_user_collateral() public {
        _addCollateral(alice, 5e18);

        vm.startPrank(admin);

        vm.expectRevert(abi.encodeWithSelector(Errors.NotAuthorized.selector, admin));
        vault.seize(alice, admin, 5e18);

        vm.stopPrank();

        assertEq(vault.collateralOf(alice), 5e18);
        assertEq(weth.balanceOf(address(vault)), 5e18);
    }

    function testFuzz_deposit_withdraw_roundtrip_without_debt(uint256 amount) public {
        amount = bound(amount, 1, 1000e18);

        uint256 before = weth.balanceOf(alice);

        _addCollateral(alice, amount);

        vm.prank(alice);
        vault.withdrawCollateral(amount);

        assertEq(weth.balanceOf(alice), before);
        assertEq(vault.collateralOf(alice), 0);
    }

    function testFuzz_total_tracks_sum_of_balances(uint256 first, uint256 second) public {
        first = bound(first, 1, 500e18);
        second = bound(second, 1, 500e18);

        _addCollateral(alice, first);
        _addCollateral(bob, second);

        assertEq(vault.totalCollateral(), vault.collateralOf(alice) + vault.collateralOf(bob));
        assertEq(weth.balanceOf(address(vault)), vault.totalCollateral());
    }

    function testFuzz_safe_withdrawal_never_makes_position_liquidatable(uint256 collateral, uint256 debt)
        public
    {
        collateral = bound(collateral, 1e18, 100e18);

        _deposit(lender, 500_000e6);
        _addCollateral(alice, collateral);

        uint256 borrowable = controller.maxBorrowable(alice);
        if (borrowable == 0) {
            return;
        }

        debt = bound(debt, 1, borrowable);
        _borrow(alice, debt);

        uint256 safeAmount = controller.maxWithdrawableCollateral(alice);
        if (safeAmount == 0) {
            return;
        }

        vm.prank(alice);
        vault.withdrawCollateral(safeAmount);

        assertFalse(manager.isLiquidatable(alice));
    }
}
