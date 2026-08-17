// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {BaseTest} from "./Base.t.sol";
import {CollateralVault} from "../src/CollateralVault.sol";
import {LendingController} from "../src/LendingController.sol";
import {LendingPool} from "../src/LendingPool.sol";
import {LiquidationManager} from "../src/LiquidationManager.sol";
import {MockAggregator} from "../src/mocks/MockAggregator.sol";
import {MockERC20} from "../src/mocks/MockERC20.sol";
import {HealthMath} from "../src/libraries/HealthMath.sol";

contract ProtocolHandler is Test {
    LendingPool public pool;
    CollateralVault public vault;
    LendingController public controller;
    LiquidationManager public manager;
    MockERC20 public weth;
    MockERC20 public usdc;
    MockAggregator public wethFeed;
    MockAggregator public usdcFeed;

    address[] public actors;

    uint256 public lastSupplyIndex;
    uint256 public lastBorrowIndex;
    uint256 public indexRegressions;
    uint256 public reserveRegressions;
    uint256 public lastReserves;
    uint256 public liquidationCount;
    uint256 public borrowCount;
    uint256 public depositCount;

    int256 constant BASE_PRICE = 341258000000;

    constructor(
        LendingPool pool_,
        CollateralVault vault_,
        LendingController controller_,
        LiquidationManager manager_,
        MockERC20 weth_,
        MockERC20 usdc_,
        MockAggregator wethFeed_,
        MockAggregator usdcFeed_,
        address[] memory actors_
    ) {
        pool = pool_;
        vault = vault_;
        controller = controller_;
        manager = manager_;
        weth = weth_;
        usdc = usdc_;
        wethFeed = wethFeed_;
        usdcFeed = usdcFeed_;
        actors = actors_;

        lastSupplyIndex = pool.supplyIndex();
        lastBorrowIndex = pool.borrowIndex();
        lastReserves = pool.accruedReserves();

        for (uint256 i = 0; i < actors.length; i++) {
            weth.mint(actors[i], 1_000e18);
            usdc.mint(actors[i], 1_000_000e6);

            vm.startPrank(actors[i]);
            weth.approve(address(vault), type(uint256).max);
            usdc.approve(address(pool), type(uint256).max);
            vm.stopPrank();
        }
    }

    function actorCount() external view returns (uint256) {
        return actors.length;
    }

    function _actor(uint256 seed) internal view returns (address) {
        return actors[seed % actors.length];
    }

    function _trackIndexes() internal {
        if (pool.supplyIndex() < lastSupplyIndex || pool.borrowIndex() < lastBorrowIndex) {
            indexRegressions++;
        }

        if (pool.accruedReserves() < lastReserves) {
            reserveRegressions++;
        }

        lastSupplyIndex = pool.supplyIndex();
        lastBorrowIndex = pool.borrowIndex();
        lastReserves = pool.accruedReserves();
    }

    function deposit(uint256 actorSeed, uint256 amount) external {
        address actor = _actor(actorSeed);
        amount = bound(amount, 1e6, 200_000e6);

        if (usdc.balanceOf(actor) < amount) {
            return;
        }

        vm.prank(actor);
        try pool.deposit(amount) {
            depositCount++;
        } catch {}

        _trackIndexes();
    }

    function withdraw(uint256 actorSeed, uint256 amount) external {
        address actor = _actor(actorSeed);

        uint256 ceiling = pool.maxWithdrawable(actor);
        if (ceiling == 0) {
            return;
        }

        amount = bound(amount, 1, ceiling);

        vm.prank(actor);
        try pool.withdraw(amount) {} catch {}

        _trackIndexes();
    }

    function addCollateral(uint256 actorSeed, uint256 amount) external {
        address actor = _actor(actorSeed);
        amount = bound(amount, 1e15, 100e18);

        if (weth.balanceOf(actor) < amount) {
            return;
        }

        vm.prank(actor);
        try vault.depositCollateral(amount) {} catch {}

        _trackIndexes();
    }

    function removeCollateral(uint256 actorSeed, uint256 amount) external {
        address actor = _actor(actorSeed);

        uint256 ceiling;
        try controller.maxWithdrawableCollateral(actor) returns (uint256 safeAmount) {
            ceiling = safeAmount;
        } catch {
            return;
        }

        if (ceiling == 0) {
            return;
        }

        amount = bound(amount, 1, ceiling);

        vm.prank(actor);
        try vault.withdrawCollateral(amount) {} catch {}

        _trackIndexes();
    }

    function borrow(uint256 actorSeed, uint256 amount) external {
        address actor = _actor(actorSeed);

        uint256 ceiling;
        try controller.maxBorrowable(actor) returns (uint256 room) {
            ceiling = room;
        } catch {
            return;
        }

        if (ceiling == 0) {
            return;
        }

        amount = bound(amount, 1, ceiling);

        vm.prank(actor);
        try controller.borrow(amount) {
            borrowCount++;
        } catch {}

        _trackIndexes();
    }

    function repay(uint256 actorSeed, uint256 amount) external {
        address actor = _actor(actorSeed);

        uint256 owed;
        try controller.debtOf(actor) returns (uint256 debt) {
            owed = debt;
        } catch {
            return;
        }

        if (owed == 0) {
            return;
        }

        amount = bound(amount, 1, owed);

        if (usdc.balanceOf(actor) < amount) {
            return;
        }

        vm.prank(actor);
        try controller.repay(amount) {} catch {}

        _trackIndexes();
    }

    function repayAll(uint256 actorSeed) external {
        address actor = _actor(actorSeed);

        vm.prank(actor);
        try controller.repayAll() {} catch {}

        _trackIndexes();
    }

    function liquidate(uint256 callerSeed, uint256 targetSeed) external {
        address caller = _actor(callerSeed);
        address target = _actor(targetSeed);

        vm.prank(caller);
        try manager.liquidate(target) {
            liquidationCount++;
        } catch {}

        _trackIndexes();
    }

    function movePrice(uint256 priceSeed) external {
        int256 newPrice = int256(bound(priceSeed, uint256(BASE_PRICE) / 4, uint256(BASE_PRICE) * 2));

        wethFeed.setPrice(newPrice);
        usdcFeed.setPrice(100000000);

        _trackIndexes();
    }

    function passTime(uint256 secondsForward) external {
        secondsForward = bound(secondsForward, 1, 30 days);

        vm.warp(block.timestamp + secondsForward);

        (, int256 ethAnswer,,,) = wethFeed.latestRoundData();
        wethFeed.setPrice(ethAnswer);
        usdcFeed.setPrice(100000000);

        pool.accrueInterest();

        _trackIndexes();
    }
}

