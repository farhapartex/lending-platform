// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";

import {WadMath} from "../../src/libraries/WadMath.sol";

contract WadMathTest is Test {
    function test_constants() public pure {
        assertEq(WadMath.WAD, 1e18);
        assertEq(WadMath.FULL_PERCENT_BPS, 10_000);
        assertEq(WadMath.PRICE_UNIT, 1e8);
        assertEq(WadMath.SECONDS_PER_YEAR, 365 days);
    }

    function test_mulDown_truncates() public pure {
        assertEq(WadMath.mulDown(3, 1, 2), 1);
        assertEq(WadMath.mulDown(999, 1, 1000), 0);
        assertEq(WadMath.mulDown(10, 3, 4), 7);
    }

    function test_mulUp_rounds_away_from_zero() public pure {
        assertEq(WadMath.mulUp(3, 1, 2), 2);
        assertEq(WadMath.mulUp(999, 1, 1000), 1);
        assertEq(WadMath.mulUp(10, 3, 4), 8);
    }

    function test_exact_division_has_no_rounding_gap() public pure {
        assertEq(WadMath.mulDown(10, 2, 4), WadMath.mulUp(10, 2, 4));
        assertEq(WadMath.mulDown(100, 1, 10), 10);
        assertEq(WadMath.mulUp(100, 1, 10), 10);
    }

    function test_takeBps() public pure {
        assertEq(WadMath.takeBpsDown(1092025600000, 7500), 819019200000);
        assertEq(WadMath.takeBpsDown(690000000000, 500), 34500000000);
        assertEq(WadMath.takeBpsDown(1, 5000), 0);
        assertEq(WadMath.takeBpsUp(1, 5000), 1);
    }

    function test_scaleUpByBps_is_inverse_of_takeBps() public pure {
        assertEq(WadMath.scaleUpByBpsUp(690000000000, 7500), 920000000000);
        assertEq(WadMath.takeBpsDown(920000000000, 7500), 690000000000);
    }

    function test_wadMul_and_wadDiv() public pure {
        assertEq(WadMath.wadMulDown(6900e6, 21280000000000000), 146832000);
        assertEq(WadMath.wadMulDown(100, WadMath.WAD), 100);
        assertEq(WadMath.wadDivDown(100, WadMath.WAD), 100);
    }

    function test_smaller_larger() public pure {
        assertEq(WadMath.smaller(3, 7), 3);
        assertEq(WadMath.smaller(7, 3), 3);
        assertEq(WadMath.smaller(5, 5), 5);
        assertEq(WadMath.larger(3, 7), 7);
        assertEq(WadMath.larger(7, 3), 7);
    }

    function test_subtractOrZero_floors_at_zero() public pure {
        assertEq(WadMath.subtractOrZero(10, 3), 7);
        assertEq(WadMath.subtractOrZero(3, 10), 0);
        assertEq(WadMath.subtractOrZero(5, 5), 0);
        assertEq(WadMath.subtractOrZero(0, type(uint256).max), 0);
    }

    function test_mulDiv_survives_intermediate_overflow() public pure {
        uint256 huge = type(uint128).max;

        assertEq(WadMath.mulDown(huge, huge, huge), huge);
    }

    function testFuzz_up_never_below_down(uint256 value, uint256 multiplier, uint256 divisor) public pure {
        value = bound(value, 0, 1e30);
        multiplier = bound(multiplier, 0, 1e30);
        divisor = bound(divisor, 1, 1e30);

        uint256 down = WadMath.mulDown(value, multiplier, divisor);
        uint256 up = WadMath.mulUp(value, multiplier, divisor);

        assertGe(up, down);
        assertLe(up - down, 1);
    }

    function testFuzz_rounding_gap_only_when_inexact(uint256 value, uint256 divisor) public pure {
        value = bound(value, 0, 1e40);
        divisor = bound(divisor, 1, 1e40);

        uint256 down = WadMath.mulDown(value, 1, divisor);
        uint256 up = WadMath.mulUp(value, 1, divisor);

        if (value % divisor == 0) {
            assertEq(down, up);
        } else {
            assertEq(up, down + 1);
        }
    }

    function testFuzz_takeBps_never_exceeds_value(uint256 value, uint256 bps) public pure {
        value = bound(value, 0, 1e40);
        bps = bound(bps, 0, WadMath.FULL_PERCENT_BPS);

        assertLe(WadMath.takeBpsDown(value, bps), value);
    }

    function testFuzz_takeBps_monotonic_in_bps(uint256 value, uint256 lowBps, uint256 highBps) public pure {
        value = bound(value, 0, 1e40);
        lowBps = bound(lowBps, 0, WadMath.FULL_PERCENT_BPS);
        highBps = bound(highBps, lowBps, WadMath.FULL_PERCENT_BPS);

        assertLe(WadMath.takeBpsDown(value, lowBps), WadMath.takeBpsDown(value, highBps));
    }

    function testFuzz_subtractOrZero_never_reverts(uint256 value, uint256 amountToRemove) public pure {
        uint256 result = WadMath.subtractOrZero(value, amountToRemove);

        if (value > amountToRemove) {
            assertEq(result, value - amountToRemove);
        } else {
            assertEq(result, 0);
        }
    }

    function testFuzz_smaller_larger_ordering(uint256 left, uint256 right) public pure {
        assertLe(WadMath.smaller(left, right), WadMath.larger(left, right));
    }

    function testFuzz_wad_roundtrip_within_one(uint256 value) public pure {
        value = bound(value, 0, 1e30);

        uint256 scaled = WadMath.wadMulDown(value, WadMath.WAD);

        assertEq(scaled, value);
    }
}
