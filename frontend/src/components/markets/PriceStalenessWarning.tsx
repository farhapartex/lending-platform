"use client";

import { BadgeTone, IconName, OracleStatus } from "@/lib/enums";
import { useOraclePrice } from "@/hooks/useOraclePrice";
import { Alert } from "@/components/ui/Alert";

export function PriceStalenessWarning() {
  const { status } = useOraclePrice();

  if (status !== OracleStatus.Stale) {
    return null;
  }

  return (
    <Alert title="The price feed has not updated recently" tone={BadgeTone.Caution} icon={IconName.Warning}>
      Actions that depend on the price are rejected while it is stale, so you do not pay gas on a transaction that
      cannot succeed. This clears on its own once a fresh price arrives.
    </Alert>
  );
}
