import { TrendDirection, ValueFormat } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { formatValue } from "@/lib/format";
import { Skeleton } from "@/components/ui/Skeleton";
import { TrendPill } from "@/components/ui/TrendPill";

type StatTileProps = {
  label: string;
  value: number;
  format: ValueFormat;
  trend?: TrendDirection;
  trendLabel?: string;
  className?: string;
};

export function StatTile({ label, value, format, trend, trendLabel, className }: StatTileProps) {
  return (
    <div className={cn("flex flex-col gap-1.5", className)}>
      <span className="text-sm text-ink-soft">{label}</span>
      <span className="text-2xl font-semibold tracking-tight text-ink tabular-nums sm:text-[1.75rem]">
        {formatValue(value, format)}
      </span>
      {trend !== undefined && trendLabel !== undefined ? <TrendPill direction={trend} label={trendLabel} /> : null}
    </div>
  );
}

export function StatTileSkeleton({ className }: { className?: string }) {
  return (
    <div className={cn("flex flex-col gap-2", className)}>
      <Skeleton className="h-4 w-24 rounded-pill" />
      <Skeleton className="h-8 w-32" />
      <Skeleton className="h-3.5 w-20 rounded-pill" />
    </div>
  );
}
