import { HowItWorksStepKey, IconName, TrustSignalKey, TrustSignalStatus, ValuePropKey } from "@/lib/enums";

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

export const heroContent = {
  eyebrow: "WETH / USDC market · FUSD in development",
  title: "Put your crypto to work, or borrow against it without selling.",
  description:
    "Lend stablecoins and earn interest that accrues every second. Or unlock liquidity from assets you already hold, with a safety score that tells you exactly where you stand at all times. A second way to borrow, minting our own dollar, is being built.",
  primaryCta: "Start lending",
  secondaryCta: "Borrow against collateral",
  custodyNote: "Non-custodial. No account and no password — connect a wallet and you are ready.",
} as const;

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
    key: ValuePropKey.Mint,
    icon: IconName.Coins,
    title: "Or mint dollars yourself",
    description:
      "FUSD is a dollar the protocol creates against your collateral rather than borrowing it from a pool. No pool to run dry, and no interest to pay. Still being built.",
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

export const practiceContent = {
  eyebrow: "No risk",
  title: "Try the whole thing with test funds first",
  description:
    "Practice mode runs the real interface on a test network with fake money. Lend, borrow, push a position to liquidation, and see what happens before a single real asset is involved.",
  cta: "Open practice mode",
} as const;

export type BorrowRoute = {
  key: string;
  name: string;
  status: string;
  isLive: boolean;
  summary: string;
  costLabel: string;
  terms: { label: string; value: string }[];
  suitsYou: string;
};

export const borrowOrMintContent = {
  eyebrow: "Two ways to borrow",
  title: "Borrow someone else's dollars, or mint your own.",
  description:
    "Both lock the same collateral and both liquidate you if it falls too far. What differs is where the dollars come from, what they cost, and how much you can take.",
  footnote:
    "The market terms are read from the deployed contracts. The FUSD terms are the published design and are not live yet.",
} as const;
