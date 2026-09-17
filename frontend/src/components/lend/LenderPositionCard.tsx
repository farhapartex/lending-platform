import { AppRoute, AssetSymbol, ButtonVariant, IconName, SurfaceElevation } from "@/lib/enums";
import { formatTokenAmount } from "@/lib/token";
import { lendAssetDecimals } from "@/content/lend";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { EmptyState } from "@/components/ui/EmptyState";

type LenderPositionCardProps = {
  depositedBalance: bigint;
};

export function LenderPositionCard({ depositedBalance }: LenderPositionCardProps) {
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

      <p className="border-t border-line pt-5 text-sm leading-relaxed text-ink-soft">
        This balance already includes the interest you have earned. The split between your principal and that
        interest needs your deposit history, which arrives with event indexing.
      </p>
    </Card>
  );
}
