// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

import {InterestRateModel} from "../src/InterestRateModel.sol";
import {RateCurve} from "../src/interfaces/IInterestRateModel.sol";
import {Errors} from "../src/libraries/Errors.sol";

contract InterestRateModelTest is Test {
    uint256 constant YEAR = 365 days;
    uint16 constant KINK = 8000;

    InterestRateModel model;

    address admin = makeAddr("admin");
    address stranger = makeAddr("stranger");

    function setUp() public {
        model = new InterestRateModel(admin, _curve());
    }

    function _curve() internal pure returns (RateCurve memory) {
        return RateCurve({
            baseRatePerSecond: uint64(uint256(1e16) / YEAR),
            slopeBelowKinkPerSecond: uint64(uint256(654e14) / YEAR),
            slopeAboveKinkPerSecond: uint64(uint256(2e18) / YEAR),
            kinkUtilizationBps: KINK
        });
    }

    function test_curve_is_stored() public view {
        RateCurve memory stored = model.curve();

        assertEq(stored.kinkUtilizationBps, KINK);
        assertEq(stored.baseRatePerSecond, uint64(uint256(1e16) / YEAR));
    }

    function test_constructor_rejects_degenerate_kink() public {
        RateCurve memory bad = _curve();

        bad.kinkUtilizationBps = 0;
        vm.expectRevert(Errors.InvalidRiskSettings.selector);
        new InterestRateModel(admin, bad);

        bad.kinkUtilizationBps = 10_000;
        vm.expectRevert(Errors.InvalidRiskSettings.selector);
        new InterestRateModel(admin, bad);

        bad.kinkUtilizationBps = 10_001;
        vm.expectRevert(Errors.InvalidRiskSettings.selector);
        new InterestRateModel(admin, bad);
    }

    function test_setCurve_is_owner_only() public {
        vm.prank(stranger);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, stranger));
        model.setCurve(_curve());
    }

    function test_setCurve_validates() public {
        RateCurve memory bad = _curve();
        bad.kinkUtilizationBps = 0;

        vm.prank(admin);
        vm.expectRevert(Errors.InvalidRiskSettings.selector);
        model.setCurve(bad);
    }

    function test_utilization_guards() public view {
        assertEq(model.utilizationBps(0, 0), 0);
        assertEq(model.utilizationBps(0, 100), 0);
        assertEq(model.utilizationBps(100, 0), 0);
    }

    function test_utilization_is_clamped_to_full() public view {
        assertEq(model.utilizationBps(100, 200), 10_000);
        assertEq(model.utilizationBps(100, 100), 10_000);
    }

    function test_utilization_matches_observed_values() public view {
        assertEq(model.utilizationBps(48_240_000, 31_580_000), 6546);
        assertEq(model.utilizationBps(50_000e6, 6_900e6), 1380);
        assertEq(model.utilizationBps(150_000e6, 15_000e6), 1000);
    }

    function test_borrow_apr_curve_points() public view {
        assertEq(model.borrowAprBps(0), 99);
        assertEq(model.borrowAprBps(1380), 212);
        assertEq(model.borrowAprBps(6546), 635);
        assertEq(model.borrowAprBps(KINK), 753);
        assertEq(model.borrowAprBps(9000), 10753);
        assertEq(model.borrowAprBps(10_000), 20753);
    }

    function test_supply_apr_is_diluted_by_utilization_and_reserves() public view {
        assertEq(model.supplyAprBps(0, 1000), 0);
        assertEq(model.supplyAprBps(6546, 1000), 374);
    }

    function test_supply_rate_never_exceeds_borrow_rate() public view {
        uint256 borrowRate = model.borrowRatePerSecond(KINK);
        uint256 supplyRate = model.supplyRatePerSecond(KINK, 0);

        assertLt(supplyRate, borrowRate);
    }

    function test_curve_is_continuous_at_kink() public view {
        uint256 justBelow = model.borrowRatePerSecond(KINK - 1);
        uint256 atKink = model.borrowRatePerSecond(KINK);
        uint256 justAbove = model.borrowRatePerSecond(KINK + 1);

        assertLe(justBelow, atKink);
        assertLe(atKink, justAbove);
        assertLt(justAbove - atKink, atKink);
    }

    function test_slope_above_kink_is_steeper() public view {
        uint256 belowSlope = model.borrowRatePerSecond(KINK) - model.borrowRatePerSecond(KINK - 100);
        uint256 aboveSlope = model.borrowRatePerSecond(KINK + 100) - model.borrowRatePerSecond(KINK);

        assertGt(aboveSlope, belowSlope);
    }

    function test_zero_reserve_factor_gives_lenders_everything() public view {
        uint256 withReserve = model.supplyRatePerSecond(5000, 1000);
        uint256 withoutReserve = model.supplyRatePerSecond(5000, 0);

        assertGt(withoutReserve, withReserve);
    }

    function testFuzz_borrow_rate_monotonic(uint256 lowUsage, uint256 highUsage) public view {
        lowUsage = bound(lowUsage, 0, 10_000);
        highUsage = bound(highUsage, lowUsage, 10_000);

        assertLe(model.borrowRatePerSecond(lowUsage), model.borrowRatePerSecond(highUsage));
    }

    function testFuzz_borrow_rate_never_below_base(uint256 usage) public view {
        usage = bound(usage, 0, 10_000);

        assertGe(model.borrowRatePerSecond(usage), model.curve().baseRatePerSecond);
    }

    function testFuzz_supply_rate_never_exceeds_borrow_rate(uint256 usage, uint256 reserveFactorBps)
        public
        view
    {
        usage = bound(usage, 0, 10_000);
        reserveFactorBps = bound(reserveFactorBps, 0, 5000);

        assertLe(model.supplyRatePerSecond(usage, reserveFactorBps), model.borrowRatePerSecond(usage));
    }

    function testFuzz_higher_reserve_factor_lowers_supply_rate(
        uint256 usage,
        uint256 lowFactor,
        uint256 highFactor
    ) public view {
        usage = bound(usage, 1, 10_000);
        lowFactor = bound(lowFactor, 0, 5000);
        highFactor = bound(highFactor, lowFactor, 5000);

        assertGe(model.supplyRatePerSecond(usage, lowFactor), model.supplyRatePerSecond(usage, highFactor));
    }

    function testFuzz_utilization_never_exceeds_full(uint256 supplied, uint256 borrowed) public view {
        supplied = bound(supplied, 0, 1e30);
        borrowed = bound(borrowed, 0, 1e30);

        assertLe(model.utilizationBps(supplied, borrowed), 10_000);
    }

    function testFuzz_utilization_monotonic_in_borrowed(
        uint256 supplied,
        uint256 lowBorrow,
        uint256 highBorrow
    ) public view {
        supplied = bound(supplied, 1, 1e30);
        lowBorrow = bound(lowBorrow, 0, 1e30);
        highBorrow = bound(highBorrow, lowBorrow, 1e30);

        assertLe(model.utilizationBps(supplied, lowBorrow), model.utilizationBps(supplied, highBorrow));
    }

    function testFuzz_any_valid_curve_is_accepted(
        uint64 baseRate,
        uint64 slopeBelow,
        uint64 slopeAbove,
        uint16 kink
    ) public {
        kink = uint16(bound(kink, 1, 9999));

        RateCurve memory candidate = RateCurve({
            baseRatePerSecond: baseRate,
            slopeBelowKinkPerSecond: slopeBelow,
            slopeAboveKinkPerSecond: slopeAbove,
            kinkUtilizationBps: kink
        });

        vm.prank(admin);
        model.setCurve(candidate);

        assertEq(model.curve().kinkUtilizationBps, kink);
        assertGe(model.borrowRatePerSecond(0), baseRate);
    }
}
