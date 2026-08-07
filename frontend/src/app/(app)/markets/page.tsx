import type { Metadata } from "next";
import { BadgeTone, DataStatus, IconName, OracleStatus } from "@/lib/enums";
import { marketDataStatus, oracleReading } from "@/content/protocol";
import { marketsPageContent } from "@/content/markets";
import { Alert } from "@/components/ui/Alert";
import { Container } from "@/components/ui/Container";
import { MarketHeader } from "@/components/markets/MarketHeader";
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

      {oracleReading.status === OracleStatus.Stale ? (
        <Container className="pt-6">
          <Alert title="The price feed has not updated recently" tone={BadgeTone.Caution} icon={IconName.ShieldCheck}>
            Deposits, borrows, and liquidations are paused while the price is stale. This protects you from acting on an
            out-of-date number, and it clears automatically once a fresh price arrives.
          </Alert>
        </Container>
      ) : null}

      <ProtocolStatsGrid />
      <MarketCard />
      <RateExplainer />
      <FeeDisclosureSummary />
    </>
  );
}
