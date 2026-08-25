"use client";

import { useQuery } from "@tanstack/react-query";

import { isRetryable } from "@/lib/api/errors";
import { fetchRecentActivity } from "@/lib/api/activity";
import type { ActivityPage } from "@/lib/api/transactionMapper";
import { queryKeys } from "@/lib/queryKeys";

export const recentActivityLimit = 5;
export const recentActivityRefetchIntervalMs = 30_000;

const maxRetries = 2;

export type RecentActivityResult = {
  page: ActivityPage | undefined;
  isLoading: boolean;
  isError: boolean;
  isEnabled: boolean;
  refetch: () => void;
};

export function useRecentActivity(
  address: string | undefined,
  limit: number = recentActivityLimit,
): RecentActivityResult {
  const isEnabled = address !== undefined;

  const query = useQuery({
    queryKey: queryKeys.activity(address, limit),
    queryFn: ({ signal }) => fetchRecentActivity(address as string, limit, signal),
    enabled: isEnabled,
    staleTime: recentActivityRefetchIntervalMs / 2,
    refetchInterval: recentActivityRefetchIntervalMs,
    retry: (failureCount, error) => isRetryable(error) && failureCount < maxRetries,
  });

  return {
    page: query.data,
    isLoading: isEnabled && query.isPending,
    isError: query.isError,
    isEnabled,
    refetch: () => {
      void query.refetch();
    },
  };
}
