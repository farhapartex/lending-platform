import { AppRoute, AssetSymbol, IconName, ValueFormat } from "@/lib/enums";
import { formatDateTimeUtc, formatValue } from "@/lib/format";
import { formatTokenAmount } from "@/lib/token";
import { collateralDecimals, debtDecimals } from "@/content/borrow";
import type { LiquidationEvent } from "@/content/dashboard";
import { MetricRow } from "@/components/ui/MetricRow";
import { TextLink } from "@/components/ui/TextLink";

type LiquidationReceiptProps = {
  event: LiquidationEvent;
};

export function LiquidationReceipt({ event }: LiquidationReceiptProps) {
  const rows = [
    {
      label: "When",
      value: formatDateTimeUtc(event.timestamp),
      hint: "The moment your health factor fell below 1.00.",
    },
    {
      label: `${AssetSymbol.Weth} price at that point`,
      value: formatValue(event.triggerPrice, ValueFormat.UsdPrice),
      hint: "This price drop is what made your position eligible.",
    },
    {
      label: "Health factor reached",
      value: event.healthFactorAtLiquidation,
      hint: "Anything below 1.00 can be liquidated by anyone.",
    },
    {
      label: "Debt repaid for you",
      value: `${formatTokenAmount(event.debtRepaid, debtDecimals, 2)} ${AssetSymbol.Usdc}`,
      hint: "A liquidator settled this part of your loan.",
    },
    {
      label: "Collateral taken",
      value: `${formatTokenAmount(event.collateralSeized, collateralDecimals, 4)} ${AssetSymbol.Weth}`,
      hint: "Transferred to the liquidator in exchange for repaying your debt.",
    },
    {
      label: "Bonus paid to the liquidator",
      value: `${formatTokenAmount(event.bonusPaid, debtDecimals, 2)} ${AssetSymbol.Usdc}`,
      hint: "The published reward for resolving an unsafe position.",
    },
  ];

  return (
    <div className="flex flex-col gap-4">
      <dl className="divide-y divide-line rounded-card border border-line bg-surface px-5">
        {rows.map((row) => (
          <MetricRow key={row.label} label={row.label} value={row.value} hint={row.hint} />
        ))}
      </dl>

      <TextLink href={AppRoute.LearnLiquidation} trailingIcon={IconName.ArrowRight}>
        Read how liquidation is calculated
      </TextLink>
    </div>
  );
}
