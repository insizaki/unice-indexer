// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// UnicePredictionMarketMetaData contains all meta data concerning the UnicePredictionMarket contract.
var UnicePredictionMarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_question\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_category\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_bettingDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_resolutionDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHALLENGE_WINDOW\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bet\",\"inputs\":[{\"name\":\"isYes\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bettingDeadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"category\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challenge\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"challenged\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimed\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"expireClaim\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getClaimable\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"claimable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"_question\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_category\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_bettingDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_resolutionDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_totalYesPool\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_totalNoPool\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_outcome\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_finalized\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOdds\",\"inputs\":[],\"outputs\":[{\"name\":\"yesPct\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"noPct\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStatus\",\"inputs\":[],\"outputs\":[{\"name\":\"status\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"noShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outcome\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumUnicePredictionMarket.Outcome\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"question\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolutionDeadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitResult\",\"inputs\":[{\"name\":\"_outcome\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submittedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalNoPool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalYesPool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"yesShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BetPlaced\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isYes\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExpireRefundClaimed\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketFinalized\",\"inputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ResultChallenged\",\"inputs\":[{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ResultSubmitted\",\"inputs\":[{\"name\":\"outcome\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WinningsClaimed\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// UnicePredictionMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use UnicePredictionMarketMetaData.ABI instead.
var UnicePredictionMarketABI = UnicePredictionMarketMetaData.ABI

// UnicePredictionMarket is an auto generated Go binding around an Ethereum contract.
type UnicePredictionMarket struct {
	UnicePredictionMarketCaller     // Read-only binding to the contract
	UnicePredictionMarketTransactor // Write-only binding to the contract
	UnicePredictionMarketFilterer   // Log filterer for contract events
}

// UnicePredictionMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnicePredictionMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnicePredictionMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnicePredictionMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnicePredictionMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnicePredictionMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnicePredictionMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnicePredictionMarketSession struct {
	Contract     *UnicePredictionMarket // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UnicePredictionMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnicePredictionMarketCallerSession struct {
	Contract *UnicePredictionMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// UnicePredictionMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnicePredictionMarketTransactorSession struct {
	Contract     *UnicePredictionMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// UnicePredictionMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnicePredictionMarketRaw struct {
	Contract *UnicePredictionMarket // Generic contract binding to access the raw methods on
}

// UnicePredictionMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnicePredictionMarketCallerRaw struct {
	Contract *UnicePredictionMarketCaller // Generic read-only contract binding to access the raw methods on
}

// UnicePredictionMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnicePredictionMarketTransactorRaw struct {
	Contract *UnicePredictionMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnicePredictionMarket creates a new instance of UnicePredictionMarket, bound to a specific deployed contract.
func NewUnicePredictionMarket(address common.Address, backend bind.ContractBackend) (*UnicePredictionMarket, error) {
	contract, err := bindUnicePredictionMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnicePredictionMarket{UnicePredictionMarketCaller: UnicePredictionMarketCaller{contract: contract}, UnicePredictionMarketTransactor: UnicePredictionMarketTransactor{contract: contract}, UnicePredictionMarketFilterer: UnicePredictionMarketFilterer{contract: contract}}, nil
}

// NewUnicePredictionMarketCaller creates a new read-only instance of UnicePredictionMarket, bound to a specific deployed contract.
func NewUnicePredictionMarketCaller(address common.Address, caller bind.ContractCaller) (*UnicePredictionMarketCaller, error) {
	contract, err := bindUnicePredictionMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnicePredictionMarketCaller{contract: contract}, nil
}

// NewUnicePredictionMarketTransactor creates a new write-only instance of UnicePredictionMarket, bound to a specific deployed contract.
func NewUnicePredictionMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*UnicePredictionMarketTransactor, error) {
	contract, err := bindUnicePredictionMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnicePredictionMarketTransactor{contract: contract}, nil
}

// NewUnicePredictionMarketFilterer creates a new log filterer instance of UnicePredictionMarket, bound to a specific deployed contract.
func NewUnicePredictionMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*UnicePredictionMarketFilterer, error) {
	contract, err := bindUnicePredictionMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnicePredictionMarketFilterer{contract: contract}, nil
}

// bindUnicePredictionMarket binds a generic wrapper to an already deployed contract.
func bindUnicePredictionMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UnicePredictionMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnicePredictionMarket *UnicePredictionMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnicePredictionMarket.Contract.UnicePredictionMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnicePredictionMarket *UnicePredictionMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.UnicePredictionMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnicePredictionMarket *UnicePredictionMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.UnicePredictionMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnicePredictionMarket *UnicePredictionMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnicePredictionMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnicePredictionMarket *UnicePredictionMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnicePredictionMarket *UnicePredictionMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.contract.Transact(opts, method, params...)
}

// CHALLENGEWINDOW is a free data retrieval call binding the contract method 0xd62aad29.
//
// Solidity: function CHALLENGE_WINDOW() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) CHALLENGEWINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "CHALLENGE_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHALLENGEWINDOW is a free data retrieval call binding the contract method 0xd62aad29.
//
// Solidity: function CHALLENGE_WINDOW() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketSession) CHALLENGEWINDOW() (*big.Int, error) {
	return _UnicePredictionMarket.Contract.CHALLENGEWINDOW(&_UnicePredictionMarket.CallOpts)
}

