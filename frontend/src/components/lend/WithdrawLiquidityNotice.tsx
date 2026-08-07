import { AssetSymbol, BadgeTone, IconName } from "@/lib/enums";
import { formatTokenAmount } from "@/lib/token";
import { depositedBalance, lendAssetDecimals, poolAvailableLiquidity } from "@/content/lend";
import { Alert } from "@/components/ui/Alert";

type WithdrawLiquidityNoticeProps = {
  withdrawable: bigint;
};

export function WithdrawLiquidityNotice({ withdrawable }: WithdrawLiquidityNoticeProps) {
  const isLiquidityConstrained = poolAvailableLiquidity < depositedBalance;

  if (isLiquidityConstrained) {
    return (
      <Alert title="The pool cannot cover your full balance right now" tone={BadgeTone.Caution} icon={IconName.Warning}>
        Most of the pool is currently lent out, so you can withdraw up to{" "}
        {formatTokenAmount(withdrawable, lendAssetDecimals, 2)} {AssetSymbol.Usdc} at this moment. The rest becomes
        available as borrowers repay.
      </Alert>
    );
  }

  return (
    <Alert title="Your full balance is available" tone={BadgeTone.Neutral} icon={IconName.Info}>
      You can withdraw up to {formatTokenAmount(withdrawable, lendAssetDecimals, 2)} {AssetSymbol.Usdc} right now.
      Withdrawals are all or nothing, so a request larger than this is rejected rather than partly filled.
    </Alert>
  );
}
