import { HealthTier, IconName } from "@/lib/enums";
import { formatHealthFactor } from "@/lib/health";
import { Icon } from "@/components/ui/Icon";
import { HealthBadge } from "@/components/borrow/HealthBadge";

type HealthImpactPreviewProps = {
  currentFactorBps: bigint | null;
  currentTier: HealthTier;
  nextFactorBps: bigint | null;
  nextTier: HealthTier;
};

export function HealthImpactPreview({
  currentFactorBps,
  currentTier,
  nextFactorBps,
  nextTier,
}: HealthImpactPreviewProps) {
  return (
    <div className="flex flex-col gap-3 rounded-card border border-line bg-surface p-4">
      <span className="text-sm font-medium text-ink">Effect on your safety score</span>

      <div className="flex flex-wrap items-center gap-3">
        <div className="flex flex-col gap-1.5">
          <span className="text-xs text-ink-faint">Now</span>
          <HealthBadge tier={currentTier} />
          <span className="text-xs text-ink-soft tabular-nums">{formatHealthFactor(currentFactorBps)}</span>
        </div>

        <Icon name={IconName.ArrowRight} className="size-4 text-ink-faint" />

        <div className="flex flex-col gap-1.5">
          <span className="text-xs text-ink-faint">After this change</span>
          <HealthBadge tier={nextTier} />
          <span className="text-xs text-ink-soft tabular-nums">{formatHealthFactor(nextFactorBps)}</span>
        </div>
      </div>
    </div>
  );
}
