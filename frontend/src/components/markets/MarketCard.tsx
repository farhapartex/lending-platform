"use client";

import {
  DataStatus,
  MarketMetricKey,
  SectionId,
  SectionTone,
  SurfaceElevation,
  ValueFormat,
} from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { bpsToRatio } from "@/lib/units";
import { dataStatusFrom, debtAmountToNumber } from "@/lib/market";
import { marketsPageContent } from "@/content/markets";
import { useMarketData } from "@/hooks/useMarketData";
import { Card } from "@/components/ui/Card";
import { MetricRow } from "@/components/ui/MetricRow";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";
import { Skeleton } from "@/components/ui/Skeleton";
import { MarketUtilization } from "@/components/markets/MarketUtilization";

const emphasisedMetrics: MarketMetricKey[] = [MarketMetricKey.MaxLtv, MarketMetricKey.LiquidationThreshold];

export function MarketCard() {
  const { data, isLoading, isError } = useMarketData();
  const status = dataStatusFrom({ isLoading, isError, hasData: data !== undefined });

  if (status === DataStatus.Unavailable) {
    return null;
  }

  const metrics =
    data === undefined
      ? []
      : [
          {
            key: MarketMetricKey.SupplyApy,
            label: "Supply APY",
            value: bpsToRatio(data.supplyAprBps),
            format: ValueFormat.Percent,
            hint: "Earned by lenders, paid by borrowers",
          },
          {
            key: MarketMetricKey.BorrowApr,
            label: "Borrow APR",
            value: bpsToRatio(data.borrowAprBps),
            format: ValueFormat.Percent,
            hint: "Rises as the pool gets more utilized",
          },
          {
            key: MarketMetricKey.MaxLtv,
            label: "Max borrow",
            value: bpsToRatio(data.maxLtvBps),
            format: ValueFormat.Percent,
            hint: "Of your collateral value",
          },
          {
            key: MarketMetricKey.LiquidationThreshold,
            label: "Liquidation at",
            value: bpsToRatio(data.liquidationThresholdBps),
            format: ValueFormat.Percent,
            hint: "Deliberate buffer above max borrow",
          },
        ];

  return (
    <Section id={SectionId.MarketSummary} tone={SectionTone.Surface} bordered>
      <SectionHeading
        sectionId={SectionId.MarketSummary}
        eyebrow="Terms"
        title={marketsPageContent.summaryTitle}
        description={marketsPageContent.summaryDescription}
      />

      <div className="mt-8 grid gap-6 lg:grid-cols-2 lg:gap-8">
        <Card elevation={SurfaceElevation.Flat} className="p-6 sm:p-7">
          {data === undefined ? (
            <div className="flex flex-col gap-5">
              <Skeleton className="h-6 w-full" />
              <Skeleton className="h-6 w-full" />
              <Skeleton className="h-6 w-full" />
              <Skeleton className="h-6 w-full" />
              <Skeleton className="h-6 w-full" />
            </div>
          ) : (
            <dl className="divide-y divide-line">
              {metrics.map((metric) => (
                <MetricRow
                  key={metric.key}
                  label={metric.label}
                  value={formatValue(metric.value, metric.format)}
                  hint={metric.hint}
                  emphasised={emphasisedMetrics.includes(metric.key)}
                />
              ))}
              <MetricRow
                label="Available to borrow now"
                value={formatValue(debtAmountToNumber(data.availableLiquidity), ValueFormat.UsdCompact)}
                hint="Deposited funds not currently lent out."
              />
            </dl>
          )}
        </Card>

        <Card elevation={SurfaceElevation.Flat} className="flex h-full flex-col justify-center p-6 sm:p-7">
          <MarketUtilization />
        </Card>
      </div>
    </Section>
  );
}
