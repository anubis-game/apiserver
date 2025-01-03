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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"own\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tok\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buy\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Address\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"Balance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Process\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pla\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"grd\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kil\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"win\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"los\",\"type\":\"address\"}],\"name\":\"Report\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"grd\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kil\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"win\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"los\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dep\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bin\",\"type\":\"uint256\"}],\"name\":\"Resolve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_FEE\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_GUARDIAN\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_PROTOCOL\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_SPLIT\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_TOTAL\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"grd\",\"type\":\"address\"}],\"name\":\"balHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"tim\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sig\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sgn\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"tim\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"wal\",\"type\":\"address\"}],\"name\":\"depositMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wog\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"kil\",\"type\":\"uint256\"}],\"name\":\"keyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"mes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sgn\",\"type\":\"bytes\"}],\"name\":\"recoverSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grd\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"kil\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"win\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"los\",\"type\":\"address\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grd\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"tim\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"wal\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sgn\",\"type\":\"bytes\"}],\"name\":\"request\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grd\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"tim\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"pla\",\"type\":\"address\"}],\"name\":\"requestMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"kil\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"win\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"los\",\"type\":\"address\"}],\"name\":\"resolve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wal\",\"type\":\"address\"}],\"name\":\"searchBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wal\",\"type\":\"address\"}],\"name\":\"searchSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"win\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"los\",\"type\":\"address\"}],\"name\":\"valHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RegistryBinding *RegistryBindingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RegistryBinding *RegistryBindingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RegistryBinding.Contract.DEFAULTADMINROLE(&_RegistryBinding.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RegistryBinding *RegistryBindingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RegistryBinding.Contract.DEFAULTADMINROLE(&_RegistryBinding.CallOpts)
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

// BalHash is a free data retrieval call binding the contract method 0xa2f9262e.
//
// Solidity: function balHash(address wal, address grd) pure returns(bytes32)
func (_RegistryBinding *RegistryBindingCaller) BalHash(opts *bind.CallOpts, wal common.Address, grd common.Address) ([32]byte, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "balHash", wal, grd)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BalHash is a free data retrieval call binding the contract method 0xa2f9262e.
//
// Solidity: function balHash(address wal, address grd) pure returns(bytes32)
func (_RegistryBinding *RegistryBindingSession) BalHash(wal common.Address, grd common.Address) ([32]byte, error) {
	return _RegistryBinding.Contract.BalHash(&_RegistryBinding.CallOpts, wal, grd)
}

