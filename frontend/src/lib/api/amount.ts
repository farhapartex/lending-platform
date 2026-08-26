import { malformedResponse } from "@/lib/api/errors";
import type { WireAmount } from "@/lib/api/wire";

export type ApiAmount = {
  amount: bigint;
  decimals: number;
  symbol: string;
};

export function toBaseUnits(raw: string, field: string): bigint {
  if (typeof raw !== "string" || !/^\d+$/.test(raw)) {
    throw malformedResponse(field);
  }

  return BigInt(raw);
}

export function toApiAmount(wire: WireAmount | undefined, field: string): ApiAmount {
  if (wire === undefined || wire === null) {
    throw malformedResponse(field);
  }

  if (!Number.isInteger(wire.decimals) || wire.decimals < 0) {
    throw malformedResponse(`${field}.decimals`);
  }

  if (typeof wire.symbol !== "string") {
    throw malformedResponse(`${field}.symbol`);
  }

  return {
    amount: toBaseUnits(wire.amount, field),
    decimals: wire.decimals,
    symbol: wire.symbol,
  };
}
