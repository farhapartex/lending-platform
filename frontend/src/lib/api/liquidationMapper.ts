import { malformedResponse } from "@/lib/api/errors";
import { toApiAmount, type ApiAmount } from "@/lib/api/amount";
import { toIndexedAt, type IndexedAt } from "@/lib/api/transactionMapper";
import type { WireLiquidation, WireLiquidationList } from "@/lib/api/wire";

export type Liquidation = {
  id: string;
  borrower: string;
  liquidator: string;
  debtRepaid: ApiAmount;
  collateralSeized: ApiAmount;
  bonusValue: ApiAmount;
  shortfallValue: ApiAmount;
  healthFactorBeforeBps: bigint | null;
  triggerPrice: ApiAmount;
  txHash: string;
  blockNumber: number;
  timestamp: string;
};

export type LiquidationHistoryPage = {
  items: Liquidation[];
  nextCursor: string | null;
  asOf: IndexedAt;
};

function requireText(value: unknown, field: string): string {
  if (typeof value !== "string" || value === "") {
    throw malformedResponse(field);
  }

  return value;
}

function toHealthFactorBps(raw: number | null | undefined): bigint | null {
  if (raw === null || raw === undefined) {
    return null;
  }

  if (!Number.isInteger(raw) || raw < 0) {
    throw malformedResponse("health_factor_before_bps");
  }

  return BigInt(raw);
}

function toBlockNumber(raw: number | undefined): number {
  if (!Number.isInteger(raw) || (raw as number) < 0) {
    throw malformedResponse("block");
  }

  return raw as number;
}

export function toLiquidation(wire: WireLiquidation): Liquidation {
  if (typeof wire?.block_time !== "string" || Number.isNaN(Date.parse(wire.block_time))) {
    throw malformedResponse("block_time");
  }

  return {
    id: requireText(wire.id, "id"),
    borrower: requireText(wire.borrower, "borrower"),
    liquidator: requireText(wire.liquidator, "liquidator"),
    debtRepaid: toApiAmount(wire.debt_repaid, "debt_repaid"),
    collateralSeized: toApiAmount(wire.collateral_seized, "collateral_seized"),
    bonusValue: toApiAmount(wire.bonus_value, "bonus_value"),
    shortfallValue: toApiAmount(wire.shortfall_value, "shortfall_value"),
    healthFactorBeforeBps: toHealthFactorBps(wire.health_factor_before_bps),
    triggerPrice: toApiAmount(wire.trigger_price, "trigger_price"),
    txHash: requireText(wire.tx_hash, "tx_hash"),
    blockNumber: toBlockNumber(wire.block),
    timestamp: wire.block_time,
  };
}

export function toLiquidationHistoryPage(wire: WireLiquidationList): LiquidationHistoryPage {
  if (!Array.isArray(wire?.items)) {
    throw malformedResponse("items");
  }

  if (wire.next_cursor !== null && typeof wire.next_cursor !== "string") {
    throw malformedResponse("next_cursor");
  }

  return {
    items: wire.items.map((item) => toLiquidation(item)),
    nextCursor: wire.next_cursor,
    asOf: toIndexedAt(wire.as_of),
  };
}

export function hasShortfall(liquidation: Liquidation): boolean {
  return liquidation.shortfallValue.amount > 0n;
}
