"use client";

import { AssetSymbol, BadgeTone, OracleStatus, ValueFormat } from "@/lib/enums";
import { formatSecondsAgo, formatValue } from "@/lib/format";
import { debtAmountToNumber } from "@/lib/market";
import { bpsToRatio } from "@/lib/units";
import { priceToNumber, secondsSince, useOraclePrice } from "@/hooks/useOraclePrice";
import { useMarketData } from "@/hooks/useMarketData";
import { Badge } from "@/components/ui/Badge";
import { Skeleton } from "@/components/ui/Skeleton";
import { MetricCard } from "@/components/portal/MetricCard";
import { MetricGrid } from "@/components/portal/MetricGrid";

const statusTones: Record<OracleStatus, BadgeTone> = {
  [OracleStatus.Fresh]: BadgeTone.Positive,
  [OracleStatus.Stale]: BadgeTone.Caution,
  [OracleStatus.Unavailable]: BadgeTone.Critical,
};

const statusLabels: Record<OracleStatus, string> = {
  [OracleStatus.Fresh]: "Live",
  [OracleStatus.Stale]: "Stale",
  [OracleStatus.Unavailable]: "Down",
};

export function MarketOverviewMetrics() {
  const { data } = useMarketData();
  const price = useOraclePrice();

  if (data === undefined) {
    return (
      <MetricGrid>
        {[0, 1, 2, 3].map((slot) => (
          <Skeleton key={slot} className="h-24 w-full rounded-card" />
        ))}
      </MetricGrid>
    );
  }

  return (
    <MetricGrid>
      <MetricCard
        label={`${AssetSymbol.Weth} price`}
        value={
          price.data === undefined
            ? "—"
            : formatValue(priceToNumber(price.data), ValueFormat.UsdPrice)
        }
        hint={
          price.data === undefined
            ? "Waiting for the price feed"
            : `Updated ${formatSecondsAgo(secondsSince(price.data.updatedAt))}`
        }
        accessory={
          price.status === undefined ? null : (
            <Badge tone={statusTones[price.status]}>{statusLabels[price.status]}</Badge>
          )
        }
      />
      <MetricCard
        label="Total deposited"
        value={formatValue(debtAmountToNumber(data.totalSupplied), ValueFormat.UsdCompact)}
        hint="Supplied by lenders"
      />
      <MetricCard
        label="Total borrowed"
        value={formatValue(debtAmountToNumber(data.totalBorrowed), ValueFormat.UsdCompact)}
        hint="Currently lent out"
      />
      <MetricCard
        label="Available to borrow"
        value={formatValue(debtAmountToNumber(data.availableLiquidity), ValueFormat.UsdCompact)}
        hint={`${formatValue(bpsToRatio(data.utilizationBps), ValueFormat.Percent)} utilized`}
        emphasis
      />
    </MetricGrid>
  );
}
