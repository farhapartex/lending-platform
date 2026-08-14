// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

import {IAggregatorV3} from "./interfaces/IAggregatorV3.sol";
import {IPriceOracle, PriceData} from "./interfaces/IPriceOracle.sol";
import {Errors} from "./libraries/Errors.sol";

contract PriceOracleAdapter is IPriceOracle, Ownable {
    struct Feed {
        address aggregator;
        uint8 decimals;
    }

    uint8 public constant REQUIRED_FEED_DECIMALS = 8;

    event FeedChanged(address indexed asset, address aggregator, uint8 decimals);
    event MaxPriceAgeChanged(uint32 previousAge, uint32 newAge);

    uint32 public maxPriceAge;

    mapping(address => Feed) private feeds;

    constructor(address owner, uint32 startingMaxPriceAge) Ownable(owner) {
        if (startingMaxPriceAge == 0) {
            revert Errors.InvalidRiskSettings();
        }

        maxPriceAge = startingMaxPriceAge;

        emit MaxPriceAgeChanged(0, startingMaxPriceAge);
    }

    function setFeed(address asset, address aggregator) external onlyOwner {
        if (asset == address(0) || aggregator == address(0)) {
            revert Errors.ZeroAddress();
        }

        uint8 feedDecimals = IAggregatorV3(aggregator).decimals();

        if (feedDecimals != REQUIRED_FEED_DECIMALS) {
            revert Errors.UnsupportedFeedDecimals(aggregator, feedDecimals, REQUIRED_FEED_DECIMALS);
        }

        feeds[asset] = Feed({aggregator: aggregator, decimals: feedDecimals});

        emit FeedChanged(asset, aggregator, feedDecimals);
    }

    function setMaxPriceAge(uint32 newAge) external onlyOwner {
        if (newAge == 0) {
            revert Errors.InvalidRiskSettings();
        }

        uint32 previousAge = maxPriceAge;
        maxPriceAge = newAge;

        emit MaxPriceAgeChanged(previousAge, newAge);
    }

    function feedOf(address asset) external view returns (address) {
        return feeds[asset].aggregator;
    }

    function readPrice(address asset) external view returns (PriceData memory) {
        return _evaluateFeed(feeds[asset]);
    }

    function getPrice(address asset) external view returns (uint256 price, uint8 decimals) {
        Feed memory feed = feeds[asset];

        if (feed.aggregator == address(0)) {
            revert Errors.PriceIsInvalid(asset);
        }

        PriceData memory data = _evaluateFeed(feed);

        if (data.price == 0) {
            revert Errors.PriceIsInvalid(asset);
        }

        if (data.isStale) {
            revert Errors.PriceIsStale(asset, data.updatedAt, maxPriceAge);
        }

        return (data.price, data.decimals);
    }

    function isStale(address asset) external view returns (bool) {
        return !_evaluateFeed(feeds[asset]).isValid;
    }

    function _evaluateFeed(Feed memory feed) private view returns (PriceData memory data) {
        if (feed.aggregator == address(0)) {
            data.isStale = true;
            return data;
        }

        data.decimals = feed.decimals;

        (bool reachable, int256 answer, uint256 updatedAt, uint80 roundId, uint80 answeredInRound) =
            _fetchLatestRound(feed.aggregator);

        if (!reachable) {
            data.isStale = true;
            return data;
        }

        data.updatedAt = updatedAt;

        if (!_isAnswerUsable(answer, updatedAt, roundId, answeredInRound)) {
            data.isStale = true;
            return data;
        }

        data.price = SafeCast.toUint256(answer);
        data.isStale = _isOlderThanMaxAge(updatedAt);
        data.isValid = !data.isStale;

        return data;
    }

    function _fetchLatestRound(address aggregator)
        private
        view
        returns (bool reachable, int256 answer, uint256 updatedAt, uint80 roundId, uint80 answeredInRound)
    {
        try IAggregatorV3(aggregator).latestRoundData() returns (
            uint80 fetchedRoundId,
            int256 fetchedAnswer,
            uint256,
            uint256 fetchedUpdatedAt,
            uint80 fetchedAnsweredInRound
        ) {
            return (true, fetchedAnswer, fetchedUpdatedAt, fetchedRoundId, fetchedAnsweredInRound);
        } catch {
            return (false, 0, 0, 0, 0);
        }
    }

    function _isAnswerUsable(int256 answer, uint256 updatedAt, uint80 roundId, uint80 answeredInRound)
        private
        view
        returns (bool)
    {
        if (answer <= 0) {
            return false;
        }

        if (updatedAt == 0 || updatedAt > block.timestamp) {
            return false;
        }

        return answeredInRound >= roundId;
    }

    function _isOlderThanMaxAge(uint256 updatedAt) private view returns (bool) {
        return block.timestamp - updatedAt > maxPriceAge;
    }
}
