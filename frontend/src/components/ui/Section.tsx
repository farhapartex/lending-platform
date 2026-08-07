import type { ReactNode } from "react";
import { SectionId, SectionSpacing, SectionTone } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { Container } from "@/components/ui/Container";

const toneClasses: Record<SectionTone, string> = {
  [SectionTone.Canvas]: "bg-canvas",
  [SectionTone.Surface]: "bg-surface",
  [SectionTone.Muted]: "bg-surface-muted",
};

const spacingClasses: Record<SectionSpacing, string> = {
  [SectionSpacing.Compact]: "py-12 sm:py-14",
  [SectionSpacing.Regular]: "py-16 sm:py-20",
  [SectionSpacing.Spacious]: "py-20 sm:py-28",
};

type SectionProps = {
  id: SectionId;
  children: ReactNode;
  tone?: SectionTone;
  spacing?: SectionSpacing;
  bordered?: boolean;
  className?: string;
};

export function Section({
  id,
  children,
  tone = SectionTone.Canvas,
  spacing = SectionSpacing.Regular,
  bordered = false,
  className,
}: SectionProps) {
  return (
    <section
      id={id}
      aria-labelledby={`${id}-heading`}
      className={cn(toneClasses[tone], spacingClasses[spacing], bordered && "border-y border-line", className)}
    >
      <Container>{children}</Container>
    </section>
  );
}
