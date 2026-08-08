"use client";

import { useState } from "react";
import { AmountValidationCode, AssetSymbol, ButtonSize, CollateralTab, StepState } from "@/lib/enums";
import {
  healthFactorBps,
  healthTier,
  maxSafeCollateralWithdrawal,
  toValueScaled,
} from "@/lib/health";
import { formatTokenAmount, parseTokenAmount } from "@/lib/token";
import { isBlockingValidation, validateCollateralAmount } from "@/lib/validation";
import { assetPrices } from "@/content/protocol";
import { walletBalances, wethAllowance } from "@/content/wallet";
import {
  collateralDecimals,
  collateralDeposited,
  collateralUnitPriceScaled,
  debtDecimals,
  debtOutstanding,
  debtUnitPriceScaled,
  estimatedGasUsd,
  liquidationThresholdBps,
  maxLtvBps,
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
import { CollateralWithdrawGuard } from "@/components/borrow/CollateralWithdrawGuard";
import { HealthImpactPreview } from "@/components/borrow/HealthImpactPreview";

const amountInputId = "collateral-amount";
const validationMessageId = "collateral-amount-validation";

const tabItems = [
  { value: CollateralTab.Deposit, label: "Add" },
  { value: CollateralTab.Withdraw, label: "Withdraw" },
];

const walletWeth = walletBalances[AssetSymbol.Weth];
const debtValueScaled = toValueScaled(debtOutstanding, debtDecimals, debtUnitPriceScaled);
const currentCollateralValue = toValueScaled(collateralDeposited, collateralDecimals, collateralUnitPriceScaled);
const currentFactor = healthFactorBps(currentCollateralValue, debtValueScaled, liquidationThresholdBps);
const currentTier = healthTier(currentFactor);
const safeWithdrawal = maxSafeCollateralWithdrawal(
  collateralDeposited,
  collateralDecimals,
  collateralUnitPriceScaled,
  debtValueScaled,
  maxLtvBps,
);

const messages: Record<AmountValidationCode, string | null> = {
  [AmountValidationCode.None]: null,
  [AmountValidationCode.Empty]: null,
  [AmountValidationCode.InvalidAmount]: "Enter an amount greater than zero.",
  [AmountValidationCode.BelowMinimum]: null,
  [AmountValidationCode.ExceedsWalletBalance]: "That is more WETH than your wallet holds.",
  [AmountValidationCode.ExceedsDeposit]: null,
  [AmountValidationCode.ExceedsAvailableLiquidity]: null,
  [AmountValidationCode.ExceedsCollateral]: "That is more collateral than you have deposited.",
  [AmountValidationCode.ExceedsSafeWithdrawal]:
    "You have this much collateral, but withdrawing it would push your loan past its borrowing limit. Repay some of the loan first.",
  [AmountValidationCode.ExceedsBorrowLimit]: null,
  [AmountValidationCode.ExceedsDebt]: null,
};

export function CollateralPanel() {
  const [tab, setTab] = useState(CollateralTab.Deposit);
  const [rawAmount, setRawAmount] = useState("");

  const isDeposit = tab === CollateralTab.Deposit;
  const amount = parseTokenAmount(rawAmount, collateralDecimals);

  const validation = validateCollateralAmount({
    tab,
    amount,
    walletBalance: walletWeth,
    collateralDeposited,
    maxSafeWithdrawal: safeWithdrawal,
  });

  const hasBlockingError = isBlockingValidation(validation);
  const canSubmit = amount !== null && !hasBlockingError;
  const needsApproval = isDeposit && amount !== null && amount > wethAllowance;

  const nextCollateral =
    amount === null ? collateralDeposited : isDeposit ? collateralDeposited + amount : collateralDeposited - amount;
  const nextCollateralValue = toValueScaled(nextCollateral, collateralDecimals, collateralUnitPriceScaled);
  const nextFactor = healthFactorBps(nextCollateralValue, debtValueScaled, liquidationThresholdBps);
  const nextTier = healthTier(nextFactor);

  const reviewRows = [
    {
      label: isDeposit ? "You add" : "You withdraw",
      value: `${formatTokenAmount(amount ?? 0n, collateralDecimals, 4)} ${AssetSymbol.Weth}`,
    },
    { label: "Platform fee", value: "None" },
    { label: "Estimated network gas", value: estimatedGasUsd },
    {
      label: "Collateral afterwards",
      value: `${formatTokenAmount(nextCollateral, collateralDecimals, 4)} ${AssetSymbol.Weth}`,
      emphasised: true,
    },
  ];

  return (
    <Card className="flex flex-col gap-6 p-6 sm:p-7">
      <div className="flex flex-wrap items-center justify-between gap-3">
        <div className="flex flex-col gap-1">
          <span className="text-sm text-ink-soft">Collateral deposited</span>
          <span className="text-xl font-semibold tracking-tight text-ink tabular-nums">
            {formatTokenAmount(collateralDeposited, collateralDecimals, 4)} {AssetSymbol.Weth}
          </span>
        </div>
        <TabBar items={tabItems} active={tab} label="Add or withdraw collateral" onChange={(value) => {
          setTab(value);
          setRawAmount("");
        }} />
      </div>

      <WalletGate>
        <div className="flex flex-col gap-5">
          <AssetAmountInput
            id={amountInputId}
            label={isDeposit ? "Amount to add" : "Amount to withdraw"}
            symbol={AssetSymbol.Weth}
            decimals={collateralDecimals}
            unitPrice={assetPrices[AssetSymbol.Weth]}
            value={rawAmount}
            onChange={setRawAmount}
            maxAmount={isDeposit ? walletWeth : safeWithdrawal}
            maxLabel={isDeposit ? "Wallet balance" : "Safe to withdraw"}
            invalid={hasBlockingError && validation !== AmountValidationCode.Empty}
            describedBy={validationMessageId}
          />

          <AmountValidationMessage id={validationMessageId} code={validation} messages={messages} />

          {isDeposit ? null : (
            <CollateralWithdrawGuard maxSafeWithdrawal={safeWithdrawal} hasDebt={debtOutstanding > 0n} />
          )}

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
                  label: `Approve ${AssetSymbol.Weth}`,
                  description: "A one-time permission letting the vault move this amount on your behalf.",
                  state: StepState.Active,
                },
                {
                  label: "Add collateral",
                  description: "The transfer into the collateral vault. Your borrowing power rises immediately.",
                  state: StepState.Upcoming,
                },
              ]}
            />
          ) : null}

          <TxStatusTracker status={txFlowStatus} />

          <Button size={ButtonSize.Lg} fullWidth disabled={!canSubmit}>
            {isDeposit ? "Add collateral" : "Withdraw collateral"}
          </Button>
        </div>
      </WalletGate>
    </Card>
  );
}
