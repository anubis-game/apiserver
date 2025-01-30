// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// RegistryBindingMetaData contains all meta data concerning the RegistryBinding contract.
var RegistryBindingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ben\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tok\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buy\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Address\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"Balance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Process\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"grd\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"kil\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"win\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"los\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dep\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buy\",\"type\":\"uint256\"}],\"name\":\"GuardianResolve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pla\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"grd\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"kil\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"win\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"los\",\"type\":\"address\"}],\"name\":\"WitnessPublish\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_FEE\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_GUARDIAN\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_PROTOCOL\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_SPLIT\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_TOTAL\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"tim\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sig\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sgn\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"tim\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"wal\",\"type\":\"address\"}],\"name\":\"DepositMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grd\",\"type\":\"address\"}],\"name\":\"Escape\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KILL_STATE_ESCAPE\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KILL_STATE_RELEASE\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"kil\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"win\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"los\",\"type\":\"address\"}],\"name\":\"Publish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"mes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sgn\",\"type\":\"bytes\"}],\"name\":\"RecoverSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"los\",\"type\":\"address\"}],\"name\":\"Release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grd\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"tim\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"wal\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sgn\",\"type\":\"bytes\"}],\"name\":\"Request\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grd\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"tim\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"pla\",\"type\":\"address\"}],\"name\":\"RequestMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"kil\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"win\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"los\",\"type\":\"address\"}],\"name\":\"Resolve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wal\",\"type\":\"address\"}],\"name\":\"SearchBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wal\",\"type\":\"address\"}],\"name\":\"SearchSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ben\",\"type\":\"address\"}],\"name\":\"UpdateBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RegistryBindingABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryBindingMetaData.ABI instead.
var RegistryBindingABI = RegistryBindingMetaData.ABI

// RegistryBinding is an auto generated Go binding around an Ethereum contract.
type RegistryBinding struct {
	RegistryBindingCaller     // Read-only binding to the contract
	RegistryBindingTransactor // Write-only binding to the contract
	RegistryBindingFilterer   // Log filterer for contract events
}

// RegistryBindingCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryBindingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryBindingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryBindingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryBindingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryBindingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryBindingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistryBindingSession struct {
	Contract     *RegistryBinding  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryBindingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryBindingCallerSession struct {
	Contract *RegistryBindingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RegistryBindingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryBindingTransactorSession struct {
	Contract     *RegistryBindingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RegistryBindingRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryBindingRaw struct {
	Contract *RegistryBinding // Generic contract binding to access the raw methods on
}

// RegistryBindingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryBindingCallerRaw struct {
	Contract *RegistryBindingCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryBindingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryBindingTransactorRaw struct {
	Contract *RegistryBindingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistryBinding creates a new instance of RegistryBinding, bound to a specific deployed contract.
func NewRegistryBinding(address common.Address, backend bind.ContractBackend) (*RegistryBinding, error) {
	contract, err := bindRegistryBinding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegistryBinding{RegistryBindingCaller: RegistryBindingCaller{contract: contract}, RegistryBindingTransactor: RegistryBindingTransactor{contract: contract}, RegistryBindingFilterer: RegistryBindingFilterer{contract: contract}}, nil
}

// NewRegistryBindingCaller creates a new read-only instance of RegistryBinding, bound to a specific deployed contract.
func NewRegistryBindingCaller(address common.Address, caller bind.ContractCaller) (*RegistryBindingCaller, error) {
	contract, err := bindRegistryBinding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryBindingCaller{contract: contract}, nil
}

// NewRegistryBindingTransactor creates a new write-only instance of RegistryBinding, bound to a specific deployed contract.
func NewRegistryBindingTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryBindingTransactor, error) {
	contract, err := bindRegistryBinding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryBindingTransactor{contract: contract}, nil
}

// NewRegistryBindingFilterer creates a new log filterer instance of RegistryBinding, bound to a specific deployed contract.
func NewRegistryBindingFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryBindingFilterer, error) {
	contract, err := bindRegistryBinding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryBindingFilterer{contract: contract}, nil
}

// bindRegistryBinding binds a generic wrapper to an already deployed contract.
func bindRegistryBinding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryBindingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistryBinding *RegistryBindingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryBinding.Contract.RegistryBindingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistryBinding *RegistryBindingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryBinding.Contract.RegistryBindingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistryBinding *RegistryBindingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryBinding.Contract.RegistryBindingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistryBinding *RegistryBindingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryBinding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistryBinding *RegistryBindingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryBinding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistryBinding *RegistryBindingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryBinding.Contract.contract.Transact(opts, method, params...)
}