// CHALLENGEWINDOW is a free data retrieval call binding the contract method 0xd62aad29.
//
// Solidity: function CHALLENGE_WINDOW() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) CHALLENGEWINDOW() (*big.Int, error) {
	return _UnicePredictionMarket.Contract.CHALLENGEWINDOW(&_UnicePredictionMarket.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_UnicePredictionMarket *UnicePredictionMarketSession) Admin() (common.Address, error) {
	return _UnicePredictionMarket.Contract.Admin(&_UnicePredictionMarket.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) Admin() (common.Address, error) {
	return _UnicePredictionMarket.Contract.Admin(&_UnicePredictionMarket.CallOpts)
}

// BettingDeadline is a free data retrieval call binding the contract method 0x2468698d.
//
// Solidity: function bettingDeadline() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) BettingDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "bettingDeadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BettingDeadline is a free data retrieval call binding the contract method 0x2468698d.
//
// Solidity: function bettingDeadline() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketSession) BettingDeadline() (*big.Int, error) {
	return _UnicePredictionMarket.Contract.BettingDeadline(&_UnicePredictionMarket.CallOpts)
}

// BettingDeadline is a free data retrieval call binding the contract method 0x2468698d.
//
// Solidity: function bettingDeadline() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) BettingDeadline() (*big.Int, error) {
	return _UnicePredictionMarket.Contract.BettingDeadline(&_UnicePredictionMarket.CallOpts)
}

// Category is a free data retrieval call binding the contract method 0xef430aa6.
//
// Solidity: function category() view returns(string)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) Category(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "category")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Category is a free data retrieval call binding the contract method 0xef430aa6.
//
// Solidity: function category() view returns(string)
func (_UnicePredictionMarket *UnicePredictionMarketSession) Category() (string, error) {
	return _UnicePredictionMarket.Contract.Category(&_UnicePredictionMarket.CallOpts)
}

// Category is a free data retrieval call binding the contract method 0xef430aa6.
//
// Solidity: function category() view returns(string)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) Category() (string, error) {
	return _UnicePredictionMarket.Contract.Category(&_UnicePredictionMarket.CallOpts)
}

// Challenged is a free data retrieval call binding the contract method 0x689006a5.
//
// Solidity: function challenged() view returns(bool)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) Challenged(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "challenged")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Challenged is a free data retrieval call binding the contract method 0x689006a5.
//
// Solidity: function challenged() view returns(bool)
func (_UnicePredictionMarket *UnicePredictionMarketSession) Challenged() (bool, error) {
	return _UnicePredictionMarket.Contract.Challenged(&_UnicePredictionMarket.CallOpts)
}

