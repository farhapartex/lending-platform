import { BadgeTone, IconName, OracleStatus } from "@/lib/enums";
import { oracleReading } from "@/content/protocol";
import { Alert } from "@/components/ui/Alert";
import { Container } from "@/components/ui/Container";

export function PriceStalenessWarning() {
  if (oracleReading.status !== OracleStatus.Stale) {
    return null;
  }

  return (
    <Container className="pt-6">
      <Alert title="The price feed has not updated recently" tone={BadgeTone.Caution} icon={IconName.Warning}>
        Actions that depend on the price are rejected while it is stale, so you do not pay gas on a transaction that
        cannot succeed. This clears on its own once a fresh price arrives.
      </Alert>
    </Container>
  );
}
