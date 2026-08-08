import { ActivityKind, AssetSymbol } from "@/lib/enums";
import { assetDecimals } from "@/content/protocol";
import { parseTokenAmount } from "@/lib/token";

const usdcDecimals = assetDecimals[AssetSymbol.Usdc];
const wethDecimals = assetDecimals[AssetSymbol.Weth];

export type ActivityEntry = {
  id: string;
  kind: ActivityKind;
  amount: bigint;
  symbol: AssetSymbol;
  decimals: number;
  timestamp: string;
};

export type LiquidationEvent = {
  id: string;
  timestamp: string;
  debtRepaid: bigint;
  collateralSeized: bigint;
  bonusPaid: bigint;
  healthFactorAtLiquidation: string;
  triggerPrice: number;
};

export const recentActivity: ActivityEntry[] = [
  {
    id: "act-1",
    kind: ActivityKind.Borrow,
    amount: parseTokenAmount("1500", usdcDecimals) ?? 0n,
    symbol: AssetSymbol.Usdc,
    decimals: usdcDecimals,
    timestamp: "2026-08-07T14:22:00Z",
  },
  {
    id: "act-2",
    kind: ActivityKind.CollateralAdded,
    amount: parseTokenAmount("0.75", wethDecimals) ?? 0n,
    symbol: AssetSymbol.Weth,
    decimals: wethDecimals,
    timestamp: "2026-08-06T09:05:00Z",
  },
  {
    id: "act-3",
    kind: ActivityKind.Deposit,
    amount: parseTokenAmount("10000", usdcDecimals) ?? 0n,
    symbol: AssetSymbol.Usdc,
    decimals: usdcDecimals,
    timestamp: "2026-08-04T18:40:00Z",
  },
  {
    id: "act-4",
    kind: ActivityKind.Liquidation,
    amount: parseTokenAmount("2100", usdcDecimals) ?? 0n,
    symbol: AssetSymbol.Usdc,
    decimals: usdcDecimals,
    timestamp: "2026-08-01T03:12:00Z",
  },
  {
    id: "act-5",
    kind: ActivityKind.Repay,
    amount: parseTokenAmount("640", usdcDecimals) ?? 0n,
    symbol: AssetSymbol.Usdc,
    decimals: usdcDecimals,
    timestamp: "2026-07-29T11:58:00Z",
  },
];

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
  simulatorTrigger: "Test a price drop",
  emptyTitle: "You do not have any positions yet",
  emptyDescription:
    "Deposit USDC to start earning interest, or post WETH as collateral to borrow against it. Both take a couple of minutes.",
} as const;
