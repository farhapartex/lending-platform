"use client";

import { BadgeTone, ButtonSize, ButtonVariant } from "@/lib/enums";
import { formatDateTimeUtc } from "@/lib/format";
import type { Liquidation } from "@/lib/api/liquidationMapper";
import { hasShortfall } from "@/lib/api/liquidationMapper";
import { liquidationsPageContent } from "@/content/liquidations";
import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { formatApiAmount, formatUsdValue } from "@/components/liquidations/liquidationPresentation";

type LiquidationHistoryTableProps = {
  rows: Liquidation[];
  onSelect: (liquidation: Liquidation) => void;
};

export function LiquidationHistoryTable({ rows, onSelect }: LiquidationHistoryTableProps) {
  return (
    <Card className="overflow-hidden p-0">
      <div className="overflow-x-auto">
        <table className="w-full min-w-3xl border-collapse text-left">
          <caption className="sr-only">{liquidationsPageContent.historyTitle}</caption>
          <thead className="bg-surface-muted text-xs font-medium uppercase tracking-wide text-ink-faint">
            <tr>
              <th scope="col" className="px-5 py-3">
                When
              </th>
              <th scope="col" className="px-5 py-3">
                Debt repaid
              </th>
              <th scope="col" className="px-5 py-3">
                Collateral seized
              </th>
              <th scope="col" className="px-5 py-3">
                Bonus
              </th>
              <th scope="col" className="px-5 py-3">
                <span className="sr-only">Receipt</span>
              </th>
            </tr>
          </thead>
          <tbody>
            {rows.map((row) => (
              <tr key={row.id} className="border-t border-line align-middle">
                <td className="whitespace-nowrap px-5 py-4">
                  <div className="flex flex-col gap-1">
                    <span className="text-sm text-ink">{formatDateTimeUtc(row.timestamp)}</span>
                    {hasShortfall(row) ? (
                      <Badge tone={BadgeTone.Critical}>{liquidationsPageContent.shortfallBadge}</Badge>
                    ) : null}
                  </div>
                </td>
                <td className="whitespace-nowrap px-5 py-4 text-sm font-medium text-ink tabular-nums">
                  {formatApiAmount(row.debtRepaid, 2)}
                </td>
                <td className="whitespace-nowrap px-5 py-4 text-sm font-medium text-ink tabular-nums">
                  {formatApiAmount(row.collateralSeized, 6)}
                </td>
                <td className="whitespace-nowrap px-5 py-4 text-sm font-medium text-mint-ink tabular-nums">
                  {formatUsdValue(row.bonusValue)}
                </td>
                <td className="whitespace-nowrap px-5 py-4 text-right">
                  <Button variant={ButtonVariant.Ghost} size={ButtonSize.Sm} onClick={() => onSelect(row)}>
                    Receipt
                  </Button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </Card>
  );
}
