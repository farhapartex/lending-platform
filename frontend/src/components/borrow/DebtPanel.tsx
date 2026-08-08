"use client";

import { useState } from "react";
import {
  AmountValidationCode,
  AssetSymbol,
  ButtonSize,
  DebtTab,
  RepayMode,
  StepState,
  ValueFormat,
} from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { borrowCapacity, healthFactorBps, healthTier, toValueScaled } from "@/lib/health";
import { formatTokenAmount, minBigInt, parseTokenAmount, toAmountInputValue } from "@/lib/token";
import { isBlockingValidation, validateDebtAmount } from "@/lib/validation";
import { assetPrices } from "@/content/protocol";
import { usdcAllowance, walletBalances } from "@/content/wallet";
import { poolAvailableLiquidity } from "@/content/lend";
import {
  borrowAprDisplayRate,
  collateralDecimals,
  collateralDeposited,
  collateralUnitPriceScaled,
  debtDecimals,
  debtOutstanding,
  debtUnitPriceScaled,
  defaultRepayMode,
  estimatedGasUsd,
  liquidationThresholdBps,
  maxLtvBps,
  recommendedLtvBps,
  txFlowStatus,
} from "@/content/borrow";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { TabBar } from "@/components/ui/TabBar";
import { AssetAmountInput } from "@/components/tx/AssetAmountInput";
import { AmountValidationMessage } from "@/components/tx/AmountValidationMessage";
import { ApprovalStep } from "@/components/tx/ApprovalStep";
import { TxReviewSheet } from "@/components/tx/TxReviewSheet";
import { TxStatusTracker } from "@/components/tx/TxStatusTracker";
import { WalletGate } from "@/components/app/WalletGate";
import { BorrowLimitSlider } from "@/components/borrow/BorrowLimitSlider";
import { HealthImpactPreview } from "@/components/borrow/HealthImpactPreview";
import { InsufficientLiquidityNotice } from "@/components/borrow/InsufficientLiquidityNotice";
import { MaxBorrowExplainer } from "@/components/borrow/MaxBorrowExplainer";
import { SafetyBufferRecommendation } from "@/components/borrow/SafetyBufferRecommendation";

const amountInputId = "debt-amount";
const validationMessageId = "debt-amount-validation";
const sliderMax = 10_000;

const tabItems = [
  { value: DebtTab.Borrow, label: "Borrow" },
  { value: DebtTab.Repay, label: "Repay" },
];

const repayModeItems = [
  { value: RepayMode.Partial, label: "Part of it" },
  { value: RepayMode.Full, label: "All of it" },
];

const walletUsdc = walletBalances[AssetSymbol.Usdc];
const collateralValueScaled = toValueScaled(collateralDeposited, collateralDecimals, collateralUnitPriceScaled);
const debtValueScaled = toValueScaled(debtOutstanding, debtDecimals, debtUnitPriceScaled);
const currentFactor = healthFactorBps(collateralValueScaled, debtValueScaled, liquidationThresholdBps);
const currentTier = healthTier(currentFactor);

const rawCapacity = borrowCapacity(
  collateralValueScaled,
  debtValueScaled,
  maxLtvBps,
  debtDecimals,
  debtUnitPriceScaled,
);
const capacity = minBigInt(rawCapacity, poolAvailableLiquidity);
const isLiquidityConstrained = poolAvailableLiquidity < rawCapacity;

const recommendedCapacity = borrowCapacity(
  collateralValueScaled,
  debtValueScaled,
  recommendedLtvBps,
  debtDecimals,
  debtUnitPriceScaled,
);
const isPastRecommended = recommendedCapacity <= 0n;

const messages: Record<AmountValidationCode, string | null> = {
  [AmountValidationCode.None]: null,
  [AmountValidationCode.Empty]: null,
  [AmountValidationCode.InvalidAmount]: "Enter an amount greater than zero.",
  [AmountValidationCode.BelowMinimum]: null,
  [AmountValidationCode.ExceedsWalletBalance]: "That is more USDC than your wallet holds to repay with.",
  [AmountValidationCode.ExceedsDeposit]: null,
  [AmountValidationCode.ExceedsAvailableLiquidity]:
    "Your collateral supports this, but the pool does not have enough liquid USDC right now. This is a pool liquidity limit, not a collateral limit.",
  [AmountValidationCode.ExceedsCollateral]: null,
  [AmountValidationCode.ExceedsSafeWithdrawal]: null,
  [AmountValidationCode.ExceedsBorrowLimit]:
    "That is more than your collateral allows you to borrow. Add collateral to raise the limit.",
  [AmountValidationCode.ExceedsDebt]: "That is more than you currently owe.",
};

