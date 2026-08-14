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

// LendingControllerMetaData contains all meta data concerning the LendingController contract.
var LendingControllerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lendingPool\",\"type\":\"address\",\"internalType\":\"contractILendingPool\"},{\"name\":\"collateralVault\",\"type\":\"address\",\"internalType\":\"contractICollateralVault\"},{\"name\":\"priceOracle\",\"type\":\"address\",\"internalType\":\"contractIPriceOracle\"},{\"name\":\"startingMaxLtvBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"startingLiquidationThresholdBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"startingLiquidationBonusBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"borrow\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"borrowPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralAsset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralValueOf\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"debtAsset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"debtDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"debtOf\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"debtValueOf\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"healthFactorBps\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLiquidatable\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidationBonusBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidationThresholdBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxBorrowable\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxLtvBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxWithdrawableCollateral\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPriceOracle\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILendingPool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"repay\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"repayAll\",\"inputs\":[],\"outputs\":[{\"name\":\"amountPaid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBorrowPaused\",\"inputs\":[{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRiskSettings\",\"inputs\":[{\"name\":\"newMaxLtvBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"newThresholdBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"newBonusBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICollateralVault\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Borrow\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newDebt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"healthFactorBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"borrowIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BorrowPausedChanged\",\"inputs\":[{\"name\":\"isPaused\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Repay\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newDebt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RiskSettingsChanged\",\"inputs\":[{\"name\":\"maxLtvBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"liquidationThresholdBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"liquidationBonusBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ExceedsBorrowLimit\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"allowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ExceedsDebt\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRiskSettings\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughLiquidity\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]}]",
}

// LendingControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use LendingControllerMetaData.ABI instead.
var LendingControllerABI = LendingControllerMetaData.ABI

// LendingController is an auto generated Go binding around an Ethereum contract.
type LendingController struct {
	LendingControllerCaller     // Read-only binding to the contract
	LendingControllerTransactor // Write-only binding to the contract
	LendingControllerFilterer   // Log filterer for contract events
}

// LendingControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingControllerSession struct {
	Contract     *LendingController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LendingControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingControllerCallerSession struct {
	Contract *LendingControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LendingControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingControllerTransactorSession struct {
	Contract     *LendingControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LendingControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingControllerRaw struct {
	Contract *LendingController // Generic contract binding to access the raw methods on
}

// LendingControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingControllerCallerRaw struct {
	Contract *LendingControllerCaller // Generic read-only contract binding to access the raw methods on
}

// LendingControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingControllerTransactorRaw struct {
	Contract *LendingControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLendingController creates a new instance of LendingController, bound to a specific deployed contract.
func NewLendingController(address common.Address, backend bind.ContractBackend) (*LendingController, error) {
	contract, err := bindLendingController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LendingController{LendingControllerCaller: LendingControllerCaller{contract: contract}, LendingControllerTransactor: LendingControllerTransactor{contract: contract}, LendingControllerFilterer: LendingControllerFilterer{contract: contract}}, nil
}

// NewLendingControllerCaller creates a new read-only instance of LendingController, bound to a specific deployed contract.
func NewLendingControllerCaller(address common.Address, caller bind.ContractCaller) (*LendingControllerCaller, error) {
	contract, err := bindLendingController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingControllerCaller{contract: contract}, nil
}

// NewLendingControllerTransactor creates a new write-only instance of LendingController, bound to a specific deployed contract.
func NewLendingControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingControllerTransactor, error) {
	contract, err := bindLendingController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingControllerTransactor{contract: contract}, nil
}

// NewLendingControllerFilterer creates a new log filterer instance of LendingController, bound to a specific deployed contract.
func NewLendingControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingControllerFilterer, error) {
	contract, err := bindLendingController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingControllerFilterer{contract: contract}, nil
}

// bindLendingController binds a generic wrapper to an already deployed contract.
func bindLendingController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LendingControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingController *LendingControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingController.Contract.LendingControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingController *LendingControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingController.Contract.LendingControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingController *LendingControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingController.Contract.LendingControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingController *LendingControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingController *LendingControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingController *LendingControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingController.Contract.contract.Transact(opts, method, params...)
}

