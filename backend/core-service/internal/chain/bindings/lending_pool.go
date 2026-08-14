// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LendingPoolMetaData contains all meta data concerning the LendingPool contract.
var LendingPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolAsset\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"startingRateModel\",\"type\":\"address\",\"internalType\":\"contractIInterestRateModel\"},{\"name\":\"startingMinDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startingReserveFactorBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"accrueInterest\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"accruedReserves\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"asset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"assetToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"availableLiquidity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOfAssets\",\"inputs\":[{\"name\":\"lender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"borrowFor\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"borrowIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collectAllReserves\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"collected\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collectReserves\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"collected\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"controller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"debtOf\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"debtSharesOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositsPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastAccrualTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkController\",\"inputs\":[{\"name\":\"newController\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"linkLiquidationManager\",\"inputs\":[{\"name\":\"newLiquidationManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"liquidationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxWithdrawable\",\"inputs\":[{\"name\":\"lender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minDeposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewAccrual\",\"inputs\":[],\"outputs\":[{\"name\":\"nextSupplyIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nextBorrowIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reservesToAdd\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rateModel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIInterestRateModel\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeemShares\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"repayAllFor\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amountPaid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"repayFor\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reserveFactorBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDepositsPaused\",\"inputs\":[{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinDeposit\",\"inputs\":[{\"name\":\"newAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateModel\",\"inputs\":[{\"name\":\"newModel\",\"type\":\"address\",\"internalType\":\"contractIInterestRateModel\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setReserveFactorBps\",\"inputs\":[{\"name\":\"newBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sharesOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supplyIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalBorrowed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalDebtShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupplied\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupplyShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"utilizationBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ControllerLinked\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"lender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"supplyIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"totalSupplied\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositsPausedChanged\",\"inputs\":[{\"name\":\"isPaused\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InterestAccrued\",\"inputs\":[{\"name\":\"supplyIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"borrowIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"totalBorrowed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reservesAccrued\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidationManagerLinked\",\"inputs\":[{\"name\":\"liquidationManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinDepositChanged\",\"inputs\":[{\"name\":\"previousAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateModelChanged\",\"inputs\":[{\"name\":\"previousModel\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newModel\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReserveFactorChanged\",\"inputs\":[{\"name\":\"previousBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReservesCollected\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"remainingReserves\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"lender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"supplyIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"totalSupplied\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BelowMinimumDeposit\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ExceedsReserves\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ExceedsSupplyBalance\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRiskSettings\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotEnoughLiquidity\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]}]",
}

// LendingPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use LendingPoolMetaData.ABI instead.
var LendingPoolABI = LendingPoolMetaData.ABI

// LendingPool is an auto generated Go binding around an Ethereum contract.
type LendingPool struct {
	LendingPoolCaller     // Read-only binding to the contract
	LendingPoolTransactor // Write-only binding to the contract
	LendingPoolFilterer   // Log filterer for contract events
}

// LendingPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingPoolSession struct {
	Contract     *LendingPool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LendingPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingPoolCallerSession struct {
	Contract *LendingPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LendingPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingPoolTransactorSession struct {
	Contract     *LendingPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LendingPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingPoolRaw struct {
	Contract *LendingPool // Generic contract binding to access the raw methods on
}

// LendingPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingPoolCallerRaw struct {
	Contract *LendingPoolCaller // Generic read-only contract binding to access the raw methods on
}

// LendingPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingPoolTransactorRaw struct {
	Contract *LendingPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLendingPool creates a new instance of LendingPool, bound to a specific deployed contract.
func NewLendingPool(address common.Address, backend bind.ContractBackend) (*LendingPool, error) {
	contract, err := bindLendingPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LendingPool{LendingPoolCaller: LendingPoolCaller{contract: contract}, LendingPoolTransactor: LendingPoolTransactor{contract: contract}, LendingPoolFilterer: LendingPoolFilterer{contract: contract}}, nil
}

// NewLendingPoolCaller creates a new read-only instance of LendingPool, bound to a specific deployed contract.
func NewLendingPoolCaller(address common.Address, caller bind.ContractCaller) (*LendingPoolCaller, error) {
	contract, err := bindLendingPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolCaller{contract: contract}, nil
}

// NewLendingPoolTransactor creates a new write-only instance of LendingPool, bound to a specific deployed contract.
func NewLendingPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingPoolTransactor, error) {
	contract, err := bindLendingPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolTransactor{contract: contract}, nil
}

// NewLendingPoolFilterer creates a new log filterer instance of LendingPool, bound to a specific deployed contract.
func NewLendingPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingPoolFilterer, error) {
	contract, err := bindLendingPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingPoolFilterer{contract: contract}, nil
}

// bindLendingPool binds a generic wrapper to an already deployed contract.
func bindLendingPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LendingPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPool *LendingPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingPool.Contract.LendingPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPool *LendingPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPool.Contract.LendingPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPool *LendingPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPool.Contract.LendingPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPool *LendingPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPool *LendingPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPool *LendingPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPool.Contract.contract.Transact(opts, method, params...)
}

// AccruedReserves is a free data retrieval call binding the contract method 0x2891d7f2.
//
// Solidity: function accruedReserves() view returns(uint256)
func (_LendingPool *LendingPoolCaller) AccruedReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "accruedReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedReserves is a free data retrieval call binding the contract method 0x2891d7f2.
//
// Solidity: function accruedReserves() view returns(uint256)
func (_LendingPool *LendingPoolSession) AccruedReserves() (*big.Int, error) {
	return _LendingPool.Contract.AccruedReserves(&_LendingPool.CallOpts)
}

// AccruedReserves is a free data retrieval call binding the contract method 0x2891d7f2.
//
// Solidity: function accruedReserves() view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) AccruedReserves() (*big.Int, error) {
	return _LendingPool.Contract.AccruedReserves(&_LendingPool.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_LendingPool *LendingPoolCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_LendingPool *LendingPoolSession) Asset() (common.Address, error) {
	return _LendingPool.Contract.Asset(&_LendingPool.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_LendingPool *LendingPoolCallerSession) Asset() (common.Address, error) {
	return _LendingPool.Contract.Asset(&_LendingPool.CallOpts)
}

// AssetToken is a free data retrieval call binding the contract method 0x1083f761.
//
// Solidity: function assetToken() view returns(address)
func (_LendingPool *LendingPoolCaller) AssetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "assetToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetToken is a free data retrieval call binding the contract method 0x1083f761.
//
// Solidity: function assetToken() view returns(address)
func (_LendingPool *LendingPoolSession) AssetToken() (common.Address, error) {
	return _LendingPool.Contract.AssetToken(&_LendingPool.CallOpts)
}

// AssetToken is a free data retrieval call binding the contract method 0x1083f761.
//
// Solidity: function assetToken() view returns(address)
func (_LendingPool *LendingPoolCallerSession) AssetToken() (common.Address, error) {
	return _LendingPool.Contract.AssetToken(&_LendingPool.CallOpts)
}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_LendingPool *LendingPoolCaller) AvailableLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "availableLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_LendingPool *LendingPoolSession) AvailableLiquidity() (*big.Int, error) {
	return _LendingPool.Contract.AvailableLiquidity(&_LendingPool.CallOpts)
}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) AvailableLiquidity() (*big.Int, error) {
	return _LendingPool.Contract.AvailableLiquidity(&_LendingPool.CallOpts)
}

