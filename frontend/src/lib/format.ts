import { ValueFormat } from "@/lib/enums";

const usdCompactFormatter = new Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
  notation: "compact",
  maximumFractionDigits: 1,
});

const usdFormatter = new Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
  maximumFractionDigits: 0,
});

const percentFormatter = new Intl.NumberFormat("en-US", {
  style: "percent",
  minimumFractionDigits: 2,
  maximumFractionDigits: 2,
});

const formatters: Record<ValueFormat, Intl.NumberFormat> = {
  [ValueFormat.UsdCompact]: usdCompactFormatter,
  [ValueFormat.Usd]: usdFormatter,
  [ValueFormat.Percent]: percentFormatter,
};

export function formatValue(value: number, format: ValueFormat): string {
  return formatters[format].format(value);
}
