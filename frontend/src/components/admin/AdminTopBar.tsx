"use client";

import { ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { useWalletState } from "@/hooks/useWalletState";
import { Button } from "@/components/ui/Button";
import { Icon } from "@/components/ui/Icon";
import { NetworkBadge } from "@/components/app/NetworkBadge";
import { WalletPill } from "@/components/admin/WalletPill";

type AdminTopBarProps = {
  onOpenSidebar: () => void;
};

export function AdminTopBar({ onOpenSidebar }: AdminTopBarProps) {
  const { address, isSettling } = useWalletState();

  return (
    <header className="sticky top-0 z-30 flex h-16 items-center justify-between gap-4 border-b border-shell-line bg-surface px-4 sm:px-6">
      <Button
        variant={ButtonVariant.Ghost}
        size={ButtonSize.Sm}
        className="lg:hidden"
        onClick={onOpenSidebar}
        ariaLabel="Open the menu"
      >
        <Icon name={IconName.Menu} className="size-4" />
      </Button>

      <div className="ml-auto flex items-center gap-3">
        <NetworkBadge />

        {isSettling || address === undefined ? (
          <span className="h-9 w-44 animate-pulse rounded-pill bg-surface-muted" />
        ) : (
          <WalletPill address={address} />
        )}
      </div>
    </header>
  );
}
