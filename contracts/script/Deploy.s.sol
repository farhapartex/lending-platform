// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Script, console2} from "forge-std/Script.sol";

import {DeployConfig, Deployment, ProtocolDeployer} from "./Deployment.sol";

contract Deploy is Script {
    function run() external returns (Deployment memory deployed) {
        uint256 deployerKey = vm.envOr("DEPLOYER_PRIVATE_KEY", uint256(0));
        address deployer = deployerKey == 0 ? msg.sender : vm.addr(deployerKey);

        DeployConfig memory config = _readConfig();

        console2.log("=== configuration ===");
        console2.log("chain id            ", block.chainid);
        console2.log("deployer            ", deployer);
        console2.log("final owner         ", config.finalOwner);
        console2.log("collateral token    ", config.collateralToken);
        console2.log("debt token          ", config.debtToken);
        console2.log("collateral feed     ", config.collateralFeed);
        console2.log("debt feed           ", config.debtFeed);
        console2.log("max price age       ", config.maxPriceAge);
        console2.log("min deposit         ", config.minDeposit);
        console2.log("reserve factor bps  ", config.reserveFactorBps);
        console2.log("max ltv bps         ", config.maxLtvBps);
        console2.log("liq threshold bps   ", config.liquidationThresholdBps);
        console2.log("liq bonus bps       ", config.liquidationBonusBps);

        ProtocolDeployer.requireEightDecimalFeeds(config);

        if (deployerKey == 0) {
            vm.startBroadcast();
        } else {
            vm.startBroadcast(deployerKey);
        }

        deployed = ProtocolDeployer.deployAll(config, deployer);
        ProtocolDeployer.wireAll(deployed, config);

        if (config.finalOwner != deployer) {
            ProtocolDeployer.handOverOwnership(deployed, config.finalOwner);
        }

        vm.stopBroadcast();

        ProtocolDeployer.verifyWiring(deployed, config);
        ProtocolDeployer.verifyOwnership(deployed, config.finalOwner);

        _report(deployed);

        return deployed;
    }

    function _readConfig() private view returns (DeployConfig memory config) {
        uint256 deployerKey = vm.envOr("DEPLOYER_PRIVATE_KEY", uint256(0));
        address deployer = deployerKey == 0 ? msg.sender : vm.addr(deployerKey);

        config.collateralToken = vm.envAddress("COLLATERAL_TOKEN");
        config.debtToken = vm.envAddress("DEBT_TOKEN");
        config.collateralFeed = vm.envAddress("COLLATERAL_FEED");
        config.debtFeed = vm.envAddress("DEBT_FEED");

        config.finalOwner = vm.envOr("PROTOCOL_OWNER", deployer);

        config.maxPriceAge = uint32(vm.envOr("MAX_PRICE_AGE", uint256(3600)));
        config.minDeposit = vm.envOr("MIN_DEPOSIT", uint256(1e6));
        config.reserveFactorBps = uint16(vm.envOr("RESERVE_FACTOR_BPS", uint256(1000)));

        config.maxLtvBps = uint16(vm.envOr("MAX_LTV_BPS", uint256(7500)));
        config.liquidationThresholdBps = uint16(vm.envOr("LIQUIDATION_THRESHOLD_BPS", uint256(8000)));
        config.liquidationBonusBps = uint16(vm.envOr("LIQUIDATION_BONUS_BPS", uint256(500)));

        config.baseRateAprBps = uint16(vm.envOr("BASE_RATE_APR_BPS", uint256(100)));
        config.slopeBelowKinkAprBps = uint16(vm.envOr("SLOPE_BELOW_KINK_APR_BPS", uint256(654)));
        config.slopeAboveKinkAprBps = uint32(vm.envOr("SLOPE_ABOVE_KINK_APR_BPS", uint256(20000)));
        config.kinkUtilizationBps = uint16(vm.envOr("KINK_UTILIZATION_BPS", uint256(8000)));

        return config;
    }

    function _report(Deployment memory deployed) private pure {
        console2.log("=== deployed addresses ===");
        console2.log("PriceOracleAdapter  ", address(deployed.oracle));
        console2.log("InterestRateModel   ", address(deployed.rateModel));
        console2.log("LendingPool         ", address(deployed.pool));
        console2.log("CollateralVault     ", address(deployed.vault));
        console2.log("LendingController   ", address(deployed.controller));
        console2.log("LiquidationManager  ", address(deployed.manager));
        console2.log("PositionLens        ", address(deployed.lens));

        console2.log("=== copy into .env ===");
        console2.log(string.concat("ORACLE_ADDRESS=", vm.toString(address(deployed.oracle))));
        console2.log(string.concat("RATE_MODEL_ADDRESS=", vm.toString(address(deployed.rateModel))));
        console2.log(string.concat("POOL_ADDRESS=", vm.toString(address(deployed.pool))));
        console2.log(string.concat("VAULT_ADDRESS=", vm.toString(address(deployed.vault))));
        console2.log(string.concat("CONTROLLER_ADDRESS=", vm.toString(address(deployed.controller))));
        console2.log(string.concat("LIQUIDATION_MANAGER_ADDRESS=", vm.toString(address(deployed.manager))));
        console2.log(string.concat("LENS_ADDRESS=", vm.toString(address(deployed.lens))));

        console2.log("wiring and ownership verified");
    }
}
