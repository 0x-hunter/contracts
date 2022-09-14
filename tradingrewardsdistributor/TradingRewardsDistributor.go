// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tradingrewardsdistributor

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

// TradingrewardsdistributorMetaData contains all meta data concerning the Tradingrewardsdistributor contract.
var TradingrewardsdistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hunterGovernanceToken_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardRound\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrawnOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardRound\",\"type\":\"uint256\"}],\"name\":\"UpdateTradingRewards\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BUFFER_ADMIN_WITHDRAW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"amountClaimedByUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"canClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRewardRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasUserClaimedForRewardRound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hunterGovernanceToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPausedTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maximumAmountPerUserInCurrentTree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleRootOfRewardRound\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"merkleRootUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseDistribution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseDistribution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"newMaximumAmountPerUser\",\"type\":\"uint256\"}],\"name\":\"updateTradingRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokenRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TradingrewardsdistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use TradingrewardsdistributorMetaData.ABI instead.
var TradingrewardsdistributorABI = TradingrewardsdistributorMetaData.ABI

// Tradingrewardsdistributor is an auto generated Go binding around an Ethereum contract.
type Tradingrewardsdistributor struct {
	TradingrewardsdistributorCaller     // Read-only binding to the contract
	TradingrewardsdistributorTransactor // Write-only binding to the contract
	TradingrewardsdistributorFilterer   // Log filterer for contract events
}

// TradingrewardsdistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradingrewardsdistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingrewardsdistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradingrewardsdistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingrewardsdistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradingrewardsdistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingrewardsdistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradingrewardsdistributorSession struct {
	Contract     *Tradingrewardsdistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TradingrewardsdistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradingrewardsdistributorCallerSession struct {
	Contract *TradingrewardsdistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// TradingrewardsdistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradingrewardsdistributorTransactorSession struct {
	Contract     *TradingrewardsdistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// TradingrewardsdistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradingrewardsdistributorRaw struct {
	Contract *Tradingrewardsdistributor // Generic contract binding to access the raw methods on
}

// TradingrewardsdistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradingrewardsdistributorCallerRaw struct {
	Contract *TradingrewardsdistributorCaller // Generic read-only contract binding to access the raw methods on
}

// TradingrewardsdistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradingrewardsdistributorTransactorRaw struct {
	Contract *TradingrewardsdistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTradingrewardsdistributor creates a new instance of Tradingrewardsdistributor, bound to a specific deployed contract.
func NewTradingrewardsdistributor(address common.Address, backend bind.ContractBackend) (*Tradingrewardsdistributor, error) {
	contract, err := bindTradingrewardsdistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tradingrewardsdistributor{TradingrewardsdistributorCaller: TradingrewardsdistributorCaller{contract: contract}, TradingrewardsdistributorTransactor: TradingrewardsdistributorTransactor{contract: contract}, TradingrewardsdistributorFilterer: TradingrewardsdistributorFilterer{contract: contract}}, nil
}

// NewTradingrewardsdistributorCaller creates a new read-only instance of Tradingrewardsdistributor, bound to a specific deployed contract.
func NewTradingrewardsdistributorCaller(address common.Address, caller bind.ContractCaller) (*TradingrewardsdistributorCaller, error) {
	contract, err := bindTradingrewardsdistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradingrewardsdistributorCaller{contract: contract}, nil
}

// NewTradingrewardsdistributorTransactor creates a new write-only instance of Tradingrewardsdistributor, bound to a specific deployed contract.
func NewTradingrewardsdistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*TradingrewardsdistributorTransactor, error) {
	contract, err := bindTradingrewardsdistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradingrewardsdistributorTransactor{contract: contract}, nil
}

// NewTradingrewardsdistributorFilterer creates a new log filterer instance of Tradingrewardsdistributor, bound to a specific deployed contract.
func NewTradingrewardsdistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*TradingrewardsdistributorFilterer, error) {
	contract, err := bindTradingrewardsdistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradingrewardsdistributorFilterer{contract: contract}, nil
}

// bindTradingrewardsdistributor binds a generic wrapper to an already deployed contract.
func bindTradingrewardsdistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TradingrewardsdistributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tradingrewardsdistributor *TradingrewardsdistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tradingrewardsdistributor.Contract.TradingrewardsdistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tradingrewardsdistributor *TradingrewardsdistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.TradingrewardsdistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tradingrewardsdistributor *TradingrewardsdistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.TradingrewardsdistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tradingrewardsdistributor *TradingrewardsdistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tradingrewardsdistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.contract.Transact(opts, method, params...)
}

