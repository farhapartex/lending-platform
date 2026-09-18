import { IconName } from "@/lib/enums";

export type FusdTerm = {
  label: string;
  value: string;
  hint: string;
};

export type FusdStep = {
  key: string;
  title: string;
  description: string;
};

export type FusdUse = {
  key: string;
  icon: IconName;
  title: string;
  description: string;
};

export const fusdPageContent = {
  eyebrow: "In development",
  title: "A dollar you mint yourself, backed by crypto you already own.",
  description:
    "FUSD is a token meant to hold a steady value of one US dollar. You create it by locking up collateral worth more than the amount you mint. No company holds reserves in a bank, and no address can freeze your balance.",
  statusNote:
    "The token contract is written and tested. The engine that holds collateral and enforces the rules is being built, so none of this is live yet.",
} as const;

export const whatItIsContent = {
  eyebrow: "The idea",
  title: "There are three ways to hold a token at a dollar. This is the second.",
  description:
    "Most tokens move a lot. That is fine for trading and painful for paying someone. A stablecoin is meant to stay still, and how it manages that is the whole story.",
} as const;

export const stablecoinKinds = [
  {
    key: "bank",
    title: "Backed by money in a bank",
    description:
      "A company takes your dollar, keeps it in an account, and issues you a token. Simple, and it holds the peg well. You have to trust the company, and it can freeze your balance.",
    verdict: "Not this",
  },
  {
    key: "crypto",
    title: "Backed by crypto in a contract",
    description:
      "You lock up something worth more than you take out, and a contract issues the token against it. Nobody holds your money. The trade-off is that you must overcollateralize, because the backing can fall in price.",
    verdict: "FUSD is this",
  },
  {
    key: "algorithmic",
    title: "Backed by an algorithm",
    description:
      "The protocol tries to hold the peg by minting and burning a second token when the price drifts. This has failed hard and publicly. We are not going near it.",
    verdict: "Not this",
  },
] as const;

export const howItWorksContent = {
  eyebrow: "The loop",
  title: "Lock collateral, mint dollars, repay when you want them back.",
  description:
    "Say you hold 1 ETH at 3,000 dollars and you want spending money without selling. The whole cycle is four steps, and you control the timing of all of them.",
} as const;

export const fusdSteps: FusdStep[] = [
  {
    key: "deposit",
    title: "Lock your collateral",
    description:
      "Send WETH or WBTC into the engine. It stays yours. The protocol prices it through a Chainlink feed and refuses to act on a price that has gone quiet.",
  },
  {
    key: "mint",
    title: "Mint FUSD against it",
    description:
      "Against 3,000 dollars of ETH you can mint up to 2,000 FUSD. The engine recalculates your health factor with the new debt included and only mints if you stay safe. Otherwise nothing happens at all.",
  },
  {
    key: "spend",
    title: "Use the dollars",
    description:
      "FUSD is a plain ERC20 with no transfer hooks, no blocklist and no fee on transfer. Every wallet and every exchange handles it without special treatment.",
  },
  {
    key: "repay",
    title: "Return them and unlock",
    description:
      "Send FUSD back, the engine burns it, and your debt falls. Clear the whole debt and you can take every bit of your collateral back out.",
  },
];

export const healthContent = {
  eyebrow: "The one number",
  title: "Everything comes back to the health factor.",
  description:
    "It tells you how safe your position is. Above 1 you are fine. Exactly 1 is the edge. Below 1 anyone in the world can close your position.",
  formula: "health factor = (collateral value in USD × 0.6667) ÷ FUSD minted",
  exampleTitle: "Worked example",
  exampleBody:
    "Deposit 1 ETH at 3,000 dollars and mint 1,000 FUSD. That gives (3000 × 0.6667) ÷ 1000, a health factor of 2.0. You could mint up to 2,000 FUSD before reaching 1.0, but sitting on the line is a bad idea — one price tick puts you under.",
  drift:
    "Unlike a loan with interest, this number does not drift down on its own. Version 1 charges nothing to borrow, so only the collateral price moves it.",
} as const;

