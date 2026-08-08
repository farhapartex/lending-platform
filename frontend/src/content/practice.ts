import { AssetSymbol, FaucetStatus, IconName } from "@/lib/enums";
import { assetDecimals } from "@/content/protocol";
import { parseTokenAmount } from "@/lib/token";

const usdcDecimals = assetDecimals[AssetSymbol.Usdc];
const wethDecimals = assetDecimals[AssetSymbol.Weth];

export type FaucetAsset = {
  symbol: AssetSymbol;
  decimals: number;
  amount: bigint;
  status: FaucetStatus;
  cooldownHoursRemaining: number | null;
};

export type PracticeIdea = {
  id: string;
  icon: IconName;
  title: string;
  description: string;
};

export const faucetAssets: FaucetAsset[] = [
  {
    symbol: AssetSymbol.Usdc,
    decimals: usdcDecimals,
    amount: parseTokenAmount("5000", usdcDecimals) ?? 0n,
    status: FaucetStatus.Ready,
    cooldownHoursRemaining: null,
  },
  {
    symbol: AssetSymbol.Weth,
    decimals: wethDecimals,
    amount: parseTokenAmount("2", wethDecimals) ?? 0n,
    status: FaucetStatus.CoolingDown,
    cooldownHoursRemaining: 19,
  },
];

export const faucetRequestsPerDay = 1;

export const realAspects: string[] = [
  "The interface is the same one you would use with real money. Nothing is simplified or hidden.",
  "Rates, safety scores and liquidation rules behave exactly as they do on the live network.",
  "Your actions are genuine blockchain transactions, so you get a feel for confirmations and gas.",
];

export const unrealAspects: string[] = [
  "The tokens are worthless. They cannot be sold, swapped, or moved to the live network.",
  "Any interest you earn or losses you take here are not real either way.",
  "Test networks are occasionally reset, which clears balances without notice.",
];

export const practiceIdeas: PracticeIdea[] = [
  {
    id: "idea-deposit",
    icon: IconName.Coins,
    title: "Deposit and watch interest arrive",
    description: "Put some test USDC in and watch the balance tick upward second by second.",
  },
  {
    id: "idea-borrow",
    icon: IconName.Wallet,
    title: "Take out a loan",
    description: "Post WETH as collateral, borrow against it, and see how the safety score responds.",
  },
  {
    id: "idea-simulate",
    icon: IconName.Sliders,
    title: "Preview a price drop",
    description: "Use the simulator on the borrow page to see what a fall in the WETH price would do to you.",
  },
  {
    id: "idea-liquidate",
    icon: IconName.ShieldCheck,
    title: "Cause a liquidation on purpose",
    description:
      "Borrow close to the limit, let the score fall under 1.00, then liquidate it yourself from the liquidations page. Far better to see it here than for real.",
  },
];

export const practicePageContent = {
  title: "Practice mode",
  description:
    "The same app running on a test network, with tokens that are worth nothing. Try anything you like, including the things you would never risk with real money.",
  explainerTitle: "What is real and what is not",
  explainerDescription: "Worth being clear about, so nothing here misleads you about the live platform.",
  realTitle: "Real",
  unrealTitle: "Not real",
  setupTitle: "Getting set up",
  setupDescription: "Two things to do, and you are ready to experiment.",
  ideasTitle: "Things worth trying",
  ideasDescription: "The point of practice mode is to make mistakes cheaply. Here is where to start.",
  faucetTitle: "Get test funds",
  faucetDescription: "Free test tokens, sent to your connected wallet.",
  faucetPendingNote: "Faucet requests go live alongside the test network deployment.",
  returnTitle: "Ready for the real thing?",
  returnDescription:
    "Nothing carries across from practice mode, so you start fresh. Take your time, and consider borrowing well under your limit for your first loan.",
  returnCta: "Switch to the live app",
} as const;
