import { AppRoute, BadgeTone, IconName, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { basisPoints } from "@/lib/health";
import { liquidationBonusBps } from "@/content/borrow";
import { Alert } from "@/components/ui/Alert";
import { TextLink } from "@/components/ui/TextLink";

export function FullLiquidationNotice() {
  return (
    <Alert title="Liquidation closes your whole position" tone={BadgeTone.Caution} icon={IconName.Warning}>
      <div className="flex flex-col gap-2">
        <span>
          If your health factor falls below 1.00, anyone can repay your loan and take your collateral, plus a{" "}
          {formatValue(Number(liquidationBonusBps) / Number(basisPoints), ValueFormat.Percent)} bonus. At this stage the
          entire position is closed rather than just the unsafe part, so even a brief price dip can end the whole loan.
        </span>
        <TextLink href={AppRoute.LearnLiquidation} trailingIcon={IconName.ArrowRight}>
          Read exactly how liquidation works
        </TextLink>
      </div>
    </Alert>
  );
}
