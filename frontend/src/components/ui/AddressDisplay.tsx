import { cn } from "@/lib/cn";

const leadingCharacters = 6;
const trailingCharacters = 4;

export function truncateAddress(address: string): string {
  if (address.length <= leadingCharacters + trailingCharacters) {
    return address;
  }

  return `${address.slice(0, leadingCharacters)}…${address.slice(-trailingCharacters)}`;
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