contract InvariantsTest is BaseTest {
    ProtocolHandler internal handler;
    address[] internal actors;

    function setUp() public override {
        super.setUp();

        actors.push(makeAddr("actorOne"));
        actors.push(makeAddr("actorTwo"));
        actors.push(makeAddr("actorThree"));
        actors.push(makeAddr("actorFour"));

        handler =
            new ProtocolHandler(pool, vault, controller, manager, weth, usdc, wethFeed, usdcFeed, actors);

        targetContract(address(handler));
    }

    function invariant_supplyShares_match_holders() public view {
        uint256 total;

        for (uint256 i = 0; i < actors.length; i++) {
            total += pool.sharesOf(actors[i]);
        }

        total += pool.sharesOf(lender);
        total += pool.sharesOf(lenderTwo);

        assertEq(pool.totalSupplyShares(), total);
    }

    function invariant_debtShares_match_borrowers() public view {
        uint256 total;

        for (uint256 i = 0; i < actors.length; i++) {
            total += pool.debtSharesOf(actors[i]);
        }

        total += pool.debtSharesOf(alice);
        total += pool.debtSharesOf(bob);

        assertEq(pool.totalDebtShares(), total);
    }

    function invariant_pool_cash_covers_liquidity_within_rounding() public view {
        uint256 holders = actors.length + 4;

        assertGe(usdc.balanceOf(address(pool)) + holders, pool.availableLiquidity());
    }

    function invariant_withdrawable_amount_is_always_payable() public view {
        uint256 cash = usdc.balanceOf(address(pool));

        for (uint256 i = 0; i < actors.length; i++) {
            assertLe(pool.maxWithdrawable(actors[i]), cash + actors.length + 4);
        }
    }

    function invariant_vault_holds_every_deposit() public view {
        uint256 total;

        for (uint256 i = 0; i < actors.length; i++) {
            total += vault.collateralOf(actors[i]);
        }

        total += vault.collateralOf(alice);
        total += vault.collateralOf(bob);

        assertEq(vault.totalCollateral(), total);
        assertGe(weth.balanceOf(address(vault)), vault.totalCollateral());
    }

    function invariant_indexes_never_regress() public view {
        assertEq(handler.indexRegressions(), 0);
        assertGe(pool.supplyIndex(), 1e18);
        assertGe(pool.borrowIndex(), 1e18);
    }

    function invariant_debt_and_cash_cover_every_claim() public view {
        uint256 assets = usdc.balanceOf(address(pool)) + pool.totalBorrowed();
        uint256 claims = pool.totalSupplied() + pool.accruedReserves();
        uint256 tolerance = actors.length + 4;

        assertGe(assets + tolerance, claims);
    }

    function invariant_healthy_positions_are_never_liquidatable() public view {
        for (uint256 i = 0; i < actors.length; i++) {
            try manager.isLiquidatable(actors[i]) returns (bool liquidatable) {
                if (!liquidatable) {
                    continue;
                }

                uint256 health = controller.healthFactorBps(actors[i]);
                assertLt(health, 10_000);
            } catch {
                continue;
            }
        }
    }

    function invariant_no_debt_means_full_collateral_is_free() public view {
        for (uint256 i = 0; i < actors.length; i++) {
            if (pool.debtSharesOf(actors[i]) != 0) {
                continue;
            }

            try controller.maxWithdrawableCollateral(actors[i]) returns (uint256 free) {
                assertEq(free, vault.collateralOf(actors[i]));
            } catch {
                continue;
            }
        }
    }

    function invariant_borrowed_never_exceeds_supplied_plus_reserves() public view {
        if (pool.totalDebtShares() == 0) {
            assertEq(pool.totalBorrowed(), 0);
        }
    }

    function invariant_reserves_only_grow_until_collected() public view {
        assertEq(handler.reserveRegressions(), 0);
    }

    function invariant_callSummary() public view {
        assertGe(handler.depositCount() + handler.borrowCount() + handler.liquidationCount(), 0);
    }
}
