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

// AccountData is an auto generated low-level Go binding around an user-defined struct.
type AccountData struct {
	SupplyShares              *big.Int
	SupplyAssets              *big.Int
	CollateralAmount          *big.Int
	CollateralValue           *big.Int
	DebtAmount                *big.Int
	DebtValue                 *big.Int
	HealthFactorBps           *big.Int
	MaxBorrowable             *big.Int
	MaxWithdrawableCollateral *big.Int
	CollateralPrice           *big.Int
	PriceUpdatedAt            *big.Int
	IsLiquidatable            bool
	PriceStale                bool
}

// MarketData is an auto generated low-level Go binding around an user-defined struct.
type MarketData struct {
	TotalSupplied           *big.Int
	TotalBorrowed           *big.Int
	AvailableLiquidity      *big.Int
	UtilizationBps          *big.Int
	SupplyRatePerSecond     *big.Int
	BorrowRatePerSecond     *big.Int
	SupplyAprBps            *big.Int
	BorrowAprBps            *big.Int
	SupplyIndex             *big.Int
	BorrowIndex             *big.Int
	MaxLtvBps               *big.Int
	LiquidationThresholdBps *big.Int
	LiquidationBonusBps     *big.Int
	KinkUtilizationBps      *big.Int
	ReserveFactorBps        *big.Int
	MinDeposit              *big.Int
	AccruedReserves         *big.Int
	DepositsPaused          bool
	BorrowPaused            bool
}

// PositionLensMetaData contains all meta data concerning the PositionLens contract.
var PositionLensMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"lendingPool\",\"type\":\"address\",\"internalType\":\"contractILendingPool\"},{\"name\":\"collateralVault\",\"type\":\"address\",\"internalType\":\"contractICollateralVault\"},{\"name\":\"priceOracle\",\"type\":\"address\",\"internalType\":\"contractIPriceOracle\"},{\"name\":\"lendingController\",\"type\":\"address\",\"internalType\":\"contractILendingController\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"accountData\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structAccountData\",\"components\":[{\"name\":\"supplyShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supplyAssets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debtAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debtValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactorBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxBorrowable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxWithdrawableCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priceUpdatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isLiquidatable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"priceStale\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralAsset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"controller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILendingController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"debtAsset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"debtDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketData\",\"inputs\":[],\"outputs\":[{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structMarketData\",\"components\":[{\"name\":\"totalSupplied\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalBorrowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"utilizationBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supplyRatePerSecond\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"borrowRatePerSecond\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supplyAprBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"borrowAprBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supplyIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"borrowIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxLtvBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"liquidationThresholdBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"liquidationBonusBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kinkUtilizationBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserveFactorBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedReserves\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"depositsPaused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"borrowPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPriceOracle\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILendingPool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICollateralVault\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// PositionLensABI is the input ABI used to generate the binding from.
// Deprecated: Use PositionLensMetaData.ABI instead.
var PositionLensABI = PositionLensMetaData.ABI

// PositionLens is an auto generated Go binding around an Ethereum contract.
type PositionLens struct {
	PositionLensCaller     // Read-only binding to the contract
	PositionLensTransactor // Write-only binding to the contract
	PositionLensFilterer   // Log filterer for contract events
}

