// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {WadMath} from "./WadMath.sol";

library HealthMath {
    uint256 internal constant LIQUIDATION_LEVEL_BPS = 10_000;

    uint256 internal constant NO_DEBT_HEALTH_FACTOR = type(uint256).max;

    function valueOfAmount(uint256 amount, uint8 tokenDecimals, uint256 unitPrice)
        internal
        pure
        returns (uint256)
    {
        return WadMath.mulDown(amount, unitPrice, 10 ** tokenDecimals);
    }

    function amountFromValueDown(uint256 value, uint8 tokenDecimals, uint256 unitPrice)
        internal
        pure
        returns (uint256)
    {
        if (unitPrice == 0) {
            return 0;
        }

        return WadMath.mulDown(value, 10 ** tokenDecimals, unitPrice);
    }

    function amountFromValueUp(uint256 value, uint8 tokenDecimals, uint256 unitPrice)
        internal
        pure
        returns (uint256)
    {
        if (unitPrice == 0) {
            return 0;
        }

        return WadMath.mulUp(value, 10 ** tokenDecimals, unitPrice);
    }

    function healthFactor(uint256 collateralValue, uint256 debtValue, uint256 liquidationThresholdBps)
        internal
        pure
        returns (uint256)
    {
        if (debtValue == 0) {
            return NO_DEBT_HEALTH_FACTOR;
        }

        return WadMath.mulDown(collateralValue, liquidationThresholdBps, debtValue);
    }

    function isLiquidatable(uint256 healthFactorBps) internal pure returns (bool) {
        return healthFactorBps < LIQUIDATION_LEVEL_BPS;
    }

    function borrowLimitValue(uint256 collateralValue, uint256 maxLtvBps) internal pure returns (uint256) {
        return WadMath.takeBpsDown(collateralValue, maxLtvBps);
    }

    function liquidationLimitValue(uint256 collateralValue, uint256 liquidationThresholdBps)
        internal
        pure
        returns (uint256)
    {
        return WadMath.takeBpsDown(collateralValue, liquidationThresholdBps);
    }

    function remainingBorrowAmount(
        uint256 collateralValue,
        uint256 debtValue,
        uint256 maxLtvBps,
        uint8 debtDecimals,
        uint256 debtUnitPrice
    ) internal pure returns (uint256) {
        uint256 limitValue = borrowLimitValue(collateralValue, maxLtvBps);

        if (limitValue <= debtValue) {
            return 0;
        }

        return amountFromValueDown(limitValue - debtValue, debtDecimals, debtUnitPrice);
    }

    function collateralValueNeeded(uint256 debtValue, uint256 maxLtvBps) internal pure returns (uint256) {
        if (debtValue == 0) {
            return 0;
        }

        return WadMath.scaleUpByBpsUp(debtValue, maxLtvBps);
    }

    function freeCollateralAmount(
        uint256 collateralAmount,
        uint8 collateralDecimals,
        uint256 collateralUnitPrice,
        uint256 debtValue,
        uint256 maxLtvBps
    ) internal pure returns (uint256) {
        if (debtValue == 0) {
            return collateralAmount;
        }

        uint256 valueNeeded = collateralValueNeeded(debtValue, maxLtvBps);
        uint256 amountNeeded = amountFromValueUp(valueNeeded, collateralDecimals, collateralUnitPrice);

        return WadMath.subtractOrZero(collateralAmount, amountNeeded);
    }
}
