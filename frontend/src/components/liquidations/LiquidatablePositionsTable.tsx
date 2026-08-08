import { SurfaceElevation } from "@/lib/enums";
import type { LiquidationRow } from "@/lib/liquidation";
import { Card } from "@/components/ui/Card";
import { LiquidatePositionRow } from "@/components/liquidations/LiquidatePositionRow";

const headers = ["Borrower", "Collateral", "Loan owed", "Health", "Your bonus", ""];

type LiquidatablePositionsTableProps = {
  rows: LiquidationRow[];
  onSelect: (row: LiquidationRow) => void;
};

export function LiquidatablePositionsTable({ rows, onSelect }: LiquidatablePositionsTableProps) {
  return (
    <Card elevation={SurfaceElevation.Raised} className="overflow-hidden">
      <div className="overflow-x-auto">
        <table className="w-full min-w-[46rem] border-collapse text-left">
          <caption className="sr-only">Positions currently eligible for liquidation</caption>
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
            {rows.map((row) => (
              <LiquidatePositionRow key={row.id} row={row} onSelect={onSelect} />
            ))}
          </tbody>
        </table>
      </div>
    </Card>
  );
}