// BalanceOfAssets is a free data retrieval call binding the contract method 0x9159b206.
//
// Solidity: function balanceOfAssets(address lender) view returns(uint256)
func (_LendingPool *LendingPoolCaller) BalanceOfAssets(opts *bind.CallOpts, lender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "balanceOfAssets", lender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAssets is a free data retrieval call binding the contract method 0x9159b206.
//
// Solidity: function balanceOfAssets(address lender) view returns(uint256)
func (_LendingPool *LendingPoolSession) BalanceOfAssets(lender common.Address) (*big.Int, error) {
	return _LendingPool.Contract.BalanceOfAssets(&_LendingPool.CallOpts, lender)
}

// BalanceOfAssets is a free data retrieval call binding the contract method 0x9159b206.
//
// Solidity: function balanceOfAssets(address lender) view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) BalanceOfAssets(lender common.Address) (*big.Int, error) {
	return _LendingPool.Contract.BalanceOfAssets(&_LendingPool.CallOpts, lender)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_LendingPool *LendingPoolCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_LendingPool *LendingPoolSession) BorrowIndex() (*big.Int, error) {
	return _LendingPool.Contract.BorrowIndex(&_LendingPool.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) BorrowIndex() (*big.Int, error) {
	return _LendingPool.Contract.BorrowIndex(&_LendingPool.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_LendingPool *LendingPoolCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_LendingPool *LendingPoolSession) Controller() (common.Address, error) {
	return _LendingPool.Contract.Controller(&_LendingPool.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_LendingPool *LendingPoolCallerSession) Controller() (common.Address, error) {
	return _LendingPool.Contract.Controller(&_LendingPool.CallOpts)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address borrower) view returns(uint256)
func (_LendingPool *LendingPoolCaller) DebtOf(opts *bind.CallOpts, borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "debtOf", borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address borrower) view returns(uint256)
func (_LendingPool *LendingPoolSession) DebtOf(borrower common.Address) (*big.Int, error) {
	return _LendingPool.Contract.DebtOf(&_LendingPool.CallOpts, borrower)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address borrower) view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) DebtOf(borrower common.Address) (*big.Int, error) {
	return _LendingPool.Contract.DebtOf(&_LendingPool.CallOpts, borrower)
}

// DebtSharesOf is a free data retrieval call binding the contract method 0x3566dae0.
//
// Solidity: function debtSharesOf(address ) view returns(uint256)
func (_LendingPool *LendingPoolCaller) DebtSharesOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "debtSharesOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtSharesOf is a free data retrieval call binding the contract method 0x3566dae0.
//
// Solidity: function debtSharesOf(address ) view returns(uint256)
func (_LendingPool *LendingPoolSession) DebtSharesOf(arg0 common.Address) (*big.Int, error) {
	return _LendingPool.Contract.DebtSharesOf(&_LendingPool.CallOpts, arg0)
}

// DebtSharesOf is a free data retrieval call binding the contract method 0x3566dae0.
//
// Solidity: function debtSharesOf(address ) view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) DebtSharesOf(arg0 common.Address) (*big.Int, error) {
	return _LendingPool.Contract.DebtSharesOf(&_LendingPool.CallOpts, arg0)
}

// DepositsPaused is a free data retrieval call binding the contract method 0x60da3e83.
//
// Solidity: function depositsPaused() view returns(bool)
func (_LendingPool *LendingPoolCaller) DepositsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "depositsPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositsPaused is a free data retrieval call binding the contract method 0x60da3e83.
//
// Solidity: function depositsPaused() view returns(bool)
func (_LendingPool *LendingPoolSession) DepositsPaused() (bool, error) {
	return _LendingPool.Contract.DepositsPaused(&_LendingPool.CallOpts)
}

// DepositsPaused is a free data retrieval call binding the contract method 0x60da3e83.
//
// Solidity: function depositsPaused() view returns(bool)
func (_LendingPool *LendingPoolCallerSession) DepositsPaused() (bool, error) {
	return _LendingPool.Contract.DepositsPaused(&_LendingPool.CallOpts)
}

// LastAccrualTimestamp is a free data retrieval call binding the contract method 0xc0789dcb.
//
// Solidity: function lastAccrualTimestamp() view returns(uint256)
func (_LendingPool *LendingPoolCaller) LastAccrualTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "lastAccrualTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastAccrualTimestamp is a free data retrieval call binding the contract method 0xc0789dcb.
//
// Solidity: function lastAccrualTimestamp() view returns(uint256)
func (_LendingPool *LendingPoolSession) LastAccrualTimestamp() (*big.Int, error) {
	return _LendingPool.Contract.LastAccrualTimestamp(&_LendingPool.CallOpts)
}

// LastAccrualTimestamp is a free data retrieval call binding the contract method 0xc0789dcb.
//
// Solidity: function lastAccrualTimestamp() view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) LastAccrualTimestamp() (*big.Int, error) {
	return _LendingPool.Contract.LastAccrualTimestamp(&_LendingPool.CallOpts)
}

// LiquidationManager is a free data retrieval call binding the contract method 0x1ef3a04c.
//
// Solidity: function liquidationManager() view returns(address)
func (_LendingPool *LendingPoolCaller) LiquidationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "liquidationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidationManager is a free data retrieval call binding the contract method 0x1ef3a04c.
//
// Solidity: function liquidationManager() view returns(address)
func (_LendingPool *LendingPoolSession) LiquidationManager() (common.Address, error) {
	return _LendingPool.Contract.LiquidationManager(&_LendingPool.CallOpts)
}

// LiquidationManager is a free data retrieval call binding the contract method 0x1ef3a04c.
//
// Solidity: function liquidationManager() view returns(address)
func (_LendingPool *LendingPoolCallerSession) LiquidationManager() (common.Address, error) {
	return _LendingPool.Contract.LiquidationManager(&_LendingPool.CallOpts)
}

