"use client";

import {
  AssetSymbol,
  BadgeTone,
  ButtonSize,
  ButtonVariant,
  IconName,
  StepState,
  WalletGatePurpose,
  WalletStatus,
} from "@/lib/enums";
import { useWalletState } from "@/hooks/useWalletState";
import { WalletGate } from "@/components/app/WalletGate";
import { formatTokenAmount } from "@/lib/token";
import { debtDecimals } from "@/content/protocol";
import { estimatedGasUsd, liquidatorUsdcBalance, txFlowStatus } from "@/content/liquidations";
import type { LiquidationRow } from "@/lib/liquidation";
import { Alert } from "@/components/ui/Alert";
import { Button } from "@/components/ui/Button";
import { Modal } from "@/components/ui/Modal";
import { AddressDisplay } from "@/components/ui/AddressDisplay";
import { ApprovalStep } from "@/components/tx/ApprovalStep";
import { TxReviewSheet } from "@/components/tx/TxReviewSheet";
import { TxStatusTracker } from "@/components/tx/TxStatusTracker";
import { EligibilityRecheckNotice } from "@/components/liquidations/EligibilityRecheckNotice";
import { LiquidationRewardBreakdown } from "@/components/liquidations/LiquidationRewardBreakdown";
import { RaceConditionNotice } from "@/components/liquidations/RaceConditionNotice";

type LiquidateModalProps = {
  row: LiquidationRow | null;
  onClose: () => void;
};

export function LiquidateModal({ row, onClose }: LiquidateModalProps) {
  const { status: walletStatus } = useWalletState();

  if (row === null) {
    return null;
  }

  const canAfford = liquidatorUsdcBalance >= row.debtAmount;
  const isConnected = walletStatus === WalletStatus.Connected;

  const reviewRows = [
    { label: "Estimated network gas", value: estimatedGasUsd },
    {
      label: "Your USDC balance",
      value: `${formatTokenAmount(liquidatorUsdcBalance, debtDecimals, 2)} ${AssetSymbol.Usdc}`,
    },
  ];

  return (
    <Modal
      open
      onClose={onClose}
      title="Liquidate this position"
      footer={
        <div className="flex flex-col gap-2 sm:flex-row sm:justify-end">
          <Button variant={ButtonVariant.Ghost} size={ButtonSize.Md} onClick={onClose}>
            Cancel
          </Button>
          {isConnected ? (
            <Button size={ButtonSize.Md} disabled={!canAfford}>
              Repay and claim collateral
            </Button>
          ) : null}
        </div>
      }
    >
      <div className="flex flex-col gap-5">
        <div className="flex flex-col gap-1.5">
          <span className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">Borrower</span>
          <AddressDisplay address={row.borrower} />
        </div>

        <LiquidationRewardBreakdown row={row} />

        {row.isUnderwater ? (
          <Alert title="This position cannot fully cover its debt" tone={BadgeTone.Critical} icon={IconName.Warning}>
            The collateral is worth less than the loan plus bonus, so you would receive everything that remains and still
            be short. Phase 1 has no reserve fund covering that gap.
          </Alert>
        ) : null}

        <WalletGate purpose={WalletGatePurpose.Liquidate} skeletonClassName="h-32 rounded-card">
          <div className="flex flex-col gap-5">
            {canAfford ? null : (
              <Alert title="Not enough USDC to repay this loan" tone={BadgeTone.Caution} icon={IconName.Warning}>
                You need {formatTokenAmount(row.debtAmount, debtDecimals, 2)} {AssetSymbol.Usdc} to settle this
                position, and your wallet holds less than that.
              </Alert>
            )}

            <ApprovalStep
              steps={[
                {
                  label: `Approve ${AssetSymbol.Usdc}`,
                  description: "A one-time permission letting the pool collect your repayment.",
                  state: StepState.Active,
                },
                {
                  label: "Liquidate",
                  description: "Repays the loan and transfers the collateral, plus your bonus, to you.",
                  state: StepState.Upcoming,
                },
              ]}
            />

            <TxReviewSheet title="Costs" rows={reviewRows} />

            <EligibilityRecheckNotice />

            <RaceConditionNotice />

            <TxStatusTracker status={txFlowStatus} />
          </div>
        </WalletGate>
      </div>
    </Modal>
  );
}
