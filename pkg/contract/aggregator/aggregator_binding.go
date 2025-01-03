// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aggregator

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

// AggregatorBindingMetaData contains all meta data concerning the AggregatorBinding contract.
var AggregatorBindingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"name\":\"target\",\"type\":\"address[]\"},{\"name\":\"value\",\"type\":\"uint256[]\"},{\"name\":\"callData\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AggregatorBindingABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregatorBindingMetaData.ABI instead.
var AggregatorBindingABI = AggregatorBindingMetaData.ABI

// AggregatorBinding is an auto generated Go binding around an Ethereum contract.
type AggregatorBinding struct {
	AggregatorBindingCaller     // Read-only binding to the contract
	AggregatorBindingTransactor // Write-only binding to the contract
	AggregatorBindingFilterer   // Log filterer for contract events
}

// AggregatorBindingCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggregatorBindingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorBindingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregatorBindingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorBindingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregatorBindingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorBindingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregatorBindingSession struct {
	Contract     *AggregatorBinding // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AggregatorBindingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregatorBindingCallerSession struct {
	Contract *AggregatorBindingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AggregatorBindingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregatorBindingTransactorSession struct {
	Contract     *AggregatorBindingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AggregatorBindingRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggregatorBindingRaw struct {
	Contract *AggregatorBinding // Generic contract binding to access the raw methods on
}

// AggregatorBindingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregatorBindingCallerRaw struct {
	Contract *AggregatorBindingCaller // Generic read-only contract binding to access the raw methods on
}

// AggregatorBindingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregatorBindingTransactorRaw struct {
	Contract *AggregatorBindingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregatorBinding creates a new instance of AggregatorBinding, bound to a specific deployed contract.
func NewAggregatorBinding(address common.Address, backend bind.ContractBackend) (*AggregatorBinding, error) {
	contract, err := bindAggregatorBinding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggregatorBinding{AggregatorBindingCaller: AggregatorBindingCaller{contract: contract}, AggregatorBindingTransactor: AggregatorBindingTransactor{contract: contract}, AggregatorBindingFilterer: AggregatorBindingFilterer{contract: contract}}, nil
}

// NewAggregatorBindingCaller creates a new read-only instance of AggregatorBinding, bound to a specific deployed contract.
func NewAggregatorBindingCaller(address common.Address, caller bind.ContractCaller) (*AggregatorBindingCaller, error) {
	contract, err := bindAggregatorBinding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorBindingCaller{contract: contract}, nil
}

// NewAggregatorBindingTransactor creates a new write-only instance of AggregatorBinding, bound to a specific deployed contract.
func NewAggregatorBindingTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregatorBindingTransactor, error) {
	contract, err := bindAggregatorBinding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorBindingTransactor{contract: contract}, nil
}

// NewAggregatorBindingFilterer creates a new log filterer instance of AggregatorBinding, bound to a specific deployed contract.
func NewAggregatorBindingFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregatorBindingFilterer, error) {
	contract, err := bindAggregatorBinding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregatorBindingFilterer{contract: contract}, nil
}

// bindAggregatorBinding binds a generic wrapper to an already deployed contract.
func bindAggregatorBinding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggregatorBindingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregatorBinding *AggregatorBindingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorBinding.Contract.AggregatorBindingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregatorBinding *AggregatorBindingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorBinding.Contract.AggregatorBindingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregatorBinding *AggregatorBindingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorBinding.Contract.AggregatorBindingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregatorBinding *AggregatorBindingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorBinding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregatorBinding *AggregatorBindingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorBinding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregatorBinding *AggregatorBindingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorBinding.Contract.contract.Transact(opts, method, params...)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] target, uint256[] value, bytes[] callData) returns()
func (_AggregatorBinding *AggregatorBindingTransactor) ExecuteBatch(opts *bind.TransactOpts, target []common.Address, value []*big.Int, callData [][]byte) (*types.Transaction, error) {
	return _AggregatorBinding.contract.Transact(opts, "executeBatch", target, value, callData)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] target, uint256[] value, bytes[] callData) returns()
func (_AggregatorBinding *AggregatorBindingSession) ExecuteBatch(target []common.Address, value []*big.Int, callData [][]byte) (*types.Transaction, error) {
	return _AggregatorBinding.Contract.ExecuteBatch(&_AggregatorBinding.TransactOpts, target, value, callData)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] target, uint256[] value, bytes[] callData) returns()
func (_AggregatorBinding *AggregatorBindingTransactorSession) ExecuteBatch(target []common.Address, value []*big.Int, callData [][]byte) (*types.Transaction, error) {
	return _AggregatorBinding.Contract.ExecuteBatch(&_AggregatorBinding.TransactOpts, target, value, callData)
}
