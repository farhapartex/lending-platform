"use client";

import type { ReactNode } from "react";
import { ButtonSize, IconName, WalletGatePurpose, WalletStatus } from "@/lib/enums";
import { appChain } from "@/lib/chain";
import { useWalletState } from "@/hooks/useWalletState";
import { EmptyState } from "@/components/ui/EmptyState";
import { Skeleton } from "@/components/ui/Skeleton";
import { WalletConnectButton } from "@/components/app/WalletConnectButton";

const disconnectedCopy: Record<WalletGatePurpose, { title: string; description: string }> = {
  [WalletGatePurpose.Action]: {
    title: "Connect a wallet to continue",
    description:
      "Everything on this page is readable without connecting. A wallet is only needed to move funds.",
  },
  [WalletGatePurpose.PersonalData]: {
    title: "Connect a wallet to see your position",
    description:
      "This page shows your own deposits, loan and safety score, so there is nothing to display until a wallet is connected.",
  },
  [WalletGatePurpose.Liquidate]: {
    title: "Connect a wallet to liquidate",
    description:
      "The list above is public, but repaying someone's loan and claiming their collateral needs a wallet with USDC in it.",
  },
  [WalletGatePurpose.Faucet]: {
    title: "Connect a wallet to get test funds",
    description: "The faucet needs an address to send the test tokens to.",
  },
};

type WalletGateProps = {
  children: ReactNode;
  purpose?: WalletGatePurpose;
  skeletonClassName?: string;
};

export function WalletGate({
  children,
  purpose = WalletGatePurpose.Action,
  skeletonClassName = "h-40 rounded-card",
}: WalletGateProps) {
  const { status, isSettling } = useWalletState();

  if (isSettling) {
    return <Skeleton className={skeletonClassName} />;
  }

  if (status === WalletStatus.Connected) {
    return <>{children}</>;
  }

  if (status === WalletStatus.WrongNetwork) {
    return (
      <EmptyState
        title={`Switch to ${appChain.name} to continue`}
        description="Your wallet is connected but pointed at a different network than this market runs on."
        icon={IconName.Warning}
      />
    );
  }

  if (status === WalletStatus.Connecting) {
    return (
      <EmptyState
        title="Waiting for your wallet"
        description="Approve the connection request in your wallet to carry on."
        icon={IconName.Wallet}
      />
    );
  }

  const copy = disconnectedCopy[purpose];

  return (
    <EmptyState
      title={copy.title}
      description={copy.description}
      icon={IconName.Wallet}
      action={<WalletConnectButton size={ButtonSize.Md} />}
    />
  );
}
