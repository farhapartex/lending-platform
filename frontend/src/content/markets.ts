import { IconName, RateExplainerPointKey } from "@/lib/enums";

export type RateExplainerPoint = {
  key: RateExplainerPointKey;
  icon: IconName;
  title: string;
  description: string;
};

export const marketsPageContent = {
  title: "WETH / USDC market",
  description:
    "One market, one set of published rules. Deposit USDC to earn interest, or post WETH as collateral to borrow against it.",
  totalsTitle: "Protocol totals",
  totalsDescription: "Live figures for the whole platform, visible without connecting a wallet.",
  summaryTitle: "Market terms",
  summaryDescription:
    "The rates you pay or earn, and the two limits that decide when a position becomes unsafe. Both limits are shown because they are deliberately different.",
  ratesTitle: "Why rates move",
  ratesDescription:
    "Nobody sets these rates by hand. They follow pool utilization, which is simply how much of the deposited money is currently borrowed.",
  feesTitle: "Fees",
  feesDescription: "The complete list, visible before you connect anything.",
  unavailableTitle: "Market data is unavailable",
  unavailableDescription:
    "We could not reach the network to read live market data. Nothing is wrong with your funds, and this page will recover on its own once the connection returns.",
} as const;

export const rateExplainerPoints: RateExplainerPoint[] = [
  {
    key: RateExplainerPointKey.Demand,
    icon: IconName.TrendUp,
    title: "More borrowing pushes rates up",
    description:
      "When a larger share of the pool is borrowed, borrowers pay more and lenders earn more. Rates fall again as borrowing eases off.",
  },
  {
    key: RateExplainerPointKey.Kink,
    icon: IconName.Sliders,
    title: "Past the kink, rates climb sharply",
    description:
      "Above a set utilization point the curve steepens hard. That is deliberate, and it discourages the pool from ever being drained.",
  },
  {
    key: RateExplainerPointKey.Withdrawals,
    icon: IconName.Lock,
    title: "It protects your ability to withdraw",
    description:
      "A fully borrowed pool would mean lenders cannot take their money out. The steep curve keeps liquidity available for withdrawals.",
  },
];
