import type { ActivityKind } from "@/lib/enums";
import { formatDateTimeUtc, truncateMiddle } from "@/lib/format";
import { formatHealthFactor } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import { historyPageContent } from "@/content/history";

export type TransactionView = {
  kind: ActivityKind;
  amount: bigint;
  symbol: string;
  decimals: number;
  timestamp: string;
  blockNumber: number;
  txHash: string;
  logIndex?: number;
  healthFactorAfterBps: bigint | null;
};

export type DetailRow = {
  label: string;
  value: string;
  hint?: string;
};

export function transactionDetailRows(view: TransactionView): DetailRow[] {
  const rows: DetailRow[] = [
    {
      label: "Amount",
      value: `${formatTokenAmount(view.amount, view.decimals, 6)} ${view.symbol}`,
    },
    {
      label: "Health factor afterwards",
      value: view.healthFactorAfterBps === null ? "Not applicable" : formatHealthFactor(view.healthFactorAfterBps),
      hint: view.healthFactorAfterBps === null ? historyPageContent.healthNotApplicable : undefined,
    },
    {
      label: "When",
      value: formatDateTimeUtc(view.timestamp),
    },
    {
      label: "Block",
      value: view.blockNumber.toLocaleString("en-US"),
    },
    {
      label: "Transaction",
      value: truncateMiddle(view.txHash, 10, 8),
    },
  ];

  if (view.logIndex !== undefined) {
    rows.push({
      label: "Event position",
      value: view.logIndex.toLocaleString("en-US"),
      hint: historyPageContent.logIndexHint,
    });
  }

  return rows;
}