// BASISFEE is a free data retrieval call binding the contract method 0x44c6d89b.
//
// Solidity: function BASIS_FEE() view returns(uint16)
func (_RegistryBinding *RegistryBindingCaller) BASISFEE(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "BASIS_FEE")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BASISFEE is a free data retrieval call binding the contract method 0x44c6d89b.
//
// Solidity: function BASIS_FEE() view returns(uint16)
func (_RegistryBinding *RegistryBindingSession) BASISFEE() (uint16, error) {
	return _RegistryBinding.Contract.BASISFEE(&_RegistryBinding.CallOpts)
}

// BASISFEE is a free data retrieval call binding the contract method 0x44c6d89b.
//
// Solidity: function BASIS_FEE() view returns(uint16)
func (_RegistryBinding *RegistryBindingCallerSession) BASISFEE() (uint16, error) {
	return _RegistryBinding.Contract.BASISFEE(&_RegistryBinding.CallOpts)
}

// BASISGUARDIAN is a free data retrieval call binding the contract method 0x949a86f4.
//
// Solidity: function BASIS_GUARDIAN() view returns(uint16)
func (_RegistryBinding *RegistryBindingCaller) BASISGUARDIAN(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "BASIS_GUARDIAN")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BASISGUARDIAN is a free data retrieval call binding the contract method 0x949a86f4.
//
// Solidity: function BASIS_GUARDIAN() view returns(uint16)
func (_RegistryBinding *RegistryBindingSession) BASISGUARDIAN() (uint16, error) {
	return _RegistryBinding.Contract.BASISGUARDIAN(&_RegistryBinding.CallOpts)
}

// BASISGUARDIAN is a free data retrieval call binding the contract method 0x949a86f4.
//
// Solidity: function BASIS_GUARDIAN() view returns(uint16)
func (_RegistryBinding *RegistryBindingCallerSession) BASISGUARDIAN() (uint16, error) {
	return _RegistryBinding.Contract.BASISGUARDIAN(&_RegistryBinding.CallOpts)
}

// BASISPROTOCOL is a free data retrieval call binding the contract method 0xc926e70f.
//
// Solidity: function BASIS_PROTOCOL() view returns(uint16)
func (_RegistryBinding *RegistryBindingCaller) BASISPROTOCOL(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "BASIS_PROTOCOL")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BASISPROTOCOL is a free data retrieval call binding the contract method 0xc926e70f.
//
// Solidity: function BASIS_PROTOCOL() view returns(uint16)
func (_RegistryBinding *RegistryBindingSession) BASISPROTOCOL() (uint16, error) {
	return _RegistryBinding.Contract.BASISPROTOCOL(&_RegistryBinding.CallOpts)
}

// BASISPROTOCOL is a free data retrieval call binding the contract method 0xc926e70f.
//
// Solidity: function BASIS_PROTOCOL() view returns(uint16)
func (_RegistryBinding *RegistryBindingCallerSession) BASISPROTOCOL() (uint16, error) {
	return _RegistryBinding.Contract.BASISPROTOCOL(&_RegistryBinding.CallOpts)
}

// BASISSPLIT is a free data retrieval call binding the contract method 0x34cb7584.
//
// Solidity: function BASIS_SPLIT() view returns(uint16)
func (_RegistryBinding *RegistryBindingCaller) BASISSPLIT(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "BASIS_SPLIT")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BASISSPLIT is a free data retrieval call binding the contract method 0x34cb7584.
//
// Solidity: function BASIS_SPLIT() view returns(uint16)
func (_RegistryBinding *RegistryBindingSession) BASISSPLIT() (uint16, error) {
	return _RegistryBinding.Contract.BASISSPLIT(&_RegistryBinding.CallOpts)
}

