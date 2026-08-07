import { cn } from "@/lib/cn";

type MetricRowProps = {
  label: string;
  value: string;
  hint?: string;
  emphasised?: boolean;
  className?: string;
};

export function MetricRow({ label, value, hint, emphasised = false, className }: MetricRowProps) {
  return (
    <div className={cn("flex flex-col gap-1.5 py-4", className)}>
      <div className="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-1">
        <span className="text-sm font-medium text-ink">{label}</span>
        <span
          className={cn(
            "text-base font-semibold tabular-nums",
            emphasised ? "text-brand-ink" : "text-ink",
          )}
        >
          {value}
        </span>
      </div>
      {hint ? <p className="max-w-xl text-sm leading-relaxed text-ink-soft">{hint}</p> : null}
    </div>
  );
}
