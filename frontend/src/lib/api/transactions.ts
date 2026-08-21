import { getJson } from "@/lib/api/client";
import { toTransactionDetail, type TransactionDetail } from "@/lib/api/transactionMapper";
import type { WireTransaction } from "@/lib/api/wire";

export function transactionDetailPath(address: string, transactionId: string): string {
  return `/accounts/${encodeURIComponent(address.toLowerCase())}/transactions/${encodeURIComponent(transactionId)}`;
}

export async function fetchTransactionDetail(
  address: string,
  transactionId: string,
  signal?: AbortSignal,
): Promise<TransactionDetail> {
  const wire = await getJson<WireTransaction>(transactionDetailPath(address, transactionId), signal);

  return toTransactionDetail(wire);
}
