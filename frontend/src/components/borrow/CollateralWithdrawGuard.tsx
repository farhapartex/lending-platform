import { AssetSymbol, BadgeTone, IconName } from "@/lib/enums";
import { formatTokenAmount } from "@/lib/token";
import { collateralDecimals } from "@/content/borrow";
import { Alert } from "@/components/ui/Alert";

type CollateralWithdrawGuardProps = {
  maxSafeWithdrawal: bigint;
  hasDebt: boolean;
};

export function CollateralWithdrawGuard({ maxSafeWithdrawal, hasDebt }: CollateralWithdrawGuardProps) {
  if (!hasDebt) {
    return (
      <Alert title="All of your collateral is free to withdraw" tone={BadgeTone.Positive} icon={IconName.Check}>
        You have no outstanding loan, so nothing restricts how much you take back.
      </Alert>
    );
  }

  if (maxSafeWithdrawal <= 0n) {
    return (
      <Alert title="No collateral can be withdrawn right now" tone={BadgeTone.Critical} icon={IconName.Warning}>
        Your loan already uses your full borrowing power. Repay part of it first, and the collateral it was supporting
        becomes available to withdraw.
      </Alert>
    );
  }

  return (
    <Alert title="Some collateral is locked by your loan" tone={BadgeTone.Neutral} icon={IconName.Info}>
      You can take back up to {formatTokenAmount(maxSafeWithdrawal, collateralDecimals, 4)} {AssetSymbol.Weth} and stay
      within your borrowing limit. The rest is backing what you have already borrowed.
    </Alert>
  );
}
