// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface ILiquidationManager {
    function isLiquidatable(address borrower) external view returns (bool);

    function previewLiquidation(address borrower)
        external
        view
        returns (uint256 debtToRepay, uint256 collateralToSeize, uint256 bonusValue, uint256 shortfall);

    function liquidate(address borrower) external;
}
