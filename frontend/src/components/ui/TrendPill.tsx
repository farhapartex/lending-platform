import { IconName, TrendDirection } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { Icon } from "@/components/ui/Icon";

const iconByDirection: Record<TrendDirection, IconName> = {
  [TrendDirection.Up]: IconName.TrendUp,
  [TrendDirection.Down]: IconName.TrendDown,
  [TrendDirection.Flat]: IconName.Minus,
};

const colorByDirection: Record<TrendDirection, string> = {
  [TrendDirection.Up]: "text-mint-ink",
  [TrendDirection.Down]: "text-rose-ink",
  [TrendDirection.Flat]: "text-ink-faint",
};

type TrendPillProps = {
  direction: TrendDirection;
  label: string;
  className?: string;
};

export function TrendPill({ direction, label, className }: TrendPillProps) {
  return (
    <span className={cn("inline-flex items-center gap-1.5 text-xs font-medium", colorByDirection[direction], className)}>
      <Icon name={iconByDirection[direction]} className="size-3.5" />
      {label}
    </span>
  );
}