// BorrowPaused is a free data retrieval call binding the contract method 0xbcb4bbea.
//
// Solidity: function borrowPaused() view returns(bool)
func (_LendingController *LendingControllerCaller) BorrowPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "borrowPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowPaused is a free data retrieval call binding the contract method 0xbcb4bbea.
//
// Solidity: function borrowPaused() view returns(bool)
func (_LendingController *LendingControllerSession) BorrowPaused() (bool, error) {
	return _LendingController.Contract.BorrowPaused(&_LendingController.CallOpts)
}

// BorrowPaused is a free data retrieval call binding the contract method 0xbcb4bbea.
//
// Solidity: function borrowPaused() view returns(bool)
func (_LendingController *LendingControllerCallerSession) BorrowPaused() (bool, error) {
	return _LendingController.Contract.BorrowPaused(&_LendingController.CallOpts)
}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_LendingController *LendingControllerCaller) CollateralAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "collateralAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_LendingController *LendingControllerSession) CollateralAsset() (common.Address, error) {
	return _LendingController.Contract.CollateralAsset(&_LendingController.CallOpts)
}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_LendingController *LendingControllerCallerSession) CollateralAsset() (common.Address, error) {
	return _LendingController.Contract.CollateralAsset(&_LendingController.CallOpts)
}

// CollateralDecimals is a free data retrieval call binding the contract method 0xec9c6c30.
//
// Solidity: function collateralDecimals() view returns(uint8)
func (_LendingController *LendingControllerCaller) CollateralDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "collateralDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CollateralDecimals is a free data retrieval call binding the contract method 0xec9c6c30.
//
// Solidity: function collateralDecimals() view returns(uint8)
func (_LendingController *LendingControllerSession) CollateralDecimals() (uint8, error) {
	return _LendingController.Contract.CollateralDecimals(&_LendingController.CallOpts)
}

// CollateralDecimals is a free data retrieval call binding the contract method 0xec9c6c30.
//
// Solidity: function collateralDecimals() view returns(uint8)
func (_LendingController *LendingControllerCallerSession) CollateralDecimals() (uint8, error) {
	return _LendingController.Contract.CollateralDecimals(&_LendingController.CallOpts)
}

