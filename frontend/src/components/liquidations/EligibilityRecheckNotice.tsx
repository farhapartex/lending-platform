import { BadgeTone, IconName } from "@/lib/enums";
import { Alert } from "@/components/ui/Alert";

export function EligibilityRecheckNotice() {
  return (
    <Alert title="Eligibility is checked again on-chain" tone={BadgeTone.Neutral} icon={IconName.Info}>
      The contract re-verifies this position at the moment your transaction runs, not when you opened this dialog. If the
      WETH price recovers in between, the liquidation is rejected cleanly and nothing moves.
    </Alert>
  );
}
