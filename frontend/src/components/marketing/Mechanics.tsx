import { SectionId, SectionTone, SurfaceElevation } from "@/lib/enums";
import { mechanics, mechanicsContent } from "@/content/landing";
import { Card } from "@/components/ui/Card";
import { Icon } from "@/components/ui/Icon";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";

export function Mechanics() {
  return (
    <Section id={SectionId.Mechanics} tone={SectionTone.Surface} bordered>
      <SectionHeading
        sectionId={SectionId.Mechanics}
        eyebrow={mechanicsContent.eyebrow}
        title={mechanicsContent.title}
        description={mechanicsContent.description}
      />

      <ul className="mt-10 grid gap-5 md:grid-cols-3">
        {mechanics.map((mechanic) => (
          <li key={mechanic.key}>
            <Card elevation={SurfaceElevation.Flat} className="flex h-full flex-col gap-3 p-6">
              <span className="grid size-10 place-items-center rounded-tile bg-accent-soft text-accent">
                <Icon name={mechanic.icon} className="size-5" />
              </span>
              <h3 className="text-base font-semibold tracking-tight text-ink">{mechanic.title}</h3>
              <p className="text-sm leading-relaxed text-ink-soft">{mechanic.description}</p>
            </Card>
          </li>
        ))}
      </ul>
    </Section>
  );
}
