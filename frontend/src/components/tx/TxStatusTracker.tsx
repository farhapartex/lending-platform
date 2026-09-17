import { BadgeTone, IconName, TxFlowStatus } from "@/lib/enums";
import { Alert } from "@/components/ui/Alert";

const titles: Record<TxFlowStatus, string> = {
  [TxFlowStatus.Idle]: "",
  [TxFlowStatus.AwaitingApproval]: "Approve the token in your wallet",
  [TxFlowStatus.AwaitingSignature]: "Confirm the transaction in your wallet",
  [TxFlowStatus.Pending]: "Transaction submitted",
  [TxFlowStatus.Confirmed]: "Transaction confirmed",
  [TxFlowStatus.Reverted]: "Transaction failed",
};

const descriptions: Record<TxFlowStatus, string> = {
  [TxFlowStatus.Idle]: "",
  [TxFlowStatus.AwaitingApproval]: "Your wallet is asking you to allow this amount to be moved.",
  [TxFlowStatus.AwaitingSignature]: "Nothing has been sent yet. You can still reject this in your wallet.",
  [TxFlowStatus.Pending]: "Waiting for the network to include it in a block. This usually takes a few seconds.",
  [TxFlowStatus.Confirmed]: "Your position has been updated.",
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
  approvalAsset?: string;
  approvalSpender?: string;
  confirmedMessage?: string;
  errorMessage?: string | null;
};

export function TxStatusTracker({
  status,
  approvalAsset,
  approvalSpender,
  confirmedMessage,
  errorMessage,
}: TxStatusTrackerProps) {
  if (status === TxFlowStatus.Idle) {
    return null;
  }

  const title =
    status === TxFlowStatus.AwaitingApproval && approvalAsset !== undefined
      ? `Approve ${approvalAsset} in your wallet`
      : titles[status];

  let description = descriptions[status];

  if (status === TxFlowStatus.AwaitingApproval && approvalAsset !== undefined && approvalSpender !== undefined) {
    description = `Your wallet is asking you to allow the ${approvalSpender} to move this amount of ${approvalAsset}.`;
  }

  if (status === TxFlowStatus.Confirmed && confirmedMessage !== undefined) {
    description = confirmedMessage;
  }

  if (status === TxFlowStatus.Reverted && errorMessage !== undefined && errorMessage !== null) {
    description = `${errorMessage} Nothing was moved, and you only paid the network gas.`;
  }

  return (
    <Alert title={title} tone={tones[status]} icon={icons[status]}>
      {description}
    </Alert>
  );
}
