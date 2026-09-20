"use client";

import { ProductKey, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { bpsToRatio } from "@/lib/units";
import { debtAmountToNumber } from "@/lib/market";
import { products } from "@/content/landing";
import { useMarketData } from "@/hooks/useMarketData";
import { Skeleton } from "@/components/ui/Skeleton";
import { ProductCard } from "@/components/marketing/ProductCard";

function LendingFigures() {
  const { data } = useMarketData();

  if (data === undefined) {
    return <Skeleton className="h-16 w-full rounded-tile" />;
  }

  const figures = [
    { label: "Supply APY", value: formatValue(bpsToRatio(data.supplyAprBps), ValueFormat.Percent) },
    { label: "Borrow APR", value: formatValue(bpsToRatio(data.borrowAprBps), ValueFormat.Percent) },
    { label: "Deposited", value: formatValue(debtAmountToNumber(data.totalSupplied), ValueFormat.UsdCompact) },
  ];

  return (
    <dl className="grid grid-cols-3 gap-3 rounded-tile border border-line bg-surface-muted p-3">
      {figures.map((figure) => (
        <div key={figure.label} className="flex flex-col gap-0.5">
          <dt className="text-xs text-ink-faint">{figure.label}</dt>
          <dd className="text-sm font-semibold text-ink tabular-nums">{figure.value}</dd>
        </div>
      ))}
    </dl>
  );
}

export function ProductShowcase() {
  return (
    <div className="grid gap-6 lg:grid-cols-2">
      {products.map((product) => (
        <ProductCard
          key={product.key}
          product={product}
          accessory={product.key === ProductKey.Lending ? <LendingFigures /> : null}
        />
      ))}
    </div>
  );
}
