import { AmountValidationCode, CollateralTab, DebtTab, LendTab } from "@/lib/enums";

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

type CollateralValidationInput = {
  tab: CollateralTab;
  amount: bigint | null;
  walletBalance: bigint;
  collateralDeposited: bigint;
  maxSafeWithdrawal: bigint;
};

export function validateCollateralAmount({
  tab,
  amount,
  walletBalance,
  collateralDeposited,
  maxSafeWithdrawal,
}: CollateralValidationInput): AmountValidationCode {
  if (amount === null) {
    return AmountValidationCode.Empty;
  }

  if (amount <= 0n) {
    return AmountValidationCode.InvalidAmount;
  }

  if (tab === CollateralTab.Deposit) {
    return amount > walletBalance ? AmountValidationCode.ExceedsWalletBalance : AmountValidationCode.None;
  }

  if (amount > collateralDeposited) {
    return AmountValidationCode.ExceedsCollateral;
  }

  if (amount > maxSafeWithdrawal) {
    return AmountValidationCode.ExceedsSafeWithdrawal;
  }

  return AmountValidationCode.None;
}

type DebtValidationInput = {
  tab: DebtTab;
  amount: bigint | null;
  borrowCapacity: bigint;
  availableLiquidity: bigint;
  debtOutstanding: bigint;
  walletBalance: bigint;
};

export function validateDebtAmount({
  tab,
  amount,
  borrowCapacity,
  availableLiquidity,
  debtOutstanding,
  walletBalance,
}: DebtValidationInput): AmountValidationCode {
  if (amount === null) {
    return AmountValidationCode.Empty;
  }

  if (amount <= 0n) {
    return AmountValidationCode.InvalidAmount;
  }

  if (tab === DebtTab.Borrow) {
    if (amount > borrowCapacity) {
      return AmountValidationCode.ExceedsBorrowLimit;
    }

    if (amount > availableLiquidity) {
      return AmountValidationCode.ExceedsAvailableLiquidity;
    }

    return AmountValidationCode.None;
  }

  if (amount > debtOutstanding) {
    return AmountValidationCode.ExceedsDebt;
  }

  if (amount > walletBalance) {
    return AmountValidationCode.ExceedsWalletBalance;
  }

  return AmountValidationCode.None;
}
