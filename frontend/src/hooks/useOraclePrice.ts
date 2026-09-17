"use client";

import { useReadContract } from "wagmi";

import { OracleStatus } from "@/lib/enums";
import { useProtocolContracts } from "@/hooks/useProtocolContracts";

export type OraclePrice = {
  price: bigint;
  decimals: number;
  updatedAt: bigint;
  isStale: boolean;
  isValid: boolean;
};

export type OraclePriceResult = {
  data: OraclePrice | undefined;
  maxPriceAge: number | undefined;
  status: OracleStatus | undefined;
  isLoading: boolean;
  isError: boolean;
  isUnsupportedChain: boolean;
  refetch: () => void;
};

export const oraclePriceRefetchInterval = 12_000;

export function oracleStatusFrom(reading: OraclePrice | undefined): OracleStatus | undefined {
  if (reading === undefined) {
    return undefined;
  }

  if (reading.price === 0n) {
    return OracleStatus.Unavailable;
  }

  return reading.isStale ? OracleStatus.Stale : OracleStatus.Fresh;
}

export type OracleAsset = "collateral" | "debt";

export function useOraclePrice(asset: OracleAsset = "collateral"): OraclePriceResult {
  const { chainId, contracts, isSupported } = useProtocolContracts();
  const token = asset === "debt" ? contracts?.debtToken.address : contracts?.collateralToken.address;

  const priceQuery = useReadContract({
    address: contracts?.oracle.address,
    abi: contracts?.oracle.abi,
    functionName: "readPrice",
    args: token === undefined ? undefined : [token],
    chainId,
    query: {
      enabled: isSupported,
      refetchInterval: oraclePriceRefetchInterval,
      staleTime: oraclePriceRefetchInterval / 2,
    },
  });

  const ageQuery = useReadContract({
    address: contracts?.oracle.address,
    abi: contracts?.oracle.abi,
    functionName: "maxPriceAge",
    chainId,
    query: { enabled: isSupported },
  });

  const data = priceQuery.data as OraclePrice | undefined;
  const maxPriceAge = ageQuery.data as number | undefined;

  return {
    data,
    maxPriceAge,
    status: oracleStatusFrom(data),
    isLoading: isSupported && priceQuery.isPending,
    isError: priceQuery.isError,
    isUnsupportedChain: !isSupported,
    refetch: () => {
      void priceQuery.refetch();
    },
  };
}

export function priceToNumber(reading: OraclePrice): number {
  return Number(reading.price) / 10 ** reading.decimals;
}

export function secondsSince(updatedAt: bigint): number {
  const elapsed = Math.floor(Date.now() / 1000) - Number(updatedAt);

  return elapsed < 0 ? 0 : elapsed;
}
