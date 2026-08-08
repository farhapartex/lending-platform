import { BadgeTone, IconName, WelcomeStepKey } from "@/lib/enums";

export type WelcomeStep = {
  key: WelcomeStepKey;
  shortLabel: string;
  eyebrow: string;
  title: string;
  takeaway: string;
  icon: IconName;
  paragraphs: string[];
  bullets?: string[];
  callout?: {
    tone: BadgeTone;
    icon: IconName;
    title: string;
    body: string;
  };
};

export const welcomeSteps: WelcomeStep[] = [
  {
    key: WelcomeStepKey.Overview,
    shortLabel: "Overview",
    eyebrow: "Welcome",
    title: "Two ways to put your crypto to work",
    takeaway: "Nothing moves until you choose it.",
    icon: IconName.Sparkles,
    paragraphs: [
      "You can do two things here. Earn interest on crypto you are already holding, or borrow cash against it without having to sell.",
      "Have a look around as long as you like. Nothing happens automatically, and every page works without connecting anything.",
    ],
  },
  {
    key: WelcomeStepKey.Wallet,
    shortLabel: "Your wallet",
    eyebrow: "Signing in",
    title: "There is no account to create",
    takeaway: "You are already set up.",
    icon: IconName.Lock,
    paragraphs: [
      "No form to fill in, no password to invent, and no waiting for approval. Connecting your wallet is the whole of signing in.",
      "Your assets stay in your own wallet throughout. They move only when you approve a transaction yourself.",
    ],
    bullets: [
      "No email address and no personal details required.",
      "No minimum balance and no account fees.",
      "You can disconnect at any moment and nothing changes.",
    ],
    callout: {
      tone: BadgeTone.Brand,
      icon: IconName.ShieldCheck,
      title: "One thing worth setting up properly",
      body: "Keep your wallet's recovery phrase written down somewhere safe and offline. It is how you get back in on a new device, and it is the one thing nobody can replace for you.",
    },
  },
  {
    key: WelcomeStepKey.Lending,
    shortLabel: "Lending",
    eyebrow: "Option one",
    title: "Earn on crypto you are already holding",
    takeaway: "Withdraw whenever you like.",
    icon: IconName.Coins,
    paragraphs: [
      "Deposit USDC and borrowers pay you interest to use it. Your balance grows continuously rather than in monthly instalments.",
    ],
    bullets: [
      "No lock-up period and no minimum term.",
      "Interest is added to your balance automatically, so there is nothing to claim.",
      "Take your money out whenever the pool has funds available.",
    ],
  },
  {
    key: WelcomeStepKey.Borrowing,
    shortLabel: "Borrowing",
    eyebrow: "Option two",
    title: "Borrow without selling what you own",
    takeaway: "Keep your WETH and its upside.",
    icon: IconName.Wallet,
    paragraphs: [
      "Lock WETH as collateral and borrow USDC against it. No credit check, no paperwork, and no fixed repayment date.",
    ],
    bullets: [
      "Borrow up to a published share of what your collateral is worth.",
      "Repay any amount, whenever suits you.",
      "Add more collateral at any time to give yourself more room.",
    ],
  },
  {
    key: WelcomeStepKey.Health,
    shortLabel: "Safety score",
    eyebrow: "Staying informed",
    title: "One number tells you how you are doing",
    takeaway: "Above 1.00 means you are fine.",
    icon: IconName.Gauge,
    paragraphs: [
      "Every loan has a safety score comparing your collateral against what you owe. We show it in plain words next to the number, and it updates live as prices move.",
    ],
    bullets: [
      "Safe, Caution and At risk are the three states you will see.",
      "You are warned early, well before anything can happen to your loan.",
      "There is a preview tool that shows what a price drop would do, before it happens.",
    ],
  },
  {
    key: WelcomeStepKey.Liquidation,
    shortLabel: "Staying safe",
    eyebrow: "Good habits",
    title: "How to stay on the safe side",
    takeaway: "Borrow less than the maximum and you have room to breathe.",
    icon: IconName.ShieldCheck,
    paragraphs: [
      "If a loan's safety score ever reaches 1.00, someone else is allowed to repay it and take the collateral. That rule is what protects the people whose money you borrowed.",
      "Avoiding it is straightforward, and the app is built to help you.",
    ],
    bullets: [
      "Borrow comfortably under your limit rather than right up to it.",
      "The borrow screen suggests a safer figure than the maximum it permits.",
      "Topping up collateral or repaying a little takes one transaction.",
    ],
    callout: {
      tone: BadgeTone.Neutral,
      icon: IconName.Info,
      title: "Worth knowing while we are early",
      body: "For now, a liquidation closes the whole loan rather than just enough to make it safe again. Partial liquidation arrives in a later update, so leaving yourself a decent buffer matters more in the meantime.",
    },
  },
  {
    key: WelcomeStepKey.Ready,
    shortLabel: "Try it",
    eyebrow: "You are ready",
    title: "Have a go with play money first",
    takeaway: "Practice mode uses worthless test tokens.",
    icon: IconName.Beaker,
    paragraphs: [
      "Practice mode is this exact app running on a test network. Deposit, borrow, and even watch a position get liquidated, all with tokens that are worth nothing.",
      "When it feels familiar, the real thing works in exactly the same way.",
    ],
  },
];

export const welcomePageContent = {
  title: "Welcome",
  description: "A quick tour of how this works. Around two minutes, and you can leave whenever you like.",
  skipLabel: "Skip the tour",
  backLabel: "Back",
  nextLabel: "Continue",
  practiceLabel: "Try practice mode",
  finishLabel: "Take me to the app",
  progressLabel: "Tour progress",
  takeawayLabel: "In short",
  reassurance: "Reading this changes nothing. No wallet is connected and no funds can move.",
} as const;
