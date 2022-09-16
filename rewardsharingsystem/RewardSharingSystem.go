// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rewardsharingsystem

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

// RewardsharingsystemMetaData contains all meta data concerning the Rewardsharingsystem contract.
var RewardsharingsystemMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hunterGovernanceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenDistributor\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestedAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestedAmount\",\"type\":\"uint256\"}],\"name\":\"HarvestFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestedAmount\",\"type\":\"uint256\"}],\"name\":\"HarvestHTG\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numberBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"NewRewardPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestedAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRewardPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimRewardToken\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestHTG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hunterGovernanceToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardAdjustment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"pendingFeeRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"pendingHTGRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodEndBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenDistributor\",\"outputs\":[{\"internalType\":\"contractTokenDistributor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDurationInBlocks\",\"type\":\"uint256\"}],\"name\":\"updateRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userRewardPerTokenPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimRewardToken\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"claimRewardToken\",\"type\":\"bool\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RewardsharingsystemABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardsharingsystemMetaData.ABI instead.
var RewardsharingsystemABI = RewardsharingsystemMetaData.ABI

// Rewardsharingsystem is an auto generated Go binding around an Ethereum contract.
type Rewardsharingsystem struct {
	RewardsharingsystemCaller     // Read-only binding to the contract
	RewardsharingsystemTransactor // Write-only binding to the contract
	RewardsharingsystemFilterer   // Log filterer for contract events
}

// RewardsharingsystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardsharingsystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsharingsystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardsharingsystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsharingsystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardsharingsystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsharingsystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardsharingsystemSession struct {
	Contract     *Rewardsharingsystem // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RewardsharingsystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardsharingsystemCallerSession struct {
	Contract *RewardsharingsystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RewardsharingsystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardsharingsystemTransactorSession struct {
	Contract     *RewardsharingsystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RewardsharingsystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardsharingsystemRaw struct {
	Contract *Rewardsharingsystem // Generic contract binding to access the raw methods on
}

// RewardsharingsystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardsharingsystemCallerRaw struct {
	Contract *RewardsharingsystemCaller // Generic read-only contract binding to access the raw methods on
}

// RewardsharingsystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardsharingsystemTransactorRaw struct {
	Contract *RewardsharingsystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardsharingsystem creates a new instance of Rewardsharingsystem, bound to a specific deployed contract.
func NewRewardsharingsystem(address common.Address, backend bind.ContractBackend) (*Rewardsharingsystem, error) {
	contract, err := bindRewardsharingsystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rewardsharingsystem{RewardsharingsystemCaller: RewardsharingsystemCaller{contract: contract}, RewardsharingsystemTransactor: RewardsharingsystemTransactor{contract: contract}, RewardsharingsystemFilterer: RewardsharingsystemFilterer{contract: contract}}, nil
}

// NewRewardsharingsystemCaller creates a new read-only instance of Rewardsharingsystem, bound to a specific deployed contract.
func NewRewardsharingsystemCaller(address common.Address, caller bind.ContractCaller) (*RewardsharingsystemCaller, error) {
	contract, err := bindRewardsharingsystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsharingsystemCaller{contract: contract}, nil
}

// NewRewardsharingsystemTransactor creates a new write-only instance of Rewardsharingsystem, bound to a specific deployed contract.
func NewRewardsharingsystemTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardsharingsystemTransactor, error) {
	contract, err := bindRewardsharingsystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsharingsystemTransactor{contract: contract}, nil
}

// NewRewardsharingsystemFilterer creates a new log filterer instance of Rewardsharingsystem, bound to a specific deployed contract.
func NewRewardsharingsystemFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardsharingsystemFilterer, error) {
	contract, err := bindRewardsharingsystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardsharingsystemFilterer{contract: contract}, nil
}

