import type { DocPage } from "@/content/learn";
import { DocBlocks } from "@/components/learn/DocBlocks";
import { NextPrevNav } from "@/components/learn/NextPrevNav";
import { TableOfContents } from "@/components/learn/TableOfContents";

type DocPageViewProps = {
  page: DocPage;
};

export function DocPageView({ page }: DocPageViewProps) {
  return (
    <article className="flex flex-col gap-8">
      <header className="flex flex-col gap-3">
        <h1 className="text-3xl font-semibold tracking-tight text-ink sm:text-4xl">{page.title}</h1>
        <p className="max-w-2xl text-pretty text-base leading-relaxed text-ink-soft">{page.summary}</p>
      </header>

      <TableOfContents sections={page.sections} />

      <div className="flex flex-col gap-10">
        {page.sections.map((section) => (
          <section key={section.id} id={section.id} aria-labelledby={`${section.id}-heading`} className="flex flex-col gap-4">
            <h2 id={`${section.id}-heading`} className="text-xl font-semibold tracking-tight text-ink">
              {section.title}
            </h2>
            <DocBlocks blocks={section.blocks} />
          </section>
        ))}
      </div>

      <NextPrevNav currentKey={page.key} />
    </article>
  );
}
