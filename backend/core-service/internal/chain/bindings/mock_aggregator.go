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

// MockAggregatorMetaData contains all meta data concerning the MockAggregator contract.
var MockAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"description_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"startingAnswer\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"description\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feedDown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoundData\",\"inputs\":[{\"name\":\"wantedRoundId\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"outputs\":[{\"name\":\"roundId\",\"type\":\"uint80\",\"internalType\":\"uint80\"},{\"name\":\"answer\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"startedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"answeredInRound\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestRoundData\",\"inputs\":[],\"outputs\":[{\"name\":\"roundId\",\"type\":\"uint80\",\"internalType\":\"uint80\"},{\"name\":\"answer\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"startedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"answeredInRound\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestRoundId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"makeStale\",\"inputs\":[{\"name\":\"secondsInThePast\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeedDown\",\"inputs\":[{\"name\":\"isDown\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIncompleteRound\",\"inputs\":[{\"name\":\"newAnswer\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPrice\",\"inputs\":[{\"name\":\"newAnswer\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceWithTimestamp\",\"inputs\":[{\"name\":\"newAnswer\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"updatedAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"AnswerRecorded\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"uint80\",\"indexed\":true,\"internalType\":\"uint80\"},{\"name\":\"answer\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"updatedAt\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"FeedIsDown\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RoundNotFound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"uint80\",\"internalType\":\"uint80\"}]}]",
}

// MockAggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use MockAggregatorMetaData.ABI instead.
var MockAggregatorABI = MockAggregatorMetaData.ABI

// MockAggregator is an auto generated Go binding around an Ethereum contract.
type MockAggregator struct {
	MockAggregatorCaller     // Read-only binding to the contract
	MockAggregatorTransactor // Write-only binding to the contract
	MockAggregatorFilterer   // Log filterer for contract events
}

// MockAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockAggregatorSession struct {
	Contract     *MockAggregator   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockAggregatorCallerSession struct {
	Contract *MockAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MockAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockAggregatorTransactorSession struct {
	Contract     *MockAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MockAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockAggregatorRaw struct {
	Contract *MockAggregator // Generic contract binding to access the raw methods on
}

// MockAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockAggregatorCallerRaw struct {
	Contract *MockAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// MockAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockAggregatorTransactorRaw struct {
	Contract *MockAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockAggregator creates a new instance of MockAggregator, bound to a specific deployed contract.
func NewMockAggregator(address common.Address, backend bind.ContractBackend) (*MockAggregator, error) {
	contract, err := bindMockAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockAggregator{MockAggregatorCaller: MockAggregatorCaller{contract: contract}, MockAggregatorTransactor: MockAggregatorTransactor{contract: contract}, MockAggregatorFilterer: MockAggregatorFilterer{contract: contract}}, nil
}

// NewMockAggregatorCaller creates a new read-only instance of MockAggregator, bound to a specific deployed contract.
func NewMockAggregatorCaller(address common.Address, caller bind.ContractCaller) (*MockAggregatorCaller, error) {
	contract, err := bindMockAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockAggregatorCaller{contract: contract}, nil
}

// NewMockAggregatorTransactor creates a new write-only instance of MockAggregator, bound to a specific deployed contract.
func NewMockAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*MockAggregatorTransactor, error) {
	contract, err := bindMockAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockAggregatorTransactor{contract: contract}, nil
}

// NewMockAggregatorFilterer creates a new log filterer instance of MockAggregator, bound to a specific deployed contract.
func NewMockAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*MockAggregatorFilterer, error) {
	contract, err := bindMockAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockAggregatorFilterer{contract: contract}, nil
}

