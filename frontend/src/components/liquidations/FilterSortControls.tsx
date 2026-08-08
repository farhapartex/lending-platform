"use client";

import { LiquidationSortKey } from "@/lib/enums";
import { Dropdown } from "@/components/ui/Dropdown";
import { AutoRefreshIndicator } from "@/components/liquidations/AutoRefreshIndicator";

const sortOptions = [
  { value: LiquidationSortKey.Health, label: "Closest to liquidation" },
  { value: LiquidationSortKey.Size, label: "Largest loan first" },
  { value: LiquidationSortKey.Reward, label: "Highest bonus first" },
];

type FilterSortControlsProps = {
  sortKey: LiquidationSortKey;
  onSortChange: (value: LiquidationSortKey) => void;
  visibleCount: number;
};

export function FilterSortControls({ sortKey, onSortChange, visibleCount }: FilterSortControlsProps) {
  return (
    <div className="flex flex-wrap items-end justify-between gap-4">
      <div className="flex flex-col gap-1.5">
        <span className="text-sm font-medium text-ink">
          {visibleCount} {visibleCount === 1 ? "position" : "positions"} eligible
        </span>
        <AutoRefreshIndicator />
      </div>

      <Dropdown
        id="liquidation-sort"
        label="Sort by"
        value={sortKey}
        options={sortOptions}
        onChange={onSortChange}
        className="w-60"
      />
    </div>
  );
}