// MaxWithdrawable is a free data retrieval call binding the contract method 0xf2f11679.
//
// Solidity: function maxWithdrawable(address lender) view returns(uint256)
func (_LendingPool *LendingPoolCaller) MaxWithdrawable(opts *bind.CallOpts, lender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "maxWithdrawable", lender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdrawable is a free data retrieval call binding the contract method 0xf2f11679.
//
// Solidity: function maxWithdrawable(address lender) view returns(uint256)
func (_LendingPool *LendingPoolSession) MaxWithdrawable(lender common.Address) (*big.Int, error) {
	return _LendingPool.Contract.MaxWithdrawable(&_LendingPool.CallOpts, lender)
}

// MaxWithdrawable is a free data retrieval call binding the contract method 0xf2f11679.
//
// Solidity: function maxWithdrawable(address lender) view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) MaxWithdrawable(lender common.Address) (*big.Int, error) {
	return _LendingPool.Contract.MaxWithdrawable(&_LendingPool.CallOpts, lender)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_LendingPool *LendingPoolCaller) MinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "minDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_LendingPool *LendingPoolSession) MinDeposit() (*big.Int, error) {
	return _LendingPool.Contract.MinDeposit(&_LendingPool.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) MinDeposit() (*big.Int, error) {
	return _LendingPool.Contract.MinDeposit(&_LendingPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingPool *LendingPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingPool *LendingPoolSession) Owner() (common.Address, error) {
	return _LendingPool.Contract.Owner(&_LendingPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingPool *LendingPoolCallerSession) Owner() (common.Address, error) {
	return _LendingPool.Contract.Owner(&_LendingPool.CallOpts)
}

// PreviewAccrual is a free data retrieval call binding the contract method 0x37d5722a.
//
// Solidity: function previewAccrual() view returns(uint256 nextSupplyIndex, uint256 nextBorrowIndex, uint256 reservesToAdd)
func (_LendingPool *LendingPoolCaller) PreviewAccrual(opts *bind.CallOpts) (struct {
	NextSupplyIndex *big.Int
	NextBorrowIndex *big.Int
	ReservesToAdd   *big.Int
}, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "previewAccrual")

	outstruct := new(struct {
		NextSupplyIndex *big.Int
		NextBorrowIndex *big.Int
		ReservesToAdd   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NextSupplyIndex = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NextBorrowIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ReservesToAdd = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PreviewAccrual is a free data retrieval call binding the contract method 0x37d5722a.
//
// Solidity: function previewAccrual() view returns(uint256 nextSupplyIndex, uint256 nextBorrowIndex, uint256 reservesToAdd)
func (_LendingPool *LendingPoolSession) PreviewAccrual() (struct {
	NextSupplyIndex *big.Int
	NextBorrowIndex *big.Int
	ReservesToAdd   *big.Int
}, error) {
	return _LendingPool.Contract.PreviewAccrual(&_LendingPool.CallOpts)
}

// PreviewAccrual is a free data retrieval call binding the contract method 0x37d5722a.
//
// Solidity: function previewAccrual() view returns(uint256 nextSupplyIndex, uint256 nextBorrowIndex, uint256 reservesToAdd)
func (_LendingPool *LendingPoolCallerSession) PreviewAccrual() (struct {
	NextSupplyIndex *big.Int
	NextBorrowIndex *big.Int
	ReservesToAdd   *big.Int
}, error) {
	return _LendingPool.Contract.PreviewAccrual(&_LendingPool.CallOpts)
}

// RateModel is a free data retrieval call binding the contract method 0xa1088459.
//
// Solidity: function rateModel() view returns(address)
func (_LendingPool *LendingPoolCaller) RateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "rateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RateModel is a free data retrieval call binding the contract method 0xa1088459.
//
// Solidity: function rateModel() view returns(address)
func (_LendingPool *LendingPoolSession) RateModel() (common.Address, error) {
	return _LendingPool.Contract.RateModel(&_LendingPool.CallOpts)
}

// RateModel is a free data retrieval call binding the contract method 0xa1088459.
//
// Solidity: function rateModel() view returns(address)
func (_LendingPool *LendingPoolCallerSession) RateModel() (common.Address, error) {
	return _LendingPool.Contract.RateModel(&_LendingPool.CallOpts)
}

// ReserveFactorBps is a free data retrieval call binding the contract method 0xc87bed70.
//
// Solidity: function reserveFactorBps() view returns(uint16)
func (_LendingPool *LendingPoolCaller) ReserveFactorBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "reserveFactorBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ReserveFactorBps is a free data retrieval call binding the contract method 0xc87bed70.
//
// Solidity: function reserveFactorBps() view returns(uint16)
func (_LendingPool *LendingPoolSession) ReserveFactorBps() (uint16, error) {
	return _LendingPool.Contract.ReserveFactorBps(&_LendingPool.CallOpts)
}

// ReserveFactorBps is a free data retrieval call binding the contract method 0xc87bed70.
//
// Solidity: function reserveFactorBps() view returns(uint16)
func (_LendingPool *LendingPoolCallerSession) ReserveFactorBps() (uint16, error) {
	return _LendingPool.Contract.ReserveFactorBps(&_LendingPool.CallOpts)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address ) view returns(uint256)
func (_LendingPool *LendingPoolCaller) SharesOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "sharesOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address ) view returns(uint256)
func (_LendingPool *LendingPoolSession) SharesOf(arg0 common.Address) (*big.Int, error) {
	return _LendingPool.Contract.SharesOf(&_LendingPool.CallOpts, arg0)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address ) view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) SharesOf(arg0 common.Address) (*big.Int, error) {
	return _LendingPool.Contract.SharesOf(&_LendingPool.CallOpts, arg0)
}

// SupplyIndex is a free data retrieval call binding the contract method 0x98f1bc12.
//
// Solidity: function supplyIndex() view returns(uint256)
func (_LendingPool *LendingPoolCaller) SupplyIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "supplyIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyIndex is a free data retrieval call binding the contract method 0x98f1bc12.
//
// Solidity: function supplyIndex() view returns(uint256)
func (_LendingPool *LendingPoolSession) SupplyIndex() (*big.Int, error) {
	return _LendingPool.Contract.SupplyIndex(&_LendingPool.CallOpts)
}

// SupplyIndex is a free data retrieval call binding the contract method 0x98f1bc12.
//
// Solidity: function supplyIndex() view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) SupplyIndex() (*big.Int, error) {
	return _LendingPool.Contract.SupplyIndex(&_LendingPool.CallOpts)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_LendingPool *LendingPoolCaller) TotalBorrowed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "totalBorrowed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_LendingPool *LendingPoolSession) TotalBorrowed() (*big.Int, error) {
	return _LendingPool.Contract.TotalBorrowed(&_LendingPool.CallOpts)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) TotalBorrowed() (*big.Int, error) {
	return _LendingPool.Contract.TotalBorrowed(&_LendingPool.CallOpts)
}

// TotalDebtShares is a free data retrieval call binding the contract method 0x1b859e41.
//
// Solidity: function totalDebtShares() view returns(uint256)
func (_LendingPool *LendingPoolCaller) TotalDebtShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "totalDebtShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebtShares is a free data retrieval call binding the contract method 0x1b859e41.
//
// Solidity: function totalDebtShares() view returns(uint256)
func (_LendingPool *LendingPoolSession) TotalDebtShares() (*big.Int, error) {
	return _LendingPool.Contract.TotalDebtShares(&_LendingPool.CallOpts)
}

// TotalDebtShares is a free data retrieval call binding the contract method 0x1b859e41.
//
// Solidity: function totalDebtShares() view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) TotalDebtShares() (*big.Int, error) {
	return _LendingPool.Contract.TotalDebtShares(&_LendingPool.CallOpts)
}

// TotalSupplied is a free data retrieval call binding the contract method 0x630fd0ac.
//
// Solidity: function totalSupplied() view returns(uint256)
func (_LendingPool *LendingPoolCaller) TotalSupplied(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "totalSupplied")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplied is a free data retrieval call binding the contract method 0x630fd0ac.
//
// Solidity: function totalSupplied() view returns(uint256)
func (_LendingPool *LendingPoolSession) TotalSupplied() (*big.Int, error) {
	return _LendingPool.Contract.TotalSupplied(&_LendingPool.CallOpts)
}

// TotalSupplied is a free data retrieval call binding the contract method 0x630fd0ac.
//
// Solidity: function totalSupplied() view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) TotalSupplied() (*big.Int, error) {
	return _LendingPool.Contract.TotalSupplied(&_LendingPool.CallOpts)
}

// TotalSupplyShares is a free data retrieval call binding the contract method 0xba7bde55.
//
// Solidity: function totalSupplyShares() view returns(uint256)
func (_LendingPool *LendingPoolCaller) TotalSupplyShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "totalSupplyShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyShares is a free data retrieval call binding the contract method 0xba7bde55.
//
// Solidity: function totalSupplyShares() view returns(uint256)
func (_LendingPool *LendingPoolSession) TotalSupplyShares() (*big.Int, error) {
	return _LendingPool.Contract.TotalSupplyShares(&_LendingPool.CallOpts)
}

// TotalSupplyShares is a free data retrieval call binding the contract method 0xba7bde55.
//
// Solidity: function totalSupplyShares() view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) TotalSupplyShares() (*big.Int, error) {
	return _LendingPool.Contract.TotalSupplyShares(&_LendingPool.CallOpts)
}

