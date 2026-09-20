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
import { debtDecimals, defaultRepayMode, estimatedGasUsd } from "@/content/borrow";
import { bpsToRatio } from "@/lib/units";
import type { PositionView } from "@/hooks/usePositionView";
import { useProtocolContracts } from "@/hooks/useProtocolContracts";
import { useTxFlow, type TxStep } from "@/hooks/useTxFlow";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { TabBar } from "@/components/ui/TabBar";
import { AssetAmountInput } from "@/components/tx/AssetAmountInput";
import { AmountValidationMessage } from "@/components/tx/AmountValidationMessage";
import { ApprovalStep } from "@/components/tx/ApprovalStep";
import { TxReviewSheet } from "@/components/tx/TxReviewSheet";
import { TxStatusTracker } from "@/components/tx/TxStatusTracker";
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

type DebtPanelProps = {
  view: PositionView;
};

export function DebtPanel({ view }: DebtPanelProps) {
  const { contracts } = useProtocolContracts();
  const walletUsdc = view.walletUsdc;
  const debtOutstanding = view.debtOutstanding;
  const debtValueScaled = view.debtValueScaled;
  const collateralValueScaled = view.collateralValueScaled;
  const currentFactor = view.factorBps;
  const currentTier = view.tier;

  const rawCapacity = borrowCapacity(
    collateralValueScaled,
    debtValueScaled,
    view.maxLtvBps,
    debtDecimals,
    view.debtUnitPriceScaled,
  );
  const capacity = minBigInt(rawCapacity, view.availableLiquidity);
  const isLiquidityConstrained = view.availableLiquidity < rawCapacity;

  const recommendedCapacity = borrowCapacity(
    collateralValueScaled,
    debtValueScaled,
    view.recommendedLtvBps,
    debtDecimals,
    view.debtUnitPriceScaled,
  );
  const isPastRecommended = recommendedCapacity <= 0n;

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
    availableLiquidity: view.availableLiquidity,
    debtOutstanding,
    walletBalance: walletUsdc,
  });

  const hasBlockingError = isBlockingValidation(validation);
  const canSubmit = amount !== null && !hasBlockingError;
  const needsApproval = !isBorrow && amount !== null && amount > view.usdcAllowance;

  const nextDebt = amount === null ? debtOutstanding : isBorrow ? debtOutstanding + amount : debtOutstanding - amount;
  const nextDebtValue = toValueScaled(nextDebt, debtDecimals, view.debtUnitPriceScaled);
  const nextFactor = healthFactorBps(collateralValueScaled, nextDebtValue, view.liquidationThresholdBps);
  const nextTier = healthTier(nextFactor);

  const sliderAmount = typedAmount === null ? 0n : minBigInt(typedAmount, capacity);
  const sliderBps =
    capacity <= 0n || typedAmount === null
      ? 0
      : Math.min(sliderMax, Number((typedAmount * BigInt(sliderMax)) / capacity));

  const handleSliderChange = (valueBps: number) => {
    const nextAmount = (capacity * BigInt(valueBps)) / BigInt(sliderMax);
    setRawAmount(toAmountInputValue(nextAmount, debtDecimals));
  };

  const approval: TxStep | null =
    contracts === null || isBorrow || amount === null || amount <= view.usdcAllowance
      ? null
      : {
          address: contracts.debtToken.address,
          abi: contracts.debtToken.abi,
          functionName: "approve",
          args: [contracts.pool.address, amount],
        };

  const action: TxStep | null =
    contracts === null || amount === null
      ? null
      : {
          address: contracts.controller.address,
          abi: contracts.controller.abi,
          functionName: isBorrow ? "borrow" : isFullRepay ? "repayAll" : "repay",
          args: isFullRepay ? [] : [amount],
        };

  const tx = useTxFlow({
    approval,
    action,
    onConfirmed: () => setRawAmount(""),
  });

  const reviewRows = [
    {
      label: isBorrow ? "You borrow" : "You repay",
      value: `${formatTokenAmount(amount ?? 0n, debtDecimals, 2)} ${AssetSymbol.Usdc}`,
    },
    { label: "Borrow APR", value: formatValue(bpsToRatio(view.borrowAprBps), ValueFormat.Percent) },
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
            tx.reset();
          }}
        />
      </div>

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
            unitPrice={view.debtPrice}
            value={rawAmount}
            onChange={setRawAmount}
            maxAmount={isBorrow ? capacity : minBigInt(debtOutstanding, walletUsdc)}
            maxLabel={isBorrow ? "Available to borrow" : "You owe"}
            invalid={hasBlockingError && validation !== AmountValidationCode.Empty}
            describedBy={validationMessageId}
          />
        )}

        {isBorrow ? (
          <BorrowLimitSlider valueBps={sliderBps} onChange={handleSliderChange} selectedAmount={sliderAmount} />
        ) : null}

        <AmountValidationMessage id={validationMessageId} code={validation} messages={messages} />

        {isBorrow ? (
          <SafetyBufferRecommendation recommendedCapacity={recommendedCapacity} isExceeded={isPastRecommended} />
        ) : null}

        {isBorrow && isLiquidityConstrained ? (
          <InsufficientLiquidityNotice availableLiquidity={view.availableLiquidity} />
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

        <TxStatusTracker
          status={tx.status}
          approvalAsset={AssetSymbol.Usdc}
          approvalSpender="pool"
          confirmedMessage={
            isBorrow
              ? "The USDC is in your wallet and interest has started accruing."
              : "Your debt has gone down and your safety score has gone up."
          }
          errorMessage={tx.error}
        />

        <Button size={ButtonSize.Lg} fullWidth disabled={!canSubmit || tx.isBusy} onClick={tx.submit}>
          {tx.isBusy ? "Working" : isBorrow ? "Borrow" : "Repay"}
        </Button>
      </div>
    </Card>
  );
}
