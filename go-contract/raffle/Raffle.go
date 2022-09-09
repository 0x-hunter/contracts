// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package raffle

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

// RaffleMetaData contains all meta data concerning the Raffle contract.
var RaffleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usd\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_uniV3Factory\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_currencies\",\"type\":\"address[]\"},{\"internalType\":\"uint24[]\",\"name\":\"_feeRates\",\"type\":\"uint24[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundsAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FundsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"PrizeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"RaffleCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"RaffleClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"RaffleCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"}],\"name\":\"RaffleEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winningTicket\",\"type\":\"uint256\"}],\"name\":\"WinningTicketPicked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_raffleId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_raffleId\",\"type\":\"uint256\"}],\"name\":\"claimFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_raffleId\",\"type\":\"uint256\"}],\"name\":\"claimPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_raffleId\",\"type\":\"uint256\"}],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_ticketPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_ticketSize\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currencyInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"feeRate\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTaskId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_raffleId\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"_amount\",\"type\":\"uint96\"}],\"name\":\"enter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"}],\"name\":\"getCurrencyPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prizeInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"winningNumber\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"prizeClaimed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"fundsClaimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"raffleInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"ticketSize\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ticketPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint160\",\"name\":\"funds\",\"type\":\"uint160\"},{\"internalType\":\"uint96\",\"name\":\"soldTickets\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"enumRaffle.TaskStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"_feeRate\",\"type\":\"uint24\"}],\"name\":\"setCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ticketAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RaffleABI is the input ABI used to generate the binding from.
// Deprecated: Use RaffleMetaData.ABI instead.
var RaffleABI = RaffleMetaData.ABI

// Raffle is an auto generated Go binding around an Ethereum contract.
type Raffle struct {
	RaffleCaller     // Read-only binding to the contract
	RaffleTransactor // Write-only binding to the contract
	RaffleFilterer   // Log filterer for contract events
}

// RaffleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RaffleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RaffleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RaffleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RaffleSession struct {
	Contract     *Raffle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RaffleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RaffleCallerSession struct {
	Contract *RaffleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RaffleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RaffleTransactorSession struct {
	Contract     *RaffleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RaffleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RaffleRaw struct {
	Contract *Raffle // Generic contract binding to access the raw methods on
}

// RaffleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RaffleCallerRaw struct {
	Contract *RaffleCaller // Generic read-only contract binding to access the raw methods on
}

// RaffleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RaffleTransactorRaw struct {
	Contract *RaffleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRaffle creates a new instance of Raffle, bound to a specific deployed contract.
func NewRaffle(address common.Address, backend bind.ContractBackend) (*Raffle, error) {
	contract, err := bindRaffle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Raffle{RaffleCaller: RaffleCaller{contract: contract}, RaffleTransactor: RaffleTransactor{contract: contract}, RaffleFilterer: RaffleFilterer{contract: contract}}, nil
}

// NewRaffleCaller creates a new read-only instance of Raffle, bound to a specific deployed contract.
func NewRaffleCaller(address common.Address, caller bind.ContractCaller) (*RaffleCaller, error) {
	contract, err := bindRaffle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RaffleCaller{contract: contract}, nil
}

// NewRaffleTransactor creates a new write-only instance of Raffle, bound to a specific deployed contract.
func NewRaffleTransactor(address common.Address, transactor bind.ContractTransactor) (*RaffleTransactor, error) {
	contract, err := bindRaffle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RaffleTransactor{contract: contract}, nil
}

// NewRaffleFilterer creates a new log filterer instance of Raffle, bound to a specific deployed contract.
func NewRaffleFilterer(address common.Address, filterer bind.ContractFilterer) (*RaffleFilterer, error) {
	contract, err := bindRaffle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RaffleFilterer{contract: contract}, nil
}

// bindRaffle binds a generic wrapper to an already deployed contract.
func bindRaffle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RaffleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Raffle *RaffleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Raffle.Contract.RaffleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Raffle *RaffleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Raffle.Contract.RaffleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Raffle *RaffleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Raffle.Contract.RaffleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Raffle *RaffleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Raffle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Raffle *RaffleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Raffle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Raffle *RaffleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Raffle.Contract.contract.Transact(opts, method, params...)
}

// CurrencyInfo is a free data retrieval call binding the contract method 0x6d31cb2a.
//
// Solidity: function currencyInfo(address ) view returns(bool enabled, uint24 feeRate)
func (_Raffle *RaffleCaller) CurrencyInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Enabled bool
	FeeRate *big.Int
}, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "currencyInfo", arg0)

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
func (_Raffle *RaffleSession) CurrencyInfo(arg0 common.Address) (struct {
	Enabled bool
	FeeRate *big.Int
}, error) {
	return _Raffle.Contract.CurrencyInfo(&_Raffle.CallOpts, arg0)
}

