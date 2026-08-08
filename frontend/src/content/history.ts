import { ActivityKind, AssetSymbol, DataStatus } from "@/lib/enums";
import { assetDecimals } from "@/content/protocol";
import { parseTokenAmount } from "@/lib/token";

const usdcDecimals = assetDecimals[AssetSymbol.Usdc];
const wethDecimals = assetDecimals[AssetSymbol.Weth];

export type HistoryEntry = {
  id: string;
  kind: ActivityKind;
  amount: bigint;
  symbol: AssetSymbol;
  decimals: number;
  timestamp: string;
  blockNumber: number;
  txHash: string;
  healthFactorAfterBps: bigint | null;
};

export const historyDataStatus: DataStatus = DataStatus.Ready;

export const historyAsOf = "2026-08-08T12:00:00Z";

export const historyPageSize = 6;

function usdc(amount: string): bigint {
  return parseTokenAmount(amount, usdcDecimals) ?? 0n;
}

function weth(amount: string): bigint {
  return parseTokenAmount(amount, wethDecimals) ?? 0n;
}

export const historyEntries: HistoryEntry[] = [
  {
    id: "tx-01",
    kind: ActivityKind.Borrow,
    amount: usdc("1500"),
    symbol: AssetSymbol.Usdc,
    decimals: usdcDecimals,
    timestamp: "2026-08-07T14:22:00Z",
    blockNumber: 21_486_912,
    txHash: "0x8f3a1c9d7e2b45a06c8d1f9e4a7b23c05d6e83fa19b7c4e0a2d5f6b8c1e3a90d4",
    healthFactorAfterBps: 12_661n,
  },
  {
    id: "tx-02",
    kind: ActivityKind.CollateralAdded,
    amount: weth("0.75"),
    symbol: AssetSymbol.Weth,
    decimals: wethDecimals,
    timestamp: "2026-08-06T09:05:00Z",
    blockNumber: 21_479_344,
    txHash: "0x2b7c4e0a9d5f1e8a3c06b2d4f7e9a15c8b0d3e6f4a7c92b5d1e08f3a6c4b9d72e",
    healthFactorAfterBps: 15_820n,
  },
  {
    id: "tx-03",
    kind: ActivityKind.Deposit,
    amount: usdc("10000"),
    symbol: AssetSymbol.Usdc,
    decimals: usdcDecimals,
    timestamp: "2026-08-04T18:40:00Z",
    blockNumber: 21_461_207,
    txHash: "0x5d1e08f3a6c4b9d72e2b7c4e0a9d5f1e8a3c06b2d4f7e9a15c8b0d3e6f4a7c92b",
    healthFactorAfterBps: null,
  },
  {
    id: "tx-04",
    kind: ActivityKind.Liquidation,
    amount: usdc("2100"),
    symbol: AssetSymbol.Usdc,
    decimals: usdcDecimals,
    timestamp: "2026-08-01T03:12:00Z",
    blockNumber: 21_439_885,
    txHash: "0xa7c92b5d1e08f3a6c4b9d72e2b7c4e0a9d5f1e8a3c06b2d4f7e9a15c8b0d3e6f4",
    healthFactorAfterBps: 9_800n,
  },
  {
    id: "tx-05",
    kind: ActivityKind.Repay,
    amount: usdc("640"),
    symbol: AssetSymbol.Usdc,
    decimals: usdcDecimals,
    timestamp: "2026-07-29T11:58:00Z",
    blockNumber: 21_419_002,
    txHash: "0x9d5f1e8a3c06b2d4f7e9a15c8b0d3e6f4a7c92b5d1e08f3a6c4b9d72e2b7c4e0a",
    healthFactorAfterBps: 11_240n,
  },
  {
    id: "tx-06",
    kind: ActivityKind.Withdraw,
    amount: usdc("3200"),
    symbol: AssetSymbol.Usdc,
    decimals: usdcDecimals,
    timestamp: "2026-07-24T16:30:00Z",
    blockNumber: 21_383_551,
    txHash: "0xf7e9a15c8b0d3e6f4a7c92b5d1e08f3a6c4b9d72e2b7c4e0a9d5f1e8a3c06b2d4",
    healthFactorAfterBps: null,
  },
  {
    id: "tx-07",
    kind: ActivityKind.CollateralWithdrawn,
    amount: weth("0.4"),
    symbol: AssetSymbol.Weth,
    decimals: wethDecimals,
    timestamp: "2026-07-20T08:14:00Z",
    blockNumber: 21_354_889,
    txHash: "0x3c06b2d4f7e9a15c8b0d3e6f4a7c92b5d1e08f3a6c4b9d72e2b7c4e0a9d5f1e8a",
    healthFactorAfterBps: 13_015n,
  },
  {
    id: "tx-08",
    kind: ActivityKind.Borrow,
    amount: usdc("4800"),
    symbol: AssetSymbol.Usdc,
    decimals: usdcDecimals,
    timestamp: "2026-07-15T13:47:00Z",
    blockNumber: 21_318_446,
    txHash: "0xb0d3e6f4a7c92b5d1e08f3a6c4b9d72e2b7c4e0a9d5f1e8a3c06b2d4f7e9a15c8",
    healthFactorAfterBps: 14_402n,
  },
  {
    id: "tx-09",
    kind: ActivityKind.CollateralAdded,
    amount: weth("2.85"),
    symbol: AssetSymbol.Weth,
    decimals: wethDecimals,
    timestamp: "2026-07-11T10:02:00Z",
    blockNumber: 21_289_120,
    txHash: "0xc4b9d72e2b7c4e0a9d5f1e8a3c06b2d4f7e9a15c8b0d3e6f4a7c92b5d1e08f3a6",
    healthFactorAfterBps: 21_770n,
  },
  {
    id: "tx-10",
    kind: ActivityKind.Deposit,
    amount: usdc("15000"),
    symbol: AssetSymbol.Usdc,
    decimals: usdcDecimals,
    timestamp: "2026-07-06T19:25:00Z",
    blockNumber: 21_252_733,
    txHash: "0x1e08f3a6c4b9d72e2b7c4e0a9d5f1e8a3c06b2d4f7e9a15c8b0d3e6f4a7c92b5d",
    healthFactorAfterBps: null,
  },
  {
    id: "tx-11",
    kind: ActivityKind.Repay,
    amount: usdc("1250"),
    symbol: AssetSymbol.Usdc,
    decimals: usdcDecimals,
    timestamp: "2026-06-28T07:41:00Z",
    blockNumber: 21_195_318,
    txHash: "0x6f4a7c92b5d1e08f3a6c4b9d72e2b7c4e0a9d5f1e8a3c06b2d4f7e9a15c8b0d3e",
    healthFactorAfterBps: 16_930n,
  },
  {
    id: "tx-12",
    kind: ActivityKind.Deposit,
    amount: usdc("2500"),
    symbol: AssetSymbol.Usdc,
    decimals: usdcDecimals,
    timestamp: "2026-06-19T12:09:00Z",
    blockNumber: 21_131_664,
    txHash: "0xd72e2b7c4e0a9d5f1e8a3c06b2d4f7e9a15c8b0d3e6f4a7c92b5d1e08f3a6c4b9",
    healthFactorAfterBps: null,
  },
  {
    id: "tx-13",
    kind: ActivityKind.CollateralAdded,
    amount: weth("1.2"),
    symbol: AssetSymbol.Weth,
    decimals: wethDecimals,
    timestamp: "2026-06-11T15:53:00Z",
    blockNumber: 21_074_902,
    txHash: "0xe0a9d5f1e8a3c06b2d4f7e9a15c8b0d3e6f4a7c92b5d1e08f3a6c4b9d72e2b7c4",
    healthFactorAfterBps: null,
  },
  {
    id: "tx-14",
    kind: ActivityKind.Deposit,
    amount: usdc("5000"),
    symbol: AssetSymbol.Usdc,
    decimals: usdcDecimals,
    timestamp: "2026-06-02T09:18:00Z",
    blockNumber: 21_010_477,
    txHash: "0xa15c8b0d3e6f4a7c92b5d1e08f3a6c4b9d72e2b7c4e0a9d5f1e8a3c06b2d4f7e9",
    healthFactorAfterBps: null,
  },
];

export const historyPageContent = {
  title: "Transaction history",
  description:
    "Every deposit, withdrawal, borrow, repayment, and liquidation for the connected wallet, rebuilt from on-chain events.",
  listTitle: "Your transactions",
  emptyTitle: "No transactions yet",
  emptyDescription: "Once you deposit, borrow, or repay, each action will be listed here with a link to the explorer.",
  noMatchTitle: "No transactions match these filters",
  noMatchDescription: "Try widening the date range or choosing a different type.",
  unavailableTitle: "Cannot load your history",
  unavailableDescription:
    "History is rebuilt from indexed contract events, and that service is unreachable right now. Your funds and positions are unaffected.",
  healthNotApplicable: "Lending actions do not affect a health factor, so none is recorded.",
} as const;