// BASISSPLIT is a free data retrieval call binding the contract method 0x34cb7584.
//
// Solidity: function BASIS_SPLIT() view returns(uint16)
func (_RegistryBinding *RegistryBindingCallerSession) BASISSPLIT() (uint16, error) {
	return _RegistryBinding.Contract.BASISSPLIT(&_RegistryBinding.CallOpts)
}

// BASISTOTAL is a free data retrieval call binding the contract method 0x9a61d8c3.
//
// Solidity: function BASIS_TOTAL() view returns(uint16)
func (_RegistryBinding *RegistryBindingCaller) BASISTOTAL(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "BASIS_TOTAL")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BASISTOTAL is a free data retrieval call binding the contract method 0x9a61d8c3.
//
// Solidity: function BASIS_TOTAL() view returns(uint16)
func (_RegistryBinding *RegistryBindingSession) BASISTOTAL() (uint16, error) {
	return _RegistryBinding.Contract.BASISTOTAL(&_RegistryBinding.CallOpts)
}

// BASISTOTAL is a free data retrieval call binding the contract method 0x9a61d8c3.
//
// Solidity: function BASIS_TOTAL() view returns(uint16)
func (_RegistryBinding *RegistryBindingCallerSession) BASISTOTAL() (uint16, error) {
	return _RegistryBinding.Contract.BASISTOTAL(&_RegistryBinding.CallOpts)
}

// DepositMessage is a free data retrieval call binding the contract method 0x0dc9f0b0.
//
// Solidity: function DepositMessage(uint64 tim, address wal) pure returns(bytes)
func (_RegistryBinding *RegistryBindingCaller) DepositMessage(opts *bind.CallOpts, tim uint64, wal common.Address) ([]byte, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "DepositMessage", tim, wal)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DepositMessage is a free data retrieval call binding the contract method 0x0dc9f0b0.
//
// Solidity: function DepositMessage(uint64 tim, address wal) pure returns(bytes)
func (_RegistryBinding *RegistryBindingSession) DepositMessage(tim uint64, wal common.Address) ([]byte, error) {
	return _RegistryBinding.Contract.DepositMessage(&_RegistryBinding.CallOpts, tim, wal)
}

// DepositMessage is a free data retrieval call binding the contract method 0x0dc9f0b0.
//
// Solidity: function DepositMessage(uint64 tim, address wal) pure returns(bytes)
func (_RegistryBinding *RegistryBindingCallerSession) DepositMessage(tim uint64, wal common.Address) ([]byte, error) {
	return _RegistryBinding.Contract.DepositMessage(&_RegistryBinding.CallOpts, tim, wal)
}

// KILLSTATEESCAPE is a free data retrieval call binding the contract method 0xe58c988c.
//
// Solidity: function KILL_STATE_ESCAPE() view returns(bytes16)
func (_RegistryBinding *RegistryBindingCaller) KILLSTATEESCAPE(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "KILL_STATE_ESCAPE")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// KILLSTATEESCAPE is a free data retrieval call binding the contract method 0xe58c988c.
//
// Solidity: function KILL_STATE_ESCAPE() view returns(bytes16)
func (_RegistryBinding *RegistryBindingSession) KILLSTATEESCAPE() ([16]byte, error) {
	return _RegistryBinding.Contract.KILLSTATEESCAPE(&_RegistryBinding.CallOpts)
}

// KILLSTATEESCAPE is a free data retrieval call binding the contract method 0xe58c988c.
//
// Solidity: function KILL_STATE_ESCAPE() view returns(bytes16)
func (_RegistryBinding *RegistryBindingCallerSession) KILLSTATEESCAPE() ([16]byte, error) {
	return _RegistryBinding.Contract.KILLSTATEESCAPE(&_RegistryBinding.CallOpts)
}

// KILLSTATERELEASE is a free data retrieval call binding the contract method 0xf1e2ab6c.
//
// Solidity: function KILL_STATE_RELEASE() view returns(bytes16)
func (_RegistryBinding *RegistryBindingCaller) KILLSTATERELEASE(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "KILL_STATE_RELEASE")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// KILLSTATERELEASE is a free data retrieval call binding the contract method 0xf1e2ab6c.
//
// Solidity: function KILL_STATE_RELEASE() view returns(bytes16)
func (_RegistryBinding *RegistryBindingSession) KILLSTATERELEASE() ([16]byte, error) {
	return _RegistryBinding.Contract.KILLSTATERELEASE(&_RegistryBinding.CallOpts)
}

