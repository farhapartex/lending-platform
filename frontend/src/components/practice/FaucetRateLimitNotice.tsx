import { BadgeTone, IconName } from "@/lib/enums";
import { faucetRequestsPerDay } from "@/content/practice";
import { Alert } from "@/components/ui/Alert";

export function FaucetRateLimitNotice() {
  return (
    <Alert title="One request per day, per wallet" tone={BadgeTone.Neutral} icon={IconName.Info}>
      The faucet allows {faucetRequestsPerDay} request per asset every 24 hours, which keeps enough test tokens available
      for everyone else learning here. If you run out sooner, the wait is the only way to top up.
    </Alert>
  );
}
