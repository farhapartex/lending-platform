import type { Metadata } from "next";
import { BadgeTone, DataStatus, IconName } from "@/lib/enums";
import { marketDataStatus } from "@/content/protocol";
import { marketsPageContent } from "@/content/markets";
import { Alert } from "@/components/ui/Alert";
import { Container } from "@/components/ui/Container";
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
  if (marketDataStatus === DataStatus.Unavailable) {
    return (
      <Container className="py-16">
        <Alert title={marketsPageContent.unavailableTitle} tone={BadgeTone.Caution} icon={IconName.ShieldCheck}>
          {marketsPageContent.unavailableDescription}
        </Alert>
      </Container>
    );
  }

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
