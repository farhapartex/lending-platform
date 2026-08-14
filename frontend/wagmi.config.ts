import { readFileSync } from "node:fs";
import { resolve } from "node:path";

import { defineConfig } from "@wagmi/cli";
import { react } from "@wagmi/cli/plugins";
import type { Abi } from "viem";

const abiDir = resolve(process.cwd(), "..", "abi");

const contractNames = [
  "LendingPool",
  "CollateralVault",
  "LendingController",
  "LiquidationManager",
  "PositionLens",
  "PriceOracleAdapter",
  "InterestRateModel",
  "MockERC20",
  "MockAggregator",
] as const;

function loadAbi(name: string): Abi {
  return JSON.parse(readFileSync(resolve(abiDir, `${name}.json`), "utf8")) as Abi;
}

export default defineConfig({
  out: "src/lib/contracts.generated.ts",
  contracts: contractNames.map((name) => ({
    name,
    abi: loadAbi(name),
  })),
  plugins: [react()],
});
