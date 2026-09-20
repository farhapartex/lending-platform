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
import { estimatedGasUsd, lendAssetDecimals } from "@/content/lend";
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
import { WithdrawLiquidityNotice } from "@/components/lend/WithdrawLiquidityNotice";

const amountInputId = "lend-amount";
const validationMessageId = "lend-amount-validation";

const tabItems = [
  { value: LendTab.Deposit, label: "Deposit" },
  { value: LendTab.Withdraw, label: "Withdraw" },
];

function depositMessagesFor(minimumDeposit: bigint): Record<AmountValidationCode, string | null> {
  return {
    [AmountValidationCode.None]: null,
    [AmountValidationCode.Empty]: null,
    [AmountValidationCode.InvalidAmount]: "Enter an amount greater than zero.",
    [AmountValidationCode.BelowMinimum]: `The minimum deposit is ${formatTokenAmount(minimumDeposit, lendAssetDecimals)} ${AssetSymbol.Usdc}. Smaller amounts are rejected to keep dust out of the pool.`,
    [AmountValidationCode.ExceedsWalletBalance]: "That is more than your wallet holds.",
    [AmountValidationCode.ExceedsDeposit]: null,
    [AmountValidationCode.ExceedsAvailableLiquidity]: null,
    [AmountValidationCode.ExceedsCollateral]: null,
    [AmountValidationCode.ExceedsSafeWithdrawal]: null,
    [AmountValidationCode.ExceedsBorrowLimit]: null,
    [AmountValidationCode.ExceedsDebt]: null,
  };
}

const withdrawMessages: Record<AmountValidationCode, string | null> = {
  [AmountValidationCode.None]: null,
  [AmountValidationCode.Empty]: null,
  [AmountValidationCode.InvalidAmount]: "Enter an amount greater than zero.",
  [AmountValidationCode.BelowMinimum]: null,
  [AmountValidationCode.ExceedsWalletBalance]: null,
  [AmountValidationCode.ExceedsDeposit]: "That is more than you have deposited.",
  [AmountValidationCode.ExceedsAvailableLiquidity]:
    "Your balance covers this, but the pool does not have enough liquid funds right now. This is a pool liquidity limit, not a limit on what you own.",
  [AmountValidationCode.ExceedsCollateral]: null,
  [AmountValidationCode.ExceedsSafeWithdrawal]: null,
  [AmountValidationCode.ExceedsBorrowLimit]: null,
  [AmountValidationCode.ExceedsDebt]: null,
};

type LendActionPanelProps = {
  view: PositionView;
};

export function LendActionPanel({ view }: LendActionPanelProps) {
  const { contracts } = useProtocolContracts();
  const [tab, setTab] = useState(LendTab.Deposit);
  const [rawAmount, setRawAmount] = useState("");

  const walletBalance = view.walletUsdc;
  const unitPrice = view.debtPrice;
  const depositedBalance = view.suppliedBalance;
  const poolAvailableLiquidity = view.availableLiquidity;
  const minimumDeposit = view.minDeposit;
  const withdrawable = minBigInt(depositedBalance, poolAvailableLiquidity);

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
    [tab, amount, walletBalance, depositedBalance, poolAvailableLiquidity, minimumDeposit],
  );

  const needsApproval = isDeposit && amount !== null && amount > view.usdcAllowance;
  const hasBlockingError = isBlockingValidation(validation);
  const canSubmit = amount !== null && !hasBlockingError;

  const resultingBalance = amount === null ? depositedBalance : isDeposit ? depositedBalance + amount : depositedBalance - amount;

  const approval: TxStep | null =
    contracts === null || !isDeposit || amount === null || amount <= view.usdcAllowance
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
          address: contracts.pool.address,
          abi: contracts.pool.abi,
          functionName: isDeposit ? "deposit" : "withdraw",
          args: [amount],
        };

  const tx = useTxFlow({
    approval,
    action,
    onConfirmed: () => setRawAmount(""),
  });

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
          messages={isDeposit ? depositMessagesFor(minimumDeposit) : withdrawMessages}
        />

        {isDeposit ? null : (
          <WithdrawLiquidityNotice
            withdrawable={withdrawable}
            isLiquidityConstrained={poolAvailableLiquidity < depositedBalance}
          />
        )}

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

        <TxStatusTracker
          status={tx.status}
          approvalAsset={AssetSymbol.Usdc}
          approvalSpender="pool"
          confirmedMessage={
            isDeposit
              ? "Your deposit is in the pool and has started earning interest."
              : "Your USDC is back in your wallet."
          }
          errorMessage={tx.error}
        />

        <Button size={ButtonSize.Lg} fullWidth disabled={!canSubmit || tx.isBusy} onClick={tx.submit}>
          {isDeposit ? "Deposit" : "Withdraw"}
        </Button>
      </div>
    </Card>
  );
}
