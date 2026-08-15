// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

import {BaseTest} from "./Base.t.sol";
import {CollateralVault} from "../src/CollateralVault.sol";
import {LendingController} from "../src/LendingController.sol";
import {LendingPool} from "../src/LendingPool.sol";
import {MockAggregator} from "../src/mocks/MockAggregator.sol";
import {ICollateralVault} from "../src/interfaces/ICollateralVault.sol";
import {IInterestRateModel} from "../src/interfaces/IInterestRateModel.sol";
import {ILendingController} from "../src/interfaces/ILendingController.sol";
import {ILendingPool} from "../src/interfaces/ILendingPool.sol";
import {IPriceOracle} from "../src/interfaces/IPriceOracle.sol";
import {Errors} from "../src/libraries/Errors.sol";

contract ReentrantToken is ERC20 {
    address public target;
    bytes public payload;
    bool public armed;
    bool public reentered;
    bytes public lastRevert;

    constructor() ERC20("Reentrant", "RE") {}

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }

    function decimals() public pure override returns (uint8) {
        return 6;
    }

    function arm(address target_, bytes calldata payload_) external {
        target = target_;
        payload = payload_;
        armed = true;
    }

    function _update(address from, address to, uint256 value) internal override {
        super._update(from, to, value);

        if (armed && target != address(0)) {
            armed = false;
            reentered = true;

            (bool ok, bytes memory returned) = target.call(payload);

            if (!ok) {
                lastRevert = returned;
            }
        }
    }
}

contract ReentrantCollateral is ERC20 {
    address public target;
    bytes public payload;
    bool public armed;
    bytes public lastRevert;

    constructor() ERC20("ReentrantCollateral", "REC") {}

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }

    function arm(address target_, bytes calldata payload_) external {
        target = target_;
        payload = payload_;
        armed = true;
    }

    function _update(address from, address to, uint256 value) internal override {
        super._update(from, to, value);

        if (armed && target != address(0)) {
            armed = false;

            (bool ok, bytes memory returned) = target.call(payload);

            if (!ok) {
                lastRevert = returned;
            }
        }
    }
}

