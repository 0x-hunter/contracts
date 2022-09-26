// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokensplitter

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

// TokensplitterMetaData contains all meta data concerning the Tokensplitter contract.
var TokensplitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_shares\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_hunterGovernanceToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"NewSharesOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TOTAL_SHARES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensDistributedToAccount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"calculatePendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hunterGovernanceToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"releaseTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTokensDistributed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_currentRecipient\",\"type\":\"address\"}],\"name\":\"updateSharesOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokensplitterABI is the input ABI used to generate the binding from.
// Deprecated: Use TokensplitterMetaData.ABI instead.
var TokensplitterABI = TokensplitterMetaData.ABI

// Tokensplitter is an auto generated Go binding around an Ethereum contract.
type Tokensplitter struct {
	TokensplitterCaller     // Read-only binding to the contract
	TokensplitterTransactor // Write-only binding to the contract
	TokensplitterFilterer   // Log filterer for contract events
}

// TokensplitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokensplitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokensplitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokensplitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokensplitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokensplitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokensplitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokensplitterSession struct {
	Contract     *Tokensplitter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokensplitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokensplitterCallerSession struct {
	Contract *TokensplitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokensplitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokensplitterTransactorSession struct {
	Contract     *TokensplitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokensplitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokensplitterRaw struct {
	Contract *Tokensplitter // Generic contract binding to access the raw methods on
}

// TokensplitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokensplitterCallerRaw struct {
	Contract *TokensplitterCaller // Generic read-only contract binding to access the raw methods on
}

// TokensplitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokensplitterTransactorRaw struct {
	Contract *TokensplitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokensplitter creates a new instance of Tokensplitter, bound to a specific deployed contract.
func NewTokensplitter(address common.Address, backend bind.ContractBackend) (*Tokensplitter, error) {
	contract, err := bindTokensplitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokensplitter{TokensplitterCaller: TokensplitterCaller{contract: contract}, TokensplitterTransactor: TokensplitterTransactor{contract: contract}, TokensplitterFilterer: TokensplitterFilterer{contract: contract}}, nil
}

// NewTokensplitterCaller creates a new read-only instance of Tokensplitter, bound to a specific deployed contract.
func NewTokensplitterCaller(address common.Address, caller bind.ContractCaller) (*TokensplitterCaller, error) {
	contract, err := bindTokensplitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokensplitterCaller{contract: contract}, nil
}

// NewTokensplitterTransactor creates a new write-only instance of Tokensplitter, bound to a specific deployed contract.
func NewTokensplitterTransactor(address common.Address, transactor bind.ContractTransactor) (*TokensplitterTransactor, error) {
	contract, err := bindTokensplitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokensplitterTransactor{contract: contract}, nil
}

// NewTokensplitterFilterer creates a new log filterer instance of Tokensplitter, bound to a specific deployed contract.
func NewTokensplitterFilterer(address common.Address, filterer bind.ContractFilterer) (*TokensplitterFilterer, error) {
	contract, err := bindTokensplitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokensplitterFilterer{contract: contract}, nil
}

// bindTokensplitter binds a generic wrapper to an already deployed contract.
func bindTokensplitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokensplitterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokensplitter *TokensplitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokensplitter.Contract.TokensplitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokensplitter *TokensplitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokensplitter.Contract.TokensplitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokensplitter *TokensplitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokensplitter.Contract.TokensplitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokensplitter *TokensplitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokensplitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokensplitter *TokensplitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokensplitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokensplitter *TokensplitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokensplitter.Contract.contract.Transact(opts, method, params...)
}

