"use client";

import { useQuery } from "@tanstack/react-query";

import { isRetryable } from "@/lib/api/errors";
import { fetchEligiblePositions, type EligiblePositionsPage } from "@/lib/api/eligible";
import { queryKeys } from "@/lib/queryKeys";

export const eligiblePositionsPageSize = 25;
export const eligiblePositionsRefetchIntervalMs = 15_000;

const maxRetries = 2;

export type EligiblePositionsResult = {
  page: EligiblePositionsPage | undefined;
  isLoading: boolean;
  isError: boolean;
  isFetching: boolean;
  refetch: () => void;
};

export function useEligiblePositions(
  limit: number = eligiblePositionsPageSize,
): EligiblePositionsResult {
  const query = useQuery({
    queryKey: queryKeys.eligiblePositions(limit),
    queryFn: ({ signal }) => fetchEligiblePositions(limit, signal),
    staleTime: eligiblePositionsRefetchIntervalMs / 2,
    refetchInterval: eligiblePositionsRefetchIntervalMs,
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
