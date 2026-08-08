import { ActivityKind, BadgeTone, IconName } from "@/lib/enums";
import { Badge } from "@/components/ui/Badge";

const labels: Record<ActivityKind, string> = {
  [ActivityKind.Deposit]: "Deposit",
  [ActivityKind.Withdraw]: "Withdrawal",
  [ActivityKind.Borrow]: "Borrow",
  [ActivityKind.Repay]: "Repayment",
  [ActivityKind.CollateralAdded]: "Collateral added",
  [ActivityKind.CollateralWithdrawn]: "Collateral withdrawn",
  [ActivityKind.Liquidation]: "Liquidation",
};

const tones: Record<ActivityKind, BadgeTone> = {
  [ActivityKind.Deposit]: BadgeTone.Positive,
  [ActivityKind.Withdraw]: BadgeTone.Neutral,
  [ActivityKind.Borrow]: BadgeTone.Brand,
  [ActivityKind.Repay]: BadgeTone.Positive,
  [ActivityKind.CollateralAdded]: BadgeTone.Positive,
  [ActivityKind.CollateralWithdrawn]: BadgeTone.Neutral,
  [ActivityKind.Liquidation]: BadgeTone.Critical,
};

const icons: Record<ActivityKind, IconName> = {
  [ActivityKind.Deposit]: IconName.Coins,
  [ActivityKind.Withdraw]: IconName.Coins,
  [ActivityKind.Borrow]: IconName.Wallet,
  [ActivityKind.Repay]: IconName.Check,
  [ActivityKind.CollateralAdded]: IconName.ShieldCheck,
  [ActivityKind.CollateralWithdrawn]: IconName.ShieldCheck,
  [ActivityKind.Liquidation]: IconName.Warning,
};

type TxTypeBadgeProps = {
  kind: ActivityKind;
};

export function TxTypeBadge({ kind }: TxTypeBadgeProps) {
  return (
    <Badge tone={tones[kind]} icon={icons[kind]}>
      {labels[kind]}
    </Badge>
  );
}