// TOTALSHARES is a free data retrieval call binding the contract method 0x85e3f997.
//
// Solidity: function TOTAL_SHARES() view returns(uint256)
func (_Tokensplitter *TokensplitterCaller) TOTALSHARES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokensplitter.contract.Call(opts, &out, "TOTAL_SHARES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALSHARES is a free data retrieval call binding the contract method 0x85e3f997.
//
// Solidity: function TOTAL_SHARES() view returns(uint256)
func (_Tokensplitter *TokensplitterSession) TOTALSHARES() (*big.Int, error) {
	return _Tokensplitter.Contract.TOTALSHARES(&_Tokensplitter.CallOpts)
}

// TOTALSHARES is a free data retrieval call binding the contract method 0x85e3f997.
//
// Solidity: function TOTAL_SHARES() view returns(uint256)
func (_Tokensplitter *TokensplitterCallerSession) TOTALSHARES() (*big.Int, error) {
	return _Tokensplitter.Contract.TOTALSHARES(&_Tokensplitter.CallOpts)
}

// AccountInfo is a free data retrieval call binding the contract method 0xa7310b58.
//
// Solidity: function accountInfo(address ) view returns(uint256 shares, uint256 tokensDistributedToAccount)
func (_Tokensplitter *TokensplitterCaller) AccountInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Shares                     *big.Int
	TokensDistributedToAccount *big.Int
}, error) {
	var out []interface{}
	err := _Tokensplitter.contract.Call(opts, &out, "accountInfo", arg0)

	outstruct := new(struct {
		Shares                     *big.Int
		TokensDistributedToAccount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Shares = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokensDistributedToAccount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AccountInfo is a free data retrieval call binding the contract method 0xa7310b58.
//
// Solidity: function accountInfo(address ) view returns(uint256 shares, uint256 tokensDistributedToAccount)
func (_Tokensplitter *TokensplitterSession) AccountInfo(arg0 common.Address) (struct {
	Shares                     *big.Int
	TokensDistributedToAccount *big.Int
}, error) {
	return _Tokensplitter.Contract.AccountInfo(&_Tokensplitter.CallOpts, arg0)
}

// AccountInfo is a free data retrieval call binding the contract method 0xa7310b58.
//
// Solidity: function accountInfo(address ) view returns(uint256 shares, uint256 tokensDistributedToAccount)
func (_Tokensplitter *TokensplitterCallerSession) AccountInfo(arg0 common.Address) (struct {
	Shares                     *big.Int
	TokensDistributedToAccount *big.Int
}, error) {
	return _Tokensplitter.Contract.AccountInfo(&_Tokensplitter.CallOpts, arg0)
}

// CalculatePendingRewards is a free data retrieval call binding the contract method 0x097aad10.
//
// Solidity: function calculatePendingRewards(address account) view returns(uint256)
func (_Tokensplitter *TokensplitterCaller) CalculatePendingRewards(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokensplitter.contract.Call(opts, &out, "calculatePendingRewards", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePendingRewards is a free data retrieval call binding the contract method 0x097aad10.
//
// Solidity: function calculatePendingRewards(address account) view returns(uint256)
func (_Tokensplitter *TokensplitterSession) CalculatePendingRewards(account common.Address) (*big.Int, error) {
	return _Tokensplitter.Contract.CalculatePendingRewards(&_Tokensplitter.CallOpts, account)
}

// CalculatePendingRewards is a free data retrieval call binding the contract method 0x097aad10.
//
// Solidity: function calculatePendingRewards(address account) view returns(uint256)
func (_Tokensplitter *TokensplitterCallerSession) CalculatePendingRewards(account common.Address) (*big.Int, error) {
	return _Tokensplitter.Contract.CalculatePendingRewards(&_Tokensplitter.CallOpts, account)
}

// HunterGovernanceToken is a free data retrieval call binding the contract method 0xf363264a.
//
// Solidity: function hunterGovernanceToken() view returns(address)
func (_Tokensplitter *TokensplitterCaller) HunterGovernanceToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokensplitter.contract.Call(opts, &out, "hunterGovernanceToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HunterGovernanceToken is a free data retrieval call binding the contract method 0xf363264a.
//
// Solidity: function hunterGovernanceToken() view returns(address)
func (_Tokensplitter *TokensplitterSession) HunterGovernanceToken() (common.Address, error) {
	return _Tokensplitter.Contract.HunterGovernanceToken(&_Tokensplitter.CallOpts)
}

// HunterGovernanceToken is a free data retrieval call binding the contract method 0xf363264a.
//
// Solidity: function hunterGovernanceToken() view returns(address)
func (_Tokensplitter *TokensplitterCallerSession) HunterGovernanceToken() (common.Address, error) {
	return _Tokensplitter.Contract.HunterGovernanceToken(&_Tokensplitter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokensplitter *TokensplitterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokensplitter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokensplitter *TokensplitterSession) Owner() (common.Address, error) {
	return _Tokensplitter.Contract.Owner(&_Tokensplitter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokensplitter *TokensplitterCallerSession) Owner() (common.Address, error) {
	return _Tokensplitter.Contract.Owner(&_Tokensplitter.CallOpts)
}

// TotalTokensDistributed is a free data retrieval call binding the contract method 0xd8f163ab.
//
// Solidity: function totalTokensDistributed() view returns(uint256)
func (_Tokensplitter *TokensplitterCaller) TotalTokensDistributed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokensplitter.contract.Call(opts, &out, "totalTokensDistributed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokensDistributed is a free data retrieval call binding the contract method 0xd8f163ab.
//
// Solidity: function totalTokensDistributed() view returns(uint256)
func (_Tokensplitter *TokensplitterSession) TotalTokensDistributed() (*big.Int, error) {
	return _Tokensplitter.Contract.TotalTokensDistributed(&_Tokensplitter.CallOpts)
}

// TotalTokensDistributed is a free data retrieval call binding the contract method 0xd8f163ab.
//
// Solidity: function totalTokensDistributed() view returns(uint256)
func (_Tokensplitter *TokensplitterCallerSession) TotalTokensDistributed() (*big.Int, error) {
	return _Tokensplitter.Contract.TotalTokensDistributed(&_Tokensplitter.CallOpts)
}

// ReleaseTokens is a paid mutator transaction binding the contract method 0x87b0be48.
//
// Solidity: function releaseTokens(address account) returns()
func (_Tokensplitter *TokensplitterTransactor) ReleaseTokens(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Tokensplitter.contract.Transact(opts, "releaseTokens", account)
}

// ReleaseTokens is a paid mutator transaction binding the contract method 0x87b0be48.
//
// Solidity: function releaseTokens(address account) returns()
func (_Tokensplitter *TokensplitterSession) ReleaseTokens(account common.Address) (*types.Transaction, error) {
	return _Tokensplitter.Contract.ReleaseTokens(&_Tokensplitter.TransactOpts, account)
}

// ReleaseTokens is a paid mutator transaction binding the contract method 0x87b0be48.
//
// Solidity: function releaseTokens(address account) returns()
func (_Tokensplitter *TokensplitterTransactorSession) ReleaseTokens(account common.Address) (*types.Transaction, error) {
	return _Tokensplitter.Contract.ReleaseTokens(&_Tokensplitter.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokensplitter *TokensplitterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokensplitter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokensplitter *TokensplitterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokensplitter.Contract.RenounceOwnership(&_Tokensplitter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokensplitter *TokensplitterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokensplitter.Contract.RenounceOwnership(&_Tokensplitter.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokensplitter *TokensplitterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tokensplitter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokensplitter *TokensplitterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokensplitter.Contract.TransferOwnership(&_Tokensplitter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokensplitter *TokensplitterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokensplitter.Contract.TransferOwnership(&_Tokensplitter.TransactOpts, newOwner)
}

// UpdateSharesOwner is a paid mutator transaction binding the contract method 0x19281afd.
//
// Solidity: function updateSharesOwner(address _newRecipient, address _currentRecipient) returns()
func (_Tokensplitter *TokensplitterTransactor) UpdateSharesOwner(opts *bind.TransactOpts, _newRecipient common.Address, _currentRecipient common.Address) (*types.Transaction, error) {
	return _Tokensplitter.contract.Transact(opts, "updateSharesOwner", _newRecipient, _currentRecipient)
}

// UpdateSharesOwner is a paid mutator transaction binding the contract method 0x19281afd.
//
// Solidity: function updateSharesOwner(address _newRecipient, address _currentRecipient) returns()
func (_Tokensplitter *TokensplitterSession) UpdateSharesOwner(_newRecipient common.Address, _currentRecipient common.Address) (*types.Transaction, error) {
	return _Tokensplitter.Contract.UpdateSharesOwner(&_Tokensplitter.TransactOpts, _newRecipient, _currentRecipient)
}

// UpdateSharesOwner is a paid mutator transaction binding the contract method 0x19281afd.
//
// Solidity: function updateSharesOwner(address _newRecipient, address _currentRecipient) returns()
func (_Tokensplitter *TokensplitterTransactorSession) UpdateSharesOwner(_newRecipient common.Address, _currentRecipient common.Address) (*types.Transaction, error) {
	return _Tokensplitter.Contract.UpdateSharesOwner(&_Tokensplitter.TransactOpts, _newRecipient, _currentRecipient)
}

// TokensplitterNewSharesOwnerIterator is returned from FilterNewSharesOwner and is used to iterate over the raw logs and unpacked data for NewSharesOwner events raised by the Tokensplitter contract.
type TokensplitterNewSharesOwnerIterator struct {
	Event *TokensplitterNewSharesOwner // Event containing the contract specifics and raw log

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
func (it *TokensplitterNewSharesOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokensplitterNewSharesOwner)
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
		it.Event = new(TokensplitterNewSharesOwner)
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
func (it *TokensplitterNewSharesOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokensplitterNewSharesOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokensplitterNewSharesOwner represents a NewSharesOwner event raised by the Tokensplitter contract.
type TokensplitterNewSharesOwner struct {
	OldRecipient common.Address
	NewRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewSharesOwner is a free log retrieval operation binding the contract event 0x47eba4993da31bfed0680c802fb8577f1a9d57bb4c0881372fbe7ad29995a880.
//
// Solidity: event NewSharesOwner(address indexed oldRecipient, address indexed newRecipient)
func (_Tokensplitter *TokensplitterFilterer) FilterNewSharesOwner(opts *bind.FilterOpts, oldRecipient []common.Address, newRecipient []common.Address) (*TokensplitterNewSharesOwnerIterator, error) {

	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}
	var newRecipientRule []interface{}
	for _, newRecipientItem := range newRecipient {
		newRecipientRule = append(newRecipientRule, newRecipientItem)
	}

	logs, sub, err := _Tokensplitter.contract.FilterLogs(opts, "NewSharesOwner", oldRecipientRule, newRecipientRule)
	if err != nil {
		return nil, err
	}
	return &TokensplitterNewSharesOwnerIterator{contract: _Tokensplitter.contract, event: "NewSharesOwner", logs: logs, sub: sub}, nil
}

// WatchNewSharesOwner is a free log subscription operation binding the contract event 0x47eba4993da31bfed0680c802fb8577f1a9d57bb4c0881372fbe7ad29995a880.
//
// Solidity: event NewSharesOwner(address indexed oldRecipient, address indexed newRecipient)
func (_Tokensplitter *TokensplitterFilterer) WatchNewSharesOwner(opts *bind.WatchOpts, sink chan<- *TokensplitterNewSharesOwner, oldRecipient []common.Address, newRecipient []common.Address) (event.Subscription, error) {

	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}
	var newRecipientRule []interface{}
	for _, newRecipientItem := range newRecipient {
		newRecipientRule = append(newRecipientRule, newRecipientItem)
	}

	logs, sub, err := _Tokensplitter.contract.WatchLogs(opts, "NewSharesOwner", oldRecipientRule, newRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokensplitterNewSharesOwner)
				if err := _Tokensplitter.contract.UnpackLog(event, "NewSharesOwner", log); err != nil {
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

// ParseNewSharesOwner is a log parse operation binding the contract event 0x47eba4993da31bfed0680c802fb8577f1a9d57bb4c0881372fbe7ad29995a880.
//
// Solidity: event NewSharesOwner(address indexed oldRecipient, address indexed newRecipient)
func (_Tokensplitter *TokensplitterFilterer) ParseNewSharesOwner(log types.Log) (*TokensplitterNewSharesOwner, error) {
	event := new(TokensplitterNewSharesOwner)
	if err := _Tokensplitter.contract.UnpackLog(event, "NewSharesOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokensplitterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tokensplitter contract.
type TokensplitterOwnershipTransferredIterator struct {
	Event *TokensplitterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokensplitterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokensplitterOwnershipTransferred)
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
		it.Event = new(TokensplitterOwnershipTransferred)
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
func (it *TokensplitterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokensplitterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokensplitterOwnershipTransferred represents a OwnershipTransferred event raised by the Tokensplitter contract.
type TokensplitterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokensplitter *TokensplitterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokensplitterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokensplitter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokensplitterOwnershipTransferredIterator{contract: _Tokensplitter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokensplitter *TokensplitterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokensplitterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokensplitter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokensplitterOwnershipTransferred)
				if err := _Tokensplitter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Tokensplitter *TokensplitterFilterer) ParseOwnershipTransferred(log types.Log) (*TokensplitterOwnershipTransferred, error) {
	event := new(TokensplitterOwnershipTransferred)
	if err := _Tokensplitter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokensplitterTokensTransferredIterator is returned from FilterTokensTransferred and is used to iterate over the raw logs and unpacked data for TokensTransferred events raised by the Tokensplitter contract.
type TokensplitterTokensTransferredIterator struct {
	Event *TokensplitterTokensTransferred // Event containing the contract specifics and raw log

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
func (it *TokensplitterTokensTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokensplitterTokensTransferred)
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
		it.Event = new(TokensplitterTokensTransferred)
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
func (it *TokensplitterTokensTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokensplitterTokensTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokensplitterTokensTransferred represents a TokensTransferred event raised by the Tokensplitter contract.
type TokensplitterTokensTransferred struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokensTransferred is a free log retrieval operation binding the contract event 0x12f4533b5cbd2c9f8a0752a2d0b16379af992dbb2a0844a5007a19d983b3a934.
//
// Solidity: event TokensTransferred(address indexed account, uint256 amount)
func (_Tokensplitter *TokensplitterFilterer) FilterTokensTransferred(opts *bind.FilterOpts, account []common.Address) (*TokensplitterTokensTransferredIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Tokensplitter.contract.FilterLogs(opts, "TokensTransferred", accountRule)
	if err != nil {
		return nil, err
	}
	return &TokensplitterTokensTransferredIterator{contract: _Tokensplitter.contract, event: "TokensTransferred", logs: logs, sub: sub}, nil
}

// WatchTokensTransferred is a free log subscription operation binding the contract event 0x12f4533b5cbd2c9f8a0752a2d0b16379af992dbb2a0844a5007a19d983b3a934.
//
// Solidity: event TokensTransferred(address indexed account, uint256 amount)
func (_Tokensplitter *TokensplitterFilterer) WatchTokensTransferred(opts *bind.WatchOpts, sink chan<- *TokensplitterTokensTransferred, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Tokensplitter.contract.WatchLogs(opts, "TokensTransferred", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokensplitterTokensTransferred)
				if err := _Tokensplitter.contract.UnpackLog(event, "TokensTransferred", log); err != nil {
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

// ParseTokensTransferred is a log parse operation binding the contract event 0x12f4533b5cbd2c9f8a0752a2d0b16379af992dbb2a0844a5007a19d983b3a934.
//
// Solidity: event TokensTransferred(address indexed account, uint256 amount)
func (_Tokensplitter *TokensplitterFilterer) ParseTokensTransferred(log types.Log) (*TokensplitterTokensTransferred, error) {
	event := new(TokensplitterTokensTransferred)
	if err := _Tokensplitter.contract.UnpackLog(event, "TokensTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