// BUFFERADMINWITHDRAW is a free data retrieval call binding the contract method 0x141fd8c8.
//
// Solidity: function BUFFER_ADMIN_WITHDRAW() view returns(uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCaller) BUFFERADMINWITHDRAW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tradingrewardsdistributor.contract.Call(opts, &out, "BUFFER_ADMIN_WITHDRAW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BUFFERADMINWITHDRAW is a free data retrieval call binding the contract method 0x141fd8c8.
//
// Solidity: function BUFFER_ADMIN_WITHDRAW() view returns(uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) BUFFERADMINWITHDRAW() (*big.Int, error) {
	return _Tradingrewardsdistributor.Contract.BUFFERADMINWITHDRAW(&_Tradingrewardsdistributor.CallOpts)
}

// BUFFERADMINWITHDRAW is a free data retrieval call binding the contract method 0x141fd8c8.
//
// Solidity: function BUFFER_ADMIN_WITHDRAW() view returns(uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCallerSession) BUFFERADMINWITHDRAW() (*big.Int, error) {
	return _Tradingrewardsdistributor.Contract.BUFFERADMINWITHDRAW(&_Tradingrewardsdistributor.CallOpts)
}

// AmountClaimedByUser is a free data retrieval call binding the contract method 0x3b66f49d.
//
// Solidity: function amountClaimedByUser(address ) view returns(uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCaller) AmountClaimedByUser(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tradingrewardsdistributor.contract.Call(opts, &out, "amountClaimedByUser", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountClaimedByUser is a free data retrieval call binding the contract method 0x3b66f49d.
//
// Solidity: function amountClaimedByUser(address ) view returns(uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) AmountClaimedByUser(arg0 common.Address) (*big.Int, error) {
	return _Tradingrewardsdistributor.Contract.AmountClaimedByUser(&_Tradingrewardsdistributor.CallOpts, arg0)
}

// AmountClaimedByUser is a free data retrieval call binding the contract method 0x3b66f49d.
//
// Solidity: function amountClaimedByUser(address ) view returns(uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCallerSession) AmountClaimedByUser(arg0 common.Address) (*big.Int, error) {
	return _Tradingrewardsdistributor.Contract.AmountClaimedByUser(&_Tradingrewardsdistributor.CallOpts, arg0)
}

// CanClaim is a free data retrieval call binding the contract method 0xdc38bdb5.
//
// Solidity: function canClaim(address user, uint256 amount, bytes32[] merkleProof) view returns(bool, uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCaller) CanClaim(opts *bind.CallOpts, user common.Address, amount *big.Int, merkleProof [][32]byte) (bool, *big.Int, error) {
	var out []interface{}
	err := _Tradingrewardsdistributor.contract.Call(opts, &out, "canClaim", user, amount, merkleProof)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// CanClaim is a free data retrieval call binding the contract method 0xdc38bdb5.
//
// Solidity: function canClaim(address user, uint256 amount, bytes32[] merkleProof) view returns(bool, uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) CanClaim(user common.Address, amount *big.Int, merkleProof [][32]byte) (bool, *big.Int, error) {
	return _Tradingrewardsdistributor.Contract.CanClaim(&_Tradingrewardsdistributor.CallOpts, user, amount, merkleProof)
}

// CanClaim is a free data retrieval call binding the contract method 0xdc38bdb5.
//
// Solidity: function canClaim(address user, uint256 amount, bytes32[] merkleProof) view returns(bool, uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCallerSession) CanClaim(user common.Address, amount *big.Int, merkleProof [][32]byte) (bool, *big.Int, error) {
	return _Tradingrewardsdistributor.Contract.CanClaim(&_Tradingrewardsdistributor.CallOpts, user, amount, merkleProof)
}

// CurrentRewardRound is a free data retrieval call binding the contract method 0x1040faf9.
//
// Solidity: function currentRewardRound() view returns(uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCaller) CurrentRewardRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tradingrewardsdistributor.contract.Call(opts, &out, "currentRewardRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRewardRound is a free data retrieval call binding the contract method 0x1040faf9.
//
// Solidity: function currentRewardRound() view returns(uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) CurrentRewardRound() (*big.Int, error) {
	return _Tradingrewardsdistributor.Contract.CurrentRewardRound(&_Tradingrewardsdistributor.CallOpts)
}

