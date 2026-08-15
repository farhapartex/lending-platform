// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {BaseTest} from "./Base.t.sol";
import {LiquidationManager} from "../src/LiquidationManager.sol";
import {ICollateralVault} from "../src/interfaces/ICollateralVault.sol";
import {ILendingController} from "../src/interfaces/ILendingController.sol";
import {ILendingPool} from "../src/interfaces/ILendingPool.sol";
import {IPriceOracle} from "../src/interfaces/IPriceOracle.sol";
import {Errors} from "../src/libraries/Errors.sol";

contract LiquidationManagerTest is BaseTest {
    int256 constant STRESSED_PRICE = 260000000000;
    int256 constant CRASH_PRICE = 200000000000;

    function _underwaterPosition() internal {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);
        _setEthPrice(STRESSED_PRICE);
    }

    function test_constructor_rejects_zero_dependencies() public {
        vm.expectRevert(Errors.ZeroAddress.selector);
        new LiquidationManager(
            ILendingPool(address(0)),
            ICollateralVault(address(vault)),
            IPriceOracle(address(oracle)),
            ILendingController(address(controller))
        );

        vm.expectRevert(Errors.ZeroAddress.selector);
        new LiquidationManager(
            ILendingPool(address(pool)),
            ICollateralVault(address(vault)),
            IPriceOracle(address(oracle)),
            ILendingController(address(0))
        );
    }

    function test_manager_holds_no_storage_and_no_owner() public view {
        assertEq(manager.collateralAsset(), address(weth));
        assertEq(manager.debtAsset(), address(usdc));
        assertEq(address(manager.controller()), address(controller));
    }

    function test_healthy_position_cannot_be_liquidated() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        assertFalse(manager.isLiquidatable(alice));

        vm.prank(liquidator);
        vm.expectRevert(abi.encodeWithSelector(Errors.PositionIsHealthy.selector, 12661));
        manager.liquidate(alice);
    }

    function test_position_with_no_debt_cannot_be_liquidated() public {
        _addCollateral(alice, 3.2e18);

        assertFalse(manager.isLiquidatable(alice));

        vm.prank(liquidator);
        vm.expectRevert(abi.encodeWithSelector(Errors.PositionIsHealthy.selector, type(uint256).max));
        manager.liquidate(alice);
    }

    function test_preview_matches_expected_figures() public {
        _underwaterPosition();

        assertEq(controller.healthFactorBps(alice), 9646);
        assertTrue(manager.isLiquidatable(alice));

        (uint256 debtToRepay, uint256 seize, uint256 bonus, uint256 shortfall) =
            manager.previewLiquidation(alice);

        assertEq(debtToRepay, 6_900e6);
        assertEq(seize, 2786538461538461539);
        assertEq(bonus, 34500000000);
        assertEq(shortfall, 0);
    }

    function test_liquidation_pays_the_bonus_exactly() public {
        _underwaterPosition();

        uint256 liquidatorUsdcBefore = usdc.balanceOf(liquidator);

        vm.prank(liquidator);
        manager.liquidate(alice);

        uint256 spent = liquidatorUsdcBefore - usdc.balanceOf(liquidator);
        uint256 received = weth.balanceOf(liquidator);
        uint256 receivedValue = (received * uint256(STRESSED_PRICE)) / 1e18;

        assertEq(spent, 6_900e6);
        assertEq(received, 2786538461538461539);
        assertEq(receivedValue, 724500000000);
        assertEq(receivedValue - 690000000000, 34500000000);
    }

    function test_liquidation_clears_debt_and_returns_surplus_collateral() public {
        _underwaterPosition();

        vm.prank(liquidator);
        manager.liquidate(alice);

        assertEq(controller.debtOf(alice), 0);
        assertEq(pool.debtSharesOf(alice), 0);
        assertEq(vault.collateralOf(alice), 3.2e18 - 2786538461538461539);
        assertEq(controller.healthFactorBps(alice), type(uint256).max);
    }

    function test_second_liquidator_loses_the_race() public {
        _underwaterPosition();

        vm.prank(liquidator);
        manager.liquidate(alice);

        vm.prank(bob);
        vm.expectRevert(abi.encodeWithSelector(Errors.PositionIsHealthy.selector, type(uint256).max));
        manager.liquidate(alice);
    }

    function test_price_recovery_before_the_transaction_lands_blocks_liquidation() public {
        _underwaterPosition();

        assertTrue(manager.isLiquidatable(alice));

        _setEthPrice(ETH_PRICE);

        vm.prank(liquidator);
        vm.expectRevert(abi.encodeWithSelector(Errors.PositionIsHealthy.selector, 12661));
        manager.liquidate(alice);
    }

    function test_shortfall_case_seizes_everything_and_records_the_gap() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);
        _setEthPrice(CRASH_PRICE);

        (, uint256 seize, uint256 bonus, uint256 shortfall) = manager.previewLiquidation(alice);

        assertEq(seize, 3.2e18);
        assertEq(bonus, 34500000000);
        assertEq(shortfall, 84500000000);

        vm.prank(liquidator);
        manager.liquidate(alice);

        assertEq(controller.debtOf(alice), 0);
        assertEq(vault.collateralOf(alice), 0);
        assertEq(weth.balanceOf(liquidator), 3.2e18);
    }

    function test_liquidation_accrues_interest_first() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);
        _skip(YEAR);
        _setEthPrice(STRESSED_PRICE);

        uint256 owedWithInterest = controller.debtOf(alice);
        assertGt(owedWithInterest, 6_900e6);

        uint256 before = usdc.balanceOf(liquidator);

        vm.prank(liquidator);
        manager.liquidate(alice);

        assertEq(before - usdc.balanceOf(liquidator), owedWithInterest);
    }

    function test_liquidation_is_blocked_when_the_price_is_stale() public {
        _underwaterPosition();

        vm.warp(block.timestamp + MAX_PRICE_AGE + 1);

        vm.prank(liquidator);
        vm.expectRevert();
        manager.liquidate(alice);
    }

    function test_liquidation_needs_liquidator_approval_and_balance() public {
        _underwaterPosition();

        vm.prank(stranger);
        vm.expectRevert();
        manager.liquidate(alice);
    }

    function test_anyone_may_liquidate() public {
        _underwaterPosition();

        _fund(bob, 0, 0);

        vm.prank(bob);
        manager.liquidate(alice);

        assertEq(controller.debtOf(alice), 0);
        assertGt(weth.balanceOf(bob), 1000e18);
    }

    function test_preview_and_execute_agree() public {
        _underwaterPosition();

        (uint256 debtToRepay, uint256 seize,,) = manager.previewLiquidation(alice);

        uint256 usdcBefore = usdc.balanceOf(liquidator);

        vm.prank(liquidator);
        manager.liquidate(alice);

        assertEq(usdcBefore - usdc.balanceOf(liquidator), debtToRepay);
        assertEq(weth.balanceOf(liquidator), seize);
    }

    function testFuzz_healthy_positions_are_never_liquidatable(uint256 collateral, uint256 portionBps)
        public
    {
        collateral = bound(collateral, 1e18, 100e18);
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

        vm.prank(liquidator);
        vm.expectRevert();
        manager.liquidate(alice);
    }

    function testFuzz_liquidation_always_clears_the_debt(uint256 collateral, uint256 dropBps) public {
        collateral = bound(collateral, 1e18, 50e18);
        dropBps = bound(dropBps, 3000, 9000);

        _deposit(lender, 500_000e6);
        _addCollateral(alice, collateral);

        uint256 room = controller.maxBorrowable(alice);
        if (room < 1e6) {
            return;
        }

        _borrow(alice, room);

        int256 droppedPrice = (ETH_PRICE * int256(dropBps)) / 10_000;
        _setEthPrice(droppedPrice);

        if (!manager.isLiquidatable(alice)) {
            return;
        }

        vm.prank(liquidator);
        manager.liquidate(alice);

        assertEq(controller.debtOf(alice), 0);
        assertEq(pool.debtSharesOf(alice), 0);
    }

    function testFuzz_liquidation_never_seizes_more_than_held(uint256 collateral, uint256 dropBps) public {
        collateral = bound(collateral, 1e18, 50e18);
        dropBps = bound(dropBps, 1000, 9000);

        _deposit(lender, 500_000e6);
        _addCollateral(alice, collateral);

        uint256 room = controller.maxBorrowable(alice);
        if (room < 1e6) {
            return;
        }

        _borrow(alice, room);
        _setEthPrice((ETH_PRICE * int256(dropBps)) / 10_000);

        if (!manager.isLiquidatable(alice)) {
            return;
        }

        uint256 held = vault.collateralOf(alice);

        vm.prank(liquidator);
        manager.liquidate(alice);

        assertLe(weth.balanceOf(liquidator), held);
        assertEq(vault.collateralOf(alice), held - weth.balanceOf(liquidator));
    }

    function testFuzz_liquidation_is_profitable_when_there_is_no_shortfall(
        uint256 collateral,
        uint256 dropBps
    ) public {
        collateral = bound(collateral, 1e18, 50e18);
        dropBps = bound(dropBps, 7000, 9500);

        _deposit(lender, 500_000e6);
        _addCollateral(alice, collateral);

        uint256 room = controller.maxBorrowable(alice);
        if (room < 1e6) {
            return;
        }

        _borrow(alice, room);

        int256 droppedPrice = (ETH_PRICE * int256(dropBps)) / 10_000;
        _setEthPrice(droppedPrice);

        if (!manager.isLiquidatable(alice)) {
            return;
        }

        (,,, uint256 shortfall) = manager.previewLiquidation(alice);
        if (shortfall != 0) {
            return;
        }

        uint256 usdcBefore = usdc.balanceOf(liquidator);

        vm.prank(liquidator);
        manager.liquidate(alice);

        uint256 spent = usdcBefore - usdc.balanceOf(liquidator);
        uint256 receivedValue = (weth.balanceOf(liquidator) * uint256(droppedPrice)) / 1e18;
        uint256 spentValue = spent * 100;

        assertGe(receivedValue, spentValue);
    }
}
