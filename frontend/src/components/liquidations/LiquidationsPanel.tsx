"use client";

import { useMemo, useState } from "react";
import { BadgeTone, ButtonVariant, IconName, LiquidationSortKey } from "@/lib/enums";
import { collateralDecimals, debtDecimals } from "@/content/protocol";
import { liquidationsPageContent } from "@/content/liquidations";
import { buildLiquidationRow, compareBigInt, isLiquidatable, type LiquidationRow } from "@/lib/liquidation";
import { useEligiblePositions } from "@/hooks/useEligiblePositions";
import { usePositionView } from "@/hooks/usePositionView";
import { Alert } from "@/components/ui/Alert";
import { Button } from "@/components/ui/Button";
import { EmptyState } from "@/components/ui/EmptyState";
import { Skeleton } from "@/components/ui/Skeleton";
import { FilterSortControls } from "@/components/liquidations/FilterSortControls";
import { LiquidatablePositionsTable } from "@/components/liquidations/LiquidatablePositionsTable";
import { LiquidateModal } from "@/components/liquidations/LiquidateModal";

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

  const eligible = useEligiblePositions();
  const { view } = usePositionView();

  const rows = useMemo(() => {
    if (eligible.page === undefined || view === undefined) {
      return [];
    }

    const built = eligible.page.items.map((candidate) =>
      buildLiquidationRow(candidate, {
        collateralDecimals,
        debtDecimals,
        collateralUnitPriceScaled: view.collateralUnitPriceScaled,
        debtUnitPriceScaled: view.debtUnitPriceScaled,
        liquidationThresholdBps: view.liquidationThresholdBps,
        bonusBps: view.liquidationBonusBps,
      }),
    );

    return sortRows(
      built.filter((row) => isLiquidatable(row.factorBps)),
      sortKey,
    );
  }, [eligible.page, view, sortKey]);

  if (eligible.isError) {
    return (
      <EmptyState
        title={liquidationsPageContent.unavailableTitle}
        description={liquidationsPageContent.unavailableDescription}
        icon={IconName.Warning}
        action={
          <Button variant={ButtonVariant.Subtle} onClick={eligible.refetch}>
            Try again
          </Button>
        }
      />
    );
  }

  if (eligible.isLoading || view === undefined) {
    return <Skeleton className="h-48 w-full rounded-card" />;
  }

  if (!view.isValued) {
    return (
      <Alert title="Liquidations are paused while the price feed is quiet" tone={BadgeTone.Caution} icon={IconName.Warning}>
        The WETH price has not been reported recently, so no position can be valued and the protocol will refuse to
        settle one. Any attempt would be rejected on chain before it cost you gas. This clears on its own once a fresh
        price arrives.
      </Alert>
    );
  }

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
