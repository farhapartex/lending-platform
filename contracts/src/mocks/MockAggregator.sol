// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {IAggregatorV3} from "../interfaces/IAggregatorV3.sol";

contract MockAggregator is IAggregatorV3 {
    struct Round {
        int256 answer;
        uint80 roundId;
        uint80 answeredInRound;
        uint40 startedAt;
        uint40 updatedAt;
    }

    error FeedIsDown();
    error RoundNotFound(uint80 roundId);

    event AnswerRecorded(uint80 indexed roundId, int256 answer, uint40 updatedAt);

    uint256 private constant FEED_VERSION = 4;

    uint8 private immutable feedDecimals;

    string private feedDescription;

    uint80 public latestRoundId;

    bool public feedDown;

    mapping(uint80 => Round) private rounds;

    constructor(string memory description_, uint8 decimals_, int256 startingAnswer) {
        feedDescription = description_;
        feedDecimals = decimals_;

        _recordRound(startingAnswer, uint40(block.timestamp), 1);
    }

    function decimals() external view returns (uint8) {
        return feedDecimals;
    }

    function description() external view returns (string memory) {
        return feedDescription;
    }

    function version() external pure returns (uint256) {
        return FEED_VERSION;
    }

    function setPrice(int256 newAnswer) external {
        _recordRound(newAnswer, uint40(block.timestamp), latestRoundId + 1);
    }

    function setPriceWithTimestamp(int256 newAnswer, uint40 updatedAt) external {
        _recordRound(newAnswer, updatedAt, latestRoundId + 1);
    }

    function setIncompleteRound(int256 newAnswer) external {
        uint80 nextRoundId = latestRoundId + 1;

        _recordRoundWithAnsweredIn(newAnswer, uint40(block.timestamp), nextRoundId, latestRoundId);
    }

    function makeStale(uint40 secondsInThePast) external {
        Round memory current = rounds[latestRoundId];

        rounds[latestRoundId].updatedAt = uint40(block.timestamp) - secondsInThePast;

        emit AnswerRecorded(current.roundId, current.answer, uint40(block.timestamp) - secondsInThePast);
    }

    function setFeedDown(bool isDown) external {
        feedDown = isDown;
    }

    function latestRoundData()
        external
        view
        returns (uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
    {
        return _readRound(latestRoundId);
    }

    function getRoundData(uint80 wantedRoundId)
        external
        view
        returns (uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
    {
        return _readRound(wantedRoundId);
    }

    function _readRound(uint80 wantedRoundId)
        private
        view
        returns (uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
    {
        if (feedDown) {
            revert FeedIsDown();
        }

        Round memory round = rounds[wantedRoundId];

        if (round.roundId == 0) {
            revert RoundNotFound(wantedRoundId);
        }

        return (round.roundId, round.answer, round.startedAt, round.updatedAt, round.answeredInRound);
    }

    function _recordRound(int256 answer, uint40 updatedAt, uint80 roundId) private {
        _recordRoundWithAnsweredIn(answer, updatedAt, roundId, roundId);
    }

    function _recordRoundWithAnsweredIn(
        int256 answer,
        uint40 updatedAt,
        uint80 roundId,
        uint80 answeredInRound
    ) private {
        rounds[roundId] = Round({
            answer: answer,
            roundId: roundId,
            answeredInRound: answeredInRound,
            startedAt: updatedAt,
            updatedAt: updatedAt
        });

        latestRoundId = roundId;

        emit AnswerRecorded(roundId, answer, updatedAt);
    }
}