// PositionLensCaller is an auto generated read-only Go binding around an Ethereum contract.
type PositionLensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionLensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PositionLensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionLensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PositionLensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionLensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PositionLensSession struct {
	Contract     *PositionLens     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PositionLensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PositionLensCallerSession struct {
	Contract *PositionLensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PositionLensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PositionLensTransactorSession struct {
	Contract     *PositionLensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PositionLensRaw is an auto generated low-level Go binding around an Ethereum contract.
type PositionLensRaw struct {
	Contract *PositionLens // Generic contract binding to access the raw methods on
}

// PositionLensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PositionLensCallerRaw struct {
	Contract *PositionLensCaller // Generic read-only contract binding to access the raw methods on
}

// PositionLensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PositionLensTransactorRaw struct {
	Contract *PositionLensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPositionLens creates a new instance of PositionLens, bound to a specific deployed contract.
func NewPositionLens(address common.Address, backend bind.ContractBackend) (*PositionLens, error) {
	contract, err := bindPositionLens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PositionLens{PositionLensCaller: PositionLensCaller{contract: contract}, PositionLensTransactor: PositionLensTransactor{contract: contract}, PositionLensFilterer: PositionLensFilterer{contract: contract}}, nil
}

// NewPositionLensCaller creates a new read-only instance of PositionLens, bound to a specific deployed contract.
func NewPositionLensCaller(address common.Address, caller bind.ContractCaller) (*PositionLensCaller, error) {
	contract, err := bindPositionLens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PositionLensCaller{contract: contract}, nil
}

// NewPositionLensTransactor creates a new write-only instance of PositionLens, bound to a specific deployed contract.
func NewPositionLensTransactor(address common.Address, transactor bind.ContractTransactor) (*PositionLensTransactor, error) {
	contract, err := bindPositionLens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PositionLensTransactor{contract: contract}, nil
}

// NewPositionLensFilterer creates a new log filterer instance of PositionLens, bound to a specific deployed contract.
func NewPositionLensFilterer(address common.Address, filterer bind.ContractFilterer) (*PositionLensFilterer, error) {
	contract, err := bindPositionLens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PositionLensFilterer{contract: contract}, nil
}

// bindPositionLens binds a generic wrapper to an already deployed contract.
func bindPositionLens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PositionLensMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionLens *PositionLensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionLens.Contract.PositionLensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionLens *PositionLensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionLens.Contract.PositionLensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionLens *PositionLensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionLens.Contract.PositionLensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionLens *PositionLensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionLens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionLens *PositionLensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionLens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionLens *PositionLensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionLens.Contract.contract.Transact(opts, method, params...)
}

// AccountData is a free data retrieval call binding the contract method 0xdeb906e7.
//
// Solidity: function accountData(address account) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool) data)
func (_PositionLens *PositionLensCaller) AccountData(opts *bind.CallOpts, account common.Address) (AccountData, error) {
	var out []interface{}
	err := _PositionLens.contract.Call(opts, &out, "accountData", account)

	if err != nil {
		return *new(AccountData), err
	}

	out0 := *abi.ConvertType(out[0], new(AccountData)).(*AccountData)

	return out0, err

}

// AccountData is a free data retrieval call binding the contract method 0xdeb906e7.
//
// Solidity: function accountData(address account) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool) data)
func (_PositionLens *PositionLensSession) AccountData(account common.Address) (AccountData, error) {
	return _PositionLens.Contract.AccountData(&_PositionLens.CallOpts, account)
}

// AccountData is a free data retrieval call binding the contract method 0xdeb906e7.
//
// Solidity: function accountData(address account) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool) data)
func (_PositionLens *PositionLensCallerSession) AccountData(account common.Address) (AccountData, error) {
	return _PositionLens.Contract.AccountData(&_PositionLens.CallOpts, account)
}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_PositionLens *PositionLensCaller) CollateralAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionLens.contract.Call(opts, &out, "collateralAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_PositionLens *PositionLensSession) CollateralAsset() (common.Address, error) {
	return _PositionLens.Contract.CollateralAsset(&_PositionLens.CallOpts)
}

// CollateralAsset is a free data retrieval call binding the contract method 0xaabaecd6.
//
// Solidity: function collateralAsset() view returns(address)
func (_PositionLens *PositionLensCallerSession) CollateralAsset() (common.Address, error) {
	return _PositionLens.Contract.CollateralAsset(&_PositionLens.CallOpts)
}

