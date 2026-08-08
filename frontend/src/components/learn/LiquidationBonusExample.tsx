import { AssetSymbol, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import {
  basisPoints,
  formatHealthFactor,
  healthFactorBps,
  scaledValueToUsd,
  toValueScaled,
} from "@/lib/health";
import { bonusValueScaled, seizedCollateralAmount } from "@/lib/liquidation";
import { formatTokenAmount, parseTokenAmount } from "@/lib/token";
import {
  collateralDecimals,
  collateralUnitPriceScaled,
  debtDecimals,
  debtUnitPriceScaled,
  liquidationBonusBps,
  liquidationThresholdBps,
} from "@/content/protocol";
import { learnIndexContent } from "@/content/learn";
import { MetricRow } from "@/components/ui/MetricRow";

const exampleCollateral = parseTokenAmount("2", collateralDecimals) ?? 0n;
const exampleDebt = parseTokenAmount("5400", debtDecimals) ?? 0n;

export function LiquidationBonusExample() {
  const collateralValueScaled = toValueScaled(exampleCollateral, collateralDecimals, collateralUnitPriceScaled);
  const debtValueScaled = toValueScaled(exampleDebt, debtDecimals, debtUnitPriceScaled);
  const factorBps = healthFactorBps(collateralValueScaled, debtValueScaled, liquidationThresholdBps);
  const bonus = bonusValueScaled(debtValueScaled, liquidationBonusBps);
  const seized = seizedCollateralAmount(
    debtValueScaled,
    liquidationBonusBps,
    collateralUnitPriceScaled,
    collateralDecimals,
  );
  const collateralLeft = exampleCollateral > seized ? exampleCollateral - seized : 0n;

  const rows = [
    {
      label: "Borrower's collateral",
      value: `${formatTokenAmount(exampleCollateral, collateralDecimals, 4)} ${AssetSymbol.Weth}`,
      hint: `Worth ${formatValue(scaledValueToUsd(collateralValueScaled), ValueFormat.UsdPrice)} at the current price.`,
    },
    {
      label: "Borrower's loan",
      value: `${formatTokenAmount(exampleDebt, debtDecimals, 2)} ${AssetSymbol.Usdc}`,
      hint: "What they owe the pool.",
    },
    {
      label: "Health factor",
      value: formatHealthFactor(factorBps, 4),
      hint: `Collateral value times the ${formatValue(Number(liquidationThresholdBps) / Number(basisPoints), ValueFormat.Percent)} liquidation threshold, divided by the loan.`,
    },
    {
      label: "Liquidator repays",
      value: `${formatTokenAmount(exampleDebt, debtDecimals, 2)} ${AssetSymbol.Usdc}`,
      hint: "The full loan, since this phase closes the whole position.",
    },
    {
      label: "Liquidator receives",
      value: `${formatTokenAmount(seized, collateralDecimals, 4)} ${AssetSymbol.Weth}`,
      hint: `Collateral worth the repayment plus the ${formatValue(Number(liquidationBonusBps) / Number(basisPoints), ValueFormat.Percent)} bonus.`,
    },
    {
      label: "Liquidator's gross profit",
      value: formatValue(scaledValueToUsd(bonus), ValueFormat.UsdPrice),
      hint: "Before network gas, which they pay either way.",
    },
    {
      label: "Borrower keeps",
      value: `${formatTokenAmount(collateralLeft, collateralDecimals, 4)} ${AssetSymbol.Weth}`,
      hint: "Whatever collateral was not needed to cover the loan and bonus.",
    },
  ];

  return (
    <div className="flex flex-col gap-3">
      <p className="text-sm leading-relaxed text-ink-soft">{learnIndexContent.exampleNote}</p>
      <dl className="divide-y divide-line rounded-card border border-line bg-surface px-5">
        {rows.map((row) => (
          <MetricRow key={row.label} label={row.label} value={row.value} hint={row.hint} />
        ))}
      </dl>
    </div>
  );
}
