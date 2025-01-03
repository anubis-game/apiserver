// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package entrypoint

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

// BiconomyUserOperation is an auto generated low-level Go binding around an user-defined struct.
type BiconomyUserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

// BiconomyBindingMetaData contains all meta data concerning the BiconomyBinding contract.
var BiconomyBindingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBiconomyUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BiconomyBindingABI is the input ABI used to generate the binding from.
// Deprecated: Use BiconomyBindingMetaData.ABI instead.
var BiconomyBindingABI = BiconomyBindingMetaData.ABI

// BiconomyBinding is an auto generated Go binding around an Ethereum contract.
type BiconomyBinding struct {
	BiconomyBindingCaller     // Read-only binding to the contract
	BiconomyBindingTransactor // Write-only binding to the contract
	BiconomyBindingFilterer   // Log filterer for contract events
}

// BiconomyBindingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BiconomyBindingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiconomyBindingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BiconomyBindingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiconomyBindingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BiconomyBindingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiconomyBindingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BiconomyBindingSession struct {
	Contract     *BiconomyBinding  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BiconomyBindingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BiconomyBindingCallerSession struct {
	Contract *BiconomyBindingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BiconomyBindingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BiconomyBindingTransactorSession struct {
	Contract     *BiconomyBindingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BiconomyBindingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BiconomyBindingRaw struct {
	Contract *BiconomyBinding // Generic contract binding to access the raw methods on
}

// BiconomyBindingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BiconomyBindingCallerRaw struct {
	Contract *BiconomyBindingCaller // Generic read-only contract binding to access the raw methods on
}

// BiconomyBindingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BiconomyBindingTransactorRaw struct {
	Contract *BiconomyBindingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBiconomyBinding creates a new instance of BiconomyBinding, bound to a specific deployed contract.
func NewBiconomyBinding(address common.Address, backend bind.ContractBackend) (*BiconomyBinding, error) {
	contract, err := bindBiconomyBinding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BiconomyBinding{BiconomyBindingCaller: BiconomyBindingCaller{contract: contract}, BiconomyBindingTransactor: BiconomyBindingTransactor{contract: contract}, BiconomyBindingFilterer: BiconomyBindingFilterer{contract: contract}}, nil
}

// NewBiconomyBindingCaller creates a new read-only instance of BiconomyBinding, bound to a specific deployed contract.
func NewBiconomyBindingCaller(address common.Address, caller bind.ContractCaller) (*BiconomyBindingCaller, error) {
	contract, err := bindBiconomyBinding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BiconomyBindingCaller{contract: contract}, nil
}

// NewBiconomyBindingTransactor creates a new write-only instance of BiconomyBinding, bound to a specific deployed contract.
func NewBiconomyBindingTransactor(address common.Address, transactor bind.ContractTransactor) (*BiconomyBindingTransactor, error) {
	contract, err := bindBiconomyBinding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BiconomyBindingTransactor{contract: contract}, nil
}

// NewBiconomyBindingFilterer creates a new log filterer instance of BiconomyBinding, bound to a specific deployed contract.
func NewBiconomyBindingFilterer(address common.Address, filterer bind.ContractFilterer) (*BiconomyBindingFilterer, error) {
	contract, err := bindBiconomyBinding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BiconomyBindingFilterer{contract: contract}, nil
}

// bindBiconomyBinding binds a generic wrapper to an already deployed contract.
func bindBiconomyBinding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BiconomyBindingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BiconomyBinding *BiconomyBindingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BiconomyBinding.Contract.BiconomyBindingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BiconomyBinding *BiconomyBindingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BiconomyBinding.Contract.BiconomyBindingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BiconomyBinding *BiconomyBindingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BiconomyBinding.Contract.BiconomyBindingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BiconomyBinding *BiconomyBindingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BiconomyBinding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BiconomyBinding *BiconomyBindingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BiconomyBinding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BiconomyBinding *BiconomyBindingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BiconomyBinding.Contract.contract.Transact(opts, method, params...)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_BiconomyBinding *BiconomyBindingTransactor) HandleOps(opts *bind.TransactOpts, ops []BiconomyUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _BiconomyBinding.contract.Transact(opts, "handleOps", ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_BiconomyBinding *BiconomyBindingSession) HandleOps(ops []BiconomyUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _BiconomyBinding.Contract.HandleOps(&_BiconomyBinding.TransactOpts, ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_BiconomyBinding *BiconomyBindingTransactorSession) HandleOps(ops []BiconomyUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _BiconomyBinding.Contract.HandleOps(&_BiconomyBinding.TransactOpts, ops, beneficiary)
}