// KILLSTATERELEASE is a free data retrieval call binding the contract method 0xf1e2ab6c.
//
// Solidity: function KILL_STATE_RELEASE() view returns(bytes16)
func (_RegistryBinding *RegistryBindingCallerSession) KILLSTATERELEASE() ([16]byte, error) {
	return _RegistryBinding.Contract.KILLSTATERELEASE(&_RegistryBinding.CallOpts)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x073afc87.
//
// Solidity: function RecoverSigner(bytes mes, bytes sgn) pure returns(address)
func (_RegistryBinding *RegistryBindingCaller) RecoverSigner(opts *bind.CallOpts, mes []byte, sgn []byte) (common.Address, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "RecoverSigner", mes, sgn)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverSigner is a free data retrieval call binding the contract method 0x073afc87.
//
// Solidity: function RecoverSigner(bytes mes, bytes sgn) pure returns(address)
func (_RegistryBinding *RegistryBindingSession) RecoverSigner(mes []byte, sgn []byte) (common.Address, error) {
	return _RegistryBinding.Contract.RecoverSigner(&_RegistryBinding.CallOpts, mes, sgn)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x073afc87.
//
// Solidity: function RecoverSigner(bytes mes, bytes sgn) pure returns(address)
func (_RegistryBinding *RegistryBindingCallerSession) RecoverSigner(mes []byte, sgn []byte) (common.Address, error) {
	return _RegistryBinding.Contract.RecoverSigner(&_RegistryBinding.CallOpts, mes, sgn)
}

// RequestMessage is a free data retrieval call binding the contract method 0x8b601b4a.
//
// Solidity: function RequestMessage(address grd, uint64 tim, address pla) pure returns(bytes)
func (_RegistryBinding *RegistryBindingCaller) RequestMessage(opts *bind.CallOpts, grd common.Address, tim uint64, pla common.Address) ([]byte, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "RequestMessage", grd, tim, pla)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// RequestMessage is a free data retrieval call binding the contract method 0x8b601b4a.
//
// Solidity: function RequestMessage(address grd, uint64 tim, address pla) pure returns(bytes)
func (_RegistryBinding *RegistryBindingSession) RequestMessage(grd common.Address, tim uint64, pla common.Address) ([]byte, error) {
	return _RegistryBinding.Contract.RequestMessage(&_RegistryBinding.CallOpts, grd, tim, pla)
}

// RequestMessage is a free data retrieval call binding the contract method 0x8b601b4a.
//
// Solidity: function RequestMessage(address grd, uint64 tim, address pla) pure returns(bytes)
func (_RegistryBinding *RegistryBindingCallerSession) RequestMessage(grd common.Address, tim uint64, pla common.Address) ([]byte, error) {
	return _RegistryBinding.Contract.RequestMessage(&_RegistryBinding.CallOpts, grd, tim, pla)
}

// SearchBalance is a free data retrieval call binding the contract method 0x026b2666.
//
// Solidity: function SearchBalance(address wal) view returns(uint256, uint256, uint256)
func (_RegistryBinding *RegistryBindingCaller) SearchBalance(opts *bind.CallOpts, wal common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "SearchBalance", wal)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// SearchBalance is a free data retrieval call binding the contract method 0x026b2666.
//
// Solidity: function SearchBalance(address wal) view returns(uint256, uint256, uint256)
func (_RegistryBinding *RegistryBindingSession) SearchBalance(wal common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _RegistryBinding.Contract.SearchBalance(&_RegistryBinding.CallOpts, wal)
}

// SearchBalance is a free data retrieval call binding the contract method 0x026b2666.
//
// Solidity: function SearchBalance(address wal) view returns(uint256, uint256, uint256)
func (_RegistryBinding *RegistryBindingCallerSession) SearchBalance(wal common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _RegistryBinding.Contract.SearchBalance(&_RegistryBinding.CallOpts, wal)
}

// SearchSigner is a free data retrieval call binding the contract method 0xc78e95ec.
//
// Solidity: function SearchSigner(address wal) view returns(address, address)
func (_RegistryBinding *RegistryBindingCaller) SearchSigner(opts *bind.CallOpts, wal common.Address) (common.Address, common.Address, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "SearchSigner", wal)

	if err != nil {
		return *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// SearchSigner is a free data retrieval call binding the contract method 0xc78e95ec.
//
// Solidity: function SearchSigner(address wal) view returns(address, address)
func (_RegistryBinding *RegistryBindingSession) SearchSigner(wal common.Address) (common.Address, common.Address, error) {
	return _RegistryBinding.Contract.SearchSigner(&_RegistryBinding.CallOpts, wal)
}

// SearchSigner is a free data retrieval call binding the contract method 0xc78e95ec.
//
// Solidity: function SearchSigner(address wal) view returns(address, address)
func (_RegistryBinding *RegistryBindingCallerSession) SearchSigner(wal common.Address) (common.Address, common.Address, error) {
	return _RegistryBinding.Contract.SearchSigner(&_RegistryBinding.CallOpts, wal)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_RegistryBinding *RegistryBindingCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_RegistryBinding *RegistryBindingSession) VERSION() (string, error) {
	return _RegistryBinding.Contract.VERSION(&_RegistryBinding.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_RegistryBinding *RegistryBindingCallerSession) VERSION() (string, error) {
	return _RegistryBinding.Contract.VERSION(&_RegistryBinding.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_RegistryBinding *RegistryBindingCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_RegistryBinding *RegistryBindingSession) Beneficiary() (common.Address, error) {
	return _RegistryBinding.Contract.Beneficiary(&_RegistryBinding.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_RegistryBinding *RegistryBindingCallerSession) Beneficiary() (common.Address, error) {
	return _RegistryBinding.Contract.Beneficiary(&_RegistryBinding.CallOpts)
}

// Buyin is a free data retrieval call binding the contract method 0x69766079.
//
// Solidity: function buyin() view returns(uint256)
func (_RegistryBinding *RegistryBindingCaller) Buyin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "buyin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Buyin is a free data retrieval call binding the contract method 0x69766079.
//
// Solidity: function buyin() view returns(uint256)
func (_RegistryBinding *RegistryBindingSession) Buyin() (*big.Int, error) {
	return _RegistryBinding.Contract.Buyin(&_RegistryBinding.CallOpts)
}

// Buyin is a free data retrieval call binding the contract method 0x69766079.
//
// Solidity: function buyin() view returns(uint256)
func (_RegistryBinding *RegistryBindingCallerSession) Buyin() (*big.Int, error) {
	return _RegistryBinding.Contract.Buyin(&_RegistryBinding.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RegistryBinding *RegistryBindingCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RegistryBinding *RegistryBindingSession) Token() (common.Address, error) {
	return _RegistryBinding.Contract.Token(&_RegistryBinding.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RegistryBinding *RegistryBindingCallerSession) Token() (common.Address, error) {
	return _RegistryBinding.Contract.Token(&_RegistryBinding.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x96b13c54.
//
// Solidity: function Deposit(uint256 bal, uint64 tim, address sig, bytes sgn) returns()
func (_RegistryBinding *RegistryBindingTransactor) Deposit(opts *bind.TransactOpts, bal *big.Int, tim uint64, sig common.Address, sgn []byte) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "Deposit", bal, tim, sig, sgn)
}

// Deposit is a paid mutator transaction binding the contract method 0x96b13c54.
//
// Solidity: function Deposit(uint256 bal, uint64 tim, address sig, bytes sgn) returns()
func (_RegistryBinding *RegistryBindingSession) Deposit(bal *big.Int, tim uint64, sig common.Address, sgn []byte) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Deposit(&_RegistryBinding.TransactOpts, bal, tim, sig, sgn)
}

// Deposit is a paid mutator transaction binding the contract method 0x96b13c54.
//
// Solidity: function Deposit(uint256 bal, uint64 tim, address sig, bytes sgn) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) Deposit(bal *big.Int, tim uint64, sig common.Address, sgn []byte) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Deposit(&_RegistryBinding.TransactOpts, bal, tim, sig, sgn)
}

// Escape is a paid mutator transaction binding the contract method 0x0d5dd3e4.
//
// Solidity: function Escape(address grd) returns()
func (_RegistryBinding *RegistryBindingTransactor) Escape(opts *bind.TransactOpts, grd common.Address) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "Escape", grd)
}

// Escape is a paid mutator transaction binding the contract method 0x0d5dd3e4.
//
// Solidity: function Escape(address grd) returns()
func (_RegistryBinding *RegistryBindingSession) Escape(grd common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Escape(&_RegistryBinding.TransactOpts, grd)
}

// Escape is a paid mutator transaction binding the contract method 0x0d5dd3e4.
//
// Solidity: function Escape(address grd) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) Escape(grd common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Escape(&_RegistryBinding.TransactOpts, grd)
}

// Publish is a paid mutator transaction binding the contract method 0x0582ffb9.
//
// Solidity: function Publish(bytes16 kil, address win, address los) returns()
func (_RegistryBinding *RegistryBindingTransactor) Publish(opts *bind.TransactOpts, kil [16]byte, win common.Address, los common.Address) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "Publish", kil, win, los)
}

// Publish is a paid mutator transaction binding the contract method 0x0582ffb9.
//
// Solidity: function Publish(bytes16 kil, address win, address los) returns()
func (_RegistryBinding *RegistryBindingSession) Publish(kil [16]byte, win common.Address, los common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Publish(&_RegistryBinding.TransactOpts, kil, win, los)
}

// Publish is a paid mutator transaction binding the contract method 0x0582ffb9.
//
// Solidity: function Publish(bytes16 kil, address win, address los) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) Publish(kil [16]byte, win common.Address, los common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Publish(&_RegistryBinding.TransactOpts, kil, win, los)
}

