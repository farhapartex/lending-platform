import type { ReactNode } from "react";
import { cn } from "@/lib/cn";

type MetricCardProps = {
  label: string;
  value: ReactNode;
  hint?: string;
  accessory?: ReactNode;
  emphasis?: boolean;
  className?: string;
};

export function MetricCard({ label, value, hint, accessory, emphasis = false, className }: MetricCardProps) {
  return (
    <div
      className={cn(
        "flex flex-col gap-1 rounded-card border border-line bg-surface p-4",
        emphasis && "border-accent-border bg-accent-soft",
        className,
      )}
    >
      <div className="flex items-center justify-between gap-2">
        <span className="text-xs font-medium uppercase tracking-[0.06em] text-ink-faint">{label}</span>
        {accessory}
      </div>

      <span
        className={cn(
          "text-xl font-semibold tracking-tight tabular-nums",
          emphasis ? "text-accent-ink" : "text-ink",
        )}
      >
        {value}
      </span>

      {hint === undefined ? null : <span className="text-xs leading-relaxed text-ink-faint">{hint}</span>}
    </div>
  );
}
