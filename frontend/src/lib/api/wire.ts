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
