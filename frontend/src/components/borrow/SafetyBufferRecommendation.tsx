import { AssetSymbol, BadgeTone, IconName, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { basisPoints } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import { debtDecimals, recommendedLtvBps } from "@/content/borrow";
import { Alert } from "@/components/ui/Alert";

type SafetyBufferRecommendationProps = {
  recommendedCapacity: bigint;
  isExceeded: boolean;
};

export function SafetyBufferRecommendation({ recommendedCapacity, isExceeded }: SafetyBufferRecommendationProps) {
  const recommendedLtv = formatValue(Number(recommendedLtvBps) / Number(basisPoints), ValueFormat.Percent);

  if (isExceeded) {
    return (
      <Alert title="You are past our recommended buffer" tone={BadgeTone.Caution} icon={IconName.Info}>
        We suggest keeping your loan under {recommendedLtv} of your collateral value so a normal price swing cannot put
        you at risk. You are above that, which is allowed but leaves less margin than we would advise.
      </Alert>
    );
  }

  return (
    <Alert title="Recommended borrowing room" tone={BadgeTone.Brand} icon={IconName.ShieldCheck}>
      Staying under {recommendedLtv} of your collateral value keeps a healthy buffer. That works out to{" "}
      {formatTokenAmount(recommendedCapacity, debtDecimals, 2)} {AssetSymbol.Usdc} for your current collateral.
    </Alert>
  );
}
