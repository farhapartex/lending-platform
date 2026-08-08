import type { Metadata } from "next";
import { BadgeTone, DataStatus, IconName, SectionId, SectionTone } from "@/lib/enums";
import { liquidationsDataStatus, liquidationsPageContent } from "@/content/liquidations";
import { Alert } from "@/components/ui/Alert";
import { Container } from "@/components/ui/Container";
import { PageHeader } from "@/components/ui/PageHeader";
import { Section } from "@/components/ui/Section";
import { PriceStalenessWarning } from "@/components/markets/PriceStalenessWarning";
import { HowLiquidationWorksCallout } from "@/components/liquidations/HowLiquidationWorksCallout";
import { LiquidationsPanel } from "@/components/liquidations/LiquidationsPanel";

export const metadata: Metadata = {
  title: "Liquidations",
  description:
    "Positions eligible for liquidation right now, with the published bonus for resolving each one. Open to anyone, not just bots.",
};

export default function LiquidationsPage() {
  return (
    <>
      <PageHeader title={liquidationsPageContent.title} description={liquidationsPageContent.description} />

      <PriceStalenessWarning />

      <Section id={SectionId.LiquidationsList} tone={SectionTone.Canvas}>
        <h2 id={`${SectionId.LiquidationsList}-heading`} className="sr-only">
          {liquidationsPageContent.listTitle}
        </h2>

        {liquidationsDataStatus === DataStatus.Unavailable ? (
          <Alert title={liquidationsPageContent.unavailableTitle} tone={BadgeTone.Caution} icon={IconName.Warning}>
            {liquidationsPageContent.unavailableDescription}
          </Alert>
        ) : (
          <div className="grid gap-8 lg:grid-cols-[minmax(0,1fr)_minmax(0,20rem)] lg:gap-10">
            <LiquidationsPanel />
            <HowLiquidationWorksCallout />
          </div>
        )}
      </Section>

      <Container className="pb-4">
        <p className="text-xs text-ink-faint">
          This list is built from indexed contract events. Eligibility is always re-checked on-chain when a liquidation
          runs.
        </p>
      </Container>
    </>
  );
}
