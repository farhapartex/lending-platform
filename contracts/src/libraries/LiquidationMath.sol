// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {HealthMath} from "./HealthMath.sol";
import {WadMath} from "./WadMath.sol";

library LiquidationMath {
    struct SeizePlan {
        uint256 collateralToSeize;
        uint256 bonusValue;
        uint256 shortfall;
    }

    function bonusValue(uint256 debtValue, uint256 liquidationBonusBps) internal pure returns (uint256) {
        return WadMath.takeBpsDown(debtValue, liquidationBonusBps);
    }

    function rewardValue(uint256 debtValue, uint256 liquidationBonusBps) internal pure returns (uint256) {
        return debtValue + bonusValue(debtValue, liquidationBonusBps);
    }

    function planSeize(
        uint256 debtValue,
        uint256 collateralValue,
        uint256 collateralHeld,
        uint8 collateralDecimals,
        uint256 collateralPrice,
        uint256 liquidationBonusBps
    ) internal pure returns (SeizePlan memory plan) {
        plan.bonusValue = bonusValue(debtValue, liquidationBonusBps);

        uint256 targetValue = debtValue + plan.bonusValue;
        uint256 wantedAmount = HealthMath.amountFromValueUp(targetValue, collateralDecimals, collateralPrice);

        if (wantedAmount <= collateralHeld) {
            plan.collateralToSeize = wantedAmount;

            return plan;
        }

        plan.collateralToSeize = collateralHeld;
        plan.shortfall = WadMath.subtractOrZero(targetValue, collateralValue);

        return plan;
    }
}
