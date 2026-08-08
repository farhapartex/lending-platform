import { SurfaceElevation } from "@/lib/enums";
import type { HistoryEntry } from "@/content/history";
import { Card } from "@/components/ui/Card";
import { Pagination } from "@/components/ui/Pagination";
import { TxRow } from "@/components/history/TxRow";

const headers = ["Type", "Amount", "Health after", "Date", ""];

type TxHistoryTableProps = {
  entries: HistoryEntry[];
  page: number;
  pageCount: number;
  totalItems: number;
  pageSize: number;
  onPageChange: (page: number) => void;
  onSelect: (entry: HistoryEntry) => void;
};

export function TxHistoryTable({
  entries,
  page,
  pageCount,
  totalItems,
  pageSize,
  onPageChange,
  onSelect,
}: TxHistoryTableProps) {
  return (
    <Card elevation={SurfaceElevation.Raised} className="overflow-hidden">
      <div className="overflow-x-auto">
        <table className="w-full min-w-[44rem] border-collapse text-left">
          <caption className="sr-only">Your transactions, newest first</caption>
          <thead>
            <tr className="bg-surface-muted">
              {headers.map((header, index) => (
                <th
                  key={header === "" ? `actions-${index}` : header}
                  scope="col"
                  className="whitespace-nowrap px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint"
                >
                  {header === "" ? <span className="sr-only">Actions</span> : header}
                </th>
              ))}
            </tr>
          </thead>
          <tbody>
            {entries.map((entry) => (
              <TxRow key={entry.id} entry={entry} onSelect={onSelect} />
            ))}
          </tbody>
        </table>
      </div>

      <Pagination
        page={page}
        pageCount={pageCount}
        totalItems={totalItems}
        pageSize={pageSize}
        onPageChange={onPageChange}
      />
    </Card>
  );
}
