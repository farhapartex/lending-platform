export const queryKeys = {
  marketData: (chainId: number | undefined) => ["marketData", chainId] as const,
  accountData: (chainId: number | undefined, address: string | undefined) =>
    ["accountData", chainId, address] as const,
  transaction: (address: string | undefined, transactionId: string | null) =>
    ["transaction", address?.toLowerCase(), transactionId] as const,
  transactions: (address: string | undefined, filters: unknown, cursor: string | null) =>
    ["transactions", address?.toLowerCase(), filters, cursor] as const,
  activity: (address: string | undefined, limit: number) => ["activity", address?.toLowerCase(), limit] as const,
} as const;
