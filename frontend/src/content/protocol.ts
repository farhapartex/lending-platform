import {
  AssetRole,
  AssetSymbol,
  FeeKind,
  MarketMetricKey,
  ValueFormat,
} from "@/lib/enums";
import { toPriceScaled } from "@/lib/health";

export type MarketMetric = {
  key: MarketMetricKey;
  label: string;
  value: number;
  format: ValueFormat;
  hint: string;
};

export type MarketAsset = {
  symbol: AssetSymbol;
  role: AssetRole;
  name: string;
};

export type FeeItem = {
  kind: FeeKind;
  label: string;
  value: string;
  description: string;
};


export const explorerTxBaseUrl = "https://sepolia.etherscan.io/tx/";

export const marketAssets: MarketAsset[] = [
  { symbol: AssetSymbol.Weth, role: AssetRole.Collateral, name: "Wrapped Ether" },
  { symbol: AssetSymbol.Usdc, role: AssetRole.Borrowable, name: "USD Coin" },
];

export const assetDecimals: Record<AssetSymbol, number> = {
  [AssetSymbol.Weth]: 18,
  [AssetSymbol.Usdc]: 6,
};

export const assetPrices: Record<AssetSymbol, number> = {
  [AssetSymbol.Weth]: 3412.58,
  [AssetSymbol.Usdc]: 1,
};

export const borrowAprRate = 0.0635;

export const collateralAsset = AssetSymbol.Weth;

export const debtAsset = AssetSymbol.Usdc;

export const collateralDecimals = assetDecimals[AssetSymbol.Weth];

export const debtDecimals = assetDecimals[AssetSymbol.Usdc];

export const collateralUnitPriceScaled = toPriceScaled(assetPrices[AssetSymbol.Weth]);

export const debtUnitPriceScaled = toPriceScaled(assetPrices[AssetSymbol.Usdc]);

export const maxLtvBps = 7_500n;

export const liquidationThresholdBps = 8_000n;

export const recommendedLtvBps = 5_500n;

export const liquidationBonusBps = 500n;

export const availableLiquidity = 16_660_000;


export const feeItems: FeeItem[] = [
  {
    kind: FeeKind.InterestSpread,
    label: "Interest spread",
    value: "10% of interest paid",
    description: "The only ongoing fee. Taken from borrower interest, never from your deposited principal.",
  },
  {
    kind: FeeKind.LiquidationBonus,
    label: "Liquidation bonus",
    value: "5% of debt repaid",
    description: "Paid to whoever resolves an unsafe position. Applies only if your position reaches the threshold.",
  },
  {
    kind: FeeKind.PlatformFee,
    label: "Deposit & withdrawal fees",
    value: "None",
    description: "No fee to deposit, withdraw, borrow, or repay. You pay only network gas.",
  },
];
