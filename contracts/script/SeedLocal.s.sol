// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Script, console2} from "forge-std/Script.sol";

import {MockAggregator} from "../src/mocks/MockAggregator.sol";
import {MockERC20} from "../src/mocks/MockERC20.sol";
import {AccountData, MarketData} from "../src/interfaces/IPositionLens.sol";
import {DeployConfig, Deployment, ProtocolDeployer} from "./Deployment.sol";

contract SeedLocal is Script {
    uint256 private constant ANVIL_CHAIN_ID = 31337;

    int256 private constant STARTING_ETH_PRICE = 341258000000;
    int256 private constant STRESSED_ETH_PRICE = 290000000000;
    int256 private constant USDC_PRICE = 100000000;

    error WrongChain(uint256 chainId);

    MockERC20 private weth;
    MockERC20 private usdc;
    MockAggregator private wethFeed;
    MockAggregator private usdcFeed;
    Deployment private deployed;

    uint256 private deployerKey;
    uint256 private lenderOneKey;
    uint256 private lenderTwoKey;
    uint256 private aliceKey;
    uint256 private bobKey;
    uint256 private carolKey;
    uint256 private liquidatorKey;

    function run() external {
        if (block.chainid != ANVIL_CHAIN_ID) {
            revert WrongChain(block.chainid);
        }

        _loadAccounts();
        _deployMocks();
        _deployProtocol();
        _fundAccounts();
        _supplyLiquidity();
        _openPositions();
        _stressPrice();
        _report();
    }

    function _loadAccounts() private {
        string memory mnemonic =
            vm.envOr("ANVIL_MNEMONIC", string("test test test test test test test test test test test junk"));

        deployerKey = vm.deriveKey(mnemonic, 0);
        lenderOneKey = vm.deriveKey(mnemonic, 1);
        lenderTwoKey = vm.deriveKey(mnemonic, 2);
        aliceKey = vm.deriveKey(mnemonic, 3);
        bobKey = vm.deriveKey(mnemonic, 4);
        carolKey = vm.deriveKey(mnemonic, 5);
        liquidatorKey = vm.deriveKey(mnemonic, 6);
    }

    function _deployMocks() private {
        vm.startBroadcast(deployerKey);

        weth = new MockERC20("Wrapped Ether", "WETH", 18);
        usdc = new MockERC20("USD Coin", "USDC", 6);

        wethFeed = new MockAggregator("ETH / USD", 8, STARTING_ETH_PRICE);
        usdcFeed = new MockAggregator("USDC / USD", 8, USDC_PRICE);

        vm.stopBroadcast();
    }

    function _deployProtocol() private {
        DeployConfig memory config = DeployConfig({
            finalOwner: vm.addr(deployerKey),
            collateralToken: address(weth),
            debtToken: address(usdc),
            collateralFeed: address(wethFeed),
            debtFeed: address(usdcFeed),
            maxPriceAge: 3600,
            minDeposit: 1e6,
            reserveFactorBps: 1000,
            maxLtvBps: 7500,
            liquidationThresholdBps: 8000,
            liquidationBonusBps: 500,
            baseRateAprBps: 100,
            slopeBelowKinkAprBps: 654,
            slopeAboveKinkAprBps: 20000,
            kinkUtilizationBps: 8000
        });

        ProtocolDeployer.requireEightDecimalFeeds(config);

        vm.startBroadcast(deployerKey);

        deployed = ProtocolDeployer.deployAll(config, vm.addr(deployerKey));
        ProtocolDeployer.wireAll(deployed, config);

        vm.stopBroadcast();

        ProtocolDeployer.verifyWiring(deployed, config);
        ProtocolDeployer.verifyOwnership(deployed, vm.addr(deployerKey));
    }

    function _fundAccounts() private {
        vm.startBroadcast(deployerKey);

        usdc.mint(vm.addr(lenderOneKey), 200_000e6);
        usdc.mint(vm.addr(lenderTwoKey), 100_000e6);
        usdc.mint(vm.addr(liquidatorKey), 50_000e6);

        weth.mint(vm.addr(aliceKey), 20e18);
        weth.mint(vm.addr(bobKey), 20e18);
        weth.mint(vm.addr(carolKey), 20e18);

        usdc.mint(vm.addr(aliceKey), 5_000e6);
        usdc.mint(vm.addr(bobKey), 5_000e6);
        usdc.mint(vm.addr(carolKey), 5_000e6);

        vm.stopBroadcast();
    }

    function _supplyLiquidity() private {
        _deposit(lenderOneKey, 100_000e6);
        _deposit(lenderTwoKey, 50_000e6);
    }

    function _openPositions() private {
        _openPosition(aliceKey, 5e18, 3_000e6);
        _openPosition(bobKey, 3.2e18, 6_900e6);
        _openPosition(carolKey, 2e18, 5_100e6);

        _approvePool(liquidatorKey);
    }

    function _stressPrice() private {
        vm.startBroadcast(deployerKey);

        wethFeed.setPrice(STRESSED_ETH_PRICE);
        usdcFeed.setPrice(USDC_PRICE);

        vm.stopBroadcast();
    }

    function _deposit(uint256 lenderKey, uint256 amount) private {
        vm.startBroadcast(lenderKey);

        usdc.approve(address(deployed.pool), type(uint256).max);
        deployed.pool.deposit(amount);

        vm.stopBroadcast();
    }

    function _openPosition(uint256 borrowerKey, uint256 collateral, uint256 debt) private {
        vm.startBroadcast(borrowerKey);

        weth.approve(address(deployed.vault), type(uint256).max);
        usdc.approve(address(deployed.pool), type(uint256).max);

        deployed.vault.depositCollateral(collateral);
        deployed.controller.borrow(debt);

        vm.stopBroadcast();
    }

    function _approvePool(uint256 accountKey) private {
        vm.startBroadcast(accountKey);

        usdc.approve(address(deployed.pool), type(uint256).max);

        vm.stopBroadcast();
    }

    function _report() private view {
        console2.log("=== tokens and feeds ===");
        console2.log("WETH                ", address(weth));
        console2.log("USDC                ", address(usdc));
        console2.log("ETH/USD feed        ", address(wethFeed));
        console2.log("USDC/USD feed       ", address(usdcFeed));

        console2.log("=== protocol ===");
        console2.log("PriceOracleAdapter  ", address(deployed.oracle));
        console2.log("InterestRateModel   ", address(deployed.rateModel));
        console2.log("LendingPool         ", address(deployed.pool));
        console2.log("CollateralVault     ", address(deployed.vault));
        console2.log("LendingController   ", address(deployed.controller));
        console2.log("LiquidationManager  ", address(deployed.manager));
        console2.log("PositionLens        ", address(deployed.lens));

        MarketData memory market = deployed.lens.marketData();

        console2.log("=== market ===");
        console2.log("totalSupplied       ", market.totalSupplied);
        console2.log("totalBorrowed       ", market.totalBorrowed);
        console2.log("availableLiquidity  ", market.availableLiquidity);
        console2.log("utilizationBps      ", market.utilizationBps);
        console2.log("borrowAprBps        ", market.borrowAprBps);
        console2.log("supplyAprBps        ", market.supplyAprBps);

        console2.log("=== positions at ETH 2900 ===");
        _reportAccount("alice     ", vm.addr(aliceKey));
        _reportAccount("bob       ", vm.addr(bobKey));
        _reportAccount("carol     ", vm.addr(carolKey));

        console2.log("=== accounts ===");
        console2.log("deployer / owner    ", vm.addr(deployerKey));
        console2.log("lender one          ", vm.addr(lenderOneKey));
        console2.log("lender two          ", vm.addr(lenderTwoKey));
        console2.log("alice   safe        ", vm.addr(aliceKey));
        console2.log("bob     caution     ", vm.addr(bobKey));
        console2.log("carol   liquidatable", vm.addr(carolKey));
        console2.log("liquidator          ", vm.addr(liquidatorKey));
    }

    function _reportAccount(string memory label, address account) private view {
        AccountData memory data = deployed.lens.accountData(account);

        console2.log(label);
        console2.log("  collateral        ", data.collateralAmount);
        console2.log("  debt              ", data.debtAmount);
        console2.log("  healthFactorBps   ", data.healthFactorBps);
        console2.log("  maxBorrowable     ", data.maxBorrowable);
        console2.log("  freeCollateral    ", data.maxWithdrawableCollateral);
        console2.log("  isLiquidatable    ", data.isLiquidatable);
    }
}
