// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";

import {InterestMath} from "../../src/libraries/InterestMath.sol";

contract InterestMathTest is Test {
    uint256 constant YEAR = 365 days;

    function test_growthFactor_is_rate_times_time() public pure {
        assertEq(InterestMath.growthFactor(674831937, YEAR), 674831937 * YEAR);
        assertEq(InterestMath.growthFactor(0, YEAR), 0);
        assertEq(InterestMath.growthFactor(674831937, 0), 0);
    }

    function test_interestOnDebt_matches_observed_run() public pure {
        uint256 growth = InterestMath.growthFactor(674831937, YEAR);

        assertEq(InterestMath.interestOnDebt(6900e6, growth), 146842350);
    }

    function test_interestOnDebt_rounds_up() public pure {
        assertEq(InterestMath.interestOnDebt(1, 1), 1);
        assertEq(InterestMath.interestOnDebt(0, 1e18), 0);
    }

    function test_splitForReserves_is_exact() public pure {
        (uint256 reserveShare, uint256 lendersShare) = InterestMath.splitForReserves(146842350, 1000);

        assertEq(reserveShare, 14684235);
        assertEq(lendersShare, 132158115);
        assertEq(reserveShare + lendersShare, 146842350);
    }

    function test_splitForReserves_edges() public pure {
        (uint256 zeroReserve, uint256 allLenders) = InterestMath.splitForReserves(1000, 0);
        assertEq(zeroReserve, 0);
        assertEq(allLenders, 1000);

        (uint256 halfReserve, uint256 halfLenders) = InterestMath.splitForReserves(1000, 5000);
        assertEq(halfReserve, 500);
        assertEq(halfLenders, 500);

        (uint256 dustReserve, uint256 dustLenders) = InterestMath.splitForReserves(1, 1000);
        assertEq(dustReserve, 0);
        assertEq(dustLenders, 1);
    }

    function test_grownBorrowIndex_matches_observed_run() public pure {
        uint256 growth = InterestMath.growthFactor(674831937, YEAR);

        assertEq(InterestMath.grownBorrowIndex(1e18, growth), 1021281499965232000);
    }

    function test_grownSupplyIndex_matches_observed_run() public pure {
        assertEq(InterestMath.grownSupplyIndex(1e18, 132158115, 50_000e6), 1002643162300000000);
    }

    function test_grownSupplyIndex_guards() public pure {
        assertEq(InterestMath.grownSupplyIndex(1e18, 100, 0), 1e18);
        assertEq(InterestMath.grownSupplyIndex(1e18, 0, 50_000e6), 1e18);
    }

    function test_supply_index_grows_slower_than_borrow_index() public pure {
        uint256 growth = InterestMath.growthFactor(674831937, YEAR);
        uint256 interest = InterestMath.interestOnDebt(6900e6, growth);
        (, uint256 lendersShare) = InterestMath.splitForReserves(interest, 1000);

        uint256 borrowIndex = InterestMath.grownBorrowIndex(1e18, growth);
        uint256 supplyIndex = InterestMath.grownSupplyIndex(1e18, lendersShare, 50_000e6);

        assertGt(borrowIndex, supplyIndex);
    }

    function testFuzz_split_always_sums_to_whole(uint256 interest, uint256 reserveFactorBps) public pure {
        interest = bound(interest, 0, 1e30);
        reserveFactorBps = bound(reserveFactorBps, 0, 5000);

        (uint256 reserveShare, uint256 lendersShare) =
            InterestMath.splitForReserves(interest, reserveFactorBps);

        assertEq(reserveShare + lendersShare, interest);
        assertLe(reserveShare, interest);
    }

    function testFuzz_reserve_share_never_exceeds_factor(uint256 interest, uint256 reserveFactorBps)
        public
        pure
    {
        interest = bound(interest, 0, 1e30);
        reserveFactorBps = bound(reserveFactorBps, 0, 5000);

        (uint256 reserveShare,) = InterestMath.splitForReserves(interest, reserveFactorBps);

        assertLe(reserveShare * 10_000, interest * reserveFactorBps + 10_000);
    }

    function testFuzz_borrow_index_never_shrinks(uint256 index, uint256 growth) public pure {
        index = bound(index, 1e18, 1e24);
        growth = bound(growth, 0, 1e20);

        assertGe(InterestMath.grownBorrowIndex(index, growth), index);
    }

    function testFuzz_supply_index_never_shrinks(uint256 index, uint256 lendersInterest, uint256 supplied)
        public
        pure
    {
        index = bound(index, 1e18, 1e24);
        lendersInterest = bound(lendersInterest, 0, 1e24);
        supplied = bound(supplied, 1, 1e30);

        assertGe(InterestMath.grownSupplyIndex(index, lendersInterest, supplied), index);
    }

    function testFuzz_longer_time_means_more_interest(
        uint256 rate,
        uint256 shortTime,
        uint256 longTime,
        uint256 debt
    ) public pure {
        rate = bound(rate, 1, 1e12);
        shortTime = bound(shortTime, 0, YEAR);
        longTime = bound(longTime, shortTime, 10 * YEAR);
        debt = bound(debt, 1e6, 1e24);

        uint256 shortInterest = InterestMath.interestOnDebt(debt, InterestMath.growthFactor(rate, shortTime));
        uint256 longInterest = InterestMath.interestOnDebt(debt, InterestMath.growthFactor(rate, longTime));

        assertGe(longInterest, shortInterest);
    }

    function testFuzz_zero_rate_accrues_nothing(uint256 elapsed, uint256 debt) public pure {
        elapsed = bound(elapsed, 0, 100 * YEAR);
        debt = bound(debt, 0, 1e30);

        assertEq(InterestMath.growthFactor(0, elapsed), 0);
        assertEq(InterestMath.interestOnDebt(debt, 0), 0);
    }
}
