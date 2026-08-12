// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

import {ICollateralVault} from "./interfaces/ICollateralVault.sol";
import {ILendingController} from "./interfaces/ILendingController.sol";
import {ILendingPool} from "./interfaces/ILendingPool.sol";
import {IPriceOracle} from "./interfaces/IPriceOracle.sol";
import {Errors} from "./libraries/Errors.sol";
import {HealthMath} from "./libraries/HealthMath.sol";
import {ShareMath} from "./libraries/ShareMath.sol";
import {WadMath} from "./libraries/WadMath.sol";

contract LendingController is ILendingController, Ownable, ReentrancyGuard {
    event Borrow(
        address indexed borrower,
        uint256 amount,
        uint256 newDebt,
        uint256 healthFactorBps,
        uint256 borrowIndex
    );
    event Repay(address indexed borrower, address indexed payer, uint256 amount, uint256 newDebt);
    event RiskSettingsChanged(uint16 maxLtvBps, uint16 liquidationThresholdBps, uint16 liquidationBonusBps);
    event BorrowPausedChanged(bool isPaused);

    ILendingPool public immutable pool;
    ICollateralVault public immutable vault;
    IPriceOracle public immutable oracle;

    address public immutable collateralAsset;
    address public immutable debtAsset;

    uint8 public immutable collateralDecimals;
    uint8 public immutable debtDecimals;

    uint16 public maxLtvBps;
    uint16 public liquidationThresholdBps;
    uint16 public liquidationBonusBps;
    bool public borrowPaused;

    constructor(
        address owner,
        ILendingPool lendingPool,
        ICollateralVault collateralVault,
        IPriceOracle priceOracle,
        uint16 startingMaxLtvBps,
        uint16 startingLiquidationThresholdBps,
        uint16 startingLiquidationBonusBps
    ) Ownable(owner) {
        if (
            address(lendingPool) == address(0) || address(collateralVault) == address(0)
                || address(priceOracle) == address(0)
        ) {
            revert Errors.ZeroAddress();
        }

        pool = lendingPool;
        vault = collateralVault;
        oracle = priceOracle;

        collateralAsset = collateralVault.collateralAsset();
        debtAsset = lendingPool.asset();

        collateralDecimals = IERC20Metadata(collateralAsset).decimals();
        debtDecimals = IERC20Metadata(debtAsset).decimals();

        _storeRiskSettings(startingMaxLtvBps, startingLiquidationThresholdBps, startingLiquidationBonusBps);
    }

    function setRiskSettings(uint16 newMaxLtvBps, uint16 newThresholdBps, uint16 newBonusBps)
        external
        onlyOwner
    {
        pool.accrueInterest();

        _storeRiskSettings(newMaxLtvBps, newThresholdBps, newBonusBps);
    }

    function setBorrowPaused(bool isPaused) external onlyOwner {
        borrowPaused = isPaused;

        emit BorrowPausedChanged(isPaused);
    }

    function debtOf(address borrower) public view returns (uint256) {
        return _currentDebt(borrower);
    }

    function collateralValueOf(address borrower) public view returns (uint256) {
        (uint256 price,) = oracle.getPrice(collateralAsset);

        return HealthMath.valueOfAmount(vault.collateralOf(borrower), collateralDecimals, price);
    }

    function debtValueOf(address borrower) public view returns (uint256) {
        (uint256 price,) = oracle.getPrice(debtAsset);

        return HealthMath.valueOfAmount(_currentDebt(borrower), debtDecimals, price);
    }

    function healthFactorBps(address borrower) public view returns (uint256) {
        return
            HealthMath.healthFactor(
                collateralValueOf(borrower), debtValueOf(borrower), liquidationThresholdBps
            );
    }

    function isLiquidatable(address borrower) external view returns (bool) {
        return HealthMath.isLiquidatable(healthFactorBps(borrower));
    }

    function maxBorrowable(address borrower) public view returns (uint256) {
        uint256 fromCollateral = _borrowRoomFromCollateral(borrower, _currentDebt(borrower));

        return WadMath.smaller(fromCollateral, pool.availableLiquidity());
    }

    function maxWithdrawableCollateral(address borrower) external view returns (uint256) {
        uint256 owed = _currentDebt(borrower);
        uint256 held = vault.collateralOf(borrower);

        if (owed == 0) {
            return held;
        }

        (uint256 collateralPrice,) = oracle.getPrice(collateralAsset);
        (uint256 debtPrice,) = oracle.getPrice(debtAsset);

        uint256 debtValue = HealthMath.valueOfAmount(owed, debtDecimals, debtPrice);

        return
            HealthMath.freeCollateralAmount(held, collateralDecimals, collateralPrice, debtValue, maxLtvBps);
    }

    function borrow(uint256 amount) external nonReentrant {
        if (borrowPaused) {
            revert Errors.MarketPaused();
        }

        if (amount == 0) {
            revert Errors.ZeroAmount();
        }

        pool.accrueInterest();

        uint256 collateralRoom = _borrowRoomFromCollateral(msg.sender, pool.debtOf(msg.sender));
        if (amount > collateralRoom) {
            revert Errors.ExceedsBorrowLimit(amount, collateralRoom);
        }

        uint256 liquidity = pool.availableLiquidity();
        if (amount > liquidity) {
            revert Errors.NotEnoughLiquidity(amount, liquidity);
        }

        pool.borrowFor(msg.sender, msg.sender, amount);

        emit Borrow(
            msg.sender, amount, pool.debtOf(msg.sender), healthFactorBps(msg.sender), pool.borrowIndex()
        );
    }

    function repay(uint256 amount) external nonReentrant {
        if (amount == 0) {
            revert Errors.ZeroAmount();
        }

        pool.accrueInterest();

        uint256 owed = pool.debtOf(msg.sender);

        if (amount > owed) {
            revert Errors.ExceedsDebt(amount, owed);
        }

        pool.repayFor(msg.sender, msg.sender, amount);

        emit Repay(msg.sender, msg.sender, amount, pool.debtOf(msg.sender));
    }

    function repayAll() external nonReentrant returns (uint256 amountPaid) {
        pool.accrueInterest();

        if (pool.debtOf(msg.sender) == 0) {
            revert Errors.ZeroAmount();
        }

        amountPaid = pool.repayAllFor(msg.sender, msg.sender);

        emit Repay(msg.sender, msg.sender, amountPaid, 0);

        return amountPaid;
    }

    function _borrowRoomFromCollateral(address borrower, uint256 owed) private view returns (uint256) {
        (uint256 collateralPrice,) = oracle.getPrice(collateralAsset);
        (uint256 debtPrice,) = oracle.getPrice(debtAsset);

        uint256 collateralValue =
            HealthMath.valueOfAmount(vault.collateralOf(borrower), collateralDecimals, collateralPrice);
        uint256 debtValue = HealthMath.valueOfAmount(owed, debtDecimals, debtPrice);

        return
            HealthMath.remainingBorrowAmount(collateralValue, debtValue, maxLtvBps, debtDecimals, debtPrice);
    }

    function _currentDebt(address borrower) private view returns (uint256) {
        (, uint256 borrowIndexNow,) = pool.previewAccrual();

        return ShareMath.assetsFromSharesUp(pool.debtSharesOf(borrower), borrowIndexNow);
    }

    function _storeRiskSettings(uint16 newMaxLtvBps, uint16 newThresholdBps, uint16 newBonusBps) private {
        bool validLtv = newMaxLtvBps > 0 && newMaxLtvBps < newThresholdBps;
        bool validThreshold = newThresholdBps <= WadMath.FULL_PERCENT_BPS;
        bool validBonus = newBonusBps <= WadMath.FULL_PERCENT_BPS;

        if (!validLtv || !validThreshold || !validBonus) {
            revert Errors.InvalidRiskSettings();
        }

        maxLtvBps = newMaxLtvBps;
        liquidationThresholdBps = newThresholdBps;
        liquidationBonusBps = newBonusBps;

        emit RiskSettingsChanged(newMaxLtvBps, newThresholdBps, newBonusBps);
    }
}
