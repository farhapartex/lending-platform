"use client";

import { useQuery } from "@tanstack/react-query";

import { isRetryable } from "@/lib/api/errors";
import { fetchLiquidationHistory } from "@/lib/api/liquidations";
import type { LiquidationHistoryPage } from "@/lib/api/liquidationMapper";
import { queryKeys } from "@/lib/queryKeys";

export const liquidationHistoryPageSize = 10;
export const liquidationHistoryRefetchIntervalMs = 30_000;

const maxRetries = 2;

export type LiquidationHistoryResult = {
  page: LiquidationHistoryPage | undefined;
  isLoading: boolean;
  isError: boolean;
  isFetching: boolean;
  refetch: () => void;
};

export function useLiquidationHistory(
  cursor: string | null,
  market?: string,
  limit: number = liquidationHistoryPageSize,
): LiquidationHistoryResult {
  const query = useQuery({
    queryKey: queryKeys.liquidationHistory(market, cursor, limit),
    queryFn: ({ signal }) =>
      fetchLiquidationHistory({ market, cursor: cursor ?? undefined, limit }, signal),
    staleTime: liquidationHistoryRefetchIntervalMs / 2,
    refetchInterval: liquidationHistoryRefetchIntervalMs,
    retry: (failureCount, error) => isRetryable(error) && failureCount < maxRetries,
  });

  return {
    page: query.data,
    isLoading: query.isPending,
    isError: query.isError,
    isFetching: query.isFetching,
    refetch: () => {
      void query.refetch();
    },
  };
}
