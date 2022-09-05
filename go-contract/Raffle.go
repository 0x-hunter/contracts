// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usd\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_uniV3Factory\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_currencies\",\"type\":\"address[]\"},{\"internalType\":\"uint24[]\",\"name\":\"_feeRates\",\"type\":\"uint24[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundsAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FundsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"PrizeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"RaffleCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"RaffleClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"RaffleCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RaffleEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winningTicket\",\"type\":\"uint256\"}],\"name\":\"WinningTicketPicked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_raffleId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_raffleId\",\"type\":\"uint256\"}],\"name\":\"claimFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_raffleId\",\"type\":\"uint256\"}],\"name\":\"claimPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_raffleId\",\"type\":\"uint256\"}],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_ticketPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_ticketSize\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currencyInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"feeRate\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTaskId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_raffleId\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"_amount\",\"type\":\"uint96\"}],\"name\":\"enter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"}],\"name\":\"getCurrencyPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prizeInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"winningNumber\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"prizeClaimed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"fundsClaimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"raffleInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"ticketSize\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ticketPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint160\",\"name\":\"funds\",\"type\":\"uint160\"},{\"internalType\":\"uint96\",\"name\":\"soldTickets\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"enumRaffle.TaskStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"_feeRate\",\"type\":\"uint24\"}],\"name\":\"setCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ticketAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// CurrencyInfo is a free data retrieval call binding the contract method 0x6d31cb2a.
//
// Solidity: function currencyInfo(address ) view returns(bool enabled, uint24 feeRate)
func (_Contract *ContractCaller) CurrencyInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Enabled bool
	FeeRate *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currencyInfo", arg0)

	outstruct := new(struct {
		Enabled bool
		FeeRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Enabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.FeeRate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrencyInfo is a free data retrieval call binding the contract method 0x6d31cb2a.
//
// Solidity: function currencyInfo(address ) view returns(bool enabled, uint24 feeRate)
func (_Contract *ContractSession) CurrencyInfo(arg0 common.Address) (struct {
	Enabled bool
	FeeRate *big.Int
}, error) {
	return _Contract.Contract.CurrencyInfo(&_Contract.CallOpts, arg0)
}

// CurrencyInfo is a free data retrieval call binding the contract method 0x6d31cb2a.
//
// Solidity: function currencyInfo(address ) view returns(bool enabled, uint24 feeRate)
func (_Contract *ContractCallerSession) CurrencyInfo(arg0 common.Address) (struct {
	Enabled bool
	FeeRate *big.Int
}, error) {
	return _Contract.Contract.CurrencyInfo(&_Contract.CallOpts, arg0)
}

// CurrentTaskId is a free data retrieval call binding the contract method 0x98934a1f.
//
// Solidity: function currentTaskId() view returns(uint256)
func (_Contract *ContractCaller) CurrentTaskId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currentTaskId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTaskId is a free data retrieval call binding the contract method 0x98934a1f.
//
// Solidity: function currentTaskId() view returns(uint256)
func (_Contract *ContractSession) CurrentTaskId() (*big.Int, error) {
	return _Contract.Contract.CurrentTaskId(&_Contract.CallOpts)
}

// CurrentTaskId is a free data retrieval call binding the contract method 0x98934a1f.
//
// Solidity: function currentTaskId() view returns(uint256)
func (_Contract *ContractCallerSession) CurrentTaskId() (*big.Int, error) {
	return _Contract.Contract.CurrentTaskId(&_Contract.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Contract *ContractCaller) FeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Contract *ContractSession) FeeRate() (*big.Int, error) {
	return _Contract.Contract.FeeRate(&_Contract.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Contract *ContractCallerSession) FeeRate() (*big.Int, error) {
	return _Contract.Contract.FeeRate(&_Contract.CallOpts)
}

// FeeSetter is a free data retrieval call binding the contract method 0x87cf3ef4.
//
// Solidity: function feeSetter() view returns(address)
func (_Contract *ContractCaller) FeeSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "feeSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeSetter is a free data retrieval call binding the contract method 0x87cf3ef4.
//
// Solidity: function feeSetter() view returns(address)
func (_Contract *ContractSession) FeeSetter() (common.Address, error) {
	return _Contract.Contract.FeeSetter(&_Contract.CallOpts)
}

// FeeSetter is a free data retrieval call binding the contract method 0x87cf3ef4.
//
// Solidity: function feeSetter() view returns(address)
func (_Contract *ContractCallerSession) FeeSetter() (common.Address, error) {
	return _Contract.Contract.FeeSetter(&_Contract.CallOpts)
}

// GetCurrencyPrice is a free data retrieval call binding the contract method 0x7058346f.
//
// Solidity: function getCurrencyPrice(address _currency) view returns(uint256 price)
func (_Contract *ContractCaller) GetCurrencyPrice(opts *bind.CallOpts, _currency common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCurrencyPrice", _currency)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrencyPrice is a free data retrieval call binding the contract method 0x7058346f.
//
// Solidity: function getCurrencyPrice(address _currency) view returns(uint256 price)
func (_Contract *ContractSession) GetCurrencyPrice(_currency common.Address) (*big.Int, error) {
	return _Contract.Contract.GetCurrencyPrice(&_Contract.CallOpts, _currency)
}

// GetCurrencyPrice is a free data retrieval call binding the contract method 0x7058346f.
//
// Solidity: function getCurrencyPrice(address _currency) view returns(uint256 price)
func (_Contract *ContractCallerSession) GetCurrencyPrice(_currency common.Address) (*big.Int, error) {
	return _Contract.Contract.GetCurrencyPrice(&_Contract.CallOpts, _currency)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Contract *ContractCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Contract *ContractSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Contract.Contract.OnERC721Received(&_Contract.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Contract *ContractCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Contract.Contract.OnERC721Received(&_Contract.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// PrizeInfo is a free data retrieval call binding the contract method 0x726a41cf.
//
// Solidity: function prizeInfo(uint256 ) view returns(uint32 winningNumber, bool prizeClaimed, bool fundsClaimed)
func (_Contract *ContractCaller) PrizeInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	WinningNumber uint32
	PrizeClaimed  bool
	FundsClaimed  bool
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "prizeInfo", arg0)

	outstruct := new(struct {
		WinningNumber uint32
		PrizeClaimed  bool
		FundsClaimed  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WinningNumber = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.PrizeClaimed = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.FundsClaimed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// PrizeInfo is a free data retrieval call binding the contract method 0x726a41cf.
//
// Solidity: function prizeInfo(uint256 ) view returns(uint32 winningNumber, bool prizeClaimed, bool fundsClaimed)
func (_Contract *ContractSession) PrizeInfo(arg0 *big.Int) (struct {
	WinningNumber uint32
	PrizeClaimed  bool
	FundsClaimed  bool
}, error) {
	return _Contract.Contract.PrizeInfo(&_Contract.CallOpts, arg0)
}

// PrizeInfo is a free data retrieval call binding the contract method 0x726a41cf.
//
// Solidity: function prizeInfo(uint256 ) view returns(uint32 winningNumber, bool prizeClaimed, bool fundsClaimed)
func (_Contract *ContractCallerSession) PrizeInfo(arg0 *big.Int) (struct {
	WinningNumber uint32
	PrizeClaimed  bool
	FundsClaimed  bool
}, error) {
	return _Contract.Contract.PrizeInfo(&_Contract.CallOpts, arg0)
}

// RaffleInfo is a free data retrieval call binding the contract method 0xbf30d790.
//
// Solidity: function raffleInfo(uint256 ) view returns(address owner, uint64 endTime, uint32 ticketSize, uint256 tokenId, address tokenAddress, uint64 ticketPrice, uint160 funds, uint96 soldTickets, address currency, uint8 status)
func (_Contract *ContractCaller) RaffleInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner        common.Address
	EndTime      uint64
	TicketSize   uint32
	TokenId      *big.Int
	TokenAddress common.Address
	TicketPrice  uint64
	Funds        *big.Int
	SoldTickets  *big.Int
	Currency     common.Address
	Status       uint8
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "raffleInfo", arg0)

	outstruct := new(struct {
		Owner        common.Address
		EndTime      uint64
		TicketSize   uint32
		TokenId      *big.Int
		TokenAddress common.Address
		TicketPrice  uint64
		Funds        *big.Int
		SoldTickets  *big.Int
		Currency     common.Address
		Status       uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.EndTime = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.TicketSize = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.TokenId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TokenAddress = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.TicketPrice = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.Funds = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.SoldTickets = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Currency = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[9], new(uint8)).(*uint8)

	return *outstruct, err

}

// RaffleInfo is a free data retrieval call binding the contract method 0xbf30d790.
//
// Solidity: function raffleInfo(uint256 ) view returns(address owner, uint64 endTime, uint32 ticketSize, uint256 tokenId, address tokenAddress, uint64 ticketPrice, uint160 funds, uint96 soldTickets, address currency, uint8 status)
func (_Contract *ContractSession) RaffleInfo(arg0 *big.Int) (struct {
	Owner        common.Address
	EndTime      uint64
	TicketSize   uint32
	TokenId      *big.Int
	TokenAddress common.Address
	TicketPrice  uint64
	Funds        *big.Int
	SoldTickets  *big.Int
	Currency     common.Address
	Status       uint8
}, error) {
	return _Contract.Contract.RaffleInfo(&_Contract.CallOpts, arg0)
}

// RaffleInfo is a free data retrieval call binding the contract method 0xbf30d790.
//
// Solidity: function raffleInfo(uint256 ) view returns(address owner, uint64 endTime, uint32 ticketSize, uint256 tokenId, address tokenAddress, uint64 ticketPrice, uint160 funds, uint96 soldTickets, address currency, uint8 status)
func (_Contract *ContractCallerSession) RaffleInfo(arg0 *big.Int) (struct {
	Owner        common.Address
	EndTime      uint64
	TicketSize   uint32
	TokenId      *big.Int
	TokenAddress common.Address
	TicketPrice  uint64
	Funds        *big.Int
	SoldTickets  *big.Int
	Currency     common.Address
	Status       uint8
}, error) {
	return _Contract.Contract.RaffleInfo(&_Contract.CallOpts, arg0)
}

// RequestInfo is a free data retrieval call binding the contract method 0x4f588bf1.
//
// Solidity: function requestInfo(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) RequestInfo(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "requestInfo", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestInfo is a free data retrieval call binding the contract method 0x4f588bf1.
//
// Solidity: function requestInfo(uint256 ) view returns(uint256)
func (_Contract *ContractSession) RequestInfo(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.RequestInfo(&_Contract.CallOpts, arg0)
}

// RequestInfo is a free data retrieval call binding the contract method 0x4f588bf1.
//
// Solidity: function requestInfo(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) RequestInfo(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.RequestInfo(&_Contract.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 ticketAmount)
func (_Contract *ContractCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "userInfo", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 ticketAmount)
func (_Contract *ContractSession) UserInfo(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Contract.Contract.UserInfo(&_Contract.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 ticketAmount)
func (_Contract *ContractCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Contract.Contract.UserInfo(&_Contract.CallOpts, arg0, arg1)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _raffleId) returns()
func (_Contract *ContractTransactor) Cancel(opts *bind.TransactOpts, _raffleId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "cancel", _raffleId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _raffleId) returns()
func (_Contract *ContractSession) Cancel(_raffleId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Cancel(&_Contract.TransactOpts, _raffleId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _raffleId) returns()
func (_Contract *ContractTransactorSession) Cancel(_raffleId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Cancel(&_Contract.TransactOpts, _raffleId)
}

// ClaimFunds is a paid mutator transaction binding the contract method 0x1b55e338.
//
// Solidity: function claimFunds(uint256 _raffleId) returns()
func (_Contract *ContractTransactor) ClaimFunds(opts *bind.TransactOpts, _raffleId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimFunds", _raffleId)
}

// ClaimFunds is a paid mutator transaction binding the contract method 0x1b55e338.
//
// Solidity: function claimFunds(uint256 _raffleId) returns()
func (_Contract *ContractSession) ClaimFunds(_raffleId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimFunds(&_Contract.TransactOpts, _raffleId)
}

// ClaimFunds is a paid mutator transaction binding the contract method 0x1b55e338.
//
// Solidity: function claimFunds(uint256 _raffleId) returns()
func (_Contract *ContractTransactorSession) ClaimFunds(_raffleId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimFunds(&_Contract.TransactOpts, _raffleId)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0xd7098154.
//
// Solidity: function claimPrize(uint256 _raffleId) returns()
func (_Contract *ContractTransactor) ClaimPrize(opts *bind.TransactOpts, _raffleId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimPrize", _raffleId)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0xd7098154.
//
// Solidity: function claimPrize(uint256 _raffleId) returns()
func (_Contract *ContractSession) ClaimPrize(_raffleId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimPrize(&_Contract.TransactOpts, _raffleId)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0xd7098154.
//
// Solidity: function claimPrize(uint256 _raffleId) returns()
func (_Contract *ContractTransactorSession) ClaimPrize(_raffleId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimPrize(&_Contract.TransactOpts, _raffleId)
}

// Close is a paid mutator transaction binding the contract method 0x0aebeb4e.
//
// Solidity: function close(uint256 _raffleId) returns()
func (_Contract *ContractTransactor) Close(opts *bind.TransactOpts, _raffleId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "close", _raffleId)
}

// Close is a paid mutator transaction binding the contract method 0x0aebeb4e.
//
// Solidity: function close(uint256 _raffleId) returns()
func (_Contract *ContractSession) Close(_raffleId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Close(&_Contract.TransactOpts, _raffleId)
}

// Close is a paid mutator transaction binding the contract method 0x0aebeb4e.
//
// Solidity: function close(uint256 _raffleId) returns()
func (_Contract *ContractTransactorSession) Close(_raffleId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Close(&_Contract.TransactOpts, _raffleId)
}

// Create is a paid mutator transaction binding the contract method 0xa7c1891e.
//
// Solidity: function create(address _tokenAddr, uint256 _tokenId, uint256 _duration, uint64 _ticketPrice, uint32 _ticketSize, address _currency) returns()
func (_Contract *ContractTransactor) Create(opts *bind.TransactOpts, _tokenAddr common.Address, _tokenId *big.Int, _duration *big.Int, _ticketPrice uint64, _ticketSize uint32, _currency common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "create", _tokenAddr, _tokenId, _duration, _ticketPrice, _ticketSize, _currency)
}

// Create is a paid mutator transaction binding the contract method 0xa7c1891e.
//
// Solidity: function create(address _tokenAddr, uint256 _tokenId, uint256 _duration, uint64 _ticketPrice, uint32 _ticketSize, address _currency) returns()
func (_Contract *ContractSession) Create(_tokenAddr common.Address, _tokenId *big.Int, _duration *big.Int, _ticketPrice uint64, _ticketSize uint32, _currency common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Create(&_Contract.TransactOpts, _tokenAddr, _tokenId, _duration, _ticketPrice, _ticketSize, _currency)
}

// Create is a paid mutator transaction binding the contract method 0xa7c1891e.
//
// Solidity: function create(address _tokenAddr, uint256 _tokenId, uint256 _duration, uint64 _ticketPrice, uint32 _ticketSize, address _currency) returns()
func (_Contract *ContractTransactorSession) Create(_tokenAddr common.Address, _tokenId *big.Int, _duration *big.Int, _ticketPrice uint64, _ticketSize uint32, _currency common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Create(&_Contract.TransactOpts, _tokenAddr, _tokenId, _duration, _ticketPrice, _ticketSize, _currency)
}

// Enter is a paid mutator transaction binding the contract method 0x00d407d6.
//
// Solidity: function enter(uint256 _raffleId, uint96 _amount) returns()
func (_Contract *ContractTransactor) Enter(opts *bind.TransactOpts, _raffleId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "enter", _raffleId, _amount)
}

// Enter is a paid mutator transaction binding the contract method 0x00d407d6.
//
// Solidity: function enter(uint256 _raffleId, uint96 _amount) returns()
func (_Contract *ContractSession) Enter(_raffleId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Enter(&_Contract.TransactOpts, _raffleId, _amount)
}

// Enter is a paid mutator transaction binding the contract method 0x00d407d6.
//
// Solidity: function enter(uint256 _raffleId, uint96 _amount) returns()
func (_Contract *ContractTransactorSession) Enter(_raffleId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Enter(&_Contract.TransactOpts, _raffleId, _amount)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Contract *ContractTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Contract *ContractSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RawFulfillRandomWords(&_Contract.TransactOpts, requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Contract *ContractTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RawFulfillRandomWords(&_Contract.TransactOpts, requestId, randomWords)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SetCurrency is a paid mutator transaction binding the contract method 0x59991a77.
//
// Solidity: function setCurrency(address _currency, bool _enable, uint24 _feeRate) returns()
func (_Contract *ContractTransactor) SetCurrency(opts *bind.TransactOpts, _currency common.Address, _enable bool, _feeRate *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setCurrency", _currency, _enable, _feeRate)
}

// SetCurrency is a paid mutator transaction binding the contract method 0x59991a77.
//
// Solidity: function setCurrency(address _currency, bool _enable, uint24 _feeRate) returns()
func (_Contract *ContractSession) SetCurrency(_currency common.Address, _enable bool, _feeRate *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetCurrency(&_Contract.TransactOpts, _currency, _enable, _feeRate)
}

// SetCurrency is a paid mutator transaction binding the contract method 0x59991a77.
//
// Solidity: function setCurrency(address _currency, bool _enable, uint24 _feeRate) returns()
func (_Contract *ContractTransactorSession) SetCurrency(_currency common.Address, _enable bool, _feeRate *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetCurrency(&_Contract.TransactOpts, _currency, _enable, _feeRate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// ContractFundsClaimedIterator is returned from FilterFundsClaimed and is used to iterate over the raw logs and unpacked data for FundsClaimed events raised by the Contract contract.
type ContractFundsClaimedIterator struct {
	Event *ContractFundsClaimed // Event containing the contract specifics and raw log

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
func (it *ContractFundsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFundsClaimed)
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
		it.Event = new(ContractFundsClaimed)
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
func (it *ContractFundsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFundsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFundsClaimed represents a FundsClaimed event raised by the Contract contract.
type ContractFundsClaimed struct {
	TaskId      *big.Int
	FundsAmount *big.Int
	Owner       common.Address
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFundsClaimed is a free log retrieval operation binding the contract event 0x3f53f5f381ee0b1fbf730fa56efb0a2aec8a882b2847a1ec0279982d8846e4f2.
//
// Solidity: event FundsClaimed(uint256 taskId, uint256 fundsAmount, address owner, uint256 fee)
func (_Contract *ContractFilterer) FilterFundsClaimed(opts *bind.FilterOpts) (*ContractFundsClaimedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FundsClaimed")
	if err != nil {
		return nil, err
	}
	return &ContractFundsClaimedIterator{contract: _Contract.contract, event: "FundsClaimed", logs: logs, sub: sub}, nil
}

// WatchFundsClaimed is a free log subscription operation binding the contract event 0x3f53f5f381ee0b1fbf730fa56efb0a2aec8a882b2847a1ec0279982d8846e4f2.
//
// Solidity: event FundsClaimed(uint256 taskId, uint256 fundsAmount, address owner, uint256 fee)
func (_Contract *ContractFilterer) WatchFundsClaimed(opts *bind.WatchOpts, sink chan<- *ContractFundsClaimed) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FundsClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFundsClaimed)
				if err := _Contract.contract.UnpackLog(event, "FundsClaimed", log); err != nil {
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

// ParseFundsClaimed is a log parse operation binding the contract event 0x3f53f5f381ee0b1fbf730fa56efb0a2aec8a882b2847a1ec0279982d8846e4f2.
//
// Solidity: event FundsClaimed(uint256 taskId, uint256 fundsAmount, address owner, uint256 fee)
func (_Contract *ContractFilterer) ParseFundsClaimed(log types.Log) (*ContractFundsClaimed, error) {
	event := new(ContractFundsClaimed)
	if err := _Contract.contract.UnpackLog(event, "FundsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPrizeClaimedIterator is returned from FilterPrizeClaimed and is used to iterate over the raw logs and unpacked data for PrizeClaimed events raised by the Contract contract.
type ContractPrizeClaimedIterator struct {
	Event *ContractPrizeClaimed // Event containing the contract specifics and raw log

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
func (it *ContractPrizeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPrizeClaimed)
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
		it.Event = new(ContractPrizeClaimed)
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
func (it *ContractPrizeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPrizeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPrizeClaimed represents a PrizeClaimed event raised by the Contract contract.
type ContractPrizeClaimed struct {
	TaskId  *big.Int
	TokenId *big.Int
	Winner  common.Address
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPrizeClaimed is a free log retrieval operation binding the contract event 0x608c8c3c6a3d40f40b96fbee44f9b9aee57a39c2fcb8bb0f59d8892b2d5492a2.
//
// Solidity: event PrizeClaimed(uint256 taskId, uint256 tokenId, address winner, uint256 fee)
func (_Contract *ContractFilterer) FilterPrizeClaimed(opts *bind.FilterOpts) (*ContractPrizeClaimedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PrizeClaimed")
	if err != nil {
		return nil, err
	}
	return &ContractPrizeClaimedIterator{contract: _Contract.contract, event: "PrizeClaimed", logs: logs, sub: sub}, nil
}

// WatchPrizeClaimed is a free log subscription operation binding the contract event 0x608c8c3c6a3d40f40b96fbee44f9b9aee57a39c2fcb8bb0f59d8892b2d5492a2.
//
// Solidity: event PrizeClaimed(uint256 taskId, uint256 tokenId, address winner, uint256 fee)
func (_Contract *ContractFilterer) WatchPrizeClaimed(opts *bind.WatchOpts, sink chan<- *ContractPrizeClaimed) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PrizeClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPrizeClaimed)
				if err := _Contract.contract.UnpackLog(event, "PrizeClaimed", log); err != nil {
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

// ParsePrizeClaimed is a log parse operation binding the contract event 0x608c8c3c6a3d40f40b96fbee44f9b9aee57a39c2fcb8bb0f59d8892b2d5492a2.
//
// Solidity: event PrizeClaimed(uint256 taskId, uint256 tokenId, address winner, uint256 fee)
func (_Contract *ContractFilterer) ParsePrizeClaimed(log types.Log) (*ContractPrizeClaimed, error) {
	event := new(ContractPrizeClaimed)
	if err := _Contract.contract.UnpackLog(event, "PrizeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRaffleCancelledIterator is returned from FilterRaffleCancelled and is used to iterate over the raw logs and unpacked data for RaffleCancelled events raised by the Contract contract.
type ContractRaffleCancelledIterator struct {
	Event *ContractRaffleCancelled // Event containing the contract specifics and raw log

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
func (it *ContractRaffleCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRaffleCancelled)
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
		it.Event = new(ContractRaffleCancelled)
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
func (it *ContractRaffleCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRaffleCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRaffleCancelled represents a RaffleCancelled event raised by the Contract contract.
type ContractRaffleCancelled struct {
	TaskId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaffleCancelled is a free log retrieval operation binding the contract event 0x7fa78bab5d5570f162c69234a3f4d1b9294fd5bca72219cb16c1fe38839ccb8a.
//
// Solidity: event RaffleCancelled(uint256 taskId)
func (_Contract *ContractFilterer) FilterRaffleCancelled(opts *bind.FilterOpts) (*ContractRaffleCancelledIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RaffleCancelled")
	if err != nil {
		return nil, err
	}
	return &ContractRaffleCancelledIterator{contract: _Contract.contract, event: "RaffleCancelled", logs: logs, sub: sub}, nil
}

// WatchRaffleCancelled is a free log subscription operation binding the contract event 0x7fa78bab5d5570f162c69234a3f4d1b9294fd5bca72219cb16c1fe38839ccb8a.
//
// Solidity: event RaffleCancelled(uint256 taskId)
func (_Contract *ContractFilterer) WatchRaffleCancelled(opts *bind.WatchOpts, sink chan<- *ContractRaffleCancelled) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RaffleCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRaffleCancelled)
				if err := _Contract.contract.UnpackLog(event, "RaffleCancelled", log); err != nil {
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

// ParseRaffleCancelled is a log parse operation binding the contract event 0x7fa78bab5d5570f162c69234a3f4d1b9294fd5bca72219cb16c1fe38839ccb8a.
//
// Solidity: event RaffleCancelled(uint256 taskId)
func (_Contract *ContractFilterer) ParseRaffleCancelled(log types.Log) (*ContractRaffleCancelled, error) {
	event := new(ContractRaffleCancelled)
	if err := _Contract.contract.UnpackLog(event, "RaffleCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRaffleClosedIterator is returned from FilterRaffleClosed and is used to iterate over the raw logs and unpacked data for RaffleClosed events raised by the Contract contract.
type ContractRaffleClosedIterator struct {
	Event *ContractRaffleClosed // Event containing the contract specifics and raw log

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
func (it *ContractRaffleClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRaffleClosed)
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
		it.Event = new(ContractRaffleClosed)
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
func (it *ContractRaffleClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRaffleClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRaffleClosed represents a RaffleClosed event raised by the Contract contract.
type ContractRaffleClosed struct {
	TaskId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaffleClosed is a free log retrieval operation binding the contract event 0x2a7d48246875b5f17dc9944def25fc8fe90e524ca2165e290ffdc0e44a0bf8b6.
//
// Solidity: event RaffleClosed(uint256 taskId)
func (_Contract *ContractFilterer) FilterRaffleClosed(opts *bind.FilterOpts) (*ContractRaffleClosedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RaffleClosed")
	if err != nil {
		return nil, err
	}
	return &ContractRaffleClosedIterator{contract: _Contract.contract, event: "RaffleClosed", logs: logs, sub: sub}, nil
}

// WatchRaffleClosed is a free log subscription operation binding the contract event 0x2a7d48246875b5f17dc9944def25fc8fe90e524ca2165e290ffdc0e44a0bf8b6.
//
// Solidity: event RaffleClosed(uint256 taskId)
func (_Contract *ContractFilterer) WatchRaffleClosed(opts *bind.WatchOpts, sink chan<- *ContractRaffleClosed) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RaffleClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRaffleClosed)
				if err := _Contract.contract.UnpackLog(event, "RaffleClosed", log); err != nil {
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

// ParseRaffleClosed is a log parse operation binding the contract event 0x2a7d48246875b5f17dc9944def25fc8fe90e524ca2165e290ffdc0e44a0bf8b6.
//
// Solidity: event RaffleClosed(uint256 taskId)
func (_Contract *ContractFilterer) ParseRaffleClosed(log types.Log) (*ContractRaffleClosed, error) {
	event := new(ContractRaffleClosed)
	if err := _Contract.contract.UnpackLog(event, "RaffleClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRaffleCreatedIterator is returned from FilterRaffleCreated and is used to iterate over the raw logs and unpacked data for RaffleCreated events raised by the Contract contract.
type ContractRaffleCreatedIterator struct {
	Event *ContractRaffleCreated // Event containing the contract specifics and raw log

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
func (it *ContractRaffleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRaffleCreated)
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
		it.Event = new(ContractRaffleCreated)
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
func (it *ContractRaffleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRaffleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRaffleCreated represents a RaffleCreated event raised by the Contract contract.
type ContractRaffleCreated struct {
	TaskId    *big.Int
	Owner     common.Address
	StartTime *big.Int
	Duration  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRaffleCreated is a free log retrieval operation binding the contract event 0x71f77ac840b74f30ac7115cc2fca45c4157ccc2b0a6e72f4183c587d77baa0ab.
//
// Solidity: event RaffleCreated(uint256 indexed taskId, address indexed owner, uint256 startTime, uint256 duration)
func (_Contract *ContractFilterer) FilterRaffleCreated(opts *bind.FilterOpts, taskId []*big.Int, owner []common.Address) (*ContractRaffleCreatedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RaffleCreated", taskIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ContractRaffleCreatedIterator{contract: _Contract.contract, event: "RaffleCreated", logs: logs, sub: sub}, nil
}

// WatchRaffleCreated is a free log subscription operation binding the contract event 0x71f77ac840b74f30ac7115cc2fca45c4157ccc2b0a6e72f4183c587d77baa0ab.
//
// Solidity: event RaffleCreated(uint256 indexed taskId, address indexed owner, uint256 startTime, uint256 duration)
func (_Contract *ContractFilterer) WatchRaffleCreated(opts *bind.WatchOpts, sink chan<- *ContractRaffleCreated, taskId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RaffleCreated", taskIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRaffleCreated)
				if err := _Contract.contract.UnpackLog(event, "RaffleCreated", log); err != nil {
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

// ParseRaffleCreated is a log parse operation binding the contract event 0x71f77ac840b74f30ac7115cc2fca45c4157ccc2b0a6e72f4183c587d77baa0ab.
//
// Solidity: event RaffleCreated(uint256 indexed taskId, address indexed owner, uint256 startTime, uint256 duration)
func (_Contract *ContractFilterer) ParseRaffleCreated(log types.Log) (*ContractRaffleCreated, error) {
	event := new(ContractRaffleCreated)
	if err := _Contract.contract.UnpackLog(event, "RaffleCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRaffleEnteredIterator is returned from FilterRaffleEntered and is used to iterate over the raw logs and unpacked data for RaffleEntered events raised by the Contract contract.
type ContractRaffleEnteredIterator struct {
	Event *ContractRaffleEntered // Event containing the contract specifics and raw log

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
func (it *ContractRaffleEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRaffleEntered)
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
		it.Event = new(ContractRaffleEntered)
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
func (it *ContractRaffleEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRaffleEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRaffleEntered represents a RaffleEntered event raised by the Contract contract.
type ContractRaffleEntered struct {
	TaskId *big.Int
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaffleEntered is a free log retrieval operation binding the contract event 0x31ed7968c068023ef5cbf0834bb1d883bedc82dc6e9e44065bc1919184dded35.
//
// Solidity: event RaffleEntered(uint256 taskId, address user, uint256 amount)
func (_Contract *ContractFilterer) FilterRaffleEntered(opts *bind.FilterOpts) (*ContractRaffleEnteredIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RaffleEntered")
	if err != nil {
		return nil, err
	}
	return &ContractRaffleEnteredIterator{contract: _Contract.contract, event: "RaffleEntered", logs: logs, sub: sub}, nil
}

// WatchRaffleEntered is a free log subscription operation binding the contract event 0x31ed7968c068023ef5cbf0834bb1d883bedc82dc6e9e44065bc1919184dded35.
//
// Solidity: event RaffleEntered(uint256 taskId, address user, uint256 amount)
func (_Contract *ContractFilterer) WatchRaffleEntered(opts *bind.WatchOpts, sink chan<- *ContractRaffleEntered) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RaffleEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRaffleEntered)
				if err := _Contract.contract.UnpackLog(event, "RaffleEntered", log); err != nil {
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

// ParseRaffleEntered is a log parse operation binding the contract event 0x31ed7968c068023ef5cbf0834bb1d883bedc82dc6e9e44065bc1919184dded35.
//
// Solidity: event RaffleEntered(uint256 taskId, address user, uint256 amount)
func (_Contract *ContractFilterer) ParseRaffleEntered(log types.Log) (*ContractRaffleEntered, error) {
	event := new(ContractRaffleEntered)
	if err := _Contract.contract.UnpackLog(event, "RaffleEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWinningTicketPickedIterator is returned from FilterWinningTicketPicked and is used to iterate over the raw logs and unpacked data for WinningTicketPicked events raised by the Contract contract.
type ContractWinningTicketPickedIterator struct {
	Event *ContractWinningTicketPicked // Event containing the contract specifics and raw log

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
func (it *ContractWinningTicketPickedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWinningTicketPicked)
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
		it.Event = new(ContractWinningTicketPicked)
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
func (it *ContractWinningTicketPickedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWinningTicketPickedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWinningTicketPicked represents a WinningTicketPicked event raised by the Contract contract.
type ContractWinningTicketPicked struct {
	TaskId        *big.Int
	WinningTicket *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWinningTicketPicked is a free log retrieval operation binding the contract event 0x05459ed43a6fc64488d6051768b4c51dd130988284a0c8282e086078cd5ab546.
//
// Solidity: event WinningTicketPicked(uint256 taskId, uint256 winningTicket)
func (_Contract *ContractFilterer) FilterWinningTicketPicked(opts *bind.FilterOpts) (*ContractWinningTicketPickedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "WinningTicketPicked")
	if err != nil {
		return nil, err
	}
	return &ContractWinningTicketPickedIterator{contract: _Contract.contract, event: "WinningTicketPicked", logs: logs, sub: sub}, nil
}

// WatchWinningTicketPicked is a free log subscription operation binding the contract event 0x05459ed43a6fc64488d6051768b4c51dd130988284a0c8282e086078cd5ab546.
//
// Solidity: event WinningTicketPicked(uint256 taskId, uint256 winningTicket)
func (_Contract *ContractFilterer) WatchWinningTicketPicked(opts *bind.WatchOpts, sink chan<- *ContractWinningTicketPicked) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "WinningTicketPicked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWinningTicketPicked)
				if err := _Contract.contract.UnpackLog(event, "WinningTicketPicked", log); err != nil {
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

// ParseWinningTicketPicked is a log parse operation binding the contract event 0x05459ed43a6fc64488d6051768b4c51dd130988284a0c8282e086078cd5ab546.
//
// Solidity: event WinningTicketPicked(uint256 taskId, uint256 winningTicket)
func (_Contract *ContractFilterer) ParseWinningTicketPicked(log types.Log) (*ContractWinningTicketPicked, error) {
	event := new(ContractWinningTicketPicked)
	if err := _Contract.contract.UnpackLog(event, "WinningTicketPicked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
