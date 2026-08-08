export enum AppRoute {
  Home = "/",
  Markets = "/markets",
  Lend = "/lend",
  Borrow = "/borrow",
  Dashboard = "/dashboard",
  Liquidations = "/liquidations",
  History = "/history",
  Practice = "/practice",
  Welcome = "/welcome",
  Learn = "/learn",
  LearnHowItWorks = "/learn/how-it-works",
  LearnHealthScore = "/learn/health-score",
  LearnLiquidation = "/learn/liquidation",
  LearnFees = "/learn/fees",
  LearnFaq = "/learn/faq",
  LearnGlossary = "/learn/glossary",
}

export enum SectionId {
  Hero = "hero",
  ProtocolStats = "protocol-stats",
  ValueProps = "why-this-platform",
  HowItWorks = "how-it-works",
  Fees = "fees",
  Trust = "trust",
  Practice = "practice",
  MainContent = "main-content",
  ProtocolTotals = "protocol-totals",
  MarketSummary = "market-summary",
  MarketRates = "market-rates",
  MarketFees = "market-fees",
  LendPosition = "your-position",
  LendAction = "deposit-and-withdraw",
  LendMarket = "market-conditions",
  BorrowHealth = "position-health",
  BorrowCollateral = "your-collateral",
  BorrowDebt = "your-loan",
  BorrowSimulator = "price-simulator",
  DashboardOverview = "portfolio-overview",
  DashboardPositions = "your-positions",
  DashboardActivity = "recent-activity",
  LiquidationsList = "eligible-positions",
  HistoryList = "transaction-history",
}

export enum WelcomeStepKey {
  Overview = "overview",
  Wallet = "wallet",
  Lending = "lending",
  Borrowing = "borrowing",
  Health = "health",
  Liquidation = "liquidation",
  Ready = "ready",
}

export enum DocKey {
  HowItWorks = "howItWorks",
  HealthScore = "healthScore",
  Liquidation = "liquidation",
  Fees = "fees",
  Faq = "faq",
  Glossary = "glossary",
}

export enum DocBlockKind {
  Prose = "prose",
  Bullets = "bullets",
  Callout = "callout",
  FeeTable = "feeTable",
  BonusExample = "bonusExample",
  Faq = "faq",
  Glossary = "glossary",
  HealthTiers = "healthTiers",
}

export enum DateRangePreset {
  AllTime = "allTime",
  Last7Days = "last7Days",
  Last30Days = "last30Days",
  Last90Days = "last90Days",
}

export const allTypesFilter = "all";

export type TypeFilterValue = ActivityKind | typeof allTypesFilter;

export enum LiquidationSortKey {
  Health = "health",
  Size = "size",
  Reward = "reward",
}

export enum ActivityKind {
  Deposit = "deposit",
  Withdraw = "withdraw",
  Borrow = "borrow",
  Repay = "repay",
  CollateralAdded = "collateralAdded",
  CollateralWithdrawn = "collateralWithdrawn",
  Liquidation = "liquidation",
}

export enum CollateralTab {
  Deposit = "depositCollateral",
  Withdraw = "withdrawCollateral",
}

export enum DebtTab {
  Borrow = "borrow",
  Repay = "repay",
}

export enum RepayMode {
  Partial = "partial",
  Full = "full",
}

export enum HealthTier {
  Safe = "safe",
  Caution = "caution",
  AtRisk = "atRisk",
  Liquidatable = "liquidatable",
}

export enum LendTab {
  Deposit = "deposit",
  Withdraw = "withdraw",
}

export enum WalletStatus {
  Disconnected = "disconnected",
  Connecting = "connecting",
  Connected = "connected",
  WrongNetwork = "wrongNetwork",
}

export enum TxFlowStatus {
  Idle = "idle",
  AwaitingApproval = "awaitingApproval",
  AwaitingSignature = "awaitingSignature",
  Pending = "pending",
  Confirmed = "confirmed",
  Reverted = "reverted",
}

export enum StepState {
  Upcoming = "upcoming",
  Active = "active",
  Done = "done",
}

export enum AmountValidationCode {
  None = "none",
  Empty = "empty",
  InvalidAmount = "invalidAmount",
  BelowMinimum = "belowMinimum",
  ExceedsWalletBalance = "exceedsWalletBalance",
  ExceedsDeposit = "exceedsDeposit",
  ExceedsAvailableLiquidity = "exceedsAvailableLiquidity",
  ExceedsCollateral = "exceedsCollateral",
  ExceedsSafeWithdrawal = "exceedsSafeWithdrawal",
  ExceedsBorrowLimit = "exceedsBorrowLimit",
  ExceedsDebt = "exceedsDebt",
}

