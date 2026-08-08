import { ButtonSize, ButtonVariant } from "@/lib/enums";
import { formatDateTimeUtc } from "@/lib/format";
import { formatHealthFactor } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import type { HistoryEntry } from "@/content/history";
import { Button } from "@/components/ui/Button";
import { TxTypeBadge } from "@/components/tx/TxTypeBadge";

type TxRowProps = {
  entry: HistoryEntry;
  onSelect: (entry: HistoryEntry) => void;
};

export function TxRow({ entry, onSelect }: TxRowProps) {
  return (
    <tr className="border-t border-line align-middle">
      <td className="whitespace-nowrap px-5 py-4">
        <TxTypeBadge kind={entry.kind} />
      </td>

      <td className="whitespace-nowrap px-5 py-4">
        <span className="text-sm font-medium text-ink tabular-nums">
          {formatTokenAmount(entry.amount, entry.decimals, 4)} {entry.symbol}
        </span>
      </td>

      <td className="whitespace-nowrap px-5 py-4">
        {entry.healthFactorAfterBps === null ? (
          <span className="text-sm text-ink-faint">Not applicable</span>
        ) : (
          <span className="text-sm text-ink tabular-nums">{formatHealthFactor(entry.healthFactorAfterBps)}</span>
        )}
      </td>

      <td className="whitespace-nowrap px-5 py-4">
        <span className="text-sm text-ink-soft">{formatDateTimeUtc(entry.timestamp)}</span>
      </td>

      <td className="whitespace-nowrap px-5 py-4 text-right">
        <Button variant={ButtonVariant.Secondary} size={ButtonSize.Sm} onClick={() => onSelect(entry)}>
          Details
        </Button>
      </td>
    </tr>
  );
}
