import type { ReactNode } from "react";
import { BadgeTone, IconName } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { Icon } from "@/components/ui/Icon";

const toneClasses: Record<BadgeTone, string> = {
  [BadgeTone.Neutral]: "border-line bg-surface-muted text-ink-soft",
  [BadgeTone.Brand]: "border-brand-border bg-brand-soft text-brand-ink",
  [BadgeTone.Positive]: "border-mint-border bg-mint-soft text-mint-ink",
  [BadgeTone.Caution]: "border-amber-border bg-amber-soft text-amber-ink",
  [BadgeTone.Critical]: "border-rose-border bg-rose-soft text-rose-ink",
};

type BadgeProps = {
  children: ReactNode;
  tone?: BadgeTone;
  icon?: IconName;
  className?: string;
};

export function Badge({ children, tone = BadgeTone.Neutral, icon, className }: BadgeProps) {
  return (
    <span
      className={cn(
        "inline-flex items-center gap-1.5 rounded-pill border px-3 py-1 text-xs font-medium",
        toneClasses[tone],
        className,
      )}
    >
      {icon ? <Icon name={icon} className="size-3.5" /> : null}
      {children}
    </span>
  );
}
