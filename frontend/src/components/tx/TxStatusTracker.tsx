import { BadgeTone, IconName, TxFlowStatus } from "@/lib/enums";
import { Alert } from "@/components/ui/Alert";

const titles: Record<TxFlowStatus, string> = {
  [TxFlowStatus.Idle]: "",
  [TxFlowStatus.AwaitingApproval]: "Approve USDC in your wallet",
  [TxFlowStatus.AwaitingSignature]: "Confirm the transaction in your wallet",
  [TxFlowStatus.Pending]: "Transaction submitted",
  [TxFlowStatus.Confirmed]: "Transaction confirmed",
  [TxFlowStatus.Reverted]: "Transaction failed",
};

const descriptions: Record<TxFlowStatus, string> = {
  [TxFlowStatus.Idle]: "",
  [TxFlowStatus.AwaitingApproval]: "Your wallet is asking you to allow the pool to move this amount of USDC.",
  [TxFlowStatus.AwaitingSignature]: "Nothing has been sent yet. You can still reject this in your wallet.",
  [TxFlowStatus.Pending]: "Waiting for the network to include it in a block. This usually takes a few seconds.",
  [TxFlowStatus.Confirmed]: "Your balance has been updated and interest is already accruing.",
  [TxFlowStatus.Reverted]: "Nothing was moved and no funds were lost. You only paid the network gas.",
};

const tones: Record<TxFlowStatus, BadgeTone> = {
  [TxFlowStatus.Idle]: BadgeTone.Neutral,
  [TxFlowStatus.AwaitingApproval]: BadgeTone.Brand,
  [TxFlowStatus.AwaitingSignature]: BadgeTone.Brand,
  [TxFlowStatus.Pending]: BadgeTone.Brand,
  [TxFlowStatus.Confirmed]: BadgeTone.Positive,
  [TxFlowStatus.Reverted]: BadgeTone.Critical,
};

const icons: Record<TxFlowStatus, IconName> = {
  [TxFlowStatus.Idle]: IconName.Info,
  [TxFlowStatus.AwaitingApproval]: IconName.Wallet,
  [TxFlowStatus.AwaitingSignature]: IconName.Wallet,
  [TxFlowStatus.Pending]: IconName.Loader,
  [TxFlowStatus.Confirmed]: IconName.Check,
  [TxFlowStatus.Reverted]: IconName.Warning,
};

type TxStatusTrackerProps = {
  status: TxFlowStatus;
};

export function TxStatusTracker({ status }: TxStatusTrackerProps) {
  if (status === TxFlowStatus.Idle) {
    return null;
  }

  return (
    <Alert title={titles[status]} tone={tones[status]} icon={icons[status]}>
      {descriptions[status]}
    </Alert>
  );
}
