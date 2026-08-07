import { AssetSymbol, BadgeTone, SurfaceElevation, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { formatTokenAmount } from "@/lib/token";
import { supplyApyRate } from "@/content/protocol";
import { lendAssetDecimals, lendPageContent, totalSupplied } from "@/content/lend";
import { Badge } from "@/components/ui/Badge";
import { Card } from "@/components/ui/Card";
import { PageHeader } from "@/components/ui/PageHeader";

export function LendHeader() {
  return (
    <PageHeader
      title={lendPageContent.title}
      description={lendPageContent.description}
      badge={<Badge tone={BadgeTone.Positive}>{AssetSymbol.Usdc}</Badge>}
      aside={
        <Card elevation={SurfaceElevation.Flat} className="w-full min-w-56 p-5 lg:w-auto">
          <dl className="flex flex-col gap-4">
            <div className="flex flex-col gap-1">
              <dt className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">Supply APY</dt>
              <dd className="text-2xl font-semibold tracking-tight text-ink tabular-nums">
                {formatValue(supplyApyRate, ValueFormat.Percent)}
              </dd>
            </div>
            <div className="flex flex-col gap-1 border-t border-line pt-4">
              <dt className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">Total supplied</dt>
              <dd className="text-lg font-semibold text-ink tabular-nums">
                {formatTokenAmount(totalSupplied, lendAssetDecimals, 0)} {AssetSymbol.Usdc}
              </dd>
            </div>
          </dl>
        </Card>
      }
    />
  );
}
