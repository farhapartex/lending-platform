"use client";

import { useState } from "react";
import { BadgeTone, ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import type { Liquidation } from "@/lib/api/liquidationMapper";
import { liquidationsPageContent } from "@/content/liquidations";
import { useLiquidationHistory } from "@/hooks/useLiquidationHistory";
import { Alert } from "@/components/ui/Alert";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { EmptyState } from "@/components/ui/EmptyState";
import { Skeleton } from "@/components/ui/Skeleton";
import { LiquidationHistoryTable } from "@/components/liquidations/LiquidationHistoryTable";
import { LiquidationReceiptDrawer } from "@/components/liquidations/LiquidationReceiptDrawer";

export function LiquidationHistoryPanel() {
  const [cursor, setCursor] = useState<string | null>(null);
  const [selected, setSelected] = useState<Liquidation | null>(null);

  const { page, isLoading, isError, isFetching, refetch } = useLiquidationHistory(cursor);

  if (isLoading) {
    return <HistorySkeleton />;
  }

  if (isError) {
    return (
      <Alert title={liquidationsPageContent.historyUnavailableTitle} tone={BadgeTone.Neutral} icon={IconName.Info}>
        <div className="flex flex-col items-start gap-2">
          <p>{liquidationsPageContent.historyUnavailableDescription}</p>
          <Button variant={ButtonVariant.Subtle} size={ButtonSize.Sm} onClick={refetch}>
            {liquidationsPageContent.historyRetry}
          </Button>
        </div>
      </Alert>
    );
  }

  const rows = page?.items ?? [];

  if (rows.length === 0) {
    return <HistoryEmptyState indexed={page?.asOf.block !== null} />;
  }

  return (
    <div className="flex flex-col gap-4">
      <LiquidationHistoryTable rows={rows} onSelect={setSelected} />

      {page?.nextCursor === null || page?.nextCursor === undefined ? null : (
        <div className="flex justify-center">
          <Button
            variant={ButtonVariant.Secondary}
            size={ButtonSize.Sm}
            onClick={() => setCursor(page.nextCursor)}
            disabled={isFetching}
          >
            {liquidationsPageContent.loadMore}
          </Button>
        </div>
      )}

      <LiquidationReceiptDrawer summary={selected} onClose={() => setSelected(null)} />
    </div>
  );
}

function HistoryEmptyState({ indexed }: { indexed: boolean }) {
  if (!indexed) {
    return (
      <EmptyState
        title={liquidationsPageContent.historyNotIndexedTitle}
        description={liquidationsPageContent.historyNotIndexedDescription}
        icon={IconName.Info}
      />
    );
  }

  return (
    <EmptyState
      title={liquidationsPageContent.historyEmptyTitle}
      description={liquidationsPageContent.historyEmptyDescription}
      icon={IconName.ShieldCheck}
    />
  );
}

function HistorySkeleton() {
  return (
    <Card className="overflow-hidden p-0">
      <div className="flex flex-col divide-y divide-line">
        {[0, 1, 2].map((row) => (
          <div key={row} className="flex items-center justify-between gap-4 px-5 py-4">
            <Skeleton className="h-4 w-40" />
            <Skeleton className="h-4 w-24" />
            <Skeleton className="h-4 w-28" />
            <Skeleton className="h-4 w-20" />
          </div>
        ))}
      </div>
    </Card>
  );
}
