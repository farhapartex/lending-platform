import { cn } from "@/lib/cn";

type GlossaryTermProps = {
  term: string;
  definition: string;
  className?: string;
};

export function GlossaryTerm({ term, definition, className }: GlossaryTermProps) {
  return (
    <span
      tabIndex={0}
      title={definition}
      aria-label={`${term}: ${definition}`}
      className={cn(
        "cursor-help underline decoration-dotted decoration-from-font underline-offset-2 outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas",
        className,
      )}
    >
      {term}
    </span>
  );
}
