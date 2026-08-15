// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {BaseTest} from "./Base.t.sol";
import {InterestRateModel} from "../src/InterestRateModel.sol";
import {LendingPool} from "../src/LendingPool.sol";
import {IInterestRateModel, RateCurve} from "../src/interfaces/IInterestRateModel.sol";
import {Errors} from "../src/libraries/Errors.sol";
import {WadMath} from "../src/libraries/WadMath.sol";

contract EdgeCasesTest is BaseTest {
    function test_wadDivUp_rounds_away_from_zero() public pure {
        assertEq(WadMath.wadDivUp(3, 2e18), 2);
        assertEq(WadMath.wadDivDown(3, 2e18), 1);
        assertEq(WadMath.wadDivUp(0, 2e18), 0);
    }

    function test_scaleUpByBpsDown_rounds_toward_zero() public pure {
        assertEq(WadMath.scaleUpByBpsDown(690000000000, 7500), 920000000000);
        assertEq(WadMath.scaleUpByBpsDown(1, 7500), 1);
        assertEq(WadMath.scaleUpByBpsUp(1, 7500), 2);
    }

    function test_wadMulUp_rounds_away_from_zero() public pure {
        assertEq(WadMath.wadMulUp(1, 1), 1);
        assertEq(WadMath.wadMulDown(1, 1), 0);
    }

    function test_lastAccrualTimestamp_tracks_accrual() public {
        uint256 startedAt = pool.lastAccrualTimestamp();

        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);
        _skip(YEAR);
        pool.accrueInterest();

        assertEq(pool.lastAccrualTimestamp(), block.timestamp);
        assertGt(pool.lastAccrualTimestamp(), startedAt);
    }

    function test_pool_utilizationBps_view() public {
        assertEq(pool.utilizationBps(), 0);

        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);

        assertEq(pool.utilizationBps(), 1380);
    }

    function test_controller_isLiquidatable_view() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        assertFalse(controller.isLiquidatable(alice));

        _setEthPrice(260000000000);

        assertTrue(controller.isLiquidatable(alice));
    }

    function test_linkLiquidationManager_rejects_zero() public {
        LendingPool fresh =
            new LendingPool(admin, IERC20(address(usdc)), IInterestRateModel(address(rateModel)), 1e6, 1000);

        vm.prank(admin);
        vm.expectRevert(Errors.ZeroAddress.selector);
        fresh.linkLiquidationManager(address(0));
    }

    function test_setRateModel_swaps_the_curve() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);

        InterestRateModel replacement = new InterestRateModel(
            admin,
            RateCurve({
                baseRatePerSecond: uint64(uint256(2e16) / YEAR),
                slopeBelowKinkPerSecond: uint64(uint256(654e14) / YEAR),
                slopeAboveKinkPerSecond: uint64(uint256(2e18) / YEAR),
                kinkUtilizationBps: 8000
            })
        );

        uint256 rateBefore = rateModel.borrowRatePerSecond(1380);

        vm.prank(admin);
        pool.setRateModel(IInterestRateModel(address(replacement)));

        assertEq(address(pool.rateModel()), address(replacement));
        assertGt(replacement.borrowRatePerSecond(1380), rateBefore);
    }

    function test_setRateModel_accrues_before_switching() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);
        _skip(YEAR);

        InterestRateModel replacement = new InterestRateModel(admin, _defaultCurve());

        vm.prank(admin);
        pool.setRateModel(IInterestRateModel(address(replacement)));

        assertEq(pool.accruedReserves(), 14684235);
        assertEq(pool.borrowIndex(), 1021281499965232000);
    }

    function test_pool_rechecks_liquidity_even_when_the_controller_asks() public {
        _deposit(lender, 10_000e6);

        uint256 liquidity = pool.availableLiquidity();

        vm.prank(address(controller));
        vm.expectRevert(abi.encodeWithSelector(Errors.NotEnoughLiquidity.selector, liquidity + 1, liquidity));
        pool.borrowFor(alice, alice, liquidity + 1);
    }

    function test_repayFor_clamps_an_overpayment_to_the_debt() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);

        vm.prank(alice);
        usdc.approve(address(pool), type(uint256).max);

        vm.prank(address(controller));
        pool.repayFor(alice, alice, 10_000e6);

        assertEq(pool.debtSharesOf(alice), 0);
        assertEq(pool.debtOf(alice), 0);
        assertEq(pool.totalDebtShares(), 0);
    }

    function test_zero_rate_curve_accrues_nothing_over_time() public {
        InterestRateModel flat = new InterestRateModel(
            admin,
            RateCurve({
                baseRatePerSecond: 0,
                slopeBelowKinkPerSecond: 0,
                slopeAboveKinkPerSecond: 0,
                kinkUtilizationBps: 8000
            })
        );

        vm.prank(admin);
        pool.setRateModel(IInterestRateModel(address(flat)));

        _deposit(lender, 50_000e6);
        _openPosition(alice, 10e18, 6_900e6);

        _skip(YEAR);

        (uint256 supplyIndex, uint256 borrowIndex, uint256 reserves) = pool.previewAccrual();

        assertEq(supplyIndex, 1e18);
        assertEq(borrowIndex, 1e18);
        assertEq(reserves, 0);

        pool.accrueInterest();

        assertEq(pool.debtOf(alice), 6_900e6);
        assertEq(pool.balanceOfAssets(lender), 50_000e6);
    }

    function test_deposit_that_would_mint_no_shares_is_rejected() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 100e18, 40_000e6);

        _skip(YEAR);
        pool.accrueInterest();

        assertGt(pool.supplyIndex(), 1e18);

        vm.prank(admin);
        pool.setMinDeposit(1);

        vm.prank(bob);
        vm.expectRevert(Errors.ZeroAmount.selector);
        pool.deposit(1);
    }

    function test_dust_deposit_is_accepted_once_it_buys_a_share() public {
        _deposit(lender, 50_000e6);
        _openPosition(alice, 100e18, 40_000e6);

        _skip(YEAR);
        pool.accrueInterest();

        vm.prank(admin);
        pool.setMinDeposit(1);

        vm.prank(bob);
        uint256 shares = pool.deposit(2);

        assertGt(shares, 0);
    }

    function test_accrual_across_many_small_steps_matches_one_big_step() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 50e18, 20_000e6);

        uint256 checkpoint = vm.snapshotState();

        _skip(365 days);
        pool.accrueInterest();
        uint256 oneStep = pool.debtOf(alice);

        vm.revertToState(checkpoint);

        for (uint256 i = 0; i < 12; i++) {
            _skip(30 days);
            pool.accrueInterest();
        }
        _skip(5 days);
        pool.accrueInterest();

        uint256 manySteps = pool.debtOf(alice);

        assertGe(manySteps, oneStep);
    }

    function test_position_survives_a_full_lifecycle() public {
        _deposit(lender, 100_000e6);
        _deposit(lenderTwo, 50_000e6);

        _openPosition(alice, 20e18, 10_000e6);
        _openPosition(bob, 10e18, 5_000e6);

        _skip(180 days);

        vm.prank(alice);
        controller.repay(2_000e6);

        _addCollateral(alice, 5e18);

        _skip(90 days);

        vm.prank(bob);
        controller.repayAll();

        vm.prank(bob);
        vault.withdrawCollateral(10e18);

        vm.prank(alice);
        controller.repayAll();

        vm.prank(alice);
        vault.withdrawCollateral(25e18);

        vm.prank(admin);
        pool.collectAllReserves(treasury);

        uint256 lenderOwed = pool.balanceOfAssets(lender);
        uint256 lenderTwoOwed = pool.balanceOfAssets(lenderTwo);

        vm.prank(lender);
        pool.withdraw(lenderOwed);

        vm.prank(lenderTwo);
        pool.withdraw(lenderTwoOwed);

        assertEq(vault.totalCollateral(), 0);
        assertEq(pool.totalDebtShares(), 0);
        assertLe(pool.totalSupplied(), 1);
        assertGt(usdc.balanceOf(treasury), 0);
        assertLe(_poolCash(), 2);
    }
}
