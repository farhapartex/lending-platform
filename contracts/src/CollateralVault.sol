// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {ICollateralVault} from "./interfaces/ICollateralVault.sol";
import {ILendingController} from "./interfaces/ILendingController.sol";
import {Errors} from "./libraries/Errors.sol";

contract CollateralVault is ICollateralVault, Ownable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    event CollateralDeposited(address indexed borrower, uint256 amount, uint256 newCollateral);
    event CollateralWithdrawn(address indexed borrower, uint256 amount, uint256 newCollateral);
    event CollateralSeized(
        address indexed borrower, address indexed recipient, uint256 amount, uint256 newCollateral
    );
    event ControllerLinked(address controller);
    event LiquidationManagerLinked(address liquidationManager);

    IERC20 public immutable collateralToken;

    ILendingController public controller;

    address public liquidationManager;

    uint256 public totalCollateral;

    mapping(address => uint256) public collateralOf;

    constructor(address owner, IERC20 vaultAsset) Ownable(owner) {
        if (address(vaultAsset) == address(0)) {
            revert Errors.ZeroAddress();
        }

        collateralToken = vaultAsset;
    }

    function collateralAsset() external view returns (address) {
        return address(collateralToken);
    }

    function linkController(ILendingController newController) external onlyOwner {
        if (address(newController) == address(0)) {
            revert Errors.ZeroAddress();
        }

        if (address(controller) != address(0)) {
            revert Errors.AlreadyInitialized();
        }

        controller = newController;

        emit ControllerLinked(address(newController));
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

    function depositCollateral(uint256 amount) external nonReentrant {
        if (amount == 0) {
            revert Errors.ZeroAmount();
        }

        uint256 newCollateral = collateralOf[msg.sender] + amount;

        collateralOf[msg.sender] = newCollateral;
        totalCollateral += amount;

        collateralToken.safeTransferFrom(msg.sender, address(this), amount);

        emit CollateralDeposited(msg.sender, amount, newCollateral);
    }

    function withdrawCollateral(uint256 amount) external nonReentrant {
        if (amount == 0) {
            revert Errors.ZeroAmount();
        }

        uint256 held = collateralOf[msg.sender];

        if (amount > held) {
            revert Errors.ExceedsCollateralBalance(amount, held);
        }

        uint256 safeAmount = _safeWithdrawAmount(msg.sender);

        if (amount > safeAmount) {
            revert Errors.WouldBreakBorrowLimit(amount, safeAmount);
        }

        uint256 newCollateral = held - amount;

        collateralOf[msg.sender] = newCollateral;
        totalCollateral -= amount;

        collateralToken.safeTransfer(msg.sender, amount);

        emit CollateralWithdrawn(msg.sender, amount, newCollateral);
    }

    function seize(address borrower, address recipient, uint256 amount) external nonReentrant {
        if (msg.sender != liquidationManager) {
            revert Errors.NotAuthorized(msg.sender);
        }

        if (recipient == address(0)) {
            revert Errors.ZeroAddress();
        }

        uint256 held = collateralOf[borrower];

        if (amount > held) {
            revert Errors.ExceedsCollateralBalance(amount, held);
        }

        uint256 newCollateral = held - amount;

        collateralOf[borrower] = newCollateral;
        totalCollateral -= amount;

        collateralToken.safeTransfer(recipient, amount);

        emit CollateralSeized(borrower, recipient, amount, newCollateral);
    }

    function _safeWithdrawAmount(address borrower) private view returns (uint256) {
        ILendingController linkedController = controller;

        if (address(linkedController) == address(0)) {
            revert Errors.NotAuthorized(address(0));
        }

        return linkedController.maxWithdrawableCollateral(borrower);
    }
}
