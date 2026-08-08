import type { Metadata } from "next";
import Link from "next/link";
import { IconName } from "@/lib/enums";
import { docPages, learnIndexContent } from "@/content/learn";
import { Icon } from "@/components/ui/Icon";

export const metadata: Metadata = {
  title: "Learn",
  description: learnIndexContent.description,
};

export default function LearnIndexPage() {
  return (
    <div className="flex flex-col gap-8">
      <header className="flex flex-col gap-3">
        <h1 className="text-3xl font-semibold tracking-tight text-ink sm:text-4xl">{learnIndexContent.title}</h1>
        <p className="max-w-2xl text-pretty text-base leading-relaxed text-ink-soft">
          {learnIndexContent.description}
        </p>
      </header>

      <ul className="grid gap-4 sm:grid-cols-2">
        {docPages.map((page) => (
          <li key={page.key}>
            <Link
              href={page.route}
              className="flex h-full flex-col gap-2 rounded-card border border-line bg-surface p-5 transition-shadow hover:shadow-soft outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas"
            >
              <span className="flex items-center justify-between gap-3">
                <span className="text-base font-semibold tracking-tight text-ink">{page.title}</span>
                <Icon name={IconName.ArrowRight} className="size-4 text-ink-faint" />
              </span>
              <span className="text-sm leading-relaxed text-ink-soft">{page.summary}</span>
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
}
