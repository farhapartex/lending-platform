import { AppRoute, IconName, ProductKey, TrustSignalKey, TrustSignalStatus } from "@/lib/enums";

export type TrustSignal = {
  key: TrustSignalKey;
  icon: IconName;
  title: string;
  description: string;
  status: TrustSignalStatus;
};

export type ProductSummary = {
  key: ProductKey;
  eyebrow: string;
  title: string;
  description: string;
  points: string[];
  href: AppRoute;
  cta: string;
  isLive: boolean;
};

export type Mechanic = {
  key: string;
  icon: IconName;
  title: string;
  description: string;
};

export const heroContent = {
  eyebrow: "Non-custodial · WETH collateral",
  title: "Two ways to get dollars without selling your crypto.",
  description:
    "Lock what you already hold and either borrow USDC that other people have deposited, or mint FUSD, a dollar the protocol issues against your collateral. Both keep your assets in your name and both publish every rule before you commit anything.",
  note: "No account, no password, no approval. Connect a wallet and the rules are the same for everyone.",
} as const;

export const products: ProductSummary[] = [
  {
    key: ProductKey.Lending,
    eyebrow: "Live now",
    title: "Lend and borrow",
    description:
      "A shared pool of USDC. Deposit to earn interest paid by borrowers, or lock WETH and borrow against it at a rate the market sets.",
    points: [
      "Earn on idle stablecoins with no lock-up",
      "Borrow up to 75% of what your collateral is worth",
      "Interest accrues every second and compounds on its own",
    ],
    href: AppRoute.Lending,
    cta: "How lending works",
    isLive: true,
  },
  {
    key: ProductKey.Fusd,
    eyebrow: "In development",
    title: "Mint FUSD",
    description:
      "A dollar created against your collateral rather than borrowed from anybody. No pool to run dry, and version one charges no interest at all.",
    points: [
      "Mint up to two thirds of your collateral value",
      "Backed by more than a dollar of crypto, always",
      "Plain ERC20 with no freeze switch and no transfer hooks",
    ],
    href: AppRoute.Fusd,
    cta: "How FUSD works",
    isLive: false,
  },
];

export const mechanicsContent = {
  eyebrow: "Shared mechanics",
  title: "Different products, the same three rules.",
  description:
    "Whichever side you use, the protocol protects lenders the same way, and it is worth understanding before you lock anything up.",
} as const;

export const mechanics: Mechanic[] = [
  {
    key: "overcollateralised",
    icon: IconName.Lock,
    title: "You always lock more than you take",
    description:
      "That surplus is what replaces a credit check. It is also what absorbs a fall in price before anybody else is exposed to it.",
  },
  {
    key: "health",
    icon: IconName.Gauge,
    title: "One number tells you where you stand",
    description:
      "A health factor, updated live with the price. Above the line you are fine, and the interface warns you long before you approach it.",
  },
  {
    key: "liquidation",
    icon: IconName.ShieldCheck,
    title: "Anyone can close an unsafe position",
    description:
      "If a position falls through the line, a stranger repays the debt and takes the collateral plus a published bonus. Harsh, and it is what keeps every deposit backed.",
  },
];

export const trustContent = {
  eyebrow: "Where you stand",
  title: "What we can and cannot do with your money.",
} as const;

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
    description: "Every limit, threshold and bonus is documented before you commit any funds.",
    status: TrustSignalStatus.Live,
  },
  {
    key: TrustSignalKey.OpenSource,
    icon: IconName.Code,
    title: "Open source contracts",
    description: "The lending, collateral and liquidation logic is readable and verifiable by anyone.",
    status: TrustSignalStatus.Live,
  },
  {
    key: TrustSignalKey.Audit,
    icon: IconName.ShieldCheck,
    title: "Independent audit",
    description: "Not done. Nobody has reviewed these contracts, which is the single biggest reason not to use real money yet.",
    status: TrustSignalStatus.Planned,
  },
];

export const closingContent = {
  title: "Ready to look around?",
  description:
    "Connecting a wallet takes one click and creates nothing. You can read every market figure before you decide to move anything.",
  cta: "Connect a wallet",
} as const;
