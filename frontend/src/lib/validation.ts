import { AmountValidationCode, LendTab } from "@/lib/enums";

type LendValidationInput = {
  tab: LendTab;
  amount: bigint | null;
  walletBalance: bigint;
  depositedBalance: bigint;
  availableLiquidity: bigint;
  minimumDeposit: bigint;
};

export function validateLendAmount({
  tab,
  amount,
  walletBalance,
  depositedBalance,
  availableLiquidity,
  minimumDeposit,
}: LendValidationInput): AmountValidationCode {
  if (amount === null) {
    return AmountValidationCode.Empty;
  }

  if (amount <= 0n) {
    return AmountValidationCode.InvalidAmount;
  }

  if (tab === LendTab.Deposit) {
    if (amount < minimumDeposit) {
      return AmountValidationCode.BelowMinimum;
    }

    if (amount > walletBalance) {
      return AmountValidationCode.ExceedsWalletBalance;
    }

    return AmountValidationCode.None;
  }

  if (amount > depositedBalance) {
    return AmountValidationCode.ExceedsDeposit;
  }

  if (amount > availableLiquidity) {
    return AmountValidationCode.ExceedsAvailableLiquidity;
  }

  return AmountValidationCode.None;
}

export function isBlockingValidation(code: AmountValidationCode): boolean {
  return code !== AmountValidationCode.None;
}