// BalHash is a free data retrieval call binding the contract method 0xa2f9262e.
//
// Solidity: function balHash(address wal, address grd) pure returns(bytes32)
func (_RegistryBinding *RegistryBindingCallerSession) BalHash(wal common.Address, grd common.Address) ([32]byte, error) {
	return _RegistryBinding.Contract.BalHash(&_RegistryBinding.CallOpts, wal, grd)
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

// DepositMessage is a free data retrieval call binding the contract method 0x25a998c5.
//
// Solidity: function depositMessage(uint64 tim, address wal) pure returns(bytes)
func (_RegistryBinding *RegistryBindingCaller) DepositMessage(opts *bind.CallOpts, tim uint64, wal common.Address) ([]byte, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "depositMessage", tim, wal)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DepositMessage is a free data retrieval call binding the contract method 0x25a998c5.
//
// Solidity: function depositMessage(uint64 tim, address wal) pure returns(bytes)
func (_RegistryBinding *RegistryBindingSession) DepositMessage(tim uint64, wal common.Address) ([]byte, error) {
	return _RegistryBinding.Contract.DepositMessage(&_RegistryBinding.CallOpts, tim, wal)
}

// DepositMessage is a free data retrieval call binding the contract method 0x25a998c5.
//
// Solidity: function depositMessage(uint64 tim, address wal) pure returns(bytes)
func (_RegistryBinding *RegistryBindingCallerSession) DepositMessage(tim uint64, wal common.Address) ([]byte, error) {
	return _RegistryBinding.Contract.DepositMessage(&_RegistryBinding.CallOpts, tim, wal)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RegistryBinding *RegistryBindingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RegistryBinding *RegistryBindingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RegistryBinding.Contract.GetRoleAdmin(&_RegistryBinding.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RegistryBinding *RegistryBindingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RegistryBinding.Contract.GetRoleAdmin(&_RegistryBinding.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RegistryBinding *RegistryBindingCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RegistryBinding *RegistryBindingSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _RegistryBinding.Contract.GetRoleMember(&_RegistryBinding.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RegistryBinding *RegistryBindingCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _RegistryBinding.Contract.GetRoleMember(&_RegistryBinding.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RegistryBinding *RegistryBindingCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RegistryBinding *RegistryBindingSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _RegistryBinding.Contract.GetRoleMemberCount(&_RegistryBinding.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RegistryBinding *RegistryBindingCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _RegistryBinding.Contract.GetRoleMemberCount(&_RegistryBinding.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_RegistryBinding *RegistryBindingCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_RegistryBinding *RegistryBindingSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _RegistryBinding.Contract.GetRoleMembers(&_RegistryBinding.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_RegistryBinding *RegistryBindingCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _RegistryBinding.Contract.GetRoleMembers(&_RegistryBinding.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RegistryBinding *RegistryBindingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RegistryBinding *RegistryBindingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RegistryBinding.Contract.HasRole(&_RegistryBinding.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RegistryBinding *RegistryBindingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RegistryBinding.Contract.HasRole(&_RegistryBinding.CallOpts, role, account)
}

// KeyHash is a free data retrieval call binding the contract method 0x469921ed.
//
// Solidity: function keyHash(address wog, uint256 kil) pure returns(bytes32)
func (_RegistryBinding *RegistryBindingCaller) KeyHash(opts *bind.CallOpts, wog common.Address, kil *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "keyHash", wog, kil)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KeyHash is a free data retrieval call binding the contract method 0x469921ed.
//
// Solidity: function keyHash(address wog, uint256 kil) pure returns(bytes32)
func (_RegistryBinding *RegistryBindingSession) KeyHash(wog common.Address, kil *big.Int) ([32]byte, error) {
	return _RegistryBinding.Contract.KeyHash(&_RegistryBinding.CallOpts, wog, kil)
}

// KeyHash is a free data retrieval call binding the contract method 0x469921ed.
//
// Solidity: function keyHash(address wog, uint256 kil) pure returns(bytes32)
func (_RegistryBinding *RegistryBindingCallerSession) KeyHash(wog common.Address, kil *big.Int) ([32]byte, error) {
	return _RegistryBinding.Contract.KeyHash(&_RegistryBinding.CallOpts, wog, kil)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistryBinding *RegistryBindingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistryBinding *RegistryBindingSession) Owner() (common.Address, error) {
	return _RegistryBinding.Contract.Owner(&_RegistryBinding.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistryBinding *RegistryBindingCallerSession) Owner() (common.Address, error) {
	return _RegistryBinding.Contract.Owner(&_RegistryBinding.CallOpts)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x2e295ec9.
//
// Solidity: function recoverSigner(bytes mes, bytes sgn) pure returns(address)
func (_RegistryBinding *RegistryBindingCaller) RecoverSigner(opts *bind.CallOpts, mes []byte, sgn []byte) (common.Address, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "recoverSigner", mes, sgn)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverSigner is a free data retrieval call binding the contract method 0x2e295ec9.
//
// Solidity: function recoverSigner(bytes mes, bytes sgn) pure returns(address)
func (_RegistryBinding *RegistryBindingSession) RecoverSigner(mes []byte, sgn []byte) (common.Address, error) {
	return _RegistryBinding.Contract.RecoverSigner(&_RegistryBinding.CallOpts, mes, sgn)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x2e295ec9.
//
// Solidity: function recoverSigner(bytes mes, bytes sgn) pure returns(address)
func (_RegistryBinding *RegistryBindingCallerSession) RecoverSigner(mes []byte, sgn []byte) (common.Address, error) {
	return _RegistryBinding.Contract.RecoverSigner(&_RegistryBinding.CallOpts, mes, sgn)
}

// RequestMessage is a free data retrieval call binding the contract method 0x25d6fea6.
//
// Solidity: function requestMessage(address grd, uint64 tim, address pla) pure returns(bytes)
func (_RegistryBinding *RegistryBindingCaller) RequestMessage(opts *bind.CallOpts, grd common.Address, tim uint64, pla common.Address) ([]byte, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "requestMessage", grd, tim, pla)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// RequestMessage is a free data retrieval call binding the contract method 0x25d6fea6.
//
// Solidity: function requestMessage(address grd, uint64 tim, address pla) pure returns(bytes)
func (_RegistryBinding *RegistryBindingSession) RequestMessage(grd common.Address, tim uint64, pla common.Address) ([]byte, error) {
	return _RegistryBinding.Contract.RequestMessage(&_RegistryBinding.CallOpts, grd, tim, pla)
}

// RequestMessage is a free data retrieval call binding the contract method 0x25d6fea6.
//
// Solidity: function requestMessage(address grd, uint64 tim, address pla) pure returns(bytes)
func (_RegistryBinding *RegistryBindingCallerSession) RequestMessage(grd common.Address, tim uint64, pla common.Address) ([]byte, error) {
	return _RegistryBinding.Contract.RequestMessage(&_RegistryBinding.CallOpts, grd, tim, pla)
}

// SearchBalance is a free data retrieval call binding the contract method 0x778c06f2.
//
// Solidity: function searchBalance(address wal) view returns(uint256, uint256, uint256)
func (_RegistryBinding *RegistryBindingCaller) SearchBalance(opts *bind.CallOpts, wal common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "searchBalance", wal)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// SearchBalance is a free data retrieval call binding the contract method 0x778c06f2.
//
// Solidity: function searchBalance(address wal) view returns(uint256, uint256, uint256)
func (_RegistryBinding *RegistryBindingSession) SearchBalance(wal common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _RegistryBinding.Contract.SearchBalance(&_RegistryBinding.CallOpts, wal)
}

// SearchBalance is a free data retrieval call binding the contract method 0x778c06f2.
//
// Solidity: function searchBalance(address wal) view returns(uint256, uint256, uint256)
func (_RegistryBinding *RegistryBindingCallerSession) SearchBalance(wal common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _RegistryBinding.Contract.SearchBalance(&_RegistryBinding.CallOpts, wal)
}

// SearchSigner is a free data retrieval call binding the contract method 0x35b3bf67.
//
// Solidity: function searchSigner(address wal) view returns(address, address)
func (_RegistryBinding *RegistryBindingCaller) SearchSigner(opts *bind.CallOpts, wal common.Address) (common.Address, common.Address, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "searchSigner", wal)

	if err != nil {
		return *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// SearchSigner is a free data retrieval call binding the contract method 0x35b3bf67.
//
// Solidity: function searchSigner(address wal) view returns(address, address)
func (_RegistryBinding *RegistryBindingSession) SearchSigner(wal common.Address) (common.Address, common.Address, error) {
	return _RegistryBinding.Contract.SearchSigner(&_RegistryBinding.CallOpts, wal)
}

// SearchSigner is a free data retrieval call binding the contract method 0x35b3bf67.
//
// Solidity: function searchSigner(address wal) view returns(address, address)
func (_RegistryBinding *RegistryBindingCallerSession) SearchSigner(wal common.Address) (common.Address, common.Address, error) {
	return _RegistryBinding.Contract.SearchSigner(&_RegistryBinding.CallOpts, wal)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RegistryBinding *RegistryBindingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RegistryBinding *RegistryBindingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RegistryBinding.Contract.SupportsInterface(&_RegistryBinding.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RegistryBinding *RegistryBindingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RegistryBinding.Contract.SupportsInterface(&_RegistryBinding.CallOpts, interfaceId)
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

// ValHash is a free data retrieval call binding the contract method 0xe4827e2a.
//
// Solidity: function valHash(address win, address los) pure returns(bytes32)
func (_RegistryBinding *RegistryBindingCaller) ValHash(opts *bind.CallOpts, win common.Address, los common.Address) ([32]byte, error) {
	var out []interface{}
	err := _RegistryBinding.contract.Call(opts, &out, "valHash", win, los)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ValHash is a free data retrieval call binding the contract method 0xe4827e2a.
//
// Solidity: function valHash(address win, address los) pure returns(bytes32)
func (_RegistryBinding *RegistryBindingSession) ValHash(win common.Address, los common.Address) ([32]byte, error) {
	return _RegistryBinding.Contract.ValHash(&_RegistryBinding.CallOpts, win, los)
}

// ValHash is a free data retrieval call binding the contract method 0xe4827e2a.
//
// Solidity: function valHash(address win, address los) pure returns(bytes32)
func (_RegistryBinding *RegistryBindingCallerSession) ValHash(win common.Address, los common.Address) ([32]byte, error) {
	return _RegistryBinding.Contract.ValHash(&_RegistryBinding.CallOpts, win, los)
}

// Deposit is a paid mutator transaction binding the contract method 0x94cf840b.
//
// Solidity: function deposit(uint256 bal, uint64 tim, address sig, bytes sgn) returns()
func (_RegistryBinding *RegistryBindingTransactor) Deposit(opts *bind.TransactOpts, bal *big.Int, tim uint64, sig common.Address, sgn []byte) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "deposit", bal, tim, sig, sgn)
}

// Deposit is a paid mutator transaction binding the contract method 0x94cf840b.
//
// Solidity: function deposit(uint256 bal, uint64 tim, address sig, bytes sgn) returns()
func (_RegistryBinding *RegistryBindingSession) Deposit(bal *big.Int, tim uint64, sig common.Address, sgn []byte) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Deposit(&_RegistryBinding.TransactOpts, bal, tim, sig, sgn)
}

// Deposit is a paid mutator transaction binding the contract method 0x94cf840b.
//
// Solidity: function deposit(uint256 bal, uint64 tim, address sig, bytes sgn) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) Deposit(bal *big.Int, tim uint64, sig common.Address, sgn []byte) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Deposit(&_RegistryBinding.TransactOpts, bal, tim, sig, sgn)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RegistryBinding *RegistryBindingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RegistryBinding *RegistryBindingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.GrantRole(&_RegistryBinding.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.GrantRole(&_RegistryBinding.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_RegistryBinding *RegistryBindingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_RegistryBinding *RegistryBindingSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.RenounceRole(&_RegistryBinding.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.RenounceRole(&_RegistryBinding.TransactOpts, role, callerConfirmation)
}

// Report is a paid mutator transaction binding the contract method 0x80e50744.
//
// Solidity: function report(address grd, uint256 kil, address win, address los) returns()
func (_RegistryBinding *RegistryBindingTransactor) Report(opts *bind.TransactOpts, grd common.Address, kil *big.Int, win common.Address, los common.Address) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "report", grd, kil, win, los)
}

// Report is a paid mutator transaction binding the contract method 0x80e50744.
//
// Solidity: function report(address grd, uint256 kil, address win, address los) returns()
func (_RegistryBinding *RegistryBindingSession) Report(grd common.Address, kil *big.Int, win common.Address, los common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Report(&_RegistryBinding.TransactOpts, grd, kil, win, los)
}

// Report is a paid mutator transaction binding the contract method 0x80e50744.
//
// Solidity: function report(address grd, uint256 kil, address win, address los) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) Report(grd common.Address, kil *big.Int, win common.Address, los common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Report(&_RegistryBinding.TransactOpts, grd, kil, win, los)
}

// Request is a paid mutator transaction binding the contract method 0x383cb927.
//
// Solidity: function request(address grd, uint64 tim, address wal, bytes sgn) returns()
func (_RegistryBinding *RegistryBindingTransactor) Request(opts *bind.TransactOpts, grd common.Address, tim uint64, wal common.Address, sgn []byte) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "request", grd, tim, wal, sgn)
}

// Request is a paid mutator transaction binding the contract method 0x383cb927.
//
// Solidity: function request(address grd, uint64 tim, address wal, bytes sgn) returns()
func (_RegistryBinding *RegistryBindingSession) Request(grd common.Address, tim uint64, wal common.Address, sgn []byte) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Request(&_RegistryBinding.TransactOpts, grd, tim, wal, sgn)
}

// Request is a paid mutator transaction binding the contract method 0x383cb927.
//
// Solidity: function request(address grd, uint64 tim, address wal, bytes sgn) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) Request(grd common.Address, tim uint64, wal common.Address, sgn []byte) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Request(&_RegistryBinding.TransactOpts, grd, tim, wal, sgn)
}

// Resolve is a paid mutator transaction binding the contract method 0x722dcc48.
//
// Solidity: function resolve(uint256 kil, address win, address los) returns()
func (_RegistryBinding *RegistryBindingTransactor) Resolve(opts *bind.TransactOpts, kil *big.Int, win common.Address, los common.Address) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "resolve", kil, win, los)
}

// Resolve is a paid mutator transaction binding the contract method 0x722dcc48.
//
// Solidity: function resolve(uint256 kil, address win, address los) returns()
func (_RegistryBinding *RegistryBindingSession) Resolve(kil *big.Int, win common.Address, los common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Resolve(&_RegistryBinding.TransactOpts, kil, win, los)
}

// Resolve is a paid mutator transaction binding the contract method 0x722dcc48.
//
// Solidity: function resolve(uint256 kil, address win, address los) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) Resolve(kil *big.Int, win common.Address, los common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Resolve(&_RegistryBinding.TransactOpts, kil, win, los)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RegistryBinding *RegistryBindingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RegistryBinding *RegistryBindingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.RevokeRole(&_RegistryBinding.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegistryBinding.Contract.RevokeRole(&_RegistryBinding.TransactOpts, role, account)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 bal) returns()
func (_RegistryBinding *RegistryBindingTransactor) Withdraw(opts *bind.TransactOpts, bal *big.Int) (*types.Transaction, error) {
	return _RegistryBinding.contract.Transact(opts, "withdraw", bal)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 bal) returns()
func (_RegistryBinding *RegistryBindingSession) Withdraw(bal *big.Int) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Withdraw(&_RegistryBinding.TransactOpts, bal)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 bal) returns()
func (_RegistryBinding *RegistryBindingTransactorSession) Withdraw(bal *big.Int) (*types.Transaction, error) {
	return _RegistryBinding.Contract.Withdraw(&_RegistryBinding.TransactOpts, bal)
}

// RegistryBindingReportIterator is returned from FilterReport and is used to iterate over the raw logs and unpacked data for Report events raised by the RegistryBinding contract.
type RegistryBindingReportIterator struct {
	Event *RegistryBindingReport // Event containing the contract specifics and raw log

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
func (it *RegistryBindingReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryBindingReport)
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
		it.Event = new(RegistryBindingReport)
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
func (it *RegistryBindingReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryBindingReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryBindingReport represents a Report event raised by the RegistryBinding contract.
type RegistryBindingReport struct {
	Pla common.Address
	Grd common.Address
	Kil *big.Int
	Win common.Address
	Los common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterReport is a free log retrieval operation binding the contract event 0x8d50973511534af055090ee47201c6f386b11a4b61ad4665a05b955013da9c6c.
//
// Solidity: event Report(address indexed pla, address indexed grd, uint256 kil, address win, address los)
func (_RegistryBinding *RegistryBindingFilterer) FilterReport(opts *bind.FilterOpts, pla []common.Address, grd []common.Address) (*RegistryBindingReportIterator, error) {

	var plaRule []interface{}
	for _, plaItem := range pla {
		plaRule = append(plaRule, plaItem)
	}
	var grdRule []interface{}
	for _, grdItem := range grd {
		grdRule = append(grdRule, grdItem)
	}

	logs, sub, err := _RegistryBinding.contract.FilterLogs(opts, "Report", plaRule, grdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryBindingReportIterator{contract: _RegistryBinding.contract, event: "Report", logs: logs, sub: sub}, nil
}

// WatchReport is a free log subscription operation binding the contract event 0x8d50973511534af055090ee47201c6f386b11a4b61ad4665a05b955013da9c6c.
//
// Solidity: event Report(address indexed pla, address indexed grd, uint256 kil, address win, address los)
func (_RegistryBinding *RegistryBindingFilterer) WatchReport(opts *bind.WatchOpts, sink chan<- *RegistryBindingReport, pla []common.Address, grd []common.Address) (event.Subscription, error) {

	var plaRule []interface{}
	for _, plaItem := range pla {
		plaRule = append(plaRule, plaItem)
	}
	var grdRule []interface{}
	for _, grdItem := range grd {
		grdRule = append(grdRule, grdItem)
	}

	logs, sub, err := _RegistryBinding.contract.WatchLogs(opts, "Report", plaRule, grdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryBindingReport)
				if err := _RegistryBinding.contract.UnpackLog(event, "Report", log); err != nil {
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

// ParseReport is a log parse operation binding the contract event 0x8d50973511534af055090ee47201c6f386b11a4b61ad4665a05b955013da9c6c.
//
// Solidity: event Report(address indexed pla, address indexed grd, uint256 kil, address win, address los)
func (_RegistryBinding *RegistryBindingFilterer) ParseReport(log types.Log) (*RegistryBindingReport, error) {
	event := new(RegistryBindingReport)
	if err := _RegistryBinding.contract.UnpackLog(event, "Report", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryBindingResolveIterator is returned from FilterResolve and is used to iterate over the raw logs and unpacked data for Resolve events raised by the RegistryBinding contract.
type RegistryBindingResolveIterator struct {
	Event *RegistryBindingResolve // Event containing the contract specifics and raw log

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
func (it *RegistryBindingResolveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryBindingResolve)
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
		it.Event = new(RegistryBindingResolve)
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
func (it *RegistryBindingResolveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryBindingResolveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryBindingResolve represents a Resolve event raised by the RegistryBinding contract.
type RegistryBindingResolve struct {
	Grd common.Address
	Kil *big.Int
	Win common.Address
	Los common.Address
	Dep *big.Int
	Bin *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResolve is a free log retrieval operation binding the contract event 0xbed1befae87ab8df925e4bb84fac5c0d21dbbc27a1b9d379f3f731f29b0a2cde.
//
// Solidity: event Resolve(address indexed grd, uint256 kil, address win, address los, uint256 dep, uint256 bin)
func (_RegistryBinding *RegistryBindingFilterer) FilterResolve(opts *bind.FilterOpts, grd []common.Address) (*RegistryBindingResolveIterator, error) {

	var grdRule []interface{}
	for _, grdItem := range grd {
		grdRule = append(grdRule, grdItem)
	}

	logs, sub, err := _RegistryBinding.contract.FilterLogs(opts, "Resolve", grdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryBindingResolveIterator{contract: _RegistryBinding.contract, event: "Resolve", logs: logs, sub: sub}, nil
}

// WatchResolve is a free log subscription operation binding the contract event 0xbed1befae87ab8df925e4bb84fac5c0d21dbbc27a1b9d379f3f731f29b0a2cde.
//
// Solidity: event Resolve(address indexed grd, uint256 kil, address win, address los, uint256 dep, uint256 bin)
func (_RegistryBinding *RegistryBindingFilterer) WatchResolve(opts *bind.WatchOpts, sink chan<- *RegistryBindingResolve, grd []common.Address) (event.Subscription, error) {

	var grdRule []interface{}
	for _, grdItem := range grd {
		grdRule = append(grdRule, grdItem)
	}

	logs, sub, err := _RegistryBinding.contract.WatchLogs(opts, "Resolve", grdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryBindingResolve)
				if err := _RegistryBinding.contract.UnpackLog(event, "Resolve", log); err != nil {
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

// ParseResolve is a log parse operation binding the contract event 0xbed1befae87ab8df925e4bb84fac5c0d21dbbc27a1b9d379f3f731f29b0a2cde.
//
// Solidity: event Resolve(address indexed grd, uint256 kil, address win, address los, uint256 dep, uint256 bin)
func (_RegistryBinding *RegistryBindingFilterer) ParseResolve(log types.Log) (*RegistryBindingResolve, error) {
	event := new(RegistryBindingResolve)
	if err := _RegistryBinding.contract.UnpackLog(event, "Resolve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryBindingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RegistryBinding contract.
type RegistryBindingRoleAdminChangedIterator struct {
	Event *RegistryBindingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RegistryBindingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryBindingRoleAdminChanged)
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
		it.Event = new(RegistryBindingRoleAdminChanged)
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
func (it *RegistryBindingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryBindingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryBindingRoleAdminChanged represents a RoleAdminChanged event raised by the RegistryBinding contract.
type RegistryBindingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RegistryBinding *RegistryBindingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RegistryBindingRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RegistryBinding.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RegistryBindingRoleAdminChangedIterator{contract: _RegistryBinding.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RegistryBinding *RegistryBindingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RegistryBindingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RegistryBinding.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryBindingRoleAdminChanged)
				if err := _RegistryBinding.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RegistryBinding *RegistryBindingFilterer) ParseRoleAdminChanged(log types.Log) (*RegistryBindingRoleAdminChanged, error) {
	event := new(RegistryBindingRoleAdminChanged)
	if err := _RegistryBinding.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryBindingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RegistryBinding contract.
type RegistryBindingRoleGrantedIterator struct {
	Event *RegistryBindingRoleGranted // Event containing the contract specifics and raw log

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
func (it *RegistryBindingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryBindingRoleGranted)
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
		it.Event = new(RegistryBindingRoleGranted)
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
func (it *RegistryBindingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryBindingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryBindingRoleGranted represents a RoleGranted event raised by the RegistryBinding contract.
type RegistryBindingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RegistryBinding *RegistryBindingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegistryBindingRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RegistryBinding.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryBindingRoleGrantedIterator{contract: _RegistryBinding.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RegistryBinding *RegistryBindingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RegistryBindingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RegistryBinding.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryBindingRoleGranted)
				if err := _RegistryBinding.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RegistryBinding *RegistryBindingFilterer) ParseRoleGranted(log types.Log) (*RegistryBindingRoleGranted, error) {
	event := new(RegistryBindingRoleGranted)
	if err := _RegistryBinding.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryBindingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RegistryBinding contract.
type RegistryBindingRoleRevokedIterator struct {
	Event *RegistryBindingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RegistryBindingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryBindingRoleRevoked)
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
		it.Event = new(RegistryBindingRoleRevoked)
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
func (it *RegistryBindingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryBindingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryBindingRoleRevoked represents a RoleRevoked event raised by the RegistryBinding contract.
type RegistryBindingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RegistryBinding *RegistryBindingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegistryBindingRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RegistryBinding.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryBindingRoleRevokedIterator{contract: _RegistryBinding.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RegistryBinding *RegistryBindingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RegistryBindingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RegistryBinding.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryBindingRoleRevoked)
				if err := _RegistryBinding.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RegistryBinding *RegistryBindingFilterer) ParseRoleRevoked(log types.Log) (*RegistryBindingRoleRevoked, error) {
	event := new(RegistryBindingRoleRevoked)
	if err := _RegistryBinding.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
