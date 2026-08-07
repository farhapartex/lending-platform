import { MarketMetricKey, SectionId, SectionTone, SurfaceElevation, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { availableLiquidity, marketMetrics } from "@/content/protocol";
import { marketsPageContent } from "@/content/markets";
import { Card } from "@/components/ui/Card";
import { MetricRow } from "@/components/ui/MetricRow";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";
import { UtilizationBar } from "@/components/markets/UtilizationBar";

const emphasisedMetrics: MarketMetricKey[] = [MarketMetricKey.MaxLtv, MarketMetricKey.LiquidationThreshold];

export function MarketCard() {
  return (
    <Section id={SectionId.MarketSummary} tone={SectionTone.Surface} bordered>
      <SectionHeading
        sectionId={SectionId.MarketSummary}
        eyebrow="Terms"
        title={marketsPageContent.summaryTitle}
        description={marketsPageContent.summaryDescription}
      />

      <div className="mt-8 grid gap-6 lg:grid-cols-2 lg:gap-8">
        <Card elevation={SurfaceElevation.Flat} className="p-6 sm:p-7">
          <dl className="divide-y divide-line">
            {marketMetrics.map((metric) => (
              <MetricRow
                key={metric.key}
                label={metric.label}
                value={formatValue(metric.value, metric.format)}
                hint={metric.hint}
                emphasised={emphasisedMetrics.includes(metric.key)}
              />
            ))}
            <MetricRow
              label="Available to borrow now"
              value={formatValue(availableLiquidity, ValueFormat.UsdCompact)}
              hint="Deposited funds not currently lent out."
            />
          </dl>
        </Card>

        <Card elevation={SurfaceElevation.Flat} className="flex h-full flex-col justify-center p-6 sm:p-7">
          <UtilizationBar />
        </Card>
      </div>
    </Section>
  );
}
