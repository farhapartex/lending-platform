"use client";

import { useState } from "react";
import { AssetSymbol, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import {
  applyPriceDrop,
  basisPoints,
  formatHealthFactor,
  healthFactorBps,
  healthTier,
  scaledValueToUsd,
  toValueScaled,
} from "@/lib/health";
import {
  collateralDecimals,
  collateralDeposited,
  collateralUnitPriceScaled,
  debtDecimals,
  debtOutstanding,
  debtUnitPriceScaled,
  liquidationThresholdBps,
  maxSimulatedDropBps,
} from "@/content/borrow";
import { HealthBadge } from "@/components/borrow/HealthBadge";

const sliderId = "price-drop-slider";
const sliderStep = 100;

const sliderClasses =
  "w-full cursor-pointer appearance-none bg-transparent outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-surface [&::-webkit-slider-runnable-track]:h-2.5 [&::-webkit-slider-runnable-track]:rounded-pill [&::-webkit-slider-runnable-track]:bg-canvas-deep [&::-webkit-slider-thumb]:mt-[-0.3rem] [&::-webkit-slider-thumb]:size-5 [&::-webkit-slider-thumb]:appearance-none [&::-webkit-slider-thumb]:rounded-pill [&::-webkit-slider-thumb]:border-2 [&::-webkit-slider-thumb]:border-white [&::-webkit-slider-thumb]:bg-brand [&::-webkit-slider-thumb]:shadow-soft [&::-moz-range-thumb]:size-5 [&::-moz-range-thumb]:rounded-pill [&::-moz-range-thumb]:border-2 [&::-moz-range-thumb]:border-white [&::-moz-range-thumb]:bg-brand [&::-moz-range-track]:h-2.5 [&::-moz-range-track]:rounded-pill [&::-moz-range-track]:bg-canvas-deep";

export function PriceDropSimulatorBody() {
  const [dropBps, setDropBps] = useState(0);

  const simulatedPriceScaled = applyPriceDrop(collateralUnitPriceScaled, BigInt(dropBps));
  const simulatedCollateralValue = toValueScaled(collateralDeposited, collateralDecimals, simulatedPriceScaled);
  const debtValue = toValueScaled(debtOutstanding, debtDecimals, debtUnitPriceScaled);
  const simulatedFactor = healthFactorBps(simulatedCollateralValue, debtValue, liquidationThresholdBps);
  const simulatedTier = healthTier(simulatedFactor);

  return (
    <div className="flex flex-col gap-5">
      <div className="flex flex-col gap-3">
        <div className="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-1">
          <label htmlFor={sliderId} className="text-sm font-medium text-ink">
            {AssetSymbol.Weth} price falls by
          </label>
          <span className="text-sm font-semibold text-ink tabular-nums">
            {formatValue(dropBps / Number(basisPoints), ValueFormat.Percent)}
          </span>
        </div>

        <input
          id={sliderId}
          type="range"
          min={0}
          max={Number(maxSimulatedDropBps)}
          step={sliderStep}
          value={dropBps}
          onChange={(event) => setDropBps(Number(event.target.value))}
          className={sliderClasses}
        />
      </div>

      <dl className="grid gap-4 border-t border-line pt-5 sm:grid-cols-3">
        <div className="flex flex-col gap-1">
          <dt className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">Simulated price</dt>
          <dd className="text-base font-semibold text-ink tabular-nums">
            {formatValue(scaledValueToUsd(simulatedPriceScaled), ValueFormat.UsdPrice)}
          </dd>
        </div>
        <div className="flex flex-col gap-1">
          <dt className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">Collateral value</dt>
          <dd className="text-base font-semibold text-ink tabular-nums">
            {formatValue(scaledValueToUsd(simulatedCollateralValue), ValueFormat.UsdPrice)}
          </dd>
        </div>
        <div className="flex flex-col gap-1.5">
          <dt className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">Resulting health</dt>
          <dd className="flex flex-wrap items-center gap-2">
            <HealthBadge tier={simulatedTier} />
            <span className="text-sm text-ink-soft tabular-nums">{formatHealthFactor(simulatedFactor)}</span>
          </dd>
        </div>
      </dl>
    </div>
  );
}
