// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feesharingsystem

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

// FeesharingsystemMetaData contains all meta data concerning the Feesharingsystem contract.
var FeesharingsystemMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_looksRareToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenDistributor\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestedAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestedAmount\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numberBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"NewRewardPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestedAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"calculatePendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateSharePriceInLOOKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"calculateSharesValueInLOOKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRewardPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimRewardToken\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hunterGovernanceToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardAdjustment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodEndBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenDistributor\",\"outputs\":[{\"internalType\":\"contractTokenDistributor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDurationInBlocks\",\"type\":\"uint256\"}],\"name\":\"updateRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userRewardPerTokenPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimRewardToken\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"claimRewardToken\",\"type\":\"bool\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FeesharingsystemABI is the input ABI used to generate the binding from.
// Deprecated: Use FeesharingsystemMetaData.ABI instead.
var FeesharingsystemABI = FeesharingsystemMetaData.ABI

// Feesharingsystem is an auto generated Go binding around an Ethereum contract.
type Feesharingsystem struct {
	FeesharingsystemCaller     // Read-only binding to the contract
	FeesharingsystemTransactor // Write-only binding to the contract
	FeesharingsystemFilterer   // Log filterer for contract events
}

// FeesharingsystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeesharingsystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeesharingsystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeesharingsystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeesharingsystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeesharingsystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeesharingsystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeesharingsystemSession struct {
	Contract     *Feesharingsystem // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeesharingsystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeesharingsystemCallerSession struct {
	Contract *FeesharingsystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FeesharingsystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeesharingsystemTransactorSession struct {
	Contract     *FeesharingsystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FeesharingsystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeesharingsystemRaw struct {
	Contract *Feesharingsystem // Generic contract binding to access the raw methods on
}

// FeesharingsystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeesharingsystemCallerRaw struct {
	Contract *FeesharingsystemCaller // Generic read-only contract binding to access the raw methods on
}

// FeesharingsystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeesharingsystemTransactorRaw struct {
	Contract *FeesharingsystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeesharingsystem creates a new instance of Feesharingsystem, bound to a specific deployed contract.
func NewFeesharingsystem(address common.Address, backend bind.ContractBackend) (*Feesharingsystem, error) {
	contract, err := bindFeesharingsystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Feesharingsystem{FeesharingsystemCaller: FeesharingsystemCaller{contract: contract}, FeesharingsystemTransactor: FeesharingsystemTransactor{contract: contract}, FeesharingsystemFilterer: FeesharingsystemFilterer{contract: contract}}, nil
}

// NewFeesharingsystemCaller creates a new read-only instance of Feesharingsystem, bound to a specific deployed contract.
func NewFeesharingsystemCaller(address common.Address, caller bind.ContractCaller) (*FeesharingsystemCaller, error) {
	contract, err := bindFeesharingsystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeesharingsystemCaller{contract: contract}, nil
}

// NewFeesharingsystemTransactor creates a new write-only instance of Feesharingsystem, bound to a specific deployed contract.
func NewFeesharingsystemTransactor(address common.Address, transactor bind.ContractTransactor) (*FeesharingsystemTransactor, error) {
	contract, err := bindFeesharingsystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeesharingsystemTransactor{contract: contract}, nil
}

// NewFeesharingsystemFilterer creates a new log filterer instance of Feesharingsystem, bound to a specific deployed contract.
func NewFeesharingsystemFilterer(address common.Address, filterer bind.ContractFilterer) (*FeesharingsystemFilterer, error) {
	contract, err := bindFeesharingsystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeesharingsystemFilterer{contract: contract}, nil
}

// bindFeesharingsystem binds a generic wrapper to an already deployed contract.
func bindFeesharingsystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeesharingsystemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Feesharingsystem *FeesharingsystemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Feesharingsystem.Contract.FeesharingsystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Feesharingsystem *FeesharingsystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feesharingsystem.Contract.FeesharingsystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Feesharingsystem *FeesharingsystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Feesharingsystem.Contract.FeesharingsystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Feesharingsystem *FeesharingsystemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Feesharingsystem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Feesharingsystem *FeesharingsystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feesharingsystem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Feesharingsystem *FeesharingsystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Feesharingsystem.Contract.contract.Transact(opts, method, params...)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCaller) PRECISIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "PRECISION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemSession) PRECISIONFACTOR() (*big.Int, error) {
	return _Feesharingsystem.Contract.PRECISIONFACTOR(&_Feesharingsystem.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCallerSession) PRECISIONFACTOR() (*big.Int, error) {
	return _Feesharingsystem.Contract.PRECISIONFACTOR(&_Feesharingsystem.CallOpts)
}

