// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockERC20 is ERC20 {
    uint8 private immutable tokenDecimals;

    constructor(string memory tokenName, string memory tokenSymbol, uint8 tokenDecimalPlaces)
        ERC20(tokenName, tokenSymbol)
    {
        tokenDecimals = tokenDecimalPlaces;
    }

    function decimals() public view override returns (uint8) {
        return tokenDecimals;
    }

    function mint(address receiver, uint256 amount) external {
        _mint(receiver, amount);
    }

    function burn(address holder, uint256 amount) external {
        _burn(holder, amount);
    }
}
