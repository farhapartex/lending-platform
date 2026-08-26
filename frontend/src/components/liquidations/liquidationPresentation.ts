import { formatDateTimeUtc, truncateMiddle } from "@/lib/format";
import { formatHealthFactor } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import type { ApiAmount } from "@/lib/api/amount";
import type { Liquidation } from "@/lib/api/liquidationMapper";
import { liquidationsPageContent } from "@/content/liquidations";

export type ReceiptRow = {
  label: string;
  value: string;
  hint?: string;
};

export function formatApiAmount(amount: ApiAmount, maxFractionDigits = 4): string {
  const figure = formatTokenAmount(amount.amount, amount.decimals, maxFractionDigits);

  return amount.symbol === "" ? figure : `${figure} ${amount.symbol}`;
}

export function formatUsdValue(amount: ApiAmount): string {
  return `$${formatTokenAmount(amount.amount, amount.decimals, 2)}`;
}

export function receiptRows(liquidation: Liquidation): ReceiptRow[] {
  const rows: ReceiptRow[] = [
    {
      label: "When",
      value: formatDateTimeUtc(liquidation.timestamp),
    },
    {
      label: "Health factor before",
      value:
        liquidation.healthFactorBeforeBps === null
          ? "Not recorded"
          : formatHealthFactor(liquidation.healthFactorBeforeBps),
      hint: "Anything below 1.00 can be liquidated by anyone.",
    },
    {
      label: "Collateral price at that point",
      value: formatUsdValue(liquidation.triggerPrice),
      hint: "The price that pushed this position below the threshold.",
    },
    {
      label: "Debt repaid",
      value: formatApiAmount(liquidation.debtRepaid, 2),
      hint: "Paid by the liquidator to settle part of the borrower's loan.",
    },
    {
      label: "Collateral seized",
      value: formatApiAmount(liquidation.collateralSeized, 6),
      hint: "Transferred to the liquidator in exchange for repaying the debt.",
    },
    {
      label: "Bonus earned",
      value: formatUsdValue(liquidation.bonusValue),
      hint: "The published reward for resolving an unsafe position.",
    },
  ];

  if (liquidation.shortfallValue.amount > 0n) {
    rows.push({
      label: "Shortfall",
      value: formatUsdValue(liquidation.shortfallValue),
      hint: liquidationsPageContent.shortfallExplainer,
    });
  }

  rows.push(
    {
      label: "Borrower",
      value: truncateMiddle(liquidation.borrower, 10, 8),
    },
    {
      label: "Liquidator",
      value: truncateMiddle(liquidation.liquidator, 10, 8),
    },
    {
      label: "Block",
      value: liquidation.blockNumber.toLocaleString("en-US"),
    },
  );

  return rows;
}
