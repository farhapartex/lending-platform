import { AssetSymbol, WalletStatus } from "@/lib/enums";
import { assetDecimals } from "@/content/protocol";
import { parseTokenAmount } from "@/lib/token";

export const walletStatus: WalletStatus = WalletStatus.Connected;

export const walletAddress = "0x7Ac49c1f2B5D8e6A3049F1b8C2d7E5a0B4f63f21";

const usdcDecimals = assetDecimals[AssetSymbol.Usdc];
const wethDecimals = assetDecimals[AssetSymbol.Weth];

export const walletBalances: Record<AssetSymbol, bigint> = {
  [AssetSymbol.Usdc]: parseTokenAmount("12480.50", usdcDecimals) ?? 0n,
  [AssetSymbol.Weth]: parseTokenAmount("4.812", wethDecimals) ?? 0n,
};

export const usdcAllowance: bigint = 0n;
