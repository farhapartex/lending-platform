import { getJson } from "@/lib/api/client";
import { toActivityPage, type ActivityPage } from "@/lib/api/transactionMapper";
import type { WireActivity } from "@/lib/api/wire";

export function activityPath(address: string, limit?: number): string {
  const base = `/accounts/${encodeURIComponent(address.toLowerCase())}/activity`;

  if (limit === undefined) {
    return base;
  }

  return `${base}?limit=${encodeURIComponent(String(limit))}`;
}

export async function fetchRecentActivity(
  address: string,
  limit?: number,
  signal?: AbortSignal,
): Promise<ActivityPage> {
  const wire = await getJson<WireActivity>(activityPath(address, limit), signal);

  return toActivityPage(wire);
}
