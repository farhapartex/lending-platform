import { SectionId, SectionTone, SurfaceElevation } from "@/lib/enums";
import { marketsPageContent, rateExplainerPoints } from "@/content/markets";
import { Card } from "@/components/ui/Card";
import { Icon } from "@/components/ui/Icon";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";

export function RateExplainer() {
  return (
    <Section id={SectionId.MarketRates} tone={SectionTone.Canvas}>
      <SectionHeading
        sectionId={SectionId.MarketRates}
        eyebrow="Rates"
        title={marketsPageContent.ratesTitle}
        description={marketsPageContent.ratesDescription}
      />

      <ul className="mt-8 grid gap-5 md:grid-cols-3">
        {rateExplainerPoints.map((point) => (
          <li key={point.key}>
            <Card elevation={SurfaceElevation.Raised} className="flex h-full flex-col gap-3 p-6">
              <span className="grid size-10 place-items-center rounded-tile bg-brand-soft text-brand">
                <Icon name={point.icon} className="size-5" />
              </span>
              <h3 className="text-base font-semibold tracking-tight text-ink">{point.title}</h3>
              <p className="text-sm leading-relaxed text-ink-soft">{point.description}</p>
            </Card>
          </li>
        ))}
      </ul>
    </Section>
  );
}
