"use client";

import { useEffect, useState } from "react";
import { AssetSymbol } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { formatTokenAmount } from "@/lib/token";
import { lendAssetDecimals, secondsPerYear, supplyApyBasisPoints } from "@/content/lend";

const basisPointDenominator = 10_000n;
const tickIntervalMs = 1000;

function perSecondAccrual(balance: bigint): bigint {
  return (balance * supplyApyBasisPoints) / (basisPointDenominator * secondsPerYear);
}

type LiveInterestCounterProps = {
  baseInterest: bigint;
  balance: bigint;
  className?: string;
};

export function LiveInterestCounter({ baseInterest, balance, className }: LiveInterestCounterProps) {
  const [elapsedSeconds, setElapsedSeconds] = useState(0);

  useEffect(() => {
    const timer = setInterval(() => {
      setElapsedSeconds((current) => current + 1);
    }, tickIntervalMs);

    return () => clearInterval(timer);
  }, []);

  const projected = baseInterest + perSecondAccrual(balance) * BigInt(elapsedSeconds);

  return (
    <span className={cn("tabular-nums", className)}>
      {formatTokenAmount(projected, lendAssetDecimals, 6)} {AssetSymbol.Usdc}
    </span>
  );
}
