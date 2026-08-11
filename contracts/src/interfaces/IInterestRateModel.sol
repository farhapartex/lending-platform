// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

struct RateCurve {
    uint64 baseRatePerSecond;
    uint64 slopeBelowKinkPerSecond;
    uint64 slopeAboveKinkPerSecond;
    uint16 kinkUtilizationBps;
}

interface IInterestRateModel {
    function borrowRatePerSecond(uint256 utilizationBps) external view returns (uint256);

    function supplyRatePerSecond(uint256 utilizationBps, uint256 reserveFactorBps)
        external
        view
        returns (uint256);

    function curve() external view returns (RateCurve memory);

    function utilizationBps(uint256 totalSupplied, uint256 totalBorrowed) external pure returns (uint256);
}
