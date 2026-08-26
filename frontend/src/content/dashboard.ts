import { AssetSymbol } from "@/lib/enums";
import { assetDecimals } from "@/content/protocol";
import { parseTokenAmount } from "@/lib/token";

const usdcDecimals = assetDecimals[AssetSymbol.Usdc];
const wethDecimals = assetDecimals[AssetSymbol.Weth];

export type LiquidationEvent = {
  id: string;
  timestamp: string;
  debtRepaid: bigint;
  collateralSeized: bigint;
  bonusPaid: bigint;
  healthFactorAtLiquidation: string;
  triggerPrice: number;
};

export const recentLiquidation: LiquidationEvent | null = {
  id: "liq-1",
  timestamp: "2026-08-01T03:12:00Z",
  debtRepaid: parseTokenAmount("2100", usdcDecimals) ?? 0n,
  collateralSeized: parseTokenAmount("0.6455", wethDecimals) ?? 0n,
  bonusPaid: parseTokenAmount("105", usdcDecimals) ?? 0n,
  healthFactorAtLiquidation: "0.98",
  triggerPrice: 3102.44,
};

export const dashboardContent = {
  title: "Your dashboard",
  description: "Everything you have on the platform in one place, with your safety score front and centre.",
  overviewTitle: "Portfolio",
  positionsTitle: "Your positions",
  activityTitle: "Recent activity",
  activityDescription: "The last few things that happened to your positions.",
  activityEmptyTitle: "Nothing has happened yet",
  activityEmptyDescription:
    "Your deposits, loans, repayments, and any liquidations will show up here.",
  activityNotIndexedTitle: "Activity history is not ready yet",
  activityNotIndexedDescription:
    "Your positions above are read straight from the blockchain and are correct. This panel replays past events, and that record is still being built.",
  activityUnavailableTitle: "Recent activity is unavailable",
  activityUnavailableDescription:
    "We could not reach the service that keeps your history. Everything above is read from the blockchain and is unaffected.",
  activityRetry: "Try again",
  simulatorTrigger: "Test a price drop",
  emptyTitle: "You do not have any positions yet",
  emptyDescription:
    "Deposit USDC to start earning interest, or post WETH as collateral to borrow against it. Both take a couple of minutes.",
} as const;
