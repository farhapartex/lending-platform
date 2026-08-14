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

// CollateralVaultMetaData contains all meta data concerning the CollateralVault contract.
var CollateralVaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultAsset\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collateralAsset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"controller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILendingController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositCollateral\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"linkController\",\"inputs\":[{\"name\":\"newController\",\"type\":\"address\",\"internalType\":\"contractILendingController\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"linkLiquidationManager\",\"inputs\":[{\"name\":\"newLiquidationManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"liquidationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"seize\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalCollateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawCollateral\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CollateralDeposited\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newCollateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollateralSeized\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newCollateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollateralWithdrawn\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newCollateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ControllerLinked\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidationManagerLinked\",\"inputs\":[{\"name\":\"liquidationManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedsCollateralBalance\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"held\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"WouldBreakBorrowLimit\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"safeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]}]",
}

// CollateralVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use CollateralVaultMetaData.ABI instead.
var CollateralVaultABI = CollateralVaultMetaData.ABI

// CollateralVault is an auto generated Go binding around an Ethereum contract.
type CollateralVault struct {
	CollateralVaultCaller     // Read-only binding to the contract
	CollateralVaultTransactor // Write-only binding to the contract
	CollateralVaultFilterer   // Log filterer for contract events
}

// CollateralVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollateralVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollateralVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollateralVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollateralVaultSession struct {
	Contract     *CollateralVault  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CollateralVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollateralVaultCallerSession struct {
	Contract *CollateralVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CollateralVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollateralVaultTransactorSession struct {
	Contract     *CollateralVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CollateralVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollateralVaultRaw struct {
	Contract *CollateralVault // Generic contract binding to access the raw methods on
}

// CollateralVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollateralVaultCallerRaw struct {
	Contract *CollateralVaultCaller // Generic read-only contract binding to access the raw methods on
}

// CollateralVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollateralVaultTransactorRaw struct {
	Contract *CollateralVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCollateralVault creates a new instance of CollateralVault, bound to a specific deployed contract.
func NewCollateralVault(address common.Address, backend bind.ContractBackend) (*CollateralVault, error) {
	contract, err := bindCollateralVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CollateralVault{CollateralVaultCaller: CollateralVaultCaller{contract: contract}, CollateralVaultTransactor: CollateralVaultTransactor{contract: contract}, CollateralVaultFilterer: CollateralVaultFilterer{contract: contract}}, nil
}

// NewCollateralVaultCaller creates a new read-only instance of CollateralVault, bound to a specific deployed contract.
func NewCollateralVaultCaller(address common.Address, caller bind.ContractCaller) (*CollateralVaultCaller, error) {
	contract, err := bindCollateralVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultCaller{contract: contract}, nil
}

// NewCollateralVaultTransactor creates a new write-only instance of CollateralVault, bound to a specific deployed contract.
func NewCollateralVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*CollateralVaultTransactor, error) {
	contract, err := bindCollateralVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultTransactor{contract: contract}, nil
}

// NewCollateralVaultFilterer creates a new log filterer instance of CollateralVault, bound to a specific deployed contract.
func NewCollateralVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*CollateralVaultFilterer, error) {
	contract, err := bindCollateralVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultFilterer{contract: contract}, nil
}

// bindCollateralVault binds a generic wrapper to an already deployed contract.
func bindCollateralVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CollateralVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollateralVault *CollateralVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollateralVault.Contract.CollateralVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollateralVault *CollateralVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollateralVault.Contract.CollateralVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollateralVault *CollateralVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollateralVault.Contract.CollateralVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollateralVault *CollateralVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollateralVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollateralVault *CollateralVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollateralVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollateralVault *CollateralVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollateralVault.Contract.contract.Transact(opts, method, params...)
}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_CollateralVault *CollateralVaultCaller) CollateralAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "collateralAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_CollateralVault *CollateralVaultSession) CollateralAsset() (common.Address, error) {
	return _CollateralVault.Contract.CollateralAsset(&_CollateralVault.CallOpts)
}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_CollateralVault *CollateralVaultCallerSession) CollateralAsset() (common.Address, error) {
	return _CollateralVault.Contract.CollateralAsset(&_CollateralVault.CallOpts)
}

