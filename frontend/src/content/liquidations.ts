import { AssetSymbol, DataStatus, TxFlowStatus } from "@/lib/enums";
import { assetDecimals } from "@/content/protocol";
import { parseTokenAmount } from "@/lib/token";
import type { LiquidationCandidate } from "@/lib/liquidation";

const wethDecimals = assetDecimals[AssetSymbol.Weth];
const usdcDecimals = assetDecimals[AssetSymbol.Usdc];

export const liquidationsDataStatus: DataStatus = DataStatus.Ready;

export const liquidationCandidates: LiquidationCandidate[] = [
  {
    id: "pos-1",
    borrower: "0x3F8a91Cb7D2e4A05B6c8F1d9E7a4B25c0D6e83Fa",
    collateralAmount: parseTokenAmount("1.5", wethDecimals) ?? 0n,
    debtAmount: parseTokenAmount("4200", usdcDecimals) ?? 0n,
  },
  {
    id: "pos-2",
    borrower: "0xB2c47Ae91F5d38E06a7B9c2D4f1E8a35C6b0D97e",
    collateralAmount: parseTokenAmount("0.85", wethDecimals) ?? 0n,
    debtAmount: parseTokenAmount("2400", usdcDecimals) ?? 0n,
  },
  {
    id: "pos-3",
    borrower: "0x91Ea3C7b5D8f2A64c0B1e9D7a3F5c82B4e6D0a1C",
    collateralAmount: parseTokenAmount("12.4", wethDecimals) ?? 0n,
    debtAmount: parseTokenAmount("34800", usdcDecimals) ?? 0n,
  },
  {
    id: "pos-4",
    borrower: "0x6D0b8F3a9C1e7B245d0A6f2E8b9C3d7A15e4F0b2",
    collateralAmount: parseTokenAmount("3.05", wethDecimals) ?? 0n,
    debtAmount: parseTokenAmount("8500", usdcDecimals) ?? 0n,
  },
];

export const liquidatorUsdcBalance = parseTokenAmount("50000", usdcDecimals) ?? 0n;

export const lastRefreshedSecondsAgo = 12;

export const refreshIntervalSeconds = 15;

export const txFlowStatus: TxFlowStatus = TxFlowStatus.Idle;

export const estimatedGasUsd = "$1.24";

export const liquidationsPageContent = {
  title: "Liquidation centre",
  description:
    "Positions that have fallen below a health factor of 1.00 can be resolved by anyone. Repay a borrower's loan, receive their collateral plus a published bonus, and the pool stays solvent.",
  listTitle: "Positions eligible now",
  emptyTitle: "No positions are eligible right now",
  emptyDescription:
    "That is a healthy sign, not an error. Every borrower is currently above the liquidation threshold. This list refreshes on its own.",
  unavailableTitle: "Cannot load eligible positions",
  unavailableDescription:
    "This list is built from indexed contract events, and that service is unreachable right now. Positions are still safe and liquidations can still happen on-chain.",
} as const;
