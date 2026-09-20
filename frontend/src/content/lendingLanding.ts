import { IconName } from "@/lib/enums";

export type LendingPath = {
  key: string;
  icon: IconName;
  audience: string;
  title: string;
  description: string;
  steps: string[];
};

export const lendingLandingContent = {
  eyebrow: "Live now · WETH / USDC",
  title: "Put idle dollars to work, or unlock the ones inside your crypto.",
  description:
    "One shared pool. Lenders deposit USDC and earn what borrowers pay. Borrowers lock WETH and draw against it without selling. Nobody negotiates, and every rate comes from how much of the pool is in use.",
  note: "Rates and limits below are read from the deployed contracts, not written on this page.",
} as const;

export const lendingPaths: LendingPath[] = [
  {
    key: "lender",
    icon: IconName.Coins,
    audience: "If you hold stablecoins",
    title: "Deposit and earn",
    description:
      "Your USDC joins the pool and starts earning the moment it lands. Interest compounds into the balance itself, so there is nothing to claim and nothing to reinvest.",
    steps: [
      "Deposit any amount above the pool minimum",
      "Earn continuously from the people borrowing it",
      "Withdraw whenever, as long as the pool holds enough liquid USDC",
    ],
  },
  {
    key: "borrower",
    icon: IconName.Wallet,
    audience: "If you hold WETH",
    title: "Lock collateral and borrow",
    description:
      "Lock WETH, draw USDC against it, and keep the upside of the asset you did not sell. There is no fixed term and no repayment schedule.",
    steps: [
      "Lock WETH as collateral, which stays in your name",
      "Borrow up to 75% of what it is worth",
      "Repay any amount at any time to release it",
    ],
  },
];

export const lendingRiskContent = {
  eyebrow: "The part that catches people out",
  title: "A falling price can cost you your collateral.",
  description:
    "Borrowing is the easy half. The half worth reading is what happens when the collateral behind your loan loses value.",
  points: [
    {
      title: "You are liquidated at 80%, not 75%",
      body: "The gap between what you can borrow and where liquidation starts is deliberate breathing room. Borrow the maximum and you are one price move from losing it.",
    },
    {
      title: "Your score falls on its own",
      body: "Interest keeps adding to what you owe, so a loan you opened months ago and forgot about is riskier than you left it, even if the price never moved.",
    },
    {
      title: "A stranger closes the position, not us",
      body: "Whoever settles an unsafe loan repays it and takes your collateral plus a five percent bonus. Nobody is monitoring on your behalf.",
    },
  ],
} as const;
