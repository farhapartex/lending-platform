"use client";

import { useSwitchChain } from "wagmi";
import { ButtonVariant, IconName, SurfaceElevation, WalletStatus } from "@/lib/enums";
import { appChain, appChainId, chainDisplayName, isTestnetChain } from "@/lib/chain";
import { useWalletState } from "@/hooks/useWalletState";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { Icon } from "@/components/ui/Icon";

export function NetworkSwitchCard() {
  const { status, chainId } = useWalletState();
  const { switchChain, isPending } = useSwitchChain();

  const appIsTestnet = isTestnetChain(appChainId);
  const walletOnAppChain = status === WalletStatus.Connected;
  const isDisconnected = status === WalletStatus.Disconnected || status === WalletStatus.Connecting;

  return (
    <Card elevation={SurfaceElevation.Flat} className="flex flex-col gap-4 p-6">
      <div className="flex items-start gap-3">
        <span
          className={
            walletOnAppChain
              ? "grid size-9 shrink-0 place-items-center rounded-tile bg-mint-soft text-mint-ink"
              : "grid size-9 shrink-0 place-items-center rounded-tile bg-surface-muted text-ink-soft"
          }
        >
          <Icon name={walletOnAppChain ? IconName.Check : IconName.Sliders} className="size-4.5" />
        </span>

        <div className="flex flex-col gap-1">
          <h3 className="text-base font-semibold text-ink">{title(isDisconnected, walletOnAppChain, appIsTestnet)}</h3>
          <p className="text-sm leading-relaxed text-ink-soft">
            {description(isDisconnected, walletOnAppChain, chainId)}
          </p>
        </div>
      </div>

      {isDisconnected || walletOnAppChain ? null : (
        <Button
          variant={ButtonVariant.Secondary}
          disabled={isPending}
          onClick={() => switchChain({ chainId: appChainId })}
        >
          {isPending ? "Switching" : `Switch to ${appChain.name}`}
        </Button>
      )}
    </Card>
  );
}

function title(isDisconnected: boolean, walletOnAppChain: boolean, appIsTestnet: boolean): string {
  if (isDisconnected) {
    return appIsTestnet ? `This app is pointed at ${appChain.name}` : "This app is pointed at a live network";
  }

  return walletOnAppChain ? `You are on ${appChain.name}` : `Switch your wallet to ${appChain.name}`;
}

function description(isDisconnected: boolean, walletOnAppChain: boolean, chainId: number | undefined): string {
  if (isDisconnected) {
    return "Connect a wallet when you are ready. Nothing here touches real funds.";
  }

  if (walletOnAppChain) {
    return "Nothing you do from here touches real funds. Your live positions are untouched and waiting.";
  }

  return `Your wallet is on ${chainDisplayName(chainId)}, which this market does not run on.`;
}
