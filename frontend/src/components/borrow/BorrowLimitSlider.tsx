"use client";

import { AssetSymbol, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { basisPoints } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import { debtDecimals, maxLtvBps, recommendedLtvBps } from "@/content/borrow";

const sliderId = "borrow-limit-slider";
const sliderMax = 10_000;

type BorrowLimitSliderProps = {
  valueBps: number;
  onChange: (valueBps: number) => void;
  capacity: bigint;
};

export function BorrowLimitSlider({ valueBps, onChange, capacity }: BorrowLimitSliderProps) {
  const recommendedShare = (Number(recommendedLtvBps) / Number(maxLtvBps)) * 100;
  const selectedAmount = (capacity * BigInt(valueBps)) / BigInt(sliderMax);
  const isBeyondRecommended = valueBps > (Number(recommendedLtvBps) / Number(maxLtvBps)) * sliderMax;

  return (
    <div className="flex flex-col gap-3">
      <div className="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-1">
        <label htmlFor={sliderId} className="text-sm font-medium text-ink">
          Share of your available limit
        </label>
        <span className="text-sm font-semibold text-ink tabular-nums">
          {formatTokenAmount(selectedAmount, debtDecimals, 2)} {AssetSymbol.Usdc}
        </span>
      </div>

      <div className="relative">
        <div aria-hidden="true" className="absolute inset-x-0 top-1/2 h-2.5 -translate-y-1/2 overflow-hidden rounded-pill">
          <div className="flex h-full w-full">
            <div className="h-full bg-mint-soft" style={{ width: `${recommendedShare}%` }} />
            <div className="h-full flex-1 bg-amber-soft" />
          </div>
        </div>

        <input
          id={sliderId}
          type="range"
          min={0}
          max={sliderMax}
          step={50}
          value={valueBps}
          onChange={(event) => onChange(Number(event.target.value))}
          className="relative w-full cursor-pointer appearance-none bg-transparent outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-surface [&::-webkit-slider-runnable-track]:h-2.5 [&::-webkit-slider-runnable-track]:rounded-pill [&::-webkit-slider-runnable-track]:bg-transparent [&::-webkit-slider-thumb]:mt-[-0.3rem] [&::-webkit-slider-thumb]:size-5 [&::-webkit-slider-thumb]:appearance-none [&::-webkit-slider-thumb]:rounded-pill [&::-webkit-slider-thumb]:border-2 [&::-webkit-slider-thumb]:border-white [&::-webkit-slider-thumb]:bg-brand [&::-webkit-slider-thumb]:shadow-soft [&::-moz-range-thumb]:size-5 [&::-moz-range-thumb]:rounded-pill [&::-moz-range-thumb]:border-2 [&::-moz-range-thumb]:border-white [&::-moz-range-thumb]:bg-brand [&::-moz-range-track]:h-2.5 [&::-moz-range-track]:bg-transparent"
        />
      </div>

      <div className="flex flex-wrap items-center gap-x-5 gap-y-1.5 text-xs">
        <span className="flex items-center gap-1.5">
          <span aria-hidden="true" className="size-2.5 rounded-pill bg-mint-soft ring-1 ring-mint-border" />
          <span className="text-ink-soft">
            Recommended, up to {formatValue(Number(recommendedLtvBps) / Number(basisPoints), ValueFormat.Percent)} LTV
          </span>
        </span>
        <span className="flex items-center gap-1.5">
          <span aria-hidden="true" className="size-2.5 rounded-pill bg-amber-soft ring-1 ring-amber-border" />
          <span className="text-ink-soft">
            Allowed but thin, up to {formatValue(Number(maxLtvBps) / Number(basisPoints), ValueFormat.Percent)} LTV
          </span>
        </span>
      </div>

      {isBeyondRecommended ? (
        <p className="text-xs leading-relaxed text-amber-ink">
          This is past the buffer we recommend. It is your choice, but a modest fall in the WETH price could put the
          position at risk.
        </p>
      ) : null}
    </div>
  );
}
