import Link from "next/link";
import { DocKey, IconName } from "@/lib/enums";
import { docPages } from "@/content/learn";
import { Icon } from "@/components/ui/Icon";

const cardClasses =
  "flex flex-1 flex-col gap-1 rounded-card border border-line bg-surface p-4 transition-colors hover:border-brand-border outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas";

type NextPrevNavProps = {
  currentKey: DocKey;
};

export function NextPrevNav({ currentKey }: NextPrevNavProps) {
  const index = docPages.findIndex((page) => page.key === currentKey);
  const previous = index > 0 ? docPages[index - 1] : undefined;
  const next = index >= 0 && index < docPages.length - 1 ? docPages[index + 1] : undefined;

  if (previous === undefined && next === undefined) {
    return null;
  }

  return (
    <nav aria-label="Previous and next page" className="flex flex-col gap-3 sm:flex-row">
      {previous === undefined ? null : (
        <Link href={previous.route} className={cardClasses}>
          <span className="flex items-center gap-1.5 text-xs text-ink-faint">
            <Icon name={IconName.ArrowRight} className="size-3.5 rotate-180" />
            Previous
          </span>
          <span className="text-sm font-medium text-ink">{previous.title}</span>
        </Link>
      )}

      {next === undefined ? null : (
        <Link href={next.route} className={`${cardClasses} sm:items-end sm:text-right`}>
          <span className="flex items-center gap-1.5 text-xs text-ink-faint">
            Next
            <Icon name={IconName.ArrowRight} className="size-3.5" />
          </span>
          <span className="text-sm font-medium text-ink">{next.title}</span>
        </Link>
      )}
    </nav>
  );
}
