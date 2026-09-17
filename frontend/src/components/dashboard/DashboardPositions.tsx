"use client";

import { AppRoute, AssetSymbol, ButtonVariant, IconName, SectionId } from "@/lib/enums";
import { healthTier, toPriceScaled, toValueScaled } from "@/lib/health";
import { isNoDebtHealthFactor } from "@/lib/units";
import { dashboardContent } from "@/content/dashboard";
import { assetPrices, debtDecimals } from "@/content/protocol";
import { useAccountData } from "@/hooks/useAccountData";
import { useMarketData } from "@/hooks/useMarketData";
import { useOraclePrice } from "@/hooks/useOraclePrice";
import { useWalletState } from "@/hooks/useWalletState";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { EmptyState } from "@/components/ui/EmptyState";
import { Skeleton } from "@/components/ui/Skeleton";
import { HealthBar } from "@/components/borrow/HealthBar";
import { HealthScoreGauge } from "@/components/borrow/HealthScoreGauge";
import { LiquidationRiskWarning } from "@/components/borrow/LiquidationRiskWarning";
import { LenderPositionCard } from "@/components/lend/LenderPositionCard";
import { BorrowerPositionCard } from "@/components/dashboard/BorrowerPositionCard";
import { PositionOverviewHeader } from "@/components/dashboard/PositionOverviewHeader";
import { RecentActivityList } from "@/components/dashboard/RecentActivityList";
import { RiskLegend } from "@/components/dashboard/RiskLegend";

const fallbackDebtPriceScaled = toPriceScaled(assetPrices[AssetSymbol.Usdc]);

export function DashboardPositions() {
  const { address } = useWalletState();
  const account = useAccountData(address);
  const market = useMarketData();
  const debtPrice = useOraclePrice("debt");

  if (account.isError || market.isError) {
    return (
      <EmptyState
        title="We could not read your position"
        description="The blockchain node did not answer. Your funds are unaffected, and this clears once the connection recovers."
        icon={IconName.Warning}
        action={
          <Button variant={ButtonVariant.Subtle} onClick={() => account.refetch()}>
            Try again
          </Button>
        }
      />
    );
  }

  if (account.data === undefined || market.data === undefined) {
    return (
      <div className="flex flex-col gap-6">
        <Skeleton className="h-32 w-full" />
        <Skeleton className="h-64 w-full" />
      </div>
    );
  }

  const { data } = account;
  const hasNoPositions = data.supplyAssets <= 0n && data.collateralAmount <= 0n && data.debtAmount <= 0n;

  if (hasNoPositions) {
    return (
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
    );
  }

  const debtPriceScaled = debtPrice.data === undefined ? fallbackDebtPriceScaled : debtPrice.data.price;
  const suppliedValueScaled = toValueScaled(data.supplyAssets, debtDecimals, debtPriceScaled);
  const factorBps = data.debtAmount <= 0n || isNoDebtHealthFactor(data.healthFactorBps) ? null : data.healthFactorBps;
  const tier = healthTier(factorBps);

  return (
    <div className="flex flex-col gap-6">
      <PositionOverviewHeader
        suppliedValueScaled={suppliedValueScaled}
        collateralValueScaled={data.collateralValue}
        debtValueScaled={data.debtValue}
      />

      <LiquidationRiskWarning tier={tier} />

      <Card className="grid gap-8 p-6 sm:p-7 lg:grid-cols-[minmax(0,1fr)_minmax(0,18rem)]">
        <div className="flex flex-col gap-6">
          <HealthScoreGauge factorBps={factorBps} tier={tier} />
          <HealthBar
            factorBps={factorBps}
            tier={tier}
            maxLtvBps={market.data.maxLtvBps}
            liquidationThresholdBps={market.data.liquidationThresholdBps}
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
          <LenderPositionCard depositedBalance={data.supplyAssets} />
          <BorrowerPositionCard
            factorBps={factorBps}
            tier={tier}
            collateralValueScaled={data.collateralValue}
            collateralDeposited={data.collateralAmount}
            debtOutstanding={data.debtAmount}
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
  );
}
