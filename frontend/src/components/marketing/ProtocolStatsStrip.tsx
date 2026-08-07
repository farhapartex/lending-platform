import { DataStatus, SectionId, SectionSpacing, SectionTone } from "@/lib/enums";
import { protocolStats, protocolStatsStatus } from "@/content/protocol";
import { Section } from "@/components/ui/Section";
import { StatTile, StatTileSkeleton } from "@/components/ui/StatTile";

export function ProtocolStatsStrip() {
  if (protocolStatsStatus === DataStatus.Unavailable) {
    return null;
  }

  return (
    <Section
      id={SectionId.ProtocolStats}
      tone={SectionTone.Surface}
      spacing={SectionSpacing.Compact}
      bordered
      className="border-t-0"
    >
      <h2 id={`${SectionId.ProtocolStats}-heading`} className="sr-only">
        Protocol totals
      </h2>
      <dl className="grid gap-8 sm:grid-cols-3">
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
    </Section>
  );
}
