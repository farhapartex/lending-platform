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

const usdPriceFormatter = new Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
  minimumFractionDigits: 2,
  maximumFractionDigits: 2,
});

const percentFormatter = new Intl.NumberFormat("en-US", {
  style: "percent",
  minimumFractionDigits: 2,
  maximumFractionDigits: 2,
});

const formatters: Record<ValueFormat, Intl.NumberFormat> = {
  [ValueFormat.UsdCompact]: usdCompactFormatter,
  [ValueFormat.Usd]: usdFormatter,
  [ValueFormat.UsdPrice]: usdPriceFormatter,
  [ValueFormat.Percent]: percentFormatter,
};

export function formatValue(value: number, format: ValueFormat): string {
  return formatters[format].format(value);
}

const secondsPerMinute = 60;
const secondsPerHour = 3600;

export function formatSecondsAgo(seconds: number): string {
  if (seconds < secondsPerMinute) {
    return `${seconds}s ago`;
  }

  if (seconds < secondsPerHour) {
    return `${Math.floor(seconds / secondsPerMinute)}m ago`;
  }

  return `${Math.floor(seconds / secondsPerHour)}h ago`;
}
