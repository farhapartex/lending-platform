"use client";

import { useAccount } from "wagmi";
import { WalletStatus } from "@/lib/enums";
import { appChainId } from "@/lib/chain";

export type WalletState = {
  status: WalletStatus;
  address: string | undefined;
  chainId: number | undefined;
  chainName: string | undefined;
  isSettling: boolean;
};

export function useWalletState(): WalletState {
  const { address, chain, chainId, status } = useAccount();

  if (status === "connected") {
    const onExpectedChain = chainId === appChainId;

    return {
      status: onExpectedChain ? WalletStatus.Connected : WalletStatus.WrongNetwork,
      address,
      chainId,
      chainName: chain?.name,
      isSettling: false,
    };
  }

  if (status === "reconnecting") {
    return {
      status: WalletStatus.Connecting,
      address: undefined,
      chainId: undefined,
      chainName: undefined,
      isSettling: true,
    };
  }

  if (status === "connecting") {
    return {
      status: WalletStatus.Connecting,
      address: undefined,
      chainId: undefined,
      chainName: undefined,
      isSettling: false,
    };
  }

  return {
    status: WalletStatus.Disconnected,
    address: undefined,
    chainId: undefined,
    chainName: undefined,
    isSettling: false,
  };
}
