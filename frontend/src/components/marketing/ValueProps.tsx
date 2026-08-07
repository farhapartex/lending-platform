import { SectionId, SectionTone, SurfaceElevation } from "@/lib/enums";
import { valueProps } from "@/content/landing";
import { Card } from "@/components/ui/Card";
import { Icon } from "@/components/ui/Icon";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";

export function ValueProps() {
  return (
    <Section id={SectionId.ValueProps} tone={SectionTone.Canvas}>
      <SectionHeading
        sectionId={SectionId.ValueProps}
        eyebrow="Why this platform"
        title="Two ways to use it, one set of rules that never change on you."
        description="Whether you are lending or borrowing, the same published parameters apply to every user. No tiers, no special treatment, no hidden terms."
      />

      <ul className="mt-10 grid gap-5 md:grid-cols-3">
        {valueProps.map((prop) => (
          <li key={prop.key}>
            <Card elevation={SurfaceElevation.Raised} interactive className="flex h-full flex-col gap-4 p-6">
              <span className="grid size-11 place-items-center rounded-tile bg-brand-soft text-brand">
                <Icon name={prop.icon} className="size-5" />
              </span>
              <h3 className="text-lg font-semibold tracking-tight text-ink">{prop.title}</h3>
              <p className="text-sm leading-relaxed text-ink-soft">{prop.description}</p>
            </Card>
          </li>
        ))}
      </ul>
    </Section>
  );
}
