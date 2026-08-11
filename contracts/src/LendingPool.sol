// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {ILendingPool} from "./interfaces/ILendingPool.sol";
import {Errors} from "./libraries/Errors.sol";
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
    event MinDepositChanged(uint256 previousAmount, uint256 newAmount);
    event DepositsPausedChanged(bool isPaused);

    IERC20 public immutable assetToken;

    uint128 private supplyIndexValue;

    bool public depositsPaused;

    uint256 public totalSupplyShares;

    uint256 public totalBorrowed;

    uint256 public minDeposit;

    mapping(address => uint256) public sharesOf;

    constructor(address owner, IERC20 poolAsset, uint256 startingMinDeposit) Ownable(owner) {
        if (address(poolAsset) == address(0)) {
            revert Errors.ZeroAddress();
        }

        assetToken = poolAsset;
        supplyIndexValue = uint128(ShareMath.STARTING_INDEX);
        minDeposit = startingMinDeposit;

        emit MinDepositChanged(0, startingMinDeposit);
    }

    function asset() external view returns (address) {
        return address(assetToken);
    }

    function supplyIndex() public view returns (uint256) {
        return supplyIndexValue;
    }

    function totalSupplied() public view returns (uint256) {
        return ShareMath.assetsFromSharesDown(totalSupplyShares, supplyIndexValue);
    }

    function availableLiquidity() public view returns (uint256) {
        return WadMath.subtractOrZero(totalSupplied(), totalBorrowed);
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

    function deposit(uint256 assets) external nonReentrant returns (uint256 shares) {
        if (depositsPaused) {
            revert Errors.MarketPaused();
        }

        if (assets == 0) {
            revert Errors.ZeroAmount();
        }

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

        uint256 currentIndex = supplyIndexValue;
        shares = ShareMath.sharesFromAssetsUp(assets, currentIndex);

        _settleWithdraw(msg.sender, shares, assets, currentIndex);

        return shares;
    }

    function redeemShares(uint256 shares) external nonReentrant returns (uint256 assets) {
        if (shares == 0) {
            revert Errors.ZeroAmount();
        }

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
