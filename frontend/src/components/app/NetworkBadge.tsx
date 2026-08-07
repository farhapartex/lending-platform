import { BadgeTone, NetworkKind } from "@/lib/enums";
import { activeNetwork } from "@/content/protocol";
import { Badge } from "@/components/ui/Badge";

const networkLabels: Record<NetworkKind, string> = {
  [NetworkKind.Mainnet]: "Mainnet",
  [NetworkKind.Testnet]: "Testnet",
};

const networkTones: Record<NetworkKind, BadgeTone> = {
  [NetworkKind.Mainnet]: BadgeTone.Neutral,
  [NetworkKind.Testnet]: BadgeTone.Caution,
};

export function NetworkBadge() {
  return <Badge tone={networkTones[activeNetwork]}>{networkLabels[activeNetwork]}</Badge>;
}