// Release is a paid mutator transaction binding the contract method 0x79552101.
//
// Solidity: function Release(address los) returns()
func (_RegistryBinding *RegistryBindingTransactor) Release(opts *bind.TransactOpts, los common.Address) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "Release", los)
}

// Release is a paid mutator transaction binding the contract method 0x79552101.
//
// Solidity: function Release(address los) returns()
func (_RegistryBinding *RegistryBindingSession) Release(los common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Release(&_RegistryBinding.TransactOpts, los)
}

// Release is a paid mutator transaction binding the contract method 0x79552101.
//
// Solidity: function Release(address los) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) Release(los common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Release(&_RegistryBinding.TransactOpts, los)
}

// Request is a paid mutator transaction binding the contract method 0x2ef5c5a4.
//
// Solidity: function Request(address grd, uint64 tim, address wal, bytes sgn) returns()
func (_RegistryBinding *RegistryBindingTransactor) Request(opts *bind.TransactOpts, grd common.Address, tim uint64, wal common.Address, sgn []byte) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "Request", grd, tim, wal, sgn)
}

// Request is a paid mutator transaction binding the contract method 0x2ef5c5a4.
//
// Solidity: function Request(address grd, uint64 tim, address wal, bytes sgn) returns()
func (_RegistryBinding *RegistryBindingSession) Request(grd common.Address, tim uint64, wal common.Address, sgn []byte) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Request(&_RegistryBinding.TransactOpts, grd, tim, wal, sgn)
}

