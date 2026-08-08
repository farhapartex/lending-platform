"use client";

import { useMemo, useState } from "react";
import { IconName, LiquidationSortKey } from "@/lib/enums";
import {
  collateralDecimals,
  collateralUnitPriceScaled,
  debtDecimals,
  debtUnitPriceScaled,
  liquidationBonusBps,
  liquidationThresholdBps,
} from "@/content/protocol";
import { liquidationCandidates, liquidationsPageContent } from "@/content/liquidations";
import { buildLiquidationRow, compareBigInt, isLiquidatable, type LiquidationRow } from "@/lib/liquidation";
import { EmptyState } from "@/components/ui/EmptyState";
import { FilterSortControls } from "@/components/liquidations/FilterSortControls";
import { LiquidatablePositionsTable } from "@/components/liquidations/LiquidatablePositionsTable";
import { LiquidateModal } from "@/components/liquidations/LiquidateModal";

const rowParams = {
  collateralDecimals,
  debtDecimals,
  collateralUnitPriceScaled,
  debtUnitPriceScaled,
  liquidationThresholdBps,
  bonusBps: liquidationBonusBps,
};

const eligibleRows = liquidationCandidates
  .map((candidate) => buildLiquidationRow(candidate, rowParams))
  .filter((row) => isLiquidatable(row.factorBps));

function sortRows(rows: LiquidationRow[], key: LiquidationSortKey): LiquidationRow[] {
  const sorted = [...rows];

  if (key === LiquidationSortKey.Health) {
    return sorted.sort((left, right) => compareBigInt(left.factorBps ?? 0n, right.factorBps ?? 0n));
  }

  if (key === LiquidationSortKey.Size) {
    return sorted.sort((left, right) => compareBigInt(right.debtValueScaled, left.debtValueScaled));
  }

  return sorted.sort((left, right) => compareBigInt(right.bonusValueScaled, left.bonusValueScaled));
}

export function LiquidationsPanel() {
  const [sortKey, setSortKey] = useState(LiquidationSortKey.Health);
  const [selected, setSelected] = useState<LiquidationRow | null>(null);

  const rows = useMemo(() => sortRows(eligibleRows, sortKey), [sortKey]);

  if (rows.length === 0) {
    return (
      <EmptyState
        title={liquidationsPageContent.emptyTitle}
        description={liquidationsPageContent.emptyDescription}
        icon={IconName.ShieldCheck}
      />
    );
  }

  return (
    <div className="flex flex-col gap-5">
      <FilterSortControls sortKey={sortKey} onSortChange={setSortKey} visibleCount={rows.length} />
      <LiquidatablePositionsTable rows={rows} onSelect={setSelected} />
      <LiquidateModal row={selected} onClose={() => setSelected(null)} />
    </div>
  );
}
