import { basisPoints, priceScale, toValueScaled } from "@/lib/health";
import { HealthTier } from "@/lib/enums";
import type { FusdCollateral } from "@/content/fusdPreview";

export const fusdSafeTierBps = 15_000n;
export const fusdCautionTierBps = 12_000n;
export const fusdLiquidationBps = 10_000n;

export function collateralValueScaled(collateral: FusdCollateral): bigint {
  return toValueScaled(collateral.deposited, collateral.decimals, collateral.unitPriceScaled);
}

export function totalCollateralValueScaled(collaterals: FusdCollateral[]): bigint {
  return collaterals.reduce((total, entry) => total + collateralValueScaled(entry), 0n);
}

export function fusdHealthFactorBps(
  collateralValue: bigint,
  minted: bigint,
  mintedDecimals: number,
  thresholdBps: bigint,
): bigint | null {
  if (minted <= 0n) {
    return null;
  }

  const debtValue = toValueScaled(minted, mintedDecimals, priceScale);

  if (debtValue <= 0n) {
    return null;
  }

  return (collateralValue * thresholdBps * fusdLiquidationBps) / (basisPoints * debtValue);
}

export function fusdHealthTier(factorBps: bigint | null): HealthTier {
  if (factorBps === null || factorBps >= fusdSafeTierBps) {
    return HealthTier.Safe;
  }

  if (factorBps >= fusdCautionTierBps) {
    return HealthTier.Caution;
  }

  if (factorBps >= fusdLiquidationBps) {
    return HealthTier.AtRisk;
  }

  return HealthTier.Liquidatable;
}

export function mintCapacity(
  collateralValue: bigint,
  minted: bigint,
  mintedDecimals: number,
  thresholdBps: bigint,
): bigint {
  const limitValue = (collateralValue * thresholdBps) / basisPoints;
  const debtValue = toValueScaled(minted, mintedDecimals, priceScale);

  if (limitValue <= debtValue) {
    return 0n;
  }

  return ((limitValue - debtValue) * 10n ** BigInt(mintedDecimals)) / priceScale;
}

export function backingValueScaled(entries: { locked: bigint; decimals: number; unitPriceScaled: bigint }[]): bigint {
  return entries.reduce(
    (total, entry) => total + toValueScaled(entry.locked, entry.decimals, entry.unitPriceScaled),
    0n,
  );
}

export function collateralizationBps(collateralValue: bigint, supply: bigint, supplyDecimals: number): bigint | null {
  const supplyValue = toValueScaled(supply, supplyDecimals, priceScale);

  if (supplyValue <= 0n) {
    return null;
  }

  return (collateralValue * basisPoints) / supplyValue;
}

export function positionCollateralValueScaled(
  collateral: { amount: bigint; decimals: number; unitPriceScaled: bigint }[],
): bigint {
  return collateral.reduce(
    (total, entry) => total + toValueScaled(entry.amount, entry.decimals, entry.unitPriceScaled),
    0n,
  );
}

export function liquidationRewardScaled(debtValue: bigint, bonusBps: bigint): bigint {
  return (debtValue * bonusBps) / basisPoints;
}