// Request is a paid mutator transaction binding the contract method 0x2ef5c5a4.
//
// Solidity: function Request(address grd, uint64 tim, address wal, bytes sgn) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) Request(grd common.Address, tim uint64, wal common.Address, sgn []byte) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Request(&_RegistryBinding.TransactOpts, grd, tim, wal, sgn)
}

// Resolve is a paid mutator transaction binding the contract method 0x38162d19.
//
// Solidity: function Resolve(bytes16 kil, address win, address los) returns()
func (_RegistryBinding *RegistryBindingTransactor) Resolve(opts *bind.TransactOpts, kil [16]byte, win common.Address, los common.Address) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "Resolve", kil, win, los)
}

// Resolve is a paid mutator transaction binding the contract method 0x38162d19.
//
// Solidity: function Resolve(bytes16 kil, address win, address los) returns()
func (_RegistryBinding *RegistryBindingSession) Resolve(kil [16]byte, win common.Address, los common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Resolve(&_RegistryBinding.TransactOpts, kil, win, los)
}

// Resolve is a paid mutator transaction binding the contract method 0x38162d19.
//
// Solidity: function Resolve(bytes16 kil, address win, address los) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) Resolve(kil [16]byte, win common.Address, los common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Resolve(&_RegistryBinding.TransactOpts, kil, win, los)
}

