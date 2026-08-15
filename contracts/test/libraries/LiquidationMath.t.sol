// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";

import {HealthMath} from "../../src/libraries/HealthMath.sol";
import {LiquidationMath} from "../../src/libraries/LiquidationMath.sol";

contract LiquidationMathTest is Test {
    function test_bonusValue_is_five_percent() public pure {
        assertEq(LiquidationMath.bonusValue(690000000000, 500), 34500000000);
        assertEq(LiquidationMath.bonusValue(0, 500), 0);
        assertEq(LiquidationMath.bonusValue(690000000000, 0), 0);
    }

    function test_rewardValue_is_debt_plus_bonus() public pure {
        assertEq(LiquidationMath.rewardValue(690000000000, 500), 724500000000);
    }

    function test_planSeize_normal_case_matches_contract_run() public pure {
        LiquidationMath.SeizePlan memory plan =
            LiquidationMath.planSeize(690000000000, 832000000000, 3.2e18, 18, 260000000000, 500);

        assertEq(plan.collateralToSeize, 2786538461538461539);
        assertEq(plan.bonusValue, 34500000000);
        assertEq(plan.shortfall, 0);
    }

    function test_planSeize_shortfall_case_matches_contract_run() public pure {
        LiquidationMath.SeizePlan memory plan =
            LiquidationMath.planSeize(690000000000, 640000000000, 3.2e18, 18, 200000000000, 500);

        assertEq(plan.collateralToSeize, 3.2e18);
        assertEq(plan.bonusValue, 34500000000);
        assertEq(plan.shortfall, 84500000000);
    }

    function test_planSeize_exact_boundary_has_no_shortfall() public pure {
        uint256 price = 100000000;
        uint256 debtValue = 1000e8;
        uint256 target = 1050e8;
        uint256 held = 1050e18;

        LiquidationMath.SeizePlan memory plan =
            LiquidationMath.planSeize(debtValue, target, held, 18, price, 500);

        assertEq(plan.collateralToSeize, held);
        assertEq(plan.shortfall, 0);
    }

    function test_planSeize_zero_debt() public pure {
        LiquidationMath.SeizePlan memory plan =
            LiquidationMath.planSeize(0, 832000000000, 3.2e18, 18, 260000000000, 500);

        assertEq(plan.collateralToSeize, 0);
        assertEq(plan.bonusValue, 0);
        assertEq(plan.shortfall, 0);
    }

    function test_planSeize_zero_collateral_is_all_shortfall() public pure {
        LiquidationMath.SeizePlan memory plan =
            LiquidationMath.planSeize(690000000000, 0, 0, 18, 260000000000, 500);

        assertEq(plan.collateralToSeize, 0);
        assertEq(plan.shortfall, 724500000000);
    }

    function testFuzz_never_seizes_more_than_held(
        uint256 debtValue,
        uint256 held,
        uint256 price,
        uint256 bonusBps
    ) public pure {
        debtValue = bound(debtValue, 0, 1e20);
        held = bound(held, 0, 1e24);
        price = bound(price, 1e6, 1e14);
        bonusBps = bound(bonusBps, 0, 2000);

        uint256 collateralValue = HealthMath.valueOfAmount(held, 18, price);

        LiquidationMath.SeizePlan memory plan =
            LiquidationMath.planSeize(debtValue, collateralValue, held, 18, price, bonusBps);

        assertLe(plan.collateralToSeize, held);
    }

    function testFuzz_shortfall_only_when_collateral_insufficient(
        uint256 debtValue,
        uint256 held,
        uint256 price,
        uint256 bonusBps
    ) public pure {
        debtValue = bound(debtValue, 0, 1e20);
        held = bound(held, 0, 1e24);
        price = bound(price, 1e6, 1e14);
        bonusBps = bound(bonusBps, 0, 2000);

        uint256 collateralValue = HealthMath.valueOfAmount(held, 18, price);

        LiquidationMath.SeizePlan memory plan =
            LiquidationMath.planSeize(debtValue, collateralValue, held, 18, price, bonusBps);

        if (plan.shortfall > 0) {
            assertEq(plan.collateralToSeize, held);
        }
    }

    function testFuzz_liquidator_is_never_short_changed_when_solvent(
        uint256 debtValue,
        uint256 held,
        uint256 price
    ) public pure {
        debtValue = bound(debtValue, 1e8, 1e18);
        price = bound(price, 1e8, 1e14);
        held = bound(held, 1e18, 1e24);

        uint256 collateralValue = HealthMath.valueOfAmount(held, 18, price);

        LiquidationMath.SeizePlan memory plan =
            LiquidationMath.planSeize(debtValue, collateralValue, held, 18, price, 500);

        if (plan.shortfall != 0) {
            return;
        }

        uint256 seizedValue = HealthMath.valueOfAmount(plan.collateralToSeize, 18, price);

        assertGe(seizedValue, debtValue);
    }

    function testFuzz_bonus_never_exceeds_debt(uint256 debtValue, uint256 bonusBps) public pure {
        debtValue = bound(debtValue, 0, 1e30);
        bonusBps = bound(bonusBps, 0, 10_000);

        assertLe(LiquidationMath.bonusValue(debtValue, bonusBps), debtValue);
    }

    function testFuzz_higher_bonus_seizes_more(uint256 debtValue, uint256 held, uint256 price) public pure {
        debtValue = bound(debtValue, 1e8, 1e16);
        price = bound(price, 1e8, 1e14);
        held = bound(held, 1e20, 1e24);

        uint256 collateralValue = HealthMath.valueOfAmount(held, 18, price);

        LiquidationMath.SeizePlan memory low =
            LiquidationMath.planSeize(debtValue, collateralValue, held, 18, price, 500);
        LiquidationMath.SeizePlan memory high =
            LiquidationMath.planSeize(debtValue, collateralValue, held, 18, price, 1000);

        assertGe(high.collateralToSeize, low.collateralToSeize);
    }
}