// Challenged is a free data retrieval call binding the contract method 0x689006a5.
//
// Solidity: function challenged() view returns(bool)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) Challenged() (bool, error) {
	return _UnicePredictionMarket.Contract.Challenged(&_UnicePredictionMarket.CallOpts)
}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(bool)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) Claimed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "claimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(bool)
func (_UnicePredictionMarket *UnicePredictionMarketSession) Claimed(arg0 common.Address) (bool, error) {
	return _UnicePredictionMarket.Contract.Claimed(&_UnicePredictionMarket.CallOpts, arg0)
}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(bool)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) Claimed(arg0 common.Address) (bool, error) {
	return _UnicePredictionMarket.Contract.Claimed(&_UnicePredictionMarket.CallOpts, arg0)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "collateralToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_UnicePredictionMarket *UnicePredictionMarketSession) CollateralToken() (common.Address, error) {
	return _UnicePredictionMarket.Contract.CollateralToken(&_UnicePredictionMarket.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) CollateralToken() (common.Address, error) {
	return _UnicePredictionMarket.Contract.CollateralToken(&_UnicePredictionMarket.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) Finalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "finalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_UnicePredictionMarket *UnicePredictionMarketSession) Finalized() (bool, error) {
	return _UnicePredictionMarket.Contract.Finalized(&_UnicePredictionMarket.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) Finalized() (bool, error) {
	return _UnicePredictionMarket.Contract.Finalized(&_UnicePredictionMarket.CallOpts)
}

// GetClaimable is a free data retrieval call binding the contract method 0xa583024b.
//
// Solidity: function getClaimable(address user) view returns(uint256 claimable)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) GetClaimable(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "getClaimable", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimable is a free data retrieval call binding the contract method 0xa583024b.
//
// Solidity: function getClaimable(address user) view returns(uint256 claimable)
func (_UnicePredictionMarket *UnicePredictionMarketSession) GetClaimable(user common.Address) (*big.Int, error) {
	return _UnicePredictionMarket.Contract.GetClaimable(&_UnicePredictionMarket.CallOpts, user)
}

// GetClaimable is a free data retrieval call binding the contract method 0xa583024b.
//
// Solidity: function getClaimable(address user) view returns(uint256 claimable)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) GetClaimable(user common.Address) (*big.Int, error) {
	return _UnicePredictionMarket.Contract.GetClaimable(&_UnicePredictionMarket.CallOpts, user)
}

// GetMarketInfo is a free data retrieval call binding the contract method 0x23341a05.
//
// Solidity: function getMarketInfo() view returns(string _question, string _category, uint256 _bettingDeadline, uint256 _resolutionDeadline, uint256 _totalYesPool, uint256 _totalNoPool, uint8 _outcome, bool _finalized)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) GetMarketInfo(opts *bind.CallOpts) (struct {
	Question           string
	Category           string
	BettingDeadline    *big.Int
	ResolutionDeadline *big.Int
	TotalYesPool       *big.Int
	TotalNoPool        *big.Int
	Outcome            uint8
	Finalized          bool
}, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "getMarketInfo")

	outstruct := new(struct {
		Question           string
		Category           string
		BettingDeadline    *big.Int
		ResolutionDeadline *big.Int
		TotalYesPool       *big.Int
		TotalNoPool        *big.Int
		Outcome            uint8
		Finalized          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Question = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Category = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.BettingDeadline = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ResolutionDeadline = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalYesPool = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalNoPool = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Outcome = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.Finalized = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// GetMarketInfo is a free data retrieval call binding the contract method 0x23341a05.
//
// Solidity: function getMarketInfo() view returns(string _question, string _category, uint256 _bettingDeadline, uint256 _resolutionDeadline, uint256 _totalYesPool, uint256 _totalNoPool, uint8 _outcome, bool _finalized)
func (_UnicePredictionMarket *UnicePredictionMarketSession) GetMarketInfo() (struct {
	Question           string
	Category           string
	BettingDeadline    *big.Int
	ResolutionDeadline *big.Int
	TotalYesPool       *big.Int
	TotalNoPool        *big.Int
	Outcome            uint8
	Finalized          bool
}, error) {
	return _UnicePredictionMarket.Contract.GetMarketInfo(&_UnicePredictionMarket.CallOpts)
}

// GetMarketInfo is a free data retrieval call binding the contract method 0x23341a05.
//
// Solidity: function getMarketInfo() view returns(string _question, string _category, uint256 _bettingDeadline, uint256 _resolutionDeadline, uint256 _totalYesPool, uint256 _totalNoPool, uint8 _outcome, bool _finalized)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) GetMarketInfo() (struct {
	Question           string
	Category           string
	BettingDeadline    *big.Int
	ResolutionDeadline *big.Int
	TotalYesPool       *big.Int
	TotalNoPool        *big.Int
	Outcome            uint8
	Finalized          bool
}, error) {
	return _UnicePredictionMarket.Contract.GetMarketInfo(&_UnicePredictionMarket.CallOpts)
}

// GetOdds is a free data retrieval call binding the contract method 0x4c5be574.
//
// Solidity: function getOdds() view returns(uint256 yesPct, uint256 noPct)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) GetOdds(opts *bind.CallOpts) (struct {
	YesPct *big.Int
	NoPct  *big.Int
}, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "getOdds")

	outstruct := new(struct {
		YesPct *big.Int
		NoPct  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.YesPct = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NoPct = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOdds is a free data retrieval call binding the contract method 0x4c5be574.
//
// Solidity: function getOdds() view returns(uint256 yesPct, uint256 noPct)
func (_UnicePredictionMarket *UnicePredictionMarketSession) GetOdds() (struct {
	YesPct *big.Int
	NoPct  *big.Int
}, error) {
	return _UnicePredictionMarket.Contract.GetOdds(&_UnicePredictionMarket.CallOpts)
}

// GetOdds is a free data retrieval call binding the contract method 0x4c5be574.
//
// Solidity: function getOdds() view returns(uint256 yesPct, uint256 noPct)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) GetOdds() (struct {
	YesPct *big.Int
	NoPct  *big.Int
}, error) {
	return _UnicePredictionMarket.Contract.GetOdds(&_UnicePredictionMarket.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(string status)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) GetStatus(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "getStatus")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(string status)
func (_UnicePredictionMarket *UnicePredictionMarketSession) GetStatus() (string, error) {
	return _UnicePredictionMarket.Contract.GetStatus(&_UnicePredictionMarket.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(string status)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) GetStatus() (string, error) {
	return _UnicePredictionMarket.Contract.GetStatus(&_UnicePredictionMarket.CallOpts)
}

// NoShares is a free data retrieval call binding the contract method 0x56bc793e.
//
// Solidity: function noShares(address ) view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) NoShares(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "noShares", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoShares is a free data retrieval call binding the contract method 0x56bc793e.
//
// Solidity: function noShares(address ) view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketSession) NoShares(arg0 common.Address) (*big.Int, error) {
	return _UnicePredictionMarket.Contract.NoShares(&_UnicePredictionMarket.CallOpts, arg0)
}

// NoShares is a free data retrieval call binding the contract method 0x56bc793e.
//
// Solidity: function noShares(address ) view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) NoShares(arg0 common.Address) (*big.Int, error) {
	return _UnicePredictionMarket.Contract.NoShares(&_UnicePredictionMarket.CallOpts, arg0)
}

