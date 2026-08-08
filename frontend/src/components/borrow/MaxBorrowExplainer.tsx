import { AssetSymbol, IconName, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { basisPoints } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import { debtDecimals, maxLtvBps } from "@/content/borrow";
import { Icon } from "@/components/ui/Icon";

type MaxBorrowExplainerProps = {
  capacity: bigint;
};

export function MaxBorrowExplainer({ capacity }: MaxBorrowExplainerProps) {
  return (
    <p className="flex items-start gap-2 text-sm leading-relaxed text-ink-soft">
      <Icon name={IconName.Info} className="mt-0.5 size-4 text-brand" />
      <span>
        You can borrow up to {formatValue(Number(maxLtvBps) / Number(basisPoints), ValueFormat.Percent)} of your
        collateral&apos;s value, which leaves{" "}
        <span className="font-medium text-ink tabular-nums">
          {formatTokenAmount(capacity, debtDecimals, 2)} {AssetSymbol.Usdc}
        </span>{" "}
        still available to you. Borrowing the full amount is allowed but leaves you no room for the price to move.
      </span>
    </p>
  );
}
