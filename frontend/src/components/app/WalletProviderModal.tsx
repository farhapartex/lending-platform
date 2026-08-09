"use client";

import { useConnect } from "wagmi";
import { BadgeTone, IconName } from "@/lib/enums";
import { Alert } from "@/components/ui/Alert";
import { Icon } from "@/components/ui/Icon";
import { Modal } from "@/components/ui/Modal";

type WalletProviderModalProps = {
  open: boolean;
  onClose: () => void;
};

export function WalletProviderModal({ open, onClose }: WalletProviderModalProps) {
  const { connectors, connect, status, error, reset } = useConnect();

  const isPending = status === "pending";

  const handleClose = () => {
    reset();
    onClose();
  };

  return (
    <Modal open={open} onClose={handleClose} title="Connect a wallet">
      <div className="flex flex-col gap-4">
        <p className="text-sm leading-relaxed text-ink-soft">
          Connecting lets you move funds. It does not create an account, and there is no password.
        </p>

        {connectors.length === 0 ? (
          <Alert title="No wallet detected" tone={BadgeTone.Caution} icon={IconName.Info}>
            Install a browser wallet such as MetaMask, or open this page in a wallet app, then try again.
          </Alert>
        ) : (
          <ul className="flex flex-col gap-2">
            {connectors.map((connector) => (
              <li key={connector.uid}>
                <button
                  type="button"
                  disabled={isPending}
                  onClick={() => connect({ connector })}
                  className="flex w-full items-center justify-between gap-3 rounded-card border border-line bg-surface px-4 py-3 text-left transition-colors hover:border-brand-border disabled:cursor-not-allowed disabled:opacity-60 outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas"
                >
                  <span className="flex items-center gap-3">
                    <span className="grid size-9 place-items-center rounded-tile bg-brand-soft text-brand">
                      <Icon name={IconName.Wallet} className="size-4.5" />
                    </span>
                    <span className="text-sm font-medium text-ink">{connector.name}</span>
                  </span>
                  <Icon name={IconName.ArrowRight} className="size-4 text-ink-faint" />
                </button>
              </li>
            ))}
          </ul>
        )}

        {error === null ? null : (
          <Alert title="That did not work" tone={BadgeTone.Caution} icon={IconName.Warning}>
            {error.message.includes("rejected")
              ? "The request was rejected in your wallet. Nothing was shared and nothing changed."
              : error.message}
          </Alert>
        )}

        <p className="text-xs leading-relaxed text-ink-faint">
          We never receive your recovery phrase or private key, and we cannot move your funds.
        </p>
      </div>
    </Modal>
  );
}
