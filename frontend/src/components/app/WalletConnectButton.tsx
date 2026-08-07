import { ButtonSize, ButtonVariant, IconName, WalletStatus } from "@/lib/enums";
import { walletAddress, walletStatus } from "@/content/wallet";
import { AddressDisplay } from "@/components/ui/AddressDisplay";
import { Button } from "@/components/ui/Button";
import { Icon } from "@/components/ui/Icon";

const labels: Record<WalletStatus, string> = {
  [WalletStatus.Disconnected]: "Connect wallet",
  [WalletStatus.Connecting]: "Connecting",
  [WalletStatus.Connected]: "",
  [WalletStatus.WrongNetwork]: "Wrong network",
};

export function WalletConnectButton() {
  if (walletStatus === WalletStatus.Connected) {
    return (
      <span className="inline-flex items-center gap-2 rounded-pill border border-line-strong bg-surface px-3 py-2">
        <Icon name={IconName.Wallet} className="size-4 text-mint" />
        <AddressDisplay address={walletAddress} />
      </span>
    );
  }

  return (
    <Button
      variant={walletStatus === WalletStatus.WrongNetwork ? ButtonVariant.Primary : ButtonVariant.Secondary}
      size={ButtonSize.Sm}
      leadingIcon={IconName.Wallet}
      disabled
    >
      {labels[walletStatus]}
    </Button>
  );
}
