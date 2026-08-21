"use client";

import { BadgeTone, ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { isNotFound } from "@/lib/api/errors";
import { explorerTxBaseUrl } from "@/content/protocol";
import { historyPageContent, type HistoryEntry } from "@/content/history";
import { useTransactionDetail } from "@/hooks/useTransactionDetail";
import { Alert } from "@/components/ui/Alert";
import { Button } from "@/components/ui/Button";
import { Drawer } from "@/components/ui/Drawer";
import { ExternalLink } from "@/components/ui/ExternalLink";
import { MetricRow } from "@/components/ui/MetricRow";
import { TxTypeBadge } from "@/components/tx/TxTypeBadge";
import { transactionDetailRows } from "@/components/history/txDetailPresentation";

type TxDetailDrawerProps = {
  entry: HistoryEntry | null;
  address?: string;
  onClose: () => void;
};

export function TxDetailDrawer({ entry, address, onClose }: TxDetailDrawerProps) {
  const { detail, isLoading, isError, error, isEnabled, refetch } = useTransactionDetail(address, entry?.id ?? null);

  if (entry === null) {
    return null;
  }

  const view = detail ?? entry;
  const rows = transactionDetailRows(view);
  const missing = isError && isNotFound(error);

  return (
    <Drawer open onClose={onClose} title="Transaction details">
      <div className="flex flex-col gap-5">
        <TxTypeBadge kind={view.kind} />

        {isLoading && detail === undefined ? (
          <p className="flex items-center gap-2 text-sm text-ink-soft">
            <span className="size-1.5 animate-pulse rounded-full bg-brand-ink" />
            {historyPageContent.detailRefreshing}
          </p>
        ) : null}

        {missing ? (
          <Alert title={historyPageContent.detailMissingTitle} tone={BadgeTone.Caution} icon={IconName.Info}>
            {historyPageContent.detailMissingDescription}
          </Alert>
        ) : null}

        {isError && !missing ? (
          <Alert title={historyPageContent.detailUnavailableTitle} tone={BadgeTone.Neutral} icon={IconName.Info}>
            <div className="flex flex-col items-start gap-2">
              <p>{historyPageContent.detailUnavailableDescription}</p>
              <Button variant={ButtonVariant.Subtle} size={ButtonSize.Sm} onClick={refetch}>
                {historyPageContent.retry}
              </Button>
            </div>
          </Alert>
        ) : null}

        {!isEnabled ? (
          <p className="text-sm text-ink-faint">{historyPageContent.detailSampleNotice}</p>
        ) : null}

        <dl className="divide-y divide-line rounded-card border border-line bg-surface px-5">
          {rows.map((row) => (
            <MetricRow key={row.label} label={row.label} value={row.value} hint={row.hint} />
          ))}
        </dl>

        <ExternalLink href={`${explorerTxBaseUrl}${view.txHash}`}>View on the block explorer</ExternalLink>
      </div>
    </Drawer>
  );
}