// bindRewardsharingsystem binds a generic wrapper to an already deployed contract.
func bindRewardsharingsystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardsharingsystemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewardsharingsystem *RewardsharingsystemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewardsharingsystem.Contract.RewardsharingsystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewardsharingsystem *RewardsharingsystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.RewardsharingsystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewardsharingsystem *RewardsharingsystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.RewardsharingsystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewardsharingsystem *RewardsharingsystemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewardsharingsystem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewardsharingsystem *RewardsharingsystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewardsharingsystem *RewardsharingsystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.contract.Transact(opts, method, params...)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCaller) PRECISIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewardsharingsystem.contract.Call(opts, &out, "PRECISION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemSession) PRECISIONFACTOR() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.PRECISIONFACTOR(&_Rewardsharingsystem.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCallerSession) PRECISIONFACTOR() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.PRECISIONFACTOR(&_Rewardsharingsystem.CallOpts)
}

// CurrentRewardPerBlock is a free data retrieval call binding the contract method 0xcb4aec61.
//
// Solidity: function currentRewardPerBlock() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCaller) CurrentRewardPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewardsharingsystem.contract.Call(opts, &out, "currentRewardPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRewardPerBlock is a free data retrieval call binding the contract method 0xcb4aec61.
//
// Solidity: function currentRewardPerBlock() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemSession) CurrentRewardPerBlock() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.CurrentRewardPerBlock(&_Rewardsharingsystem.CallOpts)
}

// CurrentRewardPerBlock is a free data retrieval call binding the contract method 0xcb4aec61.
//
// Solidity: function currentRewardPerBlock() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCallerSession) CurrentRewardPerBlock() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.CurrentRewardPerBlock(&_Rewardsharingsystem.CallOpts)
}

// HunterGovernanceToken is a free data retrieval call binding the contract method 0xf363264a.
//
// Solidity: function hunterGovernanceToken() view returns(address)
func (_Rewardsharingsystem *RewardsharingsystemCaller) HunterGovernanceToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardsharingsystem.contract.Call(opts, &out, "hunterGovernanceToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HunterGovernanceToken is a free data retrieval call binding the contract method 0xf363264a.
//
// Solidity: function hunterGovernanceToken() view returns(address)
func (_Rewardsharingsystem *RewardsharingsystemSession) HunterGovernanceToken() (common.Address, error) {
	return _Rewardsharingsystem.Contract.HunterGovernanceToken(&_Rewardsharingsystem.CallOpts)
}

// HunterGovernanceToken is a free data retrieval call binding the contract method 0xf363264a.
//
// Solidity: function hunterGovernanceToken() view returns(address)
func (_Rewardsharingsystem *RewardsharingsystemCallerSession) HunterGovernanceToken() (common.Address, error) {
	return _Rewardsharingsystem.Contract.HunterGovernanceToken(&_Rewardsharingsystem.CallOpts)
}

// LastRewardAdjustment is a free data retrieval call binding the contract method 0x40d2abae.
//
// Solidity: function lastRewardAdjustment() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCaller) LastRewardAdjustment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewardsharingsystem.contract.Call(opts, &out, "lastRewardAdjustment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardAdjustment is a free data retrieval call binding the contract method 0x40d2abae.
//
// Solidity: function lastRewardAdjustment() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemSession) LastRewardAdjustment() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.LastRewardAdjustment(&_Rewardsharingsystem.CallOpts)
}

// LastRewardAdjustment is a free data retrieval call binding the contract method 0x40d2abae.
//
// Solidity: function lastRewardAdjustment() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCallerSession) LastRewardAdjustment() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.LastRewardAdjustment(&_Rewardsharingsystem.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCaller) LastRewardBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewardsharingsystem.contract.Call(opts, &out, "lastRewardBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemSession) LastRewardBlock() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.LastRewardBlock(&_Rewardsharingsystem.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCallerSession) LastRewardBlock() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.LastRewardBlock(&_Rewardsharingsystem.CallOpts)
}

// LastUpdateBlock is a free data retrieval call binding the contract method 0xa218141b.
//
// Solidity: function lastUpdateBlock() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCaller) LastUpdateBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewardsharingsystem.contract.Call(opts, &out, "lastUpdateBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateBlock is a free data retrieval call binding the contract method 0xa218141b.
//
// Solidity: function lastUpdateBlock() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemSession) LastUpdateBlock() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.LastUpdateBlock(&_Rewardsharingsystem.CallOpts)
}

