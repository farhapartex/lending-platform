import { SectionId, SectionTone, SurfaceElevation } from "@/lib/enums";
import { practiceIdeas, practicePageContent } from "@/content/practice";
import { Card } from "@/components/ui/Card";
import { Icon } from "@/components/ui/Icon";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";

export function ThingsToTryList() {
  return (
    <Section id={SectionId.PracticeIdeas} tone={SectionTone.Surface} bordered>
      <SectionHeading
        sectionId={SectionId.PracticeIdeas}
        eyebrow="Suggestions"
        title={practicePageContent.ideasTitle}
        description={practicePageContent.ideasDescription}
      />

      <ul className="mt-8 grid gap-5 sm:grid-cols-2">
        {practiceIdeas.map((idea) => (
          <li key={idea.id}>
            <Card elevation={SurfaceElevation.Flat} className="flex h-full flex-col gap-3 p-6">
              <span className="grid size-10 place-items-center rounded-tile bg-brand-soft text-brand">
                <Icon name={idea.icon} className="size-5" />
              </span>
              <h3 className="text-base font-semibold tracking-tight text-ink">{idea.title}</h3>
              <p className="text-sm leading-relaxed text-ink-soft">{idea.description}</p>
            </Card>
          </li>
        ))}
      </ul>
    </Section>
  );
}
