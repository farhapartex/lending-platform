"use client";

import { ButtonVariant, IconName, SectionId, SectionTone } from "@/lib/enums";
import { borrowPageContent } from "@/content/borrow";
import { usePositionView } from "@/hooks/usePositionView";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { EmptyState } from "@/components/ui/EmptyState";
import { Section } from "@/components/ui/Section";
import { Skeleton } from "@/components/ui/Skeleton";
import { BorrowHeader } from "@/components/borrow/BorrowHeader";
import { CollateralPanel } from "@/components/borrow/CollateralPanel";
import { DebtPanel } from "@/components/borrow/DebtPanel";
import { FullLiquidationNotice } from "@/components/borrow/FullLiquidationNotice";
import { HealthBar } from "@/components/borrow/HealthBar";
import { HealthScoreGauge } from "@/components/borrow/HealthScoreGauge";
import { LiquidationRiskWarning } from "@/components/borrow/LiquidationRiskWarning";
import { PriceDropSimulator } from "@/components/borrow/PriceDropSimulator";
import { PriceStalenessWarning } from "@/components/markets/PriceStalenessWarning";

export function BorrowView() {
  const { view, isError, refetch } = usePositionView();

  if (isError) {
    return (
      <Section id={SectionId.BorrowHealth} tone={SectionTone.Canvas}>
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
      <Section id={SectionId.BorrowHealth} tone={SectionTone.Canvas}>
        <div className="flex flex-col gap-6">
          <Skeleton className="h-40 w-full" />
          <Skeleton className="h-72 w-full" />
        </div>
      </Section>
    );
  }

  return (
    <>
      <BorrowHeader
        tier={view.tier}
        collateralValueScaled={view.collateralValueScaled}
        borrowAprBps={view.borrowAprBps}
      />

      <PriceStalenessWarning />

      <Section id={SectionId.BorrowHealth} tone={SectionTone.Canvas}>
        <div className="flex flex-col gap-6">
          <h2 id={`${SectionId.BorrowHealth}-heading`} className="sr-only">
            {borrowPageContent.healthTitle}
          </h2>

          <Card className="flex flex-col gap-6 p-6 sm:p-7">
            <HealthScoreGauge factorBps={view.factorBps} tier={view.tier} />
            <HealthBar
              factorBps={view.factorBps}
              tier={view.tier}
              maxLtvBps={view.maxLtvBps}
              liquidationThresholdBps={view.liquidationThresholdBps}
            />
          </Card>

          <LiquidationRiskWarning tier={view.tier} />

          <div className="grid gap-6 lg:grid-cols-2">
            <div className="flex flex-col gap-4">
              <div className="flex flex-col gap-1.5">
                <h3 id={`${SectionId.BorrowCollateral}-heading`} className="text-lg font-semibold tracking-tight text-ink">
                  {borrowPageContent.collateralTitle}
                </h3>
                <p className="text-sm leading-relaxed text-ink-soft">{borrowPageContent.collateralDescription}</p>
              </div>
              <CollateralPanel view={view} />
            </div>

            <div className="flex flex-col gap-4">
              <div className="flex flex-col gap-1.5">
                <h3 id={`${SectionId.BorrowDebt}-heading`} className="text-lg font-semibold tracking-tight text-ink">
                  {borrowPageContent.debtTitle}
                </h3>
                <p className="text-sm leading-relaxed text-ink-soft">{borrowPageContent.debtDescription}</p>
              </div>
              <DebtPanel view={view} />
            </div>
          </div>

          <PriceDropSimulator />

          <FullLiquidationNotice />
        </div>
      </Section>
    </>
  );
}