export const fusdTerms: FusdTerm[] = [
  { label: "Collateral ratio", value: "150%", hint: "You need 150 dollars of collateral for every 100 FUSD minted" },
  { label: "Liquidation threshold", value: "6667 bps", hint: "The same rule, written as the share of collateral that counts against your debt" },
  { label: "Minimum health factor", value: "1.00", hint: "Below this, anyone can close your position" },
  { label: "Liquidator bonus", value: "8%", hint: "Paid to whoever closes an unsafe position" },
  { label: "Protocol fee", value: "2%", hint: "Taken on liquidation and kept in the reserve" },
  { label: "Oracle timeout", value: "3 hours", hint: "Older than this and the protocol treats the price as unusable" },
  { label: "Cost to borrow", value: "None", hint: "Version 1 charges no interest and no stability fee" },
];

export const usesContent = {
  eyebrow: "What it is for",
  title: "Five reasons somebody would want this.",
} as const;

export const fusdUses: FusdUse[] = [
  {
    key: "liquidity",
    icon: IconName.Wallet,
    title: "Liquidity without selling",
    description:
      "The main use. You hold ETH you do not want to sell, for tax reasons or conviction. Lock it, mint dollars, repay whenever. No credit check and no fixed term.",
  },
  {
    key: "leverage",
    icon: IconName.Gauge,
    title: "Leverage, carefully",
    description:
      "Mint, buy more collateral, deposit that too. Each loop increases your exposure and how fast a falling price reaches you. For people who understand exactly what they are doing.",
  },
  {
    key: "dollars",
    icon: IconName.Coins,
    title: "Dollars that settle in seconds",
    description:
      "Hold it, send it, pay with it. Useful for settling invoices across borders or sitting out a volatile stretch without going back to a bank.",
  },
  {
    key: "liquidating",
    icon: IconName.ShieldCheck,
    title: "Earning by liquidating",
    description:
      "Watch for positions below a health factor of 1 and close them. You put up FUSD and receive collateral worth 8 percent more. The protocol needs people doing this.",
  },
  {
    key: "building",
    icon: IconName.Receipt,
    title: "Something to build on",
    description:
      "No transfer hooks, no rebasing, no fee on transfer. It drops into pools, payment flows and vaults without special handling, and every useful piece of state has a view function.",
  },
];

export const pegContent = {
  eyebrow: "Why it holds",
  title: "Two forces pull it back to a dollar.",
  above:
    "If FUSD trades above a dollar, minting becomes profitable. People mint and sell, supply rises, and the price comes down.",
  below:
    "If FUSD trades below a dollar, anyone holding debt can buy it cheaply and repay at full value. They buy, supply shrinks, and the price goes back up.",
  floor:
    "Underneath both, every FUSD is always claimable against more than a dollar of collateral. That sets a floor on what the market believes it is worth.",
} as const;

export const limitsContent = {
  eyebrow: "The limits",
  title: "What you cannot do.",
  items: [
    "Mint FUSD without posting collateral first.",
    "Take out more than two thirds of your collateral value.",
    "Avoid liquidation by being early or well connected. The rule is identical for everyone.",
    "Use collateral the protocol has not listed. Each accepted token needs a price feed configured first.",
  ],
} as const;

export const statusContent = {
  eyebrow: "Where it stands",
  title: "Honest about what exists.",
  done: [
    "The FUSD token itself, with unit tests covering minting, burning and access control",
    "An oracle library that refuses a price older than three hours",
    "The engine's storage layout, risk constants and access rules",
  ],
  next: [
    "Deposit, mint, burn and redeem, plus the combined convenience calls",
    "The health factor and USD conversion maths",
    "Liquidation, with the check that a position ends healthier than it started",
    "A full test suite including handler-based invariants",
  ],
  after: [
    "Internal security review, then an external audit",
    "A testnet deployment with a public faucet",
    "These screens, reading real balances instead of explaining a plan",
  ],
} as const;
