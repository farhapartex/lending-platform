import type { Metadata } from "next";
import { BadgeTone, DataStatus, IconName, SectionId, SectionTone } from "@/lib/enums";
import { marketDataStatus } from "@/content/protocol";
import { marketsPageContent } from "@/content/markets";
import {
  borrowPageContent,
  collateralDecimals,
  collateralDeposited,
  collateralUnitPriceScaled,
  debtDecimals,
  debtOutstanding,
  debtUnitPriceScaled,
  liquidationThresholdBps,
  maxLtvBps,
} from "@/content/borrow";
import { healthFactorBps, healthTier, toValueScaled } from "@/lib/health";
import { Alert } from "@/components/ui/Alert";
import { Card } from "@/components/ui/Card";
import { Container } from "@/components/ui/Container";
import { Section } from "@/components/ui/Section";
import { PriceStalenessWarning } from "@/components/markets/PriceStalenessWarning";
import { BorrowHeader } from "@/components/borrow/BorrowHeader";
import { CollateralPanel } from "@/components/borrow/CollateralPanel";
import { DebtPanel } from "@/components/borrow/DebtPanel";
import { FullLiquidationNotice } from "@/components/borrow/FullLiquidationNotice";
import { HealthBar } from "@/components/borrow/HealthBar";
import { HealthScoreGauge } from "@/components/borrow/HealthScoreGauge";
import { LiquidationRiskWarning } from "@/components/borrow/LiquidationRiskWarning";
import { PriceDropSimulator } from "@/components/borrow/PriceDropSimulator";

export const metadata: Metadata = {
  title: "Borrow",
  description:
    "Borrow USDC against WETH collateral, with a live safety score, a price-drop simulator, and warnings well before liquidation.",
};

const collateralValueScaled = toValueScaled(collateralDeposited, collateralDecimals, collateralUnitPriceScaled);
const debtValueScaled = toValueScaled(debtOutstanding, debtDecimals, debtUnitPriceScaled);
const factorBps = healthFactorBps(collateralValueScaled, debtValueScaled, liquidationThresholdBps);
const tier = healthTier(factorBps);

export default function BorrowPage() {
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
      <BorrowHeader tier={tier} collateralValueScaled={collateralValueScaled} />

      <PriceStalenessWarning />

      <Section id={SectionId.BorrowHealth} tone={SectionTone.Canvas}>
        <div className="flex flex-col gap-6">
          <h2 id={`${SectionId.BorrowHealth}-heading`} className="sr-only">
            {borrowPageContent.healthTitle}
          </h2>

          <Card className="flex flex-col gap-6 p-6 sm:p-7">
            <HealthScoreGauge factorBps={factorBps} tier={tier} />
            <HealthBar
              factorBps={factorBps}
              tier={tier}
              maxLtvBps={maxLtvBps}
              liquidationThresholdBps={liquidationThresholdBps}
            />
          </Card>

          <LiquidationRiskWarning tier={tier} />

          <div className="grid gap-6 lg:grid-cols-2">
            <div className="flex flex-col gap-4">
              <div className="flex flex-col gap-1.5">
                <h3 id={`${SectionId.BorrowCollateral}-heading`} className="text-lg font-semibold tracking-tight text-ink">
                  {borrowPageContent.collateralTitle}
                </h3>
                <p className="text-sm leading-relaxed text-ink-soft">{borrowPageContent.collateralDescription}</p>
              </div>
              <CollateralPanel />
            </div>

            <div className="flex flex-col gap-4">
              <div className="flex flex-col gap-1.5">
                <h3 id={`${SectionId.BorrowDebt}-heading`} className="text-lg font-semibold tracking-tight text-ink">
                  {borrowPageContent.debtTitle}
                </h3>
                <p className="text-sm leading-relaxed text-ink-soft">{borrowPageContent.debtDescription}</p>
              </div>
              <DebtPanel />
            </div>
          </div>

          <PriceDropSimulator />

          <FullLiquidationNotice />
        </div>
      </Section>
    </>
  );
}
