import { AppRoute, IconName } from "@/lib/enums";
import { practiceContent } from "@/content/landing";
import { Icon } from "@/components/ui/Icon";
import { TextLink } from "@/components/ui/TextLink";

export function PracticeModeCard() {
  return (
    <div className="flex flex-col gap-3 rounded-card border border-brand-border bg-brand-soft p-5">
      <span className="flex items-center gap-2 text-xs font-semibold uppercase tracking-[0.14em] text-brand-ink">
        <Icon name={IconName.Beaker} className="size-4" />
        {practiceContent.eyebrow}
      </span>
      <p className="text-sm leading-relaxed text-ink-soft">{practiceContent.description}</p>
      <TextLink href={AppRoute.Learn} trailingIcon={IconName.ArrowRight}>
        Or read the documentation first
      </TextLink>
    </div>
  );
}
