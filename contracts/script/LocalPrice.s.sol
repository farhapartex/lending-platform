// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Script, console2} from "forge-std/Script.sol";

import {MockAggregator} from "../src/mocks/MockAggregator.sol";
import {IPriceOracle, PriceData} from "../src/interfaces/IPriceOracle.sol";

contract LocalPrice is Script {
    uint256 private constant ANVIL_CHAIN_ID = 31337;

    error WrongChain(uint256 chainId);

    function run() external {
        if (block.chainid != ANVIL_CHAIN_ID) {
            revert WrongChain(block.chainid);
        }

        MockAggregator collateralFeed = MockAggregator(vm.envAddress("COLLATERAL_FEED"));
        MockAggregator debtFeed = MockAggregator(vm.envAddress("DEBT_FEED"));

        (, int256 currentCollateralAnswer,,,) = collateralFeed.latestRoundData();
        (, int256 currentDebtAnswer,,,) = debtFeed.latestRoundData();

        int256 collateralAnswer = int256(vm.envOr("COLLATERAL_PRICE", uint256(0)));
        int256 debtAnswer = int256(vm.envOr("DEBT_PRICE", uint256(0)));

        if (collateralAnswer == 0) {
            collateralAnswer = currentCollateralAnswer;
        }

        if (debtAnswer == 0) {
            debtAnswer = currentDebtAnswer;
        }

        uint256 deployerKey = vm.envOr("DEPLOYER_PRIVATE_KEY", uint256(0));

        if (deployerKey == 0) {
            vm.startBroadcast();
        } else {
            vm.startBroadcast(deployerKey);
        }

        collateralFeed.setPrice(collateralAnswer);
        debtFeed.setPrice(debtAnswer);

        vm.stopBroadcast();

        console2.log("collateral price     ", collateralAnswer);
        console2.log("debt price           ", debtAnswer);

        address oracleAddress = vm.envOr("ORACLE_ADDRESS", address(0));

        if (oracleAddress != address(0)) {
            _reportStaleness(IPriceOracle(oracleAddress));
        }
    }

    function _reportStaleness(IPriceOracle oracle) private view {
        PriceData memory collateral = oracle.readPrice(vm.envAddress("COLLATERAL_TOKEN"));
        PriceData memory debt = oracle.readPrice(vm.envAddress("DEBT_TOKEN"));

        console2.log("collateral isValid   ", collateral.isValid);
        console2.log("collateral updatedAt ", collateral.updatedAt);
        console2.log("debt isValid         ", debt.isValid);
        console2.log("debt updatedAt       ", debt.updatedAt);
        console2.log("max price age        ", oracle.maxPriceAge());
    }
}
