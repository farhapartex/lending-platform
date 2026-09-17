"use client";

import { useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { useConfig, useWriteContract } from "wagmi";
import { waitForTransactionReceipt } from "wagmi/actions";
import type { Abi, Address, Hash } from "viem";

import { TxFlowStatus } from "@/lib/enums";

export type TxStep = {
  address: Address;
  abi: Abi;
  functionName: string;
  args: readonly unknown[];
};

export type TxFlowParams = {
  approval: TxStep | null;
  action: TxStep | null;
  onConfirmed?: () => void;
};

export type TxFlowResult = {
  status: TxFlowStatus;
  error: string | null;
  hash: Hash | undefined;
  isBusy: boolean;
  submit: () => void;
  reset: () => void;
};

const rejectionPatterns = ["User rejected", "User denied", "rejected the request"];

function isRejection(message: string): boolean {
  return rejectionPatterns.some((pattern) => message.includes(pattern));
}

function readableError(error: unknown): string {
  if (typeof error === "object" && error !== null) {
    const candidate = error as { shortMessage?: string; message?: string };

    if (typeof candidate.shortMessage === "string") {
      return candidate.shortMessage;
    }

    if (typeof candidate.message === "string") {
      return candidate.message.split("\n")[0];
    }
  }

  return "The transaction could not be completed.";
}

export function useTxFlow({ approval, action, onConfirmed }: TxFlowParams): TxFlowResult {
  const config = useConfig();
  const queryClient = useQueryClient();
  const { writeContractAsync } = useWriteContract();

  const [status, setStatus] = useState(TxFlowStatus.Idle);
  const [error, setError] = useState<string | null>(null);
  const [hash, setHash] = useState<Hash | undefined>(undefined);

  const isBusy =
    status === TxFlowStatus.AwaitingApproval ||
    status === TxFlowStatus.AwaitingSignature ||
    status === TxFlowStatus.Pending;

  const reset = () => {
    setStatus(TxFlowStatus.Idle);
    setError(null);
    setHash(undefined);
  };

  const send = async (step: TxStep): Promise<Hash> =>
    writeContractAsync({
      address: step.address,
      abi: step.abi,
      functionName: step.functionName,
      args: step.args,
    } as never);

  const submit = () => {
    if (action === null || isBusy) {
      return;
    }

    setError(null);
    setHash(undefined);

    void (async () => {
      try {
        if (approval !== null) {
          setStatus(TxFlowStatus.AwaitingApproval);
          const approvalHash = await send(approval);
          setStatus(TxFlowStatus.Pending);
          await waitForTransactionReceipt(config, { hash: approvalHash });
        }

        setStatus(TxFlowStatus.AwaitingSignature);
        const actionHash = await send(action);
        setHash(actionHash);

        setStatus(TxFlowStatus.Pending);
        const receipt = await waitForTransactionReceipt(config, { hash: actionHash });

        if (receipt.status === "reverted") {
          setStatus(TxFlowStatus.Reverted);
          setError("The network rejected this transaction.");

          return;
        }

        setStatus(TxFlowStatus.Confirmed);
        await queryClient.invalidateQueries();
        onConfirmed?.();
      } catch (caught) {
        const message = readableError(caught);

        if (isRejection(message)) {
          reset();

          return;
        }

        setStatus(TxFlowStatus.Reverted);
        setError(message);
      }
    })();
  };

  return { status, error, hash, isBusy, submit, reset };
}
