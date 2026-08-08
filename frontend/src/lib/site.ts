import { AppRoute } from "@/lib/enums";

export const siteName = "Lending Platform";

export const siteUrl = process.env.NEXT_PUBLIC_SITE_URL ?? "http://localhost:5173";

export const indexableRoutes: AppRoute[] = [
  AppRoute.Home,
  AppRoute.Markets,
  AppRoute.Lend,
  AppRoute.Borrow,
  AppRoute.Liquidations,
  AppRoute.Practice,
  AppRoute.Welcome,
  AppRoute.Learn,
  AppRoute.LearnHowItWorks,
  AppRoute.LearnHealthScore,
  AppRoute.LearnLiquidation,
  AppRoute.LearnFees,
  AppRoute.LearnFaq,
  AppRoute.LearnGlossary,
];

export const nonIndexableRoutes: AppRoute[] = [AppRoute.Dashboard, AppRoute.History];

export function absoluteUrl(route: AppRoute): string {
  return new URL(route, siteUrl).toString();
}