// CalculatePendingRewards is a free data retrieval call binding the contract method 0x097aad10.
//
// Solidity: function calculatePendingRewards(address user) view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCaller) CalculatePendingRewards(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "calculatePendingRewards", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePendingRewards is a free data retrieval call binding the contract method 0x097aad10.
//
// Solidity: function calculatePendingRewards(address user) view returns(uint256)
func (_Feesharingsystem *FeesharingsystemSession) CalculatePendingRewards(user common.Address) (*big.Int, error) {
	return _Feesharingsystem.Contract.CalculatePendingRewards(&_Feesharingsystem.CallOpts, user)
}

// CalculatePendingRewards is a free data retrieval call binding the contract method 0x097aad10.
//
// Solidity: function calculatePendingRewards(address user) view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCallerSession) CalculatePendingRewards(user common.Address) (*big.Int, error) {
	return _Feesharingsystem.Contract.CalculatePendingRewards(&_Feesharingsystem.CallOpts, user)
}

// CalculateSharePriceInLOOKS is a free data retrieval call binding the contract method 0x6de26e38.
//
// Solidity: function calculateSharePriceInLOOKS() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCaller) CalculateSharePriceInLOOKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "calculateSharePriceInLOOKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSharePriceInLOOKS is a free data retrieval call binding the contract method 0x6de26e38.
//
// Solidity: function calculateSharePriceInLOOKS() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemSession) CalculateSharePriceInLOOKS() (*big.Int, error) {
	return _Feesharingsystem.Contract.CalculateSharePriceInLOOKS(&_Feesharingsystem.CallOpts)
}

// CalculateSharePriceInLOOKS is a free data retrieval call binding the contract method 0x6de26e38.
//
// Solidity: function calculateSharePriceInLOOKS() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCallerSession) CalculateSharePriceInLOOKS() (*big.Int, error) {
	return _Feesharingsystem.Contract.CalculateSharePriceInLOOKS(&_Feesharingsystem.CallOpts)
}

// CalculateSharesValueInLOOKS is a free data retrieval call binding the contract method 0xab5e32af.
//
// Solidity: function calculateSharesValueInLOOKS(address user) view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCaller) CalculateSharesValueInLOOKS(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "calculateSharesValueInLOOKS", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSharesValueInLOOKS is a free data retrieval call binding the contract method 0xab5e32af.
//
// Solidity: function calculateSharesValueInLOOKS(address user) view returns(uint256)
func (_Feesharingsystem *FeesharingsystemSession) CalculateSharesValueInLOOKS(user common.Address) (*big.Int, error) {
	return _Feesharingsystem.Contract.CalculateSharesValueInLOOKS(&_Feesharingsystem.CallOpts, user)
}

// CalculateSharesValueInLOOKS is a free data retrieval call binding the contract method 0xab5e32af.
//
// Solidity: function calculateSharesValueInLOOKS(address user) view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCallerSession) CalculateSharesValueInLOOKS(user common.Address) (*big.Int, error) {
	return _Feesharingsystem.Contract.CalculateSharesValueInLOOKS(&_Feesharingsystem.CallOpts, user)
}

// CurrentRewardPerBlock is a free data retrieval call binding the contract method 0xcb4aec61.
//
// Solidity: function currentRewardPerBlock() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCaller) CurrentRewardPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "currentRewardPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRewardPerBlock is a free data retrieval call binding the contract method 0xcb4aec61.
//
// Solidity: function currentRewardPerBlock() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemSession) CurrentRewardPerBlock() (*big.Int, error) {
	return _Feesharingsystem.Contract.CurrentRewardPerBlock(&_Feesharingsystem.CallOpts)
}

// CurrentRewardPerBlock is a free data retrieval call binding the contract method 0xcb4aec61.
//
// Solidity: function currentRewardPerBlock() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCallerSession) CurrentRewardPerBlock() (*big.Int, error) {
	return _Feesharingsystem.Contract.CurrentRewardPerBlock(&_Feesharingsystem.CallOpts)
}

