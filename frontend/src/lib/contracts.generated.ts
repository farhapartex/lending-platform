import {
  createUseReadContract,
  createUseWriteContract,
  createUseSimulateContract,
  createUseWatchContractEvent,
} from 'wagmi/codegen'

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// CollateralVault
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const collateralVaultAbi = [
  {
    type: 'constructor',
    inputs: [
      { name: 'owner', internalType: 'address', type: 'address' },
      { name: 'vaultAsset', internalType: 'contract IERC20', type: 'address' },
    ],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'collateralAsset',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: '', internalType: 'address', type: 'address' }],
    name: 'collateralOf',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'collateralToken',
    outputs: [{ name: '', internalType: 'contract IERC20', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'controller',
    outputs: [
      {
        name: '',
        internalType: 'contract ILendingController',
        type: 'address',
      },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'amount', internalType: 'uint256', type: 'uint256' }],
    name: 'depositCollateral',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      {
        name: 'newController',
        internalType: 'contract ILendingController',
        type: 'address',
      },
    ],
    name: 'linkController',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      {
        name: 'newLiquidationManager',
        internalType: 'address',
        type: 'address',
      },
    ],
    name: 'linkLiquidationManager',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'liquidationManager',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'borrower', internalType: 'address', type: 'address' },
      { name: 'recipient', internalType: 'address', type: 'address' },
      { name: 'amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'seize',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'totalCollateral',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'amount', internalType: 'uint256', type: 'uint256' }],
    name: 'withdrawCollateral',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'borrower',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'newCollateral',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'CollateralDeposited',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'borrower',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'newCollateral',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'CollateralSeized',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'borrower',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'newCollateral',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'CollateralWithdrawn',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'controller',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'ControllerLinked',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'liquidationManager',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'LiquidationManagerLinked',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  { type: 'error', inputs: [], name: 'AlreadyInitialized' },
  {
    type: 'error',
    inputs: [
      { name: 'requested', internalType: 'uint256', type: 'uint256' },
      { name: 'held', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'ExceedsCollateralBalance',
  },
  {
    type: 'error',
    inputs: [{ name: 'caller', internalType: 'address', type: 'address' }],
    name: 'NotAuthorized',
  },
  {
    type: 'error',
    inputs: [{ name: 'owner', internalType: 'address', type: 'address' }],
    name: 'OwnableInvalidOwner',
  },
  {
    type: 'error',
    inputs: [{ name: 'account', internalType: 'address', type: 'address' }],
    name: 'OwnableUnauthorizedAccount',
  },
  { type: 'error', inputs: [], name: 'ReentrancyGuardReentrantCall' },
  {
    type: 'error',
    inputs: [{ name: 'token', internalType: 'address', type: 'address' }],
    name: 'SafeERC20FailedOperation',
  },
  {
    type: 'error',
    inputs: [
      { name: 'requested', internalType: 'uint256', type: 'uint256' },
      { name: 'safeAmount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'WouldBreakBorrowLimit',
  },
  { type: 'error', inputs: [], name: 'ZeroAddress' },
  { type: 'error', inputs: [], name: 'ZeroAmount' },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// InterestRateModel
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const interestRateModelAbi = [
  {
    type: 'constructor',
    inputs: [
      { name: 'owner', internalType: 'address', type: 'address' },
      {
        name: 'startingCurve',
        internalType: 'struct RateCurve',
        type: 'tuple',
        components: [
          { name: 'baseRatePerSecond', internalType: 'uint64', type: 'uint64' },
          {
            name: 'slopeBelowKinkPerSecond',
            internalType: 'uint64',
            type: 'uint64',
          },
          {
            name: 'slopeAboveKinkPerSecond',
            internalType: 'uint64',
            type: 'uint64',
          },
          {
            name: 'kinkUtilizationBps',
            internalType: 'uint16',
            type: 'uint16',
          },
        ],
      },
    ],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'usageBps', internalType: 'uint256', type: 'uint256' }],
    name: 'borrowAprBps',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'usageBps', internalType: 'uint256', type: 'uint256' }],
    name: 'borrowRatePerSecond',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'curve',
    outputs: [
      {
        name: '',
        internalType: 'struct RateCurve',
        type: 'tuple',
        components: [
          { name: 'baseRatePerSecond', internalType: 'uint64', type: 'uint64' },
          {
            name: 'slopeBelowKinkPerSecond',
            internalType: 'uint64',
            type: 'uint64',
          },
          {
            name: 'slopeAboveKinkPerSecond',
            internalType: 'uint64',
            type: 'uint64',
          },
          {
            name: 'kinkUtilizationBps',
            internalType: 'uint16',
            type: 'uint16',
          },
        ],
      },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      {
        name: 'newCurve',
        internalType: 'struct RateCurve',
        type: 'tuple',
        components: [
          { name: 'baseRatePerSecond', internalType: 'uint64', type: 'uint64' },
          {
            name: 'slopeBelowKinkPerSecond',
            internalType: 'uint64',
            type: 'uint64',
          },
          {
            name: 'slopeAboveKinkPerSecond',
            internalType: 'uint64',
            type: 'uint64',
          },
          {
            name: 'kinkUtilizationBps',
            internalType: 'uint16',
            type: 'uint16',
          },
        ],
      },
    ],
    name: 'setCurve',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'usageBps', internalType: 'uint256', type: 'uint256' },
      { name: 'reserveFactorBps', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'supplyAprBps',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'usageBps', internalType: 'uint256', type: 'uint256' },
      { name: 'reserveFactorBps', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'supplyRatePerSecond',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'totalSupplied', internalType: 'uint256', type: 'uint256' },
      { name: 'totalBorrowed', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'utilizationBps',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'pure',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'baseRatePerSecond',
        internalType: 'uint64',
        type: 'uint64',
        indexed: false,
      },
      {
        name: 'slopeBelowKinkPerSecond',
        internalType: 'uint64',
        type: 'uint64',
        indexed: false,
      },
      {
        name: 'slopeAboveKinkPerSecond',
        internalType: 'uint64',
        type: 'uint64',
        indexed: false,
      },
      {
        name: 'kinkUtilizationBps',
        internalType: 'uint16',
        type: 'uint16',
        indexed: false,
      },
    ],
    name: 'CurveChanged',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  { type: 'error', inputs: [], name: 'InvalidRiskSettings' },
  {
    type: 'error',
    inputs: [{ name: 'owner', internalType: 'address', type: 'address' }],
    name: 'OwnableInvalidOwner',
  },
  {
    type: 'error',
    inputs: [{ name: 'account', internalType: 'address', type: 'address' }],
    name: 'OwnableUnauthorizedAccount',
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// LendingController
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const lendingControllerAbi = [
  {
    type: 'constructor',
    inputs: [
      { name: 'owner', internalType: 'address', type: 'address' },
      {
        name: 'lendingPool',
        internalType: 'contract ILendingPool',
        type: 'address',
      },
      {
        name: 'collateralVault',
        internalType: 'contract ICollateralVault',
        type: 'address',
      },
      {
        name: 'priceOracle',
        internalType: 'contract IPriceOracle',
        type: 'address',
      },
      { name: 'startingMaxLtvBps', internalType: 'uint16', type: 'uint16' },
      {
        name: 'startingLiquidationThresholdBps',
        internalType: 'uint16',
        type: 'uint16',
      },
      {
        name: 'startingLiquidationBonusBps',
        internalType: 'uint16',
        type: 'uint16',
      },
    ],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'amount', internalType: 'uint256', type: 'uint256' }],
    name: 'borrow',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'borrowPaused',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'collateralAsset',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'collateralDecimals',
    outputs: [{ name: '', internalType: 'uint8', type: 'uint8' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'borrower', internalType: 'address', type: 'address' }],
    name: 'collateralValueOf',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'debtAsset',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'debtDecimals',
    outputs: [{ name: '', internalType: 'uint8', type: 'uint8' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'borrower', internalType: 'address', type: 'address' }],
    name: 'debtOf',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'borrower', internalType: 'address', type: 'address' }],
    name: 'debtValueOf',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'borrower', internalType: 'address', type: 'address' }],
    name: 'healthFactorBps',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'borrower', internalType: 'address', type: 'address' }],
    name: 'isLiquidatable',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'liquidationBonusBps',
    outputs: [{ name: '', internalType: 'uint16', type: 'uint16' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'liquidationThresholdBps',
    outputs: [{ name: '', internalType: 'uint16', type: 'uint16' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'borrower', internalType: 'address', type: 'address' }],
    name: 'maxBorrowable',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'maxLtvBps',
    outputs: [{ name: '', internalType: 'uint16', type: 'uint16' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'borrower', internalType: 'address', type: 'address' }],
    name: 'maxWithdrawableCollateral',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'oracle',
    outputs: [
      { name: '', internalType: 'contract IPriceOracle', type: 'address' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'pool',
    outputs: [
      { name: '', internalType: 'contract ILendingPool', type: 'address' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'amount', internalType: 'uint256', type: 'uint256' }],
    name: 'repay',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'repayAll',
    outputs: [{ name: 'amountPaid', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'isPaused', internalType: 'bool', type: 'bool' }],
    name: 'setBorrowPaused',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'newMaxLtvBps', internalType: 'uint16', type: 'uint16' },
      { name: 'newThresholdBps', internalType: 'uint16', type: 'uint16' },
      { name: 'newBonusBps', internalType: 'uint16', type: 'uint16' },
    ],
    name: 'setRiskSettings',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'vault',
    outputs: [
      { name: '', internalType: 'contract ICollateralVault', type: 'address' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'borrower',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'newDebt',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'healthFactorBps',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'borrowIndex',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'Borrow',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'isPaused', internalType: 'bool', type: 'bool', indexed: false },
    ],
    name: 'BorrowPausedChanged',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'borrower',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'payer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'newDebt',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'Repay',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'maxLtvBps',
        internalType: 'uint16',
        type: 'uint16',
        indexed: false,
      },
      {
        name: 'liquidationThresholdBps',
        internalType: 'uint16',
        type: 'uint16',
        indexed: false,
      },
      {
        name: 'liquidationBonusBps',
        internalType: 'uint16',
        type: 'uint16',
        indexed: false,
      },
    ],
    name: 'RiskSettingsChanged',
  },
  {
    type: 'error',
    inputs: [
      { name: 'requested', internalType: 'uint256', type: 'uint256' },
      { name: 'allowed', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'ExceedsBorrowLimit',
  },
  {
    type: 'error',
    inputs: [
      { name: 'requested', internalType: 'uint256', type: 'uint256' },
      { name: 'owed', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'ExceedsDebt',
  },
  { type: 'error', inputs: [], name: 'InvalidRiskSettings' },
  { type: 'error', inputs: [], name: 'MarketPaused' },
  {
    type: 'error',
    inputs: [
      { name: 'requested', internalType: 'uint256', type: 'uint256' },
      { name: 'available', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'NotEnoughLiquidity',
  },
  {
    type: 'error',
    inputs: [{ name: 'owner', internalType: 'address', type: 'address' }],
    name: 'OwnableInvalidOwner',
  },
  {
    type: 'error',
    inputs: [{ name: 'account', internalType: 'address', type: 'address' }],
    name: 'OwnableUnauthorizedAccount',
  },
  { type: 'error', inputs: [], name: 'ReentrancyGuardReentrantCall' },
  { type: 'error', inputs: [], name: 'ZeroAddress' },
  { type: 'error', inputs: [], name: 'ZeroAmount' },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// LendingPool
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const lendingPoolAbi = [
  {
    type: 'constructor',
    inputs: [
      { name: 'owner', internalType: 'address', type: 'address' },
      { name: 'poolAsset', internalType: 'contract IERC20', type: 'address' },
      {
        name: 'startingRateModel',
        internalType: 'contract IInterestRateModel',
        type: 'address',
      },
      { name: 'startingMinDeposit', internalType: 'uint256', type: 'uint256' },
      {
        name: 'startingReserveFactorBps',
        internalType: 'uint16',
        type: 'uint16',
      },
    ],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'accrueInterest',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'accruedReserves',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'asset',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'assetToken',
    outputs: [{ name: '', internalType: 'contract IERC20', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'availableLiquidity',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'lender', internalType: 'address', type: 'address' }],
    name: 'balanceOfAssets',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'borrower', internalType: 'address', type: 'address' },
      { name: 'recipient', internalType: 'address', type: 'address' },
      { name: 'amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'borrowFor',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'borrowIndex',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'recipient', internalType: 'address', type: 'address' }],
    name: 'collectAllReserves',
    outputs: [{ name: 'collected', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'recipient', internalType: 'address', type: 'address' },
      { name: 'amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'collectReserves',
    outputs: [{ name: 'collected', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'controller',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'borrower', internalType: 'address', type: 'address' }],
    name: 'debtOf',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: '', internalType: 'address', type: 'address' }],
    name: 'debtSharesOf',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'assets', internalType: 'uint256', type: 'uint256' }],
    name: 'deposit',
    outputs: [{ name: 'shares', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'depositsPaused',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'lastAccrualTimestamp',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'newController', internalType: 'address', type: 'address' },
    ],
    name: 'linkController',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      {
        name: 'newLiquidationManager',
        internalType: 'address',
        type: 'address',
      },
    ],
    name: 'linkLiquidationManager',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'liquidationManager',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'lender', internalType: 'address', type: 'address' }],
    name: 'maxWithdrawable',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'minDeposit',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'previewAccrual',
    outputs: [
      { name: 'nextSupplyIndex', internalType: 'uint256', type: 'uint256' },
      { name: 'nextBorrowIndex', internalType: 'uint256', type: 'uint256' },
      { name: 'reservesToAdd', internalType: 'uint256', type: 'uint256' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'rateModel',
    outputs: [
      {
        name: '',
        internalType: 'contract IInterestRateModel',
        type: 'address',
      },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'shares', internalType: 'uint256', type: 'uint256' }],
    name: 'redeemShares',
    outputs: [{ name: 'assets', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'borrower', internalType: 'address', type: 'address' },
      { name: 'payer', internalType: 'address', type: 'address' },
    ],
    name: 'repayAllFor',
    outputs: [{ name: 'amountPaid', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'borrower', internalType: 'address', type: 'address' },
      { name: 'payer', internalType: 'address', type: 'address' },
      { name: 'amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'repayFor',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'reserveFactorBps',
    outputs: [{ name: '', internalType: 'uint16', type: 'uint16' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'isPaused', internalType: 'bool', type: 'bool' }],
    name: 'setDepositsPaused',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'newAmount', internalType: 'uint256', type: 'uint256' }],
    name: 'setMinDeposit',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      {
        name: 'newModel',
        internalType: 'contract IInterestRateModel',
        type: 'address',
      },
    ],
    name: 'setRateModel',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'newBps', internalType: 'uint16', type: 'uint16' }],
    name: 'setReserveFactorBps',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: '', internalType: 'address', type: 'address' }],
    name: 'sharesOf',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'supplyIndex',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'totalBorrowed',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'totalDebtShares',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'totalSupplied',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'totalSupplyShares',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'utilizationBps',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'assets', internalType: 'uint256', type: 'uint256' }],
    name: 'withdraw',
    outputs: [{ name: 'shares', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'controller',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'ControllerLinked',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'lender',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'assets',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'shares',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'supplyIndex',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'totalSupplied',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'Deposit',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'isPaused', internalType: 'bool', type: 'bool', indexed: false },
    ],
    name: 'DepositsPausedChanged',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'supplyIndex',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'borrowIndex',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'totalBorrowed',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'reservesAccrued',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'InterestAccrued',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'liquidationManager',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'LiquidationManagerLinked',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousAmount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'newAmount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'MinDepositChanged',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousModel',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
      {
        name: 'newModel',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
    ],
    name: 'RateModelChanged',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousBps',
        internalType: 'uint16',
        type: 'uint16',
        indexed: false,
      },
      {
        name: 'newBps',
        internalType: 'uint16',
        type: 'uint16',
        indexed: false,
      },
    ],
    name: 'ReserveFactorChanged',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'remainingReserves',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'ReservesCollected',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'lender',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'assets',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'shares',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'supplyIndex',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'totalSupplied',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'Withdraw',
  },
  { type: 'error', inputs: [], name: 'AlreadyInitialized' },
  {
    type: 'error',
    inputs: [
      { name: 'provided', internalType: 'uint256', type: 'uint256' },
      { name: 'minimum', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'BelowMinimumDeposit',
  },
  {
    type: 'error',
    inputs: [
      { name: 'requested', internalType: 'uint256', type: 'uint256' },
      { name: 'available', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'ExceedsReserves',
  },
  {
    type: 'error',
    inputs: [
      { name: 'requested', internalType: 'uint256', type: 'uint256' },
      { name: 'available', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'ExceedsSupplyBalance',
  },
  { type: 'error', inputs: [], name: 'InvalidRiskSettings' },
  { type: 'error', inputs: [], name: 'MarketPaused' },
  {
    type: 'error',
    inputs: [{ name: 'caller', internalType: 'address', type: 'address' }],
    name: 'NotAuthorized',
  },
  {
    type: 'error',
    inputs: [
      { name: 'requested', internalType: 'uint256', type: 'uint256' },
      { name: 'available', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'NotEnoughLiquidity',
  },
  {
    type: 'error',
    inputs: [{ name: 'owner', internalType: 'address', type: 'address' }],
    name: 'OwnableInvalidOwner',
  },
  {
    type: 'error',
    inputs: [{ name: 'account', internalType: 'address', type: 'address' }],
    name: 'OwnableUnauthorizedAccount',
  },
  { type: 'error', inputs: [], name: 'ReentrancyGuardReentrantCall' },
  {
    type: 'error',
    inputs: [
      { name: 'bits', internalType: 'uint8', type: 'uint8' },
      { name: 'value', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'SafeCastOverflowedUintDowncast',
  },
  {
    type: 'error',
    inputs: [{ name: 'token', internalType: 'address', type: 'address' }],
    name: 'SafeERC20FailedOperation',
  },
  { type: 'error', inputs: [], name: 'ZeroAddress' },
  { type: 'error', inputs: [], name: 'ZeroAmount' },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// LiquidationManager
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const liquidationManagerAbi = [
  {
    type: 'constructor',
    inputs: [
      {
        name: 'lendingPool',
        internalType: 'contract ILendingPool',
        type: 'address',
      },
      {
        name: 'collateralVault',
        internalType: 'contract ICollateralVault',
        type: 'address',
      },
      {
        name: 'priceOracle',
        internalType: 'contract IPriceOracle',
        type: 'address',
      },
      {
        name: 'lendingController',
        internalType: 'contract ILendingController',
        type: 'address',
      },
    ],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'collateralAsset',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'collateralDecimals',
    outputs: [{ name: '', internalType: 'uint8', type: 'uint8' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'controller',
    outputs: [
      {
        name: '',
        internalType: 'contract ILendingController',
        type: 'address',
      },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'debtAsset',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'debtDecimals',
    outputs: [{ name: '', internalType: 'uint8', type: 'uint8' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'borrower', internalType: 'address', type: 'address' }],
    name: 'isLiquidatable',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'borrower', internalType: 'address', type: 'address' }],
    name: 'liquidate',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'oracle',
    outputs: [
      { name: '', internalType: 'contract IPriceOracle', type: 'address' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'pool',
    outputs: [
      { name: '', internalType: 'contract ILendingPool', type: 'address' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'borrower', internalType: 'address', type: 'address' }],
    name: 'previewLiquidation',
    outputs: [
      { name: 'debtToRepay', internalType: 'uint256', type: 'uint256' },
      { name: 'collateralToSeize', internalType: 'uint256', type: 'uint256' },
      { name: 'bonusValue', internalType: 'uint256', type: 'uint256' },
      { name: 'shortfall', internalType: 'uint256', type: 'uint256' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'vault',
    outputs: [
      { name: '', internalType: 'contract ICollateralVault', type: 'address' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'borrower',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'liquidator',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'debtRepaid',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'collateralSeized',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'bonusValue',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'healthFactorBeforeBps',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'collateralPrice',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
      {
        name: 'priceDecimals',
        internalType: 'uint8',
        type: 'uint8',
        indexed: false,
      },
      {
        name: 'shortfall',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'LiquidationExecuted',
  },
  {
    type: 'error',
    inputs: [
      { name: 'healthFactorBps', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'PositionIsHealthy',
  },
  { type: 'error', inputs: [], name: 'ReentrancyGuardReentrantCall' },
  { type: 'error', inputs: [], name: 'ZeroAddress' },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// MockAggregator
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const mockAggregatorAbi = [
  {
    type: 'constructor',
    inputs: [
      { name: 'description_', internalType: 'string', type: 'string' },
      { name: 'decimals_', internalType: 'uint8', type: 'uint8' },
      { name: 'startingAnswer', internalType: 'int256', type: 'int256' },
    ],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'decimals',
    outputs: [{ name: '', internalType: 'uint8', type: 'uint8' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'description',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'feedDown',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'wantedRoundId', internalType: 'uint80', type: 'uint80' }],
    name: 'getRoundData',
    outputs: [
      { name: 'roundId', internalType: 'uint80', type: 'uint80' },
      { name: 'answer', internalType: 'int256', type: 'int256' },
      { name: 'startedAt', internalType: 'uint256', type: 'uint256' },
      { name: 'updatedAt', internalType: 'uint256', type: 'uint256' },
      { name: 'answeredInRound', internalType: 'uint80', type: 'uint80' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'latestRoundData',
    outputs: [
      { name: 'roundId', internalType: 'uint80', type: 'uint80' },
      { name: 'answer', internalType: 'int256', type: 'int256' },
      { name: 'startedAt', internalType: 'uint256', type: 'uint256' },
      { name: 'updatedAt', internalType: 'uint256', type: 'uint256' },
      { name: 'answeredInRound', internalType: 'uint80', type: 'uint80' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'latestRoundId',
    outputs: [{ name: '', internalType: 'uint80', type: 'uint80' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'secondsInThePast', internalType: 'uint40', type: 'uint40' },
    ],
    name: 'makeStale',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'isDown', internalType: 'bool', type: 'bool' }],
    name: 'setFeedDown',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'newAnswer', internalType: 'int256', type: 'int256' }],
    name: 'setIncompleteRound',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'newAnswer', internalType: 'int256', type: 'int256' }],
    name: 'setPrice',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'newAnswer', internalType: 'int256', type: 'int256' },
      { name: 'updatedAt', internalType: 'uint40', type: 'uint40' },
    ],
    name: 'setPriceWithTimestamp',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'version',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'pure',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'roundId',
        internalType: 'uint80',
        type: 'uint80',
        indexed: true,
      },
      {
        name: 'answer',
        internalType: 'int256',
        type: 'int256',
        indexed: false,
      },
      {
        name: 'updatedAt',
        internalType: 'uint40',
        type: 'uint40',
        indexed: false,
      },
    ],
    name: 'AnswerRecorded',
  },
  { type: 'error', inputs: [], name: 'FeedIsDown' },
  {
    type: 'error',
    inputs: [{ name: 'roundId', internalType: 'uint80', type: 'uint80' }],
    name: 'RoundNotFound',
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// MockERC20
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const mockErc20Abi = [
  {
    type: 'constructor',
    inputs: [
      { name: 'tokenName', internalType: 'string', type: 'string' },
      { name: 'tokenSymbol', internalType: 'string', type: 'string' },
      { name: 'tokenDecimalPlaces', internalType: 'uint8', type: 'uint8' },
    ],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'owner', internalType: 'address', type: 'address' },
      { name: 'spender', internalType: 'address', type: 'address' },
    ],
    name: 'allowance',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'spender', internalType: 'address', type: 'address' },
      { name: 'value', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'approve',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'account', internalType: 'address', type: 'address' }],
    name: 'balanceOf',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'holder', internalType: 'address', type: 'address' },
      { name: 'amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'burn',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'decimals',
    outputs: [{ name: '', internalType: 'uint8', type: 'uint8' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'receiver', internalType: 'address', type: 'address' },
      { name: 'amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'mint',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'name',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'symbol',
    outputs: [{ name: '', internalType: 'string', type: 'string' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'totalSupply',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'to', internalType: 'address', type: 'address' },
      { name: 'value', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'transfer',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'from', internalType: 'address', type: 'address' },
      { name: 'to', internalType: 'address', type: 'address' },
      { name: 'value', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'transferFrom',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'owner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'spender',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'value',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'Approval',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'value',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'Transfer',
  },
  {
    type: 'error',
    inputs: [
      { name: 'spender', internalType: 'address', type: 'address' },
      { name: 'allowance', internalType: 'uint256', type: 'uint256' },
      { name: 'needed', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'ERC20InsufficientAllowance',
  },
  {
    type: 'error',
    inputs: [
      { name: 'sender', internalType: 'address', type: 'address' },
      { name: 'balance', internalType: 'uint256', type: 'uint256' },
      { name: 'needed', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'ERC20InsufficientBalance',
  },
  {
    type: 'error',
    inputs: [{ name: 'approver', internalType: 'address', type: 'address' }],
    name: 'ERC20InvalidApprover',
  },
  {
    type: 'error',
    inputs: [{ name: 'receiver', internalType: 'address', type: 'address' }],
    name: 'ERC20InvalidReceiver',
  },
  {
    type: 'error',
    inputs: [{ name: 'sender', internalType: 'address', type: 'address' }],
    name: 'ERC20InvalidSender',
  },
  {
    type: 'error',
    inputs: [{ name: 'spender', internalType: 'address', type: 'address' }],
    name: 'ERC20InvalidSpender',
  },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PositionLens
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const positionLensAbi = [
  {
    type: 'constructor',
    inputs: [
      {
        name: 'lendingPool',
        internalType: 'contract ILendingPool',
        type: 'address',
      },
      {
        name: 'collateralVault',
        internalType: 'contract ICollateralVault',
        type: 'address',
      },
      {
        name: 'priceOracle',
        internalType: 'contract IPriceOracle',
        type: 'address',
      },
      {
        name: 'lendingController',
        internalType: 'contract ILendingController',
        type: 'address',
      },
    ],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'account', internalType: 'address', type: 'address' }],
    name: 'accountData',
    outputs: [
      {
        name: 'data',
        internalType: 'struct AccountData',
        type: 'tuple',
        components: [
          { name: 'supplyShares', internalType: 'uint256', type: 'uint256' },
          { name: 'supplyAssets', internalType: 'uint256', type: 'uint256' },
          {
            name: 'collateralAmount',
            internalType: 'uint256',
            type: 'uint256',
          },
          { name: 'collateralValue', internalType: 'uint256', type: 'uint256' },
          { name: 'debtAmount', internalType: 'uint256', type: 'uint256' },
          { name: 'debtValue', internalType: 'uint256', type: 'uint256' },
          { name: 'healthFactorBps', internalType: 'uint256', type: 'uint256' },
          { name: 'maxBorrowable', internalType: 'uint256', type: 'uint256' },
          {
            name: 'maxWithdrawableCollateral',
            internalType: 'uint256',
            type: 'uint256',
          },
          { name: 'collateralPrice', internalType: 'uint256', type: 'uint256' },
          { name: 'priceUpdatedAt', internalType: 'uint256', type: 'uint256' },
          { name: 'isLiquidatable', internalType: 'bool', type: 'bool' },
          { name: 'priceStale', internalType: 'bool', type: 'bool' },
        ],
      },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'collateralAsset',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'collateralDecimals',
    outputs: [{ name: '', internalType: 'uint8', type: 'uint8' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'controller',
    outputs: [
      {
        name: '',
        internalType: 'contract ILendingController',
        type: 'address',
      },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'debtAsset',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'debtDecimals',
    outputs: [{ name: '', internalType: 'uint8', type: 'uint8' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'marketData',
    outputs: [
      {
        name: 'data',
        internalType: 'struct MarketData',
        type: 'tuple',
        components: [
          { name: 'totalSupplied', internalType: 'uint256', type: 'uint256' },
          { name: 'totalBorrowed', internalType: 'uint256', type: 'uint256' },
          {
            name: 'availableLiquidity',
            internalType: 'uint256',
            type: 'uint256',
          },
          { name: 'utilizationBps', internalType: 'uint256', type: 'uint256' },
          {
            name: 'supplyRatePerSecond',
            internalType: 'uint256',
            type: 'uint256',
          },
          {
            name: 'borrowRatePerSecond',
            internalType: 'uint256',
            type: 'uint256',
          },
          { name: 'supplyAprBps', internalType: 'uint256', type: 'uint256' },
          { name: 'borrowAprBps', internalType: 'uint256', type: 'uint256' },
          { name: 'supplyIndex', internalType: 'uint256', type: 'uint256' },
          { name: 'borrowIndex', internalType: 'uint256', type: 'uint256' },
          { name: 'maxLtvBps', internalType: 'uint256', type: 'uint256' },
          {
            name: 'liquidationThresholdBps',
            internalType: 'uint256',
            type: 'uint256',
          },
          {
            name: 'liquidationBonusBps',
            internalType: 'uint256',
            type: 'uint256',
          },
          {
            name: 'kinkUtilizationBps',
            internalType: 'uint256',
            type: 'uint256',
          },
          {
            name: 'reserveFactorBps',
            internalType: 'uint256',
            type: 'uint256',
          },
          { name: 'minDeposit', internalType: 'uint256', type: 'uint256' },
          { name: 'accruedReserves', internalType: 'uint256', type: 'uint256' },
          { name: 'depositsPaused', internalType: 'bool', type: 'bool' },
          { name: 'borrowPaused', internalType: 'bool', type: 'bool' },
        ],
      },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'oracle',
    outputs: [
      { name: '', internalType: 'contract IPriceOracle', type: 'address' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'pool',
    outputs: [
      { name: '', internalType: 'contract ILendingPool', type: 'address' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'vault',
    outputs: [
      { name: '', internalType: 'contract ICollateralVault', type: 'address' },
    ],
    stateMutability: 'view',
  },
  { type: 'error', inputs: [], name: 'ZeroAddress' },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// PriceOracleAdapter
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

export const priceOracleAdapterAbi = [
  {
    type: 'constructor',
    inputs: [
      { name: 'owner', internalType: 'address', type: 'address' },
      { name: 'startingMaxPriceAge', internalType: 'uint32', type: 'uint32' },
    ],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'REQUIRED_FEED_DECIMALS',
    outputs: [{ name: '', internalType: 'uint8', type: 'uint8' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'asset', internalType: 'address', type: 'address' }],
    name: 'feedOf',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'asset', internalType: 'address', type: 'address' }],
    name: 'getPrice',
    outputs: [
      { name: 'price', internalType: 'uint256', type: 'uint256' },
      { name: 'decimals', internalType: 'uint8', type: 'uint8' },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'asset', internalType: 'address', type: 'address' }],
    name: 'isStale',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'maxPriceAge',
    outputs: [{ name: '', internalType: 'uint32', type: 'uint32' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'owner',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'asset', internalType: 'address', type: 'address' }],
    name: 'readPrice',
    outputs: [
      {
        name: '',
        internalType: 'struct PriceData',
        type: 'tuple',
        components: [
          { name: 'price', internalType: 'uint256', type: 'uint256' },
          { name: 'decimals', internalType: 'uint8', type: 'uint8' },
          { name: 'updatedAt', internalType: 'uint256', type: 'uint256' },
          { name: 'isStale', internalType: 'bool', type: 'bool' },
          { name: 'isValid', internalType: 'bool', type: 'bool' },
        ],
      },
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'asset', internalType: 'address', type: 'address' },
      { name: 'aggregator', internalType: 'address', type: 'address' },
    ],
    name: 'setFeed',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'newAge', internalType: 'uint32', type: 'uint32' }],
    name: 'setMaxPriceAge',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'newOwner', internalType: 'address', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'asset',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'aggregator',
        internalType: 'address',
        type: 'address',
        indexed: false,
      },
      {
        name: 'decimals',
        internalType: 'uint8',
        type: 'uint8',
        indexed: false,
      },
    ],
    name: 'FeedChanged',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousAge',
        internalType: 'uint32',
        type: 'uint32',
        indexed: false,
      },
      {
        name: 'newAge',
        internalType: 'uint32',
        type: 'uint32',
        indexed: false,
      },
    ],
    name: 'MaxPriceAgeChanged',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'previousOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'newOwner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'OwnershipTransferred',
  },
  { type: 'error', inputs: [], name: 'InvalidRiskSettings' },
  {
    type: 'error',
    inputs: [{ name: 'owner', internalType: 'address', type: 'address' }],
    name: 'OwnableInvalidOwner',
  },
  {
    type: 'error',
    inputs: [{ name: 'account', internalType: 'address', type: 'address' }],
    name: 'OwnableUnauthorizedAccount',
  },
  {
    type: 'error',
    inputs: [{ name: 'asset', internalType: 'address', type: 'address' }],
    name: 'PriceIsInvalid',
  },
  {
    type: 'error',
    inputs: [
      { name: 'asset', internalType: 'address', type: 'address' },
      { name: 'updatedAt', internalType: 'uint256', type: 'uint256' },
      { name: 'maxAge', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'PriceIsStale',
  },
  {
    type: 'error',
    inputs: [{ name: 'value', internalType: 'int256', type: 'int256' }],
    name: 'SafeCastOverflowedIntToUint',
  },
  {
    type: 'error',
    inputs: [
      { name: 'aggregator', internalType: 'address', type: 'address' },
      { name: 'provided', internalType: 'uint8', type: 'uint8' },
      { name: 'expected', internalType: 'uint8', type: 'uint8' },
    ],
    name: 'UnsupportedFeedDecimals',
  },
  { type: 'error', inputs: [], name: 'ZeroAddress' },
] as const

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// React
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link collateralVaultAbi}__
 */
export const useReadCollateralVault = /*#__PURE__*/ createUseReadContract({
  abi: collateralVaultAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"collateralAsset"`
 */
export const useReadCollateralVaultCollateralAsset =
  /*#__PURE__*/ createUseReadContract({
    abi: collateralVaultAbi,
    functionName: 'collateralAsset',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"collateralOf"`
 */
export const useReadCollateralVaultCollateralOf =
  /*#__PURE__*/ createUseReadContract({
    abi: collateralVaultAbi,
    functionName: 'collateralOf',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"collateralToken"`
 */
export const useReadCollateralVaultCollateralToken =
  /*#__PURE__*/ createUseReadContract({
    abi: collateralVaultAbi,
    functionName: 'collateralToken',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"controller"`
 */
export const useReadCollateralVaultController =
  /*#__PURE__*/ createUseReadContract({
    abi: collateralVaultAbi,
    functionName: 'controller',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"liquidationManager"`
 */
export const useReadCollateralVaultLiquidationManager =
  /*#__PURE__*/ createUseReadContract({
    abi: collateralVaultAbi,
    functionName: 'liquidationManager',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"owner"`
 */
export const useReadCollateralVaultOwner = /*#__PURE__*/ createUseReadContract({
  abi: collateralVaultAbi,
  functionName: 'owner',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"totalCollateral"`
 */
export const useReadCollateralVaultTotalCollateral =
  /*#__PURE__*/ createUseReadContract({
    abi: collateralVaultAbi,
    functionName: 'totalCollateral',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link collateralVaultAbi}__
 */
export const useWriteCollateralVault = /*#__PURE__*/ createUseWriteContract({
  abi: collateralVaultAbi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"depositCollateral"`
 */
export const useWriteCollateralVaultDepositCollateral =
  /*#__PURE__*/ createUseWriteContract({
    abi: collateralVaultAbi,
    functionName: 'depositCollateral',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"linkController"`
 */
export const useWriteCollateralVaultLinkController =
  /*#__PURE__*/ createUseWriteContract({
    abi: collateralVaultAbi,
    functionName: 'linkController',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"linkLiquidationManager"`
 */
export const useWriteCollateralVaultLinkLiquidationManager =
  /*#__PURE__*/ createUseWriteContract({
    abi: collateralVaultAbi,
    functionName: 'linkLiquidationManager',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useWriteCollateralVaultRenounceOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: collateralVaultAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"seize"`
 */
export const useWriteCollateralVaultSeize =
  /*#__PURE__*/ createUseWriteContract({
    abi: collateralVaultAbi,
    functionName: 'seize',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useWriteCollateralVaultTransferOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: collateralVaultAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"withdrawCollateral"`
 */
export const useWriteCollateralVaultWithdrawCollateral =
  /*#__PURE__*/ createUseWriteContract({
    abi: collateralVaultAbi,
    functionName: 'withdrawCollateral',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link collateralVaultAbi}__
 */
export const useSimulateCollateralVault =
  /*#__PURE__*/ createUseSimulateContract({ abi: collateralVaultAbi })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"depositCollateral"`
 */
export const useSimulateCollateralVaultDepositCollateral =
  /*#__PURE__*/ createUseSimulateContract({
    abi: collateralVaultAbi,
    functionName: 'depositCollateral',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"linkController"`
 */
export const useSimulateCollateralVaultLinkController =
  /*#__PURE__*/ createUseSimulateContract({
    abi: collateralVaultAbi,
    functionName: 'linkController',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"linkLiquidationManager"`
 */
export const useSimulateCollateralVaultLinkLiquidationManager =
  /*#__PURE__*/ createUseSimulateContract({
    abi: collateralVaultAbi,
    functionName: 'linkLiquidationManager',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useSimulateCollateralVaultRenounceOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: collateralVaultAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"seize"`
 */
export const useSimulateCollateralVaultSeize =
  /*#__PURE__*/ createUseSimulateContract({
    abi: collateralVaultAbi,
    functionName: 'seize',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useSimulateCollateralVaultTransferOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: collateralVaultAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link collateralVaultAbi}__ and `functionName` set to `"withdrawCollateral"`
 */
export const useSimulateCollateralVaultWithdrawCollateral =
  /*#__PURE__*/ createUseSimulateContract({
    abi: collateralVaultAbi,
    functionName: 'withdrawCollateral',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link collateralVaultAbi}__
 */
export const useWatchCollateralVaultEvent =
  /*#__PURE__*/ createUseWatchContractEvent({ abi: collateralVaultAbi })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link collateralVaultAbi}__ and `eventName` set to `"CollateralDeposited"`
 */
export const useWatchCollateralVaultCollateralDepositedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: collateralVaultAbi,
    eventName: 'CollateralDeposited',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link collateralVaultAbi}__ and `eventName` set to `"CollateralSeized"`
 */
export const useWatchCollateralVaultCollateralSeizedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: collateralVaultAbi,
    eventName: 'CollateralSeized',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link collateralVaultAbi}__ and `eventName` set to `"CollateralWithdrawn"`
 */
export const useWatchCollateralVaultCollateralWithdrawnEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: collateralVaultAbi,
    eventName: 'CollateralWithdrawn',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link collateralVaultAbi}__ and `eventName` set to `"ControllerLinked"`
 */
export const useWatchCollateralVaultControllerLinkedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: collateralVaultAbi,
    eventName: 'ControllerLinked',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link collateralVaultAbi}__ and `eventName` set to `"LiquidationManagerLinked"`
 */
export const useWatchCollateralVaultLiquidationManagerLinkedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: collateralVaultAbi,
    eventName: 'LiquidationManagerLinked',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link collateralVaultAbi}__ and `eventName` set to `"OwnershipTransferred"`
 */
export const useWatchCollateralVaultOwnershipTransferredEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: collateralVaultAbi,
    eventName: 'OwnershipTransferred',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link interestRateModelAbi}__
 */
export const useReadInterestRateModel = /*#__PURE__*/ createUseReadContract({
  abi: interestRateModelAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link interestRateModelAbi}__ and `functionName` set to `"borrowAprBps"`
 */
export const useReadInterestRateModelBorrowAprBps =
  /*#__PURE__*/ createUseReadContract({
    abi: interestRateModelAbi,
    functionName: 'borrowAprBps',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link interestRateModelAbi}__ and `functionName` set to `"borrowRatePerSecond"`
 */
export const useReadInterestRateModelBorrowRatePerSecond =
  /*#__PURE__*/ createUseReadContract({
    abi: interestRateModelAbi,
    functionName: 'borrowRatePerSecond',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link interestRateModelAbi}__ and `functionName` set to `"curve"`
 */
export const useReadInterestRateModelCurve =
  /*#__PURE__*/ createUseReadContract({
    abi: interestRateModelAbi,
    functionName: 'curve',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link interestRateModelAbi}__ and `functionName` set to `"owner"`
 */
export const useReadInterestRateModelOwner =
  /*#__PURE__*/ createUseReadContract({
    abi: interestRateModelAbi,
    functionName: 'owner',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link interestRateModelAbi}__ and `functionName` set to `"supplyAprBps"`
 */
export const useReadInterestRateModelSupplyAprBps =
  /*#__PURE__*/ createUseReadContract({
    abi: interestRateModelAbi,
    functionName: 'supplyAprBps',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link interestRateModelAbi}__ and `functionName` set to `"supplyRatePerSecond"`
 */
export const useReadInterestRateModelSupplyRatePerSecond =
  /*#__PURE__*/ createUseReadContract({
    abi: interestRateModelAbi,
    functionName: 'supplyRatePerSecond',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link interestRateModelAbi}__ and `functionName` set to `"utilizationBps"`
 */
export const useReadInterestRateModelUtilizationBps =
  /*#__PURE__*/ createUseReadContract({
    abi: interestRateModelAbi,
    functionName: 'utilizationBps',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link interestRateModelAbi}__
 */
export const useWriteInterestRateModel = /*#__PURE__*/ createUseWriteContract({
  abi: interestRateModelAbi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link interestRateModelAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useWriteInterestRateModelRenounceOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: interestRateModelAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link interestRateModelAbi}__ and `functionName` set to `"setCurve"`
 */
export const useWriteInterestRateModelSetCurve =
  /*#__PURE__*/ createUseWriteContract({
    abi: interestRateModelAbi,
    functionName: 'setCurve',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link interestRateModelAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useWriteInterestRateModelTransferOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: interestRateModelAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link interestRateModelAbi}__
 */
export const useSimulateInterestRateModel =
  /*#__PURE__*/ createUseSimulateContract({ abi: interestRateModelAbi })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link interestRateModelAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useSimulateInterestRateModelRenounceOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: interestRateModelAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link interestRateModelAbi}__ and `functionName` set to `"setCurve"`
 */
export const useSimulateInterestRateModelSetCurve =
  /*#__PURE__*/ createUseSimulateContract({
    abi: interestRateModelAbi,
    functionName: 'setCurve',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link interestRateModelAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useSimulateInterestRateModelTransferOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: interestRateModelAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link interestRateModelAbi}__
 */
export const useWatchInterestRateModelEvent =
  /*#__PURE__*/ createUseWatchContractEvent({ abi: interestRateModelAbi })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link interestRateModelAbi}__ and `eventName` set to `"CurveChanged"`
 */
export const useWatchInterestRateModelCurveChangedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: interestRateModelAbi,
    eventName: 'CurveChanged',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link interestRateModelAbi}__ and `eventName` set to `"OwnershipTransferred"`
 */
export const useWatchInterestRateModelOwnershipTransferredEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: interestRateModelAbi,
    eventName: 'OwnershipTransferred',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__
 */
export const useReadLendingController = /*#__PURE__*/ createUseReadContract({
  abi: lendingControllerAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"borrowPaused"`
 */
export const useReadLendingControllerBorrowPaused =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'borrowPaused',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"collateralAsset"`
 */
export const useReadLendingControllerCollateralAsset =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'collateralAsset',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"collateralDecimals"`
 */
export const useReadLendingControllerCollateralDecimals =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'collateralDecimals',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"collateralValueOf"`
 */
export const useReadLendingControllerCollateralValueOf =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'collateralValueOf',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"debtAsset"`
 */
export const useReadLendingControllerDebtAsset =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'debtAsset',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"debtDecimals"`
 */
export const useReadLendingControllerDebtDecimals =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'debtDecimals',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"debtOf"`
 */
export const useReadLendingControllerDebtOf =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'debtOf',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"debtValueOf"`
 */
export const useReadLendingControllerDebtValueOf =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'debtValueOf',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"healthFactorBps"`
 */
export const useReadLendingControllerHealthFactorBps =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'healthFactorBps',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"isLiquidatable"`
 */
export const useReadLendingControllerIsLiquidatable =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'isLiquidatable',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"liquidationBonusBps"`
 */
export const useReadLendingControllerLiquidationBonusBps =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'liquidationBonusBps',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"liquidationThresholdBps"`
 */
export const useReadLendingControllerLiquidationThresholdBps =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'liquidationThresholdBps',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"maxBorrowable"`
 */
export const useReadLendingControllerMaxBorrowable =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'maxBorrowable',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"maxLtvBps"`
 */
export const useReadLendingControllerMaxLtvBps =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'maxLtvBps',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"maxWithdrawableCollateral"`
 */
export const useReadLendingControllerMaxWithdrawableCollateral =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'maxWithdrawableCollateral',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"oracle"`
 */
export const useReadLendingControllerOracle =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'oracle',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"owner"`
 */
export const useReadLendingControllerOwner =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'owner',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"pool"`
 */
export const useReadLendingControllerPool = /*#__PURE__*/ createUseReadContract(
  { abi: lendingControllerAbi, functionName: 'pool' },
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"vault"`
 */
export const useReadLendingControllerVault =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingControllerAbi,
    functionName: 'vault',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingControllerAbi}__
 */
export const useWriteLendingController = /*#__PURE__*/ createUseWriteContract({
  abi: lendingControllerAbi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"borrow"`
 */
export const useWriteLendingControllerBorrow =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingControllerAbi,
    functionName: 'borrow',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useWriteLendingControllerRenounceOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingControllerAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"repay"`
 */
export const useWriteLendingControllerRepay =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingControllerAbi,
    functionName: 'repay',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"repayAll"`
 */
export const useWriteLendingControllerRepayAll =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingControllerAbi,
    functionName: 'repayAll',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"setBorrowPaused"`
 */
export const useWriteLendingControllerSetBorrowPaused =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingControllerAbi,
    functionName: 'setBorrowPaused',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"setRiskSettings"`
 */
export const useWriteLendingControllerSetRiskSettings =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingControllerAbi,
    functionName: 'setRiskSettings',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useWriteLendingControllerTransferOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingControllerAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingControllerAbi}__
 */
export const useSimulateLendingController =
  /*#__PURE__*/ createUseSimulateContract({ abi: lendingControllerAbi })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"borrow"`
 */
export const useSimulateLendingControllerBorrow =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingControllerAbi,
    functionName: 'borrow',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useSimulateLendingControllerRenounceOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingControllerAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"repay"`
 */
export const useSimulateLendingControllerRepay =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingControllerAbi,
    functionName: 'repay',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"repayAll"`
 */
export const useSimulateLendingControllerRepayAll =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingControllerAbi,
    functionName: 'repayAll',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"setBorrowPaused"`
 */
export const useSimulateLendingControllerSetBorrowPaused =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingControllerAbi,
    functionName: 'setBorrowPaused',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"setRiskSettings"`
 */
export const useSimulateLendingControllerSetRiskSettings =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingControllerAbi,
    functionName: 'setRiskSettings',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingControllerAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useSimulateLendingControllerTransferOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingControllerAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingControllerAbi}__
 */
export const useWatchLendingControllerEvent =
  /*#__PURE__*/ createUseWatchContractEvent({ abi: lendingControllerAbi })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingControllerAbi}__ and `eventName` set to `"Borrow"`
 */
export const useWatchLendingControllerBorrowEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingControllerAbi,
    eventName: 'Borrow',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingControllerAbi}__ and `eventName` set to `"BorrowPausedChanged"`
 */
export const useWatchLendingControllerBorrowPausedChangedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingControllerAbi,
    eventName: 'BorrowPausedChanged',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingControllerAbi}__ and `eventName` set to `"OwnershipTransferred"`
 */
export const useWatchLendingControllerOwnershipTransferredEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingControllerAbi,
    eventName: 'OwnershipTransferred',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingControllerAbi}__ and `eventName` set to `"Repay"`
 */
export const useWatchLendingControllerRepayEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingControllerAbi,
    eventName: 'Repay',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingControllerAbi}__ and `eventName` set to `"RiskSettingsChanged"`
 */
export const useWatchLendingControllerRiskSettingsChangedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingControllerAbi,
    eventName: 'RiskSettingsChanged',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__
 */
export const useReadLendingPool = /*#__PURE__*/ createUseReadContract({
  abi: lendingPoolAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"accruedReserves"`
 */
export const useReadLendingPoolAccruedReserves =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'accruedReserves',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"asset"`
 */
export const useReadLendingPoolAsset = /*#__PURE__*/ createUseReadContract({
  abi: lendingPoolAbi,
  functionName: 'asset',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"assetToken"`
 */
export const useReadLendingPoolAssetToken = /*#__PURE__*/ createUseReadContract(
  { abi: lendingPoolAbi, functionName: 'assetToken' },
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"availableLiquidity"`
 */
export const useReadLendingPoolAvailableLiquidity =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'availableLiquidity',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"balanceOfAssets"`
 */
export const useReadLendingPoolBalanceOfAssets =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'balanceOfAssets',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"borrowIndex"`
 */
export const useReadLendingPoolBorrowIndex =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'borrowIndex',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"controller"`
 */
export const useReadLendingPoolController = /*#__PURE__*/ createUseReadContract(
  { abi: lendingPoolAbi, functionName: 'controller' },
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"debtOf"`
 */
export const useReadLendingPoolDebtOf = /*#__PURE__*/ createUseReadContract({
  abi: lendingPoolAbi,
  functionName: 'debtOf',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"debtSharesOf"`
 */
export const useReadLendingPoolDebtSharesOf =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'debtSharesOf',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"depositsPaused"`
 */
export const useReadLendingPoolDepositsPaused =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'depositsPaused',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"lastAccrualTimestamp"`
 */
export const useReadLendingPoolLastAccrualTimestamp =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'lastAccrualTimestamp',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"liquidationManager"`
 */
export const useReadLendingPoolLiquidationManager =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'liquidationManager',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"maxWithdrawable"`
 */
export const useReadLendingPoolMaxWithdrawable =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'maxWithdrawable',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"minDeposit"`
 */
export const useReadLendingPoolMinDeposit = /*#__PURE__*/ createUseReadContract(
  { abi: lendingPoolAbi, functionName: 'minDeposit' },
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"owner"`
 */
export const useReadLendingPoolOwner = /*#__PURE__*/ createUseReadContract({
  abi: lendingPoolAbi,
  functionName: 'owner',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"previewAccrual"`
 */
export const useReadLendingPoolPreviewAccrual =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'previewAccrual',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"rateModel"`
 */
export const useReadLendingPoolRateModel = /*#__PURE__*/ createUseReadContract({
  abi: lendingPoolAbi,
  functionName: 'rateModel',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"reserveFactorBps"`
 */
export const useReadLendingPoolReserveFactorBps =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'reserveFactorBps',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"sharesOf"`
 */
export const useReadLendingPoolSharesOf = /*#__PURE__*/ createUseReadContract({
  abi: lendingPoolAbi,
  functionName: 'sharesOf',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"supplyIndex"`
 */
export const useReadLendingPoolSupplyIndex =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'supplyIndex',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"totalBorrowed"`
 */
export const useReadLendingPoolTotalBorrowed =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'totalBorrowed',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"totalDebtShares"`
 */
export const useReadLendingPoolTotalDebtShares =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'totalDebtShares',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"totalSupplied"`
 */
export const useReadLendingPoolTotalSupplied =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'totalSupplied',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"totalSupplyShares"`
 */
export const useReadLendingPoolTotalSupplyShares =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'totalSupplyShares',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"utilizationBps"`
 */
export const useReadLendingPoolUtilizationBps =
  /*#__PURE__*/ createUseReadContract({
    abi: lendingPoolAbi,
    functionName: 'utilizationBps',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__
 */
export const useWriteLendingPool = /*#__PURE__*/ createUseWriteContract({
  abi: lendingPoolAbi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"accrueInterest"`
 */
export const useWriteLendingPoolAccrueInterest =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingPoolAbi,
    functionName: 'accrueInterest',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"borrowFor"`
 */
export const useWriteLendingPoolBorrowFor =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingPoolAbi,
    functionName: 'borrowFor',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"collectAllReserves"`
 */
export const useWriteLendingPoolCollectAllReserves =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingPoolAbi,
    functionName: 'collectAllReserves',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"collectReserves"`
 */
export const useWriteLendingPoolCollectReserves =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingPoolAbi,
    functionName: 'collectReserves',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"deposit"`
 */
export const useWriteLendingPoolDeposit = /*#__PURE__*/ createUseWriteContract({
  abi: lendingPoolAbi,
  functionName: 'deposit',
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"linkController"`
 */
export const useWriteLendingPoolLinkController =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingPoolAbi,
    functionName: 'linkController',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"linkLiquidationManager"`
 */
export const useWriteLendingPoolLinkLiquidationManager =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingPoolAbi,
    functionName: 'linkLiquidationManager',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"redeemShares"`
 */
export const useWriteLendingPoolRedeemShares =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingPoolAbi,
    functionName: 'redeemShares',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useWriteLendingPoolRenounceOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingPoolAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"repayAllFor"`
 */
export const useWriteLendingPoolRepayAllFor =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingPoolAbi,
    functionName: 'repayAllFor',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"repayFor"`
 */
export const useWriteLendingPoolRepayFor = /*#__PURE__*/ createUseWriteContract(
  { abi: lendingPoolAbi, functionName: 'repayFor' },
)

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"setDepositsPaused"`
 */
export const useWriteLendingPoolSetDepositsPaused =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingPoolAbi,
    functionName: 'setDepositsPaused',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"setMinDeposit"`
 */
export const useWriteLendingPoolSetMinDeposit =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingPoolAbi,
    functionName: 'setMinDeposit',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"setRateModel"`
 */
export const useWriteLendingPoolSetRateModel =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingPoolAbi,
    functionName: 'setRateModel',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"setReserveFactorBps"`
 */
export const useWriteLendingPoolSetReserveFactorBps =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingPoolAbi,
    functionName: 'setReserveFactorBps',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useWriteLendingPoolTransferOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: lendingPoolAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"withdraw"`
 */
export const useWriteLendingPoolWithdraw = /*#__PURE__*/ createUseWriteContract(
  { abi: lendingPoolAbi, functionName: 'withdraw' },
)

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__
 */
export const useSimulateLendingPool = /*#__PURE__*/ createUseSimulateContract({
  abi: lendingPoolAbi,
})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"accrueInterest"`
 */
export const useSimulateLendingPoolAccrueInterest =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'accrueInterest',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"borrowFor"`
 */
export const useSimulateLendingPoolBorrowFor =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'borrowFor',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"collectAllReserves"`
 */
export const useSimulateLendingPoolCollectAllReserves =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'collectAllReserves',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"collectReserves"`
 */
export const useSimulateLendingPoolCollectReserves =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'collectReserves',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"deposit"`
 */
export const useSimulateLendingPoolDeposit =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'deposit',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"linkController"`
 */
export const useSimulateLendingPoolLinkController =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'linkController',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"linkLiquidationManager"`
 */
export const useSimulateLendingPoolLinkLiquidationManager =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'linkLiquidationManager',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"redeemShares"`
 */
export const useSimulateLendingPoolRedeemShares =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'redeemShares',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useSimulateLendingPoolRenounceOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"repayAllFor"`
 */
export const useSimulateLendingPoolRepayAllFor =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'repayAllFor',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"repayFor"`
 */
export const useSimulateLendingPoolRepayFor =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'repayFor',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"setDepositsPaused"`
 */
export const useSimulateLendingPoolSetDepositsPaused =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'setDepositsPaused',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"setMinDeposit"`
 */
export const useSimulateLendingPoolSetMinDeposit =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'setMinDeposit',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"setRateModel"`
 */
export const useSimulateLendingPoolSetRateModel =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'setRateModel',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"setReserveFactorBps"`
 */
export const useSimulateLendingPoolSetReserveFactorBps =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'setReserveFactorBps',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useSimulateLendingPoolTransferOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link lendingPoolAbi}__ and `functionName` set to `"withdraw"`
 */
export const useSimulateLendingPoolWithdraw =
  /*#__PURE__*/ createUseSimulateContract({
    abi: lendingPoolAbi,
    functionName: 'withdraw',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingPoolAbi}__
 */
export const useWatchLendingPoolEvent =
  /*#__PURE__*/ createUseWatchContractEvent({ abi: lendingPoolAbi })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingPoolAbi}__ and `eventName` set to `"ControllerLinked"`
 */
export const useWatchLendingPoolControllerLinkedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingPoolAbi,
    eventName: 'ControllerLinked',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingPoolAbi}__ and `eventName` set to `"Deposit"`
 */
export const useWatchLendingPoolDepositEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingPoolAbi,
    eventName: 'Deposit',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingPoolAbi}__ and `eventName` set to `"DepositsPausedChanged"`
 */
export const useWatchLendingPoolDepositsPausedChangedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingPoolAbi,
    eventName: 'DepositsPausedChanged',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingPoolAbi}__ and `eventName` set to `"InterestAccrued"`
 */
export const useWatchLendingPoolInterestAccruedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingPoolAbi,
    eventName: 'InterestAccrued',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingPoolAbi}__ and `eventName` set to `"LiquidationManagerLinked"`
 */
export const useWatchLendingPoolLiquidationManagerLinkedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingPoolAbi,
    eventName: 'LiquidationManagerLinked',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingPoolAbi}__ and `eventName` set to `"MinDepositChanged"`
 */
export const useWatchLendingPoolMinDepositChangedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingPoolAbi,
    eventName: 'MinDepositChanged',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingPoolAbi}__ and `eventName` set to `"OwnershipTransferred"`
 */
export const useWatchLendingPoolOwnershipTransferredEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingPoolAbi,
    eventName: 'OwnershipTransferred',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingPoolAbi}__ and `eventName` set to `"RateModelChanged"`
 */
export const useWatchLendingPoolRateModelChangedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingPoolAbi,
    eventName: 'RateModelChanged',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingPoolAbi}__ and `eventName` set to `"ReserveFactorChanged"`
 */
export const useWatchLendingPoolReserveFactorChangedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingPoolAbi,
    eventName: 'ReserveFactorChanged',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingPoolAbi}__ and `eventName` set to `"ReservesCollected"`
 */
export const useWatchLendingPoolReservesCollectedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingPoolAbi,
    eventName: 'ReservesCollected',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link lendingPoolAbi}__ and `eventName` set to `"Withdraw"`
 */
export const useWatchLendingPoolWithdrawEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: lendingPoolAbi,
    eventName: 'Withdraw',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link liquidationManagerAbi}__
 */
export const useReadLiquidationManager = /*#__PURE__*/ createUseReadContract({
  abi: liquidationManagerAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link liquidationManagerAbi}__ and `functionName` set to `"collateralAsset"`
 */
export const useReadLiquidationManagerCollateralAsset =
  /*#__PURE__*/ createUseReadContract({
    abi: liquidationManagerAbi,
    functionName: 'collateralAsset',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link liquidationManagerAbi}__ and `functionName` set to `"collateralDecimals"`
 */
export const useReadLiquidationManagerCollateralDecimals =
  /*#__PURE__*/ createUseReadContract({
    abi: liquidationManagerAbi,
    functionName: 'collateralDecimals',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link liquidationManagerAbi}__ and `functionName` set to `"controller"`
 */
export const useReadLiquidationManagerController =
  /*#__PURE__*/ createUseReadContract({
    abi: liquidationManagerAbi,
    functionName: 'controller',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link liquidationManagerAbi}__ and `functionName` set to `"debtAsset"`
 */
export const useReadLiquidationManagerDebtAsset =
  /*#__PURE__*/ createUseReadContract({
    abi: liquidationManagerAbi,
    functionName: 'debtAsset',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link liquidationManagerAbi}__ and `functionName` set to `"debtDecimals"`
 */
export const useReadLiquidationManagerDebtDecimals =
  /*#__PURE__*/ createUseReadContract({
    abi: liquidationManagerAbi,
    functionName: 'debtDecimals',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link liquidationManagerAbi}__ and `functionName` set to `"isLiquidatable"`
 */
export const useReadLiquidationManagerIsLiquidatable =
  /*#__PURE__*/ createUseReadContract({
    abi: liquidationManagerAbi,
    functionName: 'isLiquidatable',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link liquidationManagerAbi}__ and `functionName` set to `"oracle"`
 */
export const useReadLiquidationManagerOracle =
  /*#__PURE__*/ createUseReadContract({
    abi: liquidationManagerAbi,
    functionName: 'oracle',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link liquidationManagerAbi}__ and `functionName` set to `"pool"`
 */
export const useReadLiquidationManagerPool =
  /*#__PURE__*/ createUseReadContract({
    abi: liquidationManagerAbi,
    functionName: 'pool',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link liquidationManagerAbi}__ and `functionName` set to `"previewLiquidation"`
 */
export const useReadLiquidationManagerPreviewLiquidation =
  /*#__PURE__*/ createUseReadContract({
    abi: liquidationManagerAbi,
    functionName: 'previewLiquidation',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link liquidationManagerAbi}__ and `functionName` set to `"vault"`
 */
export const useReadLiquidationManagerVault =
  /*#__PURE__*/ createUseReadContract({
    abi: liquidationManagerAbi,
    functionName: 'vault',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link liquidationManagerAbi}__
 */
export const useWriteLiquidationManager = /*#__PURE__*/ createUseWriteContract({
  abi: liquidationManagerAbi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link liquidationManagerAbi}__ and `functionName` set to `"liquidate"`
 */
export const useWriteLiquidationManagerLiquidate =
  /*#__PURE__*/ createUseWriteContract({
    abi: liquidationManagerAbi,
    functionName: 'liquidate',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link liquidationManagerAbi}__
 */
export const useSimulateLiquidationManager =
  /*#__PURE__*/ createUseSimulateContract({ abi: liquidationManagerAbi })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link liquidationManagerAbi}__ and `functionName` set to `"liquidate"`
 */
export const useSimulateLiquidationManagerLiquidate =
  /*#__PURE__*/ createUseSimulateContract({
    abi: liquidationManagerAbi,
    functionName: 'liquidate',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link liquidationManagerAbi}__
 */
export const useWatchLiquidationManagerEvent =
  /*#__PURE__*/ createUseWatchContractEvent({ abi: liquidationManagerAbi })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link liquidationManagerAbi}__ and `eventName` set to `"LiquidationExecuted"`
 */
export const useWatchLiquidationManagerLiquidationExecutedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: liquidationManagerAbi,
    eventName: 'LiquidationExecuted',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link mockAggregatorAbi}__
 */
export const useReadMockAggregator = /*#__PURE__*/ createUseReadContract({
  abi: mockAggregatorAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"decimals"`
 */
export const useReadMockAggregatorDecimals =
  /*#__PURE__*/ createUseReadContract({
    abi: mockAggregatorAbi,
    functionName: 'decimals',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"description"`
 */
export const useReadMockAggregatorDescription =
  /*#__PURE__*/ createUseReadContract({
    abi: mockAggregatorAbi,
    functionName: 'description',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"feedDown"`
 */
export const useReadMockAggregatorFeedDown =
  /*#__PURE__*/ createUseReadContract({
    abi: mockAggregatorAbi,
    functionName: 'feedDown',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"getRoundData"`
 */
export const useReadMockAggregatorGetRoundData =
  /*#__PURE__*/ createUseReadContract({
    abi: mockAggregatorAbi,
    functionName: 'getRoundData',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"latestRoundData"`
 */
export const useReadMockAggregatorLatestRoundData =
  /*#__PURE__*/ createUseReadContract({
    abi: mockAggregatorAbi,
    functionName: 'latestRoundData',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"latestRoundId"`
 */
export const useReadMockAggregatorLatestRoundId =
  /*#__PURE__*/ createUseReadContract({
    abi: mockAggregatorAbi,
    functionName: 'latestRoundId',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"version"`
 */
export const useReadMockAggregatorVersion = /*#__PURE__*/ createUseReadContract(
  { abi: mockAggregatorAbi, functionName: 'version' },
)

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link mockAggregatorAbi}__
 */
export const useWriteMockAggregator = /*#__PURE__*/ createUseWriteContract({
  abi: mockAggregatorAbi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"makeStale"`
 */
export const useWriteMockAggregatorMakeStale =
  /*#__PURE__*/ createUseWriteContract({
    abi: mockAggregatorAbi,
    functionName: 'makeStale',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"setFeedDown"`
 */
export const useWriteMockAggregatorSetFeedDown =
  /*#__PURE__*/ createUseWriteContract({
    abi: mockAggregatorAbi,
    functionName: 'setFeedDown',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"setIncompleteRound"`
 */
export const useWriteMockAggregatorSetIncompleteRound =
  /*#__PURE__*/ createUseWriteContract({
    abi: mockAggregatorAbi,
    functionName: 'setIncompleteRound',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"setPrice"`
 */
export const useWriteMockAggregatorSetPrice =
  /*#__PURE__*/ createUseWriteContract({
    abi: mockAggregatorAbi,
    functionName: 'setPrice',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"setPriceWithTimestamp"`
 */
export const useWriteMockAggregatorSetPriceWithTimestamp =
  /*#__PURE__*/ createUseWriteContract({
    abi: mockAggregatorAbi,
    functionName: 'setPriceWithTimestamp',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link mockAggregatorAbi}__
 */
export const useSimulateMockAggregator =
  /*#__PURE__*/ createUseSimulateContract({ abi: mockAggregatorAbi })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"makeStale"`
 */
export const useSimulateMockAggregatorMakeStale =
  /*#__PURE__*/ createUseSimulateContract({
    abi: mockAggregatorAbi,
    functionName: 'makeStale',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"setFeedDown"`
 */
export const useSimulateMockAggregatorSetFeedDown =
  /*#__PURE__*/ createUseSimulateContract({
    abi: mockAggregatorAbi,
    functionName: 'setFeedDown',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"setIncompleteRound"`
 */
export const useSimulateMockAggregatorSetIncompleteRound =
  /*#__PURE__*/ createUseSimulateContract({
    abi: mockAggregatorAbi,
    functionName: 'setIncompleteRound',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"setPrice"`
 */
export const useSimulateMockAggregatorSetPrice =
  /*#__PURE__*/ createUseSimulateContract({
    abi: mockAggregatorAbi,
    functionName: 'setPrice',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link mockAggregatorAbi}__ and `functionName` set to `"setPriceWithTimestamp"`
 */
export const useSimulateMockAggregatorSetPriceWithTimestamp =
  /*#__PURE__*/ createUseSimulateContract({
    abi: mockAggregatorAbi,
    functionName: 'setPriceWithTimestamp',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link mockAggregatorAbi}__
 */
export const useWatchMockAggregatorEvent =
  /*#__PURE__*/ createUseWatchContractEvent({ abi: mockAggregatorAbi })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link mockAggregatorAbi}__ and `eventName` set to `"AnswerRecorded"`
 */
export const useWatchMockAggregatorAnswerRecordedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: mockAggregatorAbi,
    eventName: 'AnswerRecorded',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link mockErc20Abi}__
 */
export const useReadMockErc20 = /*#__PURE__*/ createUseReadContract({
  abi: mockErc20Abi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"allowance"`
 */
export const useReadMockErc20Allowance = /*#__PURE__*/ createUseReadContract({
  abi: mockErc20Abi,
  functionName: 'allowance',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"balanceOf"`
 */
export const useReadMockErc20BalanceOf = /*#__PURE__*/ createUseReadContract({
  abi: mockErc20Abi,
  functionName: 'balanceOf',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"decimals"`
 */
export const useReadMockErc20Decimals = /*#__PURE__*/ createUseReadContract({
  abi: mockErc20Abi,
  functionName: 'decimals',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"name"`
 */
export const useReadMockErc20Name = /*#__PURE__*/ createUseReadContract({
  abi: mockErc20Abi,
  functionName: 'name',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"symbol"`
 */
export const useReadMockErc20Symbol = /*#__PURE__*/ createUseReadContract({
  abi: mockErc20Abi,
  functionName: 'symbol',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"totalSupply"`
 */
export const useReadMockErc20TotalSupply = /*#__PURE__*/ createUseReadContract({
  abi: mockErc20Abi,
  functionName: 'totalSupply',
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link mockErc20Abi}__
 */
export const useWriteMockErc20 = /*#__PURE__*/ createUseWriteContract({
  abi: mockErc20Abi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"approve"`
 */
export const useWriteMockErc20Approve = /*#__PURE__*/ createUseWriteContract({
  abi: mockErc20Abi,
  functionName: 'approve',
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"burn"`
 */
export const useWriteMockErc20Burn = /*#__PURE__*/ createUseWriteContract({
  abi: mockErc20Abi,
  functionName: 'burn',
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"mint"`
 */
export const useWriteMockErc20Mint = /*#__PURE__*/ createUseWriteContract({
  abi: mockErc20Abi,
  functionName: 'mint',
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"transfer"`
 */
export const useWriteMockErc20Transfer = /*#__PURE__*/ createUseWriteContract({
  abi: mockErc20Abi,
  functionName: 'transfer',
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"transferFrom"`
 */
export const useWriteMockErc20TransferFrom =
  /*#__PURE__*/ createUseWriteContract({
    abi: mockErc20Abi,
    functionName: 'transferFrom',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link mockErc20Abi}__
 */
export const useSimulateMockErc20 = /*#__PURE__*/ createUseSimulateContract({
  abi: mockErc20Abi,
})

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"approve"`
 */
export const useSimulateMockErc20Approve =
  /*#__PURE__*/ createUseSimulateContract({
    abi: mockErc20Abi,
    functionName: 'approve',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"burn"`
 */
export const useSimulateMockErc20Burn = /*#__PURE__*/ createUseSimulateContract(
  { abi: mockErc20Abi, functionName: 'burn' },
)

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"mint"`
 */
export const useSimulateMockErc20Mint = /*#__PURE__*/ createUseSimulateContract(
  { abi: mockErc20Abi, functionName: 'mint' },
)

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"transfer"`
 */
export const useSimulateMockErc20Transfer =
  /*#__PURE__*/ createUseSimulateContract({
    abi: mockErc20Abi,
    functionName: 'transfer',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link mockErc20Abi}__ and `functionName` set to `"transferFrom"`
 */
export const useSimulateMockErc20TransferFrom =
  /*#__PURE__*/ createUseSimulateContract({
    abi: mockErc20Abi,
    functionName: 'transferFrom',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link mockErc20Abi}__
 */
export const useWatchMockErc20Event = /*#__PURE__*/ createUseWatchContractEvent(
  { abi: mockErc20Abi },
)

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link mockErc20Abi}__ and `eventName` set to `"Approval"`
 */
export const useWatchMockErc20ApprovalEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: mockErc20Abi,
    eventName: 'Approval',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link mockErc20Abi}__ and `eventName` set to `"Transfer"`
 */
export const useWatchMockErc20TransferEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: mockErc20Abi,
    eventName: 'Transfer',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link positionLensAbi}__
 */
export const useReadPositionLens = /*#__PURE__*/ createUseReadContract({
  abi: positionLensAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link positionLensAbi}__ and `functionName` set to `"accountData"`
 */
export const useReadPositionLensAccountData =
  /*#__PURE__*/ createUseReadContract({
    abi: positionLensAbi,
    functionName: 'accountData',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link positionLensAbi}__ and `functionName` set to `"collateralAsset"`
 */
export const useReadPositionLensCollateralAsset =
  /*#__PURE__*/ createUseReadContract({
    abi: positionLensAbi,
    functionName: 'collateralAsset',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link positionLensAbi}__ and `functionName` set to `"collateralDecimals"`
 */
export const useReadPositionLensCollateralDecimals =
  /*#__PURE__*/ createUseReadContract({
    abi: positionLensAbi,
    functionName: 'collateralDecimals',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link positionLensAbi}__ and `functionName` set to `"controller"`
 */
export const useReadPositionLensController =
  /*#__PURE__*/ createUseReadContract({
    abi: positionLensAbi,
    functionName: 'controller',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link positionLensAbi}__ and `functionName` set to `"debtAsset"`
 */
export const useReadPositionLensDebtAsset = /*#__PURE__*/ createUseReadContract(
  { abi: positionLensAbi, functionName: 'debtAsset' },
)

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link positionLensAbi}__ and `functionName` set to `"debtDecimals"`
 */
export const useReadPositionLensDebtDecimals =
  /*#__PURE__*/ createUseReadContract({
    abi: positionLensAbi,
    functionName: 'debtDecimals',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link positionLensAbi}__ and `functionName` set to `"marketData"`
 */
export const useReadPositionLensMarketData =
  /*#__PURE__*/ createUseReadContract({
    abi: positionLensAbi,
    functionName: 'marketData',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link positionLensAbi}__ and `functionName` set to `"oracle"`
 */
export const useReadPositionLensOracle = /*#__PURE__*/ createUseReadContract({
  abi: positionLensAbi,
  functionName: 'oracle',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link positionLensAbi}__ and `functionName` set to `"pool"`
 */
export const useReadPositionLensPool = /*#__PURE__*/ createUseReadContract({
  abi: positionLensAbi,
  functionName: 'pool',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link positionLensAbi}__ and `functionName` set to `"vault"`
 */
export const useReadPositionLensVault = /*#__PURE__*/ createUseReadContract({
  abi: positionLensAbi,
  functionName: 'vault',
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__
 */
export const useReadPriceOracleAdapter = /*#__PURE__*/ createUseReadContract({
  abi: priceOracleAdapterAbi,
})

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `functionName` set to `"REQUIRED_FEED_DECIMALS"`
 */
export const useReadPriceOracleAdapterRequiredFeedDecimals =
  /*#__PURE__*/ createUseReadContract({
    abi: priceOracleAdapterAbi,
    functionName: 'REQUIRED_FEED_DECIMALS',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `functionName` set to `"feedOf"`
 */
export const useReadPriceOracleAdapterFeedOf =
  /*#__PURE__*/ createUseReadContract({
    abi: priceOracleAdapterAbi,
    functionName: 'feedOf',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `functionName` set to `"getPrice"`
 */
export const useReadPriceOracleAdapterGetPrice =
  /*#__PURE__*/ createUseReadContract({
    abi: priceOracleAdapterAbi,
    functionName: 'getPrice',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `functionName` set to `"isStale"`
 */
export const useReadPriceOracleAdapterIsStale =
  /*#__PURE__*/ createUseReadContract({
    abi: priceOracleAdapterAbi,
    functionName: 'isStale',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `functionName` set to `"maxPriceAge"`
 */
export const useReadPriceOracleAdapterMaxPriceAge =
  /*#__PURE__*/ createUseReadContract({
    abi: priceOracleAdapterAbi,
    functionName: 'maxPriceAge',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `functionName` set to `"owner"`
 */
export const useReadPriceOracleAdapterOwner =
  /*#__PURE__*/ createUseReadContract({
    abi: priceOracleAdapterAbi,
    functionName: 'owner',
  })

/**
 * Wraps __{@link useReadContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `functionName` set to `"readPrice"`
 */
export const useReadPriceOracleAdapterReadPrice =
  /*#__PURE__*/ createUseReadContract({
    abi: priceOracleAdapterAbi,
    functionName: 'readPrice',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__
 */
export const useWritePriceOracleAdapter = /*#__PURE__*/ createUseWriteContract({
  abi: priceOracleAdapterAbi,
})

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useWritePriceOracleAdapterRenounceOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: priceOracleAdapterAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `functionName` set to `"setFeed"`
 */
export const useWritePriceOracleAdapterSetFeed =
  /*#__PURE__*/ createUseWriteContract({
    abi: priceOracleAdapterAbi,
    functionName: 'setFeed',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `functionName` set to `"setMaxPriceAge"`
 */
export const useWritePriceOracleAdapterSetMaxPriceAge =
  /*#__PURE__*/ createUseWriteContract({
    abi: priceOracleAdapterAbi,
    functionName: 'setMaxPriceAge',
  })

/**
 * Wraps __{@link useWriteContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useWritePriceOracleAdapterTransferOwnership =
  /*#__PURE__*/ createUseWriteContract({
    abi: priceOracleAdapterAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__
 */
export const useSimulatePriceOracleAdapter =
  /*#__PURE__*/ createUseSimulateContract({ abi: priceOracleAdapterAbi })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `functionName` set to `"renounceOwnership"`
 */
export const useSimulatePriceOracleAdapterRenounceOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: priceOracleAdapterAbi,
    functionName: 'renounceOwnership',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `functionName` set to `"setFeed"`
 */
export const useSimulatePriceOracleAdapterSetFeed =
  /*#__PURE__*/ createUseSimulateContract({
    abi: priceOracleAdapterAbi,
    functionName: 'setFeed',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `functionName` set to `"setMaxPriceAge"`
 */
export const useSimulatePriceOracleAdapterSetMaxPriceAge =
  /*#__PURE__*/ createUseSimulateContract({
    abi: priceOracleAdapterAbi,
    functionName: 'setMaxPriceAge',
  })

/**
 * Wraps __{@link useSimulateContract}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `functionName` set to `"transferOwnership"`
 */
export const useSimulatePriceOracleAdapterTransferOwnership =
  /*#__PURE__*/ createUseSimulateContract({
    abi: priceOracleAdapterAbi,
    functionName: 'transferOwnership',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link priceOracleAdapterAbi}__
 */
export const useWatchPriceOracleAdapterEvent =
  /*#__PURE__*/ createUseWatchContractEvent({ abi: priceOracleAdapterAbi })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `eventName` set to `"FeedChanged"`
 */
export const useWatchPriceOracleAdapterFeedChangedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: priceOracleAdapterAbi,
    eventName: 'FeedChanged',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `eventName` set to `"MaxPriceAgeChanged"`
 */
export const useWatchPriceOracleAdapterMaxPriceAgeChangedEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: priceOracleAdapterAbi,
    eventName: 'MaxPriceAgeChanged',
  })

/**
 * Wraps __{@link useWatchContractEvent}__ with `abi` set to __{@link priceOracleAdapterAbi}__ and `eventName` set to `"OwnershipTransferred"`
 */
export const useWatchPriceOracleAdapterOwnershipTransferredEvent =
  /*#__PURE__*/ createUseWatchContractEvent({
    abi: priceOracleAdapterAbi,
    eventName: 'OwnershipTransferred',
  })
