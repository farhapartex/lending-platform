import {
  AssetRole,
  AssetSymbol,
  DataStatus,
  FeeKind,
  HowItWorksStepKey,
  IconName,
  MarketMetricKey,
  ProtocolStatKey,
  TrendDirection,
  TrustSignalKey,
  TrustSignalStatus,
  ValueFormat,
  ValuePropKey,
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

export type ValueProp = {
  key: ValuePropKey;
  icon: IconName;
  title: string;
  description: string;
};

export type HowItWorksStep = {
  key: HowItWorksStepKey;
  icon: IconName;
  title: string;
  description: string;
};

export type TrustSignal = {
  key: TrustSignalKey;
  icon: IconName;
  title: string;
  description: string;
  status: TrustSignalStatus;
};

export type FeeItem = {
  kind: FeeKind;
  label: string;
  value: string;
  description: string;
};

export const heroContent = {
  eyebrow: "Phase 1 · WETH / USDC market",
  title: "Put your crypto to work, or borrow against it without selling.",
  description:
    "Lend stablecoins and earn interest that accrues every second. Or unlock liquidity from assets you already hold, with a safety score that tells you exactly where you stand at all times.",
  primaryCta: "Start lending",
  secondaryCta: "Borrow against collateral",
  custodyNote: "Non-custodial. No account and no password — connect a wallet and you are ready.",
} as const;

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

export const valueProps: ValueProp[] = [
  {
    key: ValuePropKey.Earn,
    icon: IconName.Coins,
    title: "Earn on idle stablecoins",
    description:
      "Deposit USDC and start earning immediately. Interest compounds automatically into your balance, with no claiming step and no lock-up.",
  },
  {
    key: ValuePropKey.Unlock,
    icon: IconName.Wallet,
    title: "Borrow without selling",
    description:
      "Use the assets you already hold as collateral to unlock liquidity. Keep your upside, skip the paperwork, and repay whenever you choose.",
  },
  {
    key: ValuePropKey.Visibility,
    icon: IconName.Gauge,
    title: "Risk you can actually read",
    description:
      "A plain-language safety score updates live with prices. See how a price drop would affect you before it happens, not after.",
  },
];

export const howItWorksSteps: HowItWorksStep[] = [
  {
    key: HowItWorksStepKey.Connect,
    icon: IconName.Wallet,
    title: "Connect a wallet",
    description:
      "No signup form and no password to forget. Your wallet is your account, and your funds never leave your control.",
  },
  {
    key: HowItWorksStepKey.Choose,
    icon: IconName.Coins,
    title: "Lend or borrow",
    description:
      "Deposit USDC to earn interest, or deposit WETH as collateral and borrow against it up to a safe limit we show you upfront.",
  },
  {
    key: HowItWorksStepKey.Monitor,
    icon: IconName.Gauge,
    title: "Stay ahead of risk",
    description:
      "Watch your safety score live, get warned well before liquidation is possible, and add collateral or repay in a single step.",
  },
];

export const trustSignals: TrustSignal[] = [
  {
    key: TrustSignalKey.SelfCustody,
    icon: IconName.Lock,
    title: "You keep custody",
    description: "No deposit is ever held by us, and no admin function can move or access your balances.",
    status: TrustSignalStatus.Live,
  },
  {
    key: TrustSignalKey.PublishedParameters,
    icon: IconName.Sliders,
    title: "Published risk parameters",
    description: "Every limit, threshold, and bonus percentage is documented before you commit any funds.",
    status: TrustSignalStatus.Live,
  },
  {
    key: TrustSignalKey.OpenSource,
    icon: IconName.Code,
    title: "Open source contracts",
    description: "The lending, collateral, and liquidation logic is readable and verifiable by anyone.",
    status: TrustSignalStatus.Live,
  },
  {
    key: TrustSignalKey.Audit,
    icon: IconName.ShieldCheck,
    title: "Independent audit",
    description: "Scheduled ahead of mainnet launch. The full report will be published here when complete.",
    status: TrustSignalStatus.Planned,
  },
];

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

export const practiceContent = {
  eyebrow: "No risk",
  title: "Try the whole thing with test funds first",
  description:
    "Practice mode runs the real interface on a test network with fake money. Lend, borrow, push a position to liquidation, and see what happens before a single real asset is involved.",
  cta: "Open practice mode",
} as const;
