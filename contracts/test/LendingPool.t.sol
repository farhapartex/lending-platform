// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {BaseTest} from "./Base.t.sol";
import {LendingPool} from "../src/LendingPool.sol";
import {IInterestRateModel} from "../src/interfaces/IInterestRateModel.sol";
import {Errors} from "../src/libraries/Errors.sol";
import {ShareMath} from "../src/libraries/ShareMath.sol";

contract LendingPoolTest is BaseTest {
    function test_constructor_rejects_zero_addresses() public {
        vm.expectRevert(Errors.ZeroAddress.selector);
        new LendingPool(admin, IERC20(address(0)), IInterestRateModel(address(rateModel)), 1e6, 1000);

        vm.expectRevert(Errors.ZeroAddress.selector);
        new LendingPool(admin, IERC20(address(usdc)), IInterestRateModel(address(0)), 1e6, 1000);
    }

    function test_constructor_caps_reserve_factor() public {
        vm.expectRevert(Errors.InvalidRiskSettings.selector);
        new LendingPool(admin, IERC20(address(usdc)), IInterestRateModel(address(rateModel)), 1e6, 5001);
    }

    function test_initial_state() public view {
        assertEq(pool.supplyIndex(), 1e18);
        assertEq(pool.borrowIndex(), 1e18);
        assertEq(pool.totalSupplied(), 0);
        assertEq(pool.totalBorrowed(), 0);
        assertEq(pool.availableLiquidity(), 0);
        assertEq(pool.accruedReserves(), 0);
        assertEq(pool.asset(), address(usdc));
    }

    function test_deposit_mints_shares_one_to_one_at_genesis() public {
        uint256 shares = _deposit(lender, 50_000e6);

        assertEq(shares, 50_000e6);
        assertEq(pool.sharesOf(lender), 50_000e6);
        assertEq(pool.balanceOfAssets(lender), 50_000e6);
        assertEq(pool.totalSupplied(), 50_000e6);
        assertEq(_poolCash(), 50_000e6);
    }

    function test_deposit_rejects_zero() public {
        vm.prank(lender);
        vm.expectRevert(Errors.ZeroAmount.selector);
        pool.deposit(0);
    }

    function test_deposit_enforces_minimum() public {
        vm.prank(lender);
        vm.expectRevert(abi.encodeWithSelector(Errors.BelowMinimumDeposit.selector, 1e6 - 1, 1e6));
        pool.deposit(1e6 - 1);
    }

    function test_deposit_respects_pause() public {
        vm.prank(admin);
        pool.setDepositsPaused(true);

        vm.prank(lender);
        vm.expectRevert(Errors.MarketPaused.selector);
        pool.deposit(1000e6);

        vm.prank(admin);
        pool.setDepositsPaused(false);

        _deposit(lender, 1000e6);
        assertEq(pool.balanceOfAssets(lender), 1000e6);
    }

    function test_withdraw_returns_principal() public {
        _deposit(lender, 50_000e6);

        uint256 balanceBefore = usdc.balanceOf(lender);

        vm.prank(lender);
        pool.withdraw(20_000e6);

        assertEq(usdc.balanceOf(lender) - balanceBefore, 20_000e6);
        assertEq(pool.balanceOfAssets(lender), 30_000e6);
    }

    function test_withdraw_rejects_zero() public {
        _deposit(lender, 1000e6);

        vm.prank(lender);
        vm.expectRevert(Errors.ZeroAmount.selector);
        pool.withdraw(0);
    }

    function test_withdraw_beyond_balance_is_distinct_from_liquidity() public {
        _deposit(lender, 10_000e6);

        vm.prank(lender);
        vm.expectRevert(abi.encodeWithSelector(Errors.ExceedsSupplyBalance.selector, 20_000e6, 10_000e6));
        pool.withdraw(20_000e6);
    }

    function test_withdraw_beyond_liquidity_reports_liquidity_error() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 30e18, 40_000e6);

        vm.prank(lender);
        vm.expectRevert(abi.encodeWithSelector(Errors.NotEnoughLiquidity.selector, 45_000e6, 10_000e6));
        pool.withdraw(45_000e6);
    }

    function test_maxWithdrawable_is_capped_by_liquidity() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 30e18, 40_000e6);

        assertEq(pool.balanceOfAssets(lender), 50_000e6);
        assertEq(pool.maxWithdrawable(lender), 10_000e6);
    }

    function test_redeemShares_pays_out_rounded_down() public {
        _deposit(lender, 50_000e6);

        vm.prank(lender);
        uint256 assets = pool.redeemShares(10_000e6);

        assertEq(assets, 10_000e6);
        assertEq(pool.sharesOf(lender), 40_000e6);
    }

    function test_redeemShares_rejects_zero() public {
        _deposit(lender, 1000e6);

        vm.prank(lender);
        vm.expectRevert(Errors.ZeroAmount.selector);
        pool.redeemShares(0);
    }

    function test_withdraw_and_redeem_round_toward_the_pool() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);
        _skip(YEAR);
        pool.accrueInterest();

        uint256 index = pool.supplyIndex();

        assertEq(
            ShareMath.sharesFromAssetsUp(1000e6, index), ShareMath.sharesFromAssetsDown(1000e6, index) + 1
        );
    }

    function test_accrual_is_idempotent_within_a_block() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);
        _skip(YEAR);

        pool.accrueInterest();
        uint256 borrowIndex = pool.borrowIndex();
        uint256 reserves = pool.accruedReserves();

        pool.accrueInterest();
        pool.accrueInterest();

        assertEq(pool.borrowIndex(), borrowIndex);
        assertEq(pool.accruedReserves(), reserves);
    }

    function test_accrual_does_nothing_without_borrowers() public {
        _deposit(lender, 50_000e6);
        _skip(YEAR);

        pool.accrueInterest();

        assertEq(pool.supplyIndex(), 1e18);
        assertEq(pool.borrowIndex(), 1e18);
        assertEq(pool.accruedReserves(), 0);
        assertEq(pool.balanceOfAssets(lender), 50_000e6);
    }

    function test_accrual_matches_known_figures() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);
        _skip(YEAR);

        pool.accrueInterest();

        assertEq(pool.supplyIndex(), 1002643162300000000);
        assertEq(pool.borrowIndex(), 1021281499965232000);
        assertEq(pool.accruedReserves(), 14684235);
        assertEq(pool.debtOf(alice), 7046842350);
        assertEq(pool.balanceOfAssets(lender), 50132158115);
    }

    function test_reserve_split_is_exact() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);

        uint256 debtBefore = pool.debtOf(alice);
        uint256 lenderBefore = pool.balanceOfAssets(lender);

        _skip(YEAR);
        pool.accrueInterest();

        uint256 interest = pool.debtOf(alice) - debtBefore;
        uint256 lendersGain = pool.balanceOfAssets(lender) - lenderBefore;

        assertEq(lendersGain + pool.accruedReserves(), interest);
    }

    function test_previewAccrual_predicts_accrueInterest() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);
        _skip(YEAR);

        (uint256 supplyIndex, uint256 borrowIndex, uint256 reserves) = pool.previewAccrual();

        pool.accrueInterest();

        assertEq(pool.supplyIndex(), supplyIndex);
        assertEq(pool.borrowIndex(), borrowIndex);
        assertEq(pool.accruedReserves(), reserves);
    }

    function test_previewAccrual_is_flat_when_nothing_pending() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);
        pool.accrueInterest();

        (uint256 supplyIndex, uint256 borrowIndex, uint256 reserves) = pool.previewAccrual();

        assertEq(supplyIndex, pool.supplyIndex());
        assertEq(borrowIndex, pool.borrowIndex());
        assertEq(reserves, 0);
    }

    function test_debt_primitives_are_controller_only() public {
        vm.startPrank(stranger);

        vm.expectRevert(abi.encodeWithSelector(Errors.NotAuthorized.selector, stranger));
        pool.borrowFor(stranger, stranger, 1e6);

        vm.expectRevert(abi.encodeWithSelector(Errors.NotAuthorized.selector, stranger));
        pool.repayFor(stranger, stranger, 1e6);

        vm.expectRevert(abi.encodeWithSelector(Errors.NotAuthorized.selector, stranger));
        pool.repayAllFor(stranger, stranger);

        vm.stopPrank();
    }

    function test_liquidation_manager_cannot_mint_debt() public {
        vm.startPrank(address(manager));

        vm.expectRevert(abi.encodeWithSelector(Errors.NotAuthorized.selector, address(manager)));
        pool.borrowFor(alice, alice, 1e6);

        vm.expectRevert(abi.encodeWithSelector(Errors.NotAuthorized.selector, address(manager)));
        pool.repayFor(alice, alice, 1e6);

        vm.stopPrank();
    }

    function test_links_are_one_time_only() public {
        vm.startPrank(admin);

        vm.expectRevert(Errors.AlreadyInitialized.selector);
        pool.linkController(stranger);

        vm.expectRevert(Errors.AlreadyInitialized.selector);
        pool.linkLiquidationManager(stranger);

        vm.stopPrank();
    }

    function test_links_reject_zero_and_non_owner() public {
        LendingPool fresh =
            new LendingPool(admin, IERC20(address(usdc)), IInterestRateModel(address(rateModel)), 1e6, 1000);

        vm.prank(stranger);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, stranger));
        fresh.linkController(address(controller));

        vm.prank(admin);
        vm.expectRevert(Errors.ZeroAddress.selector);
        fresh.linkController(address(0));
    }

    function test_owner_settings_and_caps() public {
        vm.startPrank(admin);

        pool.setMinDeposit(5e6);
        assertEq(pool.minDeposit(), 5e6);

        vm.expectRevert(Errors.InvalidRiskSettings.selector);
        pool.setReserveFactorBps(5001);

        pool.setReserveFactorBps(2000);
        assertEq(pool.reserveFactorBps(), 2000);

        vm.expectRevert(Errors.ZeroAddress.selector);
        pool.setRateModel(IInterestRateModel(address(0)));

        vm.stopPrank();

        vm.prank(stranger);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, stranger));
        pool.setMinDeposit(1);
    }

    function test_parameter_changes_accrue_first() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);
        _skip(YEAR);

        vm.prank(admin);
        pool.setReserveFactorBps(2000);

        assertEq(pool.accruedReserves(), 14684235);
    }

    function test_collectReserves_rules() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);
        _skip(YEAR);
        pool.accrueInterest();

        uint256 reserves = pool.accruedReserves();

        vm.prank(stranger);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, stranger));
        pool.collectReserves(stranger, 1);

        vm.startPrank(admin);

        vm.expectRevert(Errors.ZeroAmount.selector);
        pool.collectReserves(treasury, 0);

        vm.expectRevert(Errors.ZeroAddress.selector);
        pool.collectReserves(address(0), 1);

        vm.expectRevert(abi.encodeWithSelector(Errors.ExceedsReserves.selector, reserves + 1, reserves));
        pool.collectReserves(treasury, reserves + 1);

        pool.collectReserves(treasury, reserves);
        vm.stopPrank();

        assertEq(usdc.balanceOf(treasury), reserves);
        assertEq(pool.accruedReserves(), 0);
    }

    function test_collectAllReserves_reverts_when_empty() public {
        vm.prank(admin);
        vm.expectRevert(Errors.ZeroAmount.selector);
        pool.collectAllReserves(treasury);
    }

    function test_collecting_reserves_leaves_lender_liquidity_intact() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);
        _skip(YEAR);
        pool.accrueInterest();

        uint256 liquidityBefore = pool.availableLiquidity();
        uint256 suppliedBefore = pool.totalSupplied();

        vm.prank(admin);
        pool.collectAllReserves(treasury);

        assertEq(pool.availableLiquidity(), liquidityBefore);
        assertEq(pool.totalSupplied(), suppliedBefore);
        assertEq(_poolCash(), liquidityBefore);
    }

    function test_solvency_identity_holds() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);
        _skip(YEAR);
        pool.accrueInterest();

        assertEq(_poolCash(), pool.availableLiquidity() + pool.accruedReserves());
    }

    function test_direct_transfer_cannot_move_share_value() public {
        _deposit(lender, 50_000e6);

        uint256 valueBefore = pool.balanceOfAssets(lender);
        uint256 suppliedBefore = pool.totalSupplied();

        vm.prank(alice);
        usdc.transfer(address(pool), 100_000e6);

        assertEq(pool.balanceOfAssets(lender), valueBefore);
        assertEq(pool.totalSupplied(), suppliedBefore);
        assertEq(pool.supplyIndex(), 1e18);
    }

    function test_first_depositor_cannot_be_inflated_out() public {
        vm.prank(alice);
        usdc.transfer(address(pool), 100_000e6);

        uint256 shares = _deposit(lender, 1000e6);

        assertEq(shares, 1000e6);
        assertEq(pool.balanceOfAssets(lender), 1000e6);
    }

    function testFuzz_deposit_then_withdraw_never_profits(uint256 amount) public {
        amount = bound(amount, MIN_DEPOSIT, 500_000e6);

        uint256 before = usdc.balanceOf(lender);

        _deposit(lender, amount);

        uint256 shares = pool.sharesOf(lender);

        vm.prank(lender);
        pool.redeemShares(shares);

        assertLe(usdc.balanceOf(lender), before);
    }

    function testFuzz_shares_track_deposits_proportionally(uint256 first, uint256 second) public {
        first = bound(first, MIN_DEPOSIT, 500_000e6);
        second = bound(second, MIN_DEPOSIT, 500_000e6);

        _deposit(lender, first);
        _deposit(lenderTwo, second);

        assertEq(pool.totalSupplyShares(), pool.sharesOf(lender) + pool.sharesOf(lenderTwo));

        if (first > second) {
            assertGe(pool.sharesOf(lender), pool.sharesOf(lenderTwo));
        }
    }

    function testFuzz_accrual_never_shrinks_indexes(uint256 elapsed, uint256 borrowAmount) public {
        elapsed = bound(elapsed, 0, 10 * YEAR);
        borrowAmount = bound(borrowAmount, 1e6, 100_000e6);

        _deposit(lender, 500_000e6);
        _openPosition(alice, 200e18, borrowAmount);

        uint256 supplyBefore = pool.supplyIndex();
        uint256 borrowBefore = pool.borrowIndex();

        _skip(elapsed);
        pool.accrueInterest();

        assertGe(pool.supplyIndex(), supplyBefore);
        assertGe(pool.borrowIndex(), borrowBefore);
        assertGe(pool.borrowIndex(), pool.supplyIndex());
    }

    function testFuzz_pool_cash_always_covers_liquidity(uint256 depositAmount, uint256 borrowAmount) public {
        depositAmount = bound(depositAmount, 10_000e6, 500_000e6);
        borrowAmount = bound(borrowAmount, 1e6, depositAmount / 2);

        _deposit(lender, depositAmount);
        _openPosition(alice, 500e18, borrowAmount);

        assertGe(_poolCash(), pool.availableLiquidity());
    }
}
