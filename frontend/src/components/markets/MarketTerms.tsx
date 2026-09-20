"use client";

import { MarketMetricKey, SurfaceElevation, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { bpsToRatio } from "@/lib/units";
import { useMarketData } from "@/hooks/useMarketData";
import { Card } from "@/components/ui/Card";
import { MetricRow } from "@/components/ui/MetricRow";
import { Skeleton } from "@/components/ui/Skeleton";
import { MarketUtilization } from "@/components/markets/MarketUtilization";

const emphasisedMetrics: MarketMetricKey[] = [MarketMetricKey.MaxLtv, MarketMetricKey.LiquidationThreshold];

export function MarketTerms() {
  const { data } = useMarketData();

  if (data === undefined) {
    return (
      <div className="grid gap-4 lg:grid-cols-2">
        <Skeleton className="h-64 w-full rounded-card" />
        <Skeleton className="h-64 w-full rounded-card" />
      </div>
    );
  }

  const metrics = [
    {
      key: MarketMetricKey.SupplyApy,
      label: "Supply APY",
      value: bpsToRatio(data.supplyAprBps),
      hint: "Earned by lenders, paid by borrowers",
    },
    {
      key: MarketMetricKey.BorrowApr,
      label: "Borrow APR",
      value: bpsToRatio(data.borrowAprBps),
      hint: "Rises as the pool gets more utilized",
    },
    {
      key: MarketMetricKey.MaxLtv,
      label: "Max borrow",
      value: bpsToRatio(data.maxLtvBps),
      hint: "Of your collateral value",
    },
    {
      key: MarketMetricKey.LiquidationThreshold,
      label: "Liquidation at",
      value: bpsToRatio(data.liquidationThresholdBps),
      hint: "Deliberate buffer above max borrow",
    },
  ];

  return (
    <div className="grid gap-4 lg:grid-cols-2">
      <Card elevation={SurfaceElevation.Flat} className="p-5">
        <dl className="divide-y divide-line">
          {metrics.map((metric) => (
            <MetricRow
              key={metric.key}
              label={metric.label}
              value={formatValue(metric.value, ValueFormat.Percent)}
              hint={metric.hint}
              emphasised={emphasisedMetrics.includes(metric.key)}
            />
          ))}
        </dl>
      </Card>

      <Card elevation={SurfaceElevation.Flat} className="flex flex-col justify-center p-5">
        <MarketUtilization />
      </Card>
    </div>
  );
}
