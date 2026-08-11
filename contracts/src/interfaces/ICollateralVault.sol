// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface ICollateralVault {
    function collateralAsset() external view returns (address);

    function depositCollateral(uint256 amount) external;

    function withdrawCollateral(uint256 amount) external;

    function collateralOf(address borrower) external view returns (uint256);

    function totalCollateral() external view returns (uint256);

    function seize(address borrower, address recipient, uint256 amount) external;
}