// CurrencyInfo is a free data retrieval call binding the contract method 0x6d31cb2a.
//
// Solidity: function currencyInfo(address ) view returns(bool enabled, uint24 feeRate)
func (_Raffle *RaffleCallerSession) CurrencyInfo(arg0 common.Address) (struct {
	Enabled bool
	FeeRate *big.Int
}, error) {
	return _Raffle.Contract.CurrencyInfo(&_Raffle.CallOpts, arg0)
}

// CurrentTaskId is a free data retrieval call binding the contract method 0x98934a1f.
//
// Solidity: function currentTaskId() view returns(uint256)
func (_Raffle *RaffleCaller) CurrentTaskId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "currentTaskId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTaskId is a free data retrieval call binding the contract method 0x98934a1f.
//
// Solidity: function currentTaskId() view returns(uint256)
func (_Raffle *RaffleSession) CurrentTaskId() (*big.Int, error) {
	return _Raffle.Contract.CurrentTaskId(&_Raffle.CallOpts)
}

// CurrentTaskId is a free data retrieval call binding the contract method 0x98934a1f.
//
// Solidity: function currentTaskId() view returns(uint256)
func (_Raffle *RaffleCallerSession) CurrentTaskId() (*big.Int, error) {
	return _Raffle.Contract.CurrentTaskId(&_Raffle.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Raffle *RaffleCaller) FeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Raffle *RaffleSession) FeeRate() (*big.Int, error) {
	return _Raffle.Contract.FeeRate(&_Raffle.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Raffle *RaffleCallerSession) FeeRate() (*big.Int, error) {
	return _Raffle.Contract.FeeRate(&_Raffle.CallOpts)
}

// FeeSetter is a free data retrieval call binding the contract method 0x87cf3ef4.
//
// Solidity: function feeSetter() view returns(address)
func (_Raffle *RaffleCaller) FeeSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "feeSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeSetter is a free data retrieval call binding the contract method 0x87cf3ef4.
//
// Solidity: function feeSetter() view returns(address)
func (_Raffle *RaffleSession) FeeSetter() (common.Address, error) {
	return _Raffle.Contract.FeeSetter(&_Raffle.CallOpts)
}

// FeeSetter is a free data retrieval call binding the contract method 0x87cf3ef4.
//
// Solidity: function feeSetter() view returns(address)
func (_Raffle *RaffleCallerSession) FeeSetter() (common.Address, error) {
	return _Raffle.Contract.FeeSetter(&_Raffle.CallOpts)
}

// GetCurrencyPrice is a free data retrieval call binding the contract method 0x7058346f.
//
// Solidity: function getCurrencyPrice(address _currency) view returns(uint256 price)
func (_Raffle *RaffleCaller) GetCurrencyPrice(opts *bind.CallOpts, _currency common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "getCurrencyPrice", _currency)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrencyPrice is a free data retrieval call binding the contract method 0x7058346f.
//
// Solidity: function getCurrencyPrice(address _currency) view returns(uint256 price)
func (_Raffle *RaffleSession) GetCurrencyPrice(_currency common.Address) (*big.Int, error) {
	return _Raffle.Contract.GetCurrencyPrice(&_Raffle.CallOpts, _currency)
}

// GetCurrencyPrice is a free data retrieval call binding the contract method 0x7058346f.
//
// Solidity: function getCurrencyPrice(address _currency) view returns(uint256 price)
func (_Raffle *RaffleCallerSession) GetCurrencyPrice(_currency common.Address) (*big.Int, error) {
	return _Raffle.Contract.GetCurrencyPrice(&_Raffle.CallOpts, _currency)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Raffle *RaffleCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Raffle *RaffleSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Raffle.Contract.OnERC721Received(&_Raffle.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Raffle *RaffleCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Raffle.Contract.OnERC721Received(&_Raffle.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Raffle *RaffleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Raffle *RaffleSession) Owner() (common.Address, error) {
	return _Raffle.Contract.Owner(&_Raffle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Raffle *RaffleCallerSession) Owner() (common.Address, error) {
	return _Raffle.Contract.Owner(&_Raffle.CallOpts)
}

// PrizeInfo is a free data retrieval call binding the contract method 0x726a41cf.
//
// Solidity: function prizeInfo(uint256 ) view returns(uint32 winningNumber, bool prizeClaimed, bool fundsClaimed)
func (_Raffle *RaffleCaller) PrizeInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	WinningNumber uint32
	PrizeClaimed  bool
	FundsClaimed  bool
}, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "prizeInfo", arg0)

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
func (_Raffle *RaffleSession) PrizeInfo(arg0 *big.Int) (struct {
	WinningNumber uint32
	PrizeClaimed  bool
	FundsClaimed  bool
}, error) {
	return _Raffle.Contract.PrizeInfo(&_Raffle.CallOpts, arg0)
}

// PrizeInfo is a free data retrieval call binding the contract method 0x726a41cf.
//
// Solidity: function prizeInfo(uint256 ) view returns(uint32 winningNumber, bool prizeClaimed, bool fundsClaimed)
func (_Raffle *RaffleCallerSession) PrizeInfo(arg0 *big.Int) (struct {
	WinningNumber uint32
	PrizeClaimed  bool
	FundsClaimed  bool
}, error) {
	return _Raffle.Contract.PrizeInfo(&_Raffle.CallOpts, arg0)
}

// RaffleInfo is a free data retrieval call binding the contract method 0xbf30d790.
//
// Solidity: function raffleInfo(uint256 ) view returns(address owner, uint64 endTime, uint32 ticketSize, uint256 tokenId, address tokenAddress, uint64 ticketPrice, uint160 funds, uint96 soldTickets, address currency, uint8 status)
func (_Raffle *RaffleCaller) RaffleInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _Raffle.contract.Call(opts, &out, "raffleInfo", arg0)

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
func (_Raffle *RaffleSession) RaffleInfo(arg0 *big.Int) (struct {
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
	return _Raffle.Contract.RaffleInfo(&_Raffle.CallOpts, arg0)
}

// RaffleInfo is a free data retrieval call binding the contract method 0xbf30d790.
//
// Solidity: function raffleInfo(uint256 ) view returns(address owner, uint64 endTime, uint32 ticketSize, uint256 tokenId, address tokenAddress, uint64 ticketPrice, uint160 funds, uint96 soldTickets, address currency, uint8 status)
func (_Raffle *RaffleCallerSession) RaffleInfo(arg0 *big.Int) (struct {
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
	return _Raffle.Contract.RaffleInfo(&_Raffle.CallOpts, arg0)
}

// RequestInfo is a free data retrieval call binding the contract method 0x4f588bf1.
//
// Solidity: function requestInfo(uint256 ) view returns(uint256)
func (_Raffle *RaffleCaller) RequestInfo(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "requestInfo", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestInfo is a free data retrieval call binding the contract method 0x4f588bf1.
//
// Solidity: function requestInfo(uint256 ) view returns(uint256)
func (_Raffle *RaffleSession) RequestInfo(arg0 *big.Int) (*big.Int, error) {
	return _Raffle.Contract.RequestInfo(&_Raffle.CallOpts, arg0)
}

// RequestInfo is a free data retrieval call binding the contract method 0x4f588bf1.
//
// Solidity: function requestInfo(uint256 ) view returns(uint256)
func (_Raffle *RaffleCallerSession) RequestInfo(arg0 *big.Int) (*big.Int, error) {
	return _Raffle.Contract.RequestInfo(&_Raffle.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 ticketAmount)
func (_Raffle *RaffleCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "userInfo", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 ticketAmount)
func (_Raffle *RaffleSession) UserInfo(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Raffle.Contract.UserInfo(&_Raffle.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 ticketAmount)
func (_Raffle *RaffleCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Raffle.Contract.UserInfo(&_Raffle.CallOpts, arg0, arg1)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _raffleId) returns()
func (_Raffle *RaffleTransactor) Cancel(opts *bind.TransactOpts, _raffleId *big.Int) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "cancel", _raffleId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _raffleId) returns()
func (_Raffle *RaffleSession) Cancel(_raffleId *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.Cancel(&_Raffle.TransactOpts, _raffleId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _raffleId) returns()
func (_Raffle *RaffleTransactorSession) Cancel(_raffleId *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.Cancel(&_Raffle.TransactOpts, _raffleId)
}

// ClaimFunds is a paid mutator transaction binding the contract method 0x1b55e338.
//
// Solidity: function claimFunds(uint256 _raffleId) returns()
func (_Raffle *RaffleTransactor) ClaimFunds(opts *bind.TransactOpts, _raffleId *big.Int) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "claimFunds", _raffleId)
}

// ClaimFunds is a paid mutator transaction binding the contract method 0x1b55e338.
//
// Solidity: function claimFunds(uint256 _raffleId) returns()
func (_Raffle *RaffleSession) ClaimFunds(_raffleId *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.ClaimFunds(&_Raffle.TransactOpts, _raffleId)
}

// ClaimFunds is a paid mutator transaction binding the contract method 0x1b55e338.
//
// Solidity: function claimFunds(uint256 _raffleId) returns()
func (_Raffle *RaffleTransactorSession) ClaimFunds(_raffleId *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.ClaimFunds(&_Raffle.TransactOpts, _raffleId)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0xd7098154.
//
// Solidity: function claimPrize(uint256 _raffleId) returns()
func (_Raffle *RaffleTransactor) ClaimPrize(opts *bind.TransactOpts, _raffleId *big.Int) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "claimPrize", _raffleId)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0xd7098154.
//
// Solidity: function claimPrize(uint256 _raffleId) returns()
func (_Raffle *RaffleSession) ClaimPrize(_raffleId *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.ClaimPrize(&_Raffle.TransactOpts, _raffleId)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0xd7098154.
//
// Solidity: function claimPrize(uint256 _raffleId) returns()
func (_Raffle *RaffleTransactorSession) ClaimPrize(_raffleId *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.ClaimPrize(&_Raffle.TransactOpts, _raffleId)
}

// Close is a paid mutator transaction binding the contract method 0x0aebeb4e.
//
// Solidity: function close(uint256 _raffleId) returns()
func (_Raffle *RaffleTransactor) Close(opts *bind.TransactOpts, _raffleId *big.Int) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "close", _raffleId)
}

// Close is a paid mutator transaction binding the contract method 0x0aebeb4e.
//
// Solidity: function close(uint256 _raffleId) returns()
func (_Raffle *RaffleSession) Close(_raffleId *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.Close(&_Raffle.TransactOpts, _raffleId)
}

// Close is a paid mutator transaction binding the contract method 0x0aebeb4e.
//
// Solidity: function close(uint256 _raffleId) returns()
func (_Raffle *RaffleTransactorSession) Close(_raffleId *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.Close(&_Raffle.TransactOpts, _raffleId)
}

// Create is a paid mutator transaction binding the contract method 0xa7c1891e.
//
// Solidity: function create(address _tokenAddr, uint256 _tokenId, uint256 _duration, uint64 _ticketPrice, uint32 _ticketSize, address _currency) returns()
func (_Raffle *RaffleTransactor) Create(opts *bind.TransactOpts, _tokenAddr common.Address, _tokenId *big.Int, _duration *big.Int, _ticketPrice uint64, _ticketSize uint32, _currency common.Address) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "create", _tokenAddr, _tokenId, _duration, _ticketPrice, _ticketSize, _currency)
}

// Create is a paid mutator transaction binding the contract method 0xa7c1891e.
//
// Solidity: function create(address _tokenAddr, uint256 _tokenId, uint256 _duration, uint64 _ticketPrice, uint32 _ticketSize, address _currency) returns()
func (_Raffle *RaffleSession) Create(_tokenAddr common.Address, _tokenId *big.Int, _duration *big.Int, _ticketPrice uint64, _ticketSize uint32, _currency common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.Create(&_Raffle.TransactOpts, _tokenAddr, _tokenId, _duration, _ticketPrice, _ticketSize, _currency)
}

// Create is a paid mutator transaction binding the contract method 0xa7c1891e.
//
// Solidity: function create(address _tokenAddr, uint256 _tokenId, uint256 _duration, uint64 _ticketPrice, uint32 _ticketSize, address _currency) returns()
func (_Raffle *RaffleTransactorSession) Create(_tokenAddr common.Address, _tokenId *big.Int, _duration *big.Int, _ticketPrice uint64, _ticketSize uint32, _currency common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.Create(&_Raffle.TransactOpts, _tokenAddr, _tokenId, _duration, _ticketPrice, _ticketSize, _currency)
}

// Enter is a paid mutator transaction binding the contract method 0x00d407d6.
//
// Solidity: function enter(uint256 _raffleId, uint96 _amount) returns()
func (_Raffle *RaffleTransactor) Enter(opts *bind.TransactOpts, _raffleId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "enter", _raffleId, _amount)
}

// Enter is a paid mutator transaction binding the contract method 0x00d407d6.
//
// Solidity: function enter(uint256 _raffleId, uint96 _amount) returns()
func (_Raffle *RaffleSession) Enter(_raffleId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.Enter(&_Raffle.TransactOpts, _raffleId, _amount)
}

// Enter is a paid mutator transaction binding the contract method 0x00d407d6.
//
// Solidity: function enter(uint256 _raffleId, uint96 _amount) returns()
func (_Raffle *RaffleTransactorSession) Enter(_raffleId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.Enter(&_Raffle.TransactOpts, _raffleId, _amount)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Raffle *RaffleTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Raffle *RaffleSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.RawFulfillRandomWords(&_Raffle.TransactOpts, requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Raffle *RaffleTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.RawFulfillRandomWords(&_Raffle.TransactOpts, requestId, randomWords)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Raffle *RaffleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Raffle *RaffleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Raffle.Contract.RenounceOwnership(&_Raffle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Raffle *RaffleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Raffle.Contract.RenounceOwnership(&_Raffle.TransactOpts)
}

// SetCurrency is a paid mutator transaction binding the contract method 0x59991a77.
//
// Solidity: function setCurrency(address _currency, bool _enable, uint24 _feeRate) returns()
func (_Raffle *RaffleTransactor) SetCurrency(opts *bind.TransactOpts, _currency common.Address, _enable bool, _feeRate *big.Int) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "setCurrency", _currency, _enable, _feeRate)
}

// SetCurrency is a paid mutator transaction binding the contract method 0x59991a77.
//
// Solidity: function setCurrency(address _currency, bool _enable, uint24 _feeRate) returns()
func (_Raffle *RaffleSession) SetCurrency(_currency common.Address, _enable bool, _feeRate *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.SetCurrency(&_Raffle.TransactOpts, _currency, _enable, _feeRate)
}

// SetCurrency is a paid mutator transaction binding the contract method 0x59991a77.
//
// Solidity: function setCurrency(address _currency, bool _enable, uint24 _feeRate) returns()
func (_Raffle *RaffleTransactorSession) SetCurrency(_currency common.Address, _enable bool, _feeRate *big.Int) (*types.Transaction, error) {
	return _Raffle.Contract.SetCurrency(&_Raffle.TransactOpts, _currency, _enable, _feeRate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Raffle *RaffleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Raffle *RaffleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.TransferOwnership(&_Raffle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Raffle *RaffleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.TransferOwnership(&_Raffle.TransactOpts, newOwner)
}

// RaffleFundsClaimedIterator is returned from FilterFundsClaimed and is used to iterate over the raw logs and unpacked data for FundsClaimed events raised by the Raffle contract.
type RaffleFundsClaimedIterator struct {
	Event *RaffleFundsClaimed // Event containing the contract specifics and raw log

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
func (it *RaffleFundsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleFundsClaimed)
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
		it.Event = new(RaffleFundsClaimed)
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
func (it *RaffleFundsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleFundsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleFundsClaimed represents a FundsClaimed event raised by the Raffle contract.
type RaffleFundsClaimed struct {
	TaskId      *big.Int
	FundsAmount *big.Int
	Owner       common.Address
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFundsClaimed is a free log retrieval operation binding the contract event 0x3f53f5f381ee0b1fbf730fa56efb0a2aec8a882b2847a1ec0279982d8846e4f2.
//
// Solidity: event FundsClaimed(uint256 taskId, uint256 fundsAmount, address owner, uint256 fee)
func (_Raffle *RaffleFilterer) FilterFundsClaimed(opts *bind.FilterOpts) (*RaffleFundsClaimedIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "FundsClaimed")
	if err != nil {
		return nil, err
	}
	return &RaffleFundsClaimedIterator{contract: _Raffle.contract, event: "FundsClaimed", logs: logs, sub: sub}, nil
}

// WatchFundsClaimed is a free log subscription operation binding the contract event 0x3f53f5f381ee0b1fbf730fa56efb0a2aec8a882b2847a1ec0279982d8846e4f2.
//
// Solidity: event FundsClaimed(uint256 taskId, uint256 fundsAmount, address owner, uint256 fee)
func (_Raffle *RaffleFilterer) WatchFundsClaimed(opts *bind.WatchOpts, sink chan<- *RaffleFundsClaimed) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "FundsClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleFundsClaimed)
				if err := _Raffle.contract.UnpackLog(event, "FundsClaimed", log); err != nil {
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
func (_Raffle *RaffleFilterer) ParseFundsClaimed(log types.Log) (*RaffleFundsClaimed, error) {
	event := new(RaffleFundsClaimed)
	if err := _Raffle.contract.UnpackLog(event, "FundsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Raffle contract.
type RaffleOwnershipTransferredIterator struct {
	Event *RaffleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RaffleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleOwnershipTransferred)
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
		it.Event = new(RaffleOwnershipTransferred)
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
func (it *RaffleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleOwnershipTransferred represents a OwnershipTransferred event raised by the Raffle contract.
type RaffleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Raffle *RaffleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RaffleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RaffleOwnershipTransferredIterator{contract: _Raffle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Raffle *RaffleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RaffleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleOwnershipTransferred)
				if err := _Raffle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Raffle *RaffleFilterer) ParseOwnershipTransferred(log types.Log) (*RaffleOwnershipTransferred, error) {
	event := new(RaffleOwnershipTransferred)
	if err := _Raffle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RafflePrizeClaimedIterator is returned from FilterPrizeClaimed and is used to iterate over the raw logs and unpacked data for PrizeClaimed events raised by the Raffle contract.
type RafflePrizeClaimedIterator struct {
	Event *RafflePrizeClaimed // Event containing the contract specifics and raw log

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
func (it *RafflePrizeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RafflePrizeClaimed)
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
		it.Event = new(RafflePrizeClaimed)
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
func (it *RafflePrizeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RafflePrizeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RafflePrizeClaimed represents a PrizeClaimed event raised by the Raffle contract.
type RafflePrizeClaimed struct {
	TaskId  *big.Int
	TokenId *big.Int
	Winner  common.Address
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPrizeClaimed is a free log retrieval operation binding the contract event 0x608c8c3c6a3d40f40b96fbee44f9b9aee57a39c2fcb8bb0f59d8892b2d5492a2.
//
// Solidity: event PrizeClaimed(uint256 taskId, uint256 tokenId, address winner, uint256 fee)
func (_Raffle *RaffleFilterer) FilterPrizeClaimed(opts *bind.FilterOpts) (*RafflePrizeClaimedIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "PrizeClaimed")
	if err != nil {
		return nil, err
	}
	return &RafflePrizeClaimedIterator{contract: _Raffle.contract, event: "PrizeClaimed", logs: logs, sub: sub}, nil
}

// WatchPrizeClaimed is a free log subscription operation binding the contract event 0x608c8c3c6a3d40f40b96fbee44f9b9aee57a39c2fcb8bb0f59d8892b2d5492a2.
//
// Solidity: event PrizeClaimed(uint256 taskId, uint256 tokenId, address winner, uint256 fee)
func (_Raffle *RaffleFilterer) WatchPrizeClaimed(opts *bind.WatchOpts, sink chan<- *RafflePrizeClaimed) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "PrizeClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RafflePrizeClaimed)
				if err := _Raffle.contract.UnpackLog(event, "PrizeClaimed", log); err != nil {
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
func (_Raffle *RaffleFilterer) ParsePrizeClaimed(log types.Log) (*RafflePrizeClaimed, error) {
	event := new(RafflePrizeClaimed)
	if err := _Raffle.contract.UnpackLog(event, "PrizeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleRaffleCancelledIterator is returned from FilterRaffleCancelled and is used to iterate over the raw logs and unpacked data for RaffleCancelled events raised by the Raffle contract.
type RaffleRaffleCancelledIterator struct {
	Event *RaffleRaffleCancelled // Event containing the contract specifics and raw log

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
func (it *RaffleRaffleCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleRaffleCancelled)
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
		it.Event = new(RaffleRaffleCancelled)
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
func (it *RaffleRaffleCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleRaffleCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleRaffleCancelled represents a RaffleCancelled event raised by the Raffle contract.
type RaffleRaffleCancelled struct {
	TaskId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaffleCancelled is a free log retrieval operation binding the contract event 0x7fa78bab5d5570f162c69234a3f4d1b9294fd5bca72219cb16c1fe38839ccb8a.
//
// Solidity: event RaffleCancelled(uint256 taskId)
func (_Raffle *RaffleFilterer) FilterRaffleCancelled(opts *bind.FilterOpts) (*RaffleRaffleCancelledIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "RaffleCancelled")
	if err != nil {
		return nil, err
	}
	return &RaffleRaffleCancelledIterator{contract: _Raffle.contract, event: "RaffleCancelled", logs: logs, sub: sub}, nil
}

// WatchRaffleCancelled is a free log subscription operation binding the contract event 0x7fa78bab5d5570f162c69234a3f4d1b9294fd5bca72219cb16c1fe38839ccb8a.
//
// Solidity: event RaffleCancelled(uint256 taskId)
func (_Raffle *RaffleFilterer) WatchRaffleCancelled(opts *bind.WatchOpts, sink chan<- *RaffleRaffleCancelled) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "RaffleCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleRaffleCancelled)
				if err := _Raffle.contract.UnpackLog(event, "RaffleCancelled", log); err != nil {
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
func (_Raffle *RaffleFilterer) ParseRaffleCancelled(log types.Log) (*RaffleRaffleCancelled, error) {
	event := new(RaffleRaffleCancelled)
	if err := _Raffle.contract.UnpackLog(event, "RaffleCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleRaffleClosedIterator is returned from FilterRaffleClosed and is used to iterate over the raw logs and unpacked data for RaffleClosed events raised by the Raffle contract.
type RaffleRaffleClosedIterator struct {
	Event *RaffleRaffleClosed // Event containing the contract specifics and raw log

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
func (it *RaffleRaffleClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleRaffleClosed)
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
		it.Event = new(RaffleRaffleClosed)
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
func (it *RaffleRaffleClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleRaffleClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleRaffleClosed represents a RaffleClosed event raised by the Raffle contract.
type RaffleRaffleClosed struct {
	TaskId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaffleClosed is a free log retrieval operation binding the contract event 0x2a7d48246875b5f17dc9944def25fc8fe90e524ca2165e290ffdc0e44a0bf8b6.
//
// Solidity: event RaffleClosed(uint256 taskId)
func (_Raffle *RaffleFilterer) FilterRaffleClosed(opts *bind.FilterOpts) (*RaffleRaffleClosedIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "RaffleClosed")
	if err != nil {
		return nil, err
	}
	return &RaffleRaffleClosedIterator{contract: _Raffle.contract, event: "RaffleClosed", logs: logs, sub: sub}, nil
}

// WatchRaffleClosed is a free log subscription operation binding the contract event 0x2a7d48246875b5f17dc9944def25fc8fe90e524ca2165e290ffdc0e44a0bf8b6.
//
// Solidity: event RaffleClosed(uint256 taskId)
func (_Raffle *RaffleFilterer) WatchRaffleClosed(opts *bind.WatchOpts, sink chan<- *RaffleRaffleClosed) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "RaffleClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleRaffleClosed)
				if err := _Raffle.contract.UnpackLog(event, "RaffleClosed", log); err != nil {
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
func (_Raffle *RaffleFilterer) ParseRaffleClosed(log types.Log) (*RaffleRaffleClosed, error) {
	event := new(RaffleRaffleClosed)
	if err := _Raffle.contract.UnpackLog(event, "RaffleClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleRaffleCreatedIterator is returned from FilterRaffleCreated and is used to iterate over the raw logs and unpacked data for RaffleCreated events raised by the Raffle contract.
type RaffleRaffleCreatedIterator struct {
	Event *RaffleRaffleCreated // Event containing the contract specifics and raw log

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
func (it *RaffleRaffleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleRaffleCreated)
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
		it.Event = new(RaffleRaffleCreated)
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
func (it *RaffleRaffleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleRaffleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleRaffleCreated represents a RaffleCreated event raised by the Raffle contract.
type RaffleRaffleCreated struct {
	TaskId    *big.Int
	Owner     common.Address
	StartTime *big.Int
	Duration  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRaffleCreated is a free log retrieval operation binding the contract event 0x71f77ac840b74f30ac7115cc2fca45c4157ccc2b0a6e72f4183c587d77baa0ab.
//
// Solidity: event RaffleCreated(uint256 indexed taskId, address indexed owner, uint256 startTime, uint256 duration)
func (_Raffle *RaffleFilterer) FilterRaffleCreated(opts *bind.FilterOpts, taskId []*big.Int, owner []common.Address) (*RaffleRaffleCreatedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "RaffleCreated", taskIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &RaffleRaffleCreatedIterator{contract: _Raffle.contract, event: "RaffleCreated", logs: logs, sub: sub}, nil
}

// WatchRaffleCreated is a free log subscription operation binding the contract event 0x71f77ac840b74f30ac7115cc2fca45c4157ccc2b0a6e72f4183c587d77baa0ab.
//
// Solidity: event RaffleCreated(uint256 indexed taskId, address indexed owner, uint256 startTime, uint256 duration)
func (_Raffle *RaffleFilterer) WatchRaffleCreated(opts *bind.WatchOpts, sink chan<- *RaffleRaffleCreated, taskId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "RaffleCreated", taskIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleRaffleCreated)
				if err := _Raffle.contract.UnpackLog(event, "RaffleCreated", log); err != nil {
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
func (_Raffle *RaffleFilterer) ParseRaffleCreated(log types.Log) (*RaffleRaffleCreated, error) {
	event := new(RaffleRaffleCreated)
	if err := _Raffle.contract.UnpackLog(event, "RaffleCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleRaffleEnteredIterator is returned from FilterRaffleEntered and is used to iterate over the raw logs and unpacked data for RaffleEntered events raised by the Raffle contract.
type RaffleRaffleEnteredIterator struct {
	Event *RaffleRaffleEntered // Event containing the contract specifics and raw log

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
func (it *RaffleRaffleEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleRaffleEntered)
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
		it.Event = new(RaffleRaffleEntered)
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
func (it *RaffleRaffleEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleRaffleEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleRaffleEntered represents a RaffleEntered event raised by the Raffle contract.
type RaffleRaffleEntered struct {
	TaskId     *big.Int
	User       common.Address
	LowerBound *big.Int
	UpperBound *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRaffleEntered is a free log retrieval operation binding the contract event 0xeb79e1c4c2e9267b8cc3b7b7df468f4a50a3fb1c8093d604f2c281ebdf2ae829.
//
// Solidity: event RaffleEntered(uint256 taskId, address user, uint256 lowerBound, uint256 upperBound)
func (_Raffle *RaffleFilterer) FilterRaffleEntered(opts *bind.FilterOpts) (*RaffleRaffleEnteredIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "RaffleEntered")
	if err != nil {
		return nil, err
	}
	return &RaffleRaffleEnteredIterator{contract: _Raffle.contract, event: "RaffleEntered", logs: logs, sub: sub}, nil
}

// WatchRaffleEntered is a free log subscription operation binding the contract event 0xeb79e1c4c2e9267b8cc3b7b7df468f4a50a3fb1c8093d604f2c281ebdf2ae829.
//
// Solidity: event RaffleEntered(uint256 taskId, address user, uint256 lowerBound, uint256 upperBound)
func (_Raffle *RaffleFilterer) WatchRaffleEntered(opts *bind.WatchOpts, sink chan<- *RaffleRaffleEntered) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "RaffleEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleRaffleEntered)
				if err := _Raffle.contract.UnpackLog(event, "RaffleEntered", log); err != nil {
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

// ParseRaffleEntered is a log parse operation binding the contract event 0xeb79e1c4c2e9267b8cc3b7b7df468f4a50a3fb1c8093d604f2c281ebdf2ae829.
//
// Solidity: event RaffleEntered(uint256 taskId, address user, uint256 lowerBound, uint256 upperBound)
func (_Raffle *RaffleFilterer) ParseRaffleEntered(log types.Log) (*RaffleRaffleEntered, error) {
	event := new(RaffleRaffleEntered)
	if err := _Raffle.contract.UnpackLog(event, "RaffleEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleWinningTicketPickedIterator is returned from FilterWinningTicketPicked and is used to iterate over the raw logs and unpacked data for WinningTicketPicked events raised by the Raffle contract.
type RaffleWinningTicketPickedIterator struct {
	Event *RaffleWinningTicketPicked // Event containing the contract specifics and raw log

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
func (it *RaffleWinningTicketPickedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleWinningTicketPicked)
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
		it.Event = new(RaffleWinningTicketPicked)
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
func (it *RaffleWinningTicketPickedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleWinningTicketPickedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleWinningTicketPicked represents a WinningTicketPicked event raised by the Raffle contract.
type RaffleWinningTicketPicked struct {
	TaskId        *big.Int
	WinningTicket *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWinningTicketPicked is a free log retrieval operation binding the contract event 0x05459ed43a6fc64488d6051768b4c51dd130988284a0c8282e086078cd5ab546.
//
// Solidity: event WinningTicketPicked(uint256 taskId, uint256 winningTicket)
func (_Raffle *RaffleFilterer) FilterWinningTicketPicked(opts *bind.FilterOpts) (*RaffleWinningTicketPickedIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "WinningTicketPicked")
	if err != nil {
		return nil, err
	}
	return &RaffleWinningTicketPickedIterator{contract: _Raffle.contract, event: "WinningTicketPicked", logs: logs, sub: sub}, nil
}

// WatchWinningTicketPicked is a free log subscription operation binding the contract event 0x05459ed43a6fc64488d6051768b4c51dd130988284a0c8282e086078cd5ab546.
//
// Solidity: event WinningTicketPicked(uint256 taskId, uint256 winningTicket)
func (_Raffle *RaffleFilterer) WatchWinningTicketPicked(opts *bind.WatchOpts, sink chan<- *RaffleWinningTicketPicked) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "WinningTicketPicked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleWinningTicketPicked)
				if err := _Raffle.contract.UnpackLog(event, "WinningTicketPicked", log); err != nil {
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
func (_Raffle *RaffleFilterer) ParseWinningTicketPicked(log types.Log) (*RaffleWinningTicketPicked, error) {
	event := new(RaffleWinningTicketPicked)
	if err := _Raffle.contract.UnpackLog(event, "WinningTicketPicked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
