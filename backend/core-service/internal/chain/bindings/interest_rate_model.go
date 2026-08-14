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

// RateCurve is an auto generated low-level Go binding around an user-defined struct.
type RateCurve struct {
	BaseRatePerSecond       uint64
	SlopeBelowKinkPerSecond uint64
	SlopeAboveKinkPerSecond uint64
	KinkUtilizationBps      uint16
}

// InterestRateModelMetaData contains all meta data concerning the InterestRateModel contract.
var InterestRateModelMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"startingCurve\",\"type\":\"tuple\",\"internalType\":\"structRateCurve\",\"components\":[{\"name\":\"baseRatePerSecond\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slopeBelowKinkPerSecond\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slopeAboveKinkPerSecond\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"kinkUtilizationBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"borrowAprBps\",\"inputs\":[{\"name\":\"usageBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"borrowRatePerSecond\",\"inputs\":[{\"name\":\"usageBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"curve\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateCurve\",\"components\":[{\"name\":\"baseRatePerSecond\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slopeBelowKinkPerSecond\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slopeAboveKinkPerSecond\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"kinkUtilizationBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCurve\",\"inputs\":[{\"name\":\"newCurve\",\"type\":\"tuple\",\"internalType\":\"structRateCurve\",\"components\":[{\"name\":\"baseRatePerSecond\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slopeBelowKinkPerSecond\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"slopeAboveKinkPerSecond\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"kinkUtilizationBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supplyAprBps\",\"inputs\":[{\"name\":\"usageBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserveFactorBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supplyRatePerSecond\",\"inputs\":[{\"name\":\"usageBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserveFactorBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"utilizationBps\",\"inputs\":[{\"name\":\"totalSupplied\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalBorrowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"CurveChanged\",\"inputs\":[{\"name\":\"baseRatePerSecond\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"slopeBelowKinkPerSecond\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"slopeAboveKinkPerSecond\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"kinkUtilizationBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidRiskSettings\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// InterestRateModelABI is the input ABI used to generate the binding from.
// Deprecated: Use InterestRateModelMetaData.ABI instead.
var InterestRateModelABI = InterestRateModelMetaData.ABI

// InterestRateModel is an auto generated Go binding around an Ethereum contract.
type InterestRateModel struct {
	InterestRateModelCaller     // Read-only binding to the contract
	InterestRateModelTransactor // Write-only binding to the contract
	InterestRateModelFilterer   // Log filterer for contract events
}

// InterestRateModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterestRateModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterestRateModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterestRateModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterestRateModelSession struct {
	Contract     *InterestRateModel // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// InterestRateModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterestRateModelCallerSession struct {
	Contract *InterestRateModelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// InterestRateModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterestRateModelTransactorSession struct {
	Contract     *InterestRateModelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// InterestRateModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterestRateModelRaw struct {
	Contract *InterestRateModel // Generic contract binding to access the raw methods on
}

// InterestRateModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterestRateModelCallerRaw struct {
	Contract *InterestRateModelCaller // Generic read-only contract binding to access the raw methods on
}

// InterestRateModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterestRateModelTransactorRaw struct {
	Contract *InterestRateModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterestRateModel creates a new instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModel(address common.Address, backend bind.ContractBackend) (*InterestRateModel, error) {
	contract, err := bindInterestRateModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterestRateModel{InterestRateModelCaller: InterestRateModelCaller{contract: contract}, InterestRateModelTransactor: InterestRateModelTransactor{contract: contract}, InterestRateModelFilterer: InterestRateModelFilterer{contract: contract}}, nil
}

// NewInterestRateModelCaller creates a new read-only instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModelCaller(address common.Address, caller bind.ContractCaller) (*InterestRateModelCaller, error) {
	contract, err := bindInterestRateModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelCaller{contract: contract}, nil
}

// NewInterestRateModelTransactor creates a new write-only instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModelTransactor(address common.Address, transactor bind.ContractTransactor) (*InterestRateModelTransactor, error) {
	contract, err := bindInterestRateModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelTransactor{contract: contract}, nil
}

// NewInterestRateModelFilterer creates a new log filterer instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModelFilterer(address common.Address, filterer bind.ContractFilterer) (*InterestRateModelFilterer, error) {
	contract, err := bindInterestRateModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelFilterer{contract: contract}, nil
}

