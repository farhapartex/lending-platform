import { UtilizationZone, ValueFormat } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { formatValue } from "@/lib/format";
import { utilizationModel } from "@/content/protocol";

const fillClasses: Record<UtilizationZone, string> = {
  [UtilizationZone.BelowKink]: "bg-brand",
  [UtilizationZone.AboveKink]: "bg-amber",
};

const zoneMessages: Record<UtilizationZone, string> = {
  [UtilizationZone.BelowKink]: "Below the kink, so rates are still in their gentle range.",
  [UtilizationZone.AboveKink]: "Above the kink, so rates are climbing steeply to protect withdrawals.",
};

type UtilizationBarProps = {
  className?: string;
};

export function UtilizationBar({ className }: UtilizationBarProps) {
  const { current, kink } = utilizationModel;
  const zone = current < kink ? UtilizationZone.BelowKink : UtilizationZone.AboveKink;
  const currentPercent = current * 100;
  const kinkPercent = kink * 100;

  return (
    <div className={cn("flex flex-col gap-3", className)}>
      <div className="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-1">
        <span className="text-sm font-medium text-ink">Pool utilization</span>
        <span className="text-base font-semibold text-ink tabular-nums">
          {formatValue(current, ValueFormat.Percent)}
        </span>
      </div>

      <div className="relative">
        <div
          role="progressbar"
          aria-label="Pool utilization"
          aria-valuemin={0}
          aria-valuemax={100}
          aria-valuenow={Number(currentPercent.toFixed(2))}
          className="h-2.5 w-full overflow-hidden rounded-pill bg-canvas-deep"
        >
          <div className={cn("h-full rounded-pill", fillClasses[zone])} style={{ width: `${currentPercent}%` }} />
        </div>

        <div
          aria-hidden="true"
          className="absolute -top-1 h-4.5 w-0.5 -translate-x-1/2 rounded-pill bg-ink-faint"
          style={{ left: `${kinkPercent}%` }}
        />
      </div>

      <div className="flex flex-wrap items-center justify-between gap-x-4 gap-y-1 text-xs text-ink-faint">
        <span>0%</span>
        <span className="font-medium text-ink-soft">Rate kink at {formatValue(kink, ValueFormat.Percent)}</span>
        <span>100%</span>
      </div>

      <p className="text-sm leading-relaxed text-ink-soft">{zoneMessages[zone]}</p>
    </div>
  );
}
