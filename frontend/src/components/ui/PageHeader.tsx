import type { ReactNode } from "react";
import { Container } from "@/components/ui/Container";

type PageHeaderProps = {
  title: string;
  description?: string;
  badge?: ReactNode;
  aside?: ReactNode;
  children?: ReactNode;
};

export function PageHeader({ title, description, badge, aside, children }: PageHeaderProps) {
  return (
    <div className="border-b border-line bg-surface">
      <Container className="flex flex-col gap-8 py-10 lg:flex-row lg:items-start lg:justify-between">
        <div className="flex flex-col items-start gap-4">
          <div className="flex flex-wrap items-center gap-3">
            <h1 className="text-3xl font-semibold tracking-tight text-ink sm:text-4xl">{title}</h1>
            {badge}
          </div>
          {description ? (
            <p className="max-w-2xl text-pretty text-base leading-relaxed text-ink-soft">{description}</p>
          ) : null}
          {children}
        </div>
        {aside}
      </Container>
    </div>
  );
}
