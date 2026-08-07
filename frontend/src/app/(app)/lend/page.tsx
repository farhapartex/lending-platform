import type { Metadata } from "next";
import { BadgeTone, DataStatus, IconName, OracleStatus, SectionId, SectionTone } from "@/lib/enums";
import { marketDataStatus, oracleReading } from "@/content/protocol";
import { lendPageContent } from "@/content/lend";
import { marketsPageContent } from "@/content/markets";
import { Alert } from "@/components/ui/Alert";
import { Container } from "@/components/ui/Container";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";
import { UtilizationBar } from "@/components/markets/UtilizationBar";
import { LendHeader } from "@/components/lend/LendHeader";
import { LenderPositionCard } from "@/components/lend/LenderPositionCard";
import { LendActionPanel } from "@/components/lend/LendActionPanel";
import { SupplyApyCard } from "@/components/lend/SupplyApyCard";

export const metadata: Metadata = {
  title: "Lend",
  description: "Deposit USDC to earn interest that accrues every second, and withdraw at any time.",
};

export default function LendPage() {
  if (marketDataStatus === DataStatus.Unavailable) {
    return (
      <Container className="py-16">
        <Alert title={marketsPageContent.unavailableTitle} tone={BadgeTone.Caution} icon={IconName.Warning}>
          {marketsPageContent.unavailableDescription}
        </Alert>
      </Container>
    );
  }

  return (
    <>
      <LendHeader />

      {oracleReading.status === OracleStatus.Stale ? (
        <Container className="pt-6">
          <Alert title="The price feed has not updated recently" tone={BadgeTone.Caution} icon={IconName.Warning}>
            Deposits and withdrawals are paused while the price is stale, so you do not pay gas on a transaction that
            would be rejected. This clears automatically once a fresh price arrives.
          </Alert>
        </Container>
      ) : null}

      <Section id={SectionId.LendAction} tone={SectionTone.Canvas}>
        <div className="grid gap-8 lg:grid-cols-[minmax(0,1fr)_minmax(0,26rem)] lg:gap-10">
          <div className="flex flex-col gap-8">
            <div className="flex flex-col gap-4">
              <SectionHeading
                sectionId={SectionId.LendPosition}
                eyebrow="Position"
                title={lendPageContent.positionTitle}
                description={lendPageContent.positionDescription}
              />
              <LenderPositionCard />
            </div>

            <div className="flex flex-col gap-4">
              <h2 id={`${SectionId.LendMarket}-heading`} className="text-lg font-semibold tracking-tight text-ink">
                {lendPageContent.marketTitle}
              </h2>
              <p className="max-w-2xl text-sm leading-relaxed text-ink-soft">{lendPageContent.marketDescription}</p>
              <div className="grid gap-5 sm:grid-cols-2">
                <SupplyApyCard />
                <div className="rounded-card border border-line bg-surface p-6">
                  <UtilizationBar />
                </div>
              </div>
            </div>
          </div>

          <div className="flex flex-col gap-4">
            <h2 id={`${SectionId.LendAction}-heading`} className="text-lg font-semibold tracking-tight text-ink">
              {lendPageContent.actionTitle}
            </h2>
            <p className="text-sm leading-relaxed text-ink-soft">{lendPageContent.actionDescription}</p>
            <LendActionPanel />
          </div>
        </div>
      </Section>
    </>
  );
}