// HunterGovernanceToken is a free data retrieval call binding the contract method 0xf363264a.
//
// Solidity: function hunterGovernanceToken() view returns(address)
func (_Feesharingsystem *FeesharingsystemCaller) HunterGovernanceToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "hunterGovernanceToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HunterGovernanceToken is a free data retrieval call binding the contract method 0xf363264a.
//
// Solidity: function hunterGovernanceToken() view returns(address)
func (_Feesharingsystem *FeesharingsystemSession) HunterGovernanceToken() (common.Address, error) {
	return _Feesharingsystem.Contract.HunterGovernanceToken(&_Feesharingsystem.CallOpts)
}

// HunterGovernanceToken is a free data retrieval call binding the contract method 0xf363264a.
//
// Solidity: function hunterGovernanceToken() view returns(address)
func (_Feesharingsystem *FeesharingsystemCallerSession) HunterGovernanceToken() (common.Address, error) {
	return _Feesharingsystem.Contract.HunterGovernanceToken(&_Feesharingsystem.CallOpts)
}

// LastRewardAdjustment is a free data retrieval call binding the contract method 0x40d2abae.
//
// Solidity: function lastRewardAdjustment() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCaller) LastRewardAdjustment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "lastRewardAdjustment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardAdjustment is a free data retrieval call binding the contract method 0x40d2abae.
//
// Solidity: function lastRewardAdjustment() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemSession) LastRewardAdjustment() (*big.Int, error) {
	return _Feesharingsystem.Contract.LastRewardAdjustment(&_Feesharingsystem.CallOpts)
}

// LastRewardAdjustment is a free data retrieval call binding the contract method 0x40d2abae.
//
// Solidity: function lastRewardAdjustment() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCallerSession) LastRewardAdjustment() (*big.Int, error) {
	return _Feesharingsystem.Contract.LastRewardAdjustment(&_Feesharingsystem.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCaller) LastRewardBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "lastRewardBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemSession) LastRewardBlock() (*big.Int, error) {
	return _Feesharingsystem.Contract.LastRewardBlock(&_Feesharingsystem.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCallerSession) LastRewardBlock() (*big.Int, error) {
	return _Feesharingsystem.Contract.LastRewardBlock(&_Feesharingsystem.CallOpts)
}

// LastUpdateBlock is a free data retrieval call binding the contract method 0xa218141b.
//
// Solidity: function lastUpdateBlock() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCaller) LastUpdateBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "lastUpdateBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateBlock is a free data retrieval call binding the contract method 0xa218141b.
//
// Solidity: function lastUpdateBlock() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemSession) LastUpdateBlock() (*big.Int, error) {
	return _Feesharingsystem.Contract.LastUpdateBlock(&_Feesharingsystem.CallOpts)
}

// LastUpdateBlock is a free data retrieval call binding the contract method 0xa218141b.
//
// Solidity: function lastUpdateBlock() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCallerSession) LastUpdateBlock() (*big.Int, error) {
	return _Feesharingsystem.Contract.LastUpdateBlock(&_Feesharingsystem.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Feesharingsystem *FeesharingsystemCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Feesharingsystem *FeesharingsystemSession) Owner() (common.Address, error) {
	return _Feesharingsystem.Contract.Owner(&_Feesharingsystem.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Feesharingsystem *FeesharingsystemCallerSession) Owner() (common.Address, error) {
	return _Feesharingsystem.Contract.Owner(&_Feesharingsystem.CallOpts)
}

// PeriodEndBlock is a free data retrieval call binding the contract method 0x442da82f.
//
// Solidity: function periodEndBlock() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCaller) PeriodEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "periodEndBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodEndBlock is a free data retrieval call binding the contract method 0x442da82f.
//
// Solidity: function periodEndBlock() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemSession) PeriodEndBlock() (*big.Int, error) {
	return _Feesharingsystem.Contract.PeriodEndBlock(&_Feesharingsystem.CallOpts)
}

