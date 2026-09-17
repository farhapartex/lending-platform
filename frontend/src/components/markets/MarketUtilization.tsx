"use client";

import { bpsToRatio } from "@/lib/units";
import { useMarketData } from "@/hooks/useMarketData";
import { Skeleton } from "@/components/ui/Skeleton";
import { UtilizationBar } from "@/components/markets/UtilizationBar";

type MarketUtilizationProps = {
  className?: string;
};

export function MarketUtilization({ className }: MarketUtilizationProps) {
  const { data } = useMarketData();

  if (data === undefined) {
    return <Skeleton className="h-24 w-full" />;
  }

  return (
    <UtilizationBar
      current={bpsToRatio(data.utilizationBps)}
      kink={bpsToRatio(data.kinkUtilizationBps)}
      className={className}
    />
  );
}
