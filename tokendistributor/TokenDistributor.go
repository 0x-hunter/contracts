// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokendistributor

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

// TokendistributorMetaData contains all meta data concerning the Tokendistributor contract.
var TokendistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_looksRareToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenSplitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_rewardsPerBlockForStaking\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_rewardsPerBlockForOthers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_periodLengthesInBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_numberPeriods\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestedAmount\",\"type\":\"uint256\"}],\"name\":\"Compound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestedAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"currentPhase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerBlockForStaking\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerBlockForOthers\",\"type\":\"uint256\"}],\"name\":\"NewRewardsPerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestedAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUMBER_PERIODS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"START_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accTokenPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"calculatePendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPhase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestAndCompound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"looksRareToken\",\"outputs\":[{\"internalType\":\"contractIHunterGovernanceToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerBlockForOthers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerBlockForStaking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardPerBlockForStaking\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerBlockForOthers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periodLengthInBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenSplitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAmountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokendistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use TokendistributorMetaData.ABI instead.
var TokendistributorABI = TokendistributorMetaData.ABI

// Tokendistributor is an auto generated Go binding around an Ethereum contract.
type Tokendistributor struct {
	TokendistributorCaller     // Read-only binding to the contract
	TokendistributorTransactor // Write-only binding to the contract
	TokendistributorFilterer   // Log filterer for contract events
}

// TokendistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokendistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokendistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokendistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokendistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokendistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokendistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokendistributorSession struct {
	Contract     *Tokendistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokendistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokendistributorCallerSession struct {
	Contract *TokendistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TokendistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokendistributorTransactorSession struct {
	Contract     *TokendistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TokendistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokendistributorRaw struct {
	Contract *Tokendistributor // Generic contract binding to access the raw methods on
}

// TokendistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokendistributorCallerRaw struct {
	Contract *TokendistributorCaller // Generic read-only contract binding to access the raw methods on
}

// TokendistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokendistributorTransactorRaw struct {
	Contract *TokendistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokendistributor creates a new instance of Tokendistributor, bound to a specific deployed contract.
func NewTokendistributor(address common.Address, backend bind.ContractBackend) (*Tokendistributor, error) {
	contract, err := bindTokendistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokendistributor{TokendistributorCaller: TokendistributorCaller{contract: contract}, TokendistributorTransactor: TokendistributorTransactor{contract: contract}, TokendistributorFilterer: TokendistributorFilterer{contract: contract}}, nil
}

// NewTokendistributorCaller creates a new read-only instance of Tokendistributor, bound to a specific deployed contract.
func NewTokendistributorCaller(address common.Address, caller bind.ContractCaller) (*TokendistributorCaller, error) {
	contract, err := bindTokendistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokendistributorCaller{contract: contract}, nil
}

// NewTokendistributorTransactor creates a new write-only instance of Tokendistributor, bound to a specific deployed contract.
func NewTokendistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*TokendistributorTransactor, error) {
	contract, err := bindTokendistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokendistributorTransactor{contract: contract}, nil
}

// NewTokendistributorFilterer creates a new log filterer instance of Tokendistributor, bound to a specific deployed contract.
func NewTokendistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*TokendistributorFilterer, error) {
	contract, err := bindTokendistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokendistributorFilterer{contract: contract}, nil
}

// bindTokendistributor binds a generic wrapper to an already deployed contract.
func bindTokendistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokendistributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokendistributor *TokendistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokendistributor.Contract.TokendistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokendistributor *TokendistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokendistributor.Contract.TokendistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokendistributor *TokendistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokendistributor.Contract.TokendistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokendistributor *TokendistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokendistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokendistributor *TokendistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokendistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokendistributor *TokendistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokendistributor.Contract.contract.Transact(opts, method, params...)
}

