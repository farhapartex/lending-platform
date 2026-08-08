import { basisPoints, fromValueScaled, healthFactorBps, liquidationBps, toValueScaled } from "@/lib/health";

export function isLiquidatable(factorBps: bigint | null): boolean {
  return factorBps !== null && factorBps < liquidationBps;
}

export function bonusValueScaled(debtValueScaled: bigint, bonusBps: bigint): bigint {
  return (debtValueScaled * bonusBps) / basisPoints;
}

export function seizedCollateralAmount(
  debtValueScaled: bigint,
  bonusBps: bigint,
  collateralUnitPriceScaled: bigint,
  collateralDecimals: number,
): bigint {
  const totalOwed = debtValueScaled + bonusValueScaled(debtValueScaled, bonusBps);

  return fromValueScaled(totalOwed, collateralDecimals, collateralUnitPriceScaled);
}

export type LiquidationCandidate = {
  id: string;
  borrower: string;
  collateralAmount: bigint;
  debtAmount: bigint;
};

export type LiquidationRow = LiquidationCandidate & {
  collateralValueScaled: bigint;
  debtValueScaled: bigint;
  factorBps: bigint | null;
  bonusValueScaled: bigint;
  seizedCollateral: bigint;
  isUnderwater: boolean;
};

type BuildRowParams = {
  collateralDecimals: number;
  debtDecimals: number;
  collateralUnitPriceScaled: bigint;
  debtUnitPriceScaled: bigint;
  liquidationThresholdBps: bigint;
  bonusBps: bigint;
};

export function buildLiquidationRow(candidate: LiquidationCandidate, params: BuildRowParams): LiquidationRow {
  const collateralValueScaled = toValueScaled(
    candidate.collateralAmount,
    params.collateralDecimals,
    params.collateralUnitPriceScaled,
  );
  const debtValueScaled = toValueScaled(candidate.debtAmount, params.debtDecimals, params.debtUnitPriceScaled);
  const factorBps = healthFactorBps(collateralValueScaled, debtValueScaled, params.liquidationThresholdBps);
  const bonus = bonusValueScaled(debtValueScaled, params.bonusBps);
  const seizedCollateral = seizedCollateralAmount(
    debtValueScaled,
    params.bonusBps,
    params.collateralUnitPriceScaled,
    params.collateralDecimals,
  );

  return {
    ...candidate,
    collateralValueScaled,
    debtValueScaled,
    factorBps,
    bonusValueScaled: bonus,
    seizedCollateral,
    isUnderwater: seizedCollateral > candidate.collateralAmount,
  };
}

export function compareBigInt(left: bigint, right: bigint): number {
  if (left < right) {
    return -1;
  }

  return left > right ? 1 : 0;
}
