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

// AlchemyUserOperation is an auto generated low-level Go binding around an user-defined struct.
type AlchemyUserOperation struct {
	Sender             common.Address
	Nonce              *big.Int
	InitCode           []byte
	CallData           []byte
	AccountGasLimits   [32]byte
	PreVerificationGas *big.Int
	GasFees            [32]byte
	PaymasterAndData   []byte
	Signature          []byte
}

// AlchemyBindingMetaData contains all meta data concerning the AlchemyBinding contract.
var AlchemyBindingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structAlchemyUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AlchemyBindingABI is the input ABI used to generate the binding from.
// Deprecated: Use AlchemyBindingMetaData.ABI instead.
var AlchemyBindingABI = AlchemyBindingMetaData.ABI

// AlchemyBinding is an auto generated Go binding around an Ethereum contract.
type AlchemyBinding struct {
	AlchemyBindingCaller     // Read-only binding to the contract
	AlchemyBindingTransactor // Write-only binding to the contract
	AlchemyBindingFilterer   // Log filterer for contract events
}

// AlchemyBindingCaller is an auto generated read-only Go binding around an Ethereum contract.
type AlchemyBindingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlchemyBindingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AlchemyBindingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlchemyBindingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AlchemyBindingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlchemyBindingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AlchemyBindingSession struct {
	Contract     *AlchemyBinding   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AlchemyBindingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AlchemyBindingCallerSession struct {
	Contract *AlchemyBindingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AlchemyBindingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AlchemyBindingTransactorSession struct {
	Contract     *AlchemyBindingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AlchemyBindingRaw is an auto generated low-level Go binding around an Ethereum contract.
type AlchemyBindingRaw struct {
	Contract *AlchemyBinding // Generic contract binding to access the raw methods on
}

// AlchemyBindingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AlchemyBindingCallerRaw struct {
	Contract *AlchemyBindingCaller // Generic read-only contract binding to access the raw methods on
}

// AlchemyBindingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AlchemyBindingTransactorRaw struct {
	Contract *AlchemyBindingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAlchemyBinding creates a new instance of AlchemyBinding, bound to a specific deployed contract.
func NewAlchemyBinding(address common.Address, backend bind.ContractBackend) (*AlchemyBinding, error) {
	contract, err := bindAlchemyBinding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AlchemyBinding{AlchemyBindingCaller: AlchemyBindingCaller{contract: contract}, AlchemyBindingTransactor: AlchemyBindingTransactor{contract: contract}, AlchemyBindingFilterer: AlchemyBindingFilterer{contract: contract}}, nil
}

// NewAlchemyBindingCaller creates a new read-only instance of AlchemyBinding, bound to a specific deployed contract.
func NewAlchemyBindingCaller(address common.Address, caller bind.ContractCaller) (*AlchemyBindingCaller, error) {
	contract, err := bindAlchemyBinding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AlchemyBindingCaller{contract: contract}, nil
}

// NewAlchemyBindingTransactor creates a new write-only instance of AlchemyBinding, bound to a specific deployed contract.
func NewAlchemyBindingTransactor(address common.Address, transactor bind.ContractTransactor) (*AlchemyBindingTransactor, error) {
	contract, err := bindAlchemyBinding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AlchemyBindingTransactor{contract: contract}, nil
}

// NewAlchemyBindingFilterer creates a new log filterer instance of AlchemyBinding, bound to a specific deployed contract.
func NewAlchemyBindingFilterer(address common.Address, filterer bind.ContractFilterer) (*AlchemyBindingFilterer, error) {
	contract, err := bindAlchemyBinding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AlchemyBindingFilterer{contract: contract}, nil
}

// bindAlchemyBinding binds a generic wrapper to an already deployed contract.
func bindAlchemyBinding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AlchemyBindingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AlchemyBinding *AlchemyBindingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AlchemyBinding.Contract.AlchemyBindingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AlchemyBinding *AlchemyBindingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlchemyBinding.Contract.AlchemyBindingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AlchemyBinding *AlchemyBindingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AlchemyBinding.Contract.AlchemyBindingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AlchemyBinding *AlchemyBindingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AlchemyBinding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AlchemyBinding *AlchemyBindingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlchemyBinding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AlchemyBinding *AlchemyBindingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AlchemyBinding.Contract.contract.Transact(opts, method, params...)
}

// HandleOps is a paid mutator transaction binding the contract method 0x765e827f.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] ops, address beneficiary) returns()
func (_AlchemyBinding *AlchemyBindingTransactor) HandleOps(opts *bind.TransactOpts, ops []AlchemyUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _AlchemyBinding.contract.Transact(opts, "handleOps", ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x765e827f.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] ops, address beneficiary) returns()
func (_AlchemyBinding *AlchemyBindingSession) HandleOps(ops []AlchemyUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _AlchemyBinding.Contract.HandleOps(&_AlchemyBinding.TransactOpts, ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x765e827f.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes)[] ops, address beneficiary) returns()
func (_AlchemyBinding *AlchemyBindingTransactorSession) HandleOps(ops []AlchemyUserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _AlchemyBinding.Contract.HandleOps(&_AlchemyBinding.TransactOpts, ops, beneficiary)
}
