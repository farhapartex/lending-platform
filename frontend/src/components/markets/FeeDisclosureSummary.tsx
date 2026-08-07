import { AppRoute, ButtonVariant, IconName, SectionId, SectionTone } from "@/lib/enums";
import { feeItems } from "@/content/protocol";
import { marketsPageContent } from "@/content/markets";
import { Button } from "@/components/ui/Button";
import { MetricRow } from "@/components/ui/MetricRow";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";

export function FeeDisclosureSummary() {
  return (
    <Section id={SectionId.MarketFees} tone={SectionTone.Muted} bordered>
      <SectionHeading
        sectionId={SectionId.MarketFees}
        eyebrow="Costs"
        title={marketsPageContent.feesTitle}
        description={marketsPageContent.feesDescription}
      />

      <dl className="mt-8 divide-y divide-line overflow-hidden rounded-card border border-line bg-surface px-5 sm:px-6">
        {feeItems.map((fee) => (
          <MetricRow key={fee.kind} label={fee.label} value={fee.value} hint={fee.description} emphasised />
        ))}
      </dl>

      <div className="mt-6">
        <Button href={AppRoute.LearnFees} variant={ButtonVariant.Secondary} trailingIcon={IconName.ArrowRight}>
          Full fee disclosure
        </Button>
      </div>
    </Section>
  );
}
