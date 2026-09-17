import { getJson } from "@/lib/api/client";
import { toApiAmount } from "@/lib/api/amount";
import type { WireEligibleList } from "@/lib/api/wire";
import type { LiquidationCandidate } from "@/lib/liquidation";

export type EligiblePositionsPage = {
  items: LiquidationCandidate[];
  asOfBlock: number | null;
  asOfTime: string;
};

export function eligiblePositionsPath(limit?: number): string {
  return limit === undefined ? "/liquidations/eligible" : `/liquidations/eligible?limit=${limit}`;
}

export function toEligiblePositions(wire: WireEligibleList): EligiblePositionsPage {
  return {
    items: wire.items.map((item) => ({
      id: item.borrower,
      borrower: item.borrower,
      collateralAmount: toApiAmount(item.collateral_amount, "collateral_amount").amount,
      debtAmount: toApiAmount(item.debt_amount, "debt_amount").amount,
    })),
    asOfBlock: wire.as_of.block,
    asOfTime: wire.as_of.time,
  };
}

export async function fetchEligiblePositions(
  limit?: number,
  signal?: AbortSignal,
): Promise<EligiblePositionsPage> {
  const wire = await getJson<WireEligibleList>(eligiblePositionsPath(limit), signal);

  return toEligiblePositions(wire);
}
