// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package huntergovernancetoken

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

// HuntergovernancetokenMetaData contains all meta data concerning the Huntergovernancetoken contract.
var HuntergovernancetokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_premintReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_premintAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUPPLY_CAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HuntergovernancetokenABI is the input ABI used to generate the binding from.
// Deprecated: Use HuntergovernancetokenMetaData.ABI instead.
var HuntergovernancetokenABI = HuntergovernancetokenMetaData.ABI

// Huntergovernancetoken is an auto generated Go binding around an Ethereum contract.
type Huntergovernancetoken struct {
	HuntergovernancetokenCaller     // Read-only binding to the contract
	HuntergovernancetokenTransactor // Write-only binding to the contract
	HuntergovernancetokenFilterer   // Log filterer for contract events
}

// HuntergovernancetokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type HuntergovernancetokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HuntergovernancetokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HuntergovernancetokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HuntergovernancetokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HuntergovernancetokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HuntergovernancetokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HuntergovernancetokenSession struct {
	Contract     *Huntergovernancetoken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HuntergovernancetokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HuntergovernancetokenCallerSession struct {
	Contract *HuntergovernancetokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// HuntergovernancetokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HuntergovernancetokenTransactorSession struct {
	Contract     *HuntergovernancetokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// HuntergovernancetokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type HuntergovernancetokenRaw struct {
	Contract *Huntergovernancetoken // Generic contract binding to access the raw methods on
}

// HuntergovernancetokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HuntergovernancetokenCallerRaw struct {
	Contract *HuntergovernancetokenCaller // Generic read-only contract binding to access the raw methods on
}

// HuntergovernancetokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HuntergovernancetokenTransactorRaw struct {
	Contract *HuntergovernancetokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHuntergovernancetoken creates a new instance of Huntergovernancetoken, bound to a specific deployed contract.
func NewHuntergovernancetoken(address common.Address, backend bind.ContractBackend) (*Huntergovernancetoken, error) {
	contract, err := bindHuntergovernancetoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Huntergovernancetoken{HuntergovernancetokenCaller: HuntergovernancetokenCaller{contract: contract}, HuntergovernancetokenTransactor: HuntergovernancetokenTransactor{contract: contract}, HuntergovernancetokenFilterer: HuntergovernancetokenFilterer{contract: contract}}, nil
}

// NewHuntergovernancetokenCaller creates a new read-only instance of Huntergovernancetoken, bound to a specific deployed contract.
func NewHuntergovernancetokenCaller(address common.Address, caller bind.ContractCaller) (*HuntergovernancetokenCaller, error) {
	contract, err := bindHuntergovernancetoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HuntergovernancetokenCaller{contract: contract}, nil
}

// NewHuntergovernancetokenTransactor creates a new write-only instance of Huntergovernancetoken, bound to a specific deployed contract.
func NewHuntergovernancetokenTransactor(address common.Address, transactor bind.ContractTransactor) (*HuntergovernancetokenTransactor, error) {
	contract, err := bindHuntergovernancetoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HuntergovernancetokenTransactor{contract: contract}, nil
}

// NewHuntergovernancetokenFilterer creates a new log filterer instance of Huntergovernancetoken, bound to a specific deployed contract.
func NewHuntergovernancetokenFilterer(address common.Address, filterer bind.ContractFilterer) (*HuntergovernancetokenFilterer, error) {
	contract, err := bindHuntergovernancetoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HuntergovernancetokenFilterer{contract: contract}, nil
}

