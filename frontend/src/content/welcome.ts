import { BadgeTone, IconName, WelcomeStepKey } from "@/lib/enums";

export type WelcomeStep = {
  key: WelcomeStepKey;
  eyebrow: string;
  title: string;
  icon: IconName;
  paragraphs: string[];
  bullets?: string[];
  callout?: {
    tone: BadgeTone;
    title: string;
    body: string;
  };
};

export const welcomeSteps: WelcomeStep[] = [
  {
    key: WelcomeStepKey.Wallet,
    eyebrow: "First things first",
    title: "Your wallet is your account",
    icon: IconName.Wallet,
    paragraphs: [
      "There is no signup form here, and no password to choose. A crypto wallet is an app that holds your assets and proves who you are by signing transactions, so connecting it is the whole of signing in.",
      "This is worth understanding before you deposit anything, because it cuts both ways.",
    ],
    bullets: [
      "Nobody can freeze your funds or block your access, including us.",
      "Your assets move only when you approve a transaction yourself.",
      "If you lose your wallet's seed phrase, nobody can recover it for you.",
    ],
    callout: {
      tone: BadgeTone.Caution,
      title: "Back up your seed phrase first",
      body: "It is the only way back into your wallet. Write it down, keep it offline, and never share it with anyone who asks, including anyone claiming to be support.",
    },
  },
  {
    key: WelcomeStepKey.Lending,
    eyebrow: "Concept one of four",
    title: "Lending: earn on assets sitting idle",
    icon: IconName.Coins,
    paragraphs: [
      "Deposit USDC into the shared pool and borrowers pay you interest to use it. Your balance grows continuously rather than in monthly instalments.",
    ],
    bullets: [
      "No lock-up. Withdraw whenever the pool has liquid funds available.",
      "Interest compounds into your balance automatically, so there is nothing to claim.",
      "Your rate follows demand. More borrowing means you earn more.",
    ],
  },
  {
    key: WelcomeStepKey.Borrowing,
    eyebrow: "Concept two of four",
    title: "Borrowing: unlock cash without selling",
    icon: IconName.Wallet,
    paragraphs: [
      "Lock WETH as collateral and borrow USDC against it. You keep your WETH, and its future upside, instead of selling to raise funds.",
    ],
    bullets: [
      "You can borrow up to a published share of your collateral's value.",
      "There is no credit check, no paperwork, and no fixed repayment schedule.",
      "Repay any amount at any time, and add collateral whenever you want.",
    ],
  },
  {
    key: WelcomeStepKey.Health,
    eyebrow: "Concept three of four",
    title: "Health score: the number that matters",
    icon: IconName.Gauge,
    paragraphs: [
      "Every loan carries a health score comparing your collateral's value against what you owe. Above 1.00 you are safe. At or below 1.00 your position can be closed by anyone.",
      "It moves whenever the WETH price moves, so it can change without you doing anything at all.",
    ],
    bullets: [
      "The interface shows it live, in plain words, not just a number.",
      "You are warned well before liquidation becomes possible.",
      "Adding collateral or repaying part of the loan both push it back up.",
    ],
  },
  {
    key: WelcomeStepKey.Liquidation,
    eyebrow: "Concept four of four",
    title: "Liquidation: what happens if it goes wrong",
    icon: IconName.Warning,
    paragraphs: [
      "If your health score reaches 1.00, anyone may repay your loan and take your collateral plus a published bonus. This exists so lenders are never left short by a loan that has gone bad.",
    ],
    callout: {
      tone: BadgeTone.Caution,
      title: "In this phase the whole position is closed",
      body: "Liquidation currently repays your entire loan rather than just enough to make it safe, so a brief price dip can end the loan altogether. Partial liquidation comes in the next phase. Until then, borrowing well under your limit is the real protection.",
    },
  },
  {
    key: WelcomeStepKey.Ready,
    eyebrow: "That is everything",
    title: "Try it with fake money first",
    icon: IconName.Beaker,
    paragraphs: [
      "You now know as much as you need to start. The safest next step is practice mode, which runs this exact interface on a test network with worthless tokens.",
      "Deposit, borrow, push a position until it gets liquidated, and see what that feels like before any real value is involved.",
    ],
  },
];

export const welcomePageContent = {
  title: "Getting started",
  description: "Six short steps covering how this platform works and what can go wrong. You can leave at any point.",
  skipLabel: "Skip for now",
  backLabel: "Back",
  nextLabel: "Continue",
  practiceLabel: "Open practice mode",
  finishLabel: "Go to markets",
  progressLabel: "Onboarding progress",
} as const;
