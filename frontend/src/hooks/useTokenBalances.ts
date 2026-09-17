"use client";

import { useReadContracts } from "wagmi";

import { useProtocolContracts } from "@/hooks/useProtocolContracts";

export type TokenBalances = {
  walletWeth: bigint;
  walletUsdc: bigint;
  wethAllowance: bigint;
  usdcAllowance: bigint;
};

export type TokenBalancesResult = {
  data: TokenBalances | undefined;
  isLoading: boolean;
  isError: boolean;
  refetch: () => void;
};

export const tokenBalancesRefetchInterval = 12_000;

const emptyBalances: TokenBalances = {
  walletWeth: 0n,
  walletUsdc: 0n,
  wethAllowance: 0n,
  usdcAllowance: 0n,
};

export function useTokenBalances(address: string | undefined): TokenBalancesResult {
  const { chainId, contracts, isSupported } = useProtocolContracts();
  const isEnabled = isSupported && address !== undefined && contracts !== null;
  const owner = address as `0x${string}` | undefined;

  const query = useReadContracts({
    contracts:
      contracts === null || owner === undefined
        ? []
        : [
            {
              address: contracts.collateralToken.address,
              abi: contracts.collateralToken.abi,
              functionName: "balanceOf",
              args: [owner],
              chainId,
            },
            {
              address: contracts.debtToken.address,
              abi: contracts.debtToken.abi,
              functionName: "balanceOf",
              args: [owner],
              chainId,
            },
            {
              address: contracts.collateralToken.address,
              abi: contracts.collateralToken.abi,
              functionName: "allowance",
              args: [owner, contracts.vault.address],
              chainId,
            },
            {
              address: contracts.debtToken.address,
              abi: contracts.debtToken.abi,
              functionName: "allowance",
              args: [owner, contracts.pool.address],
              chainId,
            },
          ],
    query: {
      enabled: isEnabled,
      refetchInterval: tokenBalancesRefetchInterval,
      staleTime: tokenBalancesRefetchInterval / 2,
    },
  });

  const rows = query.data;

  const data =
    rows === undefined
      ? undefined
      : {
          ...emptyBalances,
          walletWeth: (rows[0]?.result as bigint | undefined) ?? 0n,
          walletUsdc: (rows[1]?.result as bigint | undefined) ?? 0n,
          wethAllowance: (rows[2]?.result as bigint | undefined) ?? 0n,
          usdcAllowance: (rows[3]?.result as bigint | undefined) ?? 0n,
        };

  return {
    data,
    isLoading: isEnabled && query.isPending,
    isError: query.isError,
    refetch: () => {
      void query.refetch();
    },
  };
}