// CurrentRewardRound is a free data retrieval call binding the contract method 0x1040faf9.
//
// Solidity: function currentRewardRound() view returns(uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCallerSession) CurrentRewardRound() (*big.Int, error) {
	return _Tradingrewardsdistributor.Contract.CurrentRewardRound(&_Tradingrewardsdistributor.CallOpts)
}

// HasUserClaimedForRewardRound is a free data retrieval call binding the contract method 0x9ea76f34.
//
// Solidity: function hasUserClaimedForRewardRound(uint256 , address ) view returns(bool)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCaller) HasUserClaimedForRewardRound(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Tradingrewardsdistributor.contract.Call(opts, &out, "hasUserClaimedForRewardRound", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUserClaimedForRewardRound is a free data retrieval call binding the contract method 0x9ea76f34.
//
// Solidity: function hasUserClaimedForRewardRound(uint256 , address ) view returns(bool)
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) HasUserClaimedForRewardRound(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Tradingrewardsdistributor.Contract.HasUserClaimedForRewardRound(&_Tradingrewardsdistributor.CallOpts, arg0, arg1)
}

// HasUserClaimedForRewardRound is a free data retrieval call binding the contract method 0x9ea76f34.
//
// Solidity: function hasUserClaimedForRewardRound(uint256 , address ) view returns(bool)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCallerSession) HasUserClaimedForRewardRound(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Tradingrewardsdistributor.Contract.HasUserClaimedForRewardRound(&_Tradingrewardsdistributor.CallOpts, arg0, arg1)
}

// HunterGovernanceToken is a free data retrieval call binding the contract method 0xf363264a.
//
// Solidity: function hunterGovernanceToken() view returns(address)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCaller) HunterGovernanceToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tradingrewardsdistributor.contract.Call(opts, &out, "hunterGovernanceToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HunterGovernanceToken is a free data retrieval call binding the contract method 0xf363264a.
//
// Solidity: function hunterGovernanceToken() view returns(address)
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) HunterGovernanceToken() (common.Address, error) {
	return _Tradingrewardsdistributor.Contract.HunterGovernanceToken(&_Tradingrewardsdistributor.CallOpts)
}

// HunterGovernanceToken is a free data retrieval call binding the contract method 0xf363264a.
//
// Solidity: function hunterGovernanceToken() view returns(address)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCallerSession) HunterGovernanceToken() (common.Address, error) {
	return _Tradingrewardsdistributor.Contract.HunterGovernanceToken(&_Tradingrewardsdistributor.CallOpts)
}

// LastPausedTimestamp is a free data retrieval call binding the contract method 0xfa461974.
//
// Solidity: function lastPausedTimestamp() view returns(uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCaller) LastPausedTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tradingrewardsdistributor.contract.Call(opts, &out, "lastPausedTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPausedTimestamp is a free data retrieval call binding the contract method 0xfa461974.
//
// Solidity: function lastPausedTimestamp() view returns(uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) LastPausedTimestamp() (*big.Int, error) {
	return _Tradingrewardsdistributor.Contract.LastPausedTimestamp(&_Tradingrewardsdistributor.CallOpts)
}

// LastPausedTimestamp is a free data retrieval call binding the contract method 0xfa461974.
//
// Solidity: function lastPausedTimestamp() view returns(uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCallerSession) LastPausedTimestamp() (*big.Int, error) {
	return _Tradingrewardsdistributor.Contract.LastPausedTimestamp(&_Tradingrewardsdistributor.CallOpts)
}

// MaximumAmountPerUserInCurrentTree is a free data retrieval call binding the contract method 0xff0a4eb0.
//
// Solidity: function maximumAmountPerUserInCurrentTree() view returns(uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCaller) MaximumAmountPerUserInCurrentTree(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tradingrewardsdistributor.contract.Call(opts, &out, "maximumAmountPerUserInCurrentTree")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumAmountPerUserInCurrentTree is a free data retrieval call binding the contract method 0xff0a4eb0.
//
// Solidity: function maximumAmountPerUserInCurrentTree() view returns(uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) MaximumAmountPerUserInCurrentTree() (*big.Int, error) {
	return _Tradingrewardsdistributor.Contract.MaximumAmountPerUserInCurrentTree(&_Tradingrewardsdistributor.CallOpts)
}

