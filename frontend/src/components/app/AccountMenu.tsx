"use client";

import { useEffect, useRef, useState } from "react";
import { useDisconnect } from "wagmi";
import { ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { chainDisplayName, explorerAddressUrl } from "@/lib/chain";
import { AddressDisplay } from "@/components/ui/AddressDisplay";
import { Button } from "@/components/ui/Button";
import { CopyButton } from "@/components/ui/CopyButton";
import { ExternalLink } from "@/components/ui/ExternalLink";
import { Icon } from "@/components/ui/Icon";

type AccountMenuProps = {
  address: string;
  chainId: number | undefined;
};

export function AccountMenu({ address, chainId }: AccountMenuProps) {
  const [isOpen, setIsOpen] = useState(false);
  const containerRef = useRef<HTMLDivElement>(null);
  const { disconnect } = useDisconnect();

  useEffect(() => {
    if (!isOpen) {
      return;
    }

    const handlePointerDown = (event: MouseEvent) => {
      if (containerRef.current !== null && !containerRef.current.contains(event.target as Node)) {
        setIsOpen(false);
      }
    };

    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.key === "Escape") {
        setIsOpen(false);
      }
    };

    document.addEventListener("mousedown", handlePointerDown);
    window.addEventListener("keydown", handleKeyDown);

    return () => {
      document.removeEventListener("mousedown", handlePointerDown);
      window.removeEventListener("keydown", handleKeyDown);
    };
  }, [isOpen]);

  const explorerUrl = explorerAddressUrl(address);

  return (
    <div ref={containerRef} className="relative">
      <button
        type="button"
        aria-haspopup="menu"
        aria-expanded={isOpen}
        onClick={() => setIsOpen(!isOpen)}
        className="inline-flex items-center gap-2 rounded-pill border border-line-strong bg-surface px-3 py-2 transition-colors hover:border-brand-border outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas"
      >
        <Icon name={IconName.Wallet} className="size-4 text-mint" />
        <AddressDisplay address={address} />
        <Icon
          name={IconName.ChevronDown}
          className={cn("size-3.5 text-ink-faint transition-transform", isOpen && "rotate-180")}
        />
      </button>

      {isOpen ? (
        <div
          role="menu"
          className="absolute right-0 top-full z-30 mt-2 w-64 overflow-hidden rounded-card border border-line bg-surface p-2 shadow-lift"
        >
          <div className="flex flex-col gap-1 border-b border-line px-2 pb-3 pt-1">
            <span className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">Connected</span>
            <span className="break-all font-mono text-xs text-ink">{address}</span>
            <span className="text-xs text-ink-soft">{chainDisplayName(chainId)}</span>
          </div>

          <div className="flex flex-col gap-1 pt-2">
            <CopyButton value={address} label="Copy address" />

            {explorerUrl === undefined ? null : (
              <ExternalLink href={explorerUrl} className="px-3 py-2 text-sm">
                View on explorer
              </ExternalLink>
            )}

            <Button
              variant={ButtonVariant.Ghost}
              size={ButtonSize.Sm}
              fullWidth
              onClick={() => {
                setIsOpen(false);
                disconnect();
              }}
            >
              Disconnect
            </Button>
          </div>
        </div>
      ) : null}
    </div>
  );
}
