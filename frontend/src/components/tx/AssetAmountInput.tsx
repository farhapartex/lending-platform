"use client";

import { AssetSymbol, ButtonSize, ButtonVariant, ValueFormat } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { formatValue } from "@/lib/format";
import {
  formatTokenAmount,
  parseTokenAmount,
  sanitizeAmountInput,
  toAmountInputValue,
  tokenAmountToUsd,
} from "@/lib/token";
import { Button } from "@/components/ui/Button";

type AssetAmountInputProps = {
  id: string;
  label: string;
  symbol: AssetSymbol;
  decimals: number;
  unitPrice: number;
  value: string;
  onChange: (value: string) => void;
  maxAmount: bigint;
  maxLabel: string;
  invalid?: boolean;
  describedBy?: string;
};

export function AssetAmountInput({
  id,
  label,
  symbol,
  decimals,
  unitPrice,
  value,
  onChange,
  maxAmount,
  maxLabel,
  invalid = false,
  describedBy,
}: AssetAmountInputProps) {
  const parsed = parseTokenAmount(value, decimals);
  const usdValue = parsed === null ? 0 : tokenAmountToUsd(parsed, decimals, unitPrice);

  return (
    <div className="flex flex-col gap-2">
      <div className="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-1">
        <label htmlFor={id} className="text-sm font-medium text-ink">
          {label}
        </label>
        <span className="text-xs text-ink-soft">
          {maxLabel} {formatTokenAmount(maxAmount, decimals)} {symbol}
        </span>
      </div>

      <div
        className={cn(
          "flex items-center gap-3 rounded-card border bg-surface px-4 py-3 transition-colors focus-within:border-brand-border",
          invalid ? "border-rose-border" : "border-line-strong",
        )}
      >
        <input
          id={id}
          inputMode="decimal"
          autoComplete="off"
          placeholder="0.00"
          value={value}
          aria-invalid={invalid}
          aria-describedby={describedBy}
          onChange={(event) => onChange(sanitizeAmountInput(event.target.value, decimals))}
          className="min-w-0 flex-1 bg-transparent text-2xl font-semibold tracking-tight text-ink tabular-nums outline-none placeholder:text-ink-faint"
        />
        <span className="text-sm font-medium text-ink-soft">{symbol}</span>
        <Button
          variant={ButtonVariant.Subtle}
          size={ButtonSize.Sm}
          onClick={() => onChange(toAmountInputValue(maxAmount, decimals))}
        >
          Max
        </Button>
      </div>

      <span className="text-xs text-ink-faint tabular-nums">≈ {formatValue(usdValue, ValueFormat.UsdPrice)}</span>
    </div>
  );
}
