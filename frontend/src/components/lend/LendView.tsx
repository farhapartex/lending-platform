"use client";

import { AssetSymbol, ButtonVariant, IconName, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { formatTokenAmount } from "@/lib/token";
import { bpsToRatio } from "@/lib/units";
import { debtAmountToNumber } from "@/lib/market";
import { lendAssetDecimals, lendPageContent } from "@/content/lend";
import { usePositionView } from "@/hooks/usePositionView";
import { Button } from "@/components/ui/Button";
import { EmptyState } from "@/components/ui/EmptyState";
import { Skeleton } from "@/components/ui/Skeleton";
import { MarketUtilization } from "@/components/markets/MarketUtilization";
import { PriceStalenessWarning } from "@/components/markets/PriceStalenessWarning";
import { LendActionPanel } from "@/components/lend/LendActionPanel";
import { LenderPositionCard } from "@/components/lend/LenderPositionCard";
import { SupplyApyCard } from "@/components/lend/SupplyApyCard";
import { MetricCard } from "@/components/portal/MetricCard";
import { MetricGrid } from "@/components/portal/MetricGrid";
import { PortalSection } from "@/components/portal/PortalSection";

export function LendView() {
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
          label="Your balance"
          value={`${formatTokenAmount(view.suppliedBalance, lendAssetDecimals, 2)} ${AssetSymbol.Usdc}`}
          hint="Interest already included"
          emphasis
        />
        <MetricCard
          label="Supply APY"
          value={formatValue(bpsToRatio(view.supplyAprBps), ValueFormat.Percent)}
          hint="Moves with pool utilization"
        />
        <MetricCard
          label="Pool size"
          value={formatValue(debtAmountToNumber(view.totalSupplied), ValueFormat.UsdCompact)}
          hint="Deposited by every lender"
        />
        <MetricCard
          label="Available to withdraw"
          value={formatValue(debtAmountToNumber(view.availableLiquidity), ValueFormat.UsdCompact)}
          hint="Not currently lent out"
        />
      </MetricGrid>

      <div className="grid gap-6 lg:grid-cols-[minmax(0,1fr)_minmax(0,24rem)]">
        <div className="flex flex-col gap-6">
          <PortalSection title={lendPageContent.positionTitle} description={lendPageContent.positionDescription}>
            <LenderPositionCard depositedBalance={view.suppliedBalance} />
          </PortalSection>

          <PortalSection title={lendPageContent.marketTitle} description={lendPageContent.marketDescription}>
            <div className="grid gap-4 sm:grid-cols-2">
              <SupplyApyCard supplyAprBps={view.supplyAprBps} />
              <div className="rounded-card border border-line bg-surface p-5">
                <MarketUtilization />
              </div>
            </div>
          </PortalSection>
        </div>

        <div className="lg:sticky lg:top-24 lg:self-start">
          <PortalSection title={lendPageContent.actionTitle} description={lendPageContent.actionDescription}>
            <LendActionPanel view={view} />
          </PortalSection>
        </div>
      </div>
    </>
  );
}
