import { AssetSymbol, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { scaledValueToUsd } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import { collateralDecimals, debtDecimals } from "@/content/protocol";
import type { LiquidationRow } from "@/lib/liquidation";
import { MetricRow } from "@/components/ui/MetricRow";

type LiquidationRewardBreakdownProps = {
  row: LiquidationRow;
};

export function LiquidationRewardBreakdown({ row }: LiquidationRewardBreakdownProps) {
  const rows = [
    {
      label: "You repay",
      value: `${formatTokenAmount(row.debtAmount, debtDecimals, 2)} ${AssetSymbol.Usdc}`,
      hint: "The borrower's full loan. Phase 1 liquidates the whole position.",
    },
    {
      label: "You receive",
      value: `${formatTokenAmount(row.seizedCollateral, collateralDecimals, 4)} ${AssetSymbol.Weth}`,
      hint: "Collateral worth your repayment plus the bonus.",
    },
    {
      label: "Bonus earned",
      value: formatValue(scaledValueToUsd(row.bonusValueScaled), ValueFormat.UsdPrice),
      hint: "Your gross reward before network gas.",
    },
  ];

  return (
    <dl className="divide-y divide-line rounded-card border border-line bg-surface px-5">
      {rows.map((entry) => (
        <MetricRow key={entry.label} label={entry.label} value={entry.value} hint={entry.hint} emphasised />
      ))}
    </dl>
  );
}