// CollateralValueOf is a free data retrieval call binding the contract method 0x3a65a350.
//
// Solidity: function collateralValueOf(address borrower) view returns(uint256)
func (_LendingController *LendingControllerCaller) CollateralValueOf(opts *bind.CallOpts, borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "collateralValueOf", borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralValueOf is a free data retrieval call binding the contract method 0x3a65a350.
//
// Solidity: function collateralValueOf(address borrower) view returns(uint256)
func (_LendingController *LendingControllerSession) CollateralValueOf(borrower common.Address) (*big.Int, error) {
	return _LendingController.Contract.CollateralValueOf(&_LendingController.CallOpts, borrower)
}

// CollateralValueOf is a free data retrieval call binding the contract method 0x3a65a350.
//
// Solidity: function collateralValueOf(address borrower) view returns(uint256)
func (_LendingController *LendingControllerCallerSession) CollateralValueOf(borrower common.Address) (*big.Int, error) {
	return _LendingController.Contract.CollateralValueOf(&_LendingController.CallOpts, borrower)
}

// DebtAsset is a free data retrieval call binding the contract method 0xa919802d.
//
// Solidity: function debtAsset() view returns(address)
func (_LendingController *LendingControllerCaller) DebtAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "debtAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DebtAsset is a free data retrieval call binding the contract method 0xa919802d.
//
// Solidity: function debtAsset() view returns(address)
func (_LendingController *LendingControllerSession) DebtAsset() (common.Address, error) {
	return _LendingController.Contract.DebtAsset(&_LendingController.CallOpts)
}

// DebtAsset is a free data retrieval call binding the contract method 0xa919802d.
//
// Solidity: function debtAsset() view returns(address)
func (_LendingController *LendingControllerCallerSession) DebtAsset() (common.Address, error) {
	return _LendingController.Contract.DebtAsset(&_LendingController.CallOpts)
}

// DebtDecimals is a free data retrieval call binding the contract method 0xf341e228.
//
// Solidity: function debtDecimals() view returns(uint8)
func (_LendingController *LendingControllerCaller) DebtDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "debtDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DebtDecimals is a free data retrieval call binding the contract method 0xf341e228.
//
// Solidity: function debtDecimals() view returns(uint8)
func (_LendingController *LendingControllerSession) DebtDecimals() (uint8, error) {
	return _LendingController.Contract.DebtDecimals(&_LendingController.CallOpts)
}

// DebtDecimals is a free data retrieval call binding the contract method 0xf341e228.
//
// Solidity: function debtDecimals() view returns(uint8)
func (_LendingController *LendingControllerCallerSession) DebtDecimals() (uint8, error) {
	return _LendingController.Contract.DebtDecimals(&_LendingController.CallOpts)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address borrower) view returns(uint256)
func (_LendingController *LendingControllerCaller) DebtOf(opts *bind.CallOpts, borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "debtOf", borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address borrower) view returns(uint256)
func (_LendingController *LendingControllerSession) DebtOf(borrower common.Address) (*big.Int, error) {
	return _LendingController.Contract.DebtOf(&_LendingController.CallOpts, borrower)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address borrower) view returns(uint256)
func (_LendingController *LendingControllerCallerSession) DebtOf(borrower common.Address) (*big.Int, error) {
	return _LendingController.Contract.DebtOf(&_LendingController.CallOpts, borrower)
}

// DebtValueOf is a free data retrieval call binding the contract method 0x2f573910.
//
// Solidity: function debtValueOf(address borrower) view returns(uint256)
func (_LendingController *LendingControllerCaller) DebtValueOf(opts *bind.CallOpts, borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "debtValueOf", borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtValueOf is a free data retrieval call binding the contract method 0x2f573910.
//
// Solidity: function debtValueOf(address borrower) view returns(uint256)
func (_LendingController *LendingControllerSession) DebtValueOf(borrower common.Address) (*big.Int, error) {
	return _LendingController.Contract.DebtValueOf(&_LendingController.CallOpts, borrower)
}

// DebtValueOf is a free data retrieval call binding the contract method 0x2f573910.
//
// Solidity: function debtValueOf(address borrower) view returns(uint256)
func (_LendingController *LendingControllerCallerSession) DebtValueOf(borrower common.Address) (*big.Int, error) {
	return _LendingController.Contract.DebtValueOf(&_LendingController.CallOpts, borrower)
}

// HealthFactorBps is a free data retrieval call binding the contract method 0xd10d7f21.
//
// Solidity: function healthFactorBps(address borrower) view returns(uint256)
func (_LendingController *LendingControllerCaller) HealthFactorBps(opts *bind.CallOpts, borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "healthFactorBps", borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HealthFactorBps is a free data retrieval call binding the contract method 0xd10d7f21.
//
// Solidity: function healthFactorBps(address borrower) view returns(uint256)
func (_LendingController *LendingControllerSession) HealthFactorBps(borrower common.Address) (*big.Int, error) {
	return _LendingController.Contract.HealthFactorBps(&_LendingController.CallOpts, borrower)
}

// HealthFactorBps is a free data retrieval call binding the contract method 0xd10d7f21.
//
// Solidity: function healthFactorBps(address borrower) view returns(uint256)
func (_LendingController *LendingControllerCallerSession) HealthFactorBps(borrower common.Address) (*big.Int, error) {
	return _LendingController.Contract.HealthFactorBps(&_LendingController.CallOpts, borrower)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address borrower) view returns(bool)
func (_LendingController *LendingControllerCaller) IsLiquidatable(opts *bind.CallOpts, borrower common.Address) (bool, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "isLiquidatable", borrower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address borrower) view returns(bool)
func (_LendingController *LendingControllerSession) IsLiquidatable(borrower common.Address) (bool, error) {
	return _LendingController.Contract.IsLiquidatable(&_LendingController.CallOpts, borrower)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address borrower) view returns(bool)
func (_LendingController *LendingControllerCallerSession) IsLiquidatable(borrower common.Address) (bool, error) {
	return _LendingController.Contract.IsLiquidatable(&_LendingController.CallOpts, borrower)
}

// LiquidationBonusBps is a free data retrieval call binding the contract method 0x19970d8e.
//
// Solidity: function liquidationBonusBps() view returns(uint16)
func (_LendingController *LendingControllerCaller) LiquidationBonusBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "liquidationBonusBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LiquidationBonusBps is a free data retrieval call binding the contract method 0x19970d8e.
//
// Solidity: function liquidationBonusBps() view returns(uint16)
func (_LendingController *LendingControllerSession) LiquidationBonusBps() (uint16, error) {
	return _LendingController.Contract.LiquidationBonusBps(&_LendingController.CallOpts)
}

// LiquidationBonusBps is a free data retrieval call binding the contract method 0x19970d8e.
//
// Solidity: function liquidationBonusBps() view returns(uint16)
func (_LendingController *LendingControllerCallerSession) LiquidationBonusBps() (uint16, error) {
	return _LendingController.Contract.LiquidationBonusBps(&_LendingController.CallOpts)
}

// LiquidationThresholdBps is a free data retrieval call binding the contract method 0xe4864731.
//
// Solidity: function liquidationThresholdBps() view returns(uint16)
func (_LendingController *LendingControllerCaller) LiquidationThresholdBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "liquidationThresholdBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LiquidationThresholdBps is a free data retrieval call binding the contract method 0xe4864731.
//
// Solidity: function liquidationThresholdBps() view returns(uint16)
func (_LendingController *LendingControllerSession) LiquidationThresholdBps() (uint16, error) {
	return _LendingController.Contract.LiquidationThresholdBps(&_LendingController.CallOpts)
}

// LiquidationThresholdBps is a free data retrieval call binding the contract method 0xe4864731.
//
// Solidity: function liquidationThresholdBps() view returns(uint16)
func (_LendingController *LendingControllerCallerSession) LiquidationThresholdBps() (uint16, error) {
	return _LendingController.Contract.LiquidationThresholdBps(&_LendingController.CallOpts)
}

// MaxBorrowable is a free data retrieval call binding the contract method 0x53696eb8.
//
// Solidity: function maxBorrowable(address borrower) view returns(uint256)
func (_LendingController *LendingControllerCaller) MaxBorrowable(opts *bind.CallOpts, borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "maxBorrowable", borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBorrowable is a free data retrieval call binding the contract method 0x53696eb8.
//
// Solidity: function maxBorrowable(address borrower) view returns(uint256)
func (_LendingController *LendingControllerSession) MaxBorrowable(borrower common.Address) (*big.Int, error) {
	return _LendingController.Contract.MaxBorrowable(&_LendingController.CallOpts, borrower)
}

// MaxBorrowable is a free data retrieval call binding the contract method 0x53696eb8.
//
// Solidity: function maxBorrowable(address borrower) view returns(uint256)
func (_LendingController *LendingControllerCallerSession) MaxBorrowable(borrower common.Address) (*big.Int, error) {
	return _LendingController.Contract.MaxBorrowable(&_LendingController.CallOpts, borrower)
}

// MaxLtvBps is a free data retrieval call binding the contract method 0xea7f52d1.
//
// Solidity: function maxLtvBps() view returns(uint16)
func (_LendingController *LendingControllerCaller) MaxLtvBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "maxLtvBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaxLtvBps is a free data retrieval call binding the contract method 0xea7f52d1.
//
// Solidity: function maxLtvBps() view returns(uint16)
func (_LendingController *LendingControllerSession) MaxLtvBps() (uint16, error) {
	return _LendingController.Contract.MaxLtvBps(&_LendingController.CallOpts)
}

// MaxLtvBps is a free data retrieval call binding the contract method 0xea7f52d1.
//
// Solidity: function maxLtvBps() view returns(uint16)
func (_LendingController *LendingControllerCallerSession) MaxLtvBps() (uint16, error) {
	return _LendingController.Contract.MaxLtvBps(&_LendingController.CallOpts)
}

// MaxWithdrawableCollateral is a free data retrieval call binding the contract method 0xc93853af.
//
// Solidity: function maxWithdrawableCollateral(address borrower) view returns(uint256)
func (_LendingController *LendingControllerCaller) MaxWithdrawableCollateral(opts *bind.CallOpts, borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "maxWithdrawableCollateral", borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdrawableCollateral is a free data retrieval call binding the contract method 0xc93853af.
//
// Solidity: function maxWithdrawableCollateral(address borrower) view returns(uint256)
func (_LendingController *LendingControllerSession) MaxWithdrawableCollateral(borrower common.Address) (*big.Int, error) {
	return _LendingController.Contract.MaxWithdrawableCollateral(&_LendingController.CallOpts, borrower)
}

// MaxWithdrawableCollateral is a free data retrieval call binding the contract method 0xc93853af.
//
// Solidity: function maxWithdrawableCollateral(address borrower) view returns(uint256)
func (_LendingController *LendingControllerCallerSession) MaxWithdrawableCollateral(borrower common.Address) (*big.Int, error) {
	return _LendingController.Contract.MaxWithdrawableCollateral(&_LendingController.CallOpts, borrower)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_LendingController *LendingControllerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_LendingController *LendingControllerSession) Oracle() (common.Address, error) {
	return _LendingController.Contract.Oracle(&_LendingController.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_LendingController *LendingControllerCallerSession) Oracle() (common.Address, error) {
	return _LendingController.Contract.Oracle(&_LendingController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingController *LendingControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingController *LendingControllerSession) Owner() (common.Address, error) {
	return _LendingController.Contract.Owner(&_LendingController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingController *LendingControllerCallerSession) Owner() (common.Address, error) {
	return _LendingController.Contract.Owner(&_LendingController.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_LendingController *LendingControllerCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_LendingController *LendingControllerSession) Pool() (common.Address, error) {
	return _LendingController.Contract.Pool(&_LendingController.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_LendingController *LendingControllerCallerSession) Pool() (common.Address, error) {
	return _LendingController.Contract.Pool(&_LendingController.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_LendingController *LendingControllerCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingController.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_LendingController *LendingControllerSession) Vault() (common.Address, error) {
	return _LendingController.Contract.Vault(&_LendingController.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_LendingController *LendingControllerCallerSession) Vault() (common.Address, error) {
	return _LendingController.Contract.Vault(&_LendingController.CallOpts)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 amount) returns()
func (_LendingController *LendingControllerTransactor) Borrow(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LendingController.contract.Transact(opts, "borrow", amount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 amount) returns()
func (_LendingController *LendingControllerSession) Borrow(amount *big.Int) (*types.Transaction, error) {
	return _LendingController.Contract.Borrow(&_LendingController.TransactOpts, amount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 amount) returns()
func (_LendingController *LendingControllerTransactorSession) Borrow(amount *big.Int) (*types.Transaction, error) {
	return _LendingController.Contract.Borrow(&_LendingController.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingController *LendingControllerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingController.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingController *LendingControllerSession) RenounceOwnership() (*types.Transaction, error) {
	return _LendingController.Contract.RenounceOwnership(&_LendingController.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingController *LendingControllerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LendingController.Contract.RenounceOwnership(&_LendingController.TransactOpts)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 amount) returns()
func (_LendingController *LendingControllerTransactor) Repay(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LendingController.contract.Transact(opts, "repay", amount)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 amount) returns()
func (_LendingController *LendingControllerSession) Repay(amount *big.Int) (*types.Transaction, error) {
	return _LendingController.Contract.Repay(&_LendingController.TransactOpts, amount)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 amount) returns()
func (_LendingController *LendingControllerTransactorSession) Repay(amount *big.Int) (*types.Transaction, error) {
	return _LendingController.Contract.Repay(&_LendingController.TransactOpts, amount)
}

// RepayAll is a paid mutator transaction binding the contract method 0xfa3ae6dc.
//
// Solidity: function repayAll() returns(uint256 amountPaid)
func (_LendingController *LendingControllerTransactor) RepayAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingController.contract.Transact(opts, "repayAll")
}

// RepayAll is a paid mutator transaction binding the contract method 0xfa3ae6dc.
//
// Solidity: function repayAll() returns(uint256 amountPaid)
func (_LendingController *LendingControllerSession) RepayAll() (*types.Transaction, error) {
	return _LendingController.Contract.RepayAll(&_LendingController.TransactOpts)
}

// RepayAll is a paid mutator transaction binding the contract method 0xfa3ae6dc.
//
// Solidity: function repayAll() returns(uint256 amountPaid)
func (_LendingController *LendingControllerTransactorSession) RepayAll() (*types.Transaction, error) {
	return _LendingController.Contract.RepayAll(&_LendingController.TransactOpts)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x939752bd.
//
// Solidity: function setBorrowPaused(bool isPaused) returns()
func (_LendingController *LendingControllerTransactor) SetBorrowPaused(opts *bind.TransactOpts, isPaused bool) (*types.Transaction, error) {
	return _LendingController.contract.Transact(opts, "setBorrowPaused", isPaused)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x939752bd.
//
// Solidity: function setBorrowPaused(bool isPaused) returns()
func (_LendingController *LendingControllerSession) SetBorrowPaused(isPaused bool) (*types.Transaction, error) {
	return _LendingController.Contract.SetBorrowPaused(&_LendingController.TransactOpts, isPaused)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x939752bd.
//
// Solidity: function setBorrowPaused(bool isPaused) returns()
func (_LendingController *LendingControllerTransactorSession) SetBorrowPaused(isPaused bool) (*types.Transaction, error) {
	return _LendingController.Contract.SetBorrowPaused(&_LendingController.TransactOpts, isPaused)
}

// SetRiskSettings is a paid mutator transaction binding the contract method 0x1a941656.
//
// Solidity: function setRiskSettings(uint16 newMaxLtvBps, uint16 newThresholdBps, uint16 newBonusBps) returns()
func (_LendingController *LendingControllerTransactor) SetRiskSettings(opts *bind.TransactOpts, newMaxLtvBps uint16, newThresholdBps uint16, newBonusBps uint16) (*types.Transaction, error) {
	return _LendingController.contract.Transact(opts, "setRiskSettings", newMaxLtvBps, newThresholdBps, newBonusBps)
}

// SetRiskSettings is a paid mutator transaction binding the contract method 0x1a941656.
//
// Solidity: function setRiskSettings(uint16 newMaxLtvBps, uint16 newThresholdBps, uint16 newBonusBps) returns()
func (_LendingController *LendingControllerSession) SetRiskSettings(newMaxLtvBps uint16, newThresholdBps uint16, newBonusBps uint16) (*types.Transaction, error) {
	return _LendingController.Contract.SetRiskSettings(&_LendingController.TransactOpts, newMaxLtvBps, newThresholdBps, newBonusBps)
}

// SetRiskSettings is a paid mutator transaction binding the contract method 0x1a941656.
//
// Solidity: function setRiskSettings(uint16 newMaxLtvBps, uint16 newThresholdBps, uint16 newBonusBps) returns()
func (_LendingController *LendingControllerTransactorSession) SetRiskSettings(newMaxLtvBps uint16, newThresholdBps uint16, newBonusBps uint16) (*types.Transaction, error) {
	return _LendingController.Contract.SetRiskSettings(&_LendingController.TransactOpts, newMaxLtvBps, newThresholdBps, newBonusBps)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingController *LendingControllerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LendingController.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingController *LendingControllerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LendingController.Contract.TransferOwnership(&_LendingController.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingController *LendingControllerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LendingController.Contract.TransferOwnership(&_LendingController.TransactOpts, newOwner)
}

// LendingControllerBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the LendingController contract.
type LendingControllerBorrowIterator struct {
	Event *LendingControllerBorrow // Event containing the contract specifics and raw log

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
func (it *LendingControllerBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingControllerBorrow)
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
		it.Event = new(LendingControllerBorrow)
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
func (it *LendingControllerBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingControllerBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingControllerBorrow represents a Borrow event raised by the LendingController contract.
type LendingControllerBorrow struct {
	Borrower        common.Address
	Amount          *big.Int
	NewDebt         *big.Int
	HealthFactorBps *big.Int
	BorrowIndex     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x2dd79f4fccfd18c360ce7f9132f3621bf05eee18f995224badb32d17f172df73.
//
// Solidity: event Borrow(address indexed borrower, uint256 amount, uint256 newDebt, uint256 healthFactorBps, uint256 borrowIndex)
func (_LendingController *LendingControllerFilterer) FilterBorrow(opts *bind.FilterOpts, borrower []common.Address) (*LendingControllerBorrowIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendingController.contract.FilterLogs(opts, "Borrow", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LendingControllerBorrowIterator{contract: _LendingController.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x2dd79f4fccfd18c360ce7f9132f3621bf05eee18f995224badb32d17f172df73.
//
// Solidity: event Borrow(address indexed borrower, uint256 amount, uint256 newDebt, uint256 healthFactorBps, uint256 borrowIndex)
func (_LendingController *LendingControllerFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *LendingControllerBorrow, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendingController.contract.WatchLogs(opts, "Borrow", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingControllerBorrow)
				if err := _LendingController.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x2dd79f4fccfd18c360ce7f9132f3621bf05eee18f995224badb32d17f172df73.
//
// Solidity: event Borrow(address indexed borrower, uint256 amount, uint256 newDebt, uint256 healthFactorBps, uint256 borrowIndex)
func (_LendingController *LendingControllerFilterer) ParseBorrow(log types.Log) (*LendingControllerBorrow, error) {
	event := new(LendingControllerBorrow)
	if err := _LendingController.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingControllerBorrowPausedChangedIterator is returned from FilterBorrowPausedChanged and is used to iterate over the raw logs and unpacked data for BorrowPausedChanged events raised by the LendingController contract.
type LendingControllerBorrowPausedChangedIterator struct {
	Event *LendingControllerBorrowPausedChanged // Event containing the contract specifics and raw log

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
func (it *LendingControllerBorrowPausedChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingControllerBorrowPausedChanged)
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
		it.Event = new(LendingControllerBorrowPausedChanged)
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
func (it *LendingControllerBorrowPausedChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingControllerBorrowPausedChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingControllerBorrowPausedChanged represents a BorrowPausedChanged event raised by the LendingController contract.
type LendingControllerBorrowPausedChanged struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBorrowPausedChanged is a free log retrieval operation binding the contract event 0xdea251e46ae106c2419bdad1fe673dbe317d9f062d6db7a89caca724e876fe93.
//
// Solidity: event BorrowPausedChanged(bool isPaused)
func (_LendingController *LendingControllerFilterer) FilterBorrowPausedChanged(opts *bind.FilterOpts) (*LendingControllerBorrowPausedChangedIterator, error) {

	logs, sub, err := _LendingController.contract.FilterLogs(opts, "BorrowPausedChanged")
	if err != nil {
		return nil, err
	}
	return &LendingControllerBorrowPausedChangedIterator{contract: _LendingController.contract, event: "BorrowPausedChanged", logs: logs, sub: sub}, nil
}

// WatchBorrowPausedChanged is a free log subscription operation binding the contract event 0xdea251e46ae106c2419bdad1fe673dbe317d9f062d6db7a89caca724e876fe93.
//
// Solidity: event BorrowPausedChanged(bool isPaused)
func (_LendingController *LendingControllerFilterer) WatchBorrowPausedChanged(opts *bind.WatchOpts, sink chan<- *LendingControllerBorrowPausedChanged) (event.Subscription, error) {

	logs, sub, err := _LendingController.contract.WatchLogs(opts, "BorrowPausedChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingControllerBorrowPausedChanged)
				if err := _LendingController.contract.UnpackLog(event, "BorrowPausedChanged", log); err != nil {
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

// ParseBorrowPausedChanged is a log parse operation binding the contract event 0xdea251e46ae106c2419bdad1fe673dbe317d9f062d6db7a89caca724e876fe93.
//
// Solidity: event BorrowPausedChanged(bool isPaused)
func (_LendingController *LendingControllerFilterer) ParseBorrowPausedChanged(log types.Log) (*LendingControllerBorrowPausedChanged, error) {
	event := new(LendingControllerBorrowPausedChanged)
	if err := _LendingController.contract.UnpackLog(event, "BorrowPausedChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingControllerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LendingController contract.
type LendingControllerOwnershipTransferredIterator struct {
	Event *LendingControllerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LendingControllerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingControllerOwnershipTransferred)
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
		it.Event = new(LendingControllerOwnershipTransferred)
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
func (it *LendingControllerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingControllerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingControllerOwnershipTransferred represents a OwnershipTransferred event raised by the LendingController contract.
type LendingControllerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingController *LendingControllerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LendingControllerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LendingController.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LendingControllerOwnershipTransferredIterator{contract: _LendingController.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingController *LendingControllerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LendingControllerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LendingController.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingControllerOwnershipTransferred)
				if err := _LendingController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LendingController *LendingControllerFilterer) ParseOwnershipTransferred(log types.Log) (*LendingControllerOwnershipTransferred, error) {
	event := new(LendingControllerOwnershipTransferred)
	if err := _LendingController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingControllerRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the LendingController contract.
type LendingControllerRepayIterator struct {
	Event *LendingControllerRepay // Event containing the contract specifics and raw log

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
func (it *LendingControllerRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingControllerRepay)
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
		it.Event = new(LendingControllerRepay)
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
func (it *LendingControllerRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingControllerRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingControllerRepay represents a Repay event raised by the LendingController contract.
type LendingControllerRepay struct {
	Borrower common.Address
	Payer    common.Address
	Amount   *big.Int
	NewDebt  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0xe4a1ae657f49cb1fb1c7d3a94ae6093565c4c8c0e03de488f79c377c3c3a24e0.
//
// Solidity: event Repay(address indexed borrower, address indexed payer, uint256 amount, uint256 newDebt)
func (_LendingController *LendingControllerFilterer) FilterRepay(opts *bind.FilterOpts, borrower []common.Address, payer []common.Address) (*LendingControllerRepayIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _LendingController.contract.FilterLogs(opts, "Repay", borrowerRule, payerRule)
	if err != nil {
		return nil, err
	}
	return &LendingControllerRepayIterator{contract: _LendingController.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0xe4a1ae657f49cb1fb1c7d3a94ae6093565c4c8c0e03de488f79c377c3c3a24e0.
//
// Solidity: event Repay(address indexed borrower, address indexed payer, uint256 amount, uint256 newDebt)
func (_LendingController *LendingControllerFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *LendingControllerRepay, borrower []common.Address, payer []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _LendingController.contract.WatchLogs(opts, "Repay", borrowerRule, payerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingControllerRepay)
				if err := _LendingController.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0xe4a1ae657f49cb1fb1c7d3a94ae6093565c4c8c0e03de488f79c377c3c3a24e0.
//
// Solidity: event Repay(address indexed borrower, address indexed payer, uint256 amount, uint256 newDebt)
func (_LendingController *LendingControllerFilterer) ParseRepay(log types.Log) (*LendingControllerRepay, error) {
	event := new(LendingControllerRepay)
	if err := _LendingController.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingControllerRiskSettingsChangedIterator is returned from FilterRiskSettingsChanged and is used to iterate over the raw logs and unpacked data for RiskSettingsChanged events raised by the LendingController contract.
type LendingControllerRiskSettingsChangedIterator struct {
	Event *LendingControllerRiskSettingsChanged // Event containing the contract specifics and raw log

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
func (it *LendingControllerRiskSettingsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingControllerRiskSettingsChanged)
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
		it.Event = new(LendingControllerRiskSettingsChanged)
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
func (it *LendingControllerRiskSettingsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingControllerRiskSettingsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingControllerRiskSettingsChanged represents a RiskSettingsChanged event raised by the LendingController contract.
type LendingControllerRiskSettingsChanged struct {
	MaxLtvBps               uint16
	LiquidationThresholdBps uint16
	LiquidationBonusBps     uint16
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterRiskSettingsChanged is a free log retrieval operation binding the contract event 0x4a77d301e7b50c3f6025fa9aa0b0fbaa8ea0f56bfb4368a779a0e7feba0906dd.
//
// Solidity: event RiskSettingsChanged(uint16 maxLtvBps, uint16 liquidationThresholdBps, uint16 liquidationBonusBps)
func (_LendingController *LendingControllerFilterer) FilterRiskSettingsChanged(opts *bind.FilterOpts) (*LendingControllerRiskSettingsChangedIterator, error) {

	logs, sub, err := _LendingController.contract.FilterLogs(opts, "RiskSettingsChanged")
	if err != nil {
		return nil, err
	}
	return &LendingControllerRiskSettingsChangedIterator{contract: _LendingController.contract, event: "RiskSettingsChanged", logs: logs, sub: sub}, nil
}

// WatchRiskSettingsChanged is a free log subscription operation binding the contract event 0x4a77d301e7b50c3f6025fa9aa0b0fbaa8ea0f56bfb4368a779a0e7feba0906dd.
//
// Solidity: event RiskSettingsChanged(uint16 maxLtvBps, uint16 liquidationThresholdBps, uint16 liquidationBonusBps)
func (_LendingController *LendingControllerFilterer) WatchRiskSettingsChanged(opts *bind.WatchOpts, sink chan<- *LendingControllerRiskSettingsChanged) (event.Subscription, error) {

	logs, sub, err := _LendingController.contract.WatchLogs(opts, "RiskSettingsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingControllerRiskSettingsChanged)
				if err := _LendingController.contract.UnpackLog(event, "RiskSettingsChanged", log); err != nil {
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

// ParseRiskSettingsChanged is a log parse operation binding the contract event 0x4a77d301e7b50c3f6025fa9aa0b0fbaa8ea0f56bfb4368a779a0e7feba0906dd.
//
// Solidity: event RiskSettingsChanged(uint16 maxLtvBps, uint16 liquidationThresholdBps, uint16 liquidationBonusBps)
func (_LendingController *LendingControllerFilterer) ParseRiskSettingsChanged(log types.Log) (*LendingControllerRiskSettingsChanged, error) {
	event := new(LendingControllerRiskSettingsChanged)
	if err := _LendingController.contract.UnpackLog(event, "RiskSettingsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