// bindMockAggregator binds a generic wrapper to an already deployed contract.
func bindMockAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockAggregatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockAggregator *MockAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockAggregator.Contract.MockAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockAggregator *MockAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockAggregator.Contract.MockAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockAggregator *MockAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockAggregator.Contract.MockAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockAggregator *MockAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockAggregator *MockAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockAggregator *MockAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockAggregator.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockAggregator *MockAggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MockAggregator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockAggregator *MockAggregatorSession) Decimals() (uint8, error) {
	return _MockAggregator.Contract.Decimals(&_MockAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockAggregator *MockAggregatorCallerSession) Decimals() (uint8, error) {
	return _MockAggregator.Contract.Decimals(&_MockAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_MockAggregator *MockAggregatorCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockAggregator.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_MockAggregator *MockAggregatorSession) Description() (string, error) {
	return _MockAggregator.Contract.Description(&_MockAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_MockAggregator *MockAggregatorCallerSession) Description() (string, error) {
	return _MockAggregator.Contract.Description(&_MockAggregator.CallOpts)
}

// FeedDown is a free data retrieval call binding the contract method 0x53d65cfb.
//
// Solidity: function feedDown() view returns(bool)
func (_MockAggregator *MockAggregatorCaller) FeedDown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockAggregator.contract.Call(opts, &out, "feedDown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FeedDown is a free data retrieval call binding the contract method 0x53d65cfb.
//
// Solidity: function feedDown() view returns(bool)
func (_MockAggregator *MockAggregatorSession) FeedDown() (bool, error) {
	return _MockAggregator.Contract.FeedDown(&_MockAggregator.CallOpts)
}

// FeedDown is a free data retrieval call binding the contract method 0x53d65cfb.
//
// Solidity: function feedDown() view returns(bool)
func (_MockAggregator *MockAggregatorCallerSession) FeedDown() (bool, error) {
	return _MockAggregator.Contract.FeedDown(&_MockAggregator.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 wantedRoundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_MockAggregator *MockAggregatorCaller) GetRoundData(opts *bind.CallOpts, wantedRoundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _MockAggregator.contract.Call(opts, &out, "getRoundData", wantedRoundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 wantedRoundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_MockAggregator *MockAggregatorSession) GetRoundData(wantedRoundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _MockAggregator.Contract.GetRoundData(&_MockAggregator.CallOpts, wantedRoundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 wantedRoundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_MockAggregator *MockAggregatorCallerSession) GetRoundData(wantedRoundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _MockAggregator.Contract.GetRoundData(&_MockAggregator.CallOpts, wantedRoundId)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_MockAggregator *MockAggregatorCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _MockAggregator.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_MockAggregator *MockAggregatorSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _MockAggregator.Contract.LatestRoundData(&_MockAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_MockAggregator *MockAggregatorCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _MockAggregator.Contract.LatestRoundData(&_MockAggregator.CallOpts)
}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_MockAggregator *MockAggregatorCaller) LatestRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockAggregator.contract.Call(opts, &out, "latestRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_MockAggregator *MockAggregatorSession) LatestRoundId() (*big.Int, error) {
	return _MockAggregator.Contract.LatestRoundId(&_MockAggregator.CallOpts)
}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_MockAggregator *MockAggregatorCallerSession) LatestRoundId() (*big.Int, error) {
	return _MockAggregator.Contract.LatestRoundId(&_MockAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_MockAggregator *MockAggregatorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockAggregator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_MockAggregator *MockAggregatorSession) Version() (*big.Int, error) {
	return _MockAggregator.Contract.Version(&_MockAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_MockAggregator *MockAggregatorCallerSession) Version() (*big.Int, error) {
	return _MockAggregator.Contract.Version(&_MockAggregator.CallOpts)
}

// MakeStale is a paid mutator transaction binding the contract method 0x6448fb78.
//
// Solidity: function makeStale(uint40 secondsInThePast) returns()
func (_MockAggregator *MockAggregatorTransactor) MakeStale(opts *bind.TransactOpts, secondsInThePast *big.Int) (*types.Transaction, error) {
	return _MockAggregator.contract.Transact(opts, "makeStale", secondsInThePast)
}

// MakeStale is a paid mutator transaction binding the contract method 0x6448fb78.
//
// Solidity: function makeStale(uint40 secondsInThePast) returns()
func (_MockAggregator *MockAggregatorSession) MakeStale(secondsInThePast *big.Int) (*types.Transaction, error) {
	return _MockAggregator.Contract.MakeStale(&_MockAggregator.TransactOpts, secondsInThePast)
}

// MakeStale is a paid mutator transaction binding the contract method 0x6448fb78.
//
// Solidity: function makeStale(uint40 secondsInThePast) returns()
func (_MockAggregator *MockAggregatorTransactorSession) MakeStale(secondsInThePast *big.Int) (*types.Transaction, error) {
	return _MockAggregator.Contract.MakeStale(&_MockAggregator.TransactOpts, secondsInThePast)
}

// SetFeedDown is a paid mutator transaction binding the contract method 0x5020bba4.
//
// Solidity: function setFeedDown(bool isDown) returns()
func (_MockAggregator *MockAggregatorTransactor) SetFeedDown(opts *bind.TransactOpts, isDown bool) (*types.Transaction, error) {
	return _MockAggregator.contract.Transact(opts, "setFeedDown", isDown)
}

// SetFeedDown is a paid mutator transaction binding the contract method 0x5020bba4.
//
// Solidity: function setFeedDown(bool isDown) returns()
func (_MockAggregator *MockAggregatorSession) SetFeedDown(isDown bool) (*types.Transaction, error) {
	return _MockAggregator.Contract.SetFeedDown(&_MockAggregator.TransactOpts, isDown)
}

// SetFeedDown is a paid mutator transaction binding the contract method 0x5020bba4.
//
// Solidity: function setFeedDown(bool isDown) returns()
func (_MockAggregator *MockAggregatorTransactorSession) SetFeedDown(isDown bool) (*types.Transaction, error) {
	return _MockAggregator.Contract.SetFeedDown(&_MockAggregator.TransactOpts, isDown)
}

// SetIncompleteRound is a paid mutator transaction binding the contract method 0xeaa5c5b7.
//
// Solidity: function setIncompleteRound(int256 newAnswer) returns()
func (_MockAggregator *MockAggregatorTransactor) SetIncompleteRound(opts *bind.TransactOpts, newAnswer *big.Int) (*types.Transaction, error) {
	return _MockAggregator.contract.Transact(opts, "setIncompleteRound", newAnswer)
}

// SetIncompleteRound is a paid mutator transaction binding the contract method 0xeaa5c5b7.
//
// Solidity: function setIncompleteRound(int256 newAnswer) returns()
func (_MockAggregator *MockAggregatorSession) SetIncompleteRound(newAnswer *big.Int) (*types.Transaction, error) {
	return _MockAggregator.Contract.SetIncompleteRound(&_MockAggregator.TransactOpts, newAnswer)
}

// SetIncompleteRound is a paid mutator transaction binding the contract method 0xeaa5c5b7.
//
// Solidity: function setIncompleteRound(int256 newAnswer) returns()
func (_MockAggregator *MockAggregatorTransactorSession) SetIncompleteRound(newAnswer *big.Int) (*types.Transaction, error) {
	return _MockAggregator.Contract.SetIncompleteRound(&_MockAggregator.TransactOpts, newAnswer)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7a30806.
//
// Solidity: function setPrice(int256 newAnswer) returns()
func (_MockAggregator *MockAggregatorTransactor) SetPrice(opts *bind.TransactOpts, newAnswer *big.Int) (*types.Transaction, error) {
	return _MockAggregator.contract.Transact(opts, "setPrice", newAnswer)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7a30806.
//
// Solidity: function setPrice(int256 newAnswer) returns()
func (_MockAggregator *MockAggregatorSession) SetPrice(newAnswer *big.Int) (*types.Transaction, error) {
	return _MockAggregator.Contract.SetPrice(&_MockAggregator.TransactOpts, newAnswer)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7a30806.
//
// Solidity: function setPrice(int256 newAnswer) returns()
func (_MockAggregator *MockAggregatorTransactorSession) SetPrice(newAnswer *big.Int) (*types.Transaction, error) {
	return _MockAggregator.Contract.SetPrice(&_MockAggregator.TransactOpts, newAnswer)
}

// SetPriceWithTimestamp is a paid mutator transaction binding the contract method 0x8fbca908.
//
// Solidity: function setPriceWithTimestamp(int256 newAnswer, uint40 updatedAt) returns()
func (_MockAggregator *MockAggregatorTransactor) SetPriceWithTimestamp(opts *bind.TransactOpts, newAnswer *big.Int, updatedAt *big.Int) (*types.Transaction, error) {
	return _MockAggregator.contract.Transact(opts, "setPriceWithTimestamp", newAnswer, updatedAt)
}

// SetPriceWithTimestamp is a paid mutator transaction binding the contract method 0x8fbca908.
//
// Solidity: function setPriceWithTimestamp(int256 newAnswer, uint40 updatedAt) returns()
func (_MockAggregator *MockAggregatorSession) SetPriceWithTimestamp(newAnswer *big.Int, updatedAt *big.Int) (*types.Transaction, error) {
	return _MockAggregator.Contract.SetPriceWithTimestamp(&_MockAggregator.TransactOpts, newAnswer, updatedAt)
}

// SetPriceWithTimestamp is a paid mutator transaction binding the contract method 0x8fbca908.
//
// Solidity: function setPriceWithTimestamp(int256 newAnswer, uint40 updatedAt) returns()
func (_MockAggregator *MockAggregatorTransactorSession) SetPriceWithTimestamp(newAnswer *big.Int, updatedAt *big.Int) (*types.Transaction, error) {
	return _MockAggregator.Contract.SetPriceWithTimestamp(&_MockAggregator.TransactOpts, newAnswer, updatedAt)
}

// MockAggregatorAnswerRecordedIterator is returned from FilterAnswerRecorded and is used to iterate over the raw logs and unpacked data for AnswerRecorded events raised by the MockAggregator contract.
type MockAggregatorAnswerRecordedIterator struct {
	Event *MockAggregatorAnswerRecorded // Event containing the contract specifics and raw log

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
func (it *MockAggregatorAnswerRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockAggregatorAnswerRecorded)
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
		it.Event = new(MockAggregatorAnswerRecorded)
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
func (it *MockAggregatorAnswerRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockAggregatorAnswerRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockAggregatorAnswerRecorded represents a AnswerRecorded event raised by the MockAggregator contract.
type MockAggregatorAnswerRecorded struct {
	RoundId   *big.Int
	Answer    *big.Int
	UpdatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerRecorded is a free log retrieval operation binding the contract event 0xedfc54408aaa709a7f9f945221150951bc6c16f2d460a8adf7bed6deab161363.
//
// Solidity: event AnswerRecorded(uint80 indexed roundId, int256 answer, uint40 updatedAt)
func (_MockAggregator *MockAggregatorFilterer) FilterAnswerRecorded(opts *bind.FilterOpts, roundId []*big.Int) (*MockAggregatorAnswerRecordedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _MockAggregator.contract.FilterLogs(opts, "AnswerRecorded", roundIdRule)
	if err != nil {
		return nil, err
	}
	return &MockAggregatorAnswerRecordedIterator{contract: _MockAggregator.contract, event: "AnswerRecorded", logs: logs, sub: sub}, nil
}

// WatchAnswerRecorded is a free log subscription operation binding the contract event 0xedfc54408aaa709a7f9f945221150951bc6c16f2d460a8adf7bed6deab161363.
//
// Solidity: event AnswerRecorded(uint80 indexed roundId, int256 answer, uint40 updatedAt)
func (_MockAggregator *MockAggregatorFilterer) WatchAnswerRecorded(opts *bind.WatchOpts, sink chan<- *MockAggregatorAnswerRecorded, roundId []*big.Int) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _MockAggregator.contract.WatchLogs(opts, "AnswerRecorded", roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockAggregatorAnswerRecorded)
				if err := _MockAggregator.contract.UnpackLog(event, "AnswerRecorded", log); err != nil {
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

// ParseAnswerRecorded is a log parse operation binding the contract event 0xedfc54408aaa709a7f9f945221150951bc6c16f2d460a8adf7bed6deab161363.
//
// Solidity: event AnswerRecorded(uint80 indexed roundId, int256 answer, uint40 updatedAt)
func (_MockAggregator *MockAggregatorFilterer) ParseAnswerRecorded(log types.Log) (*MockAggregatorAnswerRecorded, error) {
	event := new(MockAggregatorAnswerRecorded)
	if err := _MockAggregator.contract.UnpackLog(event, "AnswerRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
