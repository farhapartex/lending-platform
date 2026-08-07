import { AppRoute, ButtonSize, IconName, SectionId, SectionSpacing, SectionTone } from "@/lib/enums";
import { practiceContent } from "@/content/landing";
import { Button } from "@/components/ui/Button";
import { Icon } from "@/components/ui/Icon";
import { Section } from "@/components/ui/Section";

export function PracticeModeCta() {
  return (
    <Section id={SectionId.Practice} tone={SectionTone.Canvas} spacing={SectionSpacing.Regular}>
      <div className="relative overflow-hidden rounded-panel border border-brand-border bg-brand-soft px-6 py-10 sm:px-10 sm:py-12">
        <div
          aria-hidden="true"
          className="pointer-events-none absolute -right-16 -top-20 size-72 rounded-full bg-brand-muted/70 blur-3xl"
        />
        <div className="relative flex flex-col items-start gap-5 lg:flex-row lg:items-center lg:justify-between">
          <div className="flex max-w-2xl flex-col gap-3">
            <span className="inline-flex items-center gap-2 text-xs font-semibold uppercase tracking-[0.14em] text-brand-ink">
              <Icon name={IconName.Beaker} className="size-4" />
              {practiceContent.eyebrow}
            </span>
            <h2
              id={`${SectionId.Practice}-heading`}
              className="text-balance text-2xl font-semibold tracking-tight text-ink sm:text-3xl"
            >
              {practiceContent.title}
            </h2>
            <p className="text-pretty text-base leading-relaxed text-ink-soft">{practiceContent.description}</p>
          </div>
          <Button href={AppRoute.Practice} size={ButtonSize.Lg} trailingIcon={IconName.ArrowRight} className="shrink-0">
            {practiceContent.cta}
          </Button>
        </div>
      </div>
    </Section>
  );
}
