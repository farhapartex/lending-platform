import { AssetSymbol } from "@/lib/enums";
import { assetDecimals } from "@/content/protocol";
import { parseTokenAmount } from "@/lib/token";

const usdcDecimals = assetDecimals[AssetSymbol.Usdc];
const wethDecimals = assetDecimals[AssetSymbol.Weth];

export const walletBalances: Record<AssetSymbol, bigint> = {
  [AssetSymbol.Usdc]: parseTokenAmount("12480.50", usdcDecimals) ?? 0n,
  [AssetSymbol.Weth]: parseTokenAmount("4.812", wethDecimals) ?? 0n,
};

export const usdcAllowance: bigint = 0n;

export const wethAllowance: bigint = 0n;
