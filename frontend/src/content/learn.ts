import { AppRoute, BadgeTone, DocBlockKind, DocKey } from "@/lib/enums";

export type DocBlock =
  | { kind: DocBlockKind.Prose; paragraphs: string[] }
  | { kind: DocBlockKind.Bullets; items: string[] }
  | { kind: DocBlockKind.Callout; tone: BadgeTone; title: string; body: string }
  | { kind: DocBlockKind.FeeTable }
  | { kind: DocBlockKind.BonusExample }
  | { kind: DocBlockKind.Faq }
  | { kind: DocBlockKind.Glossary }
  | { kind: DocBlockKind.HealthTiers };

export type DocSection = {
  id: string;
  title: string;
  blocks: DocBlock[];
};

export type DocPage = {
  key: DocKey;
  route: AppRoute;
  title: string;
  summary: string;
  sections: DocSection[];
};

export type FaqItem = {
  id: string;
  question: string;
  answer: string;
};

export type GlossaryEntry = {
  term: string;
  definition: string;
};

export const faqItems: FaqItem[] = [
  {
    id: "faq-account",
    question: "Do I need an account?",
    answer:
      "No. There is no signup form, no email, and no password. Your wallet is your identity, and connecting it is the whole of signing in. Your first deposit creates your position directly on the blockchain.",
  },
  {
    id: "faq-lost-wallet",
    question: "What if I lose my wallet or seed phrase?",
    answer:
      "We cannot recover it, and neither can anyone else. That is the direct consequence of us never holding your funds. Back up your seed phrase somewhere safe and offline, because it is the only way back into your wallet.",
  },
  {
    id: "faq-recover-funds",
    question: "Can you recover or freeze my funds?",
    answer:
      "No. We never take custody of your assets, and no administrative function in the contracts can move or access your balances. That protects you from us as much as it protects you from anyone else, but it does mean mistakes cannot be undone.",
  },
  {
    id: "faq-withdraw",
    question: "Can I withdraw whenever I want?",
    answer:
      "Yes, as long as the pool holds enough liquid funds at that moment. Some of what lenders deposit is lent out to borrowers, so during periods of very high borrowing a large withdrawal may need to wait until borrowers repay. The interface always shows the amount you can withdraw right now.",
  },
  {
    id: "faq-minimum",
    question: "Is there a minimum deposit?",
    answer:
      "There is a small minimum, which exists to stop tiny amounts from clogging the pool with positions worth less than the gas needed to manage them. The deposit form states the exact figure.",
  },
  {
    id: "faq-rate-change",
    question: "Why did my interest rate change?",
    answer:
      "Rates are not set by hand. They follow how much of the pool is currently borrowed. When borrowing rises, borrowers pay more and lenders earn more; when it falls, both come back down. Nobody can give one user a better rate than another.",
  },
  {
    id: "faq-warning",
    question: "Will I be warned before I am liquidated?",
    answer:
      "Yes. Your safety score is visible at all times and the interface warns you well before liquidation is possible, with the two actions that fix it: add collateral or repay part of the loan. In this first phase those warnings appear in the app itself, not by email.",
  },
  {
    id: "faq-liquidation-amount",
    question: "If I am liquidated, do I lose everything?",
    answer:
      "In this first phase liquidation closes the entire position rather than only the unsafe part, so even a brief price dip can end the whole loan. Partial liquidation, which takes only what is needed to restore safety, is planned for the next phase. Keeping a buffer well above the limit is the best protection until then.",
  },
  {
    id: "faq-assets",
    question: "Which assets are supported?",
    answer:
      "This first phase supports one market: WETH as collateral and USDC for lending and borrowing. Additional assets arrive in a later phase, deliberately after the core lending and liquidation loop has been proven with a small surface area.",
  },
  {
    id: "faq-audit",
    question: "Has the code been audited?",
    answer:
      "An independent audit is scheduled before the mainnet launch and the full report will be published here when it is complete. Until then, treat this as unaudited software and do not commit funds you cannot afford to lose.",
  },
];

