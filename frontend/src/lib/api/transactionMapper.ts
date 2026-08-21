import { ActivityKind, ApiErrorCode } from "@/lib/enums";
import { ApiError } from "@/lib/api/errors";
import type { WireAmount, WireTransaction } from "@/lib/api/wire";

export type TransactionDetail = {
  id: string;
  kind: ActivityKind;
  amount: bigint;
  symbol: string;
  decimals: number;
  timestamp: string;
  blockNumber: number;
  txHash: string;
  logIndex: number;
  healthFactorAfterBps: bigint | null;
};

const kindsByWireValue: Record<string, ActivityKind> = {
  deposit: ActivityKind.Deposit,
  withdraw: ActivityKind.Withdraw,
  borrow: ActivityKind.Borrow,
  repay: ActivityKind.Repay,
  collateral_added: ActivityKind.CollateralAdded,
  collateral_withdrawn: ActivityKind.CollateralWithdrawn,
  liquidation: ActivityKind.Liquidation,
};

export const wireValuesByKind: Record<ActivityKind, string> = {
  [ActivityKind.Deposit]: "deposit",
  [ActivityKind.Withdraw]: "withdraw",
  [ActivityKind.Borrow]: "borrow",
  [ActivityKind.Repay]: "repay",
  [ActivityKind.CollateralAdded]: "collateral_added",
  [ActivityKind.CollateralWithdrawn]: "collateral_withdrawn",
  [ActivityKind.Liquidation]: "liquidation",
};

function malformed(field: string): ApiError {
  return new ApiError({
    code: ApiErrorCode.MalformedResponse,
    message: `The server sent a transaction we could not read (${field}).`,
  });
}

function toActivityKind(wireKind: string): ActivityKind {
  const kind = kindsByWireValue[wireKind];

  if (kind === undefined) {
    throw malformed("kind");
  }

  return kind;
}

function toBigInt(raw: string, field: string): bigint {
  if (!/^\d+$/.test(raw)) {
    throw malformed(field);
  }

  return BigInt(raw);
}

function toAmount(wireAmount: WireAmount | undefined): Pick<TransactionDetail, "amount" | "symbol" | "decimals"> {
  if (wireAmount === undefined) {
    throw malformed("amount");
  }

  if (!Number.isInteger(wireAmount.decimals) || wireAmount.decimals < 0) {
    throw malformed("decimals");
  }

  return {
    amount: toBigInt(wireAmount.amount, "amount"),
    symbol: wireAmount.symbol,
    decimals: wireAmount.decimals,
  };
}

function toHealthFactorBps(raw: number | null | undefined): bigint | null {
  if (raw === null || raw === undefined) {
    return null;
  }

  if (!Number.isInteger(raw) || raw < 0) {
    throw malformed("health_factor_after_bps");
  }

  return BigInt(raw);
}

function toWholeNumber(raw: number | undefined, field: string): number {
  if (!Number.isInteger(raw) || (raw as number) < 0) {
    throw malformed(field);
  }

  return raw as number;
}

export function toTransactionDetail(wire: WireTransaction): TransactionDetail {
  if (typeof wire?.id !== "string" || wire.id === "") {
    throw malformed("id");
  }

  if (typeof wire.tx_hash !== "string" || wire.tx_hash === "") {
    throw malformed("tx_hash");
  }

  if (typeof wire.block_time !== "string" || Number.isNaN(Date.parse(wire.block_time))) {
    throw malformed("block_time");
  }

  return {
    id: wire.id,
    kind: toActivityKind(wire.kind),
    ...toAmount(wire.amount),
    timestamp: wire.block_time,
    blockNumber: toWholeNumber(wire.block, "block"),
    txHash: wire.tx_hash,
    logIndex: toWholeNumber(wire.log_index, "log_index"),
    healthFactorAfterBps: toHealthFactorBps(wire.health_factor_after_bps),
  };
}
