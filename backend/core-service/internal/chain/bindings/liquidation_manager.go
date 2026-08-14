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

// LiquidationManagerMetaData contains all meta data concerning the LiquidationManager contract.
var LiquidationManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"lendingPool\",\"type\":\"address\",\"internalType\":\"contractILendingPool\"},{\"name\":\"collateralVault\",\"type\":\"address\",\"internalType\":\"contractICollateralVault\"},{\"name\":\"priceOracle\",\"type\":\"address\",\"internalType\":\"contractIPriceOracle\"},{\"name\":\"lendingController\",\"type\":\"address\",\"internalType\":\"contractILendingController\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collateralAsset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"controller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILendingController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"debtAsset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"debtDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLiquidatable\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidate\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPriceOracle\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILendingPool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewLiquidation\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"debtToRepay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralToSeize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bonusValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shortfall\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICollateralVault\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"LiquidationExecuted\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"liquidator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"debtRepaid\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateralSeized\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bonusValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"healthFactorBeforeBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateralPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"priceDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"shortfall\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"PositionIsHealthy\",\"inputs\":[{\"name\":\"healthFactorBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// LiquidationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidationManagerMetaData.ABI instead.
var LiquidationManagerABI = LiquidationManagerMetaData.ABI

// LiquidationManager is an auto generated Go binding around an Ethereum contract.
type LiquidationManager struct {
	LiquidationManagerCaller     // Read-only binding to the contract
	LiquidationManagerTransactor // Write-only binding to the contract
	LiquidationManagerFilterer   // Log filterer for contract events
}

// LiquidationManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidationManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidationManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidationManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidationManagerSession struct {
	Contract     *LiquidationManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LiquidationManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidationManagerCallerSession struct {
	Contract *LiquidationManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// LiquidationManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidationManagerTransactorSession struct {
	Contract     *LiquidationManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// LiquidationManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidationManagerRaw struct {
	Contract *LiquidationManager // Generic contract binding to access the raw methods on
}

// LiquidationManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidationManagerCallerRaw struct {
	Contract *LiquidationManagerCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidationManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidationManagerTransactorRaw struct {
	Contract *LiquidationManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidationManager creates a new instance of LiquidationManager, bound to a specific deployed contract.
func NewLiquidationManager(address common.Address, backend bind.ContractBackend) (*LiquidationManager, error) {
	contract, err := bindLiquidationManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidationManager{LiquidationManagerCaller: LiquidationManagerCaller{contract: contract}, LiquidationManagerTransactor: LiquidationManagerTransactor{contract: contract}, LiquidationManagerFilterer: LiquidationManagerFilterer{contract: contract}}, nil
}

// NewLiquidationManagerCaller creates a new read-only instance of LiquidationManager, bound to a specific deployed contract.
func NewLiquidationManagerCaller(address common.Address, caller bind.ContractCaller) (*LiquidationManagerCaller, error) {
	contract, err := bindLiquidationManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidationManagerCaller{contract: contract}, nil
}

// NewLiquidationManagerTransactor creates a new write-only instance of LiquidationManager, bound to a specific deployed contract.
func NewLiquidationManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidationManagerTransactor, error) {
	contract, err := bindLiquidationManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidationManagerTransactor{contract: contract}, nil
}

// NewLiquidationManagerFilterer creates a new log filterer instance of LiquidationManager, bound to a specific deployed contract.
func NewLiquidationManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidationManagerFilterer, error) {
	contract, err := bindLiquidationManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidationManagerFilterer{contract: contract}, nil
}

// bindLiquidationManager binds a generic wrapper to an already deployed contract.
func bindLiquidationManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiquidationManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidationManager *LiquidationManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidationManager.Contract.LiquidationManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidationManager *LiquidationManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationManager.Contract.LiquidationManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidationManager *LiquidationManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidationManager.Contract.LiquidationManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidationManager *LiquidationManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidationManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidationManager *LiquidationManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidationManager *LiquidationManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidationManager.Contract.contract.Transact(opts, method, params...)
}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_LiquidationManager *LiquidationManagerCaller) CollateralAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationManager.contract.Call(opts, &out, "collateralAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_LiquidationManager *LiquidationManagerSession) CollateralAsset() (common.Address, error) {
	return _LiquidationManager.Contract.CollateralAsset(&_LiquidationManager.CallOpts)
}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_LiquidationManager *LiquidationManagerCallerSession) CollateralAsset() (common.Address, error) {
	return _LiquidationManager.Contract.CollateralAsset(&_LiquidationManager.CallOpts)
}