// CollateralOf is a free data retrieval call binding the contract method 0x1aefb107.
//
// Solidity: function collateralOf(address ) view returns(uint256)
func (_CollateralVault *CollateralVaultCaller) CollateralOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "collateralOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralOf is a free data retrieval call binding the contract method 0x1aefb107.
//
// Solidity: function collateralOf(address ) view returns(uint256)
func (_CollateralVault *CollateralVaultSession) CollateralOf(arg0 common.Address) (*big.Int, error) {
	return _CollateralVault.Contract.CollateralOf(&_CollateralVault.CallOpts, arg0)
}

// CollateralOf is a free data retrieval call binding the contract method 0x1aefb107.
//
// Solidity: function collateralOf(address ) view returns(uint256)
func (_CollateralVault *CollateralVaultCallerSession) CollateralOf(arg0 common.Address) (*big.Int, error) {
	return _CollateralVault.Contract.CollateralOf(&_CollateralVault.CallOpts, arg0)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_CollateralVault *CollateralVaultCaller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "collateralToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_CollateralVault *CollateralVaultSession) CollateralToken() (common.Address, error) {
	return _CollateralVault.Contract.CollateralToken(&_CollateralVault.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_CollateralVault *CollateralVaultCallerSession) CollateralToken() (common.Address, error) {
	return _CollateralVault.Contract.CollateralToken(&_CollateralVault.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CollateralVault *CollateralVaultCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CollateralVault *CollateralVaultSession) Controller() (common.Address, error) {
	return _CollateralVault.Contract.Controller(&_CollateralVault.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CollateralVault *CollateralVaultCallerSession) Controller() (common.Address, error) {
	return _CollateralVault.Contract.Controller(&_CollateralVault.CallOpts)
}

// LiquidationManager is a free data retrieval call binding the contract method 0x1ef3a04c.
//
// Solidity: function liquidationManager() view returns(address)
func (_CollateralVault *CollateralVaultCaller) LiquidationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "liquidationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidationManager is a free data retrieval call binding the contract method 0x1ef3a04c.
//
// Solidity: function liquidationManager() view returns(address)
func (_CollateralVault *CollateralVaultSession) LiquidationManager() (common.Address, error) {
	return _CollateralVault.Contract.LiquidationManager(&_CollateralVault.CallOpts)
}

// LiquidationManager is a free data retrieval call binding the contract method 0x1ef3a04c.
//
// Solidity: function liquidationManager() view returns(address)
func (_CollateralVault *CollateralVaultCallerSession) LiquidationManager() (common.Address, error) {
	return _CollateralVault.Contract.LiquidationManager(&_CollateralVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CollateralVault *CollateralVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CollateralVault *CollateralVaultSession) Owner() (common.Address, error) {
	return _CollateralVault.Contract.Owner(&_CollateralVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CollateralVault *CollateralVaultCallerSession) Owner() (common.Address, error) {
	return _CollateralVault.Contract.Owner(&_CollateralVault.CallOpts)
}

// TotalCollateral is a free data retrieval call binding the contract method 0x4ac8eb5f.
//
// Solidity: function totalCollateral() view returns(uint256)
func (_CollateralVault *CollateralVaultCaller) TotalCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CollateralVault.contract.Call(opts, &out, "totalCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCollateral is a free data retrieval call binding the contract method 0x4ac8eb5f.
//
// Solidity: function totalCollateral() view returns(uint256)
func (_CollateralVault *CollateralVaultSession) TotalCollateral() (*big.Int, error) {
	return _CollateralVault.Contract.TotalCollateral(&_CollateralVault.CallOpts)
}

// TotalCollateral is a free data retrieval call binding the contract method 0x4ac8eb5f.
//
// Solidity: function totalCollateral() view returns(uint256)
func (_CollateralVault *CollateralVaultCallerSession) TotalCollateral() (*big.Int, error) {
	return _CollateralVault.Contract.TotalCollateral(&_CollateralVault.CallOpts)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xbad4a01f.
//
// Solidity: function depositCollateral(uint256 amount) returns()
func (_CollateralVault *CollateralVaultTransactor) DepositCollateral(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "depositCollateral", amount)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xbad4a01f.
//
// Solidity: function depositCollateral(uint256 amount) returns()
func (_CollateralVault *CollateralVaultSession) DepositCollateral(amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.DepositCollateral(&_CollateralVault.TransactOpts, amount)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xbad4a01f.
//
// Solidity: function depositCollateral(uint256 amount) returns()
func (_CollateralVault *CollateralVaultTransactorSession) DepositCollateral(amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.DepositCollateral(&_CollateralVault.TransactOpts, amount)
}

// LinkController is a paid mutator transaction binding the contract method 0xdf95ec90.
//
// Solidity: function linkController(address newController) returns()
func (_CollateralVault *CollateralVaultTransactor) LinkController(opts *bind.TransactOpts, newController common.Address) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "linkController", newController)
}

// LinkController is a paid mutator transaction binding the contract method 0xdf95ec90.
//
// Solidity: function linkController(address newController) returns()
func (_CollateralVault *CollateralVaultSession) LinkController(newController common.Address) (*types.Transaction, error) {
	return _CollateralVault.Contract.LinkController(&_CollateralVault.TransactOpts, newController)
}

// LinkController is a paid mutator transaction binding the contract method 0xdf95ec90.
//
// Solidity: function linkController(address newController) returns()
func (_CollateralVault *CollateralVaultTransactorSession) LinkController(newController common.Address) (*types.Transaction, error) {
	return _CollateralVault.Contract.LinkController(&_CollateralVault.TransactOpts, newController)
}

// LinkLiquidationManager is a paid mutator transaction binding the contract method 0xc894270c.
//
// Solidity: function linkLiquidationManager(address newLiquidationManager) returns()
func (_CollateralVault *CollateralVaultTransactor) LinkLiquidationManager(opts *bind.TransactOpts, newLiquidationManager common.Address) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "linkLiquidationManager", newLiquidationManager)
}

// LinkLiquidationManager is a paid mutator transaction binding the contract method 0xc894270c.
//
// Solidity: function linkLiquidationManager(address newLiquidationManager) returns()
func (_CollateralVault *CollateralVaultSession) LinkLiquidationManager(newLiquidationManager common.Address) (*types.Transaction, error) {
	return _CollateralVault.Contract.LinkLiquidationManager(&_CollateralVault.TransactOpts, newLiquidationManager)
}

// LinkLiquidationManager is a paid mutator transaction binding the contract method 0xc894270c.
//
// Solidity: function linkLiquidationManager(address newLiquidationManager) returns()
func (_CollateralVault *CollateralVaultTransactorSession) LinkLiquidationManager(newLiquidationManager common.Address) (*types.Transaction, error) {
	return _CollateralVault.Contract.LinkLiquidationManager(&_CollateralVault.TransactOpts, newLiquidationManager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CollateralVault *CollateralVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CollateralVault *CollateralVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _CollateralVault.Contract.RenounceOwnership(&_CollateralVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CollateralVault *CollateralVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CollateralVault.Contract.RenounceOwnership(&_CollateralVault.TransactOpts)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address borrower, address recipient, uint256 amount) returns()
func (_CollateralVault *CollateralVaultTransactor) Seize(opts *bind.TransactOpts, borrower common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "seize", borrower, recipient, amount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address borrower, address recipient, uint256 amount) returns()
func (_CollateralVault *CollateralVaultSession) Seize(borrower common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.Seize(&_CollateralVault.TransactOpts, borrower, recipient, amount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address borrower, address recipient, uint256 amount) returns()
func (_CollateralVault *CollateralVaultTransactorSession) Seize(borrower common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.Seize(&_CollateralVault.TransactOpts, borrower, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CollateralVault *CollateralVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CollateralVault *CollateralVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CollateralVault.Contract.TransferOwnership(&_CollateralVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CollateralVault *CollateralVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CollateralVault.Contract.TransferOwnership(&_CollateralVault.TransactOpts, newOwner)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 amount) returns()
func (_CollateralVault *CollateralVaultTransactor) WithdrawCollateral(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.contract.Transact(opts, "withdrawCollateral", amount)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 amount) returns()
func (_CollateralVault *CollateralVaultSession) WithdrawCollateral(amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.WithdrawCollateral(&_CollateralVault.TransactOpts, amount)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x6112fe2e.
//
// Solidity: function withdrawCollateral(uint256 amount) returns()
func (_CollateralVault *CollateralVaultTransactorSession) WithdrawCollateral(amount *big.Int) (*types.Transaction, error) {
	return _CollateralVault.Contract.WithdrawCollateral(&_CollateralVault.TransactOpts, amount)
}

// CollateralVaultCollateralDepositedIterator is returned from FilterCollateralDeposited and is used to iterate over the raw logs and unpacked data for CollateralDeposited events raised by the CollateralVault contract.
type CollateralVaultCollateralDepositedIterator struct {
	Event *CollateralVaultCollateralDeposited // Event containing the contract specifics and raw log

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
func (it *CollateralVaultCollateralDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultCollateralDeposited)
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
		it.Event = new(CollateralVaultCollateralDeposited)
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
func (it *CollateralVaultCollateralDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultCollateralDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultCollateralDeposited represents a CollateralDeposited event raised by the CollateralVault contract.
type CollateralVaultCollateralDeposited struct {
	Borrower      common.Address
	Amount        *big.Int
	NewCollateral *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollateralDeposited is a free log retrieval operation binding the contract event 0xf4d587c98d234ca4d147061e6b5167e7f41ee17f11562a9f0b49570abece859e.
//
// Solidity: event CollateralDeposited(address indexed borrower, uint256 amount, uint256 newCollateral)
func (_CollateralVault *CollateralVaultFilterer) FilterCollateralDeposited(opts *bind.FilterOpts, borrower []common.Address) (*CollateralVaultCollateralDepositedIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "CollateralDeposited", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultCollateralDepositedIterator{contract: _CollateralVault.contract, event: "CollateralDeposited", logs: logs, sub: sub}, nil
}

// WatchCollateralDeposited is a free log subscription operation binding the contract event 0xf4d587c98d234ca4d147061e6b5167e7f41ee17f11562a9f0b49570abece859e.
//
// Solidity: event CollateralDeposited(address indexed borrower, uint256 amount, uint256 newCollateral)
func (_CollateralVault *CollateralVaultFilterer) WatchCollateralDeposited(opts *bind.WatchOpts, sink chan<- *CollateralVaultCollateralDeposited, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "CollateralDeposited", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultCollateralDeposited)
				if err := _CollateralVault.contract.UnpackLog(event, "CollateralDeposited", log); err != nil {
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

// ParseCollateralDeposited is a log parse operation binding the contract event 0xf4d587c98d234ca4d147061e6b5167e7f41ee17f11562a9f0b49570abece859e.
//
// Solidity: event CollateralDeposited(address indexed borrower, uint256 amount, uint256 newCollateral)
func (_CollateralVault *CollateralVaultFilterer) ParseCollateralDeposited(log types.Log) (*CollateralVaultCollateralDeposited, error) {
	event := new(CollateralVaultCollateralDeposited)
	if err := _CollateralVault.contract.UnpackLog(event, "CollateralDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultCollateralSeizedIterator is returned from FilterCollateralSeized and is used to iterate over the raw logs and unpacked data for CollateralSeized events raised by the CollateralVault contract.
type CollateralVaultCollateralSeizedIterator struct {
	Event *CollateralVaultCollateralSeized // Event containing the contract specifics and raw log

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
func (it *CollateralVaultCollateralSeizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultCollateralSeized)
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
		it.Event = new(CollateralVaultCollateralSeized)
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
func (it *CollateralVaultCollateralSeizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultCollateralSeizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultCollateralSeized represents a CollateralSeized event raised by the CollateralVault contract.
type CollateralVaultCollateralSeized struct {
	Borrower      common.Address
	Recipient     common.Address
	Amount        *big.Int
	NewCollateral *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollateralSeized is a free log retrieval operation binding the contract event 0x2f72539f26ad4e48fa396e30d602a0397ca67099e2d58bd39663a7bc09944ee3.
//
// Solidity: event CollateralSeized(address indexed borrower, address indexed recipient, uint256 amount, uint256 newCollateral)
func (_CollateralVault *CollateralVaultFilterer) FilterCollateralSeized(opts *bind.FilterOpts, borrower []common.Address, recipient []common.Address) (*CollateralVaultCollateralSeizedIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "CollateralSeized", borrowerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultCollateralSeizedIterator{contract: _CollateralVault.contract, event: "CollateralSeized", logs: logs, sub: sub}, nil
}

// WatchCollateralSeized is a free log subscription operation binding the contract event 0x2f72539f26ad4e48fa396e30d602a0397ca67099e2d58bd39663a7bc09944ee3.
//
// Solidity: event CollateralSeized(address indexed borrower, address indexed recipient, uint256 amount, uint256 newCollateral)
func (_CollateralVault *CollateralVaultFilterer) WatchCollateralSeized(opts *bind.WatchOpts, sink chan<- *CollateralVaultCollateralSeized, borrower []common.Address, recipient []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "CollateralSeized", borrowerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultCollateralSeized)
				if err := _CollateralVault.contract.UnpackLog(event, "CollateralSeized", log); err != nil {
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

// ParseCollateralSeized is a log parse operation binding the contract event 0x2f72539f26ad4e48fa396e30d602a0397ca67099e2d58bd39663a7bc09944ee3.
//
// Solidity: event CollateralSeized(address indexed borrower, address indexed recipient, uint256 amount, uint256 newCollateral)
func (_CollateralVault *CollateralVaultFilterer) ParseCollateralSeized(log types.Log) (*CollateralVaultCollateralSeized, error) {
	event := new(CollateralVaultCollateralSeized)
	if err := _CollateralVault.contract.UnpackLog(event, "CollateralSeized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultCollateralWithdrawnIterator is returned from FilterCollateralWithdrawn and is used to iterate over the raw logs and unpacked data for CollateralWithdrawn events raised by the CollateralVault contract.
type CollateralVaultCollateralWithdrawnIterator struct {
	Event *CollateralVaultCollateralWithdrawn // Event containing the contract specifics and raw log

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
func (it *CollateralVaultCollateralWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultCollateralWithdrawn)
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
		it.Event = new(CollateralVaultCollateralWithdrawn)
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
func (it *CollateralVaultCollateralWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultCollateralWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultCollateralWithdrawn represents a CollateralWithdrawn event raised by the CollateralVault contract.
type CollateralVaultCollateralWithdrawn struct {
	Borrower      common.Address
	Amount        *big.Int
	NewCollateral *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollateralWithdrawn is a free log retrieval operation binding the contract event 0xdaed309a628faec6cab72194019e2a1a34e890ca9bf9be99788992dd54692819.
//
// Solidity: event CollateralWithdrawn(address indexed borrower, uint256 amount, uint256 newCollateral)
func (_CollateralVault *CollateralVaultFilterer) FilterCollateralWithdrawn(opts *bind.FilterOpts, borrower []common.Address) (*CollateralVaultCollateralWithdrawnIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "CollateralWithdrawn", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultCollateralWithdrawnIterator{contract: _CollateralVault.contract, event: "CollateralWithdrawn", logs: logs, sub: sub}, nil
}

// WatchCollateralWithdrawn is a free log subscription operation binding the contract event 0xdaed309a628faec6cab72194019e2a1a34e890ca9bf9be99788992dd54692819.
//
// Solidity: event CollateralWithdrawn(address indexed borrower, uint256 amount, uint256 newCollateral)
func (_CollateralVault *CollateralVaultFilterer) WatchCollateralWithdrawn(opts *bind.WatchOpts, sink chan<- *CollateralVaultCollateralWithdrawn, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "CollateralWithdrawn", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultCollateralWithdrawn)
				if err := _CollateralVault.contract.UnpackLog(event, "CollateralWithdrawn", log); err != nil {
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

// ParseCollateralWithdrawn is a log parse operation binding the contract event 0xdaed309a628faec6cab72194019e2a1a34e890ca9bf9be99788992dd54692819.
//
// Solidity: event CollateralWithdrawn(address indexed borrower, uint256 amount, uint256 newCollateral)
func (_CollateralVault *CollateralVaultFilterer) ParseCollateralWithdrawn(log types.Log) (*CollateralVaultCollateralWithdrawn, error) {
	event := new(CollateralVaultCollateralWithdrawn)
	if err := _CollateralVault.contract.UnpackLog(event, "CollateralWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultControllerLinkedIterator is returned from FilterControllerLinked and is used to iterate over the raw logs and unpacked data for ControllerLinked events raised by the CollateralVault contract.
type CollateralVaultControllerLinkedIterator struct {
	Event *CollateralVaultControllerLinked // Event containing the contract specifics and raw log

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
func (it *CollateralVaultControllerLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultControllerLinked)
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
		it.Event = new(CollateralVaultControllerLinked)
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
func (it *CollateralVaultControllerLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultControllerLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultControllerLinked represents a ControllerLinked event raised by the CollateralVault contract.
type CollateralVaultControllerLinked struct {
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterControllerLinked is a free log retrieval operation binding the contract event 0xa76ae82e550122379c3a3784cfc195ef6f1bae7f61abf8fae42a47725830c6c8.
//
// Solidity: event ControllerLinked(address controller)
func (_CollateralVault *CollateralVaultFilterer) FilterControllerLinked(opts *bind.FilterOpts) (*CollateralVaultControllerLinkedIterator, error) {

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "ControllerLinked")
	if err != nil {
		return nil, err
	}
	return &CollateralVaultControllerLinkedIterator{contract: _CollateralVault.contract, event: "ControllerLinked", logs: logs, sub: sub}, nil
}

// WatchControllerLinked is a free log subscription operation binding the contract event 0xa76ae82e550122379c3a3784cfc195ef6f1bae7f61abf8fae42a47725830c6c8.
//
// Solidity: event ControllerLinked(address controller)
func (_CollateralVault *CollateralVaultFilterer) WatchControllerLinked(opts *bind.WatchOpts, sink chan<- *CollateralVaultControllerLinked) (event.Subscription, error) {

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "ControllerLinked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultControllerLinked)
				if err := _CollateralVault.contract.UnpackLog(event, "ControllerLinked", log); err != nil {
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
func (_CollateralVault *CollateralVaultFilterer) ParseControllerLinked(log types.Log) (*CollateralVaultControllerLinked, error) {
	event := new(CollateralVaultControllerLinked)
	if err := _CollateralVault.contract.UnpackLog(event, "ControllerLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultLiquidationManagerLinkedIterator is returned from FilterLiquidationManagerLinked and is used to iterate over the raw logs and unpacked data for LiquidationManagerLinked events raised by the CollateralVault contract.
type CollateralVaultLiquidationManagerLinkedIterator struct {
	Event *CollateralVaultLiquidationManagerLinked // Event containing the contract specifics and raw log

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
func (it *CollateralVaultLiquidationManagerLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultLiquidationManagerLinked)
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
		it.Event = new(CollateralVaultLiquidationManagerLinked)
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
func (it *CollateralVaultLiquidationManagerLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultLiquidationManagerLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultLiquidationManagerLinked represents a LiquidationManagerLinked event raised by the CollateralVault contract.
type CollateralVaultLiquidationManagerLinked struct {
	LiquidationManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLiquidationManagerLinked is a free log retrieval operation binding the contract event 0x40fd76f04a4f105bfd0d978d7bc92e40784482098dc21077e350a7c121dfba1d.
//
// Solidity: event LiquidationManagerLinked(address liquidationManager)
func (_CollateralVault *CollateralVaultFilterer) FilterLiquidationManagerLinked(opts *bind.FilterOpts) (*CollateralVaultLiquidationManagerLinkedIterator, error) {

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "LiquidationManagerLinked")
	if err != nil {
		return nil, err
	}
	return &CollateralVaultLiquidationManagerLinkedIterator{contract: _CollateralVault.contract, event: "LiquidationManagerLinked", logs: logs, sub: sub}, nil
}

// WatchLiquidationManagerLinked is a free log subscription operation binding the contract event 0x40fd76f04a4f105bfd0d978d7bc92e40784482098dc21077e350a7c121dfba1d.
//
// Solidity: event LiquidationManagerLinked(address liquidationManager)
func (_CollateralVault *CollateralVaultFilterer) WatchLiquidationManagerLinked(opts *bind.WatchOpts, sink chan<- *CollateralVaultLiquidationManagerLinked) (event.Subscription, error) {

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "LiquidationManagerLinked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultLiquidationManagerLinked)
				if err := _CollateralVault.contract.UnpackLog(event, "LiquidationManagerLinked", log); err != nil {
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
func (_CollateralVault *CollateralVaultFilterer) ParseLiquidationManagerLinked(log types.Log) (*CollateralVaultLiquidationManagerLinked, error) {
	event := new(CollateralVaultLiquidationManagerLinked)
	if err := _CollateralVault.contract.UnpackLog(event, "LiquidationManagerLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CollateralVault contract.
type CollateralVaultOwnershipTransferredIterator struct {
	Event *CollateralVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CollateralVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralVaultOwnershipTransferred)
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
		it.Event = new(CollateralVaultOwnershipTransferred)
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
func (it *CollateralVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralVaultOwnershipTransferred represents a OwnershipTransferred event raised by the CollateralVault contract.
type CollateralVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CollateralVault *CollateralVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CollateralVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CollateralVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CollateralVaultOwnershipTransferredIterator{contract: _CollateralVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CollateralVault *CollateralVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CollateralVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CollateralVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralVaultOwnershipTransferred)
				if err := _CollateralVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CollateralVault *CollateralVaultFilterer) ParseOwnershipTransferred(log types.Log) (*CollateralVaultOwnershipTransferred, error) {
	event := new(CollateralVaultOwnershipTransferred)
	if err := _CollateralVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
