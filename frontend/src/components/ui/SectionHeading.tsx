import { SectionId, TextAlign } from "@/lib/enums";
import { cn } from "@/lib/cn";

const alignClasses: Record<TextAlign, string> = {
  [TextAlign.Left]: "items-start text-left",
  [TextAlign.Center]: "items-center text-center",
};

type SectionHeadingProps = {
  sectionId: SectionId;
  title: string;
  eyebrow?: string;
  description?: string;
  align?: TextAlign;
  className?: string;
};

export function SectionHeading({
  sectionId,
  title,
  eyebrow,
  description,
  align = TextAlign.Left,
  className,
}: SectionHeadingProps) {
  return (
    <div className={cn("flex max-w-2xl flex-col gap-3", alignClasses[align], align === TextAlign.Center && "mx-auto", className)}>
      {eyebrow ? (
        <span className="text-xs font-semibold uppercase tracking-[0.14em] text-brand">{eyebrow}</span>
      ) : null}
      <h2 id={`${sectionId}-heading`} className="text-balance text-2xl font-semibold tracking-tight text-ink sm:text-3xl">
        {title}
      </h2>
      {description ? <p className="text-pretty text-base leading-relaxed text-ink-soft">{description}</p> : null}
    </div>
  );
}
