// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
)

// ABiMetaData contains all meta data concerning the ABi contract.
var ABiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Account\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addrPerson\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"account\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"times\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"a\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"account\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferToAnother\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ABiABI is the input ABI used to generate the binding from.
// Deprecated: Use ABiMetaData.ABI instead.
var ABiABI = ABiMetaData.ABI

// ABi is an auto generated Go binding around an Ethereum contract.
type ABi struct {
	ABiCaller     // Read-only binding to the contract
	ABiTransactor // Write-only binding to the contract
	ABiFilterer   // Log filterer for contract events
}

// ABiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ABiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ABiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ABiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ABiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ABiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ABiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ABiSession struct {
	Contract     *ABi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ABiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ABiCallerSession struct {
	Contract *ABiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ABiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ABiTransactorSession struct {
	Contract     *ABiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ABiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ABiRaw struct {
	Contract *ABi // Generic contract binding to access the raw methods on
}

// ABiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ABiCallerRaw struct {
	Contract *ABiCaller // Generic read-only contract binding to access the raw methods on
}

// ABiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ABiTransactorRaw struct {
	Contract *ABiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewABi creates a new instance of ABi, bound to a specific deployed contract.
func NewABi(address common.Address, backend bind.ContractBackend) (*ABi, error) {
	contract, err := bindABi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ABi{ABiCaller: ABiCaller{contract: contract}, ABiTransactor: ABiTransactor{contract: contract}, ABiFilterer: ABiFilterer{contract: contract}}, nil
}

// NewABiCaller creates a new read-only instance of ABi, bound to a specific deployed contract.
func NewABiCaller(address common.Address, caller bind.ContractCaller) (*ABiCaller, error) {
	contract, err := bindABi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ABiCaller{contract: contract}, nil
}

// NewABiTransactor creates a new write-only instance of ABi, bound to a specific deployed contract.
func NewABiTransactor(address common.Address, transactor bind.ContractTransactor) (*ABiTransactor, error) {
	contract, err := bindABi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ABiTransactor{contract: contract}, nil
}

// NewABiFilterer creates a new log filterer instance of ABi, bound to a specific deployed contract.
func NewABiFilterer(address common.Address, filterer bind.ContractFilterer) (*ABiFilterer, error) {
	contract, err := bindABi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ABiFilterer{contract: contract}, nil
}

// bindABi binds a generic wrapper to an already deployed contract.
func bindABi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ABiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ABi *ABiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ABi.Contract.ABiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ABi *ABiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ABi.Contract.ABiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ABi *ABiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ABi.Contract.ABiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ABi *ABiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ABi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ABi *ABiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ABi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ABi *ABiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ABi.Contract.contract.Transact(opts, method, params...)
}

// B is a free data retrieval call binding the contract method 0x30311898.
//
// Solidity: function Account(address ) view returns(address addrPerson, uint256 account, uint256 times)
func (_ABi *ABiCaller) B(opts *bind.CallOpts, arg0 common.Address) (struct {
	AddrPerson common.Address
	Account    *big.Int
	Times      *big.Int
}, error) {
	var out []interface{}
	err := _ABi.contract.Call(opts, &out, "Account", arg0)

	outstruct := new(struct {
		AddrPerson common.Address
		Account    *big.Int
		Times      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AddrPerson = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Account = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Times = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// B is a free data retrieval call binding the contract method 0x30311898.
//
// Solidity: function Account(address ) view returns(address addrPerson, uint256 account, uint256 times)
func (_ABi *ABiSession) B(arg0 common.Address) (struct {
	AddrPerson common.Address
	Account    *big.Int
	Times      *big.Int
}, error) {
	return _ABi.Contract.B(&_ABi.CallOpts, arg0)
}

// B is a free data retrieval call binding the contract method 0x30311898.
//
// Solidity: function Account(address ) view returns(address addrPerson, uint256 account, uint256 times)
func (_ABi *ABiCallerSession) B(arg0 common.Address) (struct {
	AddrPerson common.Address
	Account    *big.Int
	Times      *big.Int
}, error) {
	return _ABi.Contract.B(&_ABi.CallOpts, arg0)
}

// A is a free data retrieval call binding the contract method 0xc68d81e0.
//
// Solidity: function a(address _to) view returns(uint256)
func (_ABi *ABiCaller) A(opts *bind.CallOpts, _to common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ABi.contract.Call(opts, &out, "a", _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xc68d81e0.
//
// Solidity: function a(address _to) view returns(uint256)
func (_ABi *ABiSession) A(_to common.Address) (*big.Int, error) {
	return _ABi.Contract.A(&_ABi.CallOpts, _to)
}

// A is a free data retrieval call binding the contract method 0xc68d81e0.
//
// Solidity: function a(address _to) view returns(uint256)
func (_ABi *ABiCallerSession) A(_to common.Address) (*big.Int, error) {
	return _ABi.Contract.A(&_ABi.CallOpts, _to)
}

// Account is a free data retrieval call binding the contract method 0x5dab2420.
//
// Solidity: function account() view returns(uint256)
func (_ABi *ABiCaller) Account(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ABi.contract.Call(opts, &out, "account")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Account is a free data retrieval call binding the contract method 0x5dab2420.
//
// Solidity: function account() view returns(uint256)
func (_ABi *ABiSession) Account() (*big.Int, error) {
	return _ABi.Contract.Account(&_ABi.CallOpts)
}

// Account is a free data retrieval call binding the contract method 0x5dab2420.
//
// Solidity: function account() view returns(uint256)
func (_ABi *ABiCallerSession) Account() (*big.Int, error) {
	return _ABi.Contract.Account(&_ABi.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_ABi *ABiCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ABi.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_ABi *ABiSession) GetBalance() (*big.Int, error) {
	return _ABi.Contract.GetBalance(&_ABi.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_ABi *ABiCallerSession) GetBalance() (*big.Int, error) {
	return _ABi.Contract.GetBalance(&_ABi.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ABi *ABiTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ABi.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ABi *ABiSession) Deposit() (*types.Transaction, error) {
	return _ABi.Contract.Deposit(&_ABi.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ABi *ABiTransactorSession) Deposit() (*types.Transaction, error) {
	return _ABi.Contract.Deposit(&_ABi.TransactOpts)
}

// TransferToAnother is a paid mutator transaction binding the contract method 0xcd6f4ec2.
//
// Solidity: function transferToAnother(address _to, uint256 _amount) returns()
func (_ABi *ABiTransactor) TransferToAnother(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ABi.contract.Transact(opts, "transferToAnother", _to, _amount)
}

// TransferToAnother is a paid mutator transaction binding the contract method 0xcd6f4ec2.
//
// Solidity: function transferToAnother(address _to, uint256 _amount) returns()
func (_ABi *ABiSession) TransferToAnother(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ABi.Contract.TransferToAnother(&_ABi.TransactOpts, _to, _amount)
}

// TransferToAnother is a paid mutator transaction binding the contract method 0xcd6f4ec2.
//
// Solidity: function transferToAnother(address _to, uint256 _amount) returns()
func (_ABi *ABiTransactorSession) TransferToAnother(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ABi.Contract.TransferToAnother(&_ABi.TransactOpts, _to, _amount)
}
