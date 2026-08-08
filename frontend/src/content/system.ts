import { AppRoute } from "@/lib/enums";

export type SystemDestination = {
  label: string;
  description: string;
  href: AppRoute;
};

export const notFoundContent = {
  title: "We cannot find that page",
  description:
    "The link may be out of date, or the address may have a typo in it. Nothing is wrong with your account or your funds.",
  primaryCta: "Go to markets",
  secondaryCta: "Read the docs",
} as const;

export const errorContent = {
  title: "Something went wrong on our side",
  fundsNote: "Your deposits, collateral and loans are held by the contracts and are completely unaffected by this error.",
  description:
    "This is a problem with the interface, not with your position. Trying again usually clears it, and refreshing the page is safe.",
  retryCta: "Try again",
  homeCta: "Go to markets",
  digestLabel: "Reference",
} as const;

export const systemDestinations: SystemDestination[] = [
  { label: "Markets", description: "Rates, limits and utilization", href: AppRoute.Markets },
  { label: "Lend", description: "Deposit and withdraw USDC", href: AppRoute.Lend },
  { label: "Borrow", description: "Collateral, loans and safety", href: AppRoute.Borrow },
  { label: "Learn", description: "Plain-language documentation", href: AppRoute.Learn },
];