// UtilizationBps is a free data retrieval call binding the contract method 0x975e900e.
//
// Solidity: function utilizationBps() view returns(uint256)
func (_LendingPool *LendingPoolCaller) UtilizationBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendingPool.contract.Call(opts, &out, "utilizationBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UtilizationBps is a free data retrieval call binding the contract method 0x975e900e.
//
// Solidity: function utilizationBps() view returns(uint256)
func (_LendingPool *LendingPoolSession) UtilizationBps() (*big.Int, error) {
	return _LendingPool.Contract.UtilizationBps(&_LendingPool.CallOpts)
}

// UtilizationBps is a free data retrieval call binding the contract method 0x975e900e.
//
// Solidity: function utilizationBps() view returns(uint256)
func (_LendingPool *LendingPoolCallerSession) UtilizationBps() (*big.Int, error) {
	return _LendingPool.Contract.UtilizationBps(&_LendingPool.CallOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns()
func (_LendingPool *LendingPoolTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns()
func (_LendingPool *LendingPoolSession) AccrueInterest() (*types.Transaction, error) {
	return _LendingPool.Contract.AccrueInterest(&_LendingPool.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns()
func (_LendingPool *LendingPoolTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _LendingPool.Contract.AccrueInterest(&_LendingPool.TransactOpts)
}

// BorrowFor is a paid mutator transaction binding the contract method 0x30880441.
//
// Solidity: function borrowFor(address borrower, address recipient, uint256 amount) returns()
func (_LendingPool *LendingPoolTransactor) BorrowFor(opts *bind.TransactOpts, borrower common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "borrowFor", borrower, recipient, amount)
}

// BorrowFor is a paid mutator transaction binding the contract method 0x30880441.
//
// Solidity: function borrowFor(address borrower, address recipient, uint256 amount) returns()
func (_LendingPool *LendingPoolSession) BorrowFor(borrower common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.BorrowFor(&_LendingPool.TransactOpts, borrower, recipient, amount)
}

// BorrowFor is a paid mutator transaction binding the contract method 0x30880441.
//
// Solidity: function borrowFor(address borrower, address recipient, uint256 amount) returns()
func (_LendingPool *LendingPoolTransactorSession) BorrowFor(borrower common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.BorrowFor(&_LendingPool.TransactOpts, borrower, recipient, amount)
}

// CollectAllReserves is a paid mutator transaction binding the contract method 0xcc863d10.
//
// Solidity: function collectAllReserves(address recipient) returns(uint256 collected)
func (_LendingPool *LendingPoolTransactor) CollectAllReserves(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "collectAllReserves", recipient)
}

// CollectAllReserves is a paid mutator transaction binding the contract method 0xcc863d10.
//
// Solidity: function collectAllReserves(address recipient) returns(uint256 collected)
func (_LendingPool *LendingPoolSession) CollectAllReserves(recipient common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.CollectAllReserves(&_LendingPool.TransactOpts, recipient)
}

// CollectAllReserves is a paid mutator transaction binding the contract method 0xcc863d10.
//
// Solidity: function collectAllReserves(address recipient) returns(uint256 collected)
func (_LendingPool *LendingPoolTransactorSession) CollectAllReserves(recipient common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.CollectAllReserves(&_LendingPool.TransactOpts, recipient)
}

// CollectReserves is a paid mutator transaction binding the contract method 0x390c4caf.
//
// Solidity: function collectReserves(address recipient, uint256 amount) returns(uint256 collected)
func (_LendingPool *LendingPoolTransactor) CollectReserves(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "collectReserves", recipient, amount)
}

// CollectReserves is a paid mutator transaction binding the contract method 0x390c4caf.
//
// Solidity: function collectReserves(address recipient, uint256 amount) returns(uint256 collected)
func (_LendingPool *LendingPoolSession) CollectReserves(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.CollectReserves(&_LendingPool.TransactOpts, recipient, amount)
}

// CollectReserves is a paid mutator transaction binding the contract method 0x390c4caf.
//
// Solidity: function collectReserves(address recipient, uint256 amount) returns(uint256 collected)
func (_LendingPool *LendingPoolTransactorSession) CollectReserves(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.CollectReserves(&_LendingPool.TransactOpts, recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 assets) returns(uint256 shares)
func (_LendingPool *LendingPoolTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "deposit", assets)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 assets) returns(uint256 shares)
func (_LendingPool *LendingPoolSession) Deposit(assets *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.Deposit(&_LendingPool.TransactOpts, assets)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 assets) returns(uint256 shares)
func (_LendingPool *LendingPoolTransactorSession) Deposit(assets *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.Deposit(&_LendingPool.TransactOpts, assets)
}

// LinkController is a paid mutator transaction binding the contract method 0xdf95ec90.
//
// Solidity: function linkController(address newController) returns()
func (_LendingPool *LendingPoolTransactor) LinkController(opts *bind.TransactOpts, newController common.Address) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "linkController", newController)
}

// LinkController is a paid mutator transaction binding the contract method 0xdf95ec90.
//
// Solidity: function linkController(address newController) returns()
func (_LendingPool *LendingPoolSession) LinkController(newController common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.LinkController(&_LendingPool.TransactOpts, newController)
}

// LinkController is a paid mutator transaction binding the contract method 0xdf95ec90.
//
// Solidity: function linkController(address newController) returns()
func (_LendingPool *LendingPoolTransactorSession) LinkController(newController common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.LinkController(&_LendingPool.TransactOpts, newController)
}

// LinkLiquidationManager is a paid mutator transaction binding the contract method 0xc894270c.
//
// Solidity: function linkLiquidationManager(address newLiquidationManager) returns()
func (_LendingPool *LendingPoolTransactor) LinkLiquidationManager(opts *bind.TransactOpts, newLiquidationManager common.Address) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "linkLiquidationManager", newLiquidationManager)
}

// LinkLiquidationManager is a paid mutator transaction binding the contract method 0xc894270c.
//
// Solidity: function linkLiquidationManager(address newLiquidationManager) returns()
func (_LendingPool *LendingPoolSession) LinkLiquidationManager(newLiquidationManager common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.LinkLiquidationManager(&_LendingPool.TransactOpts, newLiquidationManager)
}

// LinkLiquidationManager is a paid mutator transaction binding the contract method 0xc894270c.
//
// Solidity: function linkLiquidationManager(address newLiquidationManager) returns()
func (_LendingPool *LendingPoolTransactorSession) LinkLiquidationManager(newLiquidationManager common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.LinkLiquidationManager(&_LendingPool.TransactOpts, newLiquidationManager)
}

// RedeemShares is a paid mutator transaction binding the contract method 0xadba9804.
//
// Solidity: function redeemShares(uint256 shares) returns(uint256 assets)
func (_LendingPool *LendingPoolTransactor) RedeemShares(opts *bind.TransactOpts, shares *big.Int) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "redeemShares", shares)
}

// RedeemShares is a paid mutator transaction binding the contract method 0xadba9804.
//
// Solidity: function redeemShares(uint256 shares) returns(uint256 assets)
func (_LendingPool *LendingPoolSession) RedeemShares(shares *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.RedeemShares(&_LendingPool.TransactOpts, shares)
}

// RedeemShares is a paid mutator transaction binding the contract method 0xadba9804.
//
// Solidity: function redeemShares(uint256 shares) returns(uint256 assets)
func (_LendingPool *LendingPoolTransactorSession) RedeemShares(shares *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.RedeemShares(&_LendingPool.TransactOpts, shares)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingPool *LendingPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingPool *LendingPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _LendingPool.Contract.RenounceOwnership(&_LendingPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingPool *LendingPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LendingPool.Contract.RenounceOwnership(&_LendingPool.TransactOpts)
}

// RepayAllFor is a paid mutator transaction binding the contract method 0x9254702e.
//
// Solidity: function repayAllFor(address borrower, address payer) returns(uint256 amountPaid)
func (_LendingPool *LendingPoolTransactor) RepayAllFor(opts *bind.TransactOpts, borrower common.Address, payer common.Address) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "repayAllFor", borrower, payer)
}

