export type WireErrorEnvelope = {
  error?: {
    code?: string;
    message?: string;
  };
  request_id?: string;
};

export type WireAmount = {
  amount: string;
  decimals: number;
  symbol: string;
};

export type WireTransaction = {
  id: string;
  kind: string;
  amount: WireAmount;
  health_factor_after_bps: number | null;
  tx_hash: string;
  block: number;
  block_time: string;
  log_index: number;
  status: string;
};

export type WireAsOf = {
  block: number | null;
  time: string;
};

export type WireActivity = {
  items: WireTransaction[];
  as_of: WireAsOf;
};

export type WireLiquidation = {
  id: string;
  borrower: string;
  liquidator: string;
  debt_repaid: WireAmount;
  collateral_seized: WireAmount;
  bonus_value: WireAmount;
  shortfall_value: WireAmount;
  health_factor_before_bps: number | null;
  trigger_price: WireAmount;
  tx_hash: string;
  block: number;
  block_time: string;
};

export type WireLiquidationList = {
  items: WireLiquidation[];
  next_cursor: string | null;
  as_of: WireAsOf;
};