// Outcome is a free data retrieval call binding the contract method 0x27793f87.
//
// Solidity: function outcome() view returns(uint8)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) Outcome(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "outcome")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Outcome is a free data retrieval call binding the contract method 0x27793f87.
//
// Solidity: function outcome() view returns(uint8)
func (_UnicePredictionMarket *UnicePredictionMarketSession) Outcome() (uint8, error) {
	return _UnicePredictionMarket.Contract.Outcome(&_UnicePredictionMarket.CallOpts)
}

// Outcome is a free data retrieval call binding the contract method 0x27793f87.
//
// Solidity: function outcome() view returns(uint8)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) Outcome() (uint8, error) {
	return _UnicePredictionMarket.Contract.Outcome(&_UnicePredictionMarket.CallOpts)
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() view returns(string)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) Question(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "question")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() view returns(string)
func (_UnicePredictionMarket *UnicePredictionMarketSession) Question() (string, error) {
	return _UnicePredictionMarket.Contract.Question(&_UnicePredictionMarket.CallOpts)
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() view returns(string)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) Question() (string, error) {
	return _UnicePredictionMarket.Contract.Question(&_UnicePredictionMarket.CallOpts)
}

// ResolutionDeadline is a free data retrieval call binding the contract method 0x528f734f.
//
// Solidity: function resolutionDeadline() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) ResolutionDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "resolutionDeadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ResolutionDeadline is a free data retrieval call binding the contract method 0x528f734f.
//
// Solidity: function resolutionDeadline() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketSession) ResolutionDeadline() (*big.Int, error) {
	return _UnicePredictionMarket.Contract.ResolutionDeadline(&_UnicePredictionMarket.CallOpts)
}

// ResolutionDeadline is a free data retrieval call binding the contract method 0x528f734f.
//
// Solidity: function resolutionDeadline() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) ResolutionDeadline() (*big.Int, error) {
	return _UnicePredictionMarket.Contract.ResolutionDeadline(&_UnicePredictionMarket.CallOpts)
}

// SubmittedAt is a free data retrieval call binding the contract method 0x1ca10799.
//
// Solidity: function submittedAt() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) SubmittedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "submittedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmittedAt is a free data retrieval call binding the contract method 0x1ca10799.
//
// Solidity: function submittedAt() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketSession) SubmittedAt() (*big.Int, error) {
	return _UnicePredictionMarket.Contract.SubmittedAt(&_UnicePredictionMarket.CallOpts)
}

// SubmittedAt is a free data retrieval call binding the contract method 0x1ca10799.
//
// Solidity: function submittedAt() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) SubmittedAt() (*big.Int, error) {
	return _UnicePredictionMarket.Contract.SubmittedAt(&_UnicePredictionMarket.CallOpts)
}

// TotalNoPool is a free data retrieval call binding the contract method 0x6c8d6e5a.
//
// Solidity: function totalNoPool() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) TotalNoPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "totalNoPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNoPool is a free data retrieval call binding the contract method 0x6c8d6e5a.
//
// Solidity: function totalNoPool() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketSession) TotalNoPool() (*big.Int, error) {
	return _UnicePredictionMarket.Contract.TotalNoPool(&_UnicePredictionMarket.CallOpts)
}

// TotalNoPool is a free data retrieval call binding the contract method 0x6c8d6e5a.
//
// Solidity: function totalNoPool() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) TotalNoPool() (*big.Int, error) {
	return _UnicePredictionMarket.Contract.TotalNoPool(&_UnicePredictionMarket.CallOpts)
}

// TotalYesPool is a free data retrieval call binding the contract method 0x35b185b2.
//
// Solidity: function totalYesPool() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) TotalYesPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "totalYesPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalYesPool is a free data retrieval call binding the contract method 0x35b185b2.
//
// Solidity: function totalYesPool() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketSession) TotalYesPool() (*big.Int, error) {
	return _UnicePredictionMarket.Contract.TotalYesPool(&_UnicePredictionMarket.CallOpts)
}

// TotalYesPool is a free data retrieval call binding the contract method 0x35b185b2.
//
// Solidity: function totalYesPool() view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) TotalYesPool() (*big.Int, error) {
	return _UnicePredictionMarket.Contract.TotalYesPool(&_UnicePredictionMarket.CallOpts)
}

// YesShares is a free data retrieval call binding the contract method 0x818fa941.
//
// Solidity: function yesShares(address ) view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCaller) YesShares(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnicePredictionMarket.contract.Call(opts, &out, "yesShares", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YesShares is a free data retrieval call binding the contract method 0x818fa941.
//
// Solidity: function yesShares(address ) view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketSession) YesShares(arg0 common.Address) (*big.Int, error) {
	return _UnicePredictionMarket.Contract.YesShares(&_UnicePredictionMarket.CallOpts, arg0)
}

// YesShares is a free data retrieval call binding the contract method 0x818fa941.
//
// Solidity: function yesShares(address ) view returns(uint256)
func (_UnicePredictionMarket *UnicePredictionMarketCallerSession) YesShares(arg0 common.Address) (*big.Int, error) {
	return _UnicePredictionMarket.Contract.YesShares(&_UnicePredictionMarket.CallOpts, arg0)
}

