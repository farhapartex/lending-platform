import { AssetSymbol, ButtonSize, ButtonVariant, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { formatHealthFactor, healthTier, scaledValueToUsd } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import { collateralDecimals, debtDecimals } from "@/content/protocol";
import type { LiquidationRow } from "@/lib/liquidation";
import { AddressDisplay } from "@/components/ui/AddressDisplay";
import { Button } from "@/components/ui/Button";
import { HealthBadge } from "@/components/borrow/HealthBadge";

type LiquidatePositionRowProps = {
  row: LiquidationRow;
  onSelect: (row: LiquidationRow) => void;
};

export function LiquidatePositionRow({ row, onSelect }: LiquidatePositionRowProps) {
  return (
    <tr className="border-t border-line align-middle">
      <td className="px-5 py-4">
        <AddressDisplay address={row.borrower} />
      </td>

      <td className="px-5 py-4">
        <div className="flex flex-col gap-0.5">
          <span className="text-sm font-medium text-ink tabular-nums">
            {formatTokenAmount(row.collateralAmount, collateralDecimals, 4)} {AssetSymbol.Weth}
          </span>
          <span className="text-xs text-ink-faint tabular-nums">
            {formatValue(scaledValueToUsd(row.collateralValueScaled), ValueFormat.UsdPrice)}
          </span>
        </div>
      </td>

      <td className="px-5 py-4">
        <span className="text-sm font-medium text-ink tabular-nums">
          {formatTokenAmount(row.debtAmount, debtDecimals, 2)} {AssetSymbol.Usdc}
        </span>
      </td>

      <td className="px-5 py-4">
        <div className="flex flex-col items-start gap-1.5">
          <HealthBadge tier={healthTier(row.factorBps)} />
          <span className="text-xs text-ink-soft tabular-nums">{formatHealthFactor(row.factorBps, 4)}</span>
        </div>
      </td>

      <td className="px-5 py-4">
        <span className="text-sm font-semibold text-mint-ink tabular-nums">
          {formatValue(scaledValueToUsd(row.bonusValueScaled), ValueFormat.UsdPrice)}
        </span>
      </td>

      <td className="px-5 py-4 text-right">
        <Button variant={ButtonVariant.Secondary} size={ButtonSize.Sm} onClick={() => onSelect(row)}>
          Liquidate
        </Button>
      </td>
    </tr>
  );
}
