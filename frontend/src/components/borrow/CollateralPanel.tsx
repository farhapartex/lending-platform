"use client";

import { useState } from "react";
import { AmountValidationCode, AssetSymbol, ButtonSize, CollateralTab, StepState } from "@/lib/enums";
import { healthFactorBps, healthTier, toValueScaled } from "@/lib/health";
import { formatTokenAmount, parseTokenAmount } from "@/lib/token";
import { isBlockingValidation, validateCollateralAmount } from "@/lib/validation";
import { collateralDecimals, estimatedGasUsd } from "@/content/borrow";
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
import { WalletGate } from "@/components/app/WalletGate";
import { CollateralWithdrawGuard } from "@/components/borrow/CollateralWithdrawGuard";
import { HealthImpactPreview } from "@/components/borrow/HealthImpactPreview";

const amountInputId = "collateral-amount";
const validationMessageId = "collateral-amount-validation";

const tabItems = [
  { value: CollateralTab.Deposit, label: "Add" },
  { value: CollateralTab.Withdraw, label: "Withdraw" },
];

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

type CollateralPanelProps = {
  view: PositionView;
};

export function CollateralPanel({ view }: CollateralPanelProps) {
  const [tab, setTab] = useState(CollateralTab.Deposit);
  const [rawAmount, setRawAmount] = useState("");
  const { contracts } = useProtocolContracts();

  const isDeposit = tab === CollateralTab.Deposit;
  const amount = parseTokenAmount(rawAmount, collateralDecimals);
  const safeWithdrawal = view.maxWithdrawableCollateral;

  const validation = validateCollateralAmount({
    tab,
    amount,
    walletBalance: view.walletWeth,
    collateralDeposited: view.collateralDeposited,
    maxSafeWithdrawal: safeWithdrawal,
  });

  const hasBlockingError = isBlockingValidation(validation);
  const canSubmit = amount !== null && !hasBlockingError;
  const needsApproval = isDeposit && amount !== null && amount > view.wethAllowance;

  const nextCollateral =
    amount === null
      ? view.collateralDeposited
      : isDeposit
        ? view.collateralDeposited + amount
        : view.collateralDeposited - amount;
  const nextCollateralValue = toValueScaled(nextCollateral, collateralDecimals, view.collateralUnitPriceScaled);
  const nextFactor = healthFactorBps(nextCollateralValue, view.debtValueScaled, view.liquidationThresholdBps);
  const nextTier = healthTier(nextFactor);

  const approval: TxStep | null =
    contracts === null || !isDeposit || amount === null || amount <= view.wethAllowance
      ? null
      : {
          address: contracts.collateralToken.address,
          abi: contracts.collateralToken.abi,
          functionName: "approve",
          args: [contracts.vault.address, amount],
        };

  const action: TxStep | null =
    contracts === null || amount === null
      ? null
      : {
          address: contracts.vault.address,
          abi: contracts.vault.abi,
          functionName: isDeposit ? "depositCollateral" : "withdrawCollateral",
          args: [amount],
        };

  const tx = useTxFlow({
    approval,
    action,
    onConfirmed: () => setRawAmount(""),
  });

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
            {formatTokenAmount(view.collateralDeposited, collateralDecimals, 4)} {AssetSymbol.Weth}
          </span>
        </div>
        <TabBar items={tabItems} active={tab} label="Add or withdraw collateral" onChange={(value) => {
          setTab(value);
          setRawAmount("");
          tx.reset();
        }} />
      </div>

      <WalletGate>
        <div className="flex flex-col gap-5">
          <AssetAmountInput
            id={amountInputId}
            label={isDeposit ? "Amount to add" : "Amount to withdraw"}
            symbol={AssetSymbol.Weth}
            decimals={collateralDecimals}
            unitPrice={view.collateralPrice}
            value={rawAmount}
            onChange={setRawAmount}
            maxAmount={isDeposit ? view.walletWeth : safeWithdrawal}
            maxLabel={isDeposit ? "Wallet balance" : "Safe to withdraw"}
            invalid={hasBlockingError && validation !== AmountValidationCode.Empty}
            describedBy={validationMessageId}
          />

          <AmountValidationMessage id={validationMessageId} code={validation} messages={messages} />

          {isDeposit ? null : (
            <CollateralWithdrawGuard maxSafeWithdrawal={safeWithdrawal} hasDebt={view.debtOutstanding > 0n} />
          )}

          {canSubmit ? (
            <HealthImpactPreview
              currentFactorBps={view.factorBps}
              currentTier={view.tier}
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

          <TxStatusTracker
            status={tx.status}
            approvalAsset={AssetSymbol.Weth}
            approvalSpender="collateral vault"
            confirmedMessage={
              isDeposit
                ? "Your collateral is in the vault and your borrowing power has gone up."
                : "Your collateral is back in your wallet."
            }
            errorMessage={tx.error}
          />

          <Button size={ButtonSize.Lg} fullWidth disabled={!canSubmit || tx.isBusy} onClick={tx.submit}>
            {tx.isBusy ? "Working" : isDeposit ? "Add collateral" : "Withdraw collateral"}
          </Button>
        </div>
      </WalletGate>
    </Card>
  );
}
