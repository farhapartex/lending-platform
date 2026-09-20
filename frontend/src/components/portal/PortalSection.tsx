import type { ReactNode } from "react";
import { cn } from "@/lib/cn";

type PortalSectionProps = {
  title?: string;
  description?: string;
  actions?: ReactNode;
  children: ReactNode;
  className?: string;
};

export function PortalSection({ title, description, actions, children, className }: PortalSectionProps) {
  const hasHeading = title !== undefined || description !== undefined || actions !== undefined;

  return (
    <section className={cn("flex flex-col gap-4", className)}>
      {hasHeading ? (
        <div className="flex flex-wrap items-end justify-between gap-3">
          <div className="flex flex-col gap-1">
            {title === undefined ? null : (
              <h2 className="text-base font-semibold tracking-tight text-ink">{title}</h2>
            )}
            {description === undefined ? null : (
              <p className="max-w-2xl text-sm leading-relaxed text-ink-soft">{description}</p>
            )}
          </div>

          {actions === undefined ? null : <div className="flex items-center gap-2">{actions}</div>}
        </div>
      ) : null}

      {children}
    </section>
  );
}
