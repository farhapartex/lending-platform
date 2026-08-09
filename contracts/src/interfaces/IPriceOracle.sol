// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

struct PriceData {
    uint256 price;
    uint8 decimals;
    uint256 updatedAt;
    bool isStale;
    bool isValid;
}

interface IPriceOracle {
    function readPrice(address asset) external view returns (PriceData memory);

    function getPrice(address asset) external view returns (uint256 price, uint8 decimals);

    function isStale(address asset) external view returns (bool);

    function maxPriceAge() external view returns (uint32);

    function feedOf(address asset) external view returns (address);
}
