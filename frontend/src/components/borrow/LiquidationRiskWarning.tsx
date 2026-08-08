import { AppRoute, BadgeTone, ButtonSize, ButtonVariant, HealthTier, IconName } from "@/lib/enums";
import { Alert } from "@/components/ui/Alert";
import { Button } from "@/components/ui/Button";
import { healthExplanations } from "@/components/borrow/healthPresentation";

const warningTones: Record<HealthTier, BadgeTone> = {
  [HealthTier.Safe]: BadgeTone.Positive,
  [HealthTier.Caution]: BadgeTone.Caution,
  [HealthTier.AtRisk]: BadgeTone.Critical,
  [HealthTier.Liquidatable]: BadgeTone.Critical,
};

const warningTitles: Record<HealthTier, string> = {
  [HealthTier.Safe]: "",
  [HealthTier.Caution]: "Your buffer is getting thin",
  [HealthTier.AtRisk]: "Your position is close to liquidation",
  [HealthTier.Liquidatable]: "Your position can be liquidated now",
};

type LiquidationRiskWarningProps = {
  tier: HealthTier;
};

export function LiquidationRiskWarning({ tier }: LiquidationRiskWarningProps) {
  if (tier === HealthTier.Safe) {
    return null;
  }

  return (
    <Alert title={warningTitles[tier]} tone={warningTones[tier]} icon={IconName.Warning}>
      <div className="flex flex-col gap-3">
        <span>{healthExplanations[tier]}</span>
        <div className="flex flex-col gap-2 sm:flex-row">
          <Button href={AppRoute.Borrow} size={ButtonSize.Sm} variant={ButtonVariant.Secondary}>
            Add collateral
          </Button>
          <Button href={AppRoute.Borrow} size={ButtonSize.Sm} variant={ButtonVariant.Secondary}>
            Repay part of the loan
          </Button>
        </div>
      </div>
    </Alert>
  );
}
