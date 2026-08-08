import type { DocSection } from "@/content/learn";

type TableOfContentsProps = {
  sections: DocSection[];
};

export function TableOfContents({ sections }: TableOfContentsProps) {
  if (sections.length < 2) {
    return null;
  }

  return (
    <nav aria-label="On this page" className="rounded-card border border-line bg-surface-muted p-5">
      <h2 className="text-xs font-semibold uppercase tracking-[0.12em] text-ink-faint">On this page</h2>
      <ol className="mt-3 flex flex-col gap-2">
        {sections.map((section) => (
          <li key={section.id}>
            <a
              href={`#${section.id}`}
              className="rounded-sm text-sm text-ink-soft transition-colors hover:text-brand-ink outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-surface-muted"
            >
              {section.title}
            </a>
          </li>
        ))}
      </ol>
    </nav>
  );
}
