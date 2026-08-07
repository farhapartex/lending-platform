import { SectionId } from "@/lib/enums";

export function SkipLink() {
  return (
    <a
      href={`#${SectionId.MainContent}`}
      className="sr-only focus:not-sr-only focus:fixed focus:left-4 focus:top-4 focus:z-50 focus:rounded-pill focus:bg-brand focus:px-4 focus:py-2 focus:text-sm focus:font-medium focus:text-white"
    >
      Skip to main content
    </a>
  );
}
