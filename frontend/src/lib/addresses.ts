import { isAddress, type Address } from "viem";
import { foundry, sepolia } from "wagmi/chains";

export type ProtocolAddresses = {
  pool: Address;
  vault: Address;
  controller: Address;
  liquidationManager: Address;
  lens: Address;
  oracle: Address;
  rateModel: Address;
  collateralToken: Address;
  debtToken: Address;
};

export type DevAddresses = {
  collateralFeed?: Address;
  debtFeed?: Address;
};

const anvilAddresses: ProtocolAddresses = {
  pool: "0x0165878A594ca255338adfa4d48449f69242Eb8F",
  vault: "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853",
  controller: "0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6",
  liquidationManager: "0x8A791620dd6260079BF849Dc5567aDC3F2FdC318",
  lens: "0x610178dA211FEF7D417bC0e6FeD39F05609AD788",
  oracle: "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9",
  rateModel: "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707",
  collateralToken: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
  debtToken: "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512",
};

const anvilDevAddresses: DevAddresses = {
  collateralFeed: "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
  debtFeed: "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9",
};

function readAddress(raw: string | undefined): Address | undefined {
  const trimmed = raw?.trim();

  if (trimmed === undefined || trimmed === "" || !isAddress(trimmed)) {
    return undefined;
  }

  return trimmed;
}

function addressesFromEnv(): ProtocolAddresses | null {
  const candidate = {
    pool: readAddress(process.env.NEXT_PUBLIC_POOL_ADDRESS),
    vault: readAddress(process.env.NEXT_PUBLIC_VAULT_ADDRESS),
    controller: readAddress(process.env.NEXT_PUBLIC_CONTROLLER_ADDRESS),
    liquidationManager: readAddress(process.env.NEXT_PUBLIC_LIQUIDATION_MANAGER_ADDRESS),
    lens: readAddress(process.env.NEXT_PUBLIC_LENS_ADDRESS),
    oracle: readAddress(process.env.NEXT_PUBLIC_ORACLE_ADDRESS),
    rateModel: readAddress(process.env.NEXT_PUBLIC_RATE_MODEL_ADDRESS),
    collateralToken: readAddress(process.env.NEXT_PUBLIC_COLLATERAL_TOKEN_ADDRESS),
    debtToken: readAddress(process.env.NEXT_PUBLIC_DEBT_TOKEN_ADDRESS),
  };

  const missing = Object.entries(candidate)
    .filter(([, value]) => value === undefined)
    .map(([key]) => key);

  if (missing.length > 0) {
    return null;
  }

  return candidate as ProtocolAddresses;
}

function devAddressesFromEnv(): DevAddresses {
  return {
    collateralFeed: readAddress(process.env.NEXT_PUBLIC_COLLATERAL_FEED_ADDRESS),
    debtFeed: readAddress(process.env.NEXT_PUBLIC_DEBT_FEED_ADDRESS),
  };
}

export function addressesFor(chainId: number | undefined): ProtocolAddresses | null {
  if (chainId === undefined) {
    return null;
  }

  const configured = addressesFromEnv();

  if (configured !== null) {
    return configured;
  }

  if (chainId === foundry.id) {
    return anvilAddresses;
  }

  return null;
}

export function devAddressesFor(chainId: number | undefined): DevAddresses {
  if (chainId === foundry.id) {
    const configured = devAddressesFromEnv();

    return {
      collateralFeed: configured.collateralFeed ?? anvilDevAddresses.collateralFeed,
      debtFeed: configured.debtFeed ?? anvilDevAddresses.debtFeed,
    };
  }

  return devAddressesFromEnv();
}

export function isSupportedChain(chainId: number | undefined): boolean {
  return addressesFor(chainId) !== null;
}

export function supportedChainIds(): number[] {
  return [foundry.id, sepolia.id].filter(isSupportedChain);
}
