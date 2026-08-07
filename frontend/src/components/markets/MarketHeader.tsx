import { AssetRole, BadgeTone } from "@/lib/enums";
import { marketAssets } from "@/content/protocol";
import { marketsPageContent } from "@/content/markets";
import { Badge } from "@/components/ui/Badge";
import { PageHeader } from "@/components/ui/PageHeader";
import { MarketActionsBar } from "@/components/markets/MarketActionsBar";
import { OraclePriceTicker } from "@/components/markets/OraclePriceTicker";

const roleLabels: Record<AssetRole, string> = {
  [AssetRole.Collateral]: "Collateral",
  [AssetRole.Borrowable]: "Borrow",
};

export function MarketHeader() {
  return (
    <PageHeader
      title={marketsPageContent.title}
      description={marketsPageContent.description}
      badge={<Badge tone={BadgeTone.Positive}>Active</Badge>}
      aside={<OraclePriceTicker />}
    >
      <ul className="flex flex-wrap gap-2">
        {marketAssets.map((asset) => (
          <li
            key={asset.symbol}
            className="flex items-center gap-2 rounded-pill border border-line bg-surface-muted px-3 py-1.5"
          >
            <span className="text-sm font-medium text-ink">{asset.symbol}</span>
            <span className="text-xs text-ink-faint">
              {asset.name} · {roleLabels[asset.role]}
            </span>
          </li>
        ))}
      </ul>

      <MarketActionsBar className="mt-2" />
    </PageHeader>
  );
}