// Bet is a paid mutator transaction binding the contract method 0x5229a6fd.
//
// Solidity: function bet(bool isYes, uint256 amount) returns()
func (_UnicePredictionMarket *UnicePredictionMarketTransactor) Bet(opts *bind.TransactOpts, isYes bool, amount *big.Int) (*types.Transaction, error) {
	return _UnicePredictionMarket.contract.Transact(opts, "bet", isYes, amount)
}

// Bet is a paid mutator transaction binding the contract method 0x5229a6fd.
//
// Solidity: function bet(bool isYes, uint256 amount) returns()
func (_UnicePredictionMarket *UnicePredictionMarketSession) Bet(isYes bool, amount *big.Int) (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.Bet(&_UnicePredictionMarket.TransactOpts, isYes, amount)
}

// Bet is a paid mutator transaction binding the contract method 0x5229a6fd.
//
// Solidity: function bet(bool isYes, uint256 amount) returns()
func (_UnicePredictionMarket *UnicePredictionMarketTransactorSession) Bet(isYes bool, amount *big.Int) (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.Bet(&_UnicePredictionMarket.TransactOpts, isYes, amount)
}

// Challenge is a paid mutator transaction binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() returns()
func (_UnicePredictionMarket *UnicePredictionMarketTransactor) Challenge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnicePredictionMarket.contract.Transact(opts, "challenge")
}

// Challenge is a paid mutator transaction binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() returns()
func (_UnicePredictionMarket *UnicePredictionMarketSession) Challenge() (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.Challenge(&_UnicePredictionMarket.TransactOpts)
}

