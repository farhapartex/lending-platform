const groupingFormatter = new Intl.NumberFormat("en-US");

export function sanitizeAmountInput(raw: string, decimals: number): string {
  const cleaned = raw.replace(/[^0-9.]/g, "");
  const [whole = "", ...rest] = cleaned.split(".");

  if (rest.length === 0) {
    return whole;
  }

  const fraction = rest.join("").slice(0, decimals);
  return `${whole}.${fraction}`;
}

export function parseTokenAmount(input: string, decimals: number): bigint | null {
  if (input === "" || input === ".") {
    return null;
  }

  const [whole = "", fraction = ""] = input.split(".");

  if (!/^\d*$/.test(whole) || !/^\d*$/.test(fraction)) {
    return null;
  }

  const paddedFraction = fraction.padEnd(decimals, "0").slice(0, decimals);
  const base = 10n ** BigInt(decimals);

  return BigInt(whole || "0") * base + BigInt(paddedFraction || "0");
}

export function formatTokenAmount(value: bigint, decimals: number, maxFractionDigits = 4): string {
  const base = 10n ** BigInt(decimals);
  const isNegative = value < 0n;
  const absolute = isNegative ? -value : value;

  const whole = absolute / base;
  const fractionDigits = (absolute % base).toString().padStart(decimals, "0");
  const trimmedFraction = fractionDigits.slice(0, maxFractionDigits).replace(/0+$/, "");

  const wholeText = groupingFormatter.format(whole);
  const sign = isNegative ? "-" : "";

  return trimmedFraction === "" ? `${sign}${wholeText}` : `${sign}${wholeText}.${trimmedFraction}`;
}

export function toAmountInputValue(value: bigint, decimals: number): string {
  const base = 10n ** BigInt(decimals);
  const whole = (value / base).toString();
  const fraction = (value % base).toString().padStart(decimals, "0").replace(/0+$/, "");

  return fraction === "" ? whole : `${whole}.${fraction}`;
}

export function tokenAmountToUsd(value: bigint, decimals: number, unitPrice: number): number {
  const base = 10n ** BigInt(decimals);
  const whole = Number(value / base);
  const fraction = Number(value % base) / Number(base);

  return (whole + fraction) * unitPrice;
}

export function minBigInt(left: bigint, right: bigint): bigint {
  return left < right ? left : right;
}
