"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useConnect } from "wagmi";
import type { Connector } from "wagmi";
import { AppRoute, BadgeTone, IconName } from "@/lib/enums";
import { loginContent } from "@/content/login";
import { useWalletState } from "@/hooks/useWalletState";
import { Alert } from "@/components/ui/Alert";
import { Icon } from "@/components/ui/Icon";
import { ExternalLink } from "@/components/ui/ExternalLink";

function findMetaMask(connectors: readonly Connector[]): Connector | undefined {
  const byName = connectors.find((connector) => connector.name.toLowerCase().includes("metamask"));

  if (byName !== undefined) {
    return byName;
  }

  return connectors.find((connector) => connector.id === "injected");
}

export function MetaMaskOption() {
  const router = useRouter();
  const { connectors, connect, status, error, reset } = useConnect();
  const { address, isSettling } = useWalletState();

  const isSignedIn = address !== undefined;

  useEffect(() => {
    if (isSignedIn) {
      router.replace(AppRoute.Dashboard);
    }
  }, [isSignedIn, router]);

  if (isSettling || isSignedIn) {
    return (
      <div className="flex items-center gap-3 rounded-card border border-shell-line bg-shell px-4 py-4">
        <span className="size-5 shrink-0 animate-spin rounded-full border-2 border-shell-line border-t-accent" />
        <span className="text-sm text-ink-soft">
          {isSignedIn ? loginContent.redirectingLabel : "Checking your wallet"}
        </span>
      </div>
    );
  }

  const metaMask = findMetaMask(connectors);
  const isPending = status === "pending";

  if (metaMask === undefined) {
    return (
      <div className="flex flex-col gap-4">
        <Alert title={loginContent.notDetectedTitle} tone={BadgeTone.Caution} icon={IconName.Info}>
          {loginContent.notDetectedBody}
        </Alert>

        <ExternalLink href={loginContent.installUrl}>{loginContent.installLabel}</ExternalLink>
      </div>
    );
  }

  return (
    <div className="flex flex-col gap-4">
      <button
        type="button"
        disabled={isPending}
        onClick={() => {
          reset();
          connect({ connector: metaMask });
        }}
        className="flex w-full items-center justify-between gap-3 rounded-card border border-line bg-surface px-4 py-4 text-left transition-colors hover:border-accent-border disabled:cursor-not-allowed disabled:opacity-60 outline-none focus-visible:ring-2 focus-visible:ring-accent focus-visible:ring-offset-2 focus-visible:ring-offset-canvas"
      >
        <span className="flex items-center gap-3">
          <span className="grid size-10 place-items-center rounded-tile bg-accent-soft text-accent">
            <Icon name={IconName.Wallet} className="size-5" />
          </span>
          <span className="flex flex-col">
            <span className="text-sm font-medium text-ink">{metaMask.name}</span>
            <span className="text-xs text-ink-faint">
              {isPending ? "Confirm in your wallet" : "Browser extension"}
            </span>
          </span>
        </span>
        <Icon name={IconName.ArrowRight} className="size-4 text-ink-faint" />
      </button>

      {error === null ? null : (
        <Alert title="That did not work" tone={BadgeTone.Caution} icon={IconName.Warning}>
          {error.message.includes("rejected")
            ? "The request was rejected in your wallet. Nothing was shared and nothing changed."
            : error.message}
        </Alert>
      )}
    </div>
  );
}
