import {
  AssetRole,
  AssetSymbol,
  DataStatus,
  FeeKind,
  MarketMetricKey,
  NetworkKind,
  OracleStatus,
  ProtocolStatKey,
  TrendDirection,
  ValueFormat,
} from "@/lib/enums";

export type ProtocolStat = {
  key: ProtocolStatKey;
  label: string;
  value: number;
  format: ValueFormat;
  trend: TrendDirection;
  trendLabel: string;
};

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

export type OracleReading = {
  symbol: AssetSymbol;
  price: number;
  updatedSecondsAgo: number;
  stalenessThresholdSeconds: number;
  status: OracleStatus;
};

export type UtilizationModel = {
  current: number;
  kink: number;
};

export const activeNetwork: NetworkKind = NetworkKind.Testnet;

export const marketDataStatus: DataStatus = DataStatus.Ready;

export const protocolStatsStatus: DataStatus = DataStatus.Ready;

export const protocolStats: ProtocolStat[] = [
  {
    key: ProtocolStatKey.TotalDeposited,
    label: "Total deposited",
    value: 48_240_000,
    format: ValueFormat.UsdCompact,
    trend: TrendDirection.Up,
    trendLabel: "+4.1% this week",
  },
  {
    key: ProtocolStatKey.TotalBorrowed,
    label: "Total borrowed",
    value: 31_580_000,
    format: ValueFormat.UsdCompact,
    trend: TrendDirection.Up,
    trendLabel: "+2.7% this week",
  },
  {
    key: ProtocolStatKey.Utilization,
    label: "Pool utilization",
    value: 0.6546,
    format: ValueFormat.Percent,
    trend: TrendDirection.Flat,
    trendLabel: "Below the rate kink",
  },
];

export const marketAssets: MarketAsset[] = [
  { symbol: AssetSymbol.Weth, role: AssetRole.Collateral, name: "Wrapped Ether" },
  { symbol: AssetSymbol.Usdc, role: AssetRole.Borrowable, name: "USD Coin" },
];

export const marketMetrics: MarketMetric[] = [
  {
    key: MarketMetricKey.SupplyApy,
    label: "Supply APY",
    value: 0.0482,
    format: ValueFormat.Percent,
    hint: "Earned by lenders, paid by borrowers",
  },
  {
    key: MarketMetricKey.BorrowApr,
    label: "Borrow APR",
    value: 0.0635,
    format: ValueFormat.Percent,
    hint: "Rises as the pool gets more utilized",
  },
  {
    key: MarketMetricKey.MaxLtv,
    label: "Max borrow",
    value: 0.75,
    format: ValueFormat.Percent,
    hint: "Of your collateral value",
  },
  {
    key: MarketMetricKey.LiquidationThreshold,
    label: "Liquidation at",
    value: 0.8,
    format: ValueFormat.Percent,
    hint: "Deliberate buffer above max borrow",
  },
];

export const availableLiquidity = 16_660_000;

export const utilizationModel: UtilizationModel = {
  current: 0.6546,
  kink: 0.8,
};

export const oracleReading: OracleReading = {
  symbol: AssetSymbol.Weth,
  price: 3412.58,
  updatedSecondsAgo: 27,
  stalenessThresholdSeconds: 3600,
  status: OracleStatus.Fresh,
};

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
