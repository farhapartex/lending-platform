import { AssetRole, BadgeTone, SurfaceElevation, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { availableLiquidity, marketAssets, marketMetrics } from "@/content/landing";
import { Badge } from "@/components/ui/Badge";
import { Card } from "@/components/ui/Card";

const roleLabels: Record<AssetRole, string> = {
  [AssetRole.Collateral]: "Collateral",
  [AssetRole.Borrowable]: "Borrow",
};

export function MarketSnapshotCard() {
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

      <dl className="mt-6 grid grid-cols-2 gap-x-4 gap-y-5">
        {marketMetrics.map((metric) => (
          <div key={metric.key} className="flex flex-col gap-1">
            <dt className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">{metric.label}</dt>
            <dd className="text-xl font-semibold tracking-tight text-ink tabular-nums">
              {formatValue(metric.value, metric.format)}
            </dd>
            <p className="text-xs leading-relaxed text-ink-soft">{metric.hint}</p>
          </div>
        ))}
      </dl>

      <div className="mt-6 flex items-center justify-between gap-3 rounded-tile border border-line bg-surface-muted px-4 py-3">
        <span className="text-sm text-ink-soft">Available to borrow now</span>
        <span className="text-sm font-semibold text-ink tabular-nums">
          {formatValue(availableLiquidity, ValueFormat.UsdCompact)}
        </span>
      </div>
    </Card>
  );
}