export function DebtPanel() {
  const [tab, setTab] = useState(DebtTab.Borrow);
  const [repayMode, setRepayMode] = useState(defaultRepayMode);
  const [rawAmount, setRawAmount] = useState("");

  const isBorrow = tab === DebtTab.Borrow;
  const isFullRepay = !isBorrow && repayMode === RepayMode.Full;

  const typedAmount = parseTokenAmount(rawAmount, debtDecimals);
  const amount = isFullRepay ? debtOutstanding : typedAmount;

  const validation = validateDebtAmount({
    tab,
    amount,
    borrowCapacity: capacity,
    availableLiquidity: poolAvailableLiquidity,
    debtOutstanding,
    walletBalance: walletUsdc,
  });

  const hasBlockingError = isBlockingValidation(validation);
  const canSubmit = amount !== null && !hasBlockingError;
  const needsApproval = !isBorrow && amount !== null && amount > usdcAllowance;

  const nextDebt = amount === null ? debtOutstanding : isBorrow ? debtOutstanding + amount : debtOutstanding - amount;
  const nextDebtValue = toValueScaled(nextDebt, debtDecimals, debtUnitPriceScaled);
  const nextFactor = healthFactorBps(collateralValueScaled, nextDebtValue, liquidationThresholdBps);
  const nextTier = healthTier(nextFactor);

  const sliderBps = capacity <= 0n || typedAmount === null ? 0 : Number((typedAmount * BigInt(sliderMax)) / capacity);

  const handleSliderChange = (valueBps: number) => {
    const nextAmount = (capacity * BigInt(valueBps)) / BigInt(sliderMax);
    setRawAmount(toAmountInputValue(nextAmount, debtDecimals));
  };

  const reviewRows = [
    {
      label: isBorrow ? "You borrow" : "You repay",
      value: `${formatTokenAmount(amount ?? 0n, debtDecimals, 2)} ${AssetSymbol.Usdc}`,
    },
    { label: "Borrow APR", value: formatValue(borrowAprDisplayRate, ValueFormat.Percent) },
    { label: "Estimated network gas", value: estimatedGasUsd },
    {
      label: "Debt afterwards",
      value: `${formatTokenAmount(nextDebt, debtDecimals, 2)} ${AssetSymbol.Usdc}`,
      emphasised: true,
    },
  ];

  return (
    <Card className="flex flex-col gap-6 p-6 sm:p-7">
      <div className="flex flex-wrap items-center justify-between gap-3">
        <div className="flex flex-col gap-1">
          <span className="text-sm text-ink-soft">Currently borrowed</span>
          <span className="text-xl font-semibold tracking-tight text-ink tabular-nums">
            {formatTokenAmount(debtOutstanding, debtDecimals, 2)} {AssetSymbol.Usdc}
          </span>
        </div>
        <TabBar
          items={tabItems}
          active={tab}
          label="Borrow or repay"
          onChange={(value) => {
            setTab(value);
            setRawAmount("");
          }}
        />
      </div>

      <WalletGate>
        <div className="flex flex-col gap-5">
          {isBorrow ? <MaxBorrowExplainer capacity={capacity} /> : null}

          {!isBorrow ? (
            <TabBar
              items={repayModeItems}
              active={repayMode}
              label="How much to repay"
              onChange={(value) => {
                setRepayMode(value);
                setRawAmount("");
              }}
            />
          ) : null}

          {isFullRepay ? (
            <div className="flex flex-col gap-1.5 rounded-card border border-line bg-surface-muted p-4">
              <span className="text-sm text-ink-soft">Repaying your full balance</span>
              <span className="text-2xl font-semibold tracking-tight text-ink tabular-nums">
                {formatTokenAmount(debtOutstanding, debtDecimals, 2)} {AssetSymbol.Usdc}
              </span>
            </div>
          ) : (
            <AssetAmountInput
              id={amountInputId}
              label={isBorrow ? "Amount to borrow" : "Amount to repay"}
              symbol={AssetSymbol.Usdc}
              decimals={debtDecimals}
              unitPrice={assetPrices[AssetSymbol.Usdc]}
              value={rawAmount}
              onChange={setRawAmount}
              maxAmount={isBorrow ? capacity : minBigInt(debtOutstanding, walletUsdc)}
              maxLabel={isBorrow ? "Available to borrow" : "You owe"}
              invalid={hasBlockingError && validation !== AmountValidationCode.Empty}
              describedBy={validationMessageId}
            />
          )}

          {isBorrow ? (
            <BorrowLimitSlider valueBps={sliderBps} onChange={handleSliderChange} capacity={capacity} />
          ) : null}

          <AmountValidationMessage id={validationMessageId} code={validation} messages={messages} />

          {isBorrow ? (
            <SafetyBufferRecommendation recommendedCapacity={recommendedCapacity} isExceeded={isPastRecommended} />
          ) : null}

          {isBorrow && isLiquidityConstrained ? (
            <InsufficientLiquidityNotice availableLiquidity={poolAvailableLiquidity} />
          ) : null}

          {canSubmit ? (
            <HealthImpactPreview
              currentFactorBps={currentFactor}
              currentTier={currentTier}
              nextFactorBps={nextFactor}
              nextTier={nextTier}
            />
          ) : null}

          {canSubmit ? <TxReviewSheet title="Review" rows={reviewRows} /> : null}

          {needsApproval && canSubmit ? (
            <ApprovalStep
              steps={[
                {
                  label: `Approve ${AssetSymbol.Usdc}`,
                  description: "A one-time permission letting the pool collect your repayment.",
                  state: StepState.Active,
                },
                {
                  label: "Repay",
                  description: "The transfer that reduces your debt and lifts your safety score.",
                  state: StepState.Upcoming,
                },
              ]}
            />
          ) : null}

          <TxStatusTracker status={txFlowStatus} />

          <Button size={ButtonSize.Lg} fullWidth disabled={!canSubmit}>
            {isBorrow ? "Borrow" : "Repay"}
          </Button>
        </div>
      </WalletGate>
    </Card>
  );
}
