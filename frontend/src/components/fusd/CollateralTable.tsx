import { AssetSymbol, ButtonSize, ButtonVariant, SurfaceElevation, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { scaledValueToUsd } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import { collateralValueScaled } from "@/lib/fusd";
import type { FusdCollateral } from "@/content/fusdPreview";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";

type CollateralTableProps = {
  collaterals: FusdCollateral[];
};

export function CollateralTable({ collaterals }: CollateralTableProps) {
  return (
    <Card elevation={SurfaceElevation.Raised} className="overflow-hidden p-0">
      <table className="w-full border-collapse text-left">
        <thead className="border-b border-line bg-surface-muted">
          <tr>
            <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
              Collateral
            </th>
            <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
              Locked
            </th>
            <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
              In your wallet
            </th>
            <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
              Price
            </th>
            <th scope="col" className="px-5 py-3" />
          </tr>
        </thead>
        <tbody className="divide-y divide-line">
          {collaterals.map((entry) => (
            <tr key={entry.symbol}>
              <td className="px-5 py-4">
                <div className="flex flex-col">
                  <span className="text-sm font-medium text-ink">{entry.symbol}</span>
                  <span className="text-xs text-ink-faint">{entry.name}</span>
                </div>
              </td>
              <td className="px-5 py-4">
                <div className="flex flex-col">
                  <span className="text-sm font-semibold text-ink tabular-nums">
                    {formatTokenAmount(entry.deposited, entry.decimals, 4)} {entry.symbol}
                  </span>
                  <span className="text-xs text-ink-faint tabular-nums">
                    {formatValue(scaledValueToUsd(collateralValueScaled(entry)), ValueFormat.UsdPrice)}
                  </span>
                </div>
              </td>
              <td className="px-5 py-4 text-sm text-ink-soft tabular-nums">
                {formatTokenAmount(entry.walletBalance, entry.decimals, 4)}
              </td>
              <td className="px-5 py-4 text-sm text-ink-soft tabular-nums">
                {formatValue(scaledValueToUsd(entry.unitPriceScaled), ValueFormat.UsdPrice)}
              </td>
              <td className="px-5 py-4">
                <div className="flex justify-end gap-2">
                  <Button size={ButtonSize.Sm} variant={ButtonVariant.Secondary} disabled>
                    Lock
                  </Button>
                  <Button size={ButtonSize.Sm} variant={ButtonVariant.Ghost} disabled>
                    Unlock
                  </Button>
                </div>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </Card>
  );
}

export function collateralSymbols(collaterals: FusdCollateral[]): AssetSymbol[] {
  return collaterals.map((entry) => entry.symbol);
}
