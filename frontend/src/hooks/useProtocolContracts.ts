"use client";

import { useMemo } from "react";
import { useChainId } from "wagmi";

import { appChainId } from "@/lib/chain";
import { contractsFor, devContractsFor, type DevContracts, type ProtocolContracts } from "@/lib/contracts";

export type ProtocolContractsState = {
  chainId: number;
  contracts: ProtocolContracts | null;
  devContracts: DevContracts | null;
  isSupported: boolean;
  isLocalChain: boolean;
};

const localChainId = 31337;

export function useProtocolContracts(): ProtocolContractsState {
  const connectedChainId = useChainId();
  const chainId = connectedChainId ?? appChainId;

  return useMemo(() => {
    const contracts = contractsFor(chainId);

    return {
      chainId,
      contracts,
      devContracts: chainId === localChainId ? devContractsFor(chainId) : null,
      isSupported: contracts !== null,
      isLocalChain: chainId === localChainId,
    };
  }, [chainId]);
}
