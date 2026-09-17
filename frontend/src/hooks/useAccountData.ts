"use client";

import { useReadContract } from "wagmi";

import { useProtocolContracts } from "@/hooks/useProtocolContracts";

export type AccountData = {
  supplyShares: bigint;
  supplyAssets: bigint;
  collateralAmount: bigint;
  collateralValue: bigint;
  debtAmount: bigint;
  debtValue: bigint;
  healthFactorBps: bigint;
  maxBorrowable: bigint;
  maxWithdrawableCollateral: bigint;
  collateralPrice: bigint;
  priceUpdatedAt: bigint;
  isLiquidatable: boolean;
  priceStale: boolean;
};

export type AccountDataResult = {
  data: AccountData | undefined;
  isLoading: boolean;
  isError: boolean;
  isEnabled: boolean;
  refetch: () => void;
};

export const accountDataRefetchInterval = 12_000;

export function useAccountData(address: string | undefined): AccountDataResult {
  const { chainId, contracts, isSupported } = useProtocolContracts();
  const isEnabled = isSupported && address !== undefined;

  const query = useReadContract({
    address: contracts?.lens.address,
    abi: contracts?.lens.abi,
    functionName: "accountData",
    args: address === undefined ? undefined : [address as `0x${string}`],
    chainId,
    query: {
      enabled: isEnabled,
      refetchInterval: accountDataRefetchInterval,
      staleTime: accountDataRefetchInterval / 2,
    },
  });

  return {
    data: query.data as AccountData | undefined,
    isLoading: isEnabled && query.isPending,
    isError: query.isError,
    isEnabled,
    refetch: () => {
      void query.refetch();
    },
  };
}
