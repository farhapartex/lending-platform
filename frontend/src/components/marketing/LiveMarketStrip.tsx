"use client";

import { AssetSymbol, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { bpsToRatio } from "@/lib/units";
import { debtAmountToNumber } from "@/lib/market";
import { priceToNumber, useOraclePrice } from "@/hooks/useOraclePrice";
import { useMarketData } from "@/hooks/useMarketData";
import { Skeleton } from "@/components/ui/Skeleton";

export function LiveMarketStrip() {
  const { data } = useMarketData();
  const price = useOraclePrice();

  if (data === undefined) {
    return (
      <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-5">
        {[0, 1, 2, 3, 4].map((slot) => (
          <Skeleton key={slot} className="h-20 w-full rounded-card" />
        ))}
      </div>
    );
  }

  const figures = [
    { label: "Supply APY", value: formatValue(bpsToRatio(data.supplyAprBps), ValueFormat.Percent) },
    { label: "Borrow APR", value: formatValue(bpsToRatio(data.borrowAprBps), ValueFormat.Percent) },
    { label: "Total deposited", value: formatValue(debtAmountToNumber(data.totalSupplied), ValueFormat.UsdCompact) },
    {
      label: "Available to borrow",
      value: formatValue(debtAmountToNumber(data.availableLiquidity), ValueFormat.UsdCompact),
    },
    {
      label: `${AssetSymbol.Weth} price`,
      value: price.data === undefined ? "—" : formatValue(priceToNumber(price.data), ValueFormat.UsdPrice),
    },
  ];

  return (
    <dl className="grid gap-4 sm:grid-cols-2 lg:grid-cols-5">
      {figures.map((figure) => (
        <div key={figure.label} className="flex flex-col gap-1 rounded-card border border-line bg-surface p-4">
          <dt className="text-xs font-medium uppercase tracking-[0.06em] text-ink-faint">{figure.label}</dt>
          <dd className="text-xl font-semibold tracking-tight text-ink tabular-nums">{figure.value}</dd>
        </div>
      ))}
    </dl>
  );
}