// LastUpdateBlock is a free data retrieval call binding the contract method 0xa218141b.
//
// Solidity: function lastUpdateBlock() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCallerSession) LastUpdateBlock() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.LastUpdateBlock(&_Rewardsharingsystem.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewardsharingsystem *RewardsharingsystemCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardsharingsystem.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewardsharingsystem *RewardsharingsystemSession) Owner() (common.Address, error) {
	return _Rewardsharingsystem.Contract.Owner(&_Rewardsharingsystem.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewardsharingsystem *RewardsharingsystemCallerSession) Owner() (common.Address, error) {
	return _Rewardsharingsystem.Contract.Owner(&_Rewardsharingsystem.CallOpts)
}

// PendingFeeRewards is a free data retrieval call binding the contract method 0xd86e29a8.
//
// Solidity: function pendingFeeRewards(address user) view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCaller) PendingFeeRewards(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rewardsharingsystem.contract.Call(opts, &out, "pendingFeeRewards", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingFeeRewards is a free data retrieval call binding the contract method 0xd86e29a8.
//
// Solidity: function pendingFeeRewards(address user) view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemSession) PendingFeeRewards(user common.Address) (*big.Int, error) {
	return _Rewardsharingsystem.Contract.PendingFeeRewards(&_Rewardsharingsystem.CallOpts, user)
}

// PendingFeeRewards is a free data retrieval call binding the contract method 0xd86e29a8.
//
// Solidity: function pendingFeeRewards(address user) view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCallerSession) PendingFeeRewards(user common.Address) (*big.Int, error) {
	return _Rewardsharingsystem.Contract.PendingFeeRewards(&_Rewardsharingsystem.CallOpts, user)
}

// PendingHTGRewards is a free data retrieval call binding the contract method 0x66888246.
//
// Solidity: function pendingHTGRewards(address user) view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCaller) PendingHTGRewards(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rewardsharingsystem.contract.Call(opts, &out, "pendingHTGRewards", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingHTGRewards is a free data retrieval call binding the contract method 0x66888246.
//
// Solidity: function pendingHTGRewards(address user) view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemSession) PendingHTGRewards(user common.Address) (*big.Int, error) {
	return _Rewardsharingsystem.Contract.PendingHTGRewards(&_Rewardsharingsystem.CallOpts, user)
}

// PendingHTGRewards is a free data retrieval call binding the contract method 0x66888246.
//
// Solidity: function pendingHTGRewards(address user) view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCallerSession) PendingHTGRewards(user common.Address) (*big.Int, error) {
	return _Rewardsharingsystem.Contract.PendingHTGRewards(&_Rewardsharingsystem.CallOpts, user)
}

// PeriodEndBlock is a free data retrieval call binding the contract method 0x442da82f.
//
// Solidity: function periodEndBlock() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCaller) PeriodEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewardsharingsystem.contract.Call(opts, &out, "periodEndBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodEndBlock is a free data retrieval call binding the contract method 0x442da82f.
//
// Solidity: function periodEndBlock() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemSession) PeriodEndBlock() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.PeriodEndBlock(&_Rewardsharingsystem.CallOpts)
}

// PeriodEndBlock is a free data retrieval call binding the contract method 0x442da82f.
//
// Solidity: function periodEndBlock() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCallerSession) PeriodEndBlock() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.PeriodEndBlock(&_Rewardsharingsystem.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewardsharingsystem.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemSession) RewardPerTokenStored() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.RewardPerTokenStored(&_Rewardsharingsystem.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.RewardPerTokenStored(&_Rewardsharingsystem.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Rewardsharingsystem *RewardsharingsystemCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardsharingsystem.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Rewardsharingsystem *RewardsharingsystemSession) RewardToken() (common.Address, error) {
	return _Rewardsharingsystem.Contract.RewardToken(&_Rewardsharingsystem.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Rewardsharingsystem *RewardsharingsystemCallerSession) RewardToken() (common.Address, error) {
	return _Rewardsharingsystem.Contract.RewardToken(&_Rewardsharingsystem.CallOpts)
}