// MaximumAmountPerUserInCurrentTree is a free data retrieval call binding the contract method 0xff0a4eb0.
//
// Solidity: function maximumAmountPerUserInCurrentTree() view returns(uint256)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCallerSession) MaximumAmountPerUserInCurrentTree() (*big.Int, error) {
	return _Tradingrewardsdistributor.Contract.MaximumAmountPerUserInCurrentTree(&_Tradingrewardsdistributor.CallOpts)
}

// MerkleRootOfRewardRound is a free data retrieval call binding the contract method 0x84483c86.
//
// Solidity: function merkleRootOfRewardRound(uint256 ) view returns(bytes32)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCaller) MerkleRootOfRewardRound(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Tradingrewardsdistributor.contract.Call(opts, &out, "merkleRootOfRewardRound", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRootOfRewardRound is a free data retrieval call binding the contract method 0x84483c86.
//
// Solidity: function merkleRootOfRewardRound(uint256 ) view returns(bytes32)
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) MerkleRootOfRewardRound(arg0 *big.Int) ([32]byte, error) {
	return _Tradingrewardsdistributor.Contract.MerkleRootOfRewardRound(&_Tradingrewardsdistributor.CallOpts, arg0)
}

// MerkleRootOfRewardRound is a free data retrieval call binding the contract method 0x84483c86.
//
// Solidity: function merkleRootOfRewardRound(uint256 ) view returns(bytes32)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCallerSession) MerkleRootOfRewardRound(arg0 *big.Int) ([32]byte, error) {
	return _Tradingrewardsdistributor.Contract.MerkleRootOfRewardRound(&_Tradingrewardsdistributor.CallOpts, arg0)
}

// MerkleRootUsed is a free data retrieval call binding the contract method 0xba42590d.
//
// Solidity: function merkleRootUsed(bytes32 ) view returns(bool)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCaller) MerkleRootUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Tradingrewardsdistributor.contract.Call(opts, &out, "merkleRootUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MerkleRootUsed is a free data retrieval call binding the contract method 0xba42590d.
//
// Solidity: function merkleRootUsed(bytes32 ) view returns(bool)
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) MerkleRootUsed(arg0 [32]byte) (bool, error) {
	return _Tradingrewardsdistributor.Contract.MerkleRootUsed(&_Tradingrewardsdistributor.CallOpts, arg0)
}

// MerkleRootUsed is a free data retrieval call binding the contract method 0xba42590d.
//
// Solidity: function merkleRootUsed(bytes32 ) view returns(bool)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCallerSession) MerkleRootUsed(arg0 [32]byte) (bool, error) {
	return _Tradingrewardsdistributor.Contract.MerkleRootUsed(&_Tradingrewardsdistributor.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tradingrewardsdistributor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) Owner() (common.Address, error) {
	return _Tradingrewardsdistributor.Contract.Owner(&_Tradingrewardsdistributor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCallerSession) Owner() (common.Address, error) {
	return _Tradingrewardsdistributor.Contract.Owner(&_Tradingrewardsdistributor.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tradingrewardsdistributor.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) Paused() (bool, error) {
	return _Tradingrewardsdistributor.Contract.Paused(&_Tradingrewardsdistributor.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Tradingrewardsdistributor *TradingrewardsdistributorCallerSession) Paused() (bool, error) {
	return _Tradingrewardsdistributor.Contract.Paused(&_Tradingrewardsdistributor.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x2f52ebb7.
//
// Solidity: function claim(uint256 amount, bytes32[] merkleProof) returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactor) Claim(opts *bind.TransactOpts, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.contract.Transact(opts, "claim", amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x2f52ebb7.
//
// Solidity: function claim(uint256 amount, bytes32[] merkleProof) returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) Claim(amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.Claim(&_Tradingrewardsdistributor.TransactOpts, amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x2f52ebb7.
//
// Solidity: function claim(uint256 amount, bytes32[] merkleProof) returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactorSession) Claim(amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.Claim(&_Tradingrewardsdistributor.TransactOpts, amount, merkleProof)
}

// PauseDistribution is a paid mutator transaction binding the contract method 0x31cec7a3.
//
// Solidity: function pauseDistribution() returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactor) PauseDistribution(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.contract.Transact(opts, "pauseDistribution")
}

// PauseDistribution is a paid mutator transaction binding the contract method 0x31cec7a3.
//
// Solidity: function pauseDistribution() returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) PauseDistribution() (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.PauseDistribution(&_Tradingrewardsdistributor.TransactOpts)
}

// PauseDistribution is a paid mutator transaction binding the contract method 0x31cec7a3.
//
// Solidity: function pauseDistribution() returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactorSession) PauseDistribution() (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.PauseDistribution(&_Tradingrewardsdistributor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.RenounceOwnership(&_Tradingrewardsdistributor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.RenounceOwnership(&_Tradingrewardsdistributor.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.TransferOwnership(&_Tradingrewardsdistributor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.TransferOwnership(&_Tradingrewardsdistributor.TransactOpts, newOwner)
}

// UnpauseDistribution is a paid mutator transaction binding the contract method 0xb94ec9d0.
//
// Solidity: function unpauseDistribution() returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactor) UnpauseDistribution(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.contract.Transact(opts, "unpauseDistribution")
}