// Challenge is a paid mutator transaction binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() returns()
func (_UnicePredictionMarket *UnicePredictionMarketTransactorSession) Challenge() (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.Challenge(&_UnicePredictionMarket.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_UnicePredictionMarket *UnicePredictionMarketTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnicePredictionMarket.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_UnicePredictionMarket *UnicePredictionMarketSession) Claim() (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.Claim(&_UnicePredictionMarket.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_UnicePredictionMarket *UnicePredictionMarketTransactorSession) Claim() (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.Claim(&_UnicePredictionMarket.TransactOpts)
}

// ExpireClaim is a paid mutator transaction binding the contract method 0x1257d2c8.
//
// Solidity: function expireClaim() returns()
func (_UnicePredictionMarket *UnicePredictionMarketTransactor) ExpireClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnicePredictionMarket.contract.Transact(opts, "expireClaim")
}

// ExpireClaim is a paid mutator transaction binding the contract method 0x1257d2c8.
//
// Solidity: function expireClaim() returns()
func (_UnicePredictionMarket *UnicePredictionMarketSession) ExpireClaim() (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.ExpireClaim(&_UnicePredictionMarket.TransactOpts)
}

// ExpireClaim is a paid mutator transaction binding the contract method 0x1257d2c8.
//
// Solidity: function expireClaim() returns()
func (_UnicePredictionMarket *UnicePredictionMarketTransactorSession) ExpireClaim() (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.ExpireClaim(&_UnicePredictionMarket.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_UnicePredictionMarket *UnicePredictionMarketTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnicePredictionMarket.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_UnicePredictionMarket *UnicePredictionMarketSession) Finalize() (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.Finalize(&_UnicePredictionMarket.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_UnicePredictionMarket *UnicePredictionMarketTransactorSession) Finalize() (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.Finalize(&_UnicePredictionMarket.TransactOpts)
}

// SubmitResult is a paid mutator transaction binding the contract method 0x448807fe.
//
// Solidity: function submitResult(uint8 _outcome) returns()
func (_UnicePredictionMarket *UnicePredictionMarketTransactor) SubmitResult(opts *bind.TransactOpts, _outcome uint8) (*types.Transaction, error) {
	return _UnicePredictionMarket.contract.Transact(opts, "submitResult", _outcome)
}

// SubmitResult is a paid mutator transaction binding the contract method 0x448807fe.
//
// Solidity: function submitResult(uint8 _outcome) returns()
func (_UnicePredictionMarket *UnicePredictionMarketSession) SubmitResult(_outcome uint8) (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.SubmitResult(&_UnicePredictionMarket.TransactOpts, _outcome)
}

// SubmitResult is a paid mutator transaction binding the contract method 0x448807fe.
//
// Solidity: function submitResult(uint8 _outcome) returns()
func (_UnicePredictionMarket *UnicePredictionMarketTransactorSession) SubmitResult(_outcome uint8) (*types.Transaction, error) {
	return _UnicePredictionMarket.Contract.SubmitResult(&_UnicePredictionMarket.TransactOpts, _outcome)
}

// UnicePredictionMarketBetPlacedIterator is returned from FilterBetPlaced and is used to iterate over the raw logs and unpacked data for BetPlaced events raised by the UnicePredictionMarket contract.
type UnicePredictionMarketBetPlacedIterator struct {
	Event *UnicePredictionMarketBetPlaced // Event containing the contract specifics and raw log

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
func (it *UnicePredictionMarketBetPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicePredictionMarketBetPlaced)
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
		it.Event = new(UnicePredictionMarketBetPlaced)
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
func (it *UnicePredictionMarketBetPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicePredictionMarketBetPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicePredictionMarketBetPlaced represents a BetPlaced event raised by the UnicePredictionMarket contract.
type UnicePredictionMarketBetPlaced struct {
	User   common.Address
	IsYes  bool
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBetPlaced is a free log retrieval operation binding the contract event 0xe57fa5a543ebcdebf2a7d7102ac913e14454cec006468e43a6916e95bdfcea9e.
//
// Solidity: event BetPlaced(address indexed user, bool isYes, uint256 amount)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) FilterBetPlaced(opts *bind.FilterOpts, user []common.Address) (*UnicePredictionMarketBetPlacedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UnicePredictionMarket.contract.FilterLogs(opts, "BetPlaced", userRule)
	if err != nil {
		return nil, err
	}
	return &UnicePredictionMarketBetPlacedIterator{contract: _UnicePredictionMarket.contract, event: "BetPlaced", logs: logs, sub: sub}, nil
}

// WatchBetPlaced is a free log subscription operation binding the contract event 0xe57fa5a543ebcdebf2a7d7102ac913e14454cec006468e43a6916e95bdfcea9e.
//
// Solidity: event BetPlaced(address indexed user, bool isYes, uint256 amount)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) WatchBetPlaced(opts *bind.WatchOpts, sink chan<- *UnicePredictionMarketBetPlaced, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UnicePredictionMarket.contract.WatchLogs(opts, "BetPlaced", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicePredictionMarketBetPlaced)
				if err := _UnicePredictionMarket.contract.UnpackLog(event, "BetPlaced", log); err != nil {
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

// ParseBetPlaced is a log parse operation binding the contract event 0xe57fa5a543ebcdebf2a7d7102ac913e14454cec006468e43a6916e95bdfcea9e.
//
// Solidity: event BetPlaced(address indexed user, bool isYes, uint256 amount)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) ParseBetPlaced(log types.Log) (*UnicePredictionMarketBetPlaced, error) {
	event := new(UnicePredictionMarketBetPlaced)
	if err := _UnicePredictionMarket.contract.UnpackLog(event, "BetPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnicePredictionMarketExpireRefundClaimedIterator is returned from FilterExpireRefundClaimed and is used to iterate over the raw logs and unpacked data for ExpireRefundClaimed events raised by the UnicePredictionMarket contract.
type UnicePredictionMarketExpireRefundClaimedIterator struct {
	Event *UnicePredictionMarketExpireRefundClaimed // Event containing the contract specifics and raw log

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
func (it *UnicePredictionMarketExpireRefundClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicePredictionMarketExpireRefundClaimed)
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
		it.Event = new(UnicePredictionMarketExpireRefundClaimed)
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
func (it *UnicePredictionMarketExpireRefundClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicePredictionMarketExpireRefundClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicePredictionMarketExpireRefundClaimed represents a ExpireRefundClaimed event raised by the UnicePredictionMarket contract.
type UnicePredictionMarketExpireRefundClaimed struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExpireRefundClaimed is a free log retrieval operation binding the contract event 0xc5a2a463ffe81791cb601ca8a98e672d33989bf796a198f90da1b0be9e647e0e.
//
// Solidity: event ExpireRefundClaimed(address indexed user, uint256 amount)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) FilterExpireRefundClaimed(opts *bind.FilterOpts, user []common.Address) (*UnicePredictionMarketExpireRefundClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UnicePredictionMarket.contract.FilterLogs(opts, "ExpireRefundClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &UnicePredictionMarketExpireRefundClaimedIterator{contract: _UnicePredictionMarket.contract, event: "ExpireRefundClaimed", logs: logs, sub: sub}, nil
}

// WatchExpireRefundClaimed is a free log subscription operation binding the contract event 0xc5a2a463ffe81791cb601ca8a98e672d33989bf796a198f90da1b0be9e647e0e.
//
// Solidity: event ExpireRefundClaimed(address indexed user, uint256 amount)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) WatchExpireRefundClaimed(opts *bind.WatchOpts, sink chan<- *UnicePredictionMarketExpireRefundClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UnicePredictionMarket.contract.WatchLogs(opts, "ExpireRefundClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicePredictionMarketExpireRefundClaimed)
				if err := _UnicePredictionMarket.contract.UnpackLog(event, "ExpireRefundClaimed", log); err != nil {
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

// ParseExpireRefundClaimed is a log parse operation binding the contract event 0xc5a2a463ffe81791cb601ca8a98e672d33989bf796a198f90da1b0be9e647e0e.
//
// Solidity: event ExpireRefundClaimed(address indexed user, uint256 amount)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) ParseExpireRefundClaimed(log types.Log) (*UnicePredictionMarketExpireRefundClaimed, error) {
	event := new(UnicePredictionMarketExpireRefundClaimed)
	if err := _UnicePredictionMarket.contract.UnpackLog(event, "ExpireRefundClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnicePredictionMarketMarketFinalizedIterator is returned from FilterMarketFinalized and is used to iterate over the raw logs and unpacked data for MarketFinalized events raised by the UnicePredictionMarket contract.
type UnicePredictionMarketMarketFinalizedIterator struct {
	Event *UnicePredictionMarketMarketFinalized // Event containing the contract specifics and raw log

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
func (it *UnicePredictionMarketMarketFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicePredictionMarketMarketFinalized)
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
		it.Event = new(UnicePredictionMarketMarketFinalized)
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
func (it *UnicePredictionMarketMarketFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicePredictionMarketMarketFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicePredictionMarketMarketFinalized represents a MarketFinalized event raised by the UnicePredictionMarket contract.
type UnicePredictionMarketMarketFinalized struct {
	Outcome uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketFinalized is a free log retrieval operation binding the contract event 0xe22647b990bb2c9f1fab2207a70363abb43ac42adeeaa3ec22f78ad1edf183a5.
//
// Solidity: event MarketFinalized(uint8 outcome)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) FilterMarketFinalized(opts *bind.FilterOpts) (*UnicePredictionMarketMarketFinalizedIterator, error) {

	logs, sub, err := _UnicePredictionMarket.contract.FilterLogs(opts, "MarketFinalized")
	if err != nil {
		return nil, err
	}
	return &UnicePredictionMarketMarketFinalizedIterator{contract: _UnicePredictionMarket.contract, event: "MarketFinalized", logs: logs, sub: sub}, nil
}

// WatchMarketFinalized is a free log subscription operation binding the contract event 0xe22647b990bb2c9f1fab2207a70363abb43ac42adeeaa3ec22f78ad1edf183a5.
//
// Solidity: event MarketFinalized(uint8 outcome)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) WatchMarketFinalized(opts *bind.WatchOpts, sink chan<- *UnicePredictionMarketMarketFinalized) (event.Subscription, error) {

	logs, sub, err := _UnicePredictionMarket.contract.WatchLogs(opts, "MarketFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicePredictionMarketMarketFinalized)
				if err := _UnicePredictionMarket.contract.UnpackLog(event, "MarketFinalized", log); err != nil {
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

// ParseMarketFinalized is a log parse operation binding the contract event 0xe22647b990bb2c9f1fab2207a70363abb43ac42adeeaa3ec22f78ad1edf183a5.
//
// Solidity: event MarketFinalized(uint8 outcome)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) ParseMarketFinalized(log types.Log) (*UnicePredictionMarketMarketFinalized, error) {
	event := new(UnicePredictionMarketMarketFinalized)
	if err := _UnicePredictionMarket.contract.UnpackLog(event, "MarketFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnicePredictionMarketResultChallengedIterator is returned from FilterResultChallenged and is used to iterate over the raw logs and unpacked data for ResultChallenged events raised by the UnicePredictionMarket contract.
type UnicePredictionMarketResultChallengedIterator struct {
	Event *UnicePredictionMarketResultChallenged // Event containing the contract specifics and raw log

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
func (it *UnicePredictionMarketResultChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicePredictionMarketResultChallenged)
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
		it.Event = new(UnicePredictionMarketResultChallenged)
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
func (it *UnicePredictionMarketResultChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicePredictionMarketResultChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicePredictionMarketResultChallenged represents a ResultChallenged event raised by the UnicePredictionMarket contract.
type UnicePredictionMarketResultChallenged struct {
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterResultChallenged is a free log retrieval operation binding the contract event 0x9b9e39a8a790f8ce157ab4671a2c9557b9a3f395e93bc8fe89f89009034f00b3.
//
// Solidity: event ResultChallenged(address indexed challenger)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) FilterResultChallenged(opts *bind.FilterOpts, challenger []common.Address) (*UnicePredictionMarketResultChallengedIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _UnicePredictionMarket.contract.FilterLogs(opts, "ResultChallenged", challengerRule)
	if err != nil {
		return nil, err
	}
	return &UnicePredictionMarketResultChallengedIterator{contract: _UnicePredictionMarket.contract, event: "ResultChallenged", logs: logs, sub: sub}, nil
}

// WatchResultChallenged is a free log subscription operation binding the contract event 0x9b9e39a8a790f8ce157ab4671a2c9557b9a3f395e93bc8fe89f89009034f00b3.
//
// Solidity: event ResultChallenged(address indexed challenger)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) WatchResultChallenged(opts *bind.WatchOpts, sink chan<- *UnicePredictionMarketResultChallenged, challenger []common.Address) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _UnicePredictionMarket.contract.WatchLogs(opts, "ResultChallenged", challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicePredictionMarketResultChallenged)
				if err := _UnicePredictionMarket.contract.UnpackLog(event, "ResultChallenged", log); err != nil {
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

// ParseResultChallenged is a log parse operation binding the contract event 0x9b9e39a8a790f8ce157ab4671a2c9557b9a3f395e93bc8fe89f89009034f00b3.
//
// Solidity: event ResultChallenged(address indexed challenger)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) ParseResultChallenged(log types.Log) (*UnicePredictionMarketResultChallenged, error) {
	event := new(UnicePredictionMarketResultChallenged)
	if err := _UnicePredictionMarket.contract.UnpackLog(event, "ResultChallenged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnicePredictionMarketResultSubmittedIterator is returned from FilterResultSubmitted and is used to iterate over the raw logs and unpacked data for ResultSubmitted events raised by the UnicePredictionMarket contract.
type UnicePredictionMarketResultSubmittedIterator struct {
	Event *UnicePredictionMarketResultSubmitted // Event containing the contract specifics and raw log

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
func (it *UnicePredictionMarketResultSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicePredictionMarketResultSubmitted)
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
		it.Event = new(UnicePredictionMarketResultSubmitted)
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
func (it *UnicePredictionMarketResultSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicePredictionMarketResultSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicePredictionMarketResultSubmitted represents a ResultSubmitted event raised by the UnicePredictionMarket contract.
type UnicePredictionMarketResultSubmitted struct {
	Outcome   uint8
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResultSubmitted is a free log retrieval operation binding the contract event 0x059c8c9182ed93df36fc771341d64e3bd86b469eb45a1f5d52d38fe6ecdff6f2.
//
// Solidity: event ResultSubmitted(uint8 outcome, uint256 timestamp)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) FilterResultSubmitted(opts *bind.FilterOpts) (*UnicePredictionMarketResultSubmittedIterator, error) {

	logs, sub, err := _UnicePredictionMarket.contract.FilterLogs(opts, "ResultSubmitted")
	if err != nil {
		return nil, err
	}
	return &UnicePredictionMarketResultSubmittedIterator{contract: _UnicePredictionMarket.contract, event: "ResultSubmitted", logs: logs, sub: sub}, nil
}

// WatchResultSubmitted is a free log subscription operation binding the contract event 0x059c8c9182ed93df36fc771341d64e3bd86b469eb45a1f5d52d38fe6ecdff6f2.
//
// Solidity: event ResultSubmitted(uint8 outcome, uint256 timestamp)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) WatchResultSubmitted(opts *bind.WatchOpts, sink chan<- *UnicePredictionMarketResultSubmitted) (event.Subscription, error) {

	logs, sub, err := _UnicePredictionMarket.contract.WatchLogs(opts, "ResultSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicePredictionMarketResultSubmitted)
				if err := _UnicePredictionMarket.contract.UnpackLog(event, "ResultSubmitted", log); err != nil {
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

// ParseResultSubmitted is a log parse operation binding the contract event 0x059c8c9182ed93df36fc771341d64e3bd86b469eb45a1f5d52d38fe6ecdff6f2.
//
// Solidity: event ResultSubmitted(uint8 outcome, uint256 timestamp)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) ParseResultSubmitted(log types.Log) (*UnicePredictionMarketResultSubmitted, error) {
	event := new(UnicePredictionMarketResultSubmitted)
	if err := _UnicePredictionMarket.contract.UnpackLog(event, "ResultSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnicePredictionMarketWinningsClaimedIterator is returned from FilterWinningsClaimed and is used to iterate over the raw logs and unpacked data for WinningsClaimed events raised by the UnicePredictionMarket contract.
type UnicePredictionMarketWinningsClaimedIterator struct {
	Event *UnicePredictionMarketWinningsClaimed // Event containing the contract specifics and raw log

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
func (it *UnicePredictionMarketWinningsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicePredictionMarketWinningsClaimed)
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
		it.Event = new(UnicePredictionMarketWinningsClaimed)
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
func (it *UnicePredictionMarketWinningsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicePredictionMarketWinningsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicePredictionMarketWinningsClaimed represents a WinningsClaimed event raised by the UnicePredictionMarket contract.
type UnicePredictionMarketWinningsClaimed struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWinningsClaimed is a free log retrieval operation binding the contract event 0x1a31e733a172afcf46074b3106c17f0c298e226442682a03c1e99ce256139ec2.
//
// Solidity: event WinningsClaimed(address indexed user, uint256 amount)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) FilterWinningsClaimed(opts *bind.FilterOpts, user []common.Address) (*UnicePredictionMarketWinningsClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UnicePredictionMarket.contract.FilterLogs(opts, "WinningsClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &UnicePredictionMarketWinningsClaimedIterator{contract: _UnicePredictionMarket.contract, event: "WinningsClaimed", logs: logs, sub: sub}, nil
}

// WatchWinningsClaimed is a free log subscription operation binding the contract event 0x1a31e733a172afcf46074b3106c17f0c298e226442682a03c1e99ce256139ec2.
//
// Solidity: event WinningsClaimed(address indexed user, uint256 amount)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) WatchWinningsClaimed(opts *bind.WatchOpts, sink chan<- *UnicePredictionMarketWinningsClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UnicePredictionMarket.contract.WatchLogs(opts, "WinningsClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicePredictionMarketWinningsClaimed)
				if err := _UnicePredictionMarket.contract.UnpackLog(event, "WinningsClaimed", log); err != nil {
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

// ParseWinningsClaimed is a log parse operation binding the contract event 0x1a31e733a172afcf46074b3106c17f0c298e226442682a03c1e99ce256139ec2.
//
// Solidity: event WinningsClaimed(address indexed user, uint256 amount)
func (_UnicePredictionMarket *UnicePredictionMarketFilterer) ParseWinningsClaimed(log types.Log) (*UnicePredictionMarketWinningsClaimed, error) {
	event := new(UnicePredictionMarketWinningsClaimed)
	if err := _UnicePredictionMarket.contract.UnpackLog(event, "WinningsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
