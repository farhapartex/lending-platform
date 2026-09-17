"use client";

import { DataStatus, ProtocolStatKey, SectionId, SectionSpacing, SectionTone, ValueFormat } from "@/lib/enums";
import { bpsToRatio } from "@/lib/units";
import { debtAmountToNumber, dataStatusFrom } from "@/lib/market";
import { marketsPageContent } from "@/content/markets";
import { useMarketData } from "@/hooks/useMarketData";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";
import { StatTile, StatTileSkeleton } from "@/components/ui/StatTile";

const placeholderKeys = [ProtocolStatKey.TotalDeposited, ProtocolStatKey.TotalBorrowed, ProtocolStatKey.Utilization];

export function ProtocolStatsGrid() {
  const { data, isLoading, isError } = useMarketData();
  const status = dataStatusFrom({ isLoading, isError, hasData: data !== undefined });

  if (status === DataStatus.Unavailable) {
    return null;
  }

  const stats =
    data === undefined
      ? []
      : [
          {
            key: ProtocolStatKey.TotalDeposited,
            label: "Total deposited",
            value: debtAmountToNumber(data.totalSupplied),
            format: ValueFormat.UsdCompact,
          },
          {
            key: ProtocolStatKey.TotalBorrowed,
            label: "Total borrowed",
            value: debtAmountToNumber(data.totalBorrowed),
            format: ValueFormat.UsdCompact,
          },
          {
            key: ProtocolStatKey.Utilization,
            label: "Pool utilization",
            value: bpsToRatio(data.utilizationBps),
            format: ValueFormat.Percent,
          },
        ];

  return (
    <Section id={SectionId.ProtocolTotals} tone={SectionTone.Canvas} spacing={SectionSpacing.Regular}>
      <SectionHeading
        sectionId={SectionId.ProtocolTotals}
        eyebrow="Overview"
        title={marketsPageContent.totalsTitle}
        description={marketsPageContent.totalsDescription}
      />

      <dl className="mt-8 grid gap-6 rounded-card border border-line bg-surface p-6 sm:grid-cols-3 sm:gap-8 sm:p-8">
        {status === DataStatus.Loading
          ? placeholderKeys.map((key) => <StatTileSkeleton key={key} />)
          : stats.map((stat) => (
              <StatTile key={stat.key} label={stat.label} value={stat.value} format={stat.format} />
            ))}
      </dl>
    </Section>
  );
}
