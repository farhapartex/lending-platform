// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";

import {HealthMath} from "../../src/libraries/HealthMath.sol";

contract HealthMathTest is Test {
    uint256 constant ETH_PRICE = 341258000000;
    uint256 constant USDC_PRICE = 100000000;

    function test_valueOfAmount_matches_frontend() public pure {
        assertEq(HealthMath.valueOfAmount(3.2e18, 18, ETH_PRICE), 1092025600000);
        assertEq(HealthMath.valueOfAmount(6900e6, 6, USDC_PRICE), 690000000000);
    }

    function test_valueOfAmount_zero_cases() public pure {
        assertEq(HealthMath.valueOfAmount(0, 18, ETH_PRICE), 0);
        assertEq(HealthMath.valueOfAmount(3.2e18, 18, 0), 0);
    }

    function test_amountFromValue_zero_price_returns_zero_not_revert() public pure {
        assertEq(HealthMath.amountFromValueDown(1e18, 18, 0), 0);
        assertEq(HealthMath.amountFromValueUp(1e18, 18, 0), 0);
    }

    function test_healthFactor_matches_rendered_ui() public pure {
        assertEq(HealthMath.healthFactor(1092025600000, 690000000000, 8000), 12661);
    }

    function test_healthFactor_no_debt_is_sentinel() public pure {
        assertEq(HealthMath.healthFactor(1092025600000, 0, 8000), type(uint256).max);
        assertEq(HealthMath.healthFactor(0, 0, 8000), type(uint256).max);
    }

    function test_healthFactor_zero_collateral_with_debt_is_zero() public pure {
        assertEq(HealthMath.healthFactor(0, 690000000000, 8000), 0);
    }

    function test_isLiquidatable_boundary_belongs_to_borrower() public pure {
        assertFalse(HealthMath.isLiquidatable(10_000));
        assertTrue(HealthMath.isLiquidatable(9_999));
        assertFalse(HealthMath.isLiquidatable(10_001));
        assertFalse(HealthMath.isLiquidatable(type(uint256).max));
        assertTrue(HealthMath.isLiquidatable(0));
    }

    function test_borrow_and_liquidation_limits_differ() public pure {
        uint256 collateralValue = 1092025600000;

        assertEq(HealthMath.borrowLimitValue(collateralValue, 7500), 819019200000);
        assertEq(HealthMath.liquidationLimitValue(collateralValue, 8000), 873620480000);
    }

    function test_remainingBorrowAmount_matches_contract_run() public pure {
        uint256 remaining = HealthMath.remainingBorrowAmount(1092025600000, 690000000000, 7500, 6, USDC_PRICE);

        assertEq(remaining, 1290192000);
    }

    function test_remainingBorrowAmount_zero_when_at_or_over_limit() public pure {
        assertEq(HealthMath.remainingBorrowAmount(1092025600000, 819019200000, 7500, 6, USDC_PRICE), 0);
        assertEq(HealthMath.remainingBorrowAmount(1092025600000, 900000000000, 7500, 6, USDC_PRICE), 0);
        assertEq(HealthMath.remainingBorrowAmount(0, 1, 7500, 6, USDC_PRICE), 0);
    }

    function test_collateralValueNeeded_rounds_up() public pure {
        assertEq(HealthMath.collateralValueNeeded(690000000000, 7500), 920000000000);
        assertEq(HealthMath.collateralValueNeeded(0, 7500), 0);
        assertEq(HealthMath.collateralValueNeeded(1, 7500), 2);
    }

    function test_freeCollateralAmount_matches_contract_run() public pure {
        uint256 free = HealthMath.freeCollateralAmount(3.2e18, 18, ETH_PRICE, 690000000000, 7500);

        assertEq(free, 504092504791096472);
    }

    function test_freeCollateralAmount_no_debt_releases_everything() public pure {
        assertEq(HealthMath.freeCollateralAmount(3.2e18, 18, ETH_PRICE, 0, 7500), 3.2e18);
    }

    function test_freeCollateralAmount_floors_at_zero_when_underwater() public pure {
        assertEq(HealthMath.freeCollateralAmount(1e18, 18, ETH_PRICE, 10_000_000e8, 7500), 0);
    }

    function testFuzz_healthFactor_never_reverts(
        uint256 collateralValue,
        uint256 debtValue,
        uint256 thresholdBps
    ) public pure {
        collateralValue = bound(collateralValue, 0, 1e40);
        debtValue = bound(debtValue, 0, 1e40);
        thresholdBps = bound(thresholdBps, 1, 10_000);

        uint256 health = HealthMath.healthFactor(collateralValue, debtValue, thresholdBps);

        if (debtValue == 0) {
            assertEq(health, type(uint256).max);
        }
    }

    function testFuzz_more_collateral_never_lowers_health(
        uint256 collateralValue,
        uint256 extra,
        uint256 debtValue
    ) public pure {
        collateralValue = bound(collateralValue, 0, 1e30);
        extra = bound(extra, 0, 1e30);
        debtValue = bound(debtValue, 1, 1e30);

        uint256 before = HealthMath.healthFactor(collateralValue, debtValue, 8000);
        uint256 improved = HealthMath.healthFactor(collateralValue + extra, debtValue, 8000);

        assertGe(improved, before);
    }

    function testFuzz_more_debt_never_raises_health(uint256 collateralValue, uint256 debtValue, uint256 extra)
        public
        pure
    {
        collateralValue = bound(collateralValue, 0, 1e30);
        debtValue = bound(debtValue, 1, 1e30);
        extra = bound(extra, 0, 1e30);

        uint256 before = HealthMath.healthFactor(collateralValue, debtValue, 8000);
        uint256 worse = HealthMath.healthFactor(collateralValue, debtValue + extra, 8000);

        assertLe(worse, before);
    }

    function testFuzz_borrowing_the_maximum_leaves_position_safe(uint256 collateralAmount, uint256 price)
        public
        pure
    {
        collateralAmount = bound(collateralAmount, 1e12, 1e24);
        price = bound(price, 1e6, 1e14);

        uint256 collateralValue = HealthMath.valueOfAmount(collateralAmount, 18, price);
        uint256 borrowable = HealthMath.remainingBorrowAmount(collateralValue, 0, 7500, 6, USDC_PRICE);

        if (borrowable == 0) {
            return;
        }

        uint256 debtValue = HealthMath.valueOfAmount(borrowable, 6, USDC_PRICE);
        uint256 health = HealthMath.healthFactor(collateralValue, debtValue, 8000);

        assertFalse(HealthMath.isLiquidatable(health));
    }

    function testFuzz_free_collateral_never_exceeds_held(
        uint256 collateralAmount,
        uint256 debtValue,
        uint256 price
    ) public pure {
        collateralAmount = bound(collateralAmount, 0, 1e24);
        debtValue = bound(debtValue, 0, 1e20);
        price = bound(price, 1e6, 1e14);

        uint256 free = HealthMath.freeCollateralAmount(collateralAmount, 18, price, debtValue, 7500);

        assertLe(free, collateralAmount);
    }

    function testFuzz_withdrawing_free_collateral_keeps_position_safe(
        uint256 collateralAmount,
        uint256 debtAmount,
        uint256 price
    ) public pure {
        collateralAmount = bound(collateralAmount, 1e15, 1e24);
        debtAmount = bound(debtAmount, 1e6, 1e14);
        price = bound(price, 1e8, 1e14);

        uint256 debtValue = HealthMath.valueOfAmount(debtAmount, 6, USDC_PRICE);
        uint256 startingValue = HealthMath.valueOfAmount(collateralAmount, 18, price);

        if (HealthMath.isLiquidatable(HealthMath.healthFactor(startingValue, debtValue, 8000))) {
            return;
        }

        uint256 free = HealthMath.freeCollateralAmount(collateralAmount, 18, price, debtValue, 7500);

        uint256 remainingValue = HealthMath.valueOfAmount(collateralAmount - free, 18, price);
        uint256 health = HealthMath.healthFactor(remainingValue, debtValue, 8000);

        assertFalse(HealthMath.isLiquidatable(health));
    }

    function testFuzz_value_roundtrip_loses_at_most_one_unit(uint256 amount, uint256 price) public pure {
        amount = bound(amount, 0, 1e24);
        price = bound(price, 1e6, 1e14);

        uint256 value = HealthMath.valueOfAmount(amount, 18, price);
        uint256 back = HealthMath.amountFromValueUp(value, 18, price);

        assertLe(back, amount + 1);
    }
}