// CollateralDecimals is a free data retrieval call binding the contract method 0xec9c6c30.
//
// Solidity: function collateralDecimals() view returns(uint8)
func (_LiquidationManager *LiquidationManagerCaller) CollateralDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LiquidationManager.contract.Call(opts, &out, "collateralDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CollateralDecimals is a free data retrieval call binding the contract method 0xec9c6c30.
//
// Solidity: function collateralDecimals() view returns(uint8)
func (_LiquidationManager *LiquidationManagerSession) CollateralDecimals() (uint8, error) {
	return _LiquidationManager.Contract.CollateralDecimals(&_LiquidationManager.CallOpts)
}

// CollateralDecimals is a free data retrieval call binding the contract method 0xec9c6c30.
//
// Solidity: function collateralDecimals() view returns(uint8)
func (_LiquidationManager *LiquidationManagerCallerSession) CollateralDecimals() (uint8, error) {
	return _LiquidationManager.Contract.CollateralDecimals(&_LiquidationManager.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_LiquidationManager *LiquidationManagerCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationManager.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_LiquidationManager *LiquidationManagerSession) Controller() (common.Address, error) {
	return _LiquidationManager.Contract.Controller(&_LiquidationManager.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_LiquidationManager *LiquidationManagerCallerSession) Controller() (common.Address, error) {
	return _LiquidationManager.Contract.Controller(&_LiquidationManager.CallOpts)
}

// DebtAsset is a free data retrieval call binding the contract method 0xa919802d.
//
// Solidity: function debtAsset() view returns(address)
func (_LiquidationManager *LiquidationManagerCaller) DebtAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationManager.contract.Call(opts, &out, "debtAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DebtAsset is a free data retrieval call binding the contract method 0xa919802d.
//
// Solidity: function debtAsset() view returns(address)
func (_LiquidationManager *LiquidationManagerSession) DebtAsset() (common.Address, error) {
	return _LiquidationManager.Contract.DebtAsset(&_LiquidationManager.CallOpts)
}

// DebtAsset is a free data retrieval call binding the contract method 0xa919802d.
//
// Solidity: function debtAsset() view returns(address)
func (_LiquidationManager *LiquidationManagerCallerSession) DebtAsset() (common.Address, error) {
	return _LiquidationManager.Contract.DebtAsset(&_LiquidationManager.CallOpts)
}

// DebtDecimals is a free data retrieval call binding the contract method 0xf341e228.
//
// Solidity: function debtDecimals() view returns(uint8)
func (_LiquidationManager *LiquidationManagerCaller) DebtDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LiquidationManager.contract.Call(opts, &out, "debtDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DebtDecimals is a free data retrieval call binding the contract method 0xf341e228.
//
// Solidity: function debtDecimals() view returns(uint8)
func (_LiquidationManager *LiquidationManagerSession) DebtDecimals() (uint8, error) {
	return _LiquidationManager.Contract.DebtDecimals(&_LiquidationManager.CallOpts)
}

// DebtDecimals is a free data retrieval call binding the contract method 0xf341e228.
//
// Solidity: function debtDecimals() view returns(uint8)
func (_LiquidationManager *LiquidationManagerCallerSession) DebtDecimals() (uint8, error) {
	return _LiquidationManager.Contract.DebtDecimals(&_LiquidationManager.CallOpts)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address borrower) view returns(bool)
func (_LiquidationManager *LiquidationManagerCaller) IsLiquidatable(opts *bind.CallOpts, borrower common.Address) (bool, error) {
	var out []interface{}
	err := _LiquidationManager.contract.Call(opts, &out, "isLiquidatable", borrower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address borrower) view returns(bool)
func (_LiquidationManager *LiquidationManagerSession) IsLiquidatable(borrower common.Address) (bool, error) {
	return _LiquidationManager.Contract.IsLiquidatable(&_LiquidationManager.CallOpts, borrower)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address borrower) view returns(bool)
func (_LiquidationManager *LiquidationManagerCallerSession) IsLiquidatable(borrower common.Address) (bool, error) {
	return _LiquidationManager.Contract.IsLiquidatable(&_LiquidationManager.CallOpts, borrower)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_LiquidationManager *LiquidationManagerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationManager.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_LiquidationManager *LiquidationManagerSession) Oracle() (common.Address, error) {
	return _LiquidationManager.Contract.Oracle(&_LiquidationManager.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_LiquidationManager *LiquidationManagerCallerSession) Oracle() (common.Address, error) {
	return _LiquidationManager.Contract.Oracle(&_LiquidationManager.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_LiquidationManager *LiquidationManagerCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationManager.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_LiquidationManager *LiquidationManagerSession) Pool() (common.Address, error) {
	return _LiquidationManager.Contract.Pool(&_LiquidationManager.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_LiquidationManager *LiquidationManagerCallerSession) Pool() (common.Address, error) {
	return _LiquidationManager.Contract.Pool(&_LiquidationManager.CallOpts)
}

// PreviewLiquidation is a free data retrieval call binding the contract method 0xb37518cd.
//
// Solidity: function previewLiquidation(address borrower) view returns(uint256 debtToRepay, uint256 collateralToSeize, uint256 bonusValue, uint256 shortfall)
func (_LiquidationManager *LiquidationManagerCaller) PreviewLiquidation(opts *bind.CallOpts, borrower common.Address) (struct {
	DebtToRepay       *big.Int
	CollateralToSeize *big.Int
	BonusValue        *big.Int
	Shortfall         *big.Int
}, error) {
	var out []interface{}
	err := _LiquidationManager.contract.Call(opts, &out, "previewLiquidation", borrower)

	outstruct := new(struct {
		DebtToRepay       *big.Int
		CollateralToSeize *big.Int
		BonusValue        *big.Int
		Shortfall         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DebtToRepay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CollateralToSeize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BonusValue = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Shortfall = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PreviewLiquidation is a free data retrieval call binding the contract method 0xb37518cd.
//
// Solidity: function previewLiquidation(address borrower) view returns(uint256 debtToRepay, uint256 collateralToSeize, uint256 bonusValue, uint256 shortfall)
func (_LiquidationManager *LiquidationManagerSession) PreviewLiquidation(borrower common.Address) (struct {
	DebtToRepay       *big.Int
	CollateralToSeize *big.Int
	BonusValue        *big.Int
	Shortfall         *big.Int
}, error) {
	return _LiquidationManager.Contract.PreviewLiquidation(&_LiquidationManager.CallOpts, borrower)
}

// PreviewLiquidation is a free data retrieval call binding the contract method 0xb37518cd.
//
// Solidity: function previewLiquidation(address borrower) view returns(uint256 debtToRepay, uint256 collateralToSeize, uint256 bonusValue, uint256 shortfall)
func (_LiquidationManager *LiquidationManagerCallerSession) PreviewLiquidation(borrower common.Address) (struct {
	DebtToRepay       *big.Int
	CollateralToSeize *big.Int
	BonusValue        *big.Int
	Shortfall         *big.Int
}, error) {
	return _LiquidationManager.Contract.PreviewLiquidation(&_LiquidationManager.CallOpts, borrower)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_LiquidationManager *LiquidationManagerCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationManager.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_LiquidationManager *LiquidationManagerSession) Vault() (common.Address, error) {
	return _LiquidationManager.Contract.Vault(&_LiquidationManager.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_LiquidationManager *LiquidationManagerCallerSession) Vault() (common.Address, error) {
	return _LiquidationManager.Contract.Vault(&_LiquidationManager.CallOpts)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address borrower) returns()
func (_LiquidationManager *LiquidationManagerTransactor) Liquidate(opts *bind.TransactOpts, borrower common.Address) (*types.Transaction, error) {
	return _LiquidationManager.contract.Transact(opts, "liquidate", borrower)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address borrower) returns()
func (_LiquidationManager *LiquidationManagerSession) Liquidate(borrower common.Address) (*types.Transaction, error) {
	return _LiquidationManager.Contract.Liquidate(&_LiquidationManager.TransactOpts, borrower)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address borrower) returns()
func (_LiquidationManager *LiquidationManagerTransactorSession) Liquidate(borrower common.Address) (*types.Transaction, error) {
	return _LiquidationManager.Contract.Liquidate(&_LiquidationManager.TransactOpts, borrower)
}

// LiquidationManagerLiquidationExecutedIterator is returned from FilterLiquidationExecuted and is used to iterate over the raw logs and unpacked data for LiquidationExecuted events raised by the LiquidationManager contract.
type LiquidationManagerLiquidationExecutedIterator struct {
	Event *LiquidationManagerLiquidationExecuted // Event containing the contract specifics and raw log

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
func (it *LiquidationManagerLiquidationExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidationManagerLiquidationExecuted)
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
		it.Event = new(LiquidationManagerLiquidationExecuted)
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
func (it *LiquidationManagerLiquidationExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidationManagerLiquidationExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidationManagerLiquidationExecuted represents a LiquidationExecuted event raised by the LiquidationManager contract.
type LiquidationManagerLiquidationExecuted struct {
	Borrower              common.Address
	Liquidator            common.Address
	DebtRepaid            *big.Int
	CollateralSeized      *big.Int
	BonusValue            *big.Int
	HealthFactorBeforeBps *big.Int
	CollateralPrice       *big.Int
	PriceDecimals         uint8
	Shortfall             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterLiquidationExecuted is a free log retrieval operation binding the contract event 0x5bbac90bb94f42d17e2cb73e924f770f18c7ec2206d3fc6e09b9d7cbb7b53c86.
//
// Solidity: event LiquidationExecuted(address indexed borrower, address indexed liquidator, uint256 debtRepaid, uint256 collateralSeized, uint256 bonusValue, uint256 healthFactorBeforeBps, uint256 collateralPrice, uint8 priceDecimals, uint256 shortfall)
func (_LiquidationManager *LiquidationManagerFilterer) FilterLiquidationExecuted(opts *bind.FilterOpts, borrower []common.Address, liquidator []common.Address) (*LiquidationManagerLiquidationExecutedIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _LiquidationManager.contract.FilterLogs(opts, "LiquidationExecuted", borrowerRule, liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &LiquidationManagerLiquidationExecutedIterator{contract: _LiquidationManager.contract, event: "LiquidationExecuted", logs: logs, sub: sub}, nil
}

// WatchLiquidationExecuted is a free log subscription operation binding the contract event 0x5bbac90bb94f42d17e2cb73e924f770f18c7ec2206d3fc6e09b9d7cbb7b53c86.
//
// Solidity: event LiquidationExecuted(address indexed borrower, address indexed liquidator, uint256 debtRepaid, uint256 collateralSeized, uint256 bonusValue, uint256 healthFactorBeforeBps, uint256 collateralPrice, uint8 priceDecimals, uint256 shortfall)
func (_LiquidationManager *LiquidationManagerFilterer) WatchLiquidationExecuted(opts *bind.WatchOpts, sink chan<- *LiquidationManagerLiquidationExecuted, borrower []common.Address, liquidator []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _LiquidationManager.contract.WatchLogs(opts, "LiquidationExecuted", borrowerRule, liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidationManagerLiquidationExecuted)
				if err := _LiquidationManager.contract.UnpackLog(event, "LiquidationExecuted", log); err != nil {
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

// ParseLiquidationExecuted is a log parse operation binding the contract event 0x5bbac90bb94f42d17e2cb73e924f770f18c7ec2206d3fc6e09b9d7cbb7b53c86.
//
// Solidity: event LiquidationExecuted(address indexed borrower, address indexed liquidator, uint256 debtRepaid, uint256 collateralSeized, uint256 bonusValue, uint256 healthFactorBeforeBps, uint256 collateralPrice, uint8 priceDecimals, uint256 shortfall)
func (_LiquidationManager *LiquidationManagerFilterer) ParseLiquidationExecuted(log types.Log) (*LiquidationManagerLiquidationExecuted, error) {
	event := new(LiquidationManagerLiquidationExecuted)
	if err := _LiquidationManager.contract.UnpackLog(event, "LiquidationExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