export const glossaryEntries: GlossaryEntry[] = [
  {
    term: "APY",
    definition:
      "Annual percentage yield. What a lender earns over a year if the current rate held steady, including the effect of interest compounding into the balance.",
  },
  {
    term: "APR",
    definition: "Annual percentage rate. What a borrower pays over a year at the current rate, before compounding.",
  },
  {
    term: "Collateral",
    definition:
      "An asset you deposit and lock as security for a loan. You keep ownership of it unless your position becomes unsafe and is liquidated.",
  },
  {
    term: "Max borrow, or collateral factor",
    definition:
      "The largest share of your collateral's value you are allowed to borrow. Set below the liquidation threshold on purpose, so borrowing the maximum does not immediately put you at risk.",
  },
  {
    term: "Liquidation threshold",
    definition:
      "The share of your collateral's value your debt may reach before your position can be liquidated. Always higher than the max borrow limit, and the gap between them is your buffer.",
  },
  {
    term: "Health factor",
    definition:
      "A single number describing how safe your position is. Above 1.00 means safe; at or below 1.00 means anyone can liquidate you. It moves whenever the collateral price moves.",
  },
  {
    term: "Liquidation",
    definition:
      "The process that resolves an unsafe loan. Someone repays the borrower's debt and receives the borrower's collateral plus a bonus, which keeps the pool solvent so lenders are not left short.",
  },
  {
    term: "Liquidation bonus",
    definition:
      "The published extra amount of collateral a liquidator receives on top of the debt they repay. It is the incentive for anyone to do the work of keeping the pool safe.",
  },
  {
    term: "Utilization",
    definition:
      "The share of deposited funds currently lent out. It is the single input that decides interest rates, and it also determines how much lenders can withdraw at any moment.",
  },
  {
    term: "Rate kink",
    definition:
      "The utilization point above which interest rates start climbing much more steeply. It exists to discourage the pool from ever being fully drained, which would leave lenders unable to withdraw.",
  },
  {
    term: "Oracle",
    definition:
      "The external price source the protocol reads to value collateral. Every risk calculation depends on it, which is why a price that has not updated recently causes actions to be rejected rather than guessed at.",
  },
  {
    term: "Stale price",
    definition:
      "A price that has not been refreshed recently enough to be trusted. Rather than acting on an old number, the protocol rejects the action until a fresh price arrives.",
  },
  {
    term: "Non-custodial",
    definition:
      "The platform never holds or controls your assets. They move only when you sign a transaction, and no administrator has a path to your balances.",
  },
  {
    term: "Seed phrase",
    definition:
      "The list of words that controls your wallet. Anyone who has it can take your funds, and losing it means losing access permanently. Nobody, including us, can reset it.",
  },
  {
    term: "Gas",
    definition:
      "The fee the blockchain network charges to process a transaction. It goes to the network, not to this platform, and you pay it even for a transaction that ends up failing.",
  },
  {
    term: "Dust",
    definition:
      "An amount so small that managing it costs more in gas than it is worth. Minimum deposit sizes exist to keep dust positions out of the pool.",
  },
  {
    term: "Bad debt",
    definition:
      "What remains when a position falls so far that seizing all its collateral still does not cover the loan. A reserve fund to absorb this is planned for the next phase; in this phase there is none.",
  },
];

