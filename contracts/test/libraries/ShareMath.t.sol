// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";

import {ShareMath} from "../../src/libraries/ShareMath.sol";
import {WadMath} from "../../src/libraries/WadMath.sol";

contract ShareMathTest is Test {
    function test_starting_index_is_one() public pure {
        assertEq(ShareMath.STARTING_INDEX, 1e18);
    }

    function test_at_starting_index_shares_equal_assets() public pure {
        assertEq(ShareMath.sharesFromAssetsDown(1000e6, 1e18), 1000e6);
        assertEq(ShareMath.assetsFromSharesDown(1000e6, 1e18), 1000e6);
    }

    function test_grown_index_gives_fewer_shares() public pure {
        uint256 index = 1.05e18;

        assertEq(ShareMath.sharesFromAssetsDown(1000e6, index), 952380952);
        assertEq(ShareMath.assetsFromSharesDown(952380952, index), 999999999);
    }

    function test_late_joiner_does_not_capture_past_interest() public pure {
        uint256 index = 1.05e18;

        uint256 earlyShares = ShareMath.sharesFromAssetsDown(1000e6, 1e18);
        uint256 lateShares = ShareMath.sharesFromAssetsDown(1000e6, index);

        assertGt(earlyShares, lateShares);
        assertEq(ShareMath.assetsFromSharesDown(earlyShares, index), 1050e6);
    }

    function test_rounding_directions_are_opposite() public pure {
        uint256 index = 1002643162300000000;

        assertEq(ShareMath.sharesFromAssetsDown(1000e6, index), 997363805);
        assertEq(ShareMath.sharesFromAssetsUp(1000e6, index), 997363806);
    }

    function test_zero_inputs() public pure {
        assertEq(ShareMath.sharesFromAssetsDown(0, 1e18), 0);
        assertEq(ShareMath.assetsFromSharesUp(0, 1e18), 0);
    }

    function testFuzz_up_never_below_down(uint256 assets, uint256 index) public pure {
        assets = bound(assets, 0, 1e30);
        index = bound(index, 1e18, 1e24);

        assertGe(ShareMath.sharesFromAssetsUp(assets, index), ShareMath.sharesFromAssetsDown(assets, index));
        assertGe(ShareMath.assetsFromSharesUp(assets, index), ShareMath.assetsFromSharesDown(assets, index));
    }

    function testFuzz_deposit_then_withdraw_never_gains(uint256 assets, uint256 index) public pure {
        assets = bound(assets, 0, 1e30);
        index = bound(index, 1e18, 1e24);

        uint256 shares = ShareMath.sharesFromAssetsDown(assets, index);
        uint256 back = ShareMath.assetsFromSharesDown(shares, index);

        assertLe(back, assets);
    }

    function testFuzz_withdraw_burns_at_least_enough(uint256 assets, uint256 index) public pure {
        assets = bound(assets, 0, 1e30);
        index = bound(index, 1e18, 1e24);

        uint256 shares = ShareMath.sharesFromAssetsUp(assets, index);
        uint256 covered = ShareMath.assetsFromSharesDown(shares, index);

        assertGe(covered, assets);
    }

    function testFuzz_debt_never_understated(uint256 shares, uint256 index) public pure {
        shares = bound(shares, 0, 1e30);
        index = bound(index, 1e18, 1e24);

        uint256 owed = ShareMath.assetsFromSharesUp(shares, index);
        uint256 sharesForOwed = ShareMath.sharesFromAssetsDown(owed, index);

        assertLe(sharesForOwed, shares);
    }

    function testFuzz_growing_index_never_reduces_balance(uint256 shares, uint256 index, uint256 growth)
        public
        pure
    {
        shares = bound(shares, 0, 1e30);
        index = bound(index, 1e18, 1e23);
        growth = bound(growth, 0, 1e23);

        uint256 before = ShareMath.assetsFromSharesDown(shares, index);
        uint256 after_ = ShareMath.assetsFromSharesDown(shares, index + growth);

        assertGe(after_, before);
    }

    function testFuzz_proportional_split_is_fair(uint256 depositA, uint256 depositB, uint256 index)
        public
        pure
    {
        depositA = bound(depositA, 1e6, 1e24);
        depositB = bound(depositB, 1e6, 1e24);
        index = bound(index, 1e18, 1e20);

        uint256 sharesA = ShareMath.sharesFromAssetsDown(depositA, index);
        uint256 sharesB = ShareMath.sharesFromAssetsDown(depositB, index);

        if (depositA >= depositB) {
            assertGe(sharesA, sharesB);
        } else {
            assertLe(sharesA, sharesB);
        }
    }

    function testFuzz_index_growth_scales_all_holders_equally(
        uint256 sharesA,
        uint256 sharesB,
        uint256 indexBefore,
        uint256 indexAfter
    ) public pure {
        sharesA = bound(sharesA, 1e12, 1e24);
        sharesB = bound(sharesB, 1e12, 1e24);
        indexBefore = bound(indexBefore, 1e18, 1e20);
        indexAfter = bound(indexAfter, indexBefore, 1e21);

        uint256 gainA = ShareMath.assetsFromSharesDown(sharesA, indexAfter)
            - ShareMath.assetsFromSharesDown(sharesA, indexBefore);
        uint256 gainB = ShareMath.assetsFromSharesDown(sharesB, indexAfter)
            - ShareMath.assetsFromSharesDown(sharesB, indexBefore);

        if (sharesA > sharesB) {
            assertGe(gainA, gainB);
        }
    }

    function testFuzz_donation_cannot_move_share_value(uint256 shares, uint256 index) public pure {
        shares = bound(shares, 1, 1e30);
        index = bound(index, 1e18, 1e24);

        uint256 valueBefore = ShareMath.assetsFromSharesDown(shares, index);
        uint256 valueAfter = ShareMath.assetsFromSharesDown(shares, index);

        assertEq(valueBefore, valueAfter);
        assertEq(valueBefore, WadMath.mulDown(shares, index, 1e18));
    }
}
