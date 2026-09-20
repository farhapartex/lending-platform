"use client";

import { useState } from "react";
import { useDisconnect } from "wagmi";
import { IconName } from "@/lib/enums";
import { truncateMiddle } from "@/lib/format";
import { Icon } from "@/components/ui/Icon";

const feedbackDurationMs = 1600;

const iconButtonClasses =
  "grid size-7 shrink-0 place-items-center rounded-pill text-ink-faint transition-colors hover:bg-surface-muted hover:text-ink outline-none focus-visible:ring-2 focus-visible:ring-accent focus-visible:ring-offset-1 focus-visible:ring-offset-surface";

export function WalletPill({ address }: { address: string }) {
  const [copied, setCopied] = useState(false);
  const { disconnect } = useDisconnect();

  const handleCopy = () => {
    void navigator.clipboard
      .writeText(address)
      .then(() => {
        setCopied(true);
        window.setTimeout(() => setCopied(false), feedbackDurationMs);
      })
      .catch(() => setCopied(false));
  };

  return (
    <div className="flex items-center gap-1 rounded-pill border border-shell-line bg-surface py-1 pl-1 pr-1">
      <span className="grid size-7 shrink-0 place-items-center rounded-pill bg-accent-soft text-accent">
        <Icon name={IconName.Wallet} className="size-3.5" />
      </span>

      <span title={address} className="whitespace-nowrap px-1 font-mono text-sm text-ink">
        {truncateMiddle(address, 6, 4)}
      </span>

      <button
        type="button"
        onClick={handleCopy}
        aria-label={copied ? "Address copied" : "Copy the wallet address"}
        className={iconButtonClasses}
      >
        <Icon name={copied ? IconName.Check : IconName.Copy} className={copied ? "size-4 text-mint" : "size-4"} />
      </button>

      <span aria-hidden="true" className="mx-0.5 h-5 w-px bg-shell-line" />

      <button
        type="button"
        onClick={() => disconnect()}
        aria-label="Disconnect the wallet"
        title="Disconnect"
        className={`${iconButtonClasses} hover:bg-rose-soft hover:text-rose-ink`}
      >
        <Icon name={IconName.LogOut} className="size-4" />
      </button>
    </div>
  );
}
