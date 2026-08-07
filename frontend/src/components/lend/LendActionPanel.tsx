"use client";

import { useMemo, useState } from "react";
import {
  AmountValidationCode,
  AssetSymbol,
  ButtonSize,
  LendTab,
  StepState,
  ValueFormat,
} from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { formatTokenAmount, minBigInt, parseTokenAmount, tokenAmountToUsd } from "@/lib/token";
import { isBlockingValidation, validateLendAmount } from "@/lib/validation";
import { assetPrices } from "@/content/protocol";
import { usdcAllowance, walletBalances } from "@/content/wallet";
import {
  depositedBalance,
  estimatedGasUsd,
  lendAssetDecimals,
  minimumDeposit,
  poolAvailableLiquidity,
  txFlowStatus,
} from "@/content/lend";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { TabBar } from "@/components/ui/TabBar";
import { AssetAmountInput } from "@/components/tx/AssetAmountInput";
import { AmountValidationMessage } from "@/components/tx/AmountValidationMessage";
import { ApprovalStep } from "@/components/tx/ApprovalStep";
import { TxReviewSheet } from "@/components/tx/TxReviewSheet";
import { TxStatusTracker } from "@/components/tx/TxStatusTracker";
import { WalletGate } from "@/components/app/WalletGate";
import { WithdrawLiquidityNotice } from "@/components/lend/WithdrawLiquidityNotice";

const amountInputId = "lend-amount";
const validationMessageId = "lend-amount-validation";

const tabItems = [
  { value: LendTab.Deposit, label: "Deposit" },
  { value: LendTab.Withdraw, label: "Withdraw" },
];

const walletBalance = walletBalances[AssetSymbol.Usdc];
const unitPrice = assetPrices[AssetSymbol.Usdc];
const withdrawable = minBigInt(depositedBalance, poolAvailableLiquidity);

const depositMessages: Record<AmountValidationCode, string | null> = {
  [AmountValidationCode.None]: null,
  [AmountValidationCode.Empty]: null,
  [AmountValidationCode.InvalidAmount]: "Enter an amount greater than zero.",
  [AmountValidationCode.BelowMinimum]: `The minimum deposit is ${formatTokenAmount(minimumDeposit, lendAssetDecimals)} ${AssetSymbol.Usdc}. Smaller amounts are rejected to keep dust out of the pool.`,
  [AmountValidationCode.ExceedsWalletBalance]: "That is more than your wallet holds.",
  [AmountValidationCode.ExceedsDeposit]: null,
  [AmountValidationCode.ExceedsAvailableLiquidity]: null,
};

const withdrawMessages: Record<AmountValidationCode, string | null> = {
  [AmountValidationCode.None]: null,
  [AmountValidationCode.Empty]: null,
  [AmountValidationCode.InvalidAmount]: "Enter an amount greater than zero.",
  [AmountValidationCode.BelowMinimum]: null,
  [AmountValidationCode.ExceedsWalletBalance]: null,
  [AmountValidationCode.ExceedsDeposit]: "That is more than you have deposited.",
  [AmountValidationCode.ExceedsAvailableLiquidity]:
    "Your balance covers this, but the pool does not have enough liquid funds right now. This is a pool liquidity limit, not a limit on what you own.",
};

export function LendActionPanel() {
  const [tab, setTab] = useState(LendTab.Deposit);
  const [rawAmount, setRawAmount] = useState("");

  const isDeposit = tab === LendTab.Deposit;
  const amount = parseTokenAmount(rawAmount, lendAssetDecimals);

  const validation = useMemo(
    () =>
      validateLendAmount({
        tab,
        amount,
        walletBalance,
        depositedBalance,
        availableLiquidity: poolAvailableLiquidity,
        minimumDeposit,
      }),
    [tab, amount],
  );

  const needsApproval = isDeposit && amount !== null && amount > usdcAllowance;
  const hasBlockingError = isBlockingValidation(validation);
  const canSubmit = amount !== null && !hasBlockingError;

  const resultingBalance = amount === null ? depositedBalance : isDeposit ? depositedBalance + amount : depositedBalance - amount;

  const reviewRows = [
    {
      label: isDeposit ? "You deposit" : "You withdraw",
      value: `${formatTokenAmount(amount ?? 0n, lendAssetDecimals, 2)} ${AssetSymbol.Usdc}`,
    },
    {
      label: "Platform fee",
      value: "None",
    },
    {
      label: "Estimated network gas",
      value: estimatedGasUsd,
    },
    {
      label: "Balance afterwards",
      value: `${formatTokenAmount(resultingBalance, lendAssetDecimals, 2)} ${AssetSymbol.Usdc}`,
      emphasised: true,
    },
    {
      label: "Value afterwards",
      value: formatValue(tokenAmountToUsd(resultingBalance, lendAssetDecimals, unitPrice), ValueFormat.UsdPrice),
    },
  ];

  return (
    <Card className="flex flex-col gap-6 p-6 sm:p-7">
      <TabBar
        items={tabItems}
        active={tab}
        label="Deposit or withdraw"
        onChange={(value) => {
          setTab(value);
          setRawAmount("");
        }}
      />

      <WalletGate>
        <div className="flex flex-col gap-5">
          <AssetAmountInput
            id={amountInputId}
            label={isDeposit ? "Amount to deposit" : "Amount to withdraw"}
            symbol={AssetSymbol.Usdc}
            decimals={lendAssetDecimals}
            unitPrice={unitPrice}
            value={rawAmount}
            onChange={setRawAmount}
            maxAmount={isDeposit ? walletBalance : withdrawable}
            maxLabel={isDeposit ? "Wallet balance" : "Withdrawable"}
            invalid={hasBlockingError && validation !== AmountValidationCode.Empty}
            describedBy={validationMessageId}
          />

          <AmountValidationMessage
            id={validationMessageId}
            code={validation}
            messages={isDeposit ? depositMessages : withdrawMessages}
          />

          {isDeposit ? null : <WithdrawLiquidityNotice withdrawable={withdrawable} />}

          {canSubmit ? <TxReviewSheet title="Review" rows={reviewRows} /> : null}

          {needsApproval && canSubmit ? (
            <ApprovalStep
              steps={[
                {
                  label: `Approve ${AssetSymbol.Usdc}`,
                  description: "A one-time permission letting the pool move this amount on your behalf.",
                  state: StepState.Active,
                },
                {
                  label: "Deposit",
                  description: "The actual transfer into the pool. Interest starts accruing immediately.",
                  state: StepState.Upcoming,
                },
              ]}
            />
          ) : null}

          <TxStatusTracker status={txFlowStatus} />

          <Button size={ButtonSize.Lg} fullWidth disabled={!canSubmit}>
            {isDeposit ? "Deposit" : "Withdraw"}
          </Button>
        </div>
      </WalletGate>
    </Card>
  );
}
