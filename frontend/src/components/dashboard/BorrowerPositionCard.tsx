import { AppRoute, AssetSymbol, ButtonSize, ButtonVariant, HealthTier, IconName, SurfaceElevation, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { formatHealthFactor, scaledValueToUsd } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import { collateralDecimals, collateralDeposited, debtDecimals, debtOutstanding } from "@/content/borrow";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { EmptyState } from "@/components/ui/EmptyState";
import { HealthBadge } from "@/components/borrow/HealthBadge";

type BorrowerPositionCardProps = {
  factorBps: bigint | null;
  tier: HealthTier;
  collateralValueScaled: bigint;
};

export function BorrowerPositionCard({ factorBps, tier, collateralValueScaled }: BorrowerPositionCardProps) {
  if (collateralDeposited <= 0n && debtOutstanding <= 0n) {
    return (
      <EmptyState
        title="No loan yet"
        description="Post WETH as collateral and you can borrow against it without selling."
        icon={IconName.Wallet}
        action={
          <Button href={AppRoute.Borrow} variant={ButtonVariant.Subtle} trailingIcon={IconName.ArrowRight}>
            Open a loan
          </Button>
        }
      />
    );
  }

  return (
    <Card elevation={SurfaceElevation.Raised} className="flex h-full flex-col gap-5 p-6 sm:p-7">
      <div className="flex flex-wrap items-start justify-between gap-3">
        <div className="flex flex-col gap-1.5">
          <span className="text-sm text-ink-soft">Outstanding loan</span>
          <span className="text-2xl font-semibold tracking-tight text-ink tabular-nums">
            {formatTokenAmount(debtOutstanding, debtDecimals, 2)} {AssetSymbol.Usdc}
          </span>
        </div>
        <HealthBadge tier={tier} />
      </div>

      <dl className="grid gap-4 border-t border-line pt-5 sm:grid-cols-2">
        <div className="flex flex-col gap-1">
          <dt className="text-sm text-ink-soft">Collateral</dt>
          <dd className="text-base font-semibold text-ink tabular-nums">
            {formatTokenAmount(collateralDeposited, collateralDecimals, 4)} {AssetSymbol.Weth}
          </dd>
          <p className="text-xs text-ink-faint tabular-nums">
            {formatValue(scaledValueToUsd(collateralValueScaled), ValueFormat.UsdPrice)}
          </p>
        </div>

        <div className="flex flex-col gap-1">
          <dt className="text-sm text-ink-soft">Health factor</dt>
          <dd className="text-base font-semibold text-ink tabular-nums">{formatHealthFactor(factorBps)}</dd>
          <p className="text-xs text-ink-faint">Liquidation below 1.00</p>
        </div>
      </dl>

      <div className="mt-auto flex flex-col gap-2 sm:flex-row">
        <Button href={AppRoute.Borrow} size={ButtonSize.Sm} variant={ButtonVariant.Secondary}>
          Manage loan
        </Button>
      </div>
    </Card>
  );
}
