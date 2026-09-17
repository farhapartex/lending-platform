"use client";

import { AssetRole, BadgeTone, MarketMetricKey, SurfaceElevation, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { bpsToRatio } from "@/lib/units";
import { debtAmountToNumber } from "@/lib/market";
import { marketAssets } from "@/content/protocol";
import { useMarketData } from "@/hooks/useMarketData";
import { Badge } from "@/components/ui/Badge";
import { Card } from "@/components/ui/Card";
import { Skeleton } from "@/components/ui/Skeleton";

const roleLabels: Record<AssetRole, string> = {
  [AssetRole.Collateral]: "Collateral",
  [AssetRole.Borrowable]: "Borrow",
};

export function MarketSnapshotCard() {
  const { data } = useMarketData();

  const metrics =
    data === undefined
      ? []
      : [
          {
            key: MarketMetricKey.SupplyApy,
            label: "Supply APY",
            value: bpsToRatio(data.supplyAprBps),
            hint: "Earned by lenders, paid by borrowers",
          },
          {
            key: MarketMetricKey.BorrowApr,
            label: "Borrow APR",
            value: bpsToRatio(data.borrowAprBps),
            hint: "Rises as the pool gets more utilized",
          },
          {
            key: MarketMetricKey.MaxLtv,
            label: "Max borrow",
            value: bpsToRatio(data.maxLtvBps),
            hint: "Of your collateral value",
          },
          {
            key: MarketMetricKey.LiquidationThreshold,
            label: "Liquidation at",
            value: bpsToRatio(data.liquidationThresholdBps),
            hint: "Deliberate buffer above max borrow",
          },
        ];

  return (
    <Card elevation={SurfaceElevation.Lifted} className="w-full p-6 sm:p-7">
      <div className="flex items-start justify-between gap-4">
        <div className="flex flex-col gap-1">
          <span className="text-sm text-ink-soft">Live market</span>
          <span className="text-lg font-semibold tracking-tight text-ink">
            {marketAssets.map((asset) => asset.symbol).join(" / ")}
          </span>
        </div>
        <Badge tone={BadgeTone.Positive}>Active</Badge>
      </div>

      <ul className="mt-5 flex flex-wrap gap-2">
        {marketAssets.map((asset) => (
          <li
            key={asset.symbol}
            className="flex items-center gap-2 rounded-pill border border-line bg-surface-muted px-3 py-1.5"
          >
            <span className="text-sm font-medium text-ink">{asset.symbol}</span>
            <span className="text-xs text-ink-faint">{roleLabels[asset.role]}</span>
          </li>
        ))}
      </ul>

      {data === undefined ? (
        <div className="mt-6 flex flex-col gap-4">
          <Skeleton className="h-16 w-full" />
          <Skeleton className="h-16 w-full" />
        </div>
      ) : (
        <>
          <dl className="mt-6 grid grid-cols-2 gap-x-4 gap-y-5">
            {metrics.map((metric) => (
              <div key={metric.key} className="flex flex-col gap-1">
                <dt className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">{metric.label}</dt>
                <dd className="text-xl font-semibold tracking-tight text-ink tabular-nums">
                  {formatValue(metric.value, ValueFormat.Percent)}
                </dd>
                <p className="text-xs leading-relaxed text-ink-soft">{metric.hint}</p>
              </div>
            ))}
          </dl>

          <div className="mt-6 flex items-center justify-between gap-3 rounded-tile border border-line bg-surface-muted px-4 py-3">
            <span className="text-sm text-ink-soft">Available to borrow now</span>
            <span className="text-sm font-semibold text-ink tabular-nums">
              {formatValue(debtAmountToNumber(data.availableLiquidity), ValueFormat.UsdCompact)}
            </span>
          </div>
        </>
      )}
    </Card>
  );
}
