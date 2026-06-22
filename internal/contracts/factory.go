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

// UniceFactoryMetaData contains all meta data concerning the UniceFactory contract.
var UniceFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allMarkets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createMarket\",\"inputs\":[{\"name\":\"question\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"category\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"bettingDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolutionDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"marketAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllMarkets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketsByCategory\",\"inputs\":[{\"name\":\"category\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalMarkets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMarket\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketsByCategory\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MarketCreated\",\"inputs\":[{\"name\":\"market\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"question\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"category\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"bettingDeadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"resolutionDeadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// UniceFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use UniceFactoryMetaData.ABI instead.
var UniceFactoryABI = UniceFactoryMetaData.ABI

// UniceFactory is an auto generated Go binding around an Ethereum contract.
type UniceFactory struct {
	UniceFactoryCaller     // Read-only binding to the contract
	UniceFactoryTransactor // Write-only binding to the contract
	UniceFactoryFilterer   // Log filterer for contract events
}

// UniceFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniceFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniceFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniceFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniceFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniceFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniceFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniceFactorySession struct {
	Contract     *UniceFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniceFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniceFactoryCallerSession struct {
	Contract *UniceFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// UniceFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniceFactoryTransactorSession struct {
	Contract     *UniceFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UniceFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniceFactoryRaw struct {
	Contract *UniceFactory // Generic contract binding to access the raw methods on
}

// UniceFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniceFactoryCallerRaw struct {
	Contract *UniceFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// UniceFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniceFactoryTransactorRaw struct {
	Contract *UniceFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniceFactory creates a new instance of UniceFactory, bound to a specific deployed contract.
func NewUniceFactory(address common.Address, backend bind.ContractBackend) (*UniceFactory, error) {
	contract, err := bindUniceFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniceFactory{UniceFactoryCaller: UniceFactoryCaller{contract: contract}, UniceFactoryTransactor: UniceFactoryTransactor{contract: contract}, UniceFactoryFilterer: UniceFactoryFilterer{contract: contract}}, nil
}

// NewUniceFactoryCaller creates a new read-only instance of UniceFactory, bound to a specific deployed contract.
func NewUniceFactoryCaller(address common.Address, caller bind.ContractCaller) (*UniceFactoryCaller, error) {
	contract, err := bindUniceFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniceFactoryCaller{contract: contract}, nil
}

// NewUniceFactoryTransactor creates a new write-only instance of UniceFactory, bound to a specific deployed contract.
func NewUniceFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UniceFactoryTransactor, error) {
	contract, err := bindUniceFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniceFactoryTransactor{contract: contract}, nil
}

// NewUniceFactoryFilterer creates a new log filterer instance of UniceFactory, bound to a specific deployed contract.
func NewUniceFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*UniceFactoryFilterer, error) {
	contract, err := bindUniceFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniceFactoryFilterer{contract: contract}, nil
}

