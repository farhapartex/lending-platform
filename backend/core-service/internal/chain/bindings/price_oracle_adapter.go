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

// PriceData is an auto generated low-level Go binding around an user-defined struct.
type PriceData struct {
	Price     *big.Int
	Decimals  uint8
	UpdatedAt *big.Int
	IsStale   bool
	IsValid   bool
}

// PriceOracleAdapterMetaData contains all meta data concerning the PriceOracleAdapter contract.
var PriceOracleAdapterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"startingMaxPriceAge\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"REQUIRED_FEED_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feedOf\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPrice\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isStale\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxPriceAge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readPrice\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPriceData\",\"components\":[{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isStale\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isValid\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeed\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxPriceAge\",\"inputs\":[{\"name\":\"newAge\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeedChanged\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"aggregator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxPriceAgeChanged\",\"inputs\":[{\"name\":\"previousAge\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newAge\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidRiskSettings\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PriceIsInvalid\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PriceIsStale\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAge\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedIntToUint\",\"inputs\":[{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"type\":\"error\",\"name\":\"UnsupportedFeedDecimals\",\"inputs\":[{\"name\":\"aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"provided\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// PriceOracleAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceOracleAdapterMetaData.ABI instead.
var PriceOracleAdapterABI = PriceOracleAdapterMetaData.ABI

// PriceOracleAdapter is an auto generated Go binding around an Ethereum contract.
type PriceOracleAdapter struct {
	PriceOracleAdapterCaller     // Read-only binding to the contract
	PriceOracleAdapterTransactor // Write-only binding to the contract
	PriceOracleAdapterFilterer   // Log filterer for contract events
}

// PriceOracleAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceOracleAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceOracleAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceOracleAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceOracleAdapterSession struct {
	Contract     *PriceOracleAdapter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PriceOracleAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceOracleAdapterCallerSession struct {
	Contract *PriceOracleAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PriceOracleAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceOracleAdapterTransactorSession struct {
	Contract     *PriceOracleAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PriceOracleAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceOracleAdapterRaw struct {
	Contract *PriceOracleAdapter // Generic contract binding to access the raw methods on
}

// PriceOracleAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceOracleAdapterCallerRaw struct {
	Contract *PriceOracleAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// PriceOracleAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceOracleAdapterTransactorRaw struct {
	Contract *PriceOracleAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceOracleAdapter creates a new instance of PriceOracleAdapter, bound to a specific deployed contract.
func NewPriceOracleAdapter(address common.Address, backend bind.ContractBackend) (*PriceOracleAdapter, error) {
	contract, err := bindPriceOracleAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceOracleAdapter{PriceOracleAdapterCaller: PriceOracleAdapterCaller{contract: contract}, PriceOracleAdapterTransactor: PriceOracleAdapterTransactor{contract: contract}, PriceOracleAdapterFilterer: PriceOracleAdapterFilterer{contract: contract}}, nil
}

// NewPriceOracleAdapterCaller creates a new read-only instance of PriceOracleAdapter, bound to a specific deployed contract.
func NewPriceOracleAdapterCaller(address common.Address, caller bind.ContractCaller) (*PriceOracleAdapterCaller, error) {
	contract, err := bindPriceOracleAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOracleAdapterCaller{contract: contract}, nil
}

// NewPriceOracleAdapterTransactor creates a new write-only instance of PriceOracleAdapter, bound to a specific deployed contract.
func NewPriceOracleAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceOracleAdapterTransactor, error) {
	contract, err := bindPriceOracleAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOracleAdapterTransactor{contract: contract}, nil
}

// NewPriceOracleAdapterFilterer creates a new log filterer instance of PriceOracleAdapter, bound to a specific deployed contract.
func NewPriceOracleAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceOracleAdapterFilterer, error) {
	contract, err := bindPriceOracleAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceOracleAdapterFilterer{contract: contract}, nil
}

// bindPriceOracleAdapter binds a generic wrapper to an already deployed contract.
func bindPriceOracleAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceOracleAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOracleAdapter *PriceOracleAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceOracleAdapter.Contract.PriceOracleAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOracleAdapter *PriceOracleAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracleAdapter.Contract.PriceOracleAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOracleAdapter *PriceOracleAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOracleAdapter.Contract.PriceOracleAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOracleAdapter *PriceOracleAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceOracleAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOracleAdapter *PriceOracleAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracleAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOracleAdapter *PriceOracleAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOracleAdapter.Contract.contract.Transact(opts, method, params...)
}