// UpdateBeneficiary is a paid mutator transaction binding the contract method 0xe6131a64.
//
// Solidity: function UpdateBeneficiary(address ben) returns()
func (_RegistryBinding *RegistryBindingTransactor) UpdateBeneficiary(opts *bind.TransactOpts, ben common.Address) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "UpdateBeneficiary", ben)
}

// UpdateBeneficiary is a paid mutator transaction binding the contract method 0xe6131a64.
//
// Solidity: function UpdateBeneficiary(address ben) returns()
func (_RegistryBinding *RegistryBindingSession) UpdateBeneficiary(ben common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.UpdateBeneficiary(&_RegistryBinding.TransactOpts, ben)
}

// UpdateBeneficiary is a paid mutator transaction binding the contract method 0xe6131a64.
//
// Solidity: function UpdateBeneficiary(address ben) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) UpdateBeneficiary(ben common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.UpdateBeneficiary(&_RegistryBinding.TransactOpts, ben)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5b6b431d.
//
// Solidity: function Withdraw(uint256 bal) returns()
func (_RegistryBinding *RegistryBindingTransactor) Withdraw(opts *bind.TransactOpts, bal *big.Int) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "Withdraw", bal)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5b6b431d.
//
// Solidity: function Withdraw(uint256 bal) returns()
func (_RegistryBinding *RegistryBindingSession) Withdraw(bal *big.Int) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Withdraw(&_RegistryBinding.TransactOpts, bal)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5b6b431d.
//
// Solidity: function Withdraw(uint256 bal) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) Withdraw(bal *big.Int) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Withdraw(&_RegistryBinding.TransactOpts, bal)
}

