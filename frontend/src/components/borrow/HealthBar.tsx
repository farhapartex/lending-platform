import { HealthTier, ValueFormat } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { formatValue } from "@/lib/format";
import { basisPoints } from "@/lib/health";
import { healthFillClasses } from "@/components/borrow/healthPresentation";

const scaleCeilingBps = 20_000n;

function toPercentOfScale(valueBps: bigint): number {
  const clamped = valueBps > scaleCeilingBps ? scaleCeilingBps : valueBps;
  return (Number(clamped) / Number(scaleCeilingBps)) * 100;
}

type HealthBarProps = {
  factorBps: bigint | null;
  tier: HealthTier;
  maxLtvBps: bigint;
  liquidationThresholdBps: bigint;
  className?: string;
};

export function HealthBar({ factorBps, tier, maxLtvBps, liquidationThresholdBps, className }: HealthBarProps) {
  const fillPercent = factorBps === null ? 100 : toPercentOfScale(factorBps);
  const liquidationMarkerPercent = toPercentOfScale(10_000n);
  const borrowLimitMarkerPercent = toPercentOfScale((liquidationThresholdBps * 10_000n) / maxLtvBps);

  return (
    <div className={cn("flex flex-col gap-3", className)}>
      <div className="relative">
        <div
          role="progressbar"
          aria-label="Position health factor"
          aria-valuemin={0}
          aria-valuemax={Number(scaleCeilingBps) / Number(basisPoints)}
          aria-valuenow={factorBps === null ? undefined : Number(factorBps) / Number(basisPoints)}
          aria-valuetext={factorBps === null ? "No loan" : undefined}
          className="h-2.5 w-full overflow-hidden rounded-pill bg-canvas-deep"
        >
          <div className={cn("h-full rounded-pill", healthFillClasses[tier])} style={{ width: `${fillPercent}%` }} />
        </div>

        <span
          aria-hidden="true"
          className="absolute -top-1 h-4.5 w-0.5 -translate-x-1/2 rounded-pill bg-rose"
          style={{ left: `${liquidationMarkerPercent}%` }}
        />
        <span
          aria-hidden="true"
          className="absolute -top-1 h-4.5 w-0.5 -translate-x-1/2 rounded-pill bg-ink-faint"
          style={{ left: `${borrowLimitMarkerPercent}%` }}
        />
      </div>

      <dl className="flex flex-wrap gap-x-6 gap-y-1.5 text-xs">
        <div className="flex items-center gap-1.5">
          <span aria-hidden="true" className="h-2.5 w-0.5 rounded-pill bg-rose" />
          <dt className="text-ink-soft">Liquidation at</dt>
          <dd className="font-medium text-ink tabular-nums">1.00</dd>
        </div>
        <div className="flex items-center gap-1.5">
          <span aria-hidden="true" className="h-2.5 w-0.5 rounded-pill bg-ink-faint" />
          <dt className="text-ink-soft">Borrow limit reached</dt>
          <dd className="font-medium text-ink tabular-nums">
            {(Number(liquidationThresholdBps) / Number(maxLtvBps)).toFixed(2)}
          </dd>
        </div>
        <div className="flex items-center gap-1.5">
          <dt className="text-ink-soft">Max borrow / liquidation</dt>
          <dd className="font-medium text-ink tabular-nums">
            {formatValue(Number(maxLtvBps) / Number(basisPoints), ValueFormat.Percent)} /{" "}
            {formatValue(Number(liquidationThresholdBps) / Number(basisPoints), ValueFormat.Percent)}
          </dd>
        </div>
      </dl>
    </div>
  );
}
