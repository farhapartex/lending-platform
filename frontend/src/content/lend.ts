import { AssetSymbol, TxFlowStatus } from "@/lib/enums";
import { assetDecimals, availableLiquidity, supplyApyRate } from "@/content/protocol";
import { parseTokenAmount } from "@/lib/token";

const usdcDecimals = assetDecimals[AssetSymbol.Usdc];

export const lendAsset = AssetSymbol.Usdc;

export const lendAssetDecimals = usdcDecimals;

export const secondsPerYear = 31_536_000n;

export const supplyApyBasisPoints = BigInt(Math.round(supplyApyRate * 10_000));

export const depositedPrincipal = parseTokenAmount("25000", usdcDecimals) ?? 0n;

export const accruedInterest = parseTokenAmount("412.8317", usdcDecimals) ?? 0n;

export const depositedBalance = depositedPrincipal + accruedInterest;

export const minimumDeposit = parseTokenAmount("1", usdcDecimals) ?? 0n;

export const poolAvailableLiquidity = parseTokenAmount(String(availableLiquidity), usdcDecimals) ?? 0n;

export const totalSupplied = parseTokenAmount("48240000", usdcDecimals) ?? 0n;

export const txFlowStatus: TxFlowStatus = TxFlowStatus.Idle;

export const estimatedGasUsd = "$0.42";

export const lendPageContent = {
  title: "Lend USDC",
  description:
    "Deposit USDC and start earning immediately. Interest is added to your balance as it accrues, so there is nothing to claim and no lock-up.",
  positionTitle: "Your position",
  positionDescription: "Principal and earned interest, shown separately so you can see exactly what the pool has paid you.",
  actionTitle: "Deposit and withdraw",
  actionDescription: "Move funds in or out at any time, subject to the liquidity currently sitting in the pool.",
  marketTitle: "Market conditions",
  marketDescription:
    "Your rate follows pool utilization. When more of the pool is borrowed, you earn more, and withdrawals get tighter.",
  compoundingNote:
    "Interest compounds into your balance automatically. There is no claim button and no reinvest toggle, because none is needed.",
} as const;