// CollateralDecimals is a free data retrieval call binding the contract method 0xec9c6c30.
//
// Solidity: function collateralDecimals() view returns(uint8)
func (_PositionLens *PositionLensCaller) CollateralDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PositionLens.contract.Call(opts, &out, "collateralDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CollateralDecimals is a free data retrieval call binding the contract method 0xec9c6c30.
//
// Solidity: function collateralDecimals() view returns(uint8)
func (_PositionLens *PositionLensSession) CollateralDecimals() (uint8, error) {
	return _PositionLens.Contract.CollateralDecimals(&_PositionLens.CallOpts)
}

// CollateralDecimals is a free data retrieval call binding the contract method 0xec9c6c30.
//
// Solidity: function collateralDecimals() view returns(uint8)
func (_PositionLens *PositionLensCallerSession) CollateralDecimals() (uint8, error) {
	return _PositionLens.Contract.CollateralDecimals(&_PositionLens.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PositionLens *PositionLensCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionLens.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PositionLens *PositionLensSession) Controller() (common.Address, error) {
	return _PositionLens.Contract.Controller(&_PositionLens.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PositionLens *PositionLensCallerSession) Controller() (common.Address, error) {
	return _PositionLens.Contract.Controller(&_PositionLens.CallOpts)
}

// DebtAsset is a free data retrieval call binding the contract method 0xa919802d.
//
// Solidity: function debtAsset() view returns(address)
func (_PositionLens *PositionLensCaller) DebtAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionLens.contract.Call(opts, &out, "debtAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DebtAsset is a free data retrieval call binding the contract method 0xa919802d.
//
// Solidity: function debtAsset() view returns(address)
func (_PositionLens *PositionLensSession) DebtAsset() (common.Address, error) {
	return _PositionLens.Contract.DebtAsset(&_PositionLens.CallOpts)
}

// DebtAsset is a free data retrieval call binding the contract method 0xa919802d.
//
// Solidity: function debtAsset() view returns(address)
func (_PositionLens *PositionLensCallerSession) DebtAsset() (common.Address, error) {
	return _PositionLens.Contract.DebtAsset(&_PositionLens.CallOpts)
}

// DebtDecimals is a free data retrieval call binding the contract method 0xf341e228.
//
// Solidity: function debtDecimals() view returns(uint8)
func (_PositionLens *PositionLensCaller) DebtDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PositionLens.contract.Call(opts, &out, "debtDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DebtDecimals is a free data retrieval call binding the contract method 0xf341e228.
//
// Solidity: function debtDecimals() view returns(uint8)
func (_PositionLens *PositionLensSession) DebtDecimals() (uint8, error) {
	return _PositionLens.Contract.DebtDecimals(&_PositionLens.CallOpts)
}

// DebtDecimals is a free data retrieval call binding the contract method 0xf341e228.
//
// Solidity: function debtDecimals() view returns(uint8)
func (_PositionLens *PositionLensCallerSession) DebtDecimals() (uint8, error) {
	return _PositionLens.Contract.DebtDecimals(&_PositionLens.CallOpts)
}

// MarketData is a free data retrieval call binding the contract method 0xb945b527.
//
// Solidity: function marketData() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool) data)
func (_PositionLens *PositionLensCaller) MarketData(opts *bind.CallOpts) (MarketData, error) {
	var out []interface{}
	err := _PositionLens.contract.Call(opts, &out, "marketData")

	if err != nil {
		return *new(MarketData), err
	}

	out0 := *abi.ConvertType(out[0], new(MarketData)).(*MarketData)

	return out0, err

}

// MarketData is a free data retrieval call binding the contract method 0xb945b527.
//
// Solidity: function marketData() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool) data)
func (_PositionLens *PositionLensSession) MarketData() (MarketData, error) {
	return _PositionLens.Contract.MarketData(&_PositionLens.CallOpts)
}

// MarketData is a free data retrieval call binding the contract method 0xb945b527.
//
// Solidity: function marketData() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool) data)
func (_PositionLens *PositionLensCallerSession) MarketData() (MarketData, error) {
	return _PositionLens.Contract.MarketData(&_PositionLens.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PositionLens *PositionLensCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionLens.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PositionLens *PositionLensSession) Oracle() (common.Address, error) {
	return _PositionLens.Contract.Oracle(&_PositionLens.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PositionLens *PositionLensCallerSession) Oracle() (common.Address, error) {
	return _PositionLens.Contract.Oracle(&_PositionLens.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_PositionLens *PositionLensCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionLens.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_PositionLens *PositionLensSession) Pool() (common.Address, error) {
	return _PositionLens.Contract.Pool(&_PositionLens.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_PositionLens *PositionLensCallerSession) Pool() (common.Address, error) {
	return _PositionLens.Contract.Pool(&_PositionLens.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_PositionLens *PositionLensCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionLens.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_PositionLens *PositionLensSession) Vault() (common.Address, error) {
	return _PositionLens.Contract.Vault(&_PositionLens.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_PositionLens *PositionLensCallerSession) Vault() (common.Address, error) {
	return _PositionLens.Contract.Vault(&_PositionLens.CallOpts)
}
