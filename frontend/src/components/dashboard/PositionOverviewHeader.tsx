import { SurfaceElevation, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { scaledValueToUsd } from "@/lib/health";
import { Card } from "@/components/ui/Card";

type PositionOverviewHeaderProps = {
  suppliedValueScaled: bigint;
  collateralValueScaled: bigint;
  debtValueScaled: bigint;
};

export function PositionOverviewHeader({
  suppliedValueScaled,
  collateralValueScaled,
  debtValueScaled,
}: PositionOverviewHeaderProps) {
  const netValueScaled = suppliedValueScaled + collateralValueScaled - debtValueScaled;

  const entries = [
    { label: "Supplied", value: suppliedValueScaled, hint: "USDC earning interest" },
    { label: "Collateral", value: collateralValueScaled, hint: "WETH backing your loan" },
    { label: "Borrowed", value: debtValueScaled, hint: "USDC you owe" },
  ];

  return (
    <Card elevation={SurfaceElevation.Raised} className="p-6 sm:p-7">
      <dl className="grid gap-6 sm:grid-cols-4">
        {entries.map((entry) => (
          <div key={entry.label} className="flex flex-col gap-1">
            <dt className="text-sm text-ink-soft">{entry.label}</dt>
            <dd className="text-2xl font-semibold tracking-tight text-ink tabular-nums">
              {formatValue(scaledValueToUsd(entry.value), ValueFormat.UsdPrice)}
            </dd>
            <p className="text-xs text-ink-faint">{entry.hint}</p>
          </div>
        ))}

        <div className="flex flex-col gap-1 border-t border-line pt-4 sm:border-l sm:border-t-0 sm:pl-6 sm:pt-0">
          <dt className="text-sm text-ink-soft">Net position</dt>
          <dd className="text-2xl font-semibold tracking-tight text-brand-ink tabular-nums">
            {formatValue(scaledValueToUsd(netValueScaled), ValueFormat.UsdPrice)}
          </dd>
          <p className="text-xs text-ink-faint">Supplied plus collateral, less debt</p>
        </div>
      </dl>
    </Card>
  );
}
