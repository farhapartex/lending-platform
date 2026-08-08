import { HealthTier } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { basisPoints, healthTierLowerBounds, healthTierOrder, healthTierUpperBounds } from "@/lib/health";
import { healthFillClasses, healthLabels } from "@/components/borrow/healthPresentation";

function boundToText(value: bigint | null): string {
  return value === null ? "" : (Number(value) / Number(basisPoints)).toFixed(2);
}

function rangeText(tier: HealthTier): string {
  const lower = healthTierLowerBounds[tier];
  const upper = healthTierUpperBounds[tier];

  if (lower !== null && upper === null) {
    return `${boundToText(lower)} and above`;
  }

  if (lower === null && upper !== null) {
    return `below ${boundToText(upper)}`;
  }

  return `${boundToText(lower)} up to ${boundToText(upper)}`;
}

export function RiskLegend() {
  return (
    <div className="flex flex-col gap-3">
      <span className="text-sm font-medium text-ink">What the safety levels mean</span>
      <dl className="flex flex-col gap-2">
        {healthTierOrder.map((tier) => (
          <div key={tier} className="flex items-center gap-2.5">
            <span aria-hidden="true" className={cn("size-2.5 rounded-pill", healthFillClasses[tier])} />
            <dt className="text-sm text-ink">{healthLabels[tier]}</dt>
            <dd className="text-sm text-ink-soft tabular-nums">{rangeText(tier)}</dd>
          </div>
        ))}
      </dl>
    </div>
  );
}
