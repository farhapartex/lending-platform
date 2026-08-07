import { DataStatus, SectionId, SectionSpacing, SectionTone } from "@/lib/enums";
import { protocolStats, protocolStatsStatus } from "@/content/protocol";
import { marketsPageContent } from "@/content/markets";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";
import { StatTile, StatTileSkeleton } from "@/components/ui/StatTile";

export function ProtocolStatsGrid() {
  return (
    <Section id={SectionId.ProtocolTotals} tone={SectionTone.Canvas} spacing={SectionSpacing.Regular}>
      <SectionHeading
        sectionId={SectionId.ProtocolTotals}
        eyebrow="Overview"
        title={marketsPageContent.totalsTitle}
        description={marketsPageContent.totalsDescription}
      />

      {protocolStatsStatus === DataStatus.Unavailable ? null : (
        <dl className="mt-8 grid gap-6 rounded-card border border-line bg-surface p-6 sm:grid-cols-3 sm:gap-8 sm:p-8">
          {protocolStatsStatus === DataStatus.Loading
            ? protocolStats.map((stat) => <StatTileSkeleton key={stat.key} />)
            : protocolStats.map((stat) => (
                <StatTile
                  key={stat.key}
                  label={stat.label}
                  value={stat.value}
                  format={stat.format}
                  trend={stat.trend}
                  trendLabel={stat.trendLabel}
                />
              ))}
        </dl>
      )}
    </Section>
  );
}
