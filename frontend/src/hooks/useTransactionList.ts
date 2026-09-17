"use client";

import { useQuery } from "@tanstack/react-query";

import { isRetryable } from "@/lib/api/errors";
import { fetchTransactionList, type TransactionListFilters } from "@/lib/api/transactions";
import type { TransactionListPage } from "@/lib/api/transactionMapper";
import { queryKeys } from "@/lib/queryKeys";

export const transactionListPageSize = 6;
export const transactionListStaleTimeMs = 15_000;

const maxRetries = 2;

export type TransactionListResult = {
  page: TransactionListPage | undefined;
  isLoading: boolean;
  isError: boolean;
  isEnabled: boolean;
  refetch: () => void;
};

export function useTransactionList(
  address: string | undefined,
  filters: TransactionListFilters = {},
): TransactionListResult {
  const isEnabled = address !== undefined;

  const query = useQuery({
    queryKey: queryKeys.transactionList(address, filters),
    queryFn: ({ signal }) => fetchTransactionList(address as string, filters, signal),
    enabled: isEnabled,
    staleTime: transactionListStaleTimeMs,
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