export enum ButtonVariant {
  Primary = "primary",
  Secondary = "secondary",
  Subtle = "subtle",
  Ghost = "ghost",
}

export enum ButtonSize {
  Sm = "sm",
  Md = "md",
  Lg = "lg",
}

export enum BadgeTone {
  Neutral = "neutral",
  Brand = "brand",
  Positive = "positive",
  Caution = "caution",
  Critical = "critical",
}

export enum SectionTone {
  Canvas = "canvas",
  Surface = "surface",
  Muted = "muted",
}

export enum SectionSpacing {
  Compact = "compact",
  Regular = "regular",
  Spacious = "spacious",
}

export enum SurfaceElevation {
  Flat = "flat",
  Raised = "raised",
  Lifted = "lifted",
}

export enum TextAlign {
  Left = "left",
  Center = "center",
}

export enum TrendDirection {
  Up = "up",
  Down = "down",
  Flat = "flat",
}

export enum DataStatus {
  Loading = "loading",
  Ready = "ready",
  Unavailable = "unavailable",
}

export enum ValueFormat {
  UsdCompact = "usdCompact",
  Usd = "usd",
  UsdPrice = "usdPrice",
  Percent = "percent",
}

export enum OracleStatus {
  Fresh = "fresh",
  Stale = "stale",
  Unavailable = "unavailable",
}

export enum NetworkKind {
  Mainnet = "mainnet",
  Testnet = "testnet",
}

export enum UtilizationZone {
  BelowKink = "belowKink",
  AboveKink = "aboveKink",
}

export enum RateExplainerPointKey {
  Demand = "demand",
  Kink = "kink",
  Withdrawals = "withdrawals",
}

export enum AppNavLinkKey {
  Markets = "markets",
  Lend = "lend",
  Borrow = "borrow",
  Dashboard = "dashboard",
  History = "history",
  Liquidations = "liquidations",
  Learn = "learn",
}

export enum ProtocolStatKey {
  TotalDeposited = "totalDeposited",
  TotalBorrowed = "totalBorrowed",
  Utilization = "utilization",
}

export enum MarketMetricKey {
  SupplyApy = "supplyApy",
  BorrowApr = "borrowApr",
  MaxLtv = "maxLtv",
  LiquidationThreshold = "liquidationThreshold",
  AvailableLiquidity = "availableLiquidity",
}

export enum AssetSymbol {
  Weth = "WETH",
  Usdc = "USDC",
}

export enum AssetRole {
  Collateral = "collateral",
  Borrowable = "borrowable",
}

export enum ValuePropKey {
  Earn = "earn",
  Unlock = "unlock",
  Visibility = "visibility",
}

export enum HowItWorksStepKey {
  Connect = "connect",
  Choose = "choose",
  Monitor = "monitor",
}

export enum TrustSignalKey {
  SelfCustody = "selfCustody",
  Audit = "audit",
  OpenSource = "openSource",
  PublishedParameters = "publishedParameters",
}

export enum TrustSignalStatus {
  Live = "live",
  Planned = "planned",
}

export enum FeeKind {
  InterestSpread = "interestSpread",
  LiquidationBonus = "liquidationBonus",
  PlatformFee = "platformFee",
}

export enum NavLinkKey {
  HowItWorks = "howItWorks",
  Fees = "fees",
  Trust = "trust",
  Learn = "learn",
}

export enum NavLinkKind {
  Route = "route",
  Anchor = "anchor",
}

export enum FooterGroupKey {
  Product = "product",
  Learn = "learn",
  Protocol = "protocol",
}

export enum IconName {
  ArrowRight = "arrowRight",
  ArrowUpRight = "arrowUpRight",
  Wallet = "wallet",
  Coins = "coins",
  ShieldCheck = "shieldCheck",
  Gauge = "gauge",
  TrendUp = "trendUp",
  TrendDown = "trendDown",
  Minus = "minus",
  Check = "check",
  Code = "code",
  Lock = "lock",
  Sliders = "sliders",
  Beaker = "beaker",
  Menu = "menu",
  Close = "close",
  ExternalLink = "externalLink",
  Receipt = "receipt",
  Info = "info",
  Warning = "warning",
  Loader = "loader",
  ChevronDown = "chevronDown",
  Sparkles = "sparkles",
}
