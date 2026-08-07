import { AssetRole, BadgeTone } from "@/lib/enums";
import { marketAssets } from "@/content/protocol";
import { marketsPageContent } from "@/content/markets";
import { Badge } from "@/components/ui/Badge";
import { Container } from "@/components/ui/Container";
import { MarketActionsBar } from "@/components/markets/MarketActionsBar";
import { OraclePriceTicker } from "@/components/markets/OraclePriceTicker";

const roleLabels: Record<AssetRole, string> = {
  [AssetRole.Collateral]: "Collateral",
  [AssetRole.Borrowable]: "Borrow",
};

export function MarketHeader() {
  return (
    <div className="border-b border-line bg-surface">
      <Container className="flex flex-col gap-8 py-10 lg:flex-row lg:items-start lg:justify-between">
        <div className="flex flex-col items-start gap-4">
          <div className="flex flex-wrap items-center gap-3">
            <h1 className="text-3xl font-semibold tracking-tight text-ink sm:text-4xl">{marketsPageContent.title}</h1>
            <Badge tone={BadgeTone.Positive}>Active</Badge>
          </div>

          <p className="max-w-2xl text-pretty text-base leading-relaxed text-ink-soft">
            {marketsPageContent.description}
          </p>

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
        </div>

        <OraclePriceTicker />
      </Container>
    </div>
  );
}
