import { AppRoute, ButtonSize, IconName, SectionId, SectionTone } from "@/lib/enums";
import { closingContent } from "@/content/landing";
import { Button } from "@/components/ui/Button";
import { Section } from "@/components/ui/Section";

export function ClosingCta() {
  return (
    <Section id={SectionId.Practice} tone={SectionTone.Canvas}>
      <div className="flex flex-col items-center gap-5 rounded-card border border-accent-border bg-accent-soft px-6 py-12 text-center">
        <h2 id={`${SectionId.Practice}-heading`} className="max-w-2xl text-balance text-3xl font-semibold tracking-tight text-ink">
          {closingContent.title}
        </h2>
        <p className="max-w-xl text-pretty text-base leading-relaxed text-ink-soft">{closingContent.description}</p>
        <Button href={AppRoute.Login} size={ButtonSize.Lg} trailingIcon={IconName.ArrowRight}>
          {closingContent.cta}
        </Button>
      </div>
    </Section>
  );
}
