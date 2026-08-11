// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

import {IInterestRateModel, RateCurve} from "./interfaces/IInterestRateModel.sol";
import {Errors} from "./libraries/Errors.sol";
import {WadMath} from "./libraries/WadMath.sol";

contract InterestRateModel is IInterestRateModel, Ownable {
    event CurveChanged(
        uint64 baseRatePerSecond,
        uint64 slopeBelowKinkPerSecond,
        uint64 slopeAboveKinkPerSecond,
        uint16 kinkUtilizationBps
    );

    RateCurve private rateCurve;

    constructor(address owner, RateCurve memory startingCurve) Ownable(owner) {
        _storeCurve(startingCurve);
    }

    function setCurve(RateCurve memory newCurve) external onlyOwner {
        _storeCurve(newCurve);
    }

    function curve() external view returns (RateCurve memory) {
        return rateCurve;
    }

    function utilizationBps(uint256 totalSupplied, uint256 totalBorrowed) public pure returns (uint256) {
        if (totalSupplied == 0) {
            return 0;
        }

        uint256 usage = WadMath.mulDown(totalBorrowed, WadMath.FULL_PERCENT_BPS, totalSupplied);

        return WadMath.smaller(usage, WadMath.FULL_PERCENT_BPS);
    }

    function borrowRatePerSecond(uint256 usageBps) public view returns (uint256) {
        return _borrowRate(rateCurve, usageBps);
    }

    function supplyRatePerSecond(uint256 usageBps, uint256 reserveFactorBps) public view returns (uint256) {
        return _supplyRate(rateCurve, usageBps, reserveFactorBps);
    }

    function borrowAprBps(uint256 usageBps) external view returns (uint256) {
        return _toAprBps(_borrowRate(rateCurve, usageBps));
    }

    function supplyAprBps(uint256 usageBps, uint256 reserveFactorBps) external view returns (uint256) {
        return _toAprBps(_supplyRate(rateCurve, usageBps, reserveFactorBps));
    }

    function _borrowRate(RateCurve memory loadedCurve, uint256 usageBps) private pure returns (uint256) {
        if (usageBps <= loadedCurve.kinkUtilizationBps) {
            uint256 belowKink = WadMath.mulDown(
                loadedCurve.slopeBelowKinkPerSecond, usageBps, loadedCurve.kinkUtilizationBps
            );

            return loadedCurve.baseRatePerSecond + belowKink;
        }

        uint256 rangeAboveKink = WadMath.FULL_PERCENT_BPS - loadedCurve.kinkUtilizationBps;
        uint256 usageAboveKink = usageBps - loadedCurve.kinkUtilizationBps;

        uint256 aboveKink =
            WadMath.mulDown(loadedCurve.slopeAboveKinkPerSecond, usageAboveKink, rangeAboveKink);

        return loadedCurve.baseRatePerSecond + loadedCurve.slopeBelowKinkPerSecond + aboveKink;
    }

    function _supplyRate(RateCurve memory loadedCurve, uint256 usageBps, uint256 reserveFactorBps)
        private
        pure
        returns (uint256)
    {
        if (usageBps == 0) {
            return 0;
        }

        uint256 paidByBorrowers = WadMath.takeBpsDown(_borrowRate(loadedCurve, usageBps), usageBps);
        uint256 shareForLenders = WadMath.FULL_PERCENT_BPS - reserveFactorBps;

        return WadMath.takeBpsDown(paidByBorrowers, shareForLenders);
    }

    function _toAprBps(uint256 ratePerSecond) private pure returns (uint256) {
        return
            WadMath.mulDown(ratePerSecond * WadMath.SECONDS_PER_YEAR, WadMath.FULL_PERCENT_BPS, WadMath.WAD);
    }

    function _storeCurve(RateCurve memory newCurve) private {
        if (newCurve.kinkUtilizationBps == 0 || newCurve.kinkUtilizationBps >= WadMath.FULL_PERCENT_BPS) {
            revert Errors.InvalidRiskSettings();
        }

        rateCurve = newCurve;

        emit CurveChanged(
            newCurve.baseRatePerSecond,
            newCurve.slopeBelowKinkPerSecond,
            newCurve.slopeAboveKinkPerSecond,
            newCurve.kinkUtilizationBps
        );
    }
}