// bindUniceFactory binds a generic wrapper to an already deployed contract.
func bindUniceFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniceFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniceFactory *UniceFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniceFactory.Contract.UniceFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniceFactory *UniceFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniceFactory.Contract.UniceFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniceFactory *UniceFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniceFactory.Contract.UniceFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniceFactory *UniceFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniceFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniceFactory *UniceFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniceFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniceFactory *UniceFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniceFactory.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_UniceFactory *UniceFactoryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniceFactory.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_UniceFactory *UniceFactorySession) Admin() (common.Address, error) {
	return _UniceFactory.Contract.Admin(&_UniceFactory.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_UniceFactory *UniceFactoryCallerSession) Admin() (common.Address, error) {
	return _UniceFactory.Contract.Admin(&_UniceFactory.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_UniceFactory *UniceFactoryCaller) AllMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UniceFactory.contract.Call(opts, &out, "allMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_UniceFactory *UniceFactorySession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _UniceFactory.Contract.AllMarkets(&_UniceFactory.CallOpts, arg0)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_UniceFactory *UniceFactoryCallerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _UniceFactory.Contract.AllMarkets(&_UniceFactory.CallOpts, arg0)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_UniceFactory *UniceFactoryCaller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniceFactory.contract.Call(opts, &out, "collateralToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_UniceFactory *UniceFactorySession) CollateralToken() (common.Address, error) {
	return _UniceFactory.Contract.CollateralToken(&_UniceFactory.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_UniceFactory *UniceFactoryCallerSession) CollateralToken() (common.Address, error) {
	return _UniceFactory.Contract.CollateralToken(&_UniceFactory.CallOpts)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_UniceFactory *UniceFactoryCaller) GetAllMarkets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _UniceFactory.contract.Call(opts, &out, "getAllMarkets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_UniceFactory *UniceFactorySession) GetAllMarkets() ([]common.Address, error) {
	return _UniceFactory.Contract.GetAllMarkets(&_UniceFactory.CallOpts)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_UniceFactory *UniceFactoryCallerSession) GetAllMarkets() ([]common.Address, error) {
	return _UniceFactory.Contract.GetAllMarkets(&_UniceFactory.CallOpts)
}

// GetMarketsByCategory is a free data retrieval call binding the contract method 0x76e6caff.
//
// Solidity: function getMarketsByCategory(string category) view returns(address[])
func (_UniceFactory *UniceFactoryCaller) GetMarketsByCategory(opts *bind.CallOpts, category string) ([]common.Address, error) {
	var out []interface{}
	err := _UniceFactory.contract.Call(opts, &out, "getMarketsByCategory", category)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMarketsByCategory is a free data retrieval call binding the contract method 0x76e6caff.
//
// Solidity: function getMarketsByCategory(string category) view returns(address[])
func (_UniceFactory *UniceFactorySession) GetMarketsByCategory(category string) ([]common.Address, error) {
	return _UniceFactory.Contract.GetMarketsByCategory(&_UniceFactory.CallOpts, category)
}

// GetMarketsByCategory is a free data retrieval call binding the contract method 0x76e6caff.
//
// Solidity: function getMarketsByCategory(string category) view returns(address[])
func (_UniceFactory *UniceFactoryCallerSession) GetMarketsByCategory(category string) ([]common.Address, error) {
	return _UniceFactory.Contract.GetMarketsByCategory(&_UniceFactory.CallOpts, category)
}

// GetTotalMarkets is a free data retrieval call binding the contract method 0x81ebf209.
//
// Solidity: function getTotalMarkets() view returns(uint256)
func (_UniceFactory *UniceFactoryCaller) GetTotalMarkets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniceFactory.contract.Call(opts, &out, "getTotalMarkets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalMarkets is a free data retrieval call binding the contract method 0x81ebf209.
//
// Solidity: function getTotalMarkets() view returns(uint256)
func (_UniceFactory *UniceFactorySession) GetTotalMarkets() (*big.Int, error) {
	return _UniceFactory.Contract.GetTotalMarkets(&_UniceFactory.CallOpts)
}

// GetTotalMarkets is a free data retrieval call binding the contract method 0x81ebf209.
//
// Solidity: function getTotalMarkets() view returns(uint256)
func (_UniceFactory *UniceFactoryCallerSession) GetTotalMarkets() (*big.Int, error) {
	return _UniceFactory.Contract.GetTotalMarkets(&_UniceFactory.CallOpts)
}

// IsMarket is a free data retrieval call binding the contract method 0x6ec934da.
//
// Solidity: function isMarket(address ) view returns(bool)
func (_UniceFactory *UniceFactoryCaller) IsMarket(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _UniceFactory.contract.Call(opts, &out, "isMarket", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarket is a free data retrieval call binding the contract method 0x6ec934da.
//
// Solidity: function isMarket(address ) view returns(bool)
func (_UniceFactory *UniceFactorySession) IsMarket(arg0 common.Address) (bool, error) {
	return _UniceFactory.Contract.IsMarket(&_UniceFactory.CallOpts, arg0)
}

// IsMarket is a free data retrieval call binding the contract method 0x6ec934da.
//
// Solidity: function isMarket(address ) view returns(bool)
func (_UniceFactory *UniceFactoryCallerSession) IsMarket(arg0 common.Address) (bool, error) {
	return _UniceFactory.Contract.IsMarket(&_UniceFactory.CallOpts, arg0)
}

// MarketsByCategory is a free data retrieval call binding the contract method 0xaff14fe8.
//
// Solidity: function marketsByCategory(string , uint256 ) view returns(address)
func (_UniceFactory *UniceFactoryCaller) MarketsByCategory(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UniceFactory.contract.Call(opts, &out, "marketsByCategory", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketsByCategory is a free data retrieval call binding the contract method 0xaff14fe8.
//
// Solidity: function marketsByCategory(string , uint256 ) view returns(address)
func (_UniceFactory *UniceFactorySession) MarketsByCategory(arg0 string, arg1 *big.Int) (common.Address, error) {
	return _UniceFactory.Contract.MarketsByCategory(&_UniceFactory.CallOpts, arg0, arg1)
}

// MarketsByCategory is a free data retrieval call binding the contract method 0xaff14fe8.
//
// Solidity: function marketsByCategory(string , uint256 ) view returns(address)
func (_UniceFactory *UniceFactoryCallerSession) MarketsByCategory(arg0 string, arg1 *big.Int) (common.Address, error) {
	return _UniceFactory.Contract.MarketsByCategory(&_UniceFactory.CallOpts, arg0, arg1)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x39b6ee19.
//
// Solidity: function createMarket(string question, string category, uint256 bettingDeadline, uint256 resolutionDeadline) returns(address marketAddress)
func (_UniceFactory *UniceFactoryTransactor) CreateMarket(opts *bind.TransactOpts, question string, category string, bettingDeadline *big.Int, resolutionDeadline *big.Int) (*types.Transaction, error) {
	return _UniceFactory.contract.Transact(opts, "createMarket", question, category, bettingDeadline, resolutionDeadline)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x39b6ee19.
//
// Solidity: function createMarket(string question, string category, uint256 bettingDeadline, uint256 resolutionDeadline) returns(address marketAddress)
func (_UniceFactory *UniceFactorySession) CreateMarket(question string, category string, bettingDeadline *big.Int, resolutionDeadline *big.Int) (*types.Transaction, error) {
	return _UniceFactory.Contract.CreateMarket(&_UniceFactory.TransactOpts, question, category, bettingDeadline, resolutionDeadline)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x39b6ee19.
//
// Solidity: function createMarket(string question, string category, uint256 bettingDeadline, uint256 resolutionDeadline) returns(address marketAddress)
func (_UniceFactory *UniceFactoryTransactorSession) CreateMarket(question string, category string, bettingDeadline *big.Int, resolutionDeadline *big.Int) (*types.Transaction, error) {
	return _UniceFactory.Contract.CreateMarket(&_UniceFactory.TransactOpts, question, category, bettingDeadline, resolutionDeadline)
}

// UniceFactoryMarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the UniceFactory contract.
type UniceFactoryMarketCreatedIterator struct {
	Event *UniceFactoryMarketCreated // Event containing the contract specifics and raw log

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
func (it *UniceFactoryMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniceFactoryMarketCreated)
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
		it.Event = new(UniceFactoryMarketCreated)
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
func (it *UniceFactoryMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniceFactoryMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniceFactoryMarketCreated represents a MarketCreated event raised by the UniceFactory contract.
type UniceFactoryMarketCreated struct {
	Market             common.Address
	Question           string
	Category           string
	BettingDeadline    *big.Int
	ResolutionDeadline *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0xed229bcc1acd62d965455d23da5563e7082eb7006e172570c6dc39d189705b79.
//
// Solidity: event MarketCreated(address indexed market, string question, string category, uint256 bettingDeadline, uint256 resolutionDeadline)
func (_UniceFactory *UniceFactoryFilterer) FilterMarketCreated(opts *bind.FilterOpts, market []common.Address) (*UniceFactoryMarketCreatedIterator, error) {

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _UniceFactory.contract.FilterLogs(opts, "MarketCreated", marketRule)
	if err != nil {
		return nil, err
	}
	return &UniceFactoryMarketCreatedIterator{contract: _UniceFactory.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0xed229bcc1acd62d965455d23da5563e7082eb7006e172570c6dc39d189705b79.
//
// Solidity: event MarketCreated(address indexed market, string question, string category, uint256 bettingDeadline, uint256 resolutionDeadline)
func (_UniceFactory *UniceFactoryFilterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *UniceFactoryMarketCreated, market []common.Address) (event.Subscription, error) {

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _UniceFactory.contract.WatchLogs(opts, "MarketCreated", marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniceFactoryMarketCreated)
				if err := _UniceFactory.contract.UnpackLog(event, "MarketCreated", log); err != nil {
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

// ParseMarketCreated is a log parse operation binding the contract event 0xed229bcc1acd62d965455d23da5563e7082eb7006e172570c6dc39d189705b79.
//
// Solidity: event MarketCreated(address indexed market, string question, string category, uint256 bettingDeadline, uint256 resolutionDeadline)
func (_UniceFactory *UniceFactoryFilterer) ParseMarketCreated(log types.Log) (*UniceFactoryMarketCreated, error) {
	event := new(UniceFactoryMarketCreated)
	if err := _UniceFactory.contract.UnpackLog(event, "MarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
