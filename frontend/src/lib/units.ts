export const WAD = 1_000_000_000_000_000_000n;
export const PRICE_UNIT = 100_000_000n;
export const FULL_PERCENT_BPS = 10_000n;
export const SECONDS_PER_YEAR = 31_536_000n;

export const PRICE_DECIMALS = 8;
export const WAD_DECIMALS = 18;
export const BPS_DECIMALS = 4;

export const NO_DEBT_HEALTH_FACTOR = (1n << 256n) - 1n;

export function isNoDebtHealthFactor(value: bigint): boolean {
  return value === NO_DEBT_HEALTH_FACTOR;
}

export function bpsToPercent(bps: bigint): number {
  return Number(bps) / Number(FULL_PERCENT_BPS) * 100;
}

export function bpsToRatio(bps: bigint): number {
  return Number(bps) / Number(FULL_PERCENT_BPS);
}

export function ratePerSecondToAprBps(ratePerSecond: bigint): bigint {
  return (ratePerSecond * SECONDS_PER_YEAR * FULL_PERCENT_BPS) / WAD;
}

export function wadToNumber(value: bigint): number {
  return Number(value) / Number(WAD);
}
