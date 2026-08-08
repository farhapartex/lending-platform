"use client";

import { useMemo, useState } from "react";
import {
  ButtonVariant,
  DateRangePreset,
  IconName,
  allTypesFilter,
  type TypeFilterValue,
} from "@/lib/enums";
import { historyAsOf, historyEntries, historyPageContent, historyPageSize, type HistoryEntry } from "@/content/history";
import { Button } from "@/components/ui/Button";
import { EmptyState } from "@/components/ui/EmptyState";
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

const asOfMs = new Date(historyAsOf).getTime();

const sortedEntries = [...historyEntries].sort(
  (left, right) => new Date(right.timestamp).getTime() - new Date(left.timestamp).getTime(),
);

function withinRange(entry: HistoryEntry, preset: DateRangePreset): boolean {
  const days = presetDays[preset];

  if (days === null) {
    return true;
  }

  return new Date(entry.timestamp).getTime() >= asOfMs - days * millisecondsPerDay;
}

export function HistoryPanel() {
  const [typeFilter, setTypeFilter] = useState<TypeFilterValue>(allTypesFilter);
  const [rangeFilter, setRangeFilter] = useState(DateRangePreset.AllTime);
  const [page, setPage] = useState(1);
  const [selected, setSelected] = useState<HistoryEntry | null>(null);

  const filtered = useMemo(
    () =>
      sortedEntries.filter(
        (entry) => (typeFilter === allTypesFilter || entry.kind === typeFilter) && withinRange(entry, rangeFilter),
      ),
    [typeFilter, rangeFilter],
  );

  const pageCount = Math.max(1, Math.ceil(filtered.length / historyPageSize));
  const safePage = Math.min(page, pageCount);
  const visible = filtered.slice((safePage - 1) * historyPageSize, safePage * historyPageSize);

  const resetFilters = () => {
    setTypeFilter(allTypesFilter);
    setRangeFilter(DateRangePreset.AllTime);
    setPage(1);
  };

  if (sortedEntries.length === 0) {
    return (
      <EmptyState
        title={historyPageContent.emptyTitle}
        description={historyPageContent.emptyDescription}
        icon={IconName.Receipt}
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

      {filtered.length === 0 ? (
        <EmptyState
          title={historyPageContent.noMatchTitle}
          description={historyPageContent.noMatchDescription}
          icon={IconName.Info}
          action={
            <Button variant={ButtonVariant.Subtle} onClick={resetFilters}>
              Clear filters
            </Button>
          }
        />
      ) : (
        <TxHistoryTable
          entries={visible}
          page={safePage}
          pageCount={pageCount}
          totalItems={filtered.length}
          pageSize={historyPageSize}
          onPageChange={setPage}
          onSelect={setSelected}
        />
      )}

      <TxDetailDrawer entry={selected} onClose={() => setSelected(null)} />
    </div>
  );
}