// NUMBERPERIODS is a free data retrieval call binding the contract method 0x52bf348c.
//
// Solidity: function NUMBER_PERIODS() view returns(uint256)
func (_Tokendistributor *TokendistributorCaller) NUMBERPERIODS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokendistributor.contract.Call(opts, &out, "NUMBER_PERIODS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NUMBERPERIODS is a free data retrieval call binding the contract method 0x52bf348c.
//
// Solidity: function NUMBER_PERIODS() view returns(uint256)
func (_Tokendistributor *TokendistributorSession) NUMBERPERIODS() (*big.Int, error) {
	return _Tokendistributor.Contract.NUMBERPERIODS(&_Tokendistributor.CallOpts)
}

// NUMBERPERIODS is a free data retrieval call binding the contract method 0x52bf348c.
//
// Solidity: function NUMBER_PERIODS() view returns(uint256)
func (_Tokendistributor *TokendistributorCallerSession) NUMBERPERIODS() (*big.Int, error) {
	return _Tokendistributor.Contract.NUMBERPERIODS(&_Tokendistributor.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Tokendistributor *TokendistributorCaller) PRECISIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokendistributor.contract.Call(opts, &out, "PRECISION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Tokendistributor *TokendistributorSession) PRECISIONFACTOR() (*big.Int, error) {
	return _Tokendistributor.Contract.PRECISIONFACTOR(&_Tokendistributor.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Tokendistributor *TokendistributorCallerSession) PRECISIONFACTOR() (*big.Int, error) {
	return _Tokendistributor.Contract.PRECISIONFACTOR(&_Tokendistributor.CallOpts)
}

// STARTBLOCK is a free data retrieval call binding the contract method 0x39b3e826.
//
// Solidity: function START_BLOCK() view returns(uint256)
func (_Tokendistributor *TokendistributorCaller) STARTBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokendistributor.contract.Call(opts, &out, "START_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STARTBLOCK is a free data retrieval call binding the contract method 0x39b3e826.
//
// Solidity: function START_BLOCK() view returns(uint256)
func (_Tokendistributor *TokendistributorSession) STARTBLOCK() (*big.Int, error) {
	return _Tokendistributor.Contract.STARTBLOCK(&_Tokendistributor.CallOpts)
}

// STARTBLOCK is a free data retrieval call binding the contract method 0x39b3e826.
//
// Solidity: function START_BLOCK() view returns(uint256)
func (_Tokendistributor *TokendistributorCallerSession) STARTBLOCK() (*big.Int, error) {
	return _Tokendistributor.Contract.STARTBLOCK(&_Tokendistributor.CallOpts)
}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_Tokendistributor *TokendistributorCaller) AccTokenPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokendistributor.contract.Call(opts, &out, "accTokenPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_Tokendistributor *TokendistributorSession) AccTokenPerShare() (*big.Int, error) {
	return _Tokendistributor.Contract.AccTokenPerShare(&_Tokendistributor.CallOpts)
}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_Tokendistributor *TokendistributorCallerSession) AccTokenPerShare() (*big.Int, error) {
	return _Tokendistributor.Contract.AccTokenPerShare(&_Tokendistributor.CallOpts)
}

// CalculatePendingRewards is a free data retrieval call binding the contract method 0x097aad10.
//
// Solidity: function calculatePendingRewards(address user) view returns(uint256)
func (_Tokendistributor *TokendistributorCaller) CalculatePendingRewards(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokendistributor.contract.Call(opts, &out, "calculatePendingRewards", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePendingRewards is a free data retrieval call binding the contract method 0x097aad10.
//
// Solidity: function calculatePendingRewards(address user) view returns(uint256)
func (_Tokendistributor *TokendistributorSession) CalculatePendingRewards(user common.Address) (*big.Int, error) {
	return _Tokendistributor.Contract.CalculatePendingRewards(&_Tokendistributor.CallOpts, user)
}

// CalculatePendingRewards is a free data retrieval call binding the contract method 0x097aad10.
//
// Solidity: function calculatePendingRewards(address user) view returns(uint256)
func (_Tokendistributor *TokendistributorCallerSession) CalculatePendingRewards(user common.Address) (*big.Int, error) {
	return _Tokendistributor.Contract.CalculatePendingRewards(&_Tokendistributor.CallOpts, user)
}

// CurrentPhase is a free data retrieval call binding the contract method 0x055ad42e.
//
// Solidity: function currentPhase() view returns(uint256)
func (_Tokendistributor *TokendistributorCaller) CurrentPhase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokendistributor.contract.Call(opts, &out, "currentPhase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPhase is a free data retrieval call binding the contract method 0x055ad42e.
//
// Solidity: function currentPhase() view returns(uint256)
func (_Tokendistributor *TokendistributorSession) CurrentPhase() (*big.Int, error) {
	return _Tokendistributor.Contract.CurrentPhase(&_Tokendistributor.CallOpts)
}

// CurrentPhase is a free data retrieval call binding the contract method 0x055ad42e.
//
// Solidity: function currentPhase() view returns(uint256)
func (_Tokendistributor *TokendistributorCallerSession) CurrentPhase() (*big.Int, error) {
	return _Tokendistributor.Contract.CurrentPhase(&_Tokendistributor.CallOpts)
}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() view returns(uint256)
func (_Tokendistributor *TokendistributorCaller) EndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokendistributor.contract.Call(opts, &out, "endBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() view returns(uint256)
func (_Tokendistributor *TokendistributorSession) EndBlock() (*big.Int, error) {
	return _Tokendistributor.Contract.EndBlock(&_Tokendistributor.CallOpts)
}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() view returns(uint256)
func (_Tokendistributor *TokendistributorCallerSession) EndBlock() (*big.Int, error) {
	return _Tokendistributor.Contract.EndBlock(&_Tokendistributor.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Tokendistributor *TokendistributorCaller) LastRewardBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokendistributor.contract.Call(opts, &out, "lastRewardBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Tokendistributor *TokendistributorSession) LastRewardBlock() (*big.Int, error) {
	return _Tokendistributor.Contract.LastRewardBlock(&_Tokendistributor.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Tokendistributor *TokendistributorCallerSession) LastRewardBlock() (*big.Int, error) {
	return _Tokendistributor.Contract.LastRewardBlock(&_Tokendistributor.CallOpts)
}

// LooksRareToken is a free data retrieval call binding the contract method 0x36db9fb2.
//
// Solidity: function looksRareToken() view returns(address)
func (_Tokendistributor *TokendistributorCaller) LooksRareToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokendistributor.contract.Call(opts, &out, "looksRareToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LooksRareToken is a free data retrieval call binding the contract method 0x36db9fb2.
//
// Solidity: function looksRareToken() view returns(address)
func (_Tokendistributor *TokendistributorSession) LooksRareToken() (common.Address, error) {
	return _Tokendistributor.Contract.LooksRareToken(&_Tokendistributor.CallOpts)
}

// LooksRareToken is a free data retrieval call binding the contract method 0x36db9fb2.
//
// Solidity: function looksRareToken() view returns(address)
func (_Tokendistributor *TokendistributorCallerSession) LooksRareToken() (common.Address, error) {
	return _Tokendistributor.Contract.LooksRareToken(&_Tokendistributor.CallOpts)
}

// RewardPerBlockForOthers is a free data retrieval call binding the contract method 0xe683d96f.
//
// Solidity: function rewardPerBlockForOthers() view returns(uint256)
func (_Tokendistributor *TokendistributorCaller) RewardPerBlockForOthers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokendistributor.contract.Call(opts, &out, "rewardPerBlockForOthers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerBlockForOthers is a free data retrieval call binding the contract method 0xe683d96f.
//
// Solidity: function rewardPerBlockForOthers() view returns(uint256)
func (_Tokendistributor *TokendistributorSession) RewardPerBlockForOthers() (*big.Int, error) {
	return _Tokendistributor.Contract.RewardPerBlockForOthers(&_Tokendistributor.CallOpts)
}

// RewardPerBlockForOthers is a free data retrieval call binding the contract method 0xe683d96f.
//
// Solidity: function rewardPerBlockForOthers() view returns(uint256)
func (_Tokendistributor *TokendistributorCallerSession) RewardPerBlockForOthers() (*big.Int, error) {
	return _Tokendistributor.Contract.RewardPerBlockForOthers(&_Tokendistributor.CallOpts)
}

// RewardPerBlockForStaking is a free data retrieval call binding the contract method 0x5a9477e9.
//
// Solidity: function rewardPerBlockForStaking() view returns(uint256)
func (_Tokendistributor *TokendistributorCaller) RewardPerBlockForStaking(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokendistributor.contract.Call(opts, &out, "rewardPerBlockForStaking")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerBlockForStaking is a free data retrieval call binding the contract method 0x5a9477e9.
//
// Solidity: function rewardPerBlockForStaking() view returns(uint256)
func (_Tokendistributor *TokendistributorSession) RewardPerBlockForStaking() (*big.Int, error) {
	return _Tokendistributor.Contract.RewardPerBlockForStaking(&_Tokendistributor.CallOpts)
}

// RewardPerBlockForStaking is a free data retrieval call binding the contract method 0x5a9477e9.
//
// Solidity: function rewardPerBlockForStaking() view returns(uint256)
func (_Tokendistributor *TokendistributorCallerSession) RewardPerBlockForStaking() (*big.Int, error) {
	return _Tokendistributor.Contract.RewardPerBlockForStaking(&_Tokendistributor.CallOpts)
}

// StakingPeriod is a free data retrieval call binding the contract method 0xc1027c98.
//
// Solidity: function stakingPeriod(uint256 ) view returns(uint256 rewardPerBlockForStaking, uint256 rewardPerBlockForOthers, uint256 periodLengthInBlock)
func (_Tokendistributor *TokendistributorCaller) StakingPeriod(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RewardPerBlockForStaking *big.Int
	RewardPerBlockForOthers  *big.Int
	PeriodLengthInBlock      *big.Int
}, error) {
	var out []interface{}
	err := _Tokendistributor.contract.Call(opts, &out, "stakingPeriod", arg0)

	outstruct := new(struct {
		RewardPerBlockForStaking *big.Int
		RewardPerBlockForOthers  *big.Int
		PeriodLengthInBlock      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardPerBlockForStaking = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardPerBlockForOthers = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PeriodLengthInBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StakingPeriod is a free data retrieval call binding the contract method 0xc1027c98.
//
// Solidity: function stakingPeriod(uint256 ) view returns(uint256 rewardPerBlockForStaking, uint256 rewardPerBlockForOthers, uint256 periodLengthInBlock)
func (_Tokendistributor *TokendistributorSession) StakingPeriod(arg0 *big.Int) (struct {
	RewardPerBlockForStaking *big.Int
	RewardPerBlockForOthers  *big.Int
	PeriodLengthInBlock      *big.Int
}, error) {
	return _Tokendistributor.Contract.StakingPeriod(&_Tokendistributor.CallOpts, arg0)
}

// StakingPeriod is a free data retrieval call binding the contract method 0xc1027c98.
//
// Solidity: function stakingPeriod(uint256 ) view returns(uint256 rewardPerBlockForStaking, uint256 rewardPerBlockForOthers, uint256 periodLengthInBlock)
func (_Tokendistributor *TokendistributorCallerSession) StakingPeriod(arg0 *big.Int) (struct {
	RewardPerBlockForStaking *big.Int
	RewardPerBlockForOthers  *big.Int
	PeriodLengthInBlock      *big.Int
}, error) {
	return _Tokendistributor.Contract.StakingPeriod(&_Tokendistributor.CallOpts, arg0)
}

// TokenSplitter is a free data retrieval call binding the contract method 0xa46074c3.
//
// Solidity: function tokenSplitter() view returns(address)
func (_Tokendistributor *TokendistributorCaller) TokenSplitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokendistributor.contract.Call(opts, &out, "tokenSplitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenSplitter is a free data retrieval call binding the contract method 0xa46074c3.
//
// Solidity: function tokenSplitter() view returns(address)
func (_Tokendistributor *TokendistributorSession) TokenSplitter() (common.Address, error) {
	return _Tokendistributor.Contract.TokenSplitter(&_Tokendistributor.CallOpts)
}

// TokenSplitter is a free data retrieval call binding the contract method 0xa46074c3.
//
// Solidity: function tokenSplitter() view returns(address)
func (_Tokendistributor *TokendistributorCallerSession) TokenSplitter() (common.Address, error) {
	return _Tokendistributor.Contract.TokenSplitter(&_Tokendistributor.CallOpts)
}

// TotalAmountStaked is a free data retrieval call binding the contract method 0xfe961f61.
//
// Solidity: function totalAmountStaked() view returns(uint256)
func (_Tokendistributor *TokendistributorCaller) TotalAmountStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokendistributor.contract.Call(opts, &out, "totalAmountStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAmountStaked is a free data retrieval call binding the contract method 0xfe961f61.
//
// Solidity: function totalAmountStaked() view returns(uint256)
func (_Tokendistributor *TokendistributorSession) TotalAmountStaked() (*big.Int, error) {
	return _Tokendistributor.Contract.TotalAmountStaked(&_Tokendistributor.CallOpts)
}

// TotalAmountStaked is a free data retrieval call binding the contract method 0xfe961f61.
//
// Solidity: function totalAmountStaked() view returns(uint256)
func (_Tokendistributor *TokendistributorCallerSession) TotalAmountStaked() (*big.Int, error) {
	return _Tokendistributor.Contract.TotalAmountStaked(&_Tokendistributor.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Tokendistributor *TokendistributorCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _Tokendistributor.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Tokendistributor *TokendistributorSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Tokendistributor.Contract.UserInfo(&_Tokendistributor.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Tokendistributor *TokendistributorCallerSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Tokendistributor.Contract.UserInfo(&_Tokendistributor.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Tokendistributor *TokendistributorTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Tokendistributor.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Tokendistributor *TokendistributorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Tokendistributor.Contract.Deposit(&_Tokendistributor.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Tokendistributor *TokendistributorTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Tokendistributor.Contract.Deposit(&_Tokendistributor.TransactOpts, amount)
}

// HarvestAndCompound is a paid mutator transaction binding the contract method 0x2a4e051b.
//
// Solidity: function harvestAndCompound() returns()
func (_Tokendistributor *TokendistributorTransactor) HarvestAndCompound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokendistributor.contract.Transact(opts, "harvestAndCompound")
}

// HarvestAndCompound is a paid mutator transaction binding the contract method 0x2a4e051b.
//
// Solidity: function harvestAndCompound() returns()
func (_Tokendistributor *TokendistributorSession) HarvestAndCompound() (*types.Transaction, error) {
	return _Tokendistributor.Contract.HarvestAndCompound(&_Tokendistributor.TransactOpts)
}

// HarvestAndCompound is a paid mutator transaction binding the contract method 0x2a4e051b.
//
// Solidity: function harvestAndCompound() returns()
func (_Tokendistributor *TokendistributorTransactorSession) HarvestAndCompound() (*types.Transaction, error) {
	return _Tokendistributor.Contract.HarvestAndCompound(&_Tokendistributor.TransactOpts)
}

// UpdatePool is a paid mutator transaction binding the contract method 0xe3161ddd.
//
// Solidity: function updatePool() returns()
func (_Tokendistributor *TokendistributorTransactor) UpdatePool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokendistributor.contract.Transact(opts, "updatePool")
}

// UpdatePool is a paid mutator transaction binding the contract method 0xe3161ddd.
//
// Solidity: function updatePool() returns()
func (_Tokendistributor *TokendistributorSession) UpdatePool() (*types.Transaction, error) {
	return _Tokendistributor.Contract.UpdatePool(&_Tokendistributor.TransactOpts)
}

// UpdatePool is a paid mutator transaction binding the contract method 0xe3161ddd.
//
// Solidity: function updatePool() returns()
func (_Tokendistributor *TokendistributorTransactorSession) UpdatePool() (*types.Transaction, error) {
	return _Tokendistributor.Contract.UpdatePool(&_Tokendistributor.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Tokendistributor *TokendistributorTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Tokendistributor.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Tokendistributor *TokendistributorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Tokendistributor.Contract.Withdraw(&_Tokendistributor.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Tokendistributor *TokendistributorTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Tokendistributor.Contract.Withdraw(&_Tokendistributor.TransactOpts, amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Tokendistributor *TokendistributorTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokendistributor.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Tokendistributor *TokendistributorSession) WithdrawAll() (*types.Transaction, error) {
	return _Tokendistributor.Contract.WithdrawAll(&_Tokendistributor.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Tokendistributor *TokendistributorTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _Tokendistributor.Contract.WithdrawAll(&_Tokendistributor.TransactOpts)
}

// TokendistributorCompoundIterator is returned from FilterCompound and is used to iterate over the raw logs and unpacked data for Compound events raised by the Tokendistributor contract.
type TokendistributorCompoundIterator struct {
	Event *TokendistributorCompound // Event containing the contract specifics and raw log

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
func (it *TokendistributorCompoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokendistributorCompound)
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
		it.Event = new(TokendistributorCompound)
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
func (it *TokendistributorCompoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokendistributorCompoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokendistributorCompound represents a Compound event raised by the Tokendistributor contract.
type TokendistributorCompound struct {
	User            common.Address
	HarvestedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCompound is a free log retrieval operation binding the contract event 0x169f1815ebdea059aac3bb00ec9a9594c7a5ffcb64a17e8392b5d84909a14556.
//
// Solidity: event Compound(address indexed user, uint256 harvestedAmount)
func (_Tokendistributor *TokendistributorFilterer) FilterCompound(opts *bind.FilterOpts, user []common.Address) (*TokendistributorCompoundIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tokendistributor.contract.FilterLogs(opts, "Compound", userRule)
	if err != nil {
		return nil, err
	}
	return &TokendistributorCompoundIterator{contract: _Tokendistributor.contract, event: "Compound", logs: logs, sub: sub}, nil
}

// WatchCompound is a free log subscription operation binding the contract event 0x169f1815ebdea059aac3bb00ec9a9594c7a5ffcb64a17e8392b5d84909a14556.
//
// Solidity: event Compound(address indexed user, uint256 harvestedAmount)
func (_Tokendistributor *TokendistributorFilterer) WatchCompound(opts *bind.WatchOpts, sink chan<- *TokendistributorCompound, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tokendistributor.contract.WatchLogs(opts, "Compound", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokendistributorCompound)
				if err := _Tokendistributor.contract.UnpackLog(event, "Compound", log); err != nil {
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

// ParseCompound is a log parse operation binding the contract event 0x169f1815ebdea059aac3bb00ec9a9594c7a5ffcb64a17e8392b5d84909a14556.
//
// Solidity: event Compound(address indexed user, uint256 harvestedAmount)
func (_Tokendistributor *TokendistributorFilterer) ParseCompound(log types.Log) (*TokendistributorCompound, error) {
	event := new(TokendistributorCompound)
	if err := _Tokendistributor.contract.UnpackLog(event, "Compound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokendistributorDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Tokendistributor contract.
type TokendistributorDepositIterator struct {
	Event *TokendistributorDeposit // Event containing the contract specifics and raw log

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
func (it *TokendistributorDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokendistributorDeposit)
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
		it.Event = new(TokendistributorDeposit)
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
func (it *TokendistributorDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokendistributorDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokendistributorDeposit represents a Deposit event raised by the Tokendistributor contract.
type TokendistributorDeposit struct {
	User            common.Address
	Amount          *big.Int
	HarvestedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 amount, uint256 harvestedAmount)
func (_Tokendistributor *TokendistributorFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*TokendistributorDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tokendistributor.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &TokendistributorDepositIterator{contract: _Tokendistributor.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 amount, uint256 harvestedAmount)
func (_Tokendistributor *TokendistributorFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *TokendistributorDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tokendistributor.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokendistributorDeposit)
				if err := _Tokendistributor.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 amount, uint256 harvestedAmount)
func (_Tokendistributor *TokendistributorFilterer) ParseDeposit(log types.Log) (*TokendistributorDeposit, error) {
	event := new(TokendistributorDeposit)
	if err := _Tokendistributor.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokendistributorNewRewardsPerBlockIterator is returned from FilterNewRewardsPerBlock and is used to iterate over the raw logs and unpacked data for NewRewardsPerBlock events raised by the Tokendistributor contract.
type TokendistributorNewRewardsPerBlockIterator struct {
	Event *TokendistributorNewRewardsPerBlock // Event containing the contract specifics and raw log

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
func (it *TokendistributorNewRewardsPerBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokendistributorNewRewardsPerBlock)
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
		it.Event = new(TokendistributorNewRewardsPerBlock)
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
func (it *TokendistributorNewRewardsPerBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokendistributorNewRewardsPerBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokendistributorNewRewardsPerBlock represents a NewRewardsPerBlock event raised by the Tokendistributor contract.
type TokendistributorNewRewardsPerBlock struct {
	CurrentPhase             *big.Int
	StartBlock               *big.Int
	RewardPerBlockForStaking *big.Int
	RewardPerBlockForOthers  *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewRewardsPerBlock is a free log retrieval operation binding the contract event 0x40181eb77bccfdef1a73b669bb4290d98e2fbec678c7cf4578ae256210420e17.
//
// Solidity: event NewRewardsPerBlock(uint256 indexed currentPhase, uint256 startBlock, uint256 rewardPerBlockForStaking, uint256 rewardPerBlockForOthers)
func (_Tokendistributor *TokendistributorFilterer) FilterNewRewardsPerBlock(opts *bind.FilterOpts, currentPhase []*big.Int) (*TokendistributorNewRewardsPerBlockIterator, error) {

	var currentPhaseRule []interface{}
	for _, currentPhaseItem := range currentPhase {
		currentPhaseRule = append(currentPhaseRule, currentPhaseItem)
	}

	logs, sub, err := _Tokendistributor.contract.FilterLogs(opts, "NewRewardsPerBlock", currentPhaseRule)
	if err != nil {
		return nil, err
	}
	return &TokendistributorNewRewardsPerBlockIterator{contract: _Tokendistributor.contract, event: "NewRewardsPerBlock", logs: logs, sub: sub}, nil
}

// WatchNewRewardsPerBlock is a free log subscription operation binding the contract event 0x40181eb77bccfdef1a73b669bb4290d98e2fbec678c7cf4578ae256210420e17.
//
// Solidity: event NewRewardsPerBlock(uint256 indexed currentPhase, uint256 startBlock, uint256 rewardPerBlockForStaking, uint256 rewardPerBlockForOthers)
func (_Tokendistributor *TokendistributorFilterer) WatchNewRewardsPerBlock(opts *bind.WatchOpts, sink chan<- *TokendistributorNewRewardsPerBlock, currentPhase []*big.Int) (event.Subscription, error) {

	var currentPhaseRule []interface{}
	for _, currentPhaseItem := range currentPhase {
		currentPhaseRule = append(currentPhaseRule, currentPhaseItem)
	}

	logs, sub, err := _Tokendistributor.contract.WatchLogs(opts, "NewRewardsPerBlock", currentPhaseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokendistributorNewRewardsPerBlock)
				if err := _Tokendistributor.contract.UnpackLog(event, "NewRewardsPerBlock", log); err != nil {
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

// ParseNewRewardsPerBlock is a log parse operation binding the contract event 0x40181eb77bccfdef1a73b669bb4290d98e2fbec678c7cf4578ae256210420e17.
//
// Solidity: event NewRewardsPerBlock(uint256 indexed currentPhase, uint256 startBlock, uint256 rewardPerBlockForStaking, uint256 rewardPerBlockForOthers)
func (_Tokendistributor *TokendistributorFilterer) ParseNewRewardsPerBlock(log types.Log) (*TokendistributorNewRewardsPerBlock, error) {
	event := new(TokendistributorNewRewardsPerBlock)
	if err := _Tokendistributor.contract.UnpackLog(event, "NewRewardsPerBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokendistributorWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Tokendistributor contract.
type TokendistributorWithdrawIterator struct {
	Event *TokendistributorWithdraw // Event containing the contract specifics and raw log

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
func (it *TokendistributorWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokendistributorWithdraw)
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
		it.Event = new(TokendistributorWithdraw)
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
func (it *TokendistributorWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokendistributorWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokendistributorWithdraw represents a Withdraw event raised by the Tokendistributor contract.
type TokendistributorWithdraw struct {
	User            common.Address
	Amount          *big.Int
	HarvestedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 amount, uint256 harvestedAmount)
func (_Tokendistributor *TokendistributorFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*TokendistributorWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tokendistributor.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &TokendistributorWithdrawIterator{contract: _Tokendistributor.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 amount, uint256 harvestedAmount)
func (_Tokendistributor *TokendistributorFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *TokendistributorWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tokendistributor.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokendistributorWithdraw)
				if err := _Tokendistributor.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 amount, uint256 harvestedAmount)
func (_Tokendistributor *TokendistributorFilterer) ParseWithdraw(log types.Log) (*TokendistributorWithdraw, error) {
	event := new(TokendistributorWithdraw)
	if err := _Tokendistributor.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
