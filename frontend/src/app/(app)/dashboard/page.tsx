import type { Metadata } from "next";
import { AppRoute, ButtonVariant, IconName, SectionId, SectionTone } from "@/lib/enums";
import { assetPrices } from "@/content/protocol";
import { depositedBalance, lendAssetDecimals } from "@/content/lend";
import {
  collateralDecimals,
  collateralDeposited,
  collateralUnitPriceScaled,
  debtDecimals,
  debtOutstanding,
  debtUnitPriceScaled,
  liquidationThresholdBps,
  maxLtvBps,
} from "@/content/borrow";
import { dashboardContent, recentLiquidation } from "@/content/dashboard";
import { healthFactorBps, healthTier, toPriceScaled, toValueScaled } from "@/lib/health";
import { AssetSymbol } from "@/lib/enums";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { EmptyState } from "@/components/ui/EmptyState";
import { PageHeader } from "@/components/ui/PageHeader";
import { Section } from "@/components/ui/Section";
import { WalletGate } from "@/components/app/WalletGate";
import { PriceStalenessWarning } from "@/components/markets/PriceStalenessWarning";
import { HealthBar } from "@/components/borrow/HealthBar";
import { HealthScoreGauge } from "@/components/borrow/HealthScoreGauge";
import { LiquidationRiskWarning } from "@/components/borrow/LiquidationRiskWarning";
import { LenderPositionCard } from "@/components/lend/LenderPositionCard";
import { BorrowerPositionCard } from "@/components/dashboard/BorrowerPositionCard";
import { LiquidationEventBanner } from "@/components/dashboard/LiquidationEventBanner";
import { PositionOverviewHeader } from "@/components/dashboard/PositionOverviewHeader";
import { PriceDropDrawer } from "@/components/dashboard/PriceDropDrawer";
import { QuickActionBar } from "@/components/dashboard/QuickActionBar";
import { RecentActivityList } from "@/components/dashboard/RecentActivityList";
import { RiskLegend } from "@/components/dashboard/RiskLegend";

export const metadata: Metadata = {
  title: "Dashboard",
  description: "Your deposits, loan, collateral, and safety score in a single view.",
};

const suppliedValueScaled = toValueScaled(
  depositedBalance,
  lendAssetDecimals,
  toPriceScaled(assetPrices[AssetSymbol.Usdc]),
);
const collateralValueScaled = toValueScaled(collateralDeposited, collateralDecimals, collateralUnitPriceScaled);
const debtValueScaled = toValueScaled(debtOutstanding, debtDecimals, debtUnitPriceScaled);
const factorBps = healthFactorBps(collateralValueScaled, debtValueScaled, liquidationThresholdBps);
const tier = healthTier(factorBps);

const hasNoPositions = depositedBalance <= 0n && collateralDeposited <= 0n && debtOutstanding <= 0n;

export default function DashboardPage() {
  return (
    <>
      <PageHeader title={dashboardContent.title} description={dashboardContent.description} aside={<PriceDropDrawer />}>
        <QuickActionBar />
      </PageHeader>

      <PriceStalenessWarning />

      <Section id={SectionId.DashboardOverview} tone={SectionTone.Canvas}>
        <h2 id={`${SectionId.DashboardOverview}-heading`} className="sr-only">
          {dashboardContent.overviewTitle}
        </h2>

        <WalletGate>
          {hasNoPositions ? (
            <EmptyState
              title={dashboardContent.emptyTitle}
              description={dashboardContent.emptyDescription}
              icon={IconName.Coins}
              action={
                <div className="flex flex-col gap-2 sm:flex-row">
                  <Button href={AppRoute.Lend} trailingIcon={IconName.ArrowRight}>
                    Start lending
                  </Button>
                  <Button href={AppRoute.Borrow} variant={ButtonVariant.Secondary}>
                    Borrow instead
                  </Button>
                </div>
              }
            />
          ) : (
            <div className="flex flex-col gap-6">
              <PositionOverviewHeader
                suppliedValueScaled={suppliedValueScaled}
                collateralValueScaled={collateralValueScaled}
                debtValueScaled={debtValueScaled}
              />

              {recentLiquidation === null ? null : <LiquidationEventBanner event={recentLiquidation} />}

              <LiquidationRiskWarning tier={tier} />

              <Card className="grid gap-8 p-6 sm:p-7 lg:grid-cols-[minmax(0,1fr)_minmax(0,18rem)]">
                <div className="flex flex-col gap-6">
                  <HealthScoreGauge factorBps={factorBps} tier={tier} />
                  <HealthBar
                    factorBps={factorBps}
                    tier={tier}
                    maxLtvBps={maxLtvBps}
                    liquidationThresholdBps={liquidationThresholdBps}
                  />
                </div>
                <div className="border-t border-line pt-6 lg:border-l lg:border-t-0 lg:pl-8 lg:pt-0">
                  <RiskLegend />
                </div>
              </Card>

              <div className="flex flex-col gap-4">
                <h3 id={`${SectionId.DashboardPositions}-heading`} className="text-lg font-semibold tracking-tight text-ink">
                  {dashboardContent.positionsTitle}
                </h3>
                <div className="grid gap-6 lg:grid-cols-2">
                  <LenderPositionCard />
                  <BorrowerPositionCard
                    factorBps={factorBps}
                    tier={tier}
                    collateralValueScaled={collateralValueScaled}
                  />
                </div>
              </div>

              <div className="flex flex-col gap-4">
                <div className="flex flex-col gap-1.5">
                  <h3 id={`${SectionId.DashboardActivity}-heading`} className="text-lg font-semibold tracking-tight text-ink">
                    {dashboardContent.activityTitle}
                  </h3>
                  <p className="text-sm leading-relaxed text-ink-soft">{dashboardContent.activityDescription}</p>
                </div>
                <RecentActivityList />
              </div>
            </div>
          )}
        </WalletGate>
      </Section>
    </>
  );
}
