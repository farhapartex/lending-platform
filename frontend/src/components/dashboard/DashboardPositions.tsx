"use client";

import { AppRoute, AssetSymbol, BadgeTone, ButtonVariant, IconName, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { healthTier, scaledValueToUsd, toPriceScaled, toValueScaled } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import { isNoDebtHealthFactor } from "@/lib/units";
import { dashboardContent } from "@/content/dashboard";
import { assetPrices, collateralDecimals, debtDecimals } from "@/content/protocol";
import { useAccountData } from "@/hooks/useAccountData";
import { useMarketData } from "@/hooks/useMarketData";
import { useOraclePrice } from "@/hooks/useOraclePrice";
import { useWalletState } from "@/hooks/useWalletState";
import { Alert } from "@/components/ui/Alert";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { EmptyState } from "@/components/ui/EmptyState";
import { Skeleton } from "@/components/ui/Skeleton";
import { HealthBar } from "@/components/borrow/HealthBar";
import { HealthScoreGauge } from "@/components/borrow/HealthScoreGauge";
import { LiquidationRiskWarning } from "@/components/borrow/LiquidationRiskWarning";
import { RecentActivityList } from "@/components/dashboard/RecentActivityList";
import { RiskLegend } from "@/components/dashboard/RiskLegend";
import { MetricCard } from "@/components/portal/MetricCard";
import { MetricGrid } from "@/components/portal/MetricGrid";
import { PortalSection } from "@/components/portal/PortalSection";

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
      <div className="flex flex-col gap-4">
        <Skeleton className="h-24 w-full rounded-card" />
        <Skeleton className="h-56 w-full rounded-card" />
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
  const netValueScaled = suppliedValueScaled + data.collateralValue - data.debtValue;
  const isValued = !data.priceStale;
  const factorBps =
    !isValued || data.debtAmount <= 0n || isNoDebtHealthFactor(data.healthFactorBps) ? null : data.healthFactorBps;
  const tier = healthTier(factorBps);

  return (
    <>
      <MetricGrid>
        <MetricCard
          label="Supplied"
          value={formatValue(scaledValueToUsd(suppliedValueScaled), ValueFormat.UsdPrice)}
          hint={`${formatTokenAmount(data.supplyAssets, debtDecimals, 2)} ${AssetSymbol.Usdc} earning interest`}
        />
        <MetricCard
          label="Collateral"
          value={formatValue(scaledValueToUsd(data.collateralValue), ValueFormat.UsdPrice)}
          hint={`${formatTokenAmount(data.collateralAmount, collateralDecimals, 4)} ${AssetSymbol.Weth} locked`}
        />
        <MetricCard
          label="Borrowed"
          value={formatValue(scaledValueToUsd(data.debtValue), ValueFormat.UsdPrice)}
          hint={`${formatTokenAmount(data.debtAmount, debtDecimals, 2)} ${AssetSymbol.Usdc} owed`}
        />
        <MetricCard
          label="Net position"
          value={formatValue(scaledValueToUsd(netValueScaled), ValueFormat.UsdPrice)}
          hint="Supplied plus collateral, less debt"
          emphasis
        />
      </MetricGrid>

      {isValued ? null : (
        <Alert title="We cannot value your position right now" tone={BadgeTone.Caution} icon={IconName.Warning}>
          The WETH price feed has not reported recently, so your safety score cannot be calculated. Your collateral and
          loan are untouched, and nothing can be liquidated at a price the protocol does not trust.
        </Alert>
      )}

      {isValued ? <LiquidationRiskWarning tier={tier} /> : null}

      {isValued ? (
        <PortalSection title="Safety score" description={dashboardContent.overviewTitle}>
          <Card className="grid gap-6 p-5 lg:grid-cols-[minmax(0,1fr)_minmax(0,17rem)]">
            <div className="flex flex-col gap-5">
              <HealthScoreGauge factorBps={factorBps} tier={tier} />
              <HealthBar
                factorBps={factorBps}
                tier={tier}
                maxLtvBps={market.data.maxLtvBps}
                liquidationThresholdBps={market.data.liquidationThresholdBps}
              />
            </div>
            <div className="border-t border-line pt-5 lg:border-l lg:border-t-0 lg:pl-6 lg:pt-0">
              <RiskLegend />
            </div>
          </Card>
        </PortalSection>
      ) : null}

      <PortalSection
        title={dashboardContent.activityTitle}
        description={dashboardContent.activityDescription}
        actions={
          <Button href={AppRoute.History} variant={ButtonVariant.Subtle} trailingIcon={IconName.ArrowRight}>
            View all
          </Button>
        }
      >
        <RecentActivityList />
      </PortalSection>
    </>
  );
}
