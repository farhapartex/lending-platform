// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

import {PriceOracleAdapter} from "../src/PriceOracleAdapter.sol";
import {MockAggregator} from "../src/mocks/MockAggregator.sol";
import {MockERC20} from "../src/mocks/MockERC20.sol";
import {PriceData} from "../src/interfaces/IPriceOracle.sol";
import {Errors} from "../src/libraries/Errors.sol";

contract PriceOracleAdapterTest is Test {
    uint32 constant MAX_AGE = 3600;
    int256 constant ETH_PRICE = 341258000000;

    PriceOracleAdapter oracle;
    MockAggregator feed;
    MockERC20 token;

    address admin = makeAddr("admin");
    address stranger = makeAddr("stranger");

    function setUp() public {
        vm.warp(1_700_000_000);

        token = new MockERC20("Wrapped Ether", "WETH", 18);
        feed = new MockAggregator("ETH / USD", 8, ETH_PRICE);
        oracle = new PriceOracleAdapter(admin, MAX_AGE);

        vm.prank(admin);
        oracle.setFeed(address(token), address(feed));
    }

    function test_constructor_rejects_zero_max_age() public {
        vm.expectRevert(Errors.InvalidRiskSettings.selector);
        new PriceOracleAdapter(admin, 0);
    }

    function test_getPrice_happy_path() public view {
        (uint256 price, uint8 decimals) = oracle.getPrice(address(token));

        assertEq(price, uint256(ETH_PRICE));
        assertEq(decimals, 8);
    }

    function test_setFeed_requires_eight_decimals() public {
        MockAggregator eighteenDecimals = new MockAggregator("ETH / ETH", 18, 1e18);
        MockERC20 other = new MockERC20("Other", "OTH", 18);

        vm.prank(admin);
        vm.expectRevert(
            abi.encodeWithSelector(
                Errors.UnsupportedFeedDecimals.selector, address(eighteenDecimals), uint8(18), uint8(8)
            )
        );
        oracle.setFeed(address(other), address(eighteenDecimals));
    }

    function test_setFeed_rejects_zero_addresses() public {
        vm.startPrank(admin);

        vm.expectRevert(Errors.ZeroAddress.selector);
        oracle.setFeed(address(0), address(feed));

        vm.expectRevert(Errors.ZeroAddress.selector);
        oracle.setFeed(address(token), address(0));

        vm.stopPrank();
    }

    function test_setFeed_is_owner_only() public {
        vm.prank(stranger);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, stranger));
        oracle.setFeed(address(token), address(feed));
    }

    function test_setMaxPriceAge_rules() public {
        vm.prank(stranger);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, stranger));
        oracle.setMaxPriceAge(7200);

        vm.startPrank(admin);

        vm.expectRevert(Errors.InvalidRiskSettings.selector);
        oracle.setMaxPriceAge(0);

        oracle.setMaxPriceAge(7200);
        assertEq(oracle.maxPriceAge(), 7200);

        vm.stopPrank();
    }

    function test_unregistered_asset_reverts_on_getPrice() public {
        MockERC20 other = new MockERC20("Other", "OTH", 18);

        vm.expectRevert(abi.encodeWithSelector(Errors.PriceIsInvalid.selector, address(other)));
        oracle.getPrice(address(other));
    }

    function test_unregistered_asset_reads_as_stale_without_reverting() public {
        MockERC20 other = new MockERC20("Other", "OTH", 18);

        PriceData memory data = oracle.readPrice(address(other));

        assertTrue(data.isStale);
        assertFalse(data.isValid);
        assertEq(data.price, 0);
        assertTrue(oracle.isStale(address(other)));
    }

    function test_staleness_boundary() public {
        skip(MAX_AGE);
        assertFalse(oracle.isStale(address(token)));
        (uint256 price,) = oracle.getPrice(address(token));
        assertEq(price, uint256(ETH_PRICE));

        skip(1);
        assertTrue(oracle.isStale(address(token)));
    }

    function test_stale_price_reverts_but_read_survives() public {
        skip(MAX_AGE + 1);

        vm.expectRevert(
            abi.encodeWithSelector(
                Errors.PriceIsStale.selector, address(token), block.timestamp - MAX_AGE - 1, MAX_AGE
            )
        );
        oracle.getPrice(address(token));

        PriceData memory data = oracle.readPrice(address(token));
        assertTrue(data.isStale);
        assertFalse(data.isValid);
    }

    function test_zero_answer_is_invalid() public {
        feed.setPrice(0);

        vm.expectRevert(abi.encodeWithSelector(Errors.PriceIsInvalid.selector, address(token)));
        oracle.getPrice(address(token));

        assertTrue(oracle.isStale(address(token)));
    }

    function test_negative_answer_is_invalid() public {
        feed.setPrice(-1);

        vm.expectRevert(abi.encodeWithSelector(Errors.PriceIsInvalid.selector, address(token)));
        oracle.getPrice(address(token));

        PriceData memory data = oracle.readPrice(address(token));
        assertEq(data.price, 0);
        assertFalse(data.isValid);
    }

    function test_future_timestamp_is_rejected() public {
        feed.setPriceWithTimestamp(ETH_PRICE, uint40(block.timestamp + 1000));

        vm.expectRevert(abi.encodeWithSelector(Errors.PriceIsInvalid.selector, address(token)));
        oracle.getPrice(address(token));
    }

    function test_incomplete_round_is_rejected() public {
        feed.setIncompleteRound(ETH_PRICE);

        vm.expectRevert(abi.encodeWithSelector(Errors.PriceIsInvalid.selector, address(token)));
        oracle.getPrice(address(token));

        assertTrue(oracle.isStale(address(token)));
    }

    function test_reverting_aggregator_does_not_break_reads() public {
        feed.setFeedDown(true);

        vm.expectRevert(abi.encodeWithSelector(Errors.PriceIsInvalid.selector, address(token)));
        oracle.getPrice(address(token));

        PriceData memory data = oracle.readPrice(address(token));
        assertTrue(data.isStale);
        assertFalse(data.isValid);
        assertEq(data.price, 0);

        assertTrue(oracle.isStale(address(token)));
    }

    function test_feed_recovers_after_outage() public {
        feed.setFeedDown(true);
        assertTrue(oracle.isStale(address(token)));

        feed.setFeedDown(false);
        feed.setPrice(ETH_PRICE);

        (uint256 price,) = oracle.getPrice(address(token));
        assertEq(price, uint256(ETH_PRICE));
    }

    function test_feedOf_reports_registered_aggregator() public {
        assertEq(oracle.feedOf(address(token)), address(feed));
        assertEq(oracle.feedOf(makeAddr("unknown")), address(0));
    }

    function test_replacing_a_feed_takes_effect() public {
        MockAggregator replacement = new MockAggregator("ETH / USD v2", 8, 200000000000);

        vm.prank(admin);
        oracle.setFeed(address(token), address(replacement));

        (uint256 price,) = oracle.getPrice(address(token));
        assertEq(price, 200000000000);
    }

    function testFuzz_positive_fresh_price_is_always_readable(int256 answer) public {
        answer = bound(answer, 1, type(int64).max);

        feed.setPrice(answer);

        (uint256 price,) = oracle.getPrice(address(token));
        assertEq(price, uint256(answer));
        assertFalse(oracle.isStale(address(token)));
    }

    function testFuzz_non_positive_price_never_readable(int256 answer) public {
        answer = bound(answer, type(int64).min, 0);

        feed.setPrice(answer);

        vm.expectRevert(abi.encodeWithSelector(Errors.PriceIsInvalid.selector, address(token)));
        oracle.getPrice(address(token));
    }

    function testFuzz_staleness_follows_max_age(uint32 age, uint32 elapsed) public {
        age = uint32(bound(age, 1, 30 days));
        elapsed = uint32(bound(elapsed, 0, 60 days));

        vm.prank(admin);
        oracle.setMaxPriceAge(age);

        feed.setPrice(ETH_PRICE);
        skip(elapsed);

        assertEq(oracle.isStale(address(token)), elapsed > age);
    }

    function testFuzz_readPrice_never_reverts(int256 answer, uint32 elapsed, bool down) public {
        answer = bound(answer, type(int64).min, type(int64).max);
        elapsed = uint32(bound(elapsed, 0, 60 days));

        feed.setPrice(answer);
        feed.setFeedDown(down);
        skip(elapsed);

        PriceData memory data = oracle.readPrice(address(token));

        if (data.isValid) {
            assertGt(data.price, 0);
        }
    }
}