// TokenDistributor is a free data retrieval call binding the contract method 0x18a6bc32.
//
// Solidity: function tokenDistributor() view returns(address)
func (_Rewardsharingsystem *RewardsharingsystemCaller) TokenDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewardsharingsystem.contract.Call(opts, &out, "tokenDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenDistributor is a free data retrieval call binding the contract method 0x18a6bc32.
//
// Solidity: function tokenDistributor() view returns(address)
func (_Rewardsharingsystem *RewardsharingsystemSession) TokenDistributor() (common.Address, error) {
	return _Rewardsharingsystem.Contract.TokenDistributor(&_Rewardsharingsystem.CallOpts)
}

// TokenDistributor is a free data retrieval call binding the contract method 0x18a6bc32.
//
// Solidity: function tokenDistributor() view returns(address)
func (_Rewardsharingsystem *RewardsharingsystemCallerSession) TokenDistributor() (common.Address, error) {
	return _Rewardsharingsystem.Contract.TokenDistributor(&_Rewardsharingsystem.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewardsharingsystem.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemSession) TotalShares() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.TotalShares(&_Rewardsharingsystem.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Rewardsharingsystem *RewardsharingsystemCallerSession) TotalShares() (*big.Int, error) {
	return _Rewardsharingsystem.Contract.TotalShares(&_Rewardsharingsystem.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 shares, uint256 userRewardPerTokenPaid, uint256 rewards)
func (_Rewardsharingsystem *RewardsharingsystemCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount                 *big.Int
	Shares                 *big.Int
	UserRewardPerTokenPaid *big.Int
	Rewards                *big.Int
}, error) {
	var out []interface{}
	err := _Rewardsharingsystem.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		Amount                 *big.Int
		Shares                 *big.Int
		UserRewardPerTokenPaid *big.Int
		Rewards                *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Shares = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UserRewardPerTokenPaid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Rewards = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 shares, uint256 userRewardPerTokenPaid, uint256 rewards)
func (_Rewardsharingsystem *RewardsharingsystemSession) UserInfo(arg0 common.Address) (struct {
	Amount                 *big.Int
	Shares                 *big.Int
	UserRewardPerTokenPaid *big.Int
	Rewards                *big.Int
}, error) {
	return _Rewardsharingsystem.Contract.UserInfo(&_Rewardsharingsystem.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 shares, uint256 userRewardPerTokenPaid, uint256 rewards)
func (_Rewardsharingsystem *RewardsharingsystemCallerSession) UserInfo(arg0 common.Address) (struct {
	Amount                 *big.Int
	Shares                 *big.Int
	UserRewardPerTokenPaid *big.Int
	Rewards                *big.Int
}, error) {
	return _Rewardsharingsystem.Contract.UserInfo(&_Rewardsharingsystem.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0x9a408321.
//
// Solidity: function deposit(uint256 amount, bool claimRewardToken) returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, claimRewardToken bool) (*types.Transaction, error) {
	return _Rewardsharingsystem.contract.Transact(opts, "deposit", amount, claimRewardToken)
}

// Deposit is a paid mutator transaction binding the contract method 0x9a408321.
//
// Solidity: function deposit(uint256 amount, bool claimRewardToken) returns()
func (_Rewardsharingsystem *RewardsharingsystemSession) Deposit(amount *big.Int, claimRewardToken bool) (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.Deposit(&_Rewardsharingsystem.TransactOpts, amount, claimRewardToken)
}

// Deposit is a paid mutator transaction binding the contract method 0x9a408321.
//
// Solidity: function deposit(uint256 amount, bool claimRewardToken) returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactorSession) Deposit(amount *big.Int, claimRewardToken bool) (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.Deposit(&_Rewardsharingsystem.TransactOpts, amount, claimRewardToken)
}

// HarvestFee is a paid mutator transaction binding the contract method 0xb5962917.
//
// Solidity: function harvestFee() returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactor) HarvestFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewardsharingsystem.contract.Transact(opts, "harvestFee")
}

// HarvestFee is a paid mutator transaction binding the contract method 0xb5962917.
//
// Solidity: function harvestFee() returns()
func (_Rewardsharingsystem *RewardsharingsystemSession) HarvestFee() (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.HarvestFee(&_Rewardsharingsystem.TransactOpts)
}

// HarvestFee is a paid mutator transaction binding the contract method 0xb5962917.
//
// Solidity: function harvestFee() returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactorSession) HarvestFee() (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.HarvestFee(&_Rewardsharingsystem.TransactOpts)
}

// HarvestHTG is a paid mutator transaction binding the contract method 0xf3ab32c2.
//
// Solidity: function harvestHTG() returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactor) HarvestHTG(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewardsharingsystem.contract.Transact(opts, "harvestHTG")
}

