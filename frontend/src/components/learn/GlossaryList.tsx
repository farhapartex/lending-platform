import { glossaryEntries } from "@/content/learn";

export function GlossaryList() {
  return (
    <dl className="divide-y divide-line overflow-hidden rounded-card border border-line bg-surface">
      {glossaryEntries.map((entry) => (
        <div key={entry.term} className="flex flex-col gap-1.5 px-5 py-4">
          <dt className="text-sm font-semibold text-ink">{entry.term}</dt>
          <dd className="text-sm leading-relaxed text-ink-soft">{entry.definition}</dd>
        </div>
      ))}
    </dl>
  );
}
