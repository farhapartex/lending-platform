import type { Metadata } from "next";
import { MarketHeader } from "@/components/markets/MarketHeader";
import { PriceStalenessWarning } from "@/components/markets/PriceStalenessWarning";
import { ProtocolStatsGrid } from "@/components/markets/ProtocolStatsGrid";
import { MarketCard } from "@/components/markets/MarketCard";
import { RateExplainer } from "@/components/markets/RateExplainer";
import { FeeDisclosureSummary } from "@/components/markets/FeeDisclosureSummary";

export const metadata: Metadata = {
  title: "Markets",
  description:
    "Live rates, limits, utilization, and fees for the WETH / USDC lending market. Visible without connecting a wallet.",
};

export default function MarketsPage() {
  return (
    <>
      <MarketHeader />

      <PriceStalenessWarning />

      <ProtocolStatsGrid />
      <MarketCard />
      <RateExplainer />
      <FeeDisclosureSummary />
    </>
  );
}
