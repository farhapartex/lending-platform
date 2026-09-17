"use client";

import { DataStatus, ProtocolStatKey, SectionId, SectionSpacing, SectionTone, ValueFormat } from "@/lib/enums";
import { bpsToRatio } from "@/lib/units";
import { dataStatusFrom, debtAmountToNumber } from "@/lib/market";
import { useMarketData } from "@/hooks/useMarketData";
import { Section } from "@/components/ui/Section";
import { StatTile, StatTileSkeleton } from "@/components/ui/StatTile";

const placeholderKeys = [ProtocolStatKey.TotalDeposited, ProtocolStatKey.TotalBorrowed, ProtocolStatKey.Utilization];

export function ProtocolStatsStrip() {
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
        {status === DataStatus.Loading
          ? placeholderKeys.map((key) => <StatTileSkeleton key={key} />)
          : stats.map((stat) => (
              <StatTile key={stat.key} label={stat.label} value={stat.value} format={stat.format} />
            ))}
      </dl>
    </Section>
  );
}