// bindInterestRateModel binds a generic wrapper to an already deployed contract.
func bindInterestRateModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterestRateModelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateModel *InterestRateModelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateModel.Contract.InterestRateModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateModel *InterestRateModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModel.Contract.InterestRateModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateModel *InterestRateModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateModel.Contract.InterestRateModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateModel *InterestRateModelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateModel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateModel *InterestRateModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateModel *InterestRateModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateModel.Contract.contract.Transact(opts, method, params...)
}

// BorrowAprBps is a free data retrieval call binding the contract method 0x13c73aed.
//
// Solidity: function borrowAprBps(uint256 usageBps) view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) BorrowAprBps(opts *bind.CallOpts, usageBps *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "borrowAprBps", usageBps)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowAprBps is a free data retrieval call binding the contract method 0x13c73aed.
//
// Solidity: function borrowAprBps(uint256 usageBps) view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) BorrowAprBps(usageBps *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.BorrowAprBps(&_InterestRateModel.CallOpts, usageBps)
}

// BorrowAprBps is a free data retrieval call binding the contract method 0x13c73aed.
//
// Solidity: function borrowAprBps(uint256 usageBps) view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) BorrowAprBps(usageBps *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.BorrowAprBps(&_InterestRateModel.CallOpts, usageBps)
}

// BorrowRatePerSecond is a free data retrieval call binding the contract method 0x4132b337.
//
// Solidity: function borrowRatePerSecond(uint256 usageBps) view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) BorrowRatePerSecond(opts *bind.CallOpts, usageBps *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "borrowRatePerSecond", usageBps)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerSecond is a free data retrieval call binding the contract method 0x4132b337.
//
// Solidity: function borrowRatePerSecond(uint256 usageBps) view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) BorrowRatePerSecond(usageBps *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.BorrowRatePerSecond(&_InterestRateModel.CallOpts, usageBps)
}

// BorrowRatePerSecond is a free data retrieval call binding the contract method 0x4132b337.
//
// Solidity: function borrowRatePerSecond(uint256 usageBps) view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) BorrowRatePerSecond(usageBps *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.BorrowRatePerSecond(&_InterestRateModel.CallOpts, usageBps)
}

// Curve is a free data retrieval call binding the contract method 0x7165485d.
//
// Solidity: function curve() view returns((uint64,uint64,uint64,uint16))
func (_InterestRateModel *InterestRateModelCaller) Curve(opts *bind.CallOpts) (RateCurve, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "curve")

	if err != nil {
		return *new(RateCurve), err
	}

	out0 := *abi.ConvertType(out[0], new(RateCurve)).(*RateCurve)

	return out0, err

}

// Curve is a free data retrieval call binding the contract method 0x7165485d.
//
// Solidity: function curve() view returns((uint64,uint64,uint64,uint16))
func (_InterestRateModel *InterestRateModelSession) Curve() (RateCurve, error) {
	return _InterestRateModel.Contract.Curve(&_InterestRateModel.CallOpts)
}

// Curve is a free data retrieval call binding the contract method 0x7165485d.
//
// Solidity: function curve() view returns((uint64,uint64,uint64,uint16))
func (_InterestRateModel *InterestRateModelCallerSession) Curve() (RateCurve, error) {
	return _InterestRateModel.Contract.Curve(&_InterestRateModel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterestRateModel *InterestRateModelCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterestRateModel *InterestRateModelSession) Owner() (common.Address, error) {
	return _InterestRateModel.Contract.Owner(&_InterestRateModel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterestRateModel *InterestRateModelCallerSession) Owner() (common.Address, error) {
	return _InterestRateModel.Contract.Owner(&_InterestRateModel.CallOpts)
}

// SupplyAprBps is a free data retrieval call binding the contract method 0xfc80e760.
//
// Solidity: function supplyAprBps(uint256 usageBps, uint256 reserveFactorBps) view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) SupplyAprBps(opts *bind.CallOpts, usageBps *big.Int, reserveFactorBps *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "supplyAprBps", usageBps, reserveFactorBps)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyAprBps is a free data retrieval call binding the contract method 0xfc80e760.
//
// Solidity: function supplyAprBps(uint256 usageBps, uint256 reserveFactorBps) view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) SupplyAprBps(usageBps *big.Int, reserveFactorBps *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.SupplyAprBps(&_InterestRateModel.CallOpts, usageBps, reserveFactorBps)
}

// SupplyAprBps is a free data retrieval call binding the contract method 0xfc80e760.
//
// Solidity: function supplyAprBps(uint256 usageBps, uint256 reserveFactorBps) view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) SupplyAprBps(usageBps *big.Int, reserveFactorBps *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.SupplyAprBps(&_InterestRateModel.CallOpts, usageBps, reserveFactorBps)
}

