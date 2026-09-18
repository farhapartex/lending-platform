import { AssetSymbol } from "@/lib/enums";

export const fusdPreviewNotice = {
  title: "Design preview",
  body:
    "FUSD is not deployed. The figures on this page are illustrative, so the screen can be reviewed before the engine exists. Nothing here reads a contract and nothing can be signed.",
} as const;

export type FusdCollateral = {
  symbol: AssetSymbol;
  name: string;
  decimals: number;
  deposited: bigint;
  walletBalance: bigint;
  unitPriceScaled: bigint;
  depositCap: bigint;
};

export const fusdCollaterals: FusdCollateral[] = [
  {
    symbol: AssetSymbol.Weth,
    name: "Wrapped Ether",
    decimals: 18,
    deposited: 4_000000000000000000n,
    walletBalance: 11_000000000000000000n,
    unitPriceScaled: 290000000000n,
    depositCap: 500_000000000000000000n,
  },
  {
    symbol: AssetSymbol.Wbtc,
    name: "Wrapped Bitcoin",
    decimals: 8,
    deposited: 5000000n,
    walletBalance: 24000000n,
    unitPriceScaled: 6500000000000n,
    depositCap: 2000000000n,
  },
];

export const fusdMinted = 4_200_000000000000000000n;

export const fusdWalletBalance = 1_150_000000000000000000n;

export const fusdDecimals = 18;

export const fusdCollateralRatioBps = 15_000n;

export const fusdLiquidationThresholdBps = 6_667n;

export const fusdLiquidatorBonusBps = 800n;

export const fusdProtocolFeeBps = 200n;

export type FusdBackingEntry = {
  symbol: AssetSymbol;
  name: string;
  decimals: number;
  locked: bigint;
  unitPriceScaled: bigint;
  feedUpdatedSecondsAgo: number;
};

export const fusdTotalSupply = 2_480_000_000000000000000000n;

export const fusdProtocolReserve = 12_400_000000000000000000n;

export const fusdBacking: FusdBackingEntry[] = [
  {
    symbol: AssetSymbol.Weth,
    name: "Wrapped Ether",
    decimals: 18,
    locked: 1_150_000000000000000000n,
    unitPriceScaled: 290000000000n,
    feedUpdatedSecondsAgo: 42,
  },
  {
    symbol: AssetSymbol.Wbtc,
    name: "Wrapped Bitcoin",
    decimals: 8,
    locked: 1840000000n,
    unitPriceScaled: 6500000000000n,
    feedUpdatedSecondsAgo: 96,
  },
];

export const fusdOracleTimeoutSeconds = 10_800;

export type FusdUnsafePosition = {
  id: string;
  borrower: string;
  minted: bigint;
  collateral: { symbol: AssetSymbol; decimals: number; amount: bigint; unitPriceScaled: bigint }[];
};

export const fusdUnsafePositions: FusdUnsafePosition[] = [
  {
    id: "pos-a",
    borrower: "0x4C8fA2b71E9d05a3C6b48D07fE2a91B5d7c30E4a",
    minted: 18_400_000000000000000000n,
    collateral: [{ symbol: AssetSymbol.Weth, decimals: 18, amount: 9_400000000000000000n, unitPriceScaled: 290000000000n }],
  },
  {
    id: "pos-b",
    borrower: "0xD19a7F03cB6e45182aB9f5cE71d04A83b2e6C915",
    minted: 64_200_000000000000000000n,
    collateral: [
      { symbol: AssetSymbol.Weth, decimals: 18, amount: 21_000000000000000000n, unitPriceScaled: 290000000000n },
      { symbol: AssetSymbol.Wbtc, decimals: 8, amount: 55000000n, unitPriceScaled: 6500000000000n },
    ],
  },
  {
    id: "pos-c",
    borrower: "0x7Ba0E9cD83f14562aB09d7E3c05F8a26b1D4e07f",
    minted: 5_900_000000000000000000n,
    collateral: [{ symbol: AssetSymbol.Wbtc, decimals: 8, amount: 14000000n, unitPriceScaled: 6500000000000n }],
  },
];
