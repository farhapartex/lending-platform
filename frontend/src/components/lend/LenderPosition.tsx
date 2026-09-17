"use client";

import { useAccountData } from "@/hooks/useAccountData";
import { useWalletState } from "@/hooks/useWalletState";
import { Skeleton } from "@/components/ui/Skeleton";
import { LenderPositionCard } from "@/components/lend/LenderPositionCard";

export function LenderPosition() {
  const { address } = useWalletState();
  const { data } = useAccountData(address);

  if (data === undefined) {
    return <Skeleton className="h-48 w-full" />;
  }

  return <LenderPositionCard depositedBalance={data.supplyAssets} />;
}
