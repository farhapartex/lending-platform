export const maskedIdPrefixes = {
  transaction: "txn",
  market: "mkt",
  position: "pos",
  liquidation: "liq",
} as const;

export type MaskedIdKind = keyof typeof maskedIdPrefixes;

const bodyPattern = /^[a-z2-7]{16,32}$/;

export function isMaskedId(value: string | null | undefined, kind: MaskedIdKind): boolean {
  if (value === null || value === undefined) {
    return false;
  }

  const parts = value.split("_");

  if (parts.length !== 2) {
    return false;
  }

  const [prefix, body] = parts;

  return prefix === maskedIdPrefixes[kind] && bodyPattern.test(body);
}
