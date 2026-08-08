"use client";

import { formatDateTimeUtc, truncateMiddle } from "@/lib/format";
import { formatHealthFactor } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import { explorerTxBaseUrl } from "@/content/protocol";
import { historyPageContent, type HistoryEntry } from "@/content/history";
import { Drawer } from "@/components/ui/Drawer";
import { ExternalLink } from "@/components/ui/ExternalLink";
import { MetricRow } from "@/components/ui/MetricRow";
import { TxTypeBadge } from "@/components/tx/TxTypeBadge";

type TxDetailDrawerProps = {
  entry: HistoryEntry | null;
  onClose: () => void;
};

export function TxDetailDrawer({ entry, onClose }: TxDetailDrawerProps) {
  if (entry === null) {
    return null;
  }

  const rows = [
    {
      label: "Amount",
      value: `${formatTokenAmount(entry.amount, entry.decimals, 6)} ${entry.symbol}`,
    },
    {
      label: "Health factor afterwards",
      value: entry.healthFactorAfterBps === null ? "Not applicable" : formatHealthFactor(entry.healthFactorAfterBps),
      hint: entry.healthFactorAfterBps === null ? historyPageContent.healthNotApplicable : undefined,
    },
    {
      label: "When",
      value: formatDateTimeUtc(entry.timestamp),
    },
    {
      label: "Block",
      value: entry.blockNumber.toLocaleString("en-US"),
    },
    {
      label: "Transaction",
      value: truncateMiddle(entry.txHash, 10, 8),
    },
  ];

  return (
    <Drawer open onClose={onClose} title="Transaction details">
      <div className="flex flex-col gap-5">
        <TxTypeBadge kind={entry.kind} />

        <dl className="divide-y divide-line rounded-card border border-line bg-surface px-5">
          {rows.map((row) => (
            <MetricRow key={row.label} label={row.label} value={row.value} hint={row.hint} />
          ))}
        </dl>

        <ExternalLink href={`${explorerTxBaseUrl}${entry.txHash}`}>View on the block explorer</ExternalLink>
      </div>
    </Drawer>
  );
}
