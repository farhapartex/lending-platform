"use client";

import { useState } from "react";
import { ButtonSize, ButtonVariant, IconName, WalletStatus } from "@/lib/enums";
import { useWalletState } from "@/hooks/useWalletState";
import { Button } from "@/components/ui/Button";
import { Skeleton } from "@/components/ui/Skeleton";
import { AccountMenu } from "@/components/app/AccountMenu";
import { WalletProviderModal } from "@/components/app/WalletProviderModal";

type WalletConnectButtonProps = {
  size?: ButtonSize;
  fullWidth?: boolean;
};

export function WalletConnectButton({ size = ButtonSize.Sm, fullWidth = false }: WalletConnectButtonProps) {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const { status, address, chainId, isSettling } = useWalletState();

  if (isSettling) {
    return <Skeleton className="h-9 w-36 rounded-pill" />;
  }

  if (status === WalletStatus.Connected || status === WalletStatus.WrongNetwork) {
    if (address !== undefined) {
      return <AccountMenu address={address} chainId={chainId} />;
    }
  }

  return (
    <>
      <Button
        variant={status === WalletStatus.WrongNetwork ? ButtonVariant.Primary : ButtonVariant.Secondary}
        size={size}
        fullWidth={fullWidth}
        leadingIcon={IconName.Wallet}
        disabled={status === WalletStatus.Connecting}
        onClick={() => setIsModalOpen(true)}
      >
        {status === WalletStatus.Connecting ? "Connecting" : "Connect wallet"}
      </Button>

      <WalletProviderModal open={isModalOpen} onClose={() => setIsModalOpen(false)} />
    </>
  );
}
