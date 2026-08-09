import { createConfig, http } from "wagmi";
import { coinbaseWallet, injected, walletConnect } from "wagmi/connectors";
import type { CreateConnectorFn } from "wagmi";
import { appChain, appName, appRpcUrl, walletConnectProjectId } from "@/lib/chain";

function buildConnectors(): CreateConnectorFn[] {
  const connectors: CreateConnectorFn[] = [
    injected({ shimDisconnect: true }),
    coinbaseWallet({ appName }),
  ];

  if (walletConnectProjectId !== "") {
    connectors.push(walletConnect({ projectId: walletConnectProjectId, showQrModal: true }));
  }

  return connectors;
}

export const wagmiConfig = createConfig({
  chains: [appChain],
  connectors: buildConnectors(),
  transports: {
    [appChain.id]: http(appRpcUrl),
  },
  multiInjectedProviderDiscovery: true,
  ssr: true,
});

declare module "wagmi" {
  interface Register {
    config: typeof wagmiConfig;
  }
}
