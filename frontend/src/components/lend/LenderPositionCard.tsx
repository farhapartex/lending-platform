import { AppRoute, AssetSymbol, ButtonVariant, IconName, SurfaceElevation } from "@/lib/enums";
import { formatTokenAmount } from "@/lib/token";
import { accruedInterest, depositedBalance, depositedPrincipal, lendAssetDecimals } from "@/content/lend";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { EmptyState } from "@/components/ui/EmptyState";
import { LiveInterestCounter } from "@/components/lend/LiveInterestCounter";

export function LenderPositionCard() {
  if (depositedBalance <= 0n) {
    return (
      <EmptyState
        title="You have not deposited yet"
        description="Once you deposit USDC, your balance and the interest it earns will appear here."
        icon={IconName.Coins}
        action={
          <Button href={AppRoute.Lend} variant={ButtonVariant.Subtle} trailingIcon={IconName.ArrowRight}>
            Make a deposit
          </Button>
        }
      />
    );
  }

  return (
    <Card elevation={SurfaceElevation.Raised} className="flex flex-col gap-6 p-6 sm:p-7">
      <div className="flex flex-col gap-1.5">
        <span className="text-sm text-ink-soft">Current balance</span>
        <span className="text-3xl font-semibold tracking-tight text-ink tabular-nums">
          {formatTokenAmount(depositedBalance, lendAssetDecimals, 2)} {AssetSymbol.Usdc}
        </span>
      </div>

      <dl className="grid gap-5 border-t border-line pt-5 sm:grid-cols-2">
        <div className="flex flex-col gap-1">
          <dt className="text-sm text-ink-soft">Principal deposited</dt>
          <dd className="text-lg font-semibold text-ink tabular-nums">
            {formatTokenAmount(depositedPrincipal, lendAssetDecimals, 2)} {AssetSymbol.Usdc}
          </dd>
        </div>

        <div className="flex flex-col gap-1">
          <dt className="text-sm text-ink-soft">Interest earned</dt>
          <dd className="text-lg font-semibold text-mint-ink">
            <LiveInterestCounter baseInterest={accruedInterest} balance={depositedBalance} />
          </dd>
        </div>
      </dl>
    </Card>
  );
}