// bindHuntergovernancetoken binds a generic wrapper to an already deployed contract.
func bindHuntergovernancetoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HuntergovernancetokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Huntergovernancetoken *HuntergovernancetokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Huntergovernancetoken.Contract.HuntergovernancetokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Huntergovernancetoken *HuntergovernancetokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.HuntergovernancetokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Huntergovernancetoken *HuntergovernancetokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.HuntergovernancetokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Huntergovernancetoken *HuntergovernancetokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Huntergovernancetoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Huntergovernancetoken *HuntergovernancetokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Huntergovernancetoken *HuntergovernancetokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.contract.Transact(opts, method, params...)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Huntergovernancetoken *HuntergovernancetokenCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Huntergovernancetoken.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Huntergovernancetoken *HuntergovernancetokenSession) OPERATORROLE() ([32]byte, error) {
	return _Huntergovernancetoken.Contract.OPERATORROLE(&_Huntergovernancetoken.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Huntergovernancetoken *HuntergovernancetokenCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Huntergovernancetoken.Contract.OPERATORROLE(&_Huntergovernancetoken.CallOpts)
}

// SUPPLYCAP is a free data retrieval call binding the contract method 0x0cfccc83.
//
// Solidity: function SUPPLY_CAP() pure returns(uint256)
func (_Huntergovernancetoken *HuntergovernancetokenCaller) SUPPLYCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Huntergovernancetoken.contract.Call(opts, &out, "SUPPLY_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUPPLYCAP is a free data retrieval call binding the contract method 0x0cfccc83.
//
// Solidity: function SUPPLY_CAP() pure returns(uint256)
func (_Huntergovernancetoken *HuntergovernancetokenSession) SUPPLYCAP() (*big.Int, error) {
	return _Huntergovernancetoken.Contract.SUPPLYCAP(&_Huntergovernancetoken.CallOpts)
}

// SUPPLYCAP is a free data retrieval call binding the contract method 0x0cfccc83.
//
// Solidity: function SUPPLY_CAP() pure returns(uint256)
func (_Huntergovernancetoken *HuntergovernancetokenCallerSession) SUPPLYCAP() (*big.Int, error) {
	return _Huntergovernancetoken.Contract.SUPPLYCAP(&_Huntergovernancetoken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Huntergovernancetoken *HuntergovernancetokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Huntergovernancetoken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Huntergovernancetoken *HuntergovernancetokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Huntergovernancetoken.Contract.Allowance(&_Huntergovernancetoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Huntergovernancetoken *HuntergovernancetokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Huntergovernancetoken.Contract.Allowance(&_Huntergovernancetoken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Huntergovernancetoken *HuntergovernancetokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Huntergovernancetoken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Huntergovernancetoken *HuntergovernancetokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Huntergovernancetoken.Contract.BalanceOf(&_Huntergovernancetoken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Huntergovernancetoken *HuntergovernancetokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Huntergovernancetoken.Contract.BalanceOf(&_Huntergovernancetoken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Huntergovernancetoken *HuntergovernancetokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Huntergovernancetoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Huntergovernancetoken *HuntergovernancetokenSession) Decimals() (uint8, error) {
	return _Huntergovernancetoken.Contract.Decimals(&_Huntergovernancetoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Huntergovernancetoken *HuntergovernancetokenCallerSession) Decimals() (uint8, error) {
	return _Huntergovernancetoken.Contract.Decimals(&_Huntergovernancetoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Huntergovernancetoken *HuntergovernancetokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Huntergovernancetoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Huntergovernancetoken *HuntergovernancetokenSession) Name() (string, error) {
	return _Huntergovernancetoken.Contract.Name(&_Huntergovernancetoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Huntergovernancetoken *HuntergovernancetokenCallerSession) Name() (string, error) {
	return _Huntergovernancetoken.Contract.Name(&_Huntergovernancetoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Huntergovernancetoken *HuntergovernancetokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Huntergovernancetoken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Huntergovernancetoken *HuntergovernancetokenSession) Owner() (common.Address, error) {
	return _Huntergovernancetoken.Contract.Owner(&_Huntergovernancetoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Huntergovernancetoken *HuntergovernancetokenCallerSession) Owner() (common.Address, error) {
	return _Huntergovernancetoken.Contract.Owner(&_Huntergovernancetoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Huntergovernancetoken *HuntergovernancetokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Huntergovernancetoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Huntergovernancetoken *HuntergovernancetokenSession) Symbol() (string, error) {
	return _Huntergovernancetoken.Contract.Symbol(&_Huntergovernancetoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Huntergovernancetoken *HuntergovernancetokenCallerSession) Symbol() (string, error) {
	return _Huntergovernancetoken.Contract.Symbol(&_Huntergovernancetoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Huntergovernancetoken *HuntergovernancetokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Huntergovernancetoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Huntergovernancetoken *HuntergovernancetokenSession) TotalSupply() (*big.Int, error) {
	return _Huntergovernancetoken.Contract.TotalSupply(&_Huntergovernancetoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Huntergovernancetoken *HuntergovernancetokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Huntergovernancetoken.Contract.TotalSupply(&_Huntergovernancetoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Huntergovernancetoken *HuntergovernancetokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Huntergovernancetoken *HuntergovernancetokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.Approve(&_Huntergovernancetoken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Huntergovernancetoken *HuntergovernancetokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.Approve(&_Huntergovernancetoken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Huntergovernancetoken *HuntergovernancetokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Huntergovernancetoken *HuntergovernancetokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.DecreaseAllowance(&_Huntergovernancetoken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Huntergovernancetoken *HuntergovernancetokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.DecreaseAllowance(&_Huntergovernancetoken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Huntergovernancetoken *HuntergovernancetokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Huntergovernancetoken *HuntergovernancetokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.IncreaseAllowance(&_Huntergovernancetoken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Huntergovernancetoken *HuntergovernancetokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.IncreaseAllowance(&_Huntergovernancetoken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool status)
func (_Huntergovernancetoken *HuntergovernancetokenTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool status)
func (_Huntergovernancetoken *HuntergovernancetokenSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.Mint(&_Huntergovernancetoken.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool status)
func (_Huntergovernancetoken *HuntergovernancetokenTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.Mint(&_Huntergovernancetoken.TransactOpts, account, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Huntergovernancetoken *HuntergovernancetokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Huntergovernancetoken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Huntergovernancetoken *HuntergovernancetokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.RenounceOwnership(&_Huntergovernancetoken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Huntergovernancetoken *HuntergovernancetokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.RenounceOwnership(&_Huntergovernancetoken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Huntergovernancetoken *HuntergovernancetokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Huntergovernancetoken *HuntergovernancetokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.Transfer(&_Huntergovernancetoken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Huntergovernancetoken *HuntergovernancetokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.Transfer(&_Huntergovernancetoken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Huntergovernancetoken *HuntergovernancetokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Huntergovernancetoken *HuntergovernancetokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.TransferFrom(&_Huntergovernancetoken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Huntergovernancetoken *HuntergovernancetokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.TransferFrom(&_Huntergovernancetoken.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Huntergovernancetoken *HuntergovernancetokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Huntergovernancetoken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Huntergovernancetoken *HuntergovernancetokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.TransferOwnership(&_Huntergovernancetoken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Huntergovernancetoken *HuntergovernancetokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Huntergovernancetoken.Contract.TransferOwnership(&_Huntergovernancetoken.TransactOpts, newOwner)
}

// HuntergovernancetokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Huntergovernancetoken contract.
type HuntergovernancetokenApprovalIterator struct {
	Event *HuntergovernancetokenApproval // Event containing the contract specifics and raw log

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
func (it *HuntergovernancetokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HuntergovernancetokenApproval)
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
		it.Event = new(HuntergovernancetokenApproval)
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
func (it *HuntergovernancetokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HuntergovernancetokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HuntergovernancetokenApproval represents a Approval event raised by the Huntergovernancetoken contract.
type HuntergovernancetokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Huntergovernancetoken *HuntergovernancetokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*HuntergovernancetokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Huntergovernancetoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &HuntergovernancetokenApprovalIterator{contract: _Huntergovernancetoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Huntergovernancetoken *HuntergovernancetokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HuntergovernancetokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Huntergovernancetoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HuntergovernancetokenApproval)
				if err := _Huntergovernancetoken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Huntergovernancetoken *HuntergovernancetokenFilterer) ParseApproval(log types.Log) (*HuntergovernancetokenApproval, error) {
	event := new(HuntergovernancetokenApproval)
	if err := _Huntergovernancetoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HuntergovernancetokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Huntergovernancetoken contract.
type HuntergovernancetokenOwnershipTransferredIterator struct {
	Event *HuntergovernancetokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HuntergovernancetokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HuntergovernancetokenOwnershipTransferred)
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
		it.Event = new(HuntergovernancetokenOwnershipTransferred)
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
func (it *HuntergovernancetokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HuntergovernancetokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HuntergovernancetokenOwnershipTransferred represents a OwnershipTransferred event raised by the Huntergovernancetoken contract.
type HuntergovernancetokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Huntergovernancetoken *HuntergovernancetokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HuntergovernancetokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Huntergovernancetoken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HuntergovernancetokenOwnershipTransferredIterator{contract: _Huntergovernancetoken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Huntergovernancetoken *HuntergovernancetokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HuntergovernancetokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Huntergovernancetoken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HuntergovernancetokenOwnershipTransferred)
				if err := _Huntergovernancetoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Huntergovernancetoken *HuntergovernancetokenFilterer) ParseOwnershipTransferred(log types.Log) (*HuntergovernancetokenOwnershipTransferred, error) {
	event := new(HuntergovernancetokenOwnershipTransferred)
	if err := _Huntergovernancetoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HuntergovernancetokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Huntergovernancetoken contract.
type HuntergovernancetokenTransferIterator struct {
	Event *HuntergovernancetokenTransfer // Event containing the contract specifics and raw log

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
func (it *HuntergovernancetokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HuntergovernancetokenTransfer)
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
		it.Event = new(HuntergovernancetokenTransfer)
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
func (it *HuntergovernancetokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HuntergovernancetokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HuntergovernancetokenTransfer represents a Transfer event raised by the Huntergovernancetoken contract.
type HuntergovernancetokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Huntergovernancetoken *HuntergovernancetokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HuntergovernancetokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Huntergovernancetoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HuntergovernancetokenTransferIterator{contract: _Huntergovernancetoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Huntergovernancetoken *HuntergovernancetokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HuntergovernancetokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Huntergovernancetoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HuntergovernancetokenTransfer)
				if err := _Huntergovernancetoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Huntergovernancetoken *HuntergovernancetokenFilterer) ParseTransfer(log types.Log) (*HuntergovernancetokenTransfer, error) {
	event := new(HuntergovernancetokenTransfer)
	if err := _Huntergovernancetoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
