import type { ReactNode } from "react";
import { IconName, WalletStatus } from "@/lib/enums";
import { walletStatus } from "@/content/wallet";
import { EmptyState } from "@/components/ui/EmptyState";
import { WalletConnectButton } from "@/components/app/WalletConnectButton";

const gateTitles: Record<WalletStatus, string> = {
  [WalletStatus.Disconnected]: "Connect a wallet to continue",
  [WalletStatus.Connecting]: "Waiting for your wallet",
  [WalletStatus.Connected]: "",
  [WalletStatus.WrongNetwork]: "Switch to the right network",
};

const gateDescriptions: Record<WalletStatus, string> = {
  [WalletStatus.Disconnected]:
    "You can read every number on this page without connecting. A wallet is only needed to move funds.",
  [WalletStatus.Connecting]: "Approve the connection request in your wallet to carry on.",
  [WalletStatus.Connected]: "",
  [WalletStatus.WrongNetwork]: "This market lives on a different network than the one your wallet is currently using.",
};

type WalletGateProps = {
  children: ReactNode;
};

export function WalletGate({ children }: WalletGateProps) {
  if (walletStatus === WalletStatus.Connected) {
    return <>{children}</>;
  }

  return (
    <EmptyState
      title={gateTitles[walletStatus]}
      description={gateDescriptions[walletStatus]}
      icon={IconName.Wallet}
      action={<WalletConnectButton />}
    />
  );
}
