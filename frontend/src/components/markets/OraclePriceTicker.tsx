import { BadgeTone, OracleStatus, ValueFormat } from "@/lib/enums";
import { formatSecondsAgo, formatValue } from "@/lib/format";
import { oracleReading } from "@/content/protocol";
import { Badge } from "@/components/ui/Badge";

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

export function OraclePriceTicker() {
  return (
    <div className="flex flex-col items-start gap-2 rounded-card border border-line bg-surface px-4 py-3 sm:items-end">
      <Badge tone={statusTones[oracleReading.status]}>{statusLabels[oracleReading.status]}</Badge>
      {oracleReading.status === OracleStatus.Unavailable ? (
        <span className="text-sm text-ink-soft">Waiting for the price feed</span>
      ) : (
        <div className="flex flex-wrap items-baseline gap-x-2 gap-y-0.5">
          <span className="text-lg font-semibold tracking-tight text-ink tabular-nums">
            {formatValue(oracleReading.price, ValueFormat.UsdPrice)}
          </span>
          <span className="text-sm text-ink-soft">
            {oracleReading.symbol} · updated {formatSecondsAgo(oracleReading.updatedSecondsAgo)}
          </span>
        </div>
      )}
    </div>
  );
}
