import { getJson } from "@/lib/api/client";
import {
  toTransactionDetail,
  toTransactionListPage,
  type TransactionDetail,
  type TransactionListPage,
} from "@/lib/api/transactionMapper";
import type { WireTransaction, WireTransactionList } from "@/lib/api/wire";

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

export type TransactionListFilters = {
  kinds?: string[];
  from?: string;
  to?: string;
  cursor?: string;
  limit?: number;
};

export function transactionListPath(address: string, filters: TransactionListFilters = {}): string {
  const query = new URLSearchParams();

  for (const kind of filters.kinds ?? []) {
    query.append("kind", kind);
  }

  if (filters.from !== undefined) {
    query.set("from", filters.from);
  }

  if (filters.to !== undefined) {
    query.set("to", filters.to);
  }

  if (filters.cursor !== undefined) {
    query.set("cursor", filters.cursor);
  }

  if (filters.limit !== undefined) {
    query.set("limit", String(filters.limit));
  }

  const base = `/accounts/${encodeURIComponent(address.toLowerCase())}/transactions`;
  const search = query.toString();

  return search === "" ? base : `${base}?${search}`;
}

export async function fetchTransactionList(
  address: string,
  filters: TransactionListFilters = {},
  signal?: AbortSignal,
): Promise<TransactionListPage> {
  const wire = await getJson<WireTransactionList>(transactionListPath(address, filters), signal);

  return toTransactionListPage(wire);
}
