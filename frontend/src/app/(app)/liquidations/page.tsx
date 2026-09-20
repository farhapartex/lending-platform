import type { Metadata } from "next";
import { liquidationsPageContent } from "@/content/liquidations";
import { HowLiquidationWorksCallout } from "@/components/liquidations/HowLiquidationWorksCallout";
import { LiquidationHistoryPanel } from "@/components/liquidations/LiquidationHistoryPanel";
import { LiquidationsPanel } from "@/components/liquidations/LiquidationsPanel";
import { PortalPage } from "@/components/portal/PortalPage";
import { PortalSection } from "@/components/portal/PortalSection";

export const metadata: Metadata = {
  title: "Liquidations",
  description:
    "Positions eligible for liquidation right now, with the published bonus for resolving each one. Open to anyone, not just bots.",
};

export default function LiquidationsPage() {
  return (
    <PortalPage title={liquidationsPageContent.title} description={liquidationsPageContent.description}>
      <PortalSection title={liquidationsPageContent.listTitle}>
        <LiquidationsPanel />
      </PortalSection>

      <HowLiquidationWorksCallout />

      <PortalSection
        title={liquidationsPageContent.historyTitle}
        description={liquidationsPageContent.historyDescription}
      >
        <LiquidationHistoryPanel />
      </PortalSection>

      <p className="text-xs leading-relaxed text-ink-faint">
        These lists are rebuilt from indexed contract events. Eligibility is always re-checked on-chain when a
        liquidation runs, so a position that recovers in the meantime cannot be liquidated.
      </p>
    </PortalPage>
  );
}
