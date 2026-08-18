import { AssetSymbol, DataStatus } from "@/lib/enums";
import { bpsToPercent } from "@/lib/units";
import type { MarketData } from "@/hooks/useMarketData";

export const collateralDecimals = 18;
export const debtDecimals = 6;

export type QueryLikeState = {
  isLoading: boolean;
  isError: boolean;
  hasData: boolean;
};

export function dataStatusFrom(state: QueryLikeState): DataStatus {
  if (state.isError) {
    return DataStatus.Unavailable;
  }

  if (state.isLoading || !state.hasData) {
    return DataStatus.Loading;
  }

  return DataStatus.Ready;
}

export function tokenAmountToNumber(value: bigint, decimals: number): number {
  return Number(value) / 10 ** decimals;
}

export function debtAmountToNumber(value: bigint): number {
  return tokenAmountToNumber(value, debtDecimals);
}

export function collateralAmountToNumber(value: bigint): number {
  return tokenAmountToNumber(value, collateralDecimals);
}

export function supplyAprPercent(market: MarketData): number {
  return bpsToPercent(market.supplyAprBps);
}

export function borrowAprPercent(market: MarketData): number {
  return bpsToPercent(market.borrowAprBps);
}

export function utilizationPercent(market: MarketData): number {
  return bpsToPercent(market.utilizationBps);
}

export function kinkPercent(market: MarketData): number {
  return bpsToPercent(market.kinkUtilizationBps);
}

export function maxLtvPercent(market: MarketData): number {
  return bpsToPercent(market.maxLtvBps);
}

export function liquidationThresholdPercent(market: MarketData): number {
  return bpsToPercent(market.liquidationThresholdBps);
}

export function liquidationBonusPercent(market: MarketData): number {
  return bpsToPercent(market.liquidationBonusBps);
}

export function reserveFactorPercent(market: MarketData): number {
  return bpsToPercent(market.reserveFactorBps);
}

export function isMarketPaused(market: MarketData): boolean {
  return market.depositsPaused || market.borrowPaused;
}

export function decimalsFor(symbol: AssetSymbol): number {
  return symbol === AssetSymbol.Weth ? collateralDecimals : debtDecimals;
}
