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

type AlertProps = {
  title: string;
  tone?: BadgeTone;
  icon?: IconName;
  children?: ReactNode;
  className?: string;
};

export function Alert({ title, tone = BadgeTone.Neutral, icon, children, className }: AlertProps) {
  return (
    <div role="status" className={cn("flex gap-3 rounded-card border p-4", toneClasses[tone], className)}>
      {icon ? <Icon name={icon} className="mt-0.5 size-4.5" /> : null}
      <div className="flex flex-col gap-1">
        <p className="text-sm font-semibold">{title}</p>
        {children ? <div className="text-sm leading-relaxed opacity-90">{children}</div> : null}
      </div>
    </div>
  );
}
