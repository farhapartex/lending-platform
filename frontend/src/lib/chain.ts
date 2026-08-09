import { foundry, sepolia } from "wagmi/chains";
import type { Chain } from "viem";

const supportedChains: Chain[] = [sepolia, foundry];

const configuredChainId = Number(process.env.NEXT_PUBLIC_CHAIN_ID ?? sepolia.id);

export const appChain: Chain =
  supportedChains.find((chain) => chain.id === configuredChainId) ?? sepolia;

export const appChainId = appChain.id;

export const appRpcUrl = process.env.NEXT_PUBLIC_RPC_URL ?? appChain.rpcUrls.default.http[0];

export const walletConnectProjectId = process.env.NEXT_PUBLIC_WALLETCONNECT_PROJECT_ID ?? "";

export const appName = "Lending Platform";

function explorerBaseUrl(): string | undefined {
  return appChain.blockExplorers?.default.url;
}

export function explorerAddressUrl(address: string): string | undefined {
  const base = explorerBaseUrl();

  return base === undefined ? undefined : `${base}/address/${address}`;
}

export function explorerTxUrl(txHash: string): string | undefined {
  const base = explorerBaseUrl();

  return base === undefined ? undefined : `${base}/tx/${txHash}`;
}

export function chainDisplayName(chainId: number | undefined): string {
  if (chainId === undefined) {
    return "Unknown network";
  }

  const match = supportedChains.find((chain) => chain.id === chainId);

  return match?.name ?? `Chain ${chainId}`;
}

export function isTestnetChain(chainId: number | undefined): boolean {
  if (chainId === undefined) {
    return false;
  }

  const match = supportedChains.find((chain) => chain.id === chainId);

  return match?.testnet === true;
}
