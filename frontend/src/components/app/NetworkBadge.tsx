"use client";

import { BadgeTone, WalletStatus } from "@/lib/enums";
import { chainDisplayName, isTestnetChain } from "@/lib/chain";
import { useWalletState } from "@/hooks/useWalletState";
import { Badge } from "@/components/ui/Badge";

export function NetworkBadge() {
  const { status, chainId } = useWalletState();

  if (status === WalletStatus.Disconnected || status === WalletStatus.Connecting) {
    return null;
  }

  if (status === WalletStatus.WrongNetwork) {
    return <Badge tone={BadgeTone.Critical}>{chainDisplayName(chainId)}</Badge>;
  }

  return (
    <Badge tone={isTestnetChain(chainId) ? BadgeTone.Caution : BadgeTone.Neutral}>
      {chainDisplayName(chainId)}
    </Badge>
  );
}
