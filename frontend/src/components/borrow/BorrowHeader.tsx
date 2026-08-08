import { AssetSymbol, BadgeTone, HealthTier, SurfaceElevation, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { scaledValueToUsd } from "@/lib/health";
import { Badge } from "@/components/ui/Badge";
import { Card } from "@/components/ui/Card";
import { PageHeader } from "@/components/ui/PageHeader";
import { borrowAprDisplayRate, borrowPageContent } from "@/content/borrow";
import { HealthBadge } from "@/components/borrow/HealthBadge";

type BorrowHeaderProps = {
  tier: HealthTier;
  collateralValueScaled: bigint;
};

export function BorrowHeader({ tier, collateralValueScaled }: BorrowHeaderProps) {
  return (
    <PageHeader
      title={borrowPageContent.title}
      description={borrowPageContent.description}
      badge={<Badge tone={BadgeTone.Neutral}>{AssetSymbol.Weth} collateral</Badge>}
      aside={
        <Card elevation={SurfaceElevation.Flat} className="w-full min-w-56 p-5 lg:w-auto">
          <dl className="flex flex-col gap-4">
            <div className="flex flex-col gap-1.5">
              <dt className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">Position health</dt>
              <dd>
                <HealthBadge tier={tier} />
              </dd>
            </div>
            <div className="flex flex-col gap-1 border-t border-line pt-4">
              <dt className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">Collateral value</dt>
              <dd className="text-lg font-semibold text-ink tabular-nums">
                {formatValue(scaledValueToUsd(collateralValueScaled), ValueFormat.UsdPrice)}
              </dd>
            </div>
            <div className="flex flex-col gap-1 border-t border-line pt-4">
              <dt className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">Borrow APR</dt>
              <dd className="text-lg font-semibold text-ink tabular-nums">
                {formatValue(borrowAprDisplayRate, ValueFormat.Percent)}
              </dd>
            </div>
          </dl>
        </Card>
      }
    />
  );
}
