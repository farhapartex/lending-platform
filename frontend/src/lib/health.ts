import { HealthTier } from "@/lib/enums";

export const priceScale = 100_000_000n;
export const basisPoints = 10_000n;
export const healthFactorScale = 10_000n;

export const safeTierBps = 15_000n;
export const cautionTierBps = 11_500n;
export const liquidationBps = 10_000n;

export const healthTierLowerBounds: Record<HealthTier, bigint | null> = {
  [HealthTier.Safe]: safeTierBps,
  [HealthTier.Caution]: cautionTierBps,
  [HealthTier.AtRisk]: liquidationBps,
  [HealthTier.Liquidatable]: null,
};

export const healthTierUpperBounds: Record<HealthTier, bigint | null> = {
  [HealthTier.Safe]: null,
  [HealthTier.Caution]: safeTierBps,
  [HealthTier.AtRisk]: cautionTierBps,
  [HealthTier.Liquidatable]: liquidationBps,
};

export const healthTierOrder: HealthTier[] = [
  HealthTier.Safe,
  HealthTier.Caution,
  HealthTier.AtRisk,
  HealthTier.Liquidatable,
];

export function toPriceScaled(price: number): bigint {
  return BigInt(Math.round(price * Number(priceScale)));
}

export function toValueScaled(amount: bigint, decimals: number, unitPriceScaled: bigint): bigint {
  return (amount * unitPriceScaled) / 10n ** BigInt(decimals);
}

export function fromValueScaled(valueScaled: bigint, decimals: number, unitPriceScaled: bigint): bigint {
  if (unitPriceScaled <= 0n) {
    return 0n;
  }

  return (valueScaled * 10n ** BigInt(decimals)) / unitPriceScaled;
}

export function scaledValueToUsd(valueScaled: bigint): number {
  return Number(valueScaled) / Number(priceScale);
}

export function healthFactorBps(
  collateralValueScaled: bigint,
  debtValueScaled: bigint,
  liquidationThresholdBps: bigint,
): bigint | null {
  if (debtValueScaled <= 0n) {
    return null;
  }

  return (collateralValueScaled * liquidationThresholdBps * healthFactorScale) / (basisPoints * debtValueScaled);
}

export function healthTier(factorBps: bigint | null): HealthTier {
  if (factorBps === null || factorBps >= safeTierBps) {
    return HealthTier.Safe;
  }

  if (factorBps >= cautionTierBps) {
    return HealthTier.Caution;
  }

  if (factorBps >= liquidationBps) {
    return HealthTier.AtRisk;
  }

  return HealthTier.Liquidatable;
}

export function formatHealthFactor(factorBps: bigint | null, fractionDigits = 2): string {
  if (factorBps === null) {
    return "No loan";
  }

  return (Number(factorBps) / Number(healthFactorScale)).toFixed(fractionDigits);
}

export function borrowCapacity(
  collateralValueScaled: bigint,
  debtValueScaled: bigint,
  maxLtvBps: bigint,
  debtDecimals: number,
  debtUnitPriceScaled: bigint,
): bigint {
  const limitScaled = (collateralValueScaled * maxLtvBps) / basisPoints;

  if (limitScaled <= debtValueScaled) {
    return 0n;
  }

  return fromValueScaled(limitScaled - debtValueScaled, debtDecimals, debtUnitPriceScaled);
}

export function maxSafeCollateralWithdrawal(
  collateralAmount: bigint,
  collateralDecimals: number,
  collateralUnitPriceScaled: bigint,
  debtValueScaled: bigint,
  maxLtvBps: bigint,
): bigint {
  if (debtValueScaled <= 0n) {
    return collateralAmount;
  }

  if (collateralUnitPriceScaled <= 0n) {
    return 0n;
  }

  const requiredValueScaled = (debtValueScaled * basisPoints + maxLtvBps - 1n) / maxLtvBps;
  const collateralUnit = 10n ** BigInt(collateralDecimals);
  const requiredAmount =
    (requiredValueScaled * collateralUnit + collateralUnitPriceScaled - 1n) / collateralUnitPriceScaled;

  return requiredAmount >= collateralAmount ? 0n : collateralAmount - requiredAmount;
}

export function applyPriceDrop(unitPriceScaled: bigint, dropBps: bigint): bigint {
  return (unitPriceScaled * (basisPoints - dropBps)) / basisPoints;
}
