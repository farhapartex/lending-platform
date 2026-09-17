"use client";

import { BadgeTone, OracleStatus, ValueFormat } from "@/lib/enums";
import { formatSecondsAgo, formatValue } from "@/lib/format";
import { collateralAsset } from "@/content/protocol";
import { priceToNumber, secondsSince, useOraclePrice } from "@/hooks/useOraclePrice";
import { Badge } from "@/components/ui/Badge";
import { Skeleton } from "@/components/ui/Skeleton";

const statusLabels: Record<OracleStatus, string> = {
  [OracleStatus.Fresh]: "Live price",
  [OracleStatus.Stale]: "Price may be stale",
  [OracleStatus.Unavailable]: "Price unavailable",
};

const statusTones: Record<OracleStatus, BadgeTone> = {
  [OracleStatus.Fresh]: BadgeTone.Positive,
  [OracleStatus.Stale]: BadgeTone.Caution,
  [OracleStatus.Unavailable]: BadgeTone.Critical,
};

const shell = "flex flex-col items-start gap-2 rounded-card border border-line bg-surface px-4 py-3 sm:items-end";

export function OraclePriceTicker() {
  const { data, status, isLoading, isError, isUnsupportedChain } = useOraclePrice();

  if (isUnsupportedChain || isError || (!isLoading && status === undefined)) {
    return (
      <div className={shell}>
        <Badge tone={BadgeTone.Critical}>{statusLabels[OracleStatus.Unavailable]}</Badge>
        <span className="text-sm text-ink-soft">Waiting for the price feed</span>
      </div>
    );
  }

  if (isLoading || data === undefined || status === undefined) {
    return (
      <div className={shell}>
        <Skeleton className="h-5 w-24" />
        <Skeleton className="h-6 w-40" />
      </div>
    );
  }

  return (
    <div className={shell}>
      <Badge tone={statusTones[status]}>{statusLabels[status]}</Badge>
      {status === OracleStatus.Unavailable ? (
        <span className="text-sm text-ink-soft">Waiting for the price feed</span>
      ) : (
        <div className="flex flex-wrap items-baseline gap-x-2 gap-y-0.5">
          <span className="text-lg font-semibold tracking-tight text-ink tabular-nums">
            {formatValue(priceToNumber(data), ValueFormat.UsdPrice)}
          </span>
          <span className="text-sm text-ink-soft">
            {collateralAsset} · updated {formatSecondsAgo(secondsSince(data.updatedAt))}
          </span>
        </div>
      )}
    </div>
  );
}
