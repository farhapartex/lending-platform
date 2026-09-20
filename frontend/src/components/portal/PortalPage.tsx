import type { ReactNode } from "react";
import { cn } from "@/lib/cn";

type PortalPageProps = {
  title: string;
  description?: string;
  actions?: ReactNode;
  children: ReactNode;
  className?: string;
};

export function PortalPage({ title, description, actions, children, className }: PortalPageProps) {
  return (
    <div className={cn("mx-auto w-full max-w-7xl px-4 py-6 sm:px-6 sm:py-8", className)}>
      <header className="flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
        <div className="flex flex-col gap-1.5">
          <h1 className="text-2xl font-semibold tracking-tight text-ink">{title}</h1>
          {description === undefined ? null : (
            <p className="max-w-2xl text-sm leading-relaxed text-ink-soft">{description}</p>
          )}
        </div>

        {actions === undefined ? null : <div className="flex shrink-0 items-center gap-2">{actions}</div>}
      </header>

      <div className="mt-6 flex flex-col gap-6">{children}</div>
    </div>
  );
}