// REQUIREDFEEDDECIMALS is a free data retrieval call binding the contract method 0x347af4dd.
//
// Solidity: function REQUIRED_FEED_DECIMALS() view returns(uint8)
func (_PriceOracleAdapter *PriceOracleAdapterCaller) REQUIREDFEEDDECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PriceOracleAdapter.contract.Call(opts, &out, "REQUIRED_FEED_DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// REQUIREDFEEDDECIMALS is a free data retrieval call binding the contract method 0x347af4dd.
//
// Solidity: function REQUIRED_FEED_DECIMALS() view returns(uint8)
func (_PriceOracleAdapter *PriceOracleAdapterSession) REQUIREDFEEDDECIMALS() (uint8, error) {
	return _PriceOracleAdapter.Contract.REQUIREDFEEDDECIMALS(&_PriceOracleAdapter.CallOpts)
}

// REQUIREDFEEDDECIMALS is a free data retrieval call binding the contract method 0x347af4dd.
//
// Solidity: function REQUIRED_FEED_DECIMALS() view returns(uint8)
func (_PriceOracleAdapter *PriceOracleAdapterCallerSession) REQUIREDFEEDDECIMALS() (uint8, error) {
	return _PriceOracleAdapter.Contract.REQUIREDFEEDDECIMALS(&_PriceOracleAdapter.CallOpts)
}

// FeedOf is a free data retrieval call binding the contract method 0x4b45e8a6.
//
// Solidity: function feedOf(address asset) view returns(address)
func (_PriceOracleAdapter *PriceOracleAdapterCaller) FeedOf(opts *bind.CallOpts, asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _PriceOracleAdapter.contract.Call(opts, &out, "feedOf", asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeedOf is a free data retrieval call binding the contract method 0x4b45e8a6.
//
// Solidity: function feedOf(address asset) view returns(address)
func (_PriceOracleAdapter *PriceOracleAdapterSession) FeedOf(asset common.Address) (common.Address, error) {
	return _PriceOracleAdapter.Contract.FeedOf(&_PriceOracleAdapter.CallOpts, asset)
}

// FeedOf is a free data retrieval call binding the contract method 0x4b45e8a6.
//
// Solidity: function feedOf(address asset) view returns(address)
func (_PriceOracleAdapter *PriceOracleAdapterCallerSession) FeedOf(asset common.Address) (common.Address, error) {
	return _PriceOracleAdapter.Contract.FeedOf(&_PriceOracleAdapter.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256 price, uint8 decimals)
func (_PriceOracleAdapter *PriceOracleAdapterCaller) GetPrice(opts *bind.CallOpts, asset common.Address) (struct {
	Price    *big.Int
	Decimals uint8
}, error) {
	var out []interface{}
	err := _PriceOracleAdapter.contract.Call(opts, &out, "getPrice", asset)

	outstruct := new(struct {
		Price    *big.Int
		Decimals uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Decimals = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256 price, uint8 decimals)
func (_PriceOracleAdapter *PriceOracleAdapterSession) GetPrice(asset common.Address) (struct {
	Price    *big.Int
	Decimals uint8
}, error) {
	return _PriceOracleAdapter.Contract.GetPrice(&_PriceOracleAdapter.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256 price, uint8 decimals)
func (_PriceOracleAdapter *PriceOracleAdapterCallerSession) GetPrice(asset common.Address) (struct {
	Price    *big.Int
	Decimals uint8
}, error) {
	return _PriceOracleAdapter.Contract.GetPrice(&_PriceOracleAdapter.CallOpts, asset)
}

// IsStale is a free data retrieval call binding the contract method 0xf461f6e7.
//
// Solidity: function isStale(address asset) view returns(bool)
func (_PriceOracleAdapter *PriceOracleAdapterCaller) IsStale(opts *bind.CallOpts, asset common.Address) (bool, error) {
	var out []interface{}
	err := _PriceOracleAdapter.contract.Call(opts, &out, "isStale", asset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStale is a free data retrieval call binding the contract method 0xf461f6e7.
//
// Solidity: function isStale(address asset) view returns(bool)
func (_PriceOracleAdapter *PriceOracleAdapterSession) IsStale(asset common.Address) (bool, error) {
	return _PriceOracleAdapter.Contract.IsStale(&_PriceOracleAdapter.CallOpts, asset)
}

// IsStale is a free data retrieval call binding the contract method 0xf461f6e7.
//
// Solidity: function isStale(address asset) view returns(bool)
func (_PriceOracleAdapter *PriceOracleAdapterCallerSession) IsStale(asset common.Address) (bool, error) {
	return _PriceOracleAdapter.Contract.IsStale(&_PriceOracleAdapter.CallOpts, asset)
}

// MaxPriceAge is a free data retrieval call binding the contract method 0x1584410a.
//
// Solidity: function maxPriceAge() view returns(uint32)
func (_PriceOracleAdapter *PriceOracleAdapterCaller) MaxPriceAge(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PriceOracleAdapter.contract.Call(opts, &out, "maxPriceAge")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxPriceAge is a free data retrieval call binding the contract method 0x1584410a.
//
// Solidity: function maxPriceAge() view returns(uint32)
func (_PriceOracleAdapter *PriceOracleAdapterSession) MaxPriceAge() (uint32, error) {
	return _PriceOracleAdapter.Contract.MaxPriceAge(&_PriceOracleAdapter.CallOpts)
}

// MaxPriceAge is a free data retrieval call binding the contract method 0x1584410a.
//
// Solidity: function maxPriceAge() view returns(uint32)
func (_PriceOracleAdapter *PriceOracleAdapterCallerSession) MaxPriceAge() (uint32, error) {
	return _PriceOracleAdapter.Contract.MaxPriceAge(&_PriceOracleAdapter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceOracleAdapter *PriceOracleAdapterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceOracleAdapter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceOracleAdapter *PriceOracleAdapterSession) Owner() (common.Address, error) {
	return _PriceOracleAdapter.Contract.Owner(&_PriceOracleAdapter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceOracleAdapter *PriceOracleAdapterCallerSession) Owner() (common.Address, error) {
	return _PriceOracleAdapter.Contract.Owner(&_PriceOracleAdapter.CallOpts)
}

// ReadPrice is a free data retrieval call binding the contract method 0x033d497a.
//
// Solidity: function readPrice(address asset) view returns((uint256,uint8,uint256,bool,bool))
func (_PriceOracleAdapter *PriceOracleAdapterCaller) ReadPrice(opts *bind.CallOpts, asset common.Address) (PriceData, error) {
	var out []interface{}
	err := _PriceOracleAdapter.contract.Call(opts, &out, "readPrice", asset)

	if err != nil {
		return *new(PriceData), err
	}

	out0 := *abi.ConvertType(out[0], new(PriceData)).(*PriceData)

	return out0, err

}

// ReadPrice is a free data retrieval call binding the contract method 0x033d497a.
//
// Solidity: function readPrice(address asset) view returns((uint256,uint8,uint256,bool,bool))
func (_PriceOracleAdapter *PriceOracleAdapterSession) ReadPrice(asset common.Address) (PriceData, error) {
	return _PriceOracleAdapter.Contract.ReadPrice(&_PriceOracleAdapter.CallOpts, asset)
}

// ReadPrice is a free data retrieval call binding the contract method 0x033d497a.
//
// Solidity: function readPrice(address asset) view returns((uint256,uint8,uint256,bool,bool))
func (_PriceOracleAdapter *PriceOracleAdapterCallerSession) ReadPrice(asset common.Address) (PriceData, error) {
	return _PriceOracleAdapter.Contract.ReadPrice(&_PriceOracleAdapter.CallOpts, asset)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceOracleAdapter *PriceOracleAdapterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracleAdapter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceOracleAdapter *PriceOracleAdapterSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceOracleAdapter.Contract.RenounceOwnership(&_PriceOracleAdapter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceOracleAdapter *PriceOracleAdapterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceOracleAdapter.Contract.RenounceOwnership(&_PriceOracleAdapter.TransactOpts)
}

// SetFeed is a paid mutator transaction binding the contract method 0x40b1eb10.
//
// Solidity: function setFeed(address asset, address aggregator) returns()
func (_PriceOracleAdapter *PriceOracleAdapterTransactor) SetFeed(opts *bind.TransactOpts, asset common.Address, aggregator common.Address) (*types.Transaction, error) {
	return _PriceOracleAdapter.contract.Transact(opts, "setFeed", asset, aggregator)
}

// SetFeed is a paid mutator transaction binding the contract method 0x40b1eb10.
//
// Solidity: function setFeed(address asset, address aggregator) returns()
func (_PriceOracleAdapter *PriceOracleAdapterSession) SetFeed(asset common.Address, aggregator common.Address) (*types.Transaction, error) {
	return _PriceOracleAdapter.Contract.SetFeed(&_PriceOracleAdapter.TransactOpts, asset, aggregator)
}

// SetFeed is a paid mutator transaction binding the contract method 0x40b1eb10.
//
// Solidity: function setFeed(address asset, address aggregator) returns()
func (_PriceOracleAdapter *PriceOracleAdapterTransactorSession) SetFeed(asset common.Address, aggregator common.Address) (*types.Transaction, error) {
	return _PriceOracleAdapter.Contract.SetFeed(&_PriceOracleAdapter.TransactOpts, asset, aggregator)
}

// SetMaxPriceAge is a paid mutator transaction binding the contract method 0x6b64c46f.
//
// Solidity: function setMaxPriceAge(uint32 newAge) returns()
func (_PriceOracleAdapter *PriceOracleAdapterTransactor) SetMaxPriceAge(opts *bind.TransactOpts, newAge uint32) (*types.Transaction, error) {
	return _PriceOracleAdapter.contract.Transact(opts, "setMaxPriceAge", newAge)
}

// SetMaxPriceAge is a paid mutator transaction binding the contract method 0x6b64c46f.
//
// Solidity: function setMaxPriceAge(uint32 newAge) returns()
func (_PriceOracleAdapter *PriceOracleAdapterSession) SetMaxPriceAge(newAge uint32) (*types.Transaction, error) {
	return _PriceOracleAdapter.Contract.SetMaxPriceAge(&_PriceOracleAdapter.TransactOpts, newAge)
}

// SetMaxPriceAge is a paid mutator transaction binding the contract method 0x6b64c46f.
//
// Solidity: function setMaxPriceAge(uint32 newAge) returns()
func (_PriceOracleAdapter *PriceOracleAdapterTransactorSession) SetMaxPriceAge(newAge uint32) (*types.Transaction, error) {
	return _PriceOracleAdapter.Contract.SetMaxPriceAge(&_PriceOracleAdapter.TransactOpts, newAge)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceOracleAdapter *PriceOracleAdapterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PriceOracleAdapter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceOracleAdapter *PriceOracleAdapterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceOracleAdapter.Contract.TransferOwnership(&_PriceOracleAdapter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceOracleAdapter *PriceOracleAdapterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceOracleAdapter.Contract.TransferOwnership(&_PriceOracleAdapter.TransactOpts, newOwner)
}

// PriceOracleAdapterFeedChangedIterator is returned from FilterFeedChanged and is used to iterate over the raw logs and unpacked data for FeedChanged events raised by the PriceOracleAdapter contract.
type PriceOracleAdapterFeedChangedIterator struct {
	Event *PriceOracleAdapterFeedChanged // Event containing the contract specifics and raw log

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
func (it *PriceOracleAdapterFeedChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOracleAdapterFeedChanged)
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
		it.Event = new(PriceOracleAdapterFeedChanged)
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
func (it *PriceOracleAdapterFeedChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOracleAdapterFeedChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOracleAdapterFeedChanged represents a FeedChanged event raised by the PriceOracleAdapter contract.
type PriceOracleAdapterFeedChanged struct {
	Asset      common.Address
	Aggregator common.Address
	Decimals   uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeedChanged is a free log retrieval operation binding the contract event 0x242add3d2bce07a21d81982973b440003abf7f81bd8cc37563445e18241adbb5.
//
// Solidity: event FeedChanged(address indexed asset, address aggregator, uint8 decimals)
func (_PriceOracleAdapter *PriceOracleAdapterFilterer) FilterFeedChanged(opts *bind.FilterOpts, asset []common.Address) (*PriceOracleAdapterFeedChangedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _PriceOracleAdapter.contract.FilterLogs(opts, "FeedChanged", assetRule)
	if err != nil {
		return nil, err
	}
	return &PriceOracleAdapterFeedChangedIterator{contract: _PriceOracleAdapter.contract, event: "FeedChanged", logs: logs, sub: sub}, nil
}

// WatchFeedChanged is a free log subscription operation binding the contract event 0x242add3d2bce07a21d81982973b440003abf7f81bd8cc37563445e18241adbb5.
//
// Solidity: event FeedChanged(address indexed asset, address aggregator, uint8 decimals)
func (_PriceOracleAdapter *PriceOracleAdapterFilterer) WatchFeedChanged(opts *bind.WatchOpts, sink chan<- *PriceOracleAdapterFeedChanged, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _PriceOracleAdapter.contract.WatchLogs(opts, "FeedChanged", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOracleAdapterFeedChanged)
				if err := _PriceOracleAdapter.contract.UnpackLog(event, "FeedChanged", log); err != nil {
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

// ParseFeedChanged is a log parse operation binding the contract event 0x242add3d2bce07a21d81982973b440003abf7f81bd8cc37563445e18241adbb5.
//
// Solidity: event FeedChanged(address indexed asset, address aggregator, uint8 decimals)
func (_PriceOracleAdapter *PriceOracleAdapterFilterer) ParseFeedChanged(log types.Log) (*PriceOracleAdapterFeedChanged, error) {
	event := new(PriceOracleAdapterFeedChanged)
	if err := _PriceOracleAdapter.contract.UnpackLog(event, "FeedChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceOracleAdapterMaxPriceAgeChangedIterator is returned from FilterMaxPriceAgeChanged and is used to iterate over the raw logs and unpacked data for MaxPriceAgeChanged events raised by the PriceOracleAdapter contract.
type PriceOracleAdapterMaxPriceAgeChangedIterator struct {
	Event *PriceOracleAdapterMaxPriceAgeChanged // Event containing the contract specifics and raw log

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
func (it *PriceOracleAdapterMaxPriceAgeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOracleAdapterMaxPriceAgeChanged)
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
		it.Event = new(PriceOracleAdapterMaxPriceAgeChanged)
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
func (it *PriceOracleAdapterMaxPriceAgeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOracleAdapterMaxPriceAgeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOracleAdapterMaxPriceAgeChanged represents a MaxPriceAgeChanged event raised by the PriceOracleAdapter contract.
type PriceOracleAdapterMaxPriceAgeChanged struct {
	PreviousAge uint32
	NewAge      uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMaxPriceAgeChanged is a free log retrieval operation binding the contract event 0x8285b872e17e7c1a5db5d46d8c5de86aea953e03076d9b4733fe02777276b8f8.
//
// Solidity: event MaxPriceAgeChanged(uint32 previousAge, uint32 newAge)
func (_PriceOracleAdapter *PriceOracleAdapterFilterer) FilterMaxPriceAgeChanged(opts *bind.FilterOpts) (*PriceOracleAdapterMaxPriceAgeChangedIterator, error) {

	logs, sub, err := _PriceOracleAdapter.contract.FilterLogs(opts, "MaxPriceAgeChanged")
	if err != nil {
		return nil, err
	}
	return &PriceOracleAdapterMaxPriceAgeChangedIterator{contract: _PriceOracleAdapter.contract, event: "MaxPriceAgeChanged", logs: logs, sub: sub}, nil
}

// WatchMaxPriceAgeChanged is a free log subscription operation binding the contract event 0x8285b872e17e7c1a5db5d46d8c5de86aea953e03076d9b4733fe02777276b8f8.
//
// Solidity: event MaxPriceAgeChanged(uint32 previousAge, uint32 newAge)
func (_PriceOracleAdapter *PriceOracleAdapterFilterer) WatchMaxPriceAgeChanged(opts *bind.WatchOpts, sink chan<- *PriceOracleAdapterMaxPriceAgeChanged) (event.Subscription, error) {

	logs, sub, err := _PriceOracleAdapter.contract.WatchLogs(opts, "MaxPriceAgeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOracleAdapterMaxPriceAgeChanged)
				if err := _PriceOracleAdapter.contract.UnpackLog(event, "MaxPriceAgeChanged", log); err != nil {
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

// ParseMaxPriceAgeChanged is a log parse operation binding the contract event 0x8285b872e17e7c1a5db5d46d8c5de86aea953e03076d9b4733fe02777276b8f8.
//
// Solidity: event MaxPriceAgeChanged(uint32 previousAge, uint32 newAge)
func (_PriceOracleAdapter *PriceOracleAdapterFilterer) ParseMaxPriceAgeChanged(log types.Log) (*PriceOracleAdapterMaxPriceAgeChanged, error) {
	event := new(PriceOracleAdapterMaxPriceAgeChanged)
	if err := _PriceOracleAdapter.contract.UnpackLog(event, "MaxPriceAgeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceOracleAdapterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PriceOracleAdapter contract.
type PriceOracleAdapterOwnershipTransferredIterator struct {
	Event *PriceOracleAdapterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PriceOracleAdapterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOracleAdapterOwnershipTransferred)
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
		it.Event = new(PriceOracleAdapterOwnershipTransferred)
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
func (it *PriceOracleAdapterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOracleAdapterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOracleAdapterOwnershipTransferred represents a OwnershipTransferred event raised by the PriceOracleAdapter contract.
type PriceOracleAdapterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceOracleAdapter *PriceOracleAdapterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PriceOracleAdapterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceOracleAdapter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PriceOracleAdapterOwnershipTransferredIterator{contract: _PriceOracleAdapter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceOracleAdapter *PriceOracleAdapterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceOracleAdapterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceOracleAdapter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOracleAdapterOwnershipTransferred)
				if err := _PriceOracleAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PriceOracleAdapter *PriceOracleAdapterFilterer) ParseOwnershipTransferred(log types.Log) (*PriceOracleAdapterOwnershipTransferred, error) {
	event := new(PriceOracleAdapterOwnershipTransferred)
	if err := _PriceOracleAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