contract SecurityTest is BaseTest {
    function test_withdraw_cannot_be_reentered() public {
        ReentrantToken evil = new ReentrantToken();

        LendingPool evilPool =
            new LendingPool(admin, IERC20(address(evil)), IInterestRateModel(address(rateModel)), 1e6, 1000);

        evil.mint(alice, 100_000e6);

        vm.startPrank(alice);
        evil.approve(address(evilPool), type(uint256).max);
        evilPool.deposit(50_000e6);
        vm.stopPrank();

        evil.arm(address(evilPool), abi.encodeCall(LendingPool.withdraw, (1_000e6)));

        vm.prank(alice);
        evilPool.withdraw(1_000e6);

        assertTrue(evil.reentered());
        assertEq(bytes4(evil.lastRevert()), ReentrancyGuard.ReentrancyGuardReentrantCall.selector);
        assertEq(evilPool.balanceOfAssets(alice), 49_000e6);
    }

    function test_deposit_cannot_be_reentered() public {
        ReentrantToken evil = new ReentrantToken();

        LendingPool evilPool =
            new LendingPool(admin, IERC20(address(evil)), IInterestRateModel(address(rateModel)), 1e6, 1000);

        evil.mint(alice, 100_000e6);

        vm.startPrank(alice);
        evil.approve(address(evilPool), type(uint256).max);
        vm.stopPrank();

        evil.arm(address(evilPool), abi.encodeCall(LendingPool.deposit, (1_000e6)));

        vm.prank(alice);
        evilPool.deposit(10_000e6);

        assertEq(bytes4(evil.lastRevert()), ReentrancyGuard.ReentrancyGuardReentrantCall.selector);
        assertEq(evilPool.balanceOfAssets(alice), 10_000e6);
    }

    function test_collateral_withdraw_cannot_be_reentered() public {
        ReentrantCollateral evil = new ReentrantCollateral();
        MockAggregator evilFeed = new MockAggregator("REC / USD", 8, ETH_PRICE);

        CollateralVault evilVault = new CollateralVault(admin, IERC20(address(evil)));

        LendingController evilController = new LendingController(
            admin,
            ILendingPool(address(pool)),
            ICollateralVault(address(evilVault)),
            IPriceOracle(address(oracle)),
            MAX_LTV_BPS,
            LIQUIDATION_THRESHOLD_BPS,
            LIQUIDATION_BONUS_BPS
        );

        vm.startPrank(admin);
        oracle.setFeed(address(evil), address(evilFeed));
        evilVault.linkController(ILendingController(address(evilController)));
        vm.stopPrank();

        evil.mint(alice, 100e18);

        vm.startPrank(alice);
        evil.approve(address(evilVault), type(uint256).max);
        evilVault.depositCollateral(10e18);
        vm.stopPrank();

        evil.arm(address(evilVault), abi.encodeCall(CollateralVault.withdrawCollateral, (1e18)));

        vm.prank(alice);
        evilVault.withdrawCollateral(1e18);

        assertEq(bytes4(evil.lastRevert()), ReentrancyGuard.ReentrancyGuardReentrantCall.selector);
        assertEq(evilVault.collateralOf(alice), 9e18);
    }

    function test_owner_cannot_drain_the_pool() public {
        _deposit(lender, 100_000e6);

        uint256 poolCashBefore = _poolCash();

        vm.startPrank(admin);

        vm.expectRevert(Errors.ZeroAmount.selector);
        pool.collectAllReserves(admin);

        pool.setMinDeposit(type(uint256).max);
        pool.setDepositsPaused(true);
        pool.setReserveFactorBps(5000);

        vm.stopPrank();

        assertEq(_poolCash(), poolCashBefore);
        assertEq(pool.balanceOfAssets(lender), 100_000e6);

        vm.prank(lender);
        pool.withdraw(100_000e6);

        assertEq(usdc.balanceOf(lender), 1_000_000e6);
    }

    function test_owner_cannot_pause_the_exits() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 10e18, 5_000e6);

        vm.startPrank(admin);
        pool.setDepositsPaused(true);
        controller.setBorrowPaused(true);
        vm.stopPrank();

        vm.prank(alice);
        controller.repayAll();

        vm.prank(alice);
        vault.withdrawCollateral(10e18);

        vm.prank(lender);
        pool.withdraw(50_000e6);

        assertEq(controller.debtOf(alice), 0);
        assertEq(vault.collateralOf(alice), 0);
    }

    function test_owner_cannot_pause_liquidation() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        vm.startPrank(admin);
        pool.setDepositsPaused(true);
        controller.setBorrowPaused(true);
        vm.stopPrank();

        _setEthPrice(260000000000);

        vm.prank(liquidator);
        manager.liquidate(alice);

        assertEq(controller.debtOf(alice), 0);
    }

    function test_owner_cannot_repoint_the_links_to_steal() public {
        vm.startPrank(admin);

        vm.expectRevert(Errors.AlreadyInitialized.selector);
        pool.linkController(admin);

        vm.expectRevert(Errors.AlreadyInitialized.selector);
        pool.linkLiquidationManager(admin);

        vm.expectRevert(Errors.AlreadyInitialized.selector);
        vault.linkController(ILendingController(admin));

        vm.expectRevert(Errors.AlreadyInitialized.selector);
        vault.linkLiquidationManager(admin);

        vm.stopPrank();
    }

    function test_ownership_transfer_does_not_unlock_user_funds() public {
        _deposit(lender, 100_000e6);
        _addCollateral(alice, 10e18);

        vm.prank(admin);
        pool.transferOwnership(stranger);

        vm.startPrank(stranger);

        vm.expectRevert(Errors.AlreadyInitialized.selector);
        pool.linkController(stranger);

        vm.expectRevert(Errors.ZeroAmount.selector);
        pool.collectAllReserves(stranger);

        vm.stopPrank();

        assertEq(pool.balanceOfAssets(lender), 100_000e6);
        assertEq(vault.collateralOf(alice), 10e18);
    }

    function test_donation_does_not_change_share_value() public {
        _deposit(lender, 50_000e6);
        _deposit(lenderTwo, 50_000e6);

        uint256 lenderValue = pool.balanceOfAssets(lender);
        uint256 indexBefore = pool.supplyIndex();

        vm.prank(alice);
        usdc.transfer(address(pool), 100_000e6);

        assertEq(pool.balanceOfAssets(lender), lenderValue);
        assertEq(pool.supplyIndex(), indexBefore);
        assertEq(pool.totalSupplied(), 100_000e6);
    }

    function test_donated_tokens_cannot_be_withdrawn_by_lenders() public {
        _deposit(lender, 50_000e6);

        vm.prank(alice);
        usdc.transfer(address(pool), 50_000e6);

        vm.prank(lender);
        vm.expectRevert();
        pool.withdraw(100_000e6);

        vm.prank(lender);
        pool.withdraw(50_000e6);

        assertEq(_poolCash(), 50_000e6);
    }

    function test_donation_to_vault_does_not_credit_anyone() public {
        _addCollateral(alice, 5e18);

        vm.prank(bob);
        weth.transfer(address(vault), 50e18);

        assertEq(vault.collateralOf(alice), 5e18);
        assertEq(vault.collateralOf(bob), 0);
        assertEq(vault.totalCollateral(), 5e18);
    }

    function test_stale_price_blocks_every_money_path_that_needs_it() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        vm.warp(block.timestamp + MAX_PRICE_AGE + 1);

        vm.prank(alice);
        vm.expectRevert();
        controller.borrow(1e6);

        vm.prank(alice);
        vm.expectRevert();
        vault.withdrawCollateral(1e17);

        vm.prank(liquidator);
        vm.expectRevert();
        manager.liquidate(alice);
    }

    function test_broken_feed_cannot_be_used_to_liquidate() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        wethFeed.setFeedDown(true);

        vm.prank(liquidator);
        vm.expectRevert();
        manager.liquidate(alice);
    }

    function test_zero_price_cannot_be_used_to_liquidate() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        wethFeed.setPrice(0);

        vm.prank(liquidator);
        vm.expectRevert();
        manager.liquidate(alice);
    }

    function test_borrower_cannot_escape_liquidation_by_adding_dust() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);
        _setEthPrice(260000000000);

        assertTrue(manager.isLiquidatable(alice));

        _addCollateral(alice, 1);

        assertTrue(manager.isLiquidatable(alice));

        vm.prank(liquidator);
        manager.liquidate(alice);

        assertEq(controller.debtOf(alice), 0);
    }

    function test_self_liquidation_is_not_profitable_beyond_the_bonus() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);
        _setEthPrice(260000000000);

        uint256 usdcBefore = usdc.balanceOf(alice);
        uint256 wethBefore = weth.balanceOf(alice);

        vm.prank(alice);
        manager.liquidate(alice);

        assertEq(usdcBefore - usdc.balanceOf(alice), 6_900e6);
        assertEq(weth.balanceOf(alice) - wethBefore, 2786538461538461539);
        assertEq(vault.collateralOf(alice), 3.2e18 - 2786538461538461539);
        assertEq(controller.debtOf(alice), 0);
    }

    function test_third_party_cannot_move_another_users_position() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        vm.startPrank(stranger);

        vm.expectRevert();
        vault.withdrawCollateral(1e18);

        vm.expectRevert(Errors.ZeroAmount.selector);
        controller.repayAll();

        vm.stopPrank();

        assertEq(vault.collateralOf(alice), 3.2e18);
        assertEq(controller.debtOf(alice), 6_900e6);
    }

    function test_repaying_for_someone_else_is_not_possible_through_the_controller() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 3.2e18, 6_900e6);

        vm.prank(bob);
        vm.expectRevert(Errors.ZeroAmount.selector);
        controller.repayAll();

        assertEq(controller.debtOf(alice), 6_900e6);
    }

    function test_accrueInterest_is_permissionless_and_harmless() public {
        _deposit(lender, 100_000e6);
        _openPosition(alice, 10e18, 6_900e6);
        _skip(YEAR);

        vm.prank(stranger);
        pool.accrueInterest();

        uint256 index = pool.borrowIndex();

        vm.prank(stranger);
        pool.accrueInterest();

        assertEq(pool.borrowIndex(), index);
    }

    function testFuzz_no_sequence_of_owner_calls_can_reduce_lender_balance(
        uint256 minDeposit,
        uint16 reserveFactor,
        bool pauseDeposits,
        bool pauseBorrow
    ) public {
        reserveFactor = uint16(bound(reserveFactor, 0, 5000));

        _deposit(lender, 100_000e6);

        uint256 balanceBefore = pool.balanceOfAssets(lender);

        vm.startPrank(admin);
        pool.setMinDeposit(minDeposit);
        pool.setReserveFactorBps(reserveFactor);
        pool.setDepositsPaused(pauseDeposits);
        controller.setBorrowPaused(pauseBorrow);
        vm.stopPrank();

        assertEq(pool.balanceOfAssets(lender), balanceBefore);
    }

    function testFuzz_donations_never_change_anyones_balance(uint256 donation) public {
        donation = bound(donation, 1, 100_000e6);

        _deposit(lender, 50_000e6);

        uint256 before = pool.balanceOfAssets(lender);

        vm.prank(alice);
        usdc.transfer(address(pool), donation);

        assertEq(pool.balanceOfAssets(lender), before);
    }
}
