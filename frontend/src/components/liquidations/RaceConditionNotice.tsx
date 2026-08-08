import { BadgeTone, IconName } from "@/lib/enums";
import { Alert } from "@/components/ui/Alert";

export function RaceConditionNotice() {
  return (
    <Alert title="Another liquidator may get there first" tone={BadgeTone.Caution} icon={IconName.Warning}>
      Only one liquidation can succeed per position. If someone else is included in a block ahead of you, your
      transaction fails without moving any funds, and you pay only the network gas.
    </Alert>
  );
}
