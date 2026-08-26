"use client";

import { BadgeTone, ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { isNotFound } from "@/lib/api/errors";
import type { Liquidation } from "@/lib/api/liquidationMapper";
import { explorerTxBaseUrl } from "@/content/protocol";
import { liquidationsPageContent } from "@/content/liquidations";
import { useLiquidationReceipt } from "@/hooks/useLiquidationReceipt";
import { Alert } from "@/components/ui/Alert";
import { Button } from "@/components/ui/Button";
import { Drawer } from "@/components/ui/Drawer";
import { ExternalLink } from "@/components/ui/ExternalLink";
import { MetricRow } from "@/components/ui/MetricRow";
import { receiptRows } from "@/components/liquidations/liquidationPresentation";

type LiquidationReceiptDrawerProps = {
  summary: Liquidation | null;
  onClose: () => void;
};

export function LiquidationReceiptDrawer({ summary, onClose }: LiquidationReceiptDrawerProps) {
  const { receipt, isLoading, isError, error, refetch } = useLiquidationReceipt(summary?.id ?? null);

  if (summary === null) {
    return null;
  }

  const shown = receipt ?? summary;
  const rows = receiptRows(shown);
  const missing = isError && isNotFound(error);

  return (
    <Drawer open onClose={onClose} title={liquidationsPageContent.receiptTitle}>
      <div className="flex flex-col gap-5">
        {isLoading && receipt === undefined ? (
          <p className="flex items-center gap-2 text-sm text-ink-soft">
            <span className="size-1.5 animate-pulse rounded-full bg-brand-ink" />
            Loading the full record
          </p>
        ) : null}

        {missing ? (
          <Alert title={liquidationsPageContent.receiptMissingTitle} tone={BadgeTone.Caution} icon={IconName.Info}>
            {liquidationsPageContent.receiptMissingDescription}
          </Alert>
        ) : null}

        {isError && !missing ? (
          <Alert title={liquidationsPageContent.receiptUnavailableTitle} tone={BadgeTone.Neutral} icon={IconName.Info}>
            <div className="flex flex-col items-start gap-2">
              <p>{liquidationsPageContent.receiptUnavailableDescription}</p>
              <Button variant={ButtonVariant.Subtle} size={ButtonSize.Sm} onClick={refetch}>
                {liquidationsPageContent.historyRetry}
              </Button>
            </div>
          </Alert>
        ) : null}

        <dl className="divide-y divide-line rounded-card border border-line bg-surface px-5">
          {rows.map((row) => (
            <MetricRow key={row.label} label={row.label} value={row.value} hint={row.hint} />
          ))}
        </dl>

        <ExternalLink href={`${explorerTxBaseUrl}${shown.txHash}`}>View on the block explorer</ExternalLink>
      </div>
    </Drawer>
  );
}
