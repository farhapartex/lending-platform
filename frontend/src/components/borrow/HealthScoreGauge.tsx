import { HealthTier } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { formatHealthFactor } from "@/lib/health";
import { Icon } from "@/components/ui/Icon";
import {
  healthExplanations,
  healthIcons,
  healthLabels,
  healthTextClasses,
} from "@/components/borrow/healthPresentation";

type HealthScoreGaugeProps = {
  factorBps: bigint | null;
  tier: HealthTier;
  className?: string;
};

export function HealthScoreGauge({ factorBps, tier, className }: HealthScoreGaugeProps) {
  return (
    <div className={cn("flex flex-col gap-3", className)}>
      <span className="text-sm text-ink-soft">Your position is</span>

      <div className="flex items-center gap-3">
        <span className={cn("grid size-11 place-items-center rounded-tile bg-surface-muted", healthTextClasses[tier])}>
          <Icon name={healthIcons[tier]} className="size-6" />
        </span>
        <span className={cn("text-3xl font-semibold tracking-tight", healthTextClasses[tier])}>
          {healthLabels[tier]}
        </span>
      </div>

      <p className="max-w-xl text-sm leading-relaxed text-ink-soft">{healthExplanations[tier]}</p>

      <p className="text-xs text-ink-faint">
        Health factor <span className="font-medium tabular-nums text-ink-soft">{formatHealthFactor(factorBps)}</span>.
        Liquidation becomes possible below 1.00.
      </p>
    </div>
  );
}