// RepayAllFor is a paid mutator transaction binding the contract method 0x9254702e.
//
// Solidity: function repayAllFor(address borrower, address payer) returns(uint256 amountPaid)
func (_LendingPool *LendingPoolSession) RepayAllFor(borrower common.Address, payer common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.RepayAllFor(&_LendingPool.TransactOpts, borrower, payer)
}

// RepayAllFor is a paid mutator transaction binding the contract method 0x9254702e.
//
// Solidity: function repayAllFor(address borrower, address payer) returns(uint256 amountPaid)
func (_LendingPool *LendingPoolTransactorSession) RepayAllFor(borrower common.Address, payer common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.RepayAllFor(&_LendingPool.TransactOpts, borrower, payer)
}

// RepayFor is a paid mutator transaction binding the contract method 0x976ce495.
//
// Solidity: function repayFor(address borrower, address payer, uint256 amount) returns()
func (_LendingPool *LendingPoolTransactor) RepayFor(opts *bind.TransactOpts, borrower common.Address, payer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "repayFor", borrower, payer, amount)
}

// RepayFor is a paid mutator transaction binding the contract method 0x976ce495.
//
// Solidity: function repayFor(address borrower, address payer, uint256 amount) returns()
func (_LendingPool *LendingPoolSession) RepayFor(borrower common.Address, payer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.RepayFor(&_LendingPool.TransactOpts, borrower, payer, amount)
}

// RepayFor is a paid mutator transaction binding the contract method 0x976ce495.
//
// Solidity: function repayFor(address borrower, address payer, uint256 amount) returns()
func (_LendingPool *LendingPoolTransactorSession) RepayFor(borrower common.Address, payer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.RepayFor(&_LendingPool.TransactOpts, borrower, payer, amount)
}

// SetDepositsPaused is a paid mutator transaction binding the contract method 0x1c481e2d.
//
// Solidity: function setDepositsPaused(bool isPaused) returns()
func (_LendingPool *LendingPoolTransactor) SetDepositsPaused(opts *bind.TransactOpts, isPaused bool) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "setDepositsPaused", isPaused)
}

// SetDepositsPaused is a paid mutator transaction binding the contract method 0x1c481e2d.
//
// Solidity: function setDepositsPaused(bool isPaused) returns()
func (_LendingPool *LendingPoolSession) SetDepositsPaused(isPaused bool) (*types.Transaction, error) {
	return _LendingPool.Contract.SetDepositsPaused(&_LendingPool.TransactOpts, isPaused)
}

// SetDepositsPaused is a paid mutator transaction binding the contract method 0x1c481e2d.
//
// Solidity: function setDepositsPaused(bool isPaused) returns()
func (_LendingPool *LendingPoolTransactorSession) SetDepositsPaused(isPaused bool) (*types.Transaction, error) {
	return _LendingPool.Contract.SetDepositsPaused(&_LendingPool.TransactOpts, isPaused)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 newAmount) returns()
func (_LendingPool *LendingPoolTransactor) SetMinDeposit(opts *bind.TransactOpts, newAmount *big.Int) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "setMinDeposit", newAmount)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 newAmount) returns()
func (_LendingPool *LendingPoolSession) SetMinDeposit(newAmount *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.SetMinDeposit(&_LendingPool.TransactOpts, newAmount)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 newAmount) returns()
func (_LendingPool *LendingPoolTransactorSession) SetMinDeposit(newAmount *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.SetMinDeposit(&_LendingPool.TransactOpts, newAmount)
}

// SetRateModel is a paid mutator transaction binding the contract method 0x7f9028c8.
//
// Solidity: function setRateModel(address newModel) returns()
func (_LendingPool *LendingPoolTransactor) SetRateModel(opts *bind.TransactOpts, newModel common.Address) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "setRateModel", newModel)
}

// SetRateModel is a paid mutator transaction binding the contract method 0x7f9028c8.
//
// Solidity: function setRateModel(address newModel) returns()
func (_LendingPool *LendingPoolSession) SetRateModel(newModel common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.SetRateModel(&_LendingPool.TransactOpts, newModel)
}

// SetRateModel is a paid mutator transaction binding the contract method 0x7f9028c8.
//
// Solidity: function setRateModel(address newModel) returns()
func (_LendingPool *LendingPoolTransactorSession) SetRateModel(newModel common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.SetRateModel(&_LendingPool.TransactOpts, newModel)
}

// SetReserveFactorBps is a paid mutator transaction binding the contract method 0x570dbba5.
//
// Solidity: function setReserveFactorBps(uint16 newBps) returns()
func (_LendingPool *LendingPoolTransactor) SetReserveFactorBps(opts *bind.TransactOpts, newBps uint16) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "setReserveFactorBps", newBps)
}

// SetReserveFactorBps is a paid mutator transaction binding the contract method 0x570dbba5.
//
// Solidity: function setReserveFactorBps(uint16 newBps) returns()
func (_LendingPool *LendingPoolSession) SetReserveFactorBps(newBps uint16) (*types.Transaction, error) {
	return _LendingPool.Contract.SetReserveFactorBps(&_LendingPool.TransactOpts, newBps)
}

