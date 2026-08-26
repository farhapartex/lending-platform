import { getJson } from "@/lib/api/client";
import {
  toLiquidation,
  toLiquidationHistoryPage,
  type Liquidation,
  type LiquidationHistoryPage,
} from "@/lib/api/liquidationMapper";
import type { WireLiquidation, WireLiquidationList } from "@/lib/api/wire";

export type LiquidationHistoryFilters = {
  market?: string;
  cursor?: string;
  limit?: number;
};

export function liquidationHistoryPath(filters: LiquidationHistoryFilters = {}): string {
  const query = new URLSearchParams();

  if (filters.market !== undefined) {
    query.set("market", filters.market);
  }

  if (filters.cursor !== undefined) {
    query.set("cursor", filters.cursor);
  }

  if (filters.limit !== undefined) {
    query.set("limit", String(filters.limit));
  }

  const search = query.toString();

  return search === "" ? "/liquidations/history" : `/liquidations/history?${search}`;
}

export function liquidationReceiptPath(liquidationId: string): string {
  return `/liquidations/${encodeURIComponent(liquidationId)}`;
}

export async function fetchLiquidationHistory(
  filters: LiquidationHistoryFilters = {},
  signal?: AbortSignal,
): Promise<LiquidationHistoryPage> {
  const wire = await getJson<WireLiquidationList>(liquidationHistoryPath(filters), signal);

  return toLiquidationHistoryPage(wire);
}

export async function fetchLiquidationReceipt(
  liquidationId: string,
  signal?: AbortSignal,
): Promise<Liquidation> {
  const wire = await getJson<WireLiquidation>(liquidationReceiptPath(liquidationId), signal);

  return toLiquidation(wire);
}
