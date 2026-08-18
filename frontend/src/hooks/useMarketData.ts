"use client";

import { useReadContract } from "wagmi";

import { useProtocolContracts } from "@/hooks/useProtocolContracts";

export type MarketData = {
  totalSupplied: bigint;
  totalBorrowed: bigint;
  availableLiquidity: bigint;
  utilizationBps: bigint;
  supplyRatePerSecond: bigint;
  borrowRatePerSecond: bigint;
  supplyAprBps: bigint;
  borrowAprBps: bigint;
  supplyIndex: bigint;
  borrowIndex: bigint;
  maxLtvBps: bigint;
  liquidationThresholdBps: bigint;
  liquidationBonusBps: bigint;
  kinkUtilizationBps: bigint;
  reserveFactorBps: bigint;
  minDeposit: bigint;
  accruedReserves: bigint;
  depositsPaused: boolean;
  borrowPaused: boolean;
};

export type MarketDataResult = {
  data: MarketData | undefined;
  isLoading: boolean;
  isError: boolean;
  isUnsupportedChain: boolean;
  refetch: () => void;
};

export const marketDataRefetchInterval = 12_000;

export function useMarketData(): MarketDataResult {
  const { chainId, contracts, isSupported } = useProtocolContracts();

  const query = useReadContract({
    address: contracts?.lens.address,
    abi: contracts?.lens.abi,
    functionName: "marketData",
    chainId,
    query: {
      enabled: isSupported,
      refetchInterval: marketDataRefetchInterval,
      staleTime: marketDataRefetchInterval / 2,
    },
  });

  return {
    data: query.data as MarketData | undefined,
    isLoading: isSupported && query.isPending,
    isError: query.isError,
    isUnsupportedChain: !isSupported,
    refetch: () => {
      void query.refetch();
    },
  };
}
