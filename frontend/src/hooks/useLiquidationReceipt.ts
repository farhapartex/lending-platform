"use client";

import { useQuery } from "@tanstack/react-query";

import { isMaskedId } from "@/lib/api/ids";
import { isRetryable } from "@/lib/api/errors";
import { fetchLiquidationReceipt } from "@/lib/api/liquidations";
import type { Liquidation } from "@/lib/api/liquidationMapper";
import { queryKeys } from "@/lib/queryKeys";

export const liquidationReceiptStaleTimeMs = 60_000;

const maxRetries = 2;

export type LiquidationReceiptResult = {
  receipt: Liquidation | undefined;
  isLoading: boolean;
  isError: boolean;
  error: unknown;
  isEnabled: boolean;
  refetch: () => void;
};

export function useLiquidationReceipt(liquidationId: string | null): LiquidationReceiptResult {
  const isEnabled = isMaskedId(liquidationId, "liquidation");

  const query = useQuery({
    queryKey: queryKeys.liquidationReceipt(liquidationId),
    queryFn: ({ signal }) => fetchLiquidationReceipt(liquidationId as string, signal),
    enabled: isEnabled,
    staleTime: liquidationReceiptStaleTimeMs,
    retry: (failureCount, error) => isRetryable(error) && failureCount < maxRetries,
  });

  return {
    receipt: query.data,
    isLoading: isEnabled && query.isPending,
    isError: query.isError,
    error: query.error,
    isEnabled,
    refetch: () => {
      void query.refetch();
    },
  };
}
