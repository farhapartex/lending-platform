// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";

library WadMath {
    uint256 internal constant WAD = 1e18;

    uint256 internal constant FULL_PERCENT_BPS = 10_000;

    uint256 internal constant PRICE_UNIT = 1e8;

    uint256 internal constant SECONDS_PER_YEAR = 365 days;

    function mulDown(uint256 value, uint256 multiplier, uint256 divisor) internal pure returns (uint256) {
        return Math.mulDiv(value, multiplier, divisor, Math.Rounding.Floor);
    }

    function mulUp(uint256 value, uint256 multiplier, uint256 divisor) internal pure returns (uint256) {
        return Math.mulDiv(value, multiplier, divisor, Math.Rounding.Ceil);
    }

    function wadMulDown(uint256 value, uint256 wadFactor) internal pure returns (uint256) {
        return Math.mulDiv(value, wadFactor, WAD, Math.Rounding.Floor);
    }

    function wadMulUp(uint256 value, uint256 wadFactor) internal pure returns (uint256) {
        return Math.mulDiv(value, wadFactor, WAD, Math.Rounding.Ceil);
    }

    function wadDivDown(uint256 value, uint256 wadDivisor) internal pure returns (uint256) {
        return Math.mulDiv(value, WAD, wadDivisor, Math.Rounding.Floor);
    }

    function wadDivUp(uint256 value, uint256 wadDivisor) internal pure returns (uint256) {
        return Math.mulDiv(value, WAD, wadDivisor, Math.Rounding.Ceil);
    }

    function takeBpsDown(uint256 value, uint256 bps) internal pure returns (uint256) {
        return Math.mulDiv(value, bps, FULL_PERCENT_BPS, Math.Rounding.Floor);
    }

    function takeBpsUp(uint256 value, uint256 bps) internal pure returns (uint256) {
        return Math.mulDiv(value, bps, FULL_PERCENT_BPS, Math.Rounding.Ceil);
    }

    function scaleUpByBpsDown(uint256 value, uint256 bps) internal pure returns (uint256) {
        return Math.mulDiv(value, FULL_PERCENT_BPS, bps, Math.Rounding.Floor);
    }

    function scaleUpByBpsUp(uint256 value, uint256 bps) internal pure returns (uint256) {
        return Math.mulDiv(value, FULL_PERCENT_BPS, bps, Math.Rounding.Ceil);
    }

    function smaller(uint256 left, uint256 right) internal pure returns (uint256) {
        return left < right ? left : right;
    }

    function larger(uint256 left, uint256 right) internal pure returns (uint256) {
        return left > right ? left : right;
    }

    function subtractOrZero(uint256 value, uint256 amountToRemove) internal pure returns (uint256) {
        return value > amountToRemove ? value - amountToRemove : 0;
    }
}
