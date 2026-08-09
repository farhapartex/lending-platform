"use client";

import { useSwitchChain } from "wagmi";
import { ButtonSize, ButtonVariant, IconName, WalletStatus } from "@/lib/enums";
import { appChain, chainDisplayName } from "@/lib/chain";
import { useWalletState } from "@/hooks/useWalletState";
import { Button } from "@/components/ui/Button";
import { Container } from "@/components/ui/Container";
import { Icon } from "@/components/ui/Icon";

export function WrongNetworkBanner() {
  const { status, chainId } = useWalletState();
  const { switchChain, isPending } = useSwitchChain();

  if (status !== WalletStatus.WrongNetwork) {
    return null;
  }

  return (
    <div className="border-b border-amber-border bg-amber-soft">
      <Container className="flex flex-wrap items-center justify-between gap-3 py-2.5">
        <span className="flex items-center gap-2 text-sm text-amber-ink">
          <Icon name={IconName.Warning} className="size-4 shrink-0" />
          Your wallet is on {chainDisplayName(chainId)}. This market runs on {appChain.name}.
        </span>

        <Button
          variant={ButtonVariant.Secondary}
          size={ButtonSize.Sm}
          disabled={isPending}
          onClick={() => switchChain({ chainId: appChain.id })}
        >
          {isPending ? "Switching" : `Switch to ${appChain.name}`}
        </Button>
      </Container>
    </div>
  );
}
