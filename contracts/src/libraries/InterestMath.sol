// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {WadMath} from "./WadMath.sol";

library InterestMath {
    function growthFactor(uint256 ratePerSecond, uint256 elapsedSeconds) internal pure returns (uint256) {
        return ratePerSecond * elapsedSeconds;
    }

    function interestOnDebt(uint256 borrowedAmount, uint256 growth) internal pure returns (uint256) {
        return WadMath.wadMulUp(borrowedAmount, growth);
    }

    function grownBorrowIndex(uint256 currentIndex, uint256 growth) internal pure returns (uint256) {
        return currentIndex + WadMath.wadMulUp(currentIndex, growth);
    }

    function grownSupplyIndex(uint256 currentIndex, uint256 lendersInterest, uint256 suppliedBefore)
        internal
        pure
        returns (uint256)
    {
        if (suppliedBefore == 0 || lendersInterest == 0) {
            return currentIndex;
        }

        return currentIndex + WadMath.mulDown(currentIndex, lendersInterest, suppliedBefore);
    }

    function splitForReserves(uint256 interest, uint256 reserveFactorBps)
        internal
        pure
        returns (uint256 reserveShare, uint256 lendersShare)
    {
        reserveShare = WadMath.takeBpsDown(interest, reserveFactorBps);

        return (reserveShare, interest - reserveShare);
    }
}
