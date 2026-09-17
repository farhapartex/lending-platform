"use client";

import { ButtonVariant, IconName, SectionId, SectionTone, WalletGatePurpose } from "@/lib/enums";
import { lendPageContent } from "@/content/lend";
import { usePositionView } from "@/hooks/usePositionView";
import { Button } from "@/components/ui/Button";
import { EmptyState } from "@/components/ui/EmptyState";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";
import { Skeleton } from "@/components/ui/Skeleton";
import { WalletGate } from "@/components/app/WalletGate";
import { MarketUtilization } from "@/components/markets/MarketUtilization";
import { PriceStalenessWarning } from "@/components/markets/PriceStalenessWarning";
import { LendActionPanel } from "@/components/lend/LendActionPanel";
import { LendHeader } from "@/components/lend/LendHeader";
import { LenderPositionCard } from "@/components/lend/LenderPositionCard";
import { SupplyApyCard } from "@/components/lend/SupplyApyCard";

export function LendView() {
  const { view, isError, refetch } = usePositionView();

  if (isError) {
    return (
      <Section id={SectionId.LendAction} tone={SectionTone.Canvas}>
        <EmptyState
          title="We could not read this market"
          description="The blockchain node did not answer. Nothing has changed on chain, and this clears once the connection recovers."
          icon={IconName.Warning}
          action={
            <Button variant={ButtonVariant.Subtle} onClick={refetch}>
              Try again
            </Button>
          }
        />
      </Section>
    );
  }

  if (view === undefined) {
    return (
      <Section id={SectionId.LendAction} tone={SectionTone.Canvas}>
        <div className="flex flex-col gap-6">
          <Skeleton className="h-40 w-full" />
          <Skeleton className="h-72 w-full" />
        </div>
      </Section>
    );
  }

  return (
    <>
      <LendHeader supplyAprBps={view.supplyAprBps} totalSupplied={view.totalSupplied} />

      <PriceStalenessWarning />

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
              <WalletGate purpose={WalletGatePurpose.PersonalData}>
                <LenderPositionCard depositedBalance={view.suppliedBalance} />
              </WalletGate>
            </div>

            <div className="flex flex-col gap-4">
              <h2 id={`${SectionId.LendMarket}-heading`} className="text-lg font-semibold tracking-tight text-ink">
                {lendPageContent.marketTitle}
              </h2>
              <p className="max-w-2xl text-sm leading-relaxed text-ink-soft">{lendPageContent.marketDescription}</p>
              <div className="grid gap-5 sm:grid-cols-2">
                <SupplyApyCard supplyAprBps={view.supplyAprBps} />
                <div className="rounded-card border border-line bg-surface p-6">
                  <MarketUtilization />
                </div>
              </div>
            </div>
          </div>

          <div className="flex flex-col gap-4">
            <h2 id={`${SectionId.LendAction}-heading`} className="text-lg font-semibold tracking-tight text-ink">
              {lendPageContent.actionTitle}
            </h2>
            <p className="text-sm leading-relaxed text-ink-soft">{lendPageContent.actionDescription}</p>
            <LendActionPanel view={view} />
          </div>
        </div>
      </Section>
    </>
  );
}