// RegistryBindingGuardianResolveIterator is returned from FilterGuardianResolve and is used to iterate over the raw logs and unpacked data for GuardianResolve events raised by the RegistryBinding contract.
type RegistryBindingGuardianResolveIterator struct {
	Event *RegistryBindingGuardianResolve // Event containing the contract specifics and raw log

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
func (it *RegistryBindingGuardianResolveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryBindingGuardianResolve)
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
		it.Event = new(RegistryBindingGuardianResolve)
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
func (it *RegistryBindingGuardianResolveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryBindingGuardianResolveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryBindingGuardianResolve represents a GuardianResolve event raised by the RegistryBinding contract.
type RegistryBindingGuardianResolve struct {
	Grd common.Address
	Kil [16]byte
	Win common.Address
	Los common.Address
	Dep *big.Int
	Buy *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterGuardianResolve is a free log retrieval operation binding the contract event 0xa9b260a64c941e0970a4c190f5dd5772398eea222162b693f21b759b434c57a1.
//
// Solidity: event GuardianResolve(address indexed grd, bytes16 kil, address win, address los, uint256 dep, uint256 buy)
func (_RegistryBinding *RegistryBindingFilterer) FilterGuardianResolve(opts *bind.FilterOpts, grd []common.Address) (*RegistryBindingGuardianResolveIterator, error) {

	var grdRule []interface{}
	for _, grdItem := range grd {
		grdRule = append(grdRule, grdItem)
	}

	logs, sub, err := _RegistryBinding.contract.FilterLogs(opts, "GuardianResolve", grdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryBindingGuardianResolveIterator{contract: _RegistryBinding.contract, event: "GuardianResolve", logs: logs, sub: sub}, nil
}

// WatchGuardianResolve is a free log subscription operation binding the contract event 0xa9b260a64c941e0970a4c190f5dd5772398eea222162b693f21b759b434c57a1.
//
// Solidity: event GuardianResolve(address indexed grd, bytes16 kil, address win, address los, uint256 dep, uint256 buy)
func (_RegistryBinding *RegistryBindingFilterer) WatchGuardianResolve(opts *bind.WatchOpts, sink chan<- *RegistryBindingGuardianResolve, grd []common.Address) (event.Subscription, error) {

	var grdRule []interface{}
	for _, grdItem := range grd {
		grdRule = append(grdRule, grdItem)
	}

	logs, sub, err := _RegistryBinding.contract.WatchLogs(opts, "GuardianResolve", grdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryBindingGuardianResolve)
				if err := _RegistryBinding.contract.UnpackLog(event, "GuardianResolve", log); err != nil {
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

// ParseGuardianResolve is a log parse operation binding the contract event 0xa9b260a64c941e0970a4c190f5dd5772398eea222162b693f21b759b434c57a1.
//
// Solidity: event GuardianResolve(address indexed grd, bytes16 kil, address win, address los, uint256 dep, uint256 buy)
func (_RegistryBinding *RegistryBindingFilterer) ParseGuardianResolve(log types.Log) (*RegistryBindingGuardianResolve, error) {
	event := new(RegistryBindingGuardianResolve)
	if err := _RegistryBinding.contract.UnpackLog(event, "GuardianResolve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryBindingWitnessPublishIterator is returned from FilterWitnessPublish and is used to iterate over the raw logs and unpacked data for WitnessPublish events raised by the RegistryBinding contract.
type RegistryBindingWitnessPublishIterator struct {
	Event *RegistryBindingWitnessPublish // Event containing the contract specifics and raw log

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
func (it *RegistryBindingWitnessPublishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryBindingWitnessPublish)
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
		it.Event = new(RegistryBindingWitnessPublish)
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
func (it *RegistryBindingWitnessPublishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryBindingWitnessPublishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryBindingWitnessPublish represents a WitnessPublish event raised by the RegistryBinding contract.
type RegistryBindingWitnessPublish struct {
	Pla common.Address
	Grd common.Address
	Kil [16]byte
	Win common.Address
	Los common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWitnessPublish is a free log retrieval operation binding the contract event 0x5ef8e288165cc58ddbb342d50149b26e7a8f6edb49612012249638754d8a235d.
//
// Solidity: event WitnessPublish(address indexed pla, address indexed grd, bytes16 kil, address win, address los)
func (_RegistryBinding *RegistryBindingFilterer) FilterWitnessPublish(opts *bind.FilterOpts, pla []common.Address, grd []common.Address) (*RegistryBindingWitnessPublishIterator, error) {

	var plaRule []interface{}
	for _, plaItem := range pla {
		plaRule = append(plaRule, plaItem)
	}
	var grdRule []interface{}
	for _, grdItem := range grd {
		grdRule = append(grdRule, grdItem)
	}

	logs, sub, err := _RegistryBinding.contract.FilterLogs(opts, "WitnessPublish", plaRule, grdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryBindingWitnessPublishIterator{contract: _RegistryBinding.contract, event: "WitnessPublish", logs: logs, sub: sub}, nil
}

// WatchWitnessPublish is a free log subscription operation binding the contract event 0x5ef8e288165cc58ddbb342d50149b26e7a8f6edb49612012249638754d8a235d.
//
// Solidity: event WitnessPublish(address indexed pla, address indexed grd, bytes16 kil, address win, address los)
func (_RegistryBinding *RegistryBindingFilterer) WatchWitnessPublish(opts *bind.WatchOpts, sink chan<- *RegistryBindingWitnessPublish, pla []common.Address, grd []common.Address) (event.Subscription, error) {

	var plaRule []interface{}
	for _, plaItem := range pla {
		plaRule = append(plaRule, plaItem)
	}
	var grdRule []interface{}
	for _, grdItem := range grd {
		grdRule = append(grdRule, grdItem)
	}

	logs, sub, err := _RegistryBinding.contract.WatchLogs(opts, "WitnessPublish", plaRule, grdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryBindingWitnessPublish)
				if err := _RegistryBinding.contract.UnpackLog(event, "WitnessPublish", log); err != nil {
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

// ParseWitnessPublish is a log parse operation binding the contract event 0x5ef8e288165cc58ddbb342d50149b26e7a8f6edb49612012249638754d8a235d.
//
// Solidity: event WitnessPublish(address indexed pla, address indexed grd, bytes16 kil, address win, address los)
func (_RegistryBinding *RegistryBindingFilterer) ParseWitnessPublish(log types.Log) (*RegistryBindingWitnessPublish, error) {
	event := new(RegistryBindingWitnessPublish)
	if err := _RegistryBinding.contract.UnpackLog(event, "WitnessPublish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