// PeriodEndBlock is a free data retrieval call binding the contract method 0x442da82f.
//
// Solidity: function periodEndBlock() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCallerSession) PeriodEndBlock() (*big.Int, error) {
	return _Feesharingsystem.Contract.PeriodEndBlock(&_Feesharingsystem.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemSession) RewardPerTokenStored() (*big.Int, error) {
	return _Feesharingsystem.Contract.RewardPerTokenStored(&_Feesharingsystem.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _Feesharingsystem.Contract.RewardPerTokenStored(&_Feesharingsystem.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Feesharingsystem *FeesharingsystemCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Feesharingsystem *FeesharingsystemSession) RewardToken() (common.Address, error) {
	return _Feesharingsystem.Contract.RewardToken(&_Feesharingsystem.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Feesharingsystem *FeesharingsystemCallerSession) RewardToken() (common.Address, error) {
	return _Feesharingsystem.Contract.RewardToken(&_Feesharingsystem.CallOpts)
}

// TokenDistributor is a free data retrieval call binding the contract method 0x18a6bc32.
//
// Solidity: function tokenDistributor() view returns(address)
func (_Feesharingsystem *FeesharingsystemCaller) TokenDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "tokenDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenDistributor is a free data retrieval call binding the contract method 0x18a6bc32.
//
// Solidity: function tokenDistributor() view returns(address)
func (_Feesharingsystem *FeesharingsystemSession) TokenDistributor() (common.Address, error) {
	return _Feesharingsystem.Contract.TokenDistributor(&_Feesharingsystem.CallOpts)
}

// TokenDistributor is a free data retrieval call binding the contract method 0x18a6bc32.
//
// Solidity: function tokenDistributor() view returns(address)
func (_Feesharingsystem *FeesharingsystemCallerSession) TokenDistributor() (common.Address, error) {
	return _Feesharingsystem.Contract.TokenDistributor(&_Feesharingsystem.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemSession) TotalShares() (*big.Int, error) {
	return _Feesharingsystem.Contract.TotalShares(&_Feesharingsystem.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Feesharingsystem *FeesharingsystemCallerSession) TotalShares() (*big.Int, error) {
	return _Feesharingsystem.Contract.TotalShares(&_Feesharingsystem.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 shares, uint256 userRewardPerTokenPaid, uint256 rewards)
func (_Feesharingsystem *FeesharingsystemCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Shares                 *big.Int
	UserRewardPerTokenPaid *big.Int
	Rewards                *big.Int
}, error) {
	var out []interface{}
	err := _Feesharingsystem.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		Shares                 *big.Int
		UserRewardPerTokenPaid *big.Int
		Rewards                *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Shares = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UserRewardPerTokenPaid = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Rewards = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 shares, uint256 userRewardPerTokenPaid, uint256 rewards)
func (_Feesharingsystem *FeesharingsystemSession) UserInfo(arg0 common.Address) (struct {
	Shares                 *big.Int
	UserRewardPerTokenPaid *big.Int
	Rewards                *big.Int
}, error) {
	return _Feesharingsystem.Contract.UserInfo(&_Feesharingsystem.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 shares, uint256 userRewardPerTokenPaid, uint256 rewards)
func (_Feesharingsystem *FeesharingsystemCallerSession) UserInfo(arg0 common.Address) (struct {
	Shares                 *big.Int
	UserRewardPerTokenPaid *big.Int
	Rewards                *big.Int
}, error) {
	return _Feesharingsystem.Contract.UserInfo(&_Feesharingsystem.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0x9a408321.
//
// Solidity: function deposit(uint256 amount, bool claimRewardToken) returns()
func (_Feesharingsystem *FeesharingsystemTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, claimRewardToken bool) (*types.Transaction, error) {
	return _Feesharingsystem.contract.Transact(opts, "deposit", amount, claimRewardToken)
}

// Deposit is a paid mutator transaction binding the contract method 0x9a408321.
//
// Solidity: function deposit(uint256 amount, bool claimRewardToken) returns()
func (_Feesharingsystem *FeesharingsystemSession) Deposit(amount *big.Int, claimRewardToken bool) (*types.Transaction, error) {
	return _Feesharingsystem.Contract.Deposit(&_Feesharingsystem.TransactOpts, amount, claimRewardToken)
}

// Deposit is a paid mutator transaction binding the contract method 0x9a408321.
//
// Solidity: function deposit(uint256 amount, bool claimRewardToken) returns()
func (_Feesharingsystem *FeesharingsystemTransactorSession) Deposit(amount *big.Int, claimRewardToken bool) (*types.Transaction, error) {
	return _Feesharingsystem.Contract.Deposit(&_Feesharingsystem.TransactOpts, amount, claimRewardToken)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_Feesharingsystem *FeesharingsystemTransactor) Harvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feesharingsystem.contract.Transact(opts, "harvest")
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_Feesharingsystem *FeesharingsystemSession) Harvest() (*types.Transaction, error) {
	return _Feesharingsystem.Contract.Harvest(&_Feesharingsystem.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_Feesharingsystem *FeesharingsystemTransactorSession) Harvest() (*types.Transaction, error) {
	return _Feesharingsystem.Contract.Harvest(&_Feesharingsystem.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Feesharingsystem *FeesharingsystemTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feesharingsystem.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Feesharingsystem *FeesharingsystemSession) RenounceOwnership() (*types.Transaction, error) {
	return _Feesharingsystem.Contract.RenounceOwnership(&_Feesharingsystem.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Feesharingsystem *FeesharingsystemTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Feesharingsystem.Contract.RenounceOwnership(&_Feesharingsystem.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Feesharingsystem *FeesharingsystemTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Feesharingsystem.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Feesharingsystem *FeesharingsystemSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Feesharingsystem.Contract.TransferOwnership(&_Feesharingsystem.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Feesharingsystem *FeesharingsystemTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Feesharingsystem.Contract.TransferOwnership(&_Feesharingsystem.TransactOpts, newOwner)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x97e50818.
//
// Solidity: function updateRewards(uint256 reward, uint256 rewardDurationInBlocks) returns()
func (_Feesharingsystem *FeesharingsystemTransactor) UpdateRewards(opts *bind.TransactOpts, reward *big.Int, rewardDurationInBlocks *big.Int) (*types.Transaction, error) {
	return _Feesharingsystem.contract.Transact(opts, "updateRewards", reward, rewardDurationInBlocks)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x97e50818.
//
// Solidity: function updateRewards(uint256 reward, uint256 rewardDurationInBlocks) returns()
func (_Feesharingsystem *FeesharingsystemSession) UpdateRewards(reward *big.Int, rewardDurationInBlocks *big.Int) (*types.Transaction, error) {
	return _Feesharingsystem.Contract.UpdateRewards(&_Feesharingsystem.TransactOpts, reward, rewardDurationInBlocks)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x97e50818.
//
// Solidity: function updateRewards(uint256 reward, uint256 rewardDurationInBlocks) returns()
func (_Feesharingsystem *FeesharingsystemTransactorSession) UpdateRewards(reward *big.Int, rewardDurationInBlocks *big.Int) (*types.Transaction, error) {
	return _Feesharingsystem.Contract.UpdateRewards(&_Feesharingsystem.TransactOpts, reward, rewardDurationInBlocks)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 shares, bool claimRewardToken) returns()
func (_Feesharingsystem *FeesharingsystemTransactor) Withdraw(opts *bind.TransactOpts, shares *big.Int, claimRewardToken bool) (*types.Transaction, error) {
	return _Feesharingsystem.contract.Transact(opts, "withdraw", shares, claimRewardToken)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 shares, bool claimRewardToken) returns()
func (_Feesharingsystem *FeesharingsystemSession) Withdraw(shares *big.Int, claimRewardToken bool) (*types.Transaction, error) {
	return _Feesharingsystem.Contract.Withdraw(&_Feesharingsystem.TransactOpts, shares, claimRewardToken)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 shares, bool claimRewardToken) returns()
func (_Feesharingsystem *FeesharingsystemTransactorSession) Withdraw(shares *big.Int, claimRewardToken bool) (*types.Transaction, error) {
	return _Feesharingsystem.Contract.Withdraw(&_Feesharingsystem.TransactOpts, shares, claimRewardToken)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x1c1c6fe5.
//
// Solidity: function withdrawAll(bool claimRewardToken) returns()
func (_Feesharingsystem *FeesharingsystemTransactor) WithdrawAll(opts *bind.TransactOpts, claimRewardToken bool) (*types.Transaction, error) {
	return _Feesharingsystem.contract.Transact(opts, "withdrawAll", claimRewardToken)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x1c1c6fe5.
//
// Solidity: function withdrawAll(bool claimRewardToken) returns()
func (_Feesharingsystem *FeesharingsystemSession) WithdrawAll(claimRewardToken bool) (*types.Transaction, error) {
	return _Feesharingsystem.Contract.WithdrawAll(&_Feesharingsystem.TransactOpts, claimRewardToken)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x1c1c6fe5.
//
// Solidity: function withdrawAll(bool claimRewardToken) returns()
func (_Feesharingsystem *FeesharingsystemTransactorSession) WithdrawAll(claimRewardToken bool) (*types.Transaction, error) {
	return _Feesharingsystem.Contract.WithdrawAll(&_Feesharingsystem.TransactOpts, claimRewardToken)
}

// FeesharingsystemDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Feesharingsystem contract.
type FeesharingsystemDepositIterator struct {
	Event *FeesharingsystemDeposit // Event containing the contract specifics and raw log

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
func (it *FeesharingsystemDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeesharingsystemDeposit)
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
		it.Event = new(FeesharingsystemDeposit)
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
func (it *FeesharingsystemDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeesharingsystemDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeesharingsystemDeposit represents a Deposit event raised by the Feesharingsystem contract.
type FeesharingsystemDeposit struct {
	User            common.Address
	Amount          *big.Int
	HarvestedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 amount, uint256 harvestedAmount)
func (_Feesharingsystem *FeesharingsystemFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*FeesharingsystemDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Feesharingsystem.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &FeesharingsystemDepositIterator{contract: _Feesharingsystem.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 amount, uint256 harvestedAmount)
func (_Feesharingsystem *FeesharingsystemFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *FeesharingsystemDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Feesharingsystem.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeesharingsystemDeposit)
				if err := _Feesharingsystem.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Feesharingsystem *FeesharingsystemFilterer) ParseDeposit(log types.Log) (*FeesharingsystemDeposit, error) {
	event := new(FeesharingsystemDeposit)
	if err := _Feesharingsystem.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeesharingsystemHarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the Feesharingsystem contract.
type FeesharingsystemHarvestIterator struct {
	Event *FeesharingsystemHarvest // Event containing the contract specifics and raw log

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
func (it *FeesharingsystemHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeesharingsystemHarvest)
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
		it.Event = new(FeesharingsystemHarvest)
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
func (it *FeesharingsystemHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeesharingsystemHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeesharingsystemHarvest represents a Harvest event raised by the Feesharingsystem contract.
type FeesharingsystemHarvest struct {
	User            common.Address
	HarvestedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0xc9695243a805adb74c91f28311176c65b417e842d5699893cef56d18bfa48cba.
//
// Solidity: event Harvest(address indexed user, uint256 harvestedAmount)
func (_Feesharingsystem *FeesharingsystemFilterer) FilterHarvest(opts *bind.FilterOpts, user []common.Address) (*FeesharingsystemHarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Feesharingsystem.contract.FilterLogs(opts, "Harvest", userRule)
	if err != nil {
		return nil, err
	}
	return &FeesharingsystemHarvestIterator{contract: _Feesharingsystem.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0xc9695243a805adb74c91f28311176c65b417e842d5699893cef56d18bfa48cba.
//
// Solidity: event Harvest(address indexed user, uint256 harvestedAmount)
func (_Feesharingsystem *FeesharingsystemFilterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *FeesharingsystemHarvest, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Feesharingsystem.contract.WatchLogs(opts, "Harvest", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeesharingsystemHarvest)
				if err := _Feesharingsystem.contract.UnpackLog(event, "Harvest", log); err != nil {
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

// ParseHarvest is a log parse operation binding the contract event 0xc9695243a805adb74c91f28311176c65b417e842d5699893cef56d18bfa48cba.
//
// Solidity: event Harvest(address indexed user, uint256 harvestedAmount)
func (_Feesharingsystem *FeesharingsystemFilterer) ParseHarvest(log types.Log) (*FeesharingsystemHarvest, error) {
	event := new(FeesharingsystemHarvest)
	if err := _Feesharingsystem.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeesharingsystemNewRewardPeriodIterator is returned from FilterNewRewardPeriod and is used to iterate over the raw logs and unpacked data for NewRewardPeriod events raised by the Feesharingsystem contract.
type FeesharingsystemNewRewardPeriodIterator struct {
	Event *FeesharingsystemNewRewardPeriod // Event containing the contract specifics and raw log

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
func (it *FeesharingsystemNewRewardPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeesharingsystemNewRewardPeriod)
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
		it.Event = new(FeesharingsystemNewRewardPeriod)
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
func (it *FeesharingsystemNewRewardPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeesharingsystemNewRewardPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeesharingsystemNewRewardPeriod represents a NewRewardPeriod event raised by the Feesharingsystem contract.
type FeesharingsystemNewRewardPeriod struct {
	NumberBlocks   *big.Int
	RewardPerBlock *big.Int
	Reward         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewRewardPeriod is a free log retrieval operation binding the contract event 0x55b4fa63fe43865f67b4f2c4a4df1cf9e6c1f85767211b44b45cf4649b2c2b51.
//
// Solidity: event NewRewardPeriod(uint256 numberBlocks, uint256 rewardPerBlock, uint256 reward)
func (_Feesharingsystem *FeesharingsystemFilterer) FilterNewRewardPeriod(opts *bind.FilterOpts) (*FeesharingsystemNewRewardPeriodIterator, error) {

	logs, sub, err := _Feesharingsystem.contract.FilterLogs(opts, "NewRewardPeriod")
	if err != nil {
		return nil, err
	}
	return &FeesharingsystemNewRewardPeriodIterator{contract: _Feesharingsystem.contract, event: "NewRewardPeriod", logs: logs, sub: sub}, nil
}

// WatchNewRewardPeriod is a free log subscription operation binding the contract event 0x55b4fa63fe43865f67b4f2c4a4df1cf9e6c1f85767211b44b45cf4649b2c2b51.
//
// Solidity: event NewRewardPeriod(uint256 numberBlocks, uint256 rewardPerBlock, uint256 reward)
func (_Feesharingsystem *FeesharingsystemFilterer) WatchNewRewardPeriod(opts *bind.WatchOpts, sink chan<- *FeesharingsystemNewRewardPeriod) (event.Subscription, error) {

	logs, sub, err := _Feesharingsystem.contract.WatchLogs(opts, "NewRewardPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeesharingsystemNewRewardPeriod)
				if err := _Feesharingsystem.contract.UnpackLog(event, "NewRewardPeriod", log); err != nil {
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
func (_Feesharingsystem *FeesharingsystemFilterer) ParseNewRewardPeriod(log types.Log) (*FeesharingsystemNewRewardPeriod, error) {
	event := new(FeesharingsystemNewRewardPeriod)
	if err := _Feesharingsystem.contract.UnpackLog(event, "NewRewardPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeesharingsystemOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Feesharingsystem contract.
type FeesharingsystemOwnershipTransferredIterator struct {
	Event *FeesharingsystemOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeesharingsystemOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeesharingsystemOwnershipTransferred)
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
		it.Event = new(FeesharingsystemOwnershipTransferred)
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
func (it *FeesharingsystemOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeesharingsystemOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeesharingsystemOwnershipTransferred represents a OwnershipTransferred event raised by the Feesharingsystem contract.
type FeesharingsystemOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Feesharingsystem *FeesharingsystemFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeesharingsystemOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Feesharingsystem.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeesharingsystemOwnershipTransferredIterator{contract: _Feesharingsystem.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Feesharingsystem *FeesharingsystemFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeesharingsystemOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Feesharingsystem.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeesharingsystemOwnershipTransferred)
				if err := _Feesharingsystem.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Feesharingsystem *FeesharingsystemFilterer) ParseOwnershipTransferred(log types.Log) (*FeesharingsystemOwnershipTransferred, error) {
	event := new(FeesharingsystemOwnershipTransferred)
	if err := _Feesharingsystem.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeesharingsystemWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Feesharingsystem contract.
type FeesharingsystemWithdrawIterator struct {
	Event *FeesharingsystemWithdraw // Event containing the contract specifics and raw log

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
func (it *FeesharingsystemWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeesharingsystemWithdraw)
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
		it.Event = new(FeesharingsystemWithdraw)
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
func (it *FeesharingsystemWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeesharingsystemWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeesharingsystemWithdraw represents a Withdraw event raised by the Feesharingsystem contract.
type FeesharingsystemWithdraw struct {
	User            common.Address
	Amount          *big.Int
	HarvestedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 amount, uint256 harvestedAmount)
func (_Feesharingsystem *FeesharingsystemFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*FeesharingsystemWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Feesharingsystem.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &FeesharingsystemWithdrawIterator{contract: _Feesharingsystem.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 amount, uint256 harvestedAmount)
func (_Feesharingsystem *FeesharingsystemFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *FeesharingsystemWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Feesharingsystem.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeesharingsystemWithdraw)
				if err := _Feesharingsystem.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Feesharingsystem *FeesharingsystemFilterer) ParseWithdraw(log types.Log) (*FeesharingsystemWithdraw, error) {
	event := new(FeesharingsystemWithdraw)
	if err := _Feesharingsystem.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
