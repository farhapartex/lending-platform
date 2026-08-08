import { AssetSymbol, BadgeTone, IconName } from "@/lib/enums";
import { formatTokenAmount } from "@/lib/token";
import { debtDecimals } from "@/content/borrow";
import { Alert } from "@/components/ui/Alert";

type InsufficientLiquidityNoticeProps = {
  availableLiquidity: bigint;
};

export function InsufficientLiquidityNotice({ availableLiquidity }: InsufficientLiquidityNoticeProps) {
  return (
    <Alert title="Limited by pool liquidity, not by your collateral" tone={BadgeTone.Caution} icon={IconName.Warning}>
      Your collateral supports a larger loan, but only {formatTokenAmount(availableLiquidity, debtDecimals, 2)}{" "}
      {AssetSymbol.Usdc} is sitting in the pool right now. More becomes available as other borrowers repay.
    </Alert>
  );
}
