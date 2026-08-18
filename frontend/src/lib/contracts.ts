import { erc20Abi, type Abi, type Address } from "viem";

import {
  collateralVaultAbi,
  interestRateModelAbi,
  lendingControllerAbi,
  lendingPoolAbi,
  liquidationManagerAbi,
  mockAggregatorAbi,
  mockErc20Abi,
  positionLensAbi,
  priceOracleAdapterAbi,
} from "@/lib/contracts.generated";
import { addressesFor, devAddressesFor, type ProtocolAddresses } from "@/lib/addresses";

export type ContractRef<TAbi extends Abi> = {
  address: Address;
  abi: TAbi;
};

export type ProtocolContracts = {
  pool: ContractRef<typeof lendingPoolAbi>;
  vault: ContractRef<typeof collateralVaultAbi>;
  controller: ContractRef<typeof lendingControllerAbi>;
  liquidationManager: ContractRef<typeof liquidationManagerAbi>;
  lens: ContractRef<typeof positionLensAbi>;
  oracle: ContractRef<typeof priceOracleAdapterAbi>;
  rateModel: ContractRef<typeof interestRateModelAbi>;
  collateralToken: ContractRef<typeof erc20Abi>;
  debtToken: ContractRef<typeof erc20Abi>;
};

export type DevContracts = {
  collateralToken: ContractRef<typeof mockErc20Abi>;
  debtToken: ContractRef<typeof mockErc20Abi>;
  collateralFeed?: ContractRef<typeof mockAggregatorAbi>;
  debtFeed?: ContractRef<typeof mockAggregatorAbi>;
};

function ref<TAbi extends Abi>(address: Address, abi: TAbi): ContractRef<TAbi> {
  return { address, abi };
}

function build(addresses: ProtocolAddresses): ProtocolContracts {
  return {
    pool: ref(addresses.pool, lendingPoolAbi),
    vault: ref(addresses.vault, collateralVaultAbi),
    controller: ref(addresses.controller, lendingControllerAbi),
    liquidationManager: ref(addresses.liquidationManager, liquidationManagerAbi),
    lens: ref(addresses.lens, positionLensAbi),
    oracle: ref(addresses.oracle, priceOracleAdapterAbi),
    rateModel: ref(addresses.rateModel, interestRateModelAbi),
    collateralToken: ref(addresses.collateralToken, erc20Abi),
    debtToken: ref(addresses.debtToken, erc20Abi),
  };
}

export function contractsFor(chainId: number | undefined): ProtocolContracts | null {
  const addresses = addressesFor(chainId);

  if (addresses === null) {
    return null;
  }

  return build(addresses);
}

export function devContractsFor(chainId: number | undefined): DevContracts | null {
  const addresses = addressesFor(chainId);

  if (addresses === null) {
    return null;
  }

  const feeds = devAddressesFor(chainId);

  return {
    collateralToken: ref(addresses.collateralToken, mockErc20Abi),
    debtToken: ref(addresses.debtToken, mockErc20Abi),
    collateralFeed:
      feeds.collateralFeed === undefined ? undefined : ref(feeds.collateralFeed, mockAggregatorAbi),
    debtFeed: feeds.debtFeed === undefined ? undefined : ref(feeds.debtFeed, mockAggregatorAbi),
  };
}

export function spenderFor(contracts: ProtocolContracts, action: ApprovalAction): Address {
  switch (action) {
    case "deposit":
    case "repay":
    case "liquidate":
      return contracts.pool.address;
    case "addCollateral":
      return contracts.vault.address;
  }
}

export type ApprovalAction = "deposit" | "repay" | "liquidate" | "addCollateral";

export function tokenFor(contracts: ProtocolContracts, action: ApprovalAction): ContractRef<typeof erc20Abi> {
  switch (action) {
    case "deposit":
    case "repay":
    case "liquidate":
      return contracts.debtToken;
    case "addCollateral":
      return contracts.collateralToken;
  }
}