// SupplyRatePerSecond is a free data retrieval call binding the contract method 0x67635145.
//
// Solidity: function supplyRatePerSecond(uint256 usageBps, uint256 reserveFactorBps) view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) SupplyRatePerSecond(opts *bind.CallOpts, usageBps *big.Int, reserveFactorBps *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "supplyRatePerSecond", usageBps, reserveFactorBps)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerSecond is a free data retrieval call binding the contract method 0x67635145.
//
// Solidity: function supplyRatePerSecond(uint256 usageBps, uint256 reserveFactorBps) view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) SupplyRatePerSecond(usageBps *big.Int, reserveFactorBps *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.SupplyRatePerSecond(&_InterestRateModel.CallOpts, usageBps, reserveFactorBps)
}

// SupplyRatePerSecond is a free data retrieval call binding the contract method 0x67635145.
//
// Solidity: function supplyRatePerSecond(uint256 usageBps, uint256 reserveFactorBps) view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) SupplyRatePerSecond(usageBps *big.Int, reserveFactorBps *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.SupplyRatePerSecond(&_InterestRateModel.CallOpts, usageBps, reserveFactorBps)
}

// UtilizationBps is a free data retrieval call binding the contract method 0x621b464e.
//
// Solidity: function utilizationBps(uint256 totalSupplied, uint256 totalBorrowed) pure returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) UtilizationBps(opts *bind.CallOpts, totalSupplied *big.Int, totalBorrowed *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "utilizationBps", totalSupplied, totalBorrowed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UtilizationBps is a free data retrieval call binding the contract method 0x621b464e.
//
// Solidity: function utilizationBps(uint256 totalSupplied, uint256 totalBorrowed) pure returns(uint256)
func (_InterestRateModel *InterestRateModelSession) UtilizationBps(totalSupplied *big.Int, totalBorrowed *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.UtilizationBps(&_InterestRateModel.CallOpts, totalSupplied, totalBorrowed)
}

// UtilizationBps is a free data retrieval call binding the contract method 0x621b464e.
//
// Solidity: function utilizationBps(uint256 totalSupplied, uint256 totalBorrowed) pure returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) UtilizationBps(totalSupplied *big.Int, totalBorrowed *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.UtilizationBps(&_InterestRateModel.CallOpts, totalSupplied, totalBorrowed)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterestRateModel *InterestRateModelTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModel.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterestRateModel *InterestRateModelSession) RenounceOwnership() (*types.Transaction, error) {
	return _InterestRateModel.Contract.RenounceOwnership(&_InterestRateModel.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterestRateModel *InterestRateModelTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _InterestRateModel.Contract.RenounceOwnership(&_InterestRateModel.TransactOpts)
}

// SetCurve is a paid mutator transaction binding the contract method 0x1eb37dd2.
//
// Solidity: function setCurve((uint64,uint64,uint64,uint16) newCurve) returns()
func (_InterestRateModel *InterestRateModelTransactor) SetCurve(opts *bind.TransactOpts, newCurve RateCurve) (*types.Transaction, error) {
	return _InterestRateModel.contract.Transact(opts, "setCurve", newCurve)
}

// SetCurve is a paid mutator transaction binding the contract method 0x1eb37dd2.
//
// Solidity: function setCurve((uint64,uint64,uint64,uint16) newCurve) returns()
func (_InterestRateModel *InterestRateModelSession) SetCurve(newCurve RateCurve) (*types.Transaction, error) {
	return _InterestRateModel.Contract.SetCurve(&_InterestRateModel.TransactOpts, newCurve)
}

