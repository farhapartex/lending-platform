"use client";

import { HealthTier } from "@/lib/enums";
import { healthTier, priceScale } from "@/lib/health";
import { isNoDebtHealthFactor, NO_DEBT_HEALTH_FACTOR } from "@/lib/units";
import { recommendedLtvBps } from "@/content/protocol";
import { useAccountData, type AccountData } from "@/hooks/useAccountData";
import { useMarketData } from "@/hooks/useMarketData";
import { useOraclePrice } from "@/hooks/useOraclePrice";
import { useTokenBalances } from "@/hooks/useTokenBalances";
import { useWalletState } from "@/hooks/useWalletState";

export type PositionView = {
  collateralDeposited: bigint;
  debtOutstanding: bigint;
  suppliedBalance: bigint;
  totalSupplied: bigint;
  collateralValueScaled: bigint;
  debtValueScaled: bigint;
  collateralUnitPriceScaled: bigint;
  debtUnitPriceScaled: bigint;
  collateralPrice: number;
  debtPrice: number;
  maxLtvBps: bigint;
  liquidationThresholdBps: bigint;
  liquidationBonusBps: bigint;
  recommendedLtvBps: bigint;
  availableLiquidity: bigint;
  minDeposit: bigint;
  maxBorrowable: bigint;
  maxWithdrawableCollateral: bigint;
  walletWeth: bigint;
  walletUsdc: bigint;
  wethAllowance: bigint;
  usdcAllowance: bigint;
  borrowAprBps: bigint;
  supplyAprBps: bigint;
  factorBps: bigint | null;
  tier: HealthTier;
  priceStale: boolean;
  isValued: boolean;
};

export type PositionViewResult = {
  view: PositionView | undefined;
  isLoading: boolean;
  isError: boolean;
  refetch: () => void;
};

function scaledPrice(price: bigint | undefined, decimals: number | undefined): bigint {
  if (price === undefined || decimals === undefined) {
    return 0n;
  }

  return (price * priceScale) / 10n ** BigInt(decimals);
}

const disconnectedAccount: AccountData = {
  supplyShares: 0n,
  supplyAssets: 0n,
  collateralAmount: 0n,
  collateralValue: 0n,
  debtAmount: 0n,
  debtValue: 0n,
  healthFactorBps: NO_DEBT_HEALTH_FACTOR,
  maxBorrowable: 0n,
  maxWithdrawableCollateral: 0n,
  collateralPrice: 0n,
  priceUpdatedAt: 0n,
  isLiquidatable: false,
  priceStale: false,
};

export function usePositionView(): PositionViewResult {
  const { address } = useWalletState();
  const account = useAccountData(address);
  const market = useMarketData();
  const collateralPrice = useOraclePrice("collateral");
  const debtPrice = useOraclePrice("debt");
  const balances = useTokenBalances(address);

  const accountData = address === undefined ? disconnectedAccount : account.data;

  const isError = account.isError || market.isError || collateralPrice.isError || debtPrice.isError;
  const isLoading =
    account.isLoading || market.isLoading || collateralPrice.isLoading || debtPrice.isLoading || balances.isLoading;

  if (
    accountData === undefined ||
    market.data === undefined ||
    collateralPrice.data === undefined ||
    debtPrice.data === undefined
  ) {
    return {
      view: undefined,
      isLoading,
      isError,
      refetch: () => {
        account.refetch();
        market.refetch();
      },
    };
  }

  const isValued = !accountData.priceStale;

  const factorBps =
    !isValued || accountData.debtAmount <= 0n || isNoDebtHealthFactor(accountData.healthFactorBps)
      ? null
      : accountData.healthFactorBps;

  const view: PositionView = {
    collateralDeposited: accountData.collateralAmount,
    debtOutstanding: accountData.debtAmount,
    suppliedBalance: accountData.supplyAssets,
    totalSupplied: market.data.totalSupplied,
    collateralValueScaled: accountData.collateralValue,
    debtValueScaled: accountData.debtValue,
    collateralUnitPriceScaled: scaledPrice(collateralPrice.data.price, collateralPrice.data.decimals),
    debtUnitPriceScaled: scaledPrice(debtPrice.data.price, debtPrice.data.decimals),
    collateralPrice: Number(collateralPrice.data.price) / 10 ** collateralPrice.data.decimals,
    debtPrice: Number(debtPrice.data.price) / 10 ** debtPrice.data.decimals,
    maxLtvBps: market.data.maxLtvBps,
    liquidationThresholdBps: market.data.liquidationThresholdBps,
    liquidationBonusBps: market.data.liquidationBonusBps,
    recommendedLtvBps,
    availableLiquidity: market.data.availableLiquidity,
    minDeposit: market.data.minDeposit,
    maxBorrowable: accountData.maxBorrowable,
    maxWithdrawableCollateral: accountData.maxWithdrawableCollateral,
    walletWeth: balances.data?.walletWeth ?? 0n,
    walletUsdc: balances.data?.walletUsdc ?? 0n,
    wethAllowance: balances.data?.wethAllowance ?? 0n,
    usdcAllowance: balances.data?.usdcAllowance ?? 0n,
    borrowAprBps: market.data.borrowAprBps,
    supplyAprBps: market.data.supplyAprBps,
    factorBps,
    tier: healthTier(factorBps),
    priceStale: accountData.priceStale,
    isValued,
  };

  return {
    view,
    isLoading,
    isError,
    refetch: () => {
      account.refetch();
      market.refetch();
      balances.refetch();
    },
  };
}
