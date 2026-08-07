import type { ReactNode } from "react";
import { IconName } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { Icon } from "@/components/ui/Icon";

type EmptyStateProps = {
  title: string;
  description?: string;
  icon?: IconName;
  action?: ReactNode;
  className?: string;
};

export function EmptyState({ title, description, icon, action, className }: EmptyStateProps) {
  return (
    <div
      className={cn(
        "flex flex-col items-center gap-3 rounded-card border border-dashed border-line-strong bg-surface-muted px-6 py-10 text-center",
        className,
      )}
    >
      {icon ? (
        <span className="grid size-11 place-items-center rounded-tile bg-surface text-ink-faint">
          <Icon name={icon} className="size-5" />
        </span>
      ) : null}
      <p className="text-base font-semibold text-ink">{title}</p>
      {description ? <p className="max-w-md text-sm leading-relaxed text-ink-soft">{description}</p> : null}
      {action}
    </div>
  );
}
