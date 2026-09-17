"use client";

import { useMemo, useState } from "react";
import {
  ButtonVariant,
  DateRangePreset,
  IconName,
  allTypesFilter,
  type TypeFilterValue,
} from "@/lib/enums";
import { historyPageContent, historyPageSize, type HistoryEntry } from "@/content/history";
import { wireValuesByKind } from "@/lib/api/transactionMapper";
import { useTransactionList } from "@/hooks/useTransactionList";
import { useWalletState } from "@/hooks/useWalletState";
import { Button } from "@/components/ui/Button";
import { EmptyState } from "@/components/ui/EmptyState";
import { Skeleton } from "@/components/ui/Skeleton";
import { DateRangeFilter } from "@/components/history/DateRangeFilter";
import { TxDetailDrawer } from "@/components/history/TxDetailDrawer";
import { TxHistoryTable } from "@/components/history/TxHistoryTable";
import { TypeFilter } from "@/components/history/TypeFilter";

const millisecondsPerDay = 86_400_000;

const presetDays: Record<DateRangePreset, number | null> = {
  [DateRangePreset.AllTime]: null,
  [DateRangePreset.Last7Days]: 7,
  [DateRangePreset.Last30Days]: 30,
  [DateRangePreset.Last90Days]: 90,
};

function rangeStart(preset: DateRangePreset): string | undefined {
  const days = presetDays[preset];

  if (days === null) {
    return undefined;
  }

  return new Date(Date.now() - days * millisecondsPerDay).toISOString();
}

export function HistoryPanel() {
  const { address } = useWalletState();
  const [typeFilter, setTypeFilter] = useState<TypeFilterValue>(allTypesFilter);
  const [rangeFilter, setRangeFilter] = useState(DateRangePreset.AllTime);
  const [page, setPage] = useState(1);
  const [selected, setSelected] = useState<HistoryEntry | null>(null);

  const filters = useMemo(
    () => ({
      kinds: typeFilter === allTypesFilter ? undefined : [wireValuesByKind[typeFilter]],
      from: rangeStart(rangeFilter),
      limit: historyPageSize * 10,
    }),
    [typeFilter, rangeFilter],
  );

  const list = useTransactionList(address, filters);

  const entries = useMemo<HistoryEntry[]>(
    () =>
      (list.page?.items ?? []).map((item) => ({
        id: item.id,
        kind: item.kind,
        amount: item.amount,
        symbol: item.symbol as HistoryEntry["symbol"],
        decimals: item.decimals,
        timestamp: item.timestamp,
        blockNumber: item.blockNumber,
        txHash: item.txHash,
        healthFactorAfterBps: item.healthFactorAfterBps,
      })),
    [list.page],
  );

  const pageCount = Math.max(1, Math.ceil(entries.length / historyPageSize));
  const safePage = Math.min(page, pageCount);
  const visible = entries.slice((safePage - 1) * historyPageSize, safePage * historyPageSize);

  const hasFilters = typeFilter !== allTypesFilter || rangeFilter !== DateRangePreset.AllTime;

  const resetFilters = () => {
    setTypeFilter(allTypesFilter);
    setRangeFilter(DateRangePreset.AllTime);
    setPage(1);
  };

  if (list.isError) {
    return (
      <EmptyState
        title={historyPageContent.unavailableTitle}
        description={historyPageContent.unavailableDescription}
        icon={IconName.Warning}
        action={
          <Button variant={ButtonVariant.Subtle} onClick={list.refetch}>
            Try again
          </Button>
        }
      />
    );
  }

  return (
    <div className="flex flex-col gap-5">
      <div className="flex flex-wrap items-end gap-4">
        <TypeFilter
          value={typeFilter}
          onChange={(value) => {
            setTypeFilter(value);
            setPage(1);
          }}
        />
        <DateRangeFilter
          value={rangeFilter}
          onChange={(value) => {
            setRangeFilter(value);
            setPage(1);
          }}
        />
      </div>

      {list.isLoading ? <Skeleton className="h-64 w-full rounded-card" /> : null}

      {!list.isLoading && entries.length === 0 ? (
        <EmptyState
          title={hasFilters ? historyPageContent.noMatchTitle : historyPageContent.emptyTitle}
          description={hasFilters ? historyPageContent.noMatchDescription : historyPageContent.emptyDescription}
          icon={hasFilters ? IconName.Info : IconName.Receipt}
          action={
            hasFilters ? (
              <Button variant={ButtonVariant.Subtle} onClick={resetFilters}>
                Clear filters
              </Button>
            ) : undefined
          }
        />
      ) : null}

      {!list.isLoading && entries.length > 0 ? (
        <TxHistoryTable
          entries={visible}
          page={safePage}
          pageCount={pageCount}
          totalItems={entries.length}
          pageSize={historyPageSize}
          onPageChange={setPage}
          onSelect={setSelected}
        />
      ) : null}

      <TxDetailDrawer entry={selected} address={address} onClose={() => setSelected(null)} />
    </div>
  );
}