// SetCurve is a paid mutator transaction binding the contract method 0x1eb37dd2.
//
// Solidity: function setCurve((uint64,uint64,uint64,uint16) newCurve) returns()
func (_InterestRateModel *InterestRateModelTransactorSession) SetCurve(newCurve RateCurve) (*types.Transaction, error) {
	return _InterestRateModel.Contract.SetCurve(&_InterestRateModel.TransactOpts, newCurve)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterestRateModel *InterestRateModelTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModel.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterestRateModel *InterestRateModelSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModel.Contract.TransferOwnership(&_InterestRateModel.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterestRateModel *InterestRateModelTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InterestRateModel.Contract.TransferOwnership(&_InterestRateModel.TransactOpts, newOwner)
}

// InterestRateModelCurveChangedIterator is returned from FilterCurveChanged and is used to iterate over the raw logs and unpacked data for CurveChanged events raised by the InterestRateModel contract.
type InterestRateModelCurveChangedIterator struct {
	Event *InterestRateModelCurveChanged // Event containing the contract specifics and raw log

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
func (it *InterestRateModelCurveChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterestRateModelCurveChanged)
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
		it.Event = new(InterestRateModelCurveChanged)
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
func (it *InterestRateModelCurveChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterestRateModelCurveChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterestRateModelCurveChanged represents a CurveChanged event raised by the InterestRateModel contract.
type InterestRateModelCurveChanged struct {
	BaseRatePerSecond       uint64
	SlopeBelowKinkPerSecond uint64
	SlopeAboveKinkPerSecond uint64
	KinkUtilizationBps      uint16
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterCurveChanged is a free log retrieval operation binding the contract event 0xa02a04661df5a0520dc34d3810d89030727554af0a17703ba95c9504ad78bb3b.
//
// Solidity: event CurveChanged(uint64 baseRatePerSecond, uint64 slopeBelowKinkPerSecond, uint64 slopeAboveKinkPerSecond, uint16 kinkUtilizationBps)
func (_InterestRateModel *InterestRateModelFilterer) FilterCurveChanged(opts *bind.FilterOpts) (*InterestRateModelCurveChangedIterator, error) {

	logs, sub, err := _InterestRateModel.contract.FilterLogs(opts, "CurveChanged")
	if err != nil {
		return nil, err
	}
	return &InterestRateModelCurveChangedIterator{contract: _InterestRateModel.contract, event: "CurveChanged", logs: logs, sub: sub}, nil
}

// WatchCurveChanged is a free log subscription operation binding the contract event 0xa02a04661df5a0520dc34d3810d89030727554af0a17703ba95c9504ad78bb3b.
//
// Solidity: event CurveChanged(uint64 baseRatePerSecond, uint64 slopeBelowKinkPerSecond, uint64 slopeAboveKinkPerSecond, uint16 kinkUtilizationBps)
func (_InterestRateModel *InterestRateModelFilterer) WatchCurveChanged(opts *bind.WatchOpts, sink chan<- *InterestRateModelCurveChanged) (event.Subscription, error) {

	logs, sub, err := _InterestRateModel.contract.WatchLogs(opts, "CurveChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterestRateModelCurveChanged)
				if err := _InterestRateModel.contract.UnpackLog(event, "CurveChanged", log); err != nil {
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

// ParseCurveChanged is a log parse operation binding the contract event 0xa02a04661df5a0520dc34d3810d89030727554af0a17703ba95c9504ad78bb3b.
//
// Solidity: event CurveChanged(uint64 baseRatePerSecond, uint64 slopeBelowKinkPerSecond, uint64 slopeAboveKinkPerSecond, uint16 kinkUtilizationBps)
func (_InterestRateModel *InterestRateModelFilterer) ParseCurveChanged(log types.Log) (*InterestRateModelCurveChanged, error) {
	event := new(InterestRateModelCurveChanged)
	if err := _InterestRateModel.contract.UnpackLog(event, "CurveChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterestRateModelOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InterestRateModel contract.
type InterestRateModelOwnershipTransferredIterator struct {
	Event *InterestRateModelOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InterestRateModelOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterestRateModelOwnershipTransferred)
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
		it.Event = new(InterestRateModelOwnershipTransferred)
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
func (it *InterestRateModelOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterestRateModelOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterestRateModelOwnershipTransferred represents a OwnershipTransferred event raised by the InterestRateModel contract.
type InterestRateModelOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InterestRateModel *InterestRateModelFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InterestRateModelOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InterestRateModel.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelOwnershipTransferredIterator{contract: _InterestRateModel.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InterestRateModel *InterestRateModelFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InterestRateModelOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InterestRateModel.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterestRateModelOwnershipTransferred)
				if err := _InterestRateModel.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InterestRateModel *InterestRateModelFilterer) ParseOwnershipTransferred(log types.Log) (*InterestRateModelOwnershipTransferred, error) {
	event := new(InterestRateModelOwnershipTransferred)
	if err := _InterestRateModel.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
