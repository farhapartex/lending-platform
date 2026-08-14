// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

import {IInterestRateModel} from "./interfaces/IInterestRateModel.sol";
import {ILendingPool} from "./interfaces/ILendingPool.sol";
import {Errors} from "./libraries/Errors.sol";
import {InterestMath} from "./libraries/InterestMath.sol";
import {ShareMath} from "./libraries/ShareMath.sol";
import {WadMath} from "./libraries/WadMath.sol";

contract LendingPool is ILendingPool, Ownable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    event Deposit(
        address indexed lender, uint256 assets, uint256 shares, uint256 supplyIndex, uint256 totalSupplied
    );
    event Withdraw(
        address indexed lender, uint256 assets, uint256 shares, uint256 supplyIndex, uint256 totalSupplied
    );
    event InterestAccrued(
        uint256 supplyIndex, uint256 borrowIndex, uint256 totalBorrowed, uint256 reservesAccrued
    );
    event MinDepositChanged(uint256 previousAmount, uint256 newAmount);
    event DepositsPausedChanged(bool isPaused);
    event ReserveFactorChanged(uint16 previousBps, uint16 newBps);
    event RateModelChanged(address previousModel, address newModel);
    event ControllerLinked(address controller);
    event LiquidationManagerLinked(address liquidationManager);
    event ReservesCollected(address indexed recipient, uint256 amount, uint256 remainingReserves);

    uint16 private constant MAX_RESERVE_FACTOR_BPS = 5000;

    IERC20 public immutable assetToken;

    uint128 private supplyIndexValue;
    uint128 private borrowIndexValue;

    uint40 private lastAccrualAt;
    uint16 public reserveFactorBps;
    bool public depositsPaused;

    IInterestRateModel public rateModel;

    address public controller;

    address public liquidationManager;

    uint256 public totalSupplyShares;

    uint256 public totalDebtShares;

    uint256 public accruedReserves;

    uint256 public minDeposit;

    mapping(address => uint256) public sharesOf;

    mapping(address => uint256) public debtSharesOf;

    constructor(
        address owner,
        IERC20 poolAsset,
        IInterestRateModel startingRateModel,
        uint256 startingMinDeposit,
        uint16 startingReserveFactorBps
    ) Ownable(owner) {
        if (address(poolAsset) == address(0) || address(startingRateModel) == address(0)) {
            revert Errors.ZeroAddress();
        }

        if (startingReserveFactorBps > MAX_RESERVE_FACTOR_BPS) {
            revert Errors.InvalidRiskSettings();
        }

        assetToken = poolAsset;
        rateModel = startingRateModel;
        supplyIndexValue = uint128(ShareMath.STARTING_INDEX);
        borrowIndexValue = uint128(ShareMath.STARTING_INDEX);
        lastAccrualAt = uint40(block.timestamp);
        minDeposit = startingMinDeposit;
        reserveFactorBps = startingReserveFactorBps;

        emit MinDepositChanged(0, startingMinDeposit);
        emit ReserveFactorChanged(0, startingReserveFactorBps);
        emit RateModelChanged(address(0), address(startingRateModel));
    }

    function asset() external view returns (address) {
        return address(assetToken);
    }

    function linkController(address newController) external onlyOwner {
        if (newController == address(0)) {
            revert Errors.ZeroAddress();
        }

        if (controller != address(0)) {
            revert Errors.AlreadyInitialized();
        }

        controller = newController;

        emit ControllerLinked(newController);
    }

    function linkLiquidationManager(address newLiquidationManager) external onlyOwner {
        if (newLiquidationManager == address(0)) {
            revert Errors.ZeroAddress();
        }

        if (liquidationManager != address(0)) {
            revert Errors.AlreadyInitialized();
        }

        liquidationManager = newLiquidationManager;

        emit LiquidationManagerLinked(newLiquidationManager);
    }

    function debtOf(address borrower) public view returns (uint256) {
        return ShareMath.assetsFromSharesUp(debtSharesOf[borrower], borrowIndexValue);
    }

    function borrowFor(address borrower, address recipient, uint256 amount) external nonReentrant {
        _requireController();

        accrueInterest();

        uint256 liquidity = availableLiquidity();
        if (amount > liquidity) {
            revert Errors.NotEnoughLiquidity(amount, liquidity);
        }

        uint256 shares = ShareMath.sharesFromAssetsUp(amount, borrowIndexValue);

        debtSharesOf[borrower] += shares;
        totalDebtShares += shares;

        assetToken.safeTransfer(recipient, amount);
    }

    function repayFor(address borrower, address payer, uint256 amount) external nonReentrant {
        _requireController();

        accrueInterest();

        uint256 shares = ShareMath.sharesFromAssetsDown(amount, borrowIndexValue);
        uint256 ownedShares = debtSharesOf[borrower];

        if (shares > ownedShares) {
            shares = ownedShares;
        }

        debtSharesOf[borrower] = ownedShares - shares;
        totalDebtShares -= shares;

        assetToken.safeTransferFrom(payer, address(this), amount);
    }

    function repayAllFor(address borrower, address payer) external nonReentrant returns (uint256 amountPaid) {
        _requireDebtManager();

        accrueInterest();

        uint256 ownedShares = debtSharesOf[borrower];
        amountPaid = ShareMath.assetsFromSharesUp(ownedShares, borrowIndexValue);

        debtSharesOf[borrower] = 0;
        totalDebtShares -= ownedShares;

        assetToken.safeTransferFrom(payer, address(this), amountPaid);

        return amountPaid;
    }

    function _requireController() private view {
        if (msg.sender != controller) {
            revert Errors.NotAuthorized(msg.sender);
        }
    }

    function _requireDebtManager() private view {
        if (msg.sender != controller && msg.sender != liquidationManager) {
            revert Errors.NotAuthorized(msg.sender);
        }
    }

    function supplyIndex() public view returns (uint256) {
        return supplyIndexValue;
    }

    function borrowIndex() public view returns (uint256) {
        return borrowIndexValue;
    }

    function lastAccrualTimestamp() external view returns (uint256) {
        return lastAccrualAt;
    }

    function totalSupplied() public view returns (uint256) {
        return ShareMath.assetsFromSharesDown(totalSupplyShares, supplyIndexValue);
    }

    function totalBorrowed() public view returns (uint256) {
        return ShareMath.assetsFromSharesUp(totalDebtShares, borrowIndexValue);
    }

    function availableLiquidity() public view returns (uint256) {
        return WadMath.subtractOrZero(totalSupplied(), totalBorrowed());
    }

    function utilizationBps() public view returns (uint256) {
        return rateModel.utilizationBps(totalSupplied(), totalBorrowed());
    }

    function balanceOfAssets(address lender) public view returns (uint256) {
        return ShareMath.assetsFromSharesDown(sharesOf[lender], supplyIndexValue);
    }

    function maxWithdrawable(address lender) external view returns (uint256) {
        return WadMath.smaller(balanceOfAssets(lender), availableLiquidity());
    }

    function setMinDeposit(uint256 newAmount) external onlyOwner {
        uint256 previousAmount = minDeposit;
        minDeposit = newAmount;

        emit MinDepositChanged(previousAmount, newAmount);
    }

    function setDepositsPaused(bool isPaused) external onlyOwner {
        depositsPaused = isPaused;

        emit DepositsPausedChanged(isPaused);
    }

    function setReserveFactorBps(uint16 newBps) external onlyOwner {
        if (newBps > MAX_RESERVE_FACTOR_BPS) {
            revert Errors.InvalidRiskSettings();
        }

        accrueInterest();

        uint16 previousBps = reserveFactorBps;
        reserveFactorBps = newBps;

        emit ReserveFactorChanged(previousBps, newBps);
    }

    function setRateModel(IInterestRateModel newModel) external onlyOwner {
        if (address(newModel) == address(0)) {
            revert Errors.ZeroAddress();
        }

        accrueInterest();

        address previousModel = address(rateModel);
        rateModel = newModel;

        emit RateModelChanged(previousModel, address(newModel));
    }

    function collectReserves(address recipient, uint256 amount)
        external
        onlyOwner
        nonReentrant
        returns (uint256 collected)
    {
        if (amount == 0) {
            revert Errors.ZeroAmount();
        }

        return _collectReserves(recipient, amount);
    }

    function collectAllReserves(address recipient)
        external
        onlyOwner
        nonReentrant
        returns (uint256 collected)
    {
        accrueInterest();

        uint256 available = accruedReserves;

        if (available == 0) {
            revert Errors.ZeroAmount();
        }

        return _collectReserves(recipient, available);
    }

    function _collectReserves(address recipient, uint256 amount) private returns (uint256 collected) {
        if (recipient == address(0)) {
            revert Errors.ZeroAddress();
        }

        accrueInterest();

        uint256 available = accruedReserves;

        if (amount > available) {
            revert Errors.ExceedsReserves(amount, available);
        }

        uint256 remaining = available - amount;
        accruedReserves = remaining;

        assetToken.safeTransfer(recipient, amount);

        emit ReservesCollected(recipient, amount, remaining);

        return amount;
    }

    function previewAccrual()
        public
        view
        returns (uint256 nextSupplyIndex, uint256 nextBorrowIndex, uint256 reservesToAdd)
    {
        nextSupplyIndex = supplyIndexValue;
        nextBorrowIndex = borrowIndexValue;

        uint256 elapsedSeconds = block.timestamp - lastAccrualAt;

        if (elapsedSeconds == 0) {
            return (nextSupplyIndex, nextBorrowIndex, 0);
        }

        uint256 borrowedBefore = ShareMath.assetsFromSharesUp(totalDebtShares, nextBorrowIndex);

        if (borrowedBefore == 0) {
            return (nextSupplyIndex, nextBorrowIndex, 0);
        }

        uint256 suppliedBefore = ShareMath.assetsFromSharesDown(totalSupplyShares, nextSupplyIndex);
        uint256 ratePerSecond =
            rateModel.borrowRatePerSecond(rateModel.utilizationBps(suppliedBefore, borrowedBefore));

        uint256 growth = InterestMath.growthFactor(ratePerSecond, elapsedSeconds);

        if (growth == 0) {
            return (nextSupplyIndex, nextBorrowIndex, 0);
        }

        uint256 interest = InterestMath.interestOnDebt(borrowedBefore, growth);
        (uint256 reserveShare, uint256 lendersShare) =
            InterestMath.splitForReserves(interest, reserveFactorBps);

        return (
            InterestMath.grownSupplyIndex(nextSupplyIndex, lendersShare, suppliedBefore),
            InterestMath.grownBorrowIndex(nextBorrowIndex, growth),
            reserveShare
        );
    }

    function accrueInterest() public {
        if (block.timestamp == lastAccrualAt) {
            return;
        }

        (uint256 nextSupplyIndex, uint256 nextBorrowIndex, uint256 reservesToAdd) = previewAccrual();

        lastAccrualAt = uint40(block.timestamp);

        if (nextBorrowIndex == borrowIndexValue) {
            return;
        }

        borrowIndexValue = SafeCast.toUint128(nextBorrowIndex);
        supplyIndexValue = SafeCast.toUint128(nextSupplyIndex);
        accruedReserves += reservesToAdd;

        emit InterestAccrued(nextSupplyIndex, nextBorrowIndex, totalBorrowed(), reservesToAdd);
    }

    function deposit(uint256 assets) external nonReentrant returns (uint256 shares) {
        if (depositsPaused) {
            revert Errors.MarketPaused();
        }

        if (assets == 0) {
            revert Errors.ZeroAmount();
        }

        accrueInterest();

        uint256 requiredMinimum = minDeposit;
        if (assets < requiredMinimum) {
            revert Errors.BelowMinimumDeposit(assets, requiredMinimum);
        }

        uint256 currentIndex = supplyIndexValue;
        shares = ShareMath.sharesFromAssetsDown(assets, currentIndex);

        if (shares == 0) {
            revert Errors.ZeroAmount();
        }

        totalSupplyShares += shares;
        sharesOf[msg.sender] += shares;

        assetToken.safeTransferFrom(msg.sender, address(this), assets);

        emit Deposit(msg.sender, assets, shares, currentIndex, totalSupplied());

        return shares;
    }

    function withdraw(uint256 assets) external nonReentrant returns (uint256 shares) {
        if (assets == 0) {
            revert Errors.ZeroAmount();
        }

        accrueInterest();

        uint256 currentIndex = supplyIndexValue;
        shares = ShareMath.sharesFromAssetsUp(assets, currentIndex);

        _settleWithdraw(msg.sender, shares, assets, currentIndex);

        return shares;
    }

    function redeemShares(uint256 shares) external nonReentrant returns (uint256 assets) {
        if (shares == 0) {
            revert Errors.ZeroAmount();
        }

        accrueInterest();

        uint256 currentIndex = supplyIndexValue;
        assets = ShareMath.assetsFromSharesDown(shares, currentIndex);

        if (assets == 0) {
            revert Errors.ZeroAmount();
        }

        _settleWithdraw(msg.sender, shares, assets, currentIndex);

        return assets;
    }

    function _settleWithdraw(address lender, uint256 shares, uint256 assets, uint256 currentIndex) private {
        uint256 ownedShares = sharesOf[lender];

        if (shares > ownedShares) {
            revert Errors.ExceedsSupplyBalance(
                assets, ShareMath.assetsFromSharesDown(ownedShares, currentIndex)
            );
        }

        uint256 liquidity = availableLiquidity();
        if (assets > liquidity) {
            revert Errors.NotEnoughLiquidity(assets, liquidity);
        }

        sharesOf[lender] = ownedShares - shares;
        totalSupplyShares -= shares;

        assetToken.safeTransfer(lender, assets);

        emit Withdraw(lender, assets, shares, currentIndex, totalSupplied());
    }
}
