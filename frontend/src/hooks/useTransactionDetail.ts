"use client";

import { useQuery } from "@tanstack/react-query";

import { isMaskedId } from "@/lib/api/ids";
import { isRetryable } from "@/lib/api/errors";
import { fetchTransactionDetail } from "@/lib/api/transactions";
import type { TransactionDetail } from "@/lib/api/transactionMapper";
import { queryKeys } from "@/lib/queryKeys";

export const transactionDetailStaleTimeMs = 30_000;

const maxRetries = 2;

export type TransactionDetailResult = {
  detail: TransactionDetail | undefined;
  isLoading: boolean;
  isError: boolean;
  error: unknown;
  isEnabled: boolean;
  refetch: () => void;
};

export function useTransactionDetail(
  address: string | undefined,
  transactionId: string | null,
): TransactionDetailResult {
  const isEnabled = address !== undefined && isMaskedId(transactionId, "transaction");

  const query = useQuery({
    queryKey: queryKeys.transaction(address, transactionId),
    queryFn: ({ signal }) => fetchTransactionDetail(address as string, transactionId as string, signal),
    enabled: isEnabled,
    staleTime: transactionDetailStaleTimeMs,
    retry: (failureCount, error) => isRetryable(error) && failureCount < maxRetries,
  });

  return {
    detail: query.data,
    isLoading: isEnabled && query.isPending,
    isError: query.isError,
    error: query.error,
    isEnabled,
    refetch: () => {
      void query.refetch();
    },
  };
}
