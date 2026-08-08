import { cn } from "@/lib/cn";
import { truncateMiddle } from "@/lib/format";

const leadingCharacters = 6;
const trailingCharacters = 4;

export function truncateAddress(address: string): string {
  return truncateMiddle(address, leadingCharacters, trailingCharacters);
}

type AddressDisplayProps = {
  address: string;
  className?: string;
};

export function AddressDisplay({ address, className }: AddressDisplayProps) {
  return (
    <span title={address} className={cn("font-mono text-sm text-ink", className)}>
      {truncateAddress(address)}
    </span>
  );
}