// SetReserveFactorBps is a paid mutator transaction binding the contract method 0x570dbba5.
//
// Solidity: function setReserveFactorBps(uint16 newBps) returns()
func (_LendingPool *LendingPoolTransactorSession) SetReserveFactorBps(newBps uint16) (*types.Transaction, error) {
	return _LendingPool.Contract.SetReserveFactorBps(&_LendingPool.TransactOpts, newBps)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingPool *LendingPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingPool *LendingPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.TransferOwnership(&_LendingPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingPool *LendingPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LendingPool.Contract.TransferOwnership(&_LendingPool.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 assets) returns(uint256 shares)
func (_LendingPool *LendingPoolTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int) (*types.Transaction, error) {
	return _LendingPool.contract.Transact(opts, "withdraw", assets)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 assets) returns(uint256 shares)
func (_LendingPool *LendingPoolSession) Withdraw(assets *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.Withdraw(&_LendingPool.TransactOpts, assets)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 assets) returns(uint256 shares)
func (_LendingPool *LendingPoolTransactorSession) Withdraw(assets *big.Int) (*types.Transaction, error) {
	return _LendingPool.Contract.Withdraw(&_LendingPool.TransactOpts, assets)
}

// LendingPoolControllerLinkedIterator is returned from FilterControllerLinked and is used to iterate over the raw logs and unpacked data for ControllerLinked events raised by the LendingPool contract.
type LendingPoolControllerLinkedIterator struct {
	Event *LendingPoolControllerLinked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LendingPoolControllerLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolControllerLinked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LendingPoolControllerLinked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LendingPoolControllerLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolControllerLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolControllerLinked represents a ControllerLinked event raised by the LendingPool contract.
type LendingPoolControllerLinked struct {
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterControllerLinked is a free log retrieval operation binding the contract event 0xa76ae82e550122379c3a3784cfc195ef6f1bae7f61abf8fae42a47725830c6c8.
//
// Solidity: event ControllerLinked(address controller)
func (_LendingPool *LendingPoolFilterer) FilterControllerLinked(opts *bind.FilterOpts) (*LendingPoolControllerLinkedIterator, error) {

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "ControllerLinked")
	if err != nil {
		return nil, err
	}
	return &LendingPoolControllerLinkedIterator{contract: _LendingPool.contract, event: "ControllerLinked", logs: logs, sub: sub}, nil
}

// WatchControllerLinked is a free log subscription operation binding the contract event 0xa76ae82e550122379c3a3784cfc195ef6f1bae7f61abf8fae42a47725830c6c8.
//
// Solidity: event ControllerLinked(address controller)
func (_LendingPool *LendingPoolFilterer) WatchControllerLinked(opts *bind.WatchOpts, sink chan<- *LendingPoolControllerLinked) (event.Subscription, error) {

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "ControllerLinked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolControllerLinked)
				if err := _LendingPool.contract.UnpackLog(event, "ControllerLinked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseControllerLinked is a log parse operation binding the contract event 0xa76ae82e550122379c3a3784cfc195ef6f1bae7f61abf8fae42a47725830c6c8.
//
// Solidity: event ControllerLinked(address controller)
func (_LendingPool *LendingPoolFilterer) ParseControllerLinked(log types.Log) (*LendingPoolControllerLinked, error) {
	event := new(LendingPoolControllerLinked)
	if err := _LendingPool.contract.UnpackLog(event, "ControllerLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the LendingPool contract.
type LendingPoolDepositIterator struct {
	Event *LendingPoolDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LendingPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LendingPoolDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LendingPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolDeposit represents a Deposit event raised by the LendingPool contract.
type LendingPoolDeposit struct {
	Lender        common.Address
	Assets        *big.Int
	Shares        *big.Int
	SupplyIndex   *big.Int
	TotalSupplied *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7162984403f6c73c8639375d45a9187dfd04602231bd8e587c415718b5f7e5f9.
//
// Solidity: event Deposit(address indexed lender, uint256 assets, uint256 shares, uint256 supplyIndex, uint256 totalSupplied)
func (_LendingPool *LendingPoolFilterer) FilterDeposit(opts *bind.FilterOpts, lender []common.Address) (*LendingPoolDepositIterator, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "Deposit", lenderRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolDepositIterator{contract: _LendingPool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7162984403f6c73c8639375d45a9187dfd04602231bd8e587c415718b5f7e5f9.
//
// Solidity: event Deposit(address indexed lender, uint256 assets, uint256 shares, uint256 supplyIndex, uint256 totalSupplied)
func (_LendingPool *LendingPoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *LendingPoolDeposit, lender []common.Address) (event.Subscription, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "Deposit", lenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolDeposit)
				if err := _LendingPool.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x7162984403f6c73c8639375d45a9187dfd04602231bd8e587c415718b5f7e5f9.
//
// Solidity: event Deposit(address indexed lender, uint256 assets, uint256 shares, uint256 supplyIndex, uint256 totalSupplied)
func (_LendingPool *LendingPoolFilterer) ParseDeposit(log types.Log) (*LendingPoolDeposit, error) {
	event := new(LendingPoolDeposit)
	if err := _LendingPool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolDepositsPausedChangedIterator is returned from FilterDepositsPausedChanged and is used to iterate over the raw logs and unpacked data for DepositsPausedChanged events raised by the LendingPool contract.
type LendingPoolDepositsPausedChangedIterator struct {
	Event *LendingPoolDepositsPausedChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LendingPoolDepositsPausedChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolDepositsPausedChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LendingPoolDepositsPausedChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LendingPoolDepositsPausedChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolDepositsPausedChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolDepositsPausedChanged represents a DepositsPausedChanged event raised by the LendingPool contract.
type LendingPoolDepositsPausedChanged struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositsPausedChanged is a free log retrieval operation binding the contract event 0xc4eef1ae414c5813c621be366ea9462b94e6e0f3e9bad6c93bdec72e1d1b4b9c.
//
// Solidity: event DepositsPausedChanged(bool isPaused)
func (_LendingPool *LendingPoolFilterer) FilterDepositsPausedChanged(opts *bind.FilterOpts) (*LendingPoolDepositsPausedChangedIterator, error) {

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "DepositsPausedChanged")
	if err != nil {
		return nil, err
	}
	return &LendingPoolDepositsPausedChangedIterator{contract: _LendingPool.contract, event: "DepositsPausedChanged", logs: logs, sub: sub}, nil
}

// WatchDepositsPausedChanged is a free log subscription operation binding the contract event 0xc4eef1ae414c5813c621be366ea9462b94e6e0f3e9bad6c93bdec72e1d1b4b9c.
//
// Solidity: event DepositsPausedChanged(bool isPaused)
func (_LendingPool *LendingPoolFilterer) WatchDepositsPausedChanged(opts *bind.WatchOpts, sink chan<- *LendingPoolDepositsPausedChanged) (event.Subscription, error) {

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "DepositsPausedChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolDepositsPausedChanged)
				if err := _LendingPool.contract.UnpackLog(event, "DepositsPausedChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositsPausedChanged is a log parse operation binding the contract event 0xc4eef1ae414c5813c621be366ea9462b94e6e0f3e9bad6c93bdec72e1d1b4b9c.
//
// Solidity: event DepositsPausedChanged(bool isPaused)
func (_LendingPool *LendingPoolFilterer) ParseDepositsPausedChanged(log types.Log) (*LendingPoolDepositsPausedChanged, error) {
	event := new(LendingPoolDepositsPausedChanged)
	if err := _LendingPool.contract.UnpackLog(event, "DepositsPausedChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolInterestAccruedIterator is returned from FilterInterestAccrued and is used to iterate over the raw logs and unpacked data for InterestAccrued events raised by the LendingPool contract.
type LendingPoolInterestAccruedIterator struct {
	Event *LendingPoolInterestAccrued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LendingPoolInterestAccruedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolInterestAccrued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LendingPoolInterestAccrued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LendingPoolInterestAccruedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolInterestAccruedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolInterestAccrued represents a InterestAccrued event raised by the LendingPool contract.
type LendingPoolInterestAccrued struct {
	SupplyIndex     *big.Int
	BorrowIndex     *big.Int
	TotalBorrowed   *big.Int
	ReservesAccrued *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterInterestAccrued is a free log retrieval operation binding the contract event 0xe3972ae250606e82c6683843383cb8a51755588d2f443af42e8cb5b1e085392d.
//
// Solidity: event InterestAccrued(uint256 supplyIndex, uint256 borrowIndex, uint256 totalBorrowed, uint256 reservesAccrued)
func (_LendingPool *LendingPoolFilterer) FilterInterestAccrued(opts *bind.FilterOpts) (*LendingPoolInterestAccruedIterator, error) {

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "InterestAccrued")
	if err != nil {
		return nil, err
	}
	return &LendingPoolInterestAccruedIterator{contract: _LendingPool.contract, event: "InterestAccrued", logs: logs, sub: sub}, nil
}

// WatchInterestAccrued is a free log subscription operation binding the contract event 0xe3972ae250606e82c6683843383cb8a51755588d2f443af42e8cb5b1e085392d.
//
// Solidity: event InterestAccrued(uint256 supplyIndex, uint256 borrowIndex, uint256 totalBorrowed, uint256 reservesAccrued)
func (_LendingPool *LendingPoolFilterer) WatchInterestAccrued(opts *bind.WatchOpts, sink chan<- *LendingPoolInterestAccrued) (event.Subscription, error) {

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "InterestAccrued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolInterestAccrued)
				if err := _LendingPool.contract.UnpackLog(event, "InterestAccrued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInterestAccrued is a log parse operation binding the contract event 0xe3972ae250606e82c6683843383cb8a51755588d2f443af42e8cb5b1e085392d.
//
// Solidity: event InterestAccrued(uint256 supplyIndex, uint256 borrowIndex, uint256 totalBorrowed, uint256 reservesAccrued)
func (_LendingPool *LendingPoolFilterer) ParseInterestAccrued(log types.Log) (*LendingPoolInterestAccrued, error) {
	event := new(LendingPoolInterestAccrued)
	if err := _LendingPool.contract.UnpackLog(event, "InterestAccrued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolLiquidationManagerLinkedIterator is returned from FilterLiquidationManagerLinked and is used to iterate over the raw logs and unpacked data for LiquidationManagerLinked events raised by the LendingPool contract.
type LendingPoolLiquidationManagerLinkedIterator struct {
	Event *LendingPoolLiquidationManagerLinked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LendingPoolLiquidationManagerLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolLiquidationManagerLinked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LendingPoolLiquidationManagerLinked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LendingPoolLiquidationManagerLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolLiquidationManagerLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolLiquidationManagerLinked represents a LiquidationManagerLinked event raised by the LendingPool contract.
type LendingPoolLiquidationManagerLinked struct {
	LiquidationManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLiquidationManagerLinked is a free log retrieval operation binding the contract event 0x40fd76f04a4f105bfd0d978d7bc92e40784482098dc21077e350a7c121dfba1d.
//
// Solidity: event LiquidationManagerLinked(address liquidationManager)
func (_LendingPool *LendingPoolFilterer) FilterLiquidationManagerLinked(opts *bind.FilterOpts) (*LendingPoolLiquidationManagerLinkedIterator, error) {

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "LiquidationManagerLinked")
	if err != nil {
		return nil, err
	}
	return &LendingPoolLiquidationManagerLinkedIterator{contract: _LendingPool.contract, event: "LiquidationManagerLinked", logs: logs, sub: sub}, nil
}

// WatchLiquidationManagerLinked is a free log subscription operation binding the contract event 0x40fd76f04a4f105bfd0d978d7bc92e40784482098dc21077e350a7c121dfba1d.
//
// Solidity: event LiquidationManagerLinked(address liquidationManager)
func (_LendingPool *LendingPoolFilterer) WatchLiquidationManagerLinked(opts *bind.WatchOpts, sink chan<- *LendingPoolLiquidationManagerLinked) (event.Subscription, error) {

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "LiquidationManagerLinked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolLiquidationManagerLinked)
				if err := _LendingPool.contract.UnpackLog(event, "LiquidationManagerLinked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidationManagerLinked is a log parse operation binding the contract event 0x40fd76f04a4f105bfd0d978d7bc92e40784482098dc21077e350a7c121dfba1d.
//
// Solidity: event LiquidationManagerLinked(address liquidationManager)
func (_LendingPool *LendingPoolFilterer) ParseLiquidationManagerLinked(log types.Log) (*LendingPoolLiquidationManagerLinked, error) {
	event := new(LendingPoolLiquidationManagerLinked)
	if err := _LendingPool.contract.UnpackLog(event, "LiquidationManagerLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolMinDepositChangedIterator is returned from FilterMinDepositChanged and is used to iterate over the raw logs and unpacked data for MinDepositChanged events raised by the LendingPool contract.
type LendingPoolMinDepositChangedIterator struct {
	Event *LendingPoolMinDepositChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LendingPoolMinDepositChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolMinDepositChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LendingPoolMinDepositChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LendingPoolMinDepositChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolMinDepositChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolMinDepositChanged represents a MinDepositChanged event raised by the LendingPool contract.
type LendingPoolMinDepositChanged struct {
	PreviousAmount *big.Int
	NewAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMinDepositChanged is a free log retrieval operation binding the contract event 0xcacd94bd1e7bb1185c816a740d9439bc8eff8159f6f4ffad8d306b5aca2ebd92.
//
// Solidity: event MinDepositChanged(uint256 previousAmount, uint256 newAmount)
func (_LendingPool *LendingPoolFilterer) FilterMinDepositChanged(opts *bind.FilterOpts) (*LendingPoolMinDepositChangedIterator, error) {

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "MinDepositChanged")
	if err != nil {
		return nil, err
	}
	return &LendingPoolMinDepositChangedIterator{contract: _LendingPool.contract, event: "MinDepositChanged", logs: logs, sub: sub}, nil
}

// WatchMinDepositChanged is a free log subscription operation binding the contract event 0xcacd94bd1e7bb1185c816a740d9439bc8eff8159f6f4ffad8d306b5aca2ebd92.
//
// Solidity: event MinDepositChanged(uint256 previousAmount, uint256 newAmount)
func (_LendingPool *LendingPoolFilterer) WatchMinDepositChanged(opts *bind.WatchOpts, sink chan<- *LendingPoolMinDepositChanged) (event.Subscription, error) {

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "MinDepositChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolMinDepositChanged)
				if err := _LendingPool.contract.UnpackLog(event, "MinDepositChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinDepositChanged is a log parse operation binding the contract event 0xcacd94bd1e7bb1185c816a740d9439bc8eff8159f6f4ffad8d306b5aca2ebd92.
//
// Solidity: event MinDepositChanged(uint256 previousAmount, uint256 newAmount)
func (_LendingPool *LendingPoolFilterer) ParseMinDepositChanged(log types.Log) (*LendingPoolMinDepositChanged, error) {
	event := new(LendingPoolMinDepositChanged)
	if err := _LendingPool.contract.UnpackLog(event, "MinDepositChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LendingPool contract.
type LendingPoolOwnershipTransferredIterator struct {
	Event *LendingPoolOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LendingPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LendingPoolOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LendingPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolOwnershipTransferred represents a OwnershipTransferred event raised by the LendingPool contract.
type LendingPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingPool *LendingPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LendingPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolOwnershipTransferredIterator{contract: _LendingPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingPool *LendingPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LendingPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolOwnershipTransferred)
				if err := _LendingPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingPool *LendingPoolFilterer) ParseOwnershipTransferred(log types.Log) (*LendingPoolOwnershipTransferred, error) {
	event := new(LendingPoolOwnershipTransferred)
	if err := _LendingPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolRateModelChangedIterator is returned from FilterRateModelChanged and is used to iterate over the raw logs and unpacked data for RateModelChanged events raised by the LendingPool contract.
type LendingPoolRateModelChangedIterator struct {
	Event *LendingPoolRateModelChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LendingPoolRateModelChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolRateModelChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LendingPoolRateModelChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LendingPoolRateModelChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolRateModelChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolRateModelChanged represents a RateModelChanged event raised by the LendingPool contract.
type LendingPoolRateModelChanged struct {
	PreviousModel common.Address
	NewModel      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRateModelChanged is a free log retrieval operation binding the contract event 0x17f03d48d6a1d89526a5f262a5054aee66a5f171d7c97d08ea786a4b9a04bef8.
//
// Solidity: event RateModelChanged(address previousModel, address newModel)
func (_LendingPool *LendingPoolFilterer) FilterRateModelChanged(opts *bind.FilterOpts) (*LendingPoolRateModelChangedIterator, error) {

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "RateModelChanged")
	if err != nil {
		return nil, err
	}
	return &LendingPoolRateModelChangedIterator{contract: _LendingPool.contract, event: "RateModelChanged", logs: logs, sub: sub}, nil
}

// WatchRateModelChanged is a free log subscription operation binding the contract event 0x17f03d48d6a1d89526a5f262a5054aee66a5f171d7c97d08ea786a4b9a04bef8.
//
// Solidity: event RateModelChanged(address previousModel, address newModel)
func (_LendingPool *LendingPoolFilterer) WatchRateModelChanged(opts *bind.WatchOpts, sink chan<- *LendingPoolRateModelChanged) (event.Subscription, error) {

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "RateModelChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolRateModelChanged)
				if err := _LendingPool.contract.UnpackLog(event, "RateModelChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRateModelChanged is a log parse operation binding the contract event 0x17f03d48d6a1d89526a5f262a5054aee66a5f171d7c97d08ea786a4b9a04bef8.
//
// Solidity: event RateModelChanged(address previousModel, address newModel)
func (_LendingPool *LendingPoolFilterer) ParseRateModelChanged(log types.Log) (*LendingPoolRateModelChanged, error) {
	event := new(LendingPoolRateModelChanged)
	if err := _LendingPool.contract.UnpackLog(event, "RateModelChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolReserveFactorChangedIterator is returned from FilterReserveFactorChanged and is used to iterate over the raw logs and unpacked data for ReserveFactorChanged events raised by the LendingPool contract.
type LendingPoolReserveFactorChangedIterator struct {
	Event *LendingPoolReserveFactorChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LendingPoolReserveFactorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolReserveFactorChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LendingPoolReserveFactorChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LendingPoolReserveFactorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolReserveFactorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolReserveFactorChanged represents a ReserveFactorChanged event raised by the LendingPool contract.
type LendingPoolReserveFactorChanged struct {
	PreviousBps uint16
	NewBps      uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReserveFactorChanged is a free log retrieval operation binding the contract event 0x4207dd431b604e1a2ba057ca6243d2ad619677ea8f5c46e6ffdf90491d4f2078.
//
// Solidity: event ReserveFactorChanged(uint16 previousBps, uint16 newBps)
func (_LendingPool *LendingPoolFilterer) FilterReserveFactorChanged(opts *bind.FilterOpts) (*LendingPoolReserveFactorChangedIterator, error) {

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "ReserveFactorChanged")
	if err != nil {
		return nil, err
	}
	return &LendingPoolReserveFactorChangedIterator{contract: _LendingPool.contract, event: "ReserveFactorChanged", logs: logs, sub: sub}, nil
}

// WatchReserveFactorChanged is a free log subscription operation binding the contract event 0x4207dd431b604e1a2ba057ca6243d2ad619677ea8f5c46e6ffdf90491d4f2078.
//
// Solidity: event ReserveFactorChanged(uint16 previousBps, uint16 newBps)
func (_LendingPool *LendingPoolFilterer) WatchReserveFactorChanged(opts *bind.WatchOpts, sink chan<- *LendingPoolReserveFactorChanged) (event.Subscription, error) {

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "ReserveFactorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolReserveFactorChanged)
				if err := _LendingPool.contract.UnpackLog(event, "ReserveFactorChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReserveFactorChanged is a log parse operation binding the contract event 0x4207dd431b604e1a2ba057ca6243d2ad619677ea8f5c46e6ffdf90491d4f2078.
//
// Solidity: event ReserveFactorChanged(uint16 previousBps, uint16 newBps)
func (_LendingPool *LendingPoolFilterer) ParseReserveFactorChanged(log types.Log) (*LendingPoolReserveFactorChanged, error) {
	event := new(LendingPoolReserveFactorChanged)
	if err := _LendingPool.contract.UnpackLog(event, "ReserveFactorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolReservesCollectedIterator is returned from FilterReservesCollected and is used to iterate over the raw logs and unpacked data for ReservesCollected events raised by the LendingPool contract.
type LendingPoolReservesCollectedIterator struct {
	Event *LendingPoolReservesCollected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LendingPoolReservesCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolReservesCollected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LendingPoolReservesCollected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LendingPoolReservesCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolReservesCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolReservesCollected represents a ReservesCollected event raised by the LendingPool contract.
type LendingPoolReservesCollected struct {
	Recipient         common.Address
	Amount            *big.Int
	RemainingReserves *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterReservesCollected is a free log retrieval operation binding the contract event 0x059757002dee698353522edc5506fabf6296b0d7b46e966c42f204cb1faa0a02.
//
// Solidity: event ReservesCollected(address indexed recipient, uint256 amount, uint256 remainingReserves)
func (_LendingPool *LendingPoolFilterer) FilterReservesCollected(opts *bind.FilterOpts, recipient []common.Address) (*LendingPoolReservesCollectedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "ReservesCollected", recipientRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolReservesCollectedIterator{contract: _LendingPool.contract, event: "ReservesCollected", logs: logs, sub: sub}, nil
}

// WatchReservesCollected is a free log subscription operation binding the contract event 0x059757002dee698353522edc5506fabf6296b0d7b46e966c42f204cb1faa0a02.
//
// Solidity: event ReservesCollected(address indexed recipient, uint256 amount, uint256 remainingReserves)
func (_LendingPool *LendingPoolFilterer) WatchReservesCollected(opts *bind.WatchOpts, sink chan<- *LendingPoolReservesCollected, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "ReservesCollected", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolReservesCollected)
				if err := _LendingPool.contract.UnpackLog(event, "ReservesCollected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReservesCollected is a log parse operation binding the contract event 0x059757002dee698353522edc5506fabf6296b0d7b46e966c42f204cb1faa0a02.
//
// Solidity: event ReservesCollected(address indexed recipient, uint256 amount, uint256 remainingReserves)
func (_LendingPool *LendingPoolFilterer) ParseReservesCollected(log types.Log) (*LendingPoolReservesCollected, error) {
	event := new(LendingPoolReservesCollected)
	if err := _LendingPool.contract.UnpackLog(event, "ReservesCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingPoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the LendingPool contract.
type LendingPoolWithdrawIterator struct {
	Event *LendingPoolWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LendingPoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LendingPoolWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LendingPoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolWithdraw represents a Withdraw event raised by the LendingPool contract.
type LendingPoolWithdraw struct {
	Lender        common.Address
	Assets        *big.Int
	Shares        *big.Int
	SupplyIndex   *big.Int
	TotalSupplied *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xe08737ac48a1dab4b1a46c7dc9398bd5bfc6d7ad6fabb7cd8caa254de14def35.
//
// Solidity: event Withdraw(address indexed lender, uint256 assets, uint256 shares, uint256 supplyIndex, uint256 totalSupplied)
func (_LendingPool *LendingPoolFilterer) FilterWithdraw(opts *bind.FilterOpts, lender []common.Address) (*LendingPoolWithdrawIterator, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _LendingPool.contract.FilterLogs(opts, "Withdraw", lenderRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolWithdrawIterator{contract: _LendingPool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xe08737ac48a1dab4b1a46c7dc9398bd5bfc6d7ad6fabb7cd8caa254de14def35.
//
// Solidity: event Withdraw(address indexed lender, uint256 assets, uint256 shares, uint256 supplyIndex, uint256 totalSupplied)
func (_LendingPool *LendingPoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *LendingPoolWithdraw, lender []common.Address) (event.Subscription, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _LendingPool.contract.WatchLogs(opts, "Withdraw", lenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolWithdraw)
				if err := _LendingPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0xe08737ac48a1dab4b1a46c7dc9398bd5bfc6d7ad6fabb7cd8caa254de14def35.
//
// Solidity: event Withdraw(address indexed lender, uint256 assets, uint256 shares, uint256 supplyIndex, uint256 totalSupplied)
func (_LendingPool *LendingPoolFilterer) ParseWithdraw(log types.Log) (*LendingPoolWithdraw, error) {
	event := new(LendingPoolWithdraw)
	if err := _LendingPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
