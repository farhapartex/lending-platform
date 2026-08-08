import { IconName } from "@/lib/enums";
import { faqItems } from "@/content/learn";
import { Icon } from "@/components/ui/Icon";

export function FaqAccordion() {
  return (
    <div className="divide-y divide-line overflow-hidden rounded-card border border-line bg-surface">
      {faqItems.map((item) => (
        <details key={item.id} className="group">
          <summary className="flex cursor-pointer items-center justify-between gap-4 px-5 py-4 text-sm font-medium text-ink outline-none marker:content-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-inset">
            {item.question}
            <Icon
              name={IconName.ChevronDown}
              className="size-4 shrink-0 text-ink-faint transition-transform group-open:rotate-180"
            />
          </summary>
          <p className="px-5 pb-4 text-sm leading-relaxed text-ink-soft">{item.answer}</p>
        </details>
      ))}
    </div>
  );
}
