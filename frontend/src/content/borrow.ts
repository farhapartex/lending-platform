import { AssetSymbol, RepayMode, TxFlowStatus } from "@/lib/enums";
import { assetDecimals, borrowAprRate } from "@/content/protocol";
import { parseTokenAmount } from "@/lib/token";

export {
  collateralAsset,
  collateralDecimals,
  collateralUnitPriceScaled,
  debtAsset,
  debtDecimals,
  debtUnitPriceScaled,
  liquidationBonusBps,
  liquidationThresholdBps,
  maxLtvBps,
  recommendedLtvBps,
} from "@/content/protocol";

const wethDecimals = assetDecimals[AssetSymbol.Weth];
const usdcDecimals = assetDecimals[AssetSymbol.Usdc];

export const collateralDeposited = parseTokenAmount("3.2", wethDecimals) ?? 0n;

export const debtOutstanding = parseTokenAmount("6900", usdcDecimals) ?? 0n;

export const txFlowStatus: TxFlowStatus = TxFlowStatus.Idle;

export const defaultRepayMode: RepayMode = RepayMode.Partial;

export const maxSimulatedDropBps = 5_000n;

export const borrowAprDisplayRate = borrowAprRate;

export const estimatedGasUsd = "$0.58";

export const borrowPageContent = {
  title: "Borrow against WETH",
  description:
    "Post WETH as collateral and borrow USDC against it without selling. Your safety score updates with the WETH price, and you can add collateral or repay at any time to improve it.",
  healthTitle: "Position health",
  healthDescription:
    "This is the single number that decides whether your position is safe. It moves when the WETH price moves, so it is calculated live rather than stored.",
  collateralTitle: "Your collateral",
  collateralDescription: "Add WETH to raise your borrowing power, or take some back if your loan can spare it.",
  debtTitle: "Your loan",
  debtDescription: "Borrow up to your limit, or repay any amount to push your safety score back up.",
  simulatorTitle: "What if the price drops?",
  simulatorDescription:
    "Drag to see what a fall in the WETH price would do to your position, before it happens rather than after.",
} as const;
