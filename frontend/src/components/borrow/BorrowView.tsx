"use client";

import { AssetSymbol, BadgeTone, ButtonVariant, IconName, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { formatHealthFactor, scaledValueToUsd } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import { borrowPageContent, collateralDecimals, debtDecimals } from "@/content/borrow";
import { usePositionView } from "@/hooks/usePositionView";
import { Alert } from "@/components/ui/Alert";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { EmptyState } from "@/components/ui/EmptyState";
import { Skeleton } from "@/components/ui/Skeleton";
import { PriceStalenessWarning } from "@/components/markets/PriceStalenessWarning";
import { CollateralPanel } from "@/components/borrow/CollateralPanel";
import { DebtPanel } from "@/components/borrow/DebtPanel";
import { FullLiquidationNotice } from "@/components/borrow/FullLiquidationNotice";
import { HealthBadge } from "@/components/borrow/HealthBadge";
import { HealthBar } from "@/components/borrow/HealthBar";
import { HealthScoreGauge } from "@/components/borrow/HealthScoreGauge";
import { LiquidationRiskWarning } from "@/components/borrow/LiquidationRiskWarning";
import { PriceDropSimulator } from "@/components/borrow/PriceDropSimulator";
import { MetricCard } from "@/components/portal/MetricCard";
import { MetricGrid } from "@/components/portal/MetricGrid";
import { PortalSection } from "@/components/portal/PortalSection";

export function BorrowView() {
  const { view, isError, refetch } = usePositionView();

  if (isError) {
    return (
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
    );
  }

  if (view === undefined) {
    return (
      <div className="flex flex-col gap-4">
        <Skeleton className="h-24 w-full rounded-card" />
        <Skeleton className="h-96 w-full rounded-card" />
      </div>
    );
  }

  return (
    <>
      <PriceStalenessWarning />

      <MetricGrid>
        <MetricCard
          label="Collateral"
          value={formatValue(scaledValueToUsd(view.collateralValueScaled), ValueFormat.UsdPrice)}
          hint={`${formatTokenAmount(view.collateralDeposited, collateralDecimals, 4)} ${AssetSymbol.Weth} locked`}
        />
        <MetricCard
          label="Borrowed"
          value={`${formatTokenAmount(view.debtOutstanding, debtDecimals, 2)} ${AssetSymbol.Usdc}`}
          hint="Interest accrues every second"
        />
        <MetricCard
          label="Available to borrow"
          value={`${formatTokenAmount(view.maxBorrowable, debtDecimals, 2)} ${AssetSymbol.Usdc}`}
          hint="Limited by collateral and pool liquidity"
        />
        <MetricCard
          label="Health factor"
          value={view.isValued ? formatHealthFactor(view.factorBps) : "—"}
          hint={view.isValued ? "Liquidation below 1.00" : "Waiting for a usable price"}
          accessory={view.isValued ? <HealthBadge tier={view.tier} /> : null}
          emphasis
        />
      </MetricGrid>

      {view.isValued ? null : (
        <Alert title="We cannot value your position right now" tone={BadgeTone.Caution} icon={IconName.Warning}>
          The WETH price feed has not reported recently, so your safety score cannot be calculated. Your collateral and
          loan are untouched, and the protocol refuses to borrow against or release collateral at a price it cannot
          trust. Repaying still works.
        </Alert>
      )}

      {view.isValued ? <LiquidationRiskWarning tier={view.tier} /> : null}

      {view.isValued ? (
        <PortalSection title={borrowPageContent.healthTitle} description={borrowPageContent.healthDescription}>
          <Card className="flex flex-col gap-5 p-5">
            <HealthScoreGauge factorBps={view.factorBps} tier={view.tier} />
            <HealthBar
              factorBps={view.factorBps}
              tier={view.tier}
              maxLtvBps={view.maxLtvBps}
              liquidationThresholdBps={view.liquidationThresholdBps}
            />
          </Card>
        </PortalSection>
      ) : null}

      <div className="grid gap-6 lg:grid-cols-2">
        <PortalSection title={borrowPageContent.collateralTitle} description={borrowPageContent.collateralDescription}>
          <CollateralPanel view={view} />
        </PortalSection>

        <PortalSection title={borrowPageContent.debtTitle} description={borrowPageContent.debtDescription}>
          <DebtPanel view={view} />
        </PortalSection>
      </div>

      <PortalSection title={borrowPageContent.simulatorTitle} description={borrowPageContent.simulatorDescription}>
        <PriceDropSimulator />
      </PortalSection>

      <FullLiquidationNotice />
    </>
  );
}
