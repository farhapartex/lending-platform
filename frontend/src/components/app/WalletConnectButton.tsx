import { ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { Button } from "@/components/ui/Button";

export function WalletConnectButton() {
  return (
    <Button variant={ButtonVariant.Secondary} size={ButtonSize.Sm} leadingIcon={IconName.Wallet} disabled>
      Connect wallet
    </Button>
  );
}