export const docPages: DocPage[] = [
  {
    key: DocKey.HowItWorks,
    route: AppRoute.LearnHowItWorks,
    title: "How it works",
    summary: "The whole platform in plain terms: who puts money in, who takes it out, and what keeps it safe.",
    sections: [
      {
        id: "two-sides",
        title: "Two sides of one pool",
        blocks: [
          {
            kind: DocBlockKind.Prose,
            paragraphs: [
              "There is a single shared pool of USDC. Lenders put USDC in and earn interest on it. Borrowers lock up WETH as collateral and take USDC out as a loan, paying interest for the privilege.",
              "That interest is the entire mechanism. It flows from borrowers to lenders, and the platform keeps a small published share of it. There is no trading, no fund manager, and no counterparty deciding your terms.",
            ],
          },
        ],
      },
      {
        id: "lending",
        title: "Lending, step by step",
        blocks: [
          {
            kind: DocBlockKind.Bullets,
            items: [
              "Connect a wallet and deposit USDC. There is no lock-up period.",
              "Interest begins accruing immediately and is added straight to your balance.",
              "There is nothing to claim and no reinvest button, because compounding is automatic.",
              "Withdraw whenever you like, as long as the pool holds enough liquid funds at that moment.",
            ],
          },
          {
            kind: DocBlockKind.Callout,
            tone: BadgeTone.Neutral,
            title: "Your risk as a lender",
            body: "You are exposed to borrowers being liquidated too slowly in a sharp crash, which could leave the pool short. In this phase there is no reserve fund covering that gap, so it is a real risk rather than a theoretical one.",
          },
        ],
      },
      {
        id: "borrowing",
        title: "Borrowing, step by step",
        blocks: [
          {
            kind: DocBlockKind.Bullets,
            items: [
              "Deposit WETH as collateral. Collateral on its own carries no risk until you borrow against it.",
              "The interface shows the maximum you may borrow, and a smaller amount it recommends instead.",
              "Borrow USDC up to that limit. You keep your WETH and its upside.",
              "Repay any amount at any time, or add more collateral, to push your safety score back up.",
            ],
          },
        ],
      },
      {
        id: "no-signup",
        title: "Why there is no signup",
        blocks: [
          {
            kind: DocBlockKind.Prose,
            paragraphs: [
              "Your wallet already proves who you are. Signing a transaction demonstrates that you control an address, so there is nothing for a password to add and no account for us to create.",
              "This has a hard edge worth understanding before you deposit. Because we hold no credential and no custody, we cannot reset anything, reverse anything, or recover a lost wallet. The protection and the limitation are the same fact.",
            ],
          },
        ],
      },
      {
        id: "never-does",
        title: "What the protocol cannot do",
        blocks: [
          {
            kind: DocBlockKind.Bullets,
            items: [
              "It cannot move your deposits or your collateral. No administrative function touches user balances.",
              "It cannot give one user better terms than another. The same published parameters apply to everybody.",
              "It cannot change your rate by hand. Rates follow utilization, mechanically.",
              "It cannot act on a price it does not trust. A stale price causes rejection, not a guess.",
            ],
          },
        ],
      },
    ],
  },
  {
    key: DocKey.HealthScore,
    route: AppRoute.LearnHealthScore,
    title: "Your health score",
    summary: "The one number that decides whether your loan is safe, and what to do when it falls.",
    sections: [
      {
        id: "what-it-means",
        title: "What the number means",
        blocks: [
          {
            kind: DocBlockKind.Prose,
            paragraphs: [
              "Your health factor compares the value of your collateral against what you owe. Above 1.00 your position is safe. At or below 1.00 anyone is allowed to liquidate it.",
              "It is calculated fresh every time it is shown, never stored, because the collateral price moves constantly. A number that looked comfortable this morning can be tight this afternoon without you doing anything at all.",
            ],
          },
        ],
      },
      {
        id: "two-limits",
        title: "Two limits, not one",
        blocks: [
          {
            kind: DocBlockKind.Prose,
            paragraphs: [
              "There are two separate percentages, and confusing them is the most common way people get caught out.",
              "The first is the maximum you may borrow against your collateral. The second, always higher, is the point at which your debt becomes large enough relative to your collateral that liquidation is allowed. The gap between them is deliberate breathing room.",
            ],
          },
          {
            kind: DocBlockKind.Callout,
            tone: BadgeTone.Caution,
            title: "Borrowing the maximum is not the same as being safe",
            body: "If you borrow every last unit you are permitted, you start with almost no buffer. A small fall in the collateral price is then enough to put you at risk. This is why the interface recommends a lower figure than the one it allows.",
          },
        ],
      },
      {
        id: "levels",
        title: "The safety levels",
        blocks: [{ kind: DocBlockKind.HealthTiers }],
      },
      {
        id: "improve",
        title: "How to improve it",
        blocks: [
          {
            kind: DocBlockKind.Bullets,
            items: [
              "Add more collateral. This raises the value backing your loan without repaying anything.",
              "Repay part of your loan. This reduces what you owe, which lifts the score directly.",
              "Do either one early. Both take a single transaction, and both work far better before the price has already moved against you.",
            ],
          },
        ],
      },
    ],
  },
  {
    key: DocKey.Liquidation,
    route: AppRoute.LearnLiquidation,
    title: "How liquidation works",
    summary: "Why it exists, exactly when it can happen, and a worked example with real numbers.",
    sections: [
      {
        id: "why",
        title: "Why liquidation exists",
        blocks: [
          {
            kind: DocBlockKind.Prose,
            paragraphs: [
              "Lenders are owed their money back. If a borrower's collateral fell below the value of their loan and nothing happened, the shortfall would land on lenders who did nothing wrong.",
              "Liquidation prevents that by resolving a risky position before it becomes a loss. Someone repays the loan on the borrower's behalf and receives the collateral plus a bonus. The pool stays whole, and the person who did the work is paid for it.",
            ],
          },
        ],
      },
      {
        id: "when",
        title: "When a position becomes eligible",
        blocks: [
          {
            kind: DocBlockKind.Prose,
            paragraphs: [
              "A position becomes eligible the moment its health factor reaches 1.00, and not before. There is no discretion, no queue, and no manual review.",
              "Eligibility is checked again on-chain at the instant a liquidation executes. If the price recovers between someone spotting your position and their transaction running, the liquidation is rejected and nothing moves.",
            ],
          },
        ],
      },
      {
        id: "example",
        title: "A worked example",
        blocks: [{ kind: DocBlockKind.BonusExample }],
      },
      {
        id: "full-position",
        title: "This phase closes the whole position",
        blocks: [
          {
            kind: DocBlockKind.Callout,
            tone: BadgeTone.Caution,
            title: "Partial liquidation is not available yet",
            body: "In this first phase a liquidation repays the entire loan and takes the collateral backing it, rather than only the amount needed to restore safety. That means a brief dip below the threshold can end the whole loan. Partial liquidation is planned for the next phase, and until it ships the practical protection is to keep a wide buffer.",
          },
        ],
      },
      {
        id: "anyone",
        title: "Anyone can be a liquidator",
        blocks: [
          {
            kind: DocBlockKind.Prose,
            paragraphs: [
              "Liquidation is not reserved for professional operators running bots. The list of eligible positions is public, the reward is published rather than negotiated, and the platform provides a plain interface for doing it in a few clicks.",
              "You need enough USDC to repay the loan you are settling, and you pay network gas whether or not you win the race against other liquidators.",
            ],
          },
        ],
      },
    ],
  },
  {
    key: DocKey.Fees,
    route: AppRoute.LearnFees,
    title: "Fees",
    summary: "Every charge on the platform, in one place, with nothing behind a signup.",
    sections: [
      {
        id: "complete-list",
        title: "The complete list",
        blocks: [
          {
            kind: DocBlockKind.Prose,
            paragraphs: [
              "This is all of it. If a charge is not on this list, it does not exist.",
            ],
          },
          { kind: DocBlockKind.FeeTable },
        ],
      },
      {
        id: "never",
        title: "What is never charged",
        blocks: [
          {
            kind: DocBlockKind.Bullets,
            items: [
              "No fee to deposit, withdraw, borrow, or repay.",
              "No management fee and no performance fee on your deposit.",
              "No penalty for repaying early, and no minimum term.",
              "No charge taken from your principal. The interest spread comes out of borrower interest only.",
            ],
          },
        ],
      },
      {
        id: "gas",
        title: "Network gas is separate",
        blocks: [
          {
            kind: DocBlockKind.Prose,
            paragraphs: [
              "Every transaction costs gas, paid to the blockchain network rather than to us. We do not mark it up and we do not receive any part of it.",
              "Gas is charged even when a transaction fails, which is why the interface tries to warn you in advance when an action would be rejected.",
            ],
          },
        ],
      },
    ],
  },
  {
    key: DocKey.Faq,
    route: AppRoute.LearnFaq,
    title: "Frequently asked questions",
    summary: "Straight answers, including to the awkward questions about custody and recovery.",
    sections: [
      {
        id: "questions",
        title: "Questions",
        blocks: [{ kind: DocBlockKind.Faq }],
      },
    ],
  },
  {
    key: DocKey.Glossary,
    route: AppRoute.LearnGlossary,
    title: "Glossary",
    summary: "Every term the interface uses, defined without jargon.",
    sections: [
      {
        id: "terms",
        title: "Terms",
        blocks: [{ kind: DocBlockKind.Glossary }],
      },
    ],
  },
];

export const learnIndexContent = {
  title: "Learn",
  description:
    "Plain-language documentation for everything this platform does. No jargon, no marketing, and nothing hidden behind a signup.",
  exampleNote: "The numbers below are calculated from the live risk parameters, so they cannot drift out of date.",
} as const;

export function findDocPage(key: DocKey): DocPage {
  const page = docPages.find((candidate) => candidate.key === key);

  if (page === undefined) {
    throw new Error(`Unknown documentation page: ${key}`);
  }

  return page;
}