// HarvestHTG is a paid mutator transaction binding the contract method 0xf3ab32c2.
//
// Solidity: function harvestHTG() returns()
func (_Rewardsharingsystem *RewardsharingsystemSession) HarvestHTG() (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.HarvestHTG(&_Rewardsharingsystem.TransactOpts)
}

// HarvestHTG is a paid mutator transaction binding the contract method 0xf3ab32c2.
//
// Solidity: function harvestHTG() returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactorSession) HarvestHTG() (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.HarvestHTG(&_Rewardsharingsystem.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewardsharingsystem.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewardsharingsystem *RewardsharingsystemSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.RenounceOwnership(&_Rewardsharingsystem.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.RenounceOwnership(&_Rewardsharingsystem.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Rewardsharingsystem.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewardsharingsystem *RewardsharingsystemSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.TransferOwnership(&_Rewardsharingsystem.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.TransferOwnership(&_Rewardsharingsystem.TransactOpts, newOwner)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x97e50818.
//
// Solidity: function updateRewards(uint256 reward, uint256 rewardDurationInBlocks) returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactor) UpdateRewards(opts *bind.TransactOpts, reward *big.Int, rewardDurationInBlocks *big.Int) (*types.Transaction, error) {
	return _Rewardsharingsystem.contract.Transact(opts, "updateRewards", reward, rewardDurationInBlocks)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x97e50818.
//
// Solidity: function updateRewards(uint256 reward, uint256 rewardDurationInBlocks) returns()
func (_Rewardsharingsystem *RewardsharingsystemSession) UpdateRewards(reward *big.Int, rewardDurationInBlocks *big.Int) (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.UpdateRewards(&_Rewardsharingsystem.TransactOpts, reward, rewardDurationInBlocks)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x97e50818.
//
// Solidity: function updateRewards(uint256 reward, uint256 rewardDurationInBlocks) returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactorSession) UpdateRewards(reward *big.Int, rewardDurationInBlocks *big.Int) (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.UpdateRewards(&_Rewardsharingsystem.TransactOpts, reward, rewardDurationInBlocks)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool claimRewardToken) returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, claimRewardToken bool) (*types.Transaction, error) {
	return _Rewardsharingsystem.contract.Transact(opts, "withdraw", amount, claimRewardToken)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool claimRewardToken) returns()
