import type { Metadata } from "next";
import { AppRoute, ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { marketsPageContent } from "@/content/markets";
import { Button } from "@/components/ui/Button";
import { MarketOverviewMetrics } from "@/components/markets/MarketOverviewMetrics";
import { MarketTerms } from "@/components/markets/MarketTerms";
import { RateExplainer } from "@/components/markets/RateExplainer";
import { FeeDisclosureSummary } from "@/components/markets/FeeDisclosureSummary";
import { PriceStalenessWarning } from "@/components/markets/PriceStalenessWarning";
import { PortalPage } from "@/components/portal/PortalPage";
import { PortalSection } from "@/components/portal/PortalSection";

export const metadata: Metadata = {
  title: "Markets",
  description:
    "Live rates, limits, utilization, and fees for the WETH / USDC lending market.",
};

export default function MarketsPage() {
  return (
    <PortalPage
      title={marketsPageContent.title}
      description={marketsPageContent.description}
      actions={
        <>
          <Button href={AppRoute.Lend} size={ButtonSize.Sm} trailingIcon={IconName.ArrowRight}>
            Lend
          </Button>
          <Button href={AppRoute.Borrow} size={ButtonSize.Sm} variant={ButtonVariant.Secondary}>
            Borrow
          </Button>
        </>
      }
    >
      <PriceStalenessWarning />

      <MarketOverviewMetrics />

      <PortalSection title={marketsPageContent.summaryTitle} description={marketsPageContent.summaryDescription}>
        <MarketTerms />
      </PortalSection>

      <PortalSection title={marketsPageContent.ratesTitle} description={marketsPageContent.ratesDescription}>
        <RateExplainer />
      </PortalSection>

      <PortalSection
        title={marketsPageContent.feesTitle}
        description={marketsPageContent.feesDescription}
        actions={
          <Button href={AppRoute.LearnFees} size={ButtonSize.Sm} variant={ButtonVariant.Subtle} trailingIcon={IconName.ArrowRight}>
            Full disclosure
          </Button>
        }
      >
        <FeeDisclosureSummary />
      </PortalSection>
    </PortalPage>
  );
}