// UnpauseDistribution is a paid mutator transaction binding the contract method 0xb94ec9d0.
//
// Solidity: function unpauseDistribution() returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) UnpauseDistribution() (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.UnpauseDistribution(&_Tradingrewardsdistributor.TransactOpts)
}

// UnpauseDistribution is a paid mutator transaction binding the contract method 0xb94ec9d0.
//
// Solidity: function unpauseDistribution() returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactorSession) UnpauseDistribution() (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.UnpauseDistribution(&_Tradingrewardsdistributor.TransactOpts)
}

// UpdateTradingRewards is a paid mutator transaction binding the contract method 0x847fadca.
//
// Solidity: function updateTradingRewards(bytes32 merkleRoot, uint256 newMaximumAmountPerUser) returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactor) UpdateTradingRewards(opts *bind.TransactOpts, merkleRoot [32]byte, newMaximumAmountPerUser *big.Int) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.contract.Transact(opts, "updateTradingRewards", merkleRoot, newMaximumAmountPerUser)
}

// UpdateTradingRewards is a paid mutator transaction binding the contract method 0x847fadca.
//
// Solidity: function updateTradingRewards(bytes32 merkleRoot, uint256 newMaximumAmountPerUser) returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) UpdateTradingRewards(merkleRoot [32]byte, newMaximumAmountPerUser *big.Int) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.UpdateTradingRewards(&_Tradingrewardsdistributor.TransactOpts, merkleRoot, newMaximumAmountPerUser)
}

// UpdateTradingRewards is a paid mutator transaction binding the contract method 0x847fadca.
//
// Solidity: function updateTradingRewards(bytes32 merkleRoot, uint256 newMaximumAmountPerUser) returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactorSession) UpdateTradingRewards(merkleRoot [32]byte, newMaximumAmountPerUser *big.Int) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.UpdateTradingRewards(&_Tradingrewardsdistributor.TransactOpts, merkleRoot, newMaximumAmountPerUser)
}

// WithdrawTokenRewards is a paid mutator transaction binding the contract method 0x8923e2f7.
//
// Solidity: function withdrawTokenRewards(uint256 amount) returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactor) WithdrawTokenRewards(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.contract.Transact(opts, "withdrawTokenRewards", amount)
}

// WithdrawTokenRewards is a paid mutator transaction binding the contract method 0x8923e2f7.
//
// Solidity: function withdrawTokenRewards(uint256 amount) returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorSession) WithdrawTokenRewards(amount *big.Int) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.WithdrawTokenRewards(&_Tradingrewardsdistributor.TransactOpts, amount)
}

// WithdrawTokenRewards is a paid mutator transaction binding the contract method 0x8923e2f7.
//
// Solidity: function withdrawTokenRewards(uint256 amount) returns()
func (_Tradingrewardsdistributor *TradingrewardsdistributorTransactorSession) WithdrawTokenRewards(amount *big.Int) (*types.Transaction, error) {
	return _Tradingrewardsdistributor.Contract.WithdrawTokenRewards(&_Tradingrewardsdistributor.TransactOpts, amount)
}

// TradingrewardsdistributorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tradingrewardsdistributor contract.
type TradingrewardsdistributorOwnershipTransferredIterator struct {
	Event *TradingrewardsdistributorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TradingrewardsdistributorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingrewardsdistributorOwnershipTransferred)
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
		it.Event = new(TradingrewardsdistributorOwnershipTransferred)
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
func (it *TradingrewardsdistributorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingrewardsdistributorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingrewardsdistributorOwnershipTransferred represents a OwnershipTransferred event raised by the Tradingrewardsdistributor contract.
type TradingrewardsdistributorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TradingrewardsdistributorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tradingrewardsdistributor.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TradingrewardsdistributorOwnershipTransferredIterator{contract: _Tradingrewardsdistributor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TradingrewardsdistributorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tradingrewardsdistributor.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingrewardsdistributorOwnershipTransferred)
				if err := _Tradingrewardsdistributor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) ParseOwnershipTransferred(log types.Log) (*TradingrewardsdistributorOwnershipTransferred, error) {
	event := new(TradingrewardsdistributorOwnershipTransferred)
	if err := _Tradingrewardsdistributor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingrewardsdistributorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Tradingrewardsdistributor contract.
type TradingrewardsdistributorPausedIterator struct {
	Event *TradingrewardsdistributorPaused // Event containing the contract specifics and raw log

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
func (it *TradingrewardsdistributorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingrewardsdistributorPaused)
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
		it.Event = new(TradingrewardsdistributorPaused)
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
func (it *TradingrewardsdistributorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingrewardsdistributorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingrewardsdistributorPaused represents a Paused event raised by the Tradingrewardsdistributor contract.
type TradingrewardsdistributorPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) FilterPaused(opts *bind.FilterOpts) (*TradingrewardsdistributorPausedIterator, error) {

	logs, sub, err := _Tradingrewardsdistributor.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TradingrewardsdistributorPausedIterator{contract: _Tradingrewardsdistributor.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TradingrewardsdistributorPaused) (event.Subscription, error) {

	logs, sub, err := _Tradingrewardsdistributor.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingrewardsdistributorPaused)
				if err := _Tradingrewardsdistributor.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) ParsePaused(log types.Log) (*TradingrewardsdistributorPaused, error) {
	event := new(TradingrewardsdistributorPaused)
	if err := _Tradingrewardsdistributor.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingrewardsdistributorRewardsClaimIterator is returned from FilterRewardsClaim and is used to iterate over the raw logs and unpacked data for RewardsClaim events raised by the Tradingrewardsdistributor contract.
type TradingrewardsdistributorRewardsClaimIterator struct {
	Event *TradingrewardsdistributorRewardsClaim // Event containing the contract specifics and raw log

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
func (it *TradingrewardsdistributorRewardsClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingrewardsdistributorRewardsClaim)
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
		it.Event = new(TradingrewardsdistributorRewardsClaim)
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
func (it *TradingrewardsdistributorRewardsClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingrewardsdistributorRewardsClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingrewardsdistributorRewardsClaim represents a RewardsClaim event raised by the Tradingrewardsdistributor contract.
type TradingrewardsdistributorRewardsClaim struct {
	User        common.Address
	RewardRound *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaim is a free log retrieval operation binding the contract event 0xe8dbd8c18906df11a70832fe7c874d1fddf6952cee2320658d66925618387999.
//
// Solidity: event RewardsClaim(address indexed user, uint256 indexed rewardRound, uint256 amount)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) FilterRewardsClaim(opts *bind.FilterOpts, user []common.Address, rewardRound []*big.Int) (*TradingrewardsdistributorRewardsClaimIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardRoundRule []interface{}
	for _, rewardRoundItem := range rewardRound {
		rewardRoundRule = append(rewardRoundRule, rewardRoundItem)
	}

	logs, sub, err := _Tradingrewardsdistributor.contract.FilterLogs(opts, "RewardsClaim", userRule, rewardRoundRule)
	if err != nil {
		return nil, err
	}
	return &TradingrewardsdistributorRewardsClaimIterator{contract: _Tradingrewardsdistributor.contract, event: "RewardsClaim", logs: logs, sub: sub}, nil
}

// WatchRewardsClaim is a free log subscription operation binding the contract event 0xe8dbd8c18906df11a70832fe7c874d1fddf6952cee2320658d66925618387999.
//
// Solidity: event RewardsClaim(address indexed user, uint256 indexed rewardRound, uint256 amount)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) WatchRewardsClaim(opts *bind.WatchOpts, sink chan<- *TradingrewardsdistributorRewardsClaim, user []common.Address, rewardRound []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardRoundRule []interface{}
	for _, rewardRoundItem := range rewardRound {
		rewardRoundRule = append(rewardRoundRule, rewardRoundItem)
	}

	logs, sub, err := _Tradingrewardsdistributor.contract.WatchLogs(opts, "RewardsClaim", userRule, rewardRoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingrewardsdistributorRewardsClaim)
				if err := _Tradingrewardsdistributor.contract.UnpackLog(event, "RewardsClaim", log); err != nil {
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

// ParseRewardsClaim is a log parse operation binding the contract event 0xe8dbd8c18906df11a70832fe7c874d1fddf6952cee2320658d66925618387999.
//
// Solidity: event RewardsClaim(address indexed user, uint256 indexed rewardRound, uint256 amount)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) ParseRewardsClaim(log types.Log) (*TradingrewardsdistributorRewardsClaim, error) {
	event := new(TradingrewardsdistributorRewardsClaim)
	if err := _Tradingrewardsdistributor.contract.UnpackLog(event, "RewardsClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingrewardsdistributorTokenWithdrawnOwnerIterator is returned from FilterTokenWithdrawnOwner and is used to iterate over the raw logs and unpacked data for TokenWithdrawnOwner events raised by the Tradingrewardsdistributor contract.
type TradingrewardsdistributorTokenWithdrawnOwnerIterator struct {
	Event *TradingrewardsdistributorTokenWithdrawnOwner // Event containing the contract specifics and raw log

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
func (it *TradingrewardsdistributorTokenWithdrawnOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingrewardsdistributorTokenWithdrawnOwner)
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
		it.Event = new(TradingrewardsdistributorTokenWithdrawnOwner)
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
func (it *TradingrewardsdistributorTokenWithdrawnOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingrewardsdistributorTokenWithdrawnOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingrewardsdistributorTokenWithdrawnOwner represents a TokenWithdrawnOwner event raised by the Tradingrewardsdistributor contract.
type TradingrewardsdistributorTokenWithdrawnOwner struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawnOwner is a free log retrieval operation binding the contract event 0xb4cd5c4a08bbed33abfe773ece179d156730e39629e065b7dcd8263027387c1d.
//
// Solidity: event TokenWithdrawnOwner(uint256 amount)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) FilterTokenWithdrawnOwner(opts *bind.FilterOpts) (*TradingrewardsdistributorTokenWithdrawnOwnerIterator, error) {

	logs, sub, err := _Tradingrewardsdistributor.contract.FilterLogs(opts, "TokenWithdrawnOwner")
	if err != nil {
		return nil, err
	}
	return &TradingrewardsdistributorTokenWithdrawnOwnerIterator{contract: _Tradingrewardsdistributor.contract, event: "TokenWithdrawnOwner", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawnOwner is a free log subscription operation binding the contract event 0xb4cd5c4a08bbed33abfe773ece179d156730e39629e065b7dcd8263027387c1d.
//
// Solidity: event TokenWithdrawnOwner(uint256 amount)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) WatchTokenWithdrawnOwner(opts *bind.WatchOpts, sink chan<- *TradingrewardsdistributorTokenWithdrawnOwner) (event.Subscription, error) {

	logs, sub, err := _Tradingrewardsdistributor.contract.WatchLogs(opts, "TokenWithdrawnOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingrewardsdistributorTokenWithdrawnOwner)
				if err := _Tradingrewardsdistributor.contract.UnpackLog(event, "TokenWithdrawnOwner", log); err != nil {
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

// ParseTokenWithdrawnOwner is a log parse operation binding the contract event 0xb4cd5c4a08bbed33abfe773ece179d156730e39629e065b7dcd8263027387c1d.
//
// Solidity: event TokenWithdrawnOwner(uint256 amount)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) ParseTokenWithdrawnOwner(log types.Log) (*TradingrewardsdistributorTokenWithdrawnOwner, error) {
	event := new(TradingrewardsdistributorTokenWithdrawnOwner)
	if err := _Tradingrewardsdistributor.contract.UnpackLog(event, "TokenWithdrawnOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingrewardsdistributorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Tradingrewardsdistributor contract.
type TradingrewardsdistributorUnpausedIterator struct {
	Event *TradingrewardsdistributorUnpaused // Event containing the contract specifics and raw log

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
func (it *TradingrewardsdistributorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingrewardsdistributorUnpaused)
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
		it.Event = new(TradingrewardsdistributorUnpaused)
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
func (it *TradingrewardsdistributorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingrewardsdistributorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingrewardsdistributorUnpaused represents a Unpaused event raised by the Tradingrewardsdistributor contract.
type TradingrewardsdistributorUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TradingrewardsdistributorUnpausedIterator, error) {

	logs, sub, err := _Tradingrewardsdistributor.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TradingrewardsdistributorUnpausedIterator{contract: _Tradingrewardsdistributor.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TradingrewardsdistributorUnpaused) (event.Subscription, error) {

	logs, sub, err := _Tradingrewardsdistributor.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingrewardsdistributorUnpaused)
				if err := _Tradingrewardsdistributor.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) ParseUnpaused(log types.Log) (*TradingrewardsdistributorUnpaused, error) {
	event := new(TradingrewardsdistributorUnpaused)
	if err := _Tradingrewardsdistributor.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingrewardsdistributorUpdateTradingRewardsIterator is returned from FilterUpdateTradingRewards and is used to iterate over the raw logs and unpacked data for UpdateTradingRewards events raised by the Tradingrewardsdistributor contract.
type TradingrewardsdistributorUpdateTradingRewardsIterator struct {
	Event *TradingrewardsdistributorUpdateTradingRewards // Event containing the contract specifics and raw log

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
func (it *TradingrewardsdistributorUpdateTradingRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingrewardsdistributorUpdateTradingRewards)
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
		it.Event = new(TradingrewardsdistributorUpdateTradingRewards)
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
func (it *TradingrewardsdistributorUpdateTradingRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingrewardsdistributorUpdateTradingRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingrewardsdistributorUpdateTradingRewards represents a UpdateTradingRewards event raised by the Tradingrewardsdistributor contract.
type TradingrewardsdistributorUpdateTradingRewards struct {
	RewardRound *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateTradingRewards is a free log retrieval operation binding the contract event 0x0b1d6b49e69d133866d064a73f0181e1e35aa94fde3ef0095587d7ebd0423caa.
//
// Solidity: event UpdateTradingRewards(uint256 indexed rewardRound)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) FilterUpdateTradingRewards(opts *bind.FilterOpts, rewardRound []*big.Int) (*TradingrewardsdistributorUpdateTradingRewardsIterator, error) {

	var rewardRoundRule []interface{}
	for _, rewardRoundItem := range rewardRound {
		rewardRoundRule = append(rewardRoundRule, rewardRoundItem)
	}

	logs, sub, err := _Tradingrewardsdistributor.contract.FilterLogs(opts, "UpdateTradingRewards", rewardRoundRule)
	if err != nil {
		return nil, err
	}
	return &TradingrewardsdistributorUpdateTradingRewardsIterator{contract: _Tradingrewardsdistributor.contract, event: "UpdateTradingRewards", logs: logs, sub: sub}, nil
}

// WatchUpdateTradingRewards is a free log subscription operation binding the contract event 0x0b1d6b49e69d133866d064a73f0181e1e35aa94fde3ef0095587d7ebd0423caa.
//
// Solidity: event UpdateTradingRewards(uint256 indexed rewardRound)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) WatchUpdateTradingRewards(opts *bind.WatchOpts, sink chan<- *TradingrewardsdistributorUpdateTradingRewards, rewardRound []*big.Int) (event.Subscription, error) {

	var rewardRoundRule []interface{}
	for _, rewardRoundItem := range rewardRound {
		rewardRoundRule = append(rewardRoundRule, rewardRoundItem)
	}

	logs, sub, err := _Tradingrewardsdistributor.contract.WatchLogs(opts, "UpdateTradingRewards", rewardRoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingrewardsdistributorUpdateTradingRewards)
				if err := _Tradingrewardsdistributor.contract.UnpackLog(event, "UpdateTradingRewards", log); err != nil {
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

// ParseUpdateTradingRewards is a log parse operation binding the contract event 0x0b1d6b49e69d133866d064a73f0181e1e35aa94fde3ef0095587d7ebd0423caa.
//
// Solidity: event UpdateTradingRewards(uint256 indexed rewardRound)
func (_Tradingrewardsdistributor *TradingrewardsdistributorFilterer) ParseUpdateTradingRewards(log types.Log) (*TradingrewardsdistributorUpdateTradingRewards, error) {
	event := new(TradingrewardsdistributorUpdateTradingRewards)
	if err := _Tradingrewardsdistributor.contract.UnpackLog(event, "UpdateTradingRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