func (_Rewardsharingsystem *RewardsharingsystemSession) Withdraw(amount *big.Int, claimRewardToken bool) (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.Withdraw(&_Rewardsharingsystem.TransactOpts, amount, claimRewardToken)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool claimRewardToken) returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactorSession) Withdraw(amount *big.Int, claimRewardToken bool) (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.Withdraw(&_Rewardsharingsystem.TransactOpts, amount, claimRewardToken)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x1c1c6fe5.
//
// Solidity: function withdrawAll(bool claimRewardToken) returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactor) WithdrawAll(opts *bind.TransactOpts, claimRewardToken bool) (*types.Transaction, error) {
	return _Rewardsharingsystem.contract.Transact(opts, "withdrawAll", claimRewardToken)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x1c1c6fe5.
//
// Solidity: function withdrawAll(bool claimRewardToken) returns()
func (_Rewardsharingsystem *RewardsharingsystemSession) WithdrawAll(claimRewardToken bool) (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.WithdrawAll(&_Rewardsharingsystem.TransactOpts, claimRewardToken)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x1c1c6fe5.
//
// Solidity: function withdrawAll(bool claimRewardToken) returns()
func (_Rewardsharingsystem *RewardsharingsystemTransactorSession) WithdrawAll(claimRewardToken bool) (*types.Transaction, error) {
	return _Rewardsharingsystem.Contract.WithdrawAll(&_Rewardsharingsystem.TransactOpts, claimRewardToken)
}

// RewardsharingsystemDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Rewardsharingsystem contract.
type RewardsharingsystemDepositIterator struct {
	Event *RewardsharingsystemDeposit // Event containing the contract specifics and raw log

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
func (it *RewardsharingsystemDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsharingsystemDeposit)
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
		it.Event = new(RewardsharingsystemDeposit)
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
func (it *RewardsharingsystemDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsharingsystemDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsharingsystemDeposit represents a Deposit event raised by the Rewardsharingsystem contract.
type RewardsharingsystemDeposit struct {
	User            common.Address
	Amount          *big.Int
	HarvestedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 amount, uint256 harvestedAmount)
func (_Rewardsharingsystem *RewardsharingsystemFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*RewardsharingsystemDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewardsharingsystem.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &RewardsharingsystemDepositIterator{contract: _Rewardsharingsystem.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 amount, uint256 harvestedAmount)
func (_Rewardsharingsystem *RewardsharingsystemFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *RewardsharingsystemDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewardsharingsystem.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsharingsystemDeposit)
				if err := _Rewardsharingsystem.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Rewardsharingsystem *RewardsharingsystemFilterer) ParseDeposit(log types.Log) (*RewardsharingsystemDeposit, error) {
	event := new(RewardsharingsystemDeposit)
	if err := _Rewardsharingsystem.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsharingsystemHarvestFeeIterator is returned from FilterHarvestFee and is used to iterate over the raw logs and unpacked data for HarvestFee events raised by the Rewardsharingsystem contract.
type RewardsharingsystemHarvestFeeIterator struct {
	Event *RewardsharingsystemHarvestFee // Event containing the contract specifics and raw log

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
func (it *RewardsharingsystemHarvestFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsharingsystemHarvestFee)
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
		it.Event = new(RewardsharingsystemHarvestFee)
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
func (it *RewardsharingsystemHarvestFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsharingsystemHarvestFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsharingsystemHarvestFee represents a HarvestFee event raised by the Rewardsharingsystem contract.
type RewardsharingsystemHarvestFee struct {
	User            common.Address
	HarvestedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterHarvestFee is a free log retrieval operation binding the contract event 0xa0e0c5921122a66e1e5a61d9a1469a42fa3ba94092de828dc1668de8ffbef01a.
//
// Solidity: event HarvestFee(address indexed user, uint256 harvestedAmount)
func (_Rewardsharingsystem *RewardsharingsystemFilterer) FilterHarvestFee(opts *bind.FilterOpts, user []common.Address) (*RewardsharingsystemHarvestFeeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewardsharingsystem.contract.FilterLogs(opts, "HarvestFee", userRule)
	if err != nil {
		return nil, err
	}
	return &RewardsharingsystemHarvestFeeIterator{contract: _Rewardsharingsystem.contract, event: "HarvestFee", logs: logs, sub: sub}, nil
}

// WatchHarvestFee is a free log subscription operation binding the contract event 0xa0e0c5921122a66e1e5a61d9a1469a42fa3ba94092de828dc1668de8ffbef01a.
//
// Solidity: event HarvestFee(address indexed user, uint256 harvestedAmount)
func (_Rewardsharingsystem *RewardsharingsystemFilterer) WatchHarvestFee(opts *bind.WatchOpts, sink chan<- *RewardsharingsystemHarvestFee, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewardsharingsystem.contract.WatchLogs(opts, "HarvestFee", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsharingsystemHarvestFee)
				if err := _Rewardsharingsystem.contract.UnpackLog(event, "HarvestFee", log); err != nil {
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

// ParseHarvestFee is a log parse operation binding the contract event 0xa0e0c5921122a66e1e5a61d9a1469a42fa3ba94092de828dc1668de8ffbef01a.
//
// Solidity: event HarvestFee(address indexed user, uint256 harvestedAmount)
func (_Rewardsharingsystem *RewardsharingsystemFilterer) ParseHarvestFee(log types.Log) (*RewardsharingsystemHarvestFee, error) {
	event := new(RewardsharingsystemHarvestFee)
	if err := _Rewardsharingsystem.contract.UnpackLog(event, "HarvestFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsharingsystemHarvestHTGIterator is returned from FilterHarvestHTG and is used to iterate over the raw logs and unpacked data for HarvestHTG events raised by the Rewardsharingsystem contract.
type RewardsharingsystemHarvestHTGIterator struct {
	Event *RewardsharingsystemHarvestHTG // Event containing the contract specifics and raw log

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
func (it *RewardsharingsystemHarvestHTGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsharingsystemHarvestHTG)
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
		it.Event = new(RewardsharingsystemHarvestHTG)
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
func (it *RewardsharingsystemHarvestHTGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsharingsystemHarvestHTGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsharingsystemHarvestHTG represents a HarvestHTG event raised by the Rewardsharingsystem contract.
type RewardsharingsystemHarvestHTG struct {
	User            common.Address
	HarvestedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterHarvestHTG is a free log retrieval operation binding the contract event 0x17012d50b70af51869b4b6a658aa39095877d071272178dce41c0cb3722e7e05.
//
// Solidity: event HarvestHTG(address indexed user, uint256 harvestedAmount)
func (_Rewardsharingsystem *RewardsharingsystemFilterer) FilterHarvestHTG(opts *bind.FilterOpts, user []common.Address) (*RewardsharingsystemHarvestHTGIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewardsharingsystem.contract.FilterLogs(opts, "HarvestHTG", userRule)
	if err != nil {
		return nil, err
	}
	return &RewardsharingsystemHarvestHTGIterator{contract: _Rewardsharingsystem.contract, event: "HarvestHTG", logs: logs, sub: sub}, nil
}

// WatchHarvestHTG is a free log subscription operation binding the contract event 0x17012d50b70af51869b4b6a658aa39095877d071272178dce41c0cb3722e7e05.
//
// Solidity: event HarvestHTG(address indexed user, uint256 harvestedAmount)
func (_Rewardsharingsystem *RewardsharingsystemFilterer) WatchHarvestHTG(opts *bind.WatchOpts, sink chan<- *RewardsharingsystemHarvestHTG, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewardsharingsystem.contract.WatchLogs(opts, "HarvestHTG", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsharingsystemHarvestHTG)
				if err := _Rewardsharingsystem.contract.UnpackLog(event, "HarvestHTG", log); err != nil {
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

// ParseHarvestHTG is a log parse operation binding the contract event 0x17012d50b70af51869b4b6a658aa39095877d071272178dce41c0cb3722e7e05.
//
// Solidity: event HarvestHTG(address indexed user, uint256 harvestedAmount)
func (_Rewardsharingsystem *RewardsharingsystemFilterer) ParseHarvestHTG(log types.Log) (*RewardsharingsystemHarvestHTG, error) {
	event := new(RewardsharingsystemHarvestHTG)
	if err := _Rewardsharingsystem.contract.UnpackLog(event, "HarvestHTG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsharingsystemNewRewardPeriodIterator is returned from FilterNewRewardPeriod and is used to iterate over the raw logs and unpacked data for NewRewardPeriod events raised by the Rewardsharingsystem contract.
type RewardsharingsystemNewRewardPeriodIterator struct {
	Event *RewardsharingsystemNewRewardPeriod // Event containing the contract specifics and raw log

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
func (it *RewardsharingsystemNewRewardPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsharingsystemNewRewardPeriod)
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
		it.Event = new(RewardsharingsystemNewRewardPeriod)
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
func (it *RewardsharingsystemNewRewardPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsharingsystemNewRewardPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsharingsystemNewRewardPeriod represents a NewRewardPeriod event raised by the Rewardsharingsystem contract.
type RewardsharingsystemNewRewardPeriod struct {
	NumberBlocks   *big.Int
	RewardPerBlock *big.Int
	Reward         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewRewardPeriod is a free log retrieval operation binding the contract event 0x55b4fa63fe43865f67b4f2c4a4df1cf9e6c1f85767211b44b45cf4649b2c2b51.
//
// Solidity: event NewRewardPeriod(uint256 numberBlocks, uint256 rewardPerBlock, uint256 reward)
func (_Rewardsharingsystem *RewardsharingsystemFilterer) FilterNewRewardPeriod(opts *bind.FilterOpts) (*RewardsharingsystemNewRewardPeriodIterator, error) {

	logs, sub, err := _Rewardsharingsystem.contract.FilterLogs(opts, "NewRewardPeriod")
	if err != nil {
		return nil, err
	}
	return &RewardsharingsystemNewRewardPeriodIterator{contract: _Rewardsharingsystem.contract, event: "NewRewardPeriod", logs: logs, sub: sub}, nil
}

// WatchNewRewardPeriod is a free log subscription operation binding the contract event 0x55b4fa63fe43865f67b4f2c4a4df1cf9e6c1f85767211b44b45cf4649b2c2b51.
//
// Solidity: event NewRewardPeriod(uint256 numberBlocks, uint256 rewardPerBlock, uint256 reward)
func (_Rewardsharingsystem *RewardsharingsystemFilterer) WatchNewRewardPeriod(opts *bind.WatchOpts, sink chan<- *RewardsharingsystemNewRewardPeriod) (event.Subscription, error) {

	logs, sub, err := _Rewardsharingsystem.contract.WatchLogs(opts, "NewRewardPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsharingsystemNewRewardPeriod)
				if err := _Rewardsharingsystem.contract.UnpackLog(event, "NewRewardPeriod", log); err != nil {
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

// ParseNewRewardPeriod is a log parse operation binding the contract event 0x55b4fa63fe43865f67b4f2c4a4df1cf9e6c1f85767211b44b45cf4649b2c2b51.
//
// Solidity: event NewRewardPeriod(uint256 numberBlocks, uint256 rewardPerBlock, uint256 reward)
func (_Rewardsharingsystem *RewardsharingsystemFilterer) ParseNewRewardPeriod(log types.Log) (*RewardsharingsystemNewRewardPeriod, error) {
	event := new(RewardsharingsystemNewRewardPeriod)
	if err := _Rewardsharingsystem.contract.UnpackLog(event, "NewRewardPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsharingsystemOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Rewardsharingsystem contract.
type RewardsharingsystemOwnershipTransferredIterator struct {
	Event *RewardsharingsystemOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RewardsharingsystemOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsharingsystemOwnershipTransferred)
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
		it.Event = new(RewardsharingsystemOwnershipTransferred)
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
func (it *RewardsharingsystemOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsharingsystemOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsharingsystemOwnershipTransferred represents a OwnershipTransferred event raised by the Rewardsharingsystem contract.
type RewardsharingsystemOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewardsharingsystem *RewardsharingsystemFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RewardsharingsystemOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rewardsharingsystem.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RewardsharingsystemOwnershipTransferredIterator{contract: _Rewardsharingsystem.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewardsharingsystem *RewardsharingsystemFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RewardsharingsystemOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rewardsharingsystem.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsharingsystemOwnershipTransferred)
				if err := _Rewardsharingsystem.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Rewardsharingsystem *RewardsharingsystemFilterer) ParseOwnershipTransferred(log types.Log) (*RewardsharingsystemOwnershipTransferred, error) {
	event := new(RewardsharingsystemOwnershipTransferred)
	if err := _Rewardsharingsystem.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsharingsystemWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Rewardsharingsystem contract.
type RewardsharingsystemWithdrawIterator struct {
	Event *RewardsharingsystemWithdraw // Event containing the contract specifics and raw log

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
func (it *RewardsharingsystemWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsharingsystemWithdraw)
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
		it.Event = new(RewardsharingsystemWithdraw)
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
func (it *RewardsharingsystemWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsharingsystemWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsharingsystemWithdraw represents a Withdraw event raised by the Rewardsharingsystem contract.
type RewardsharingsystemWithdraw struct {
	User            common.Address
	Amount          *big.Int
	HarvestedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 amount, uint256 harvestedAmount)
func (_Rewardsharingsystem *RewardsharingsystemFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*RewardsharingsystemWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewardsharingsystem.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &RewardsharingsystemWithdrawIterator{contract: _Rewardsharingsystem.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 amount, uint256 harvestedAmount)
func (_Rewardsharingsystem *RewardsharingsystemFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *RewardsharingsystemWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewardsharingsystem.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsharingsystemWithdraw)
				if err := _Rewardsharingsystem.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Rewardsharingsystem *RewardsharingsystemFilterer) ParseWithdraw(log types.Log) (*RewardsharingsystemWithdraw, error) {
	event := new(RewardsharingsystemWithdraw)
	if err := _Rewardsharingsystem.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
