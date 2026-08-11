// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {WadMath} from "./WadMath.sol";

library ShareMath {
    uint256 internal constant STARTING_INDEX = WadMath.WAD;

    function sharesFromAssetsDown(uint256 assets, uint256 index) internal pure returns (uint256) {
        return WadMath.mulDown(assets, WadMath.WAD, index);
    }

    function sharesFromAssetsUp(uint256 assets, uint256 index) internal pure returns (uint256) {
        return WadMath.mulUp(assets, WadMath.WAD, index);
    }

    function assetsFromSharesDown(uint256 shares, uint256 index) internal pure returns (uint256) {
        return WadMath.mulDown(shares, index, WadMath.WAD);
    }

    function assetsFromSharesUp(uint256 shares, uint256 index) internal pure returns (uint256) {
        return WadMath.mulUp(shares, index, WadMath.WAD);
    }
}
