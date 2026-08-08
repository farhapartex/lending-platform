import { IconName, SectionId, SectionTone } from "@/lib/enums";
import { practicePageContent, realAspects, unrealAspects } from "@/content/practice";
import { Icon } from "@/components/ui/Icon";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";

export function PracticeModeExplainer() {
  return (
    <Section id={SectionId.PracticeExplainer} tone={SectionTone.Canvas}>
      <SectionHeading
        sectionId={SectionId.PracticeExplainer}
        eyebrow="Ground rules"
        title={practicePageContent.explainerTitle}
        description={practicePageContent.explainerDescription}
      />

      <div className="mt-8 grid gap-5 md:grid-cols-2">
        <div className="flex flex-col gap-4 rounded-card border border-mint-border bg-mint-soft p-6">
          <h3 className="text-base font-semibold text-mint-ink">{practicePageContent.realTitle}</h3>
          <ul className="flex flex-col gap-3">
            {realAspects.map((item) => (
              <li key={item} className="flex gap-3 text-sm leading-relaxed text-ink-soft">
                <Icon name={IconName.Check} className="mt-0.5 size-4 shrink-0 text-mint" />
                <span>{item}</span>
              </li>
            ))}
          </ul>
        </div>

        <div className="flex flex-col gap-4 rounded-card border border-line bg-surface-muted p-6">
          <h3 className="text-base font-semibold text-ink">{practicePageContent.unrealTitle}</h3>
          <ul className="flex flex-col gap-3">
            {unrealAspects.map((item) => (
              <li key={item} className="flex gap-3 text-sm leading-relaxed text-ink-soft">
                <Icon name={IconName.Minus} className="mt-0.5 size-4 shrink-0 text-ink-faint" />
                <span>{item}</span>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </Section>
  );
}
