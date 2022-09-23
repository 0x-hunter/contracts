// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iuniswapv3poolstate

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

// Iuniswapv3poolstateMetaData contains all meta data concerning the Iuniswapv3poolstate contract.
var Iuniswapv3poolstateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"feeGrowthGlobal0X128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGrowthGlobal1X128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"observations\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"blockTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"int56\",\"name\":\"tickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityCumulativeX128\",\"type\":\"uint160\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"_liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFees\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"token0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"token1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slot0\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"observationIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinalityNext\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"feeProtocol\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"unlocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"wordPosition\",\"type\":\"int16\"}],\"name\":\"tickBitmap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"ticks\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidityGross\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"liquidityNet\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthOutside0X128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthOutside1X128\",\"type\":\"uint256\"},{\"internalType\":\"int56\",\"name\":\"tickCumulativeOutside\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityOutsideX128\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"secondsOutside\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Iuniswapv3poolstateABI is the input ABI used to generate the binding from.
// Deprecated: Use Iuniswapv3poolstateMetaData.ABI instead.
var Iuniswapv3poolstateABI = Iuniswapv3poolstateMetaData.ABI

// Iuniswapv3poolstate is an auto generated Go binding around an Ethereum contract.
type Iuniswapv3poolstate struct {
	Iuniswapv3poolstateCaller     // Read-only binding to the contract
	Iuniswapv3poolstateTransactor // Write-only binding to the contract
	Iuniswapv3poolstateFilterer   // Log filterer for contract events
}

// Iuniswapv3poolstateCaller is an auto generated read-only Go binding around an Ethereum contract.
type Iuniswapv3poolstateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iuniswapv3poolstateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Iuniswapv3poolstateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iuniswapv3poolstateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Iuniswapv3poolstateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iuniswapv3poolstateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Iuniswapv3poolstateSession struct {
	Contract     *Iuniswapv3poolstate // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Iuniswapv3poolstateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Iuniswapv3poolstateCallerSession struct {
	Contract *Iuniswapv3poolstateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// Iuniswapv3poolstateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Iuniswapv3poolstateTransactorSession struct {
	Contract     *Iuniswapv3poolstateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// Iuniswapv3poolstateRaw is an auto generated low-level Go binding around an Ethereum contract.
type Iuniswapv3poolstateRaw struct {
	Contract *Iuniswapv3poolstate // Generic contract binding to access the raw methods on
}

// Iuniswapv3poolstateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Iuniswapv3poolstateCallerRaw struct {
	Contract *Iuniswapv3poolstateCaller // Generic read-only contract binding to access the raw methods on
}

// Iuniswapv3poolstateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Iuniswapv3poolstateTransactorRaw struct {
	Contract *Iuniswapv3poolstateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIuniswapv3poolstate creates a new instance of Iuniswapv3poolstate, bound to a specific deployed contract.
func NewIuniswapv3poolstate(address common.Address, backend bind.ContractBackend) (*Iuniswapv3poolstate, error) {
	contract, err := bindIuniswapv3poolstate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv3poolstate{Iuniswapv3poolstateCaller: Iuniswapv3poolstateCaller{contract: contract}, Iuniswapv3poolstateTransactor: Iuniswapv3poolstateTransactor{contract: contract}, Iuniswapv3poolstateFilterer: Iuniswapv3poolstateFilterer{contract: contract}}, nil
}

// NewIuniswapv3poolstateCaller creates a new read-only instance of Iuniswapv3poolstate, bound to a specific deployed contract.
func NewIuniswapv3poolstateCaller(address common.Address, caller bind.ContractCaller) (*Iuniswapv3poolstateCaller, error) {
	contract, err := bindIuniswapv3poolstate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv3poolstateCaller{contract: contract}, nil
}

// NewIuniswapv3poolstateTransactor creates a new write-only instance of Iuniswapv3poolstate, bound to a specific deployed contract.
func NewIuniswapv3poolstateTransactor(address common.Address, transactor bind.ContractTransactor) (*Iuniswapv3poolstateTransactor, error) {
	contract, err := bindIuniswapv3poolstate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv3poolstateTransactor{contract: contract}, nil
}

// NewIuniswapv3poolstateFilterer creates a new log filterer instance of Iuniswapv3poolstate, bound to a specific deployed contract.
func NewIuniswapv3poolstateFilterer(address common.Address, filterer bind.ContractFilterer) (*Iuniswapv3poolstateFilterer, error) {
	contract, err := bindIuniswapv3poolstate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv3poolstateFilterer{contract: contract}, nil
}

// bindIuniswapv3poolstate binds a generic wrapper to an already deployed contract.
func bindIuniswapv3poolstate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Iuniswapv3poolstateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iuniswapv3poolstate *Iuniswapv3poolstateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iuniswapv3poolstate.Contract.Iuniswapv3poolstateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iuniswapv3poolstate *Iuniswapv3poolstateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iuniswapv3poolstate.Contract.Iuniswapv3poolstateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iuniswapv3poolstate *Iuniswapv3poolstateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iuniswapv3poolstate.Contract.Iuniswapv3poolstateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iuniswapv3poolstate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iuniswapv3poolstate *Iuniswapv3poolstateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iuniswapv3poolstate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iuniswapv3poolstate *Iuniswapv3poolstateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iuniswapv3poolstate.Contract.contract.Transact(opts, method, params...)
}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCaller) FeeGrowthGlobal0X128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iuniswapv3poolstate.contract.Call(opts, &out, "feeGrowthGlobal0X128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateSession) FeeGrowthGlobal0X128() (*big.Int, error) {
	return _Iuniswapv3poolstate.Contract.FeeGrowthGlobal0X128(&_Iuniswapv3poolstate.CallOpts)
}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCallerSession) FeeGrowthGlobal0X128() (*big.Int, error) {
	return _Iuniswapv3poolstate.Contract.FeeGrowthGlobal0X128(&_Iuniswapv3poolstate.CallOpts)
}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCaller) FeeGrowthGlobal1X128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iuniswapv3poolstate.contract.Call(opts, &out, "feeGrowthGlobal1X128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateSession) FeeGrowthGlobal1X128() (*big.Int, error) {
	return _Iuniswapv3poolstate.Contract.FeeGrowthGlobal1X128(&_Iuniswapv3poolstate.CallOpts)
}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCallerSession) FeeGrowthGlobal1X128() (*big.Int, error) {
	return _Iuniswapv3poolstate.Contract.FeeGrowthGlobal1X128(&_Iuniswapv3poolstate.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCaller) Liquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iuniswapv3poolstate.contract.Call(opts, &out, "liquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateSession) Liquidity() (*big.Int, error) {
	return _Iuniswapv3poolstate.Contract.Liquidity(&_Iuniswapv3poolstate.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCallerSession) Liquidity() (*big.Int, error) {
	return _Iuniswapv3poolstate.Contract.Liquidity(&_Iuniswapv3poolstate.CallOpts)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 index) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCaller) Observations(opts *bind.CallOpts, index *big.Int) (struct {
	BlockTimestamp                    uint32
	TickCumulative                    *big.Int
	SecondsPerLiquidityCumulativeX128 *big.Int
	Initialized                       bool
}, error) {
	var out []interface{}
	err := _Iuniswapv3poolstate.contract.Call(opts, &out, "observations", index)

	outstruct := new(struct {
		BlockTimestamp                    uint32
		TickCumulative                    *big.Int
		SecondsPerLiquidityCumulativeX128 *big.Int
		Initialized                       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockTimestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.TickCumulative = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityCumulativeX128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Initialized = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 index) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateSession) Observations(index *big.Int) (struct {
	BlockTimestamp                    uint32
	TickCumulative                    *big.Int
	SecondsPerLiquidityCumulativeX128 *big.Int
	Initialized                       bool
}, error) {
	return _Iuniswapv3poolstate.Contract.Observations(&_Iuniswapv3poolstate.CallOpts, index)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 index) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCallerSession) Observations(index *big.Int) (struct {
	BlockTimestamp                    uint32
	TickCumulative                    *big.Int
	SecondsPerLiquidityCumulativeX128 *big.Int
	Initialized                       bool
}, error) {
	return _Iuniswapv3poolstate.Contract.Observations(&_Iuniswapv3poolstate.CallOpts, index)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 _liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCaller) Positions(opts *bind.CallOpts, key [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	var out []interface{}
	err := _Iuniswapv3poolstate.contract.Call(opts, &out, "positions", key)

	outstruct := new(struct {
		Liquidity                *big.Int
		FeeGrowthInside0LastX128 *big.Int
		FeeGrowthInside1LastX128 *big.Int
		TokensOwed0              *big.Int
		TokensOwed1              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside0LastX128 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside1LastX128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed0 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed1 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 _liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateSession) Positions(key [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _Iuniswapv3poolstate.Contract.Positions(&_Iuniswapv3poolstate.CallOpts, key)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 _liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCallerSession) Positions(key [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _Iuniswapv3poolstate.Contract.Positions(&_Iuniswapv3poolstate.CallOpts, key)
}

// ProtocolFees is a free data retrieval call binding the contract method 0x1ad8b03b.
//
// Solidity: function protocolFees() view returns(uint128 token0, uint128 token1)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCaller) ProtocolFees(opts *bind.CallOpts) (struct {
	Token0 *big.Int
	Token1 *big.Int
}, error) {
	var out []interface{}
	err := _Iuniswapv3poolstate.contract.Call(opts, &out, "protocolFees")

	outstruct := new(struct {
		Token0 *big.Int
		Token1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Token1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProtocolFees is a free data retrieval call binding the contract method 0x1ad8b03b.
//
// Solidity: function protocolFees() view returns(uint128 token0, uint128 token1)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateSession) ProtocolFees() (struct {
	Token0 *big.Int
	Token1 *big.Int
}, error) {
	return _Iuniswapv3poolstate.Contract.ProtocolFees(&_Iuniswapv3poolstate.CallOpts)
}

// ProtocolFees is a free data retrieval call binding the contract method 0x1ad8b03b.
//
// Solidity: function protocolFees() view returns(uint128 token0, uint128 token1)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCallerSession) ProtocolFees() (struct {
	Token0 *big.Int
	Token1 *big.Int
}, error) {
	return _Iuniswapv3poolstate.Contract.ProtocolFees(&_Iuniswapv3poolstate.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCaller) Slot0(opts *bind.CallOpts) (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	var out []interface{}
	err := _Iuniswapv3poolstate.contract.Call(opts, &out, "slot0")

	outstruct := new(struct {
		SqrtPriceX96               *big.Int
		Tick                       *big.Int
		ObservationIndex           uint16
		ObservationCardinality     uint16
		ObservationCardinalityNext uint16
		FeeProtocol                uint8
		Unlocked                   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPriceX96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ObservationIndex = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ObservationCardinality = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.ObservationCardinalityNext = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.FeeProtocol = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Unlocked = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _Iuniswapv3poolstate.Contract.Slot0(&_Iuniswapv3poolstate.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCallerSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _Iuniswapv3poolstate.Contract.Slot0(&_Iuniswapv3poolstate.CallOpts)
}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 wordPosition) view returns(uint256)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCaller) TickBitmap(opts *bind.CallOpts, wordPosition int16) (*big.Int, error) {
	var out []interface{}
	err := _Iuniswapv3poolstate.contract.Call(opts, &out, "tickBitmap", wordPosition)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 wordPosition) view returns(uint256)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateSession) TickBitmap(wordPosition int16) (*big.Int, error) {
	return _Iuniswapv3poolstate.Contract.TickBitmap(&_Iuniswapv3poolstate.CallOpts, wordPosition)
}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 wordPosition) view returns(uint256)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCallerSession) TickBitmap(wordPosition int16) (*big.Int, error) {
	return _Iuniswapv3poolstate.Contract.TickBitmap(&_Iuniswapv3poolstate.CallOpts, wordPosition)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCaller) Ticks(opts *bind.CallOpts, tick *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	var out []interface{}
	err := _Iuniswapv3poolstate.contract.Call(opts, &out, "ticks", tick)

	outstruct := new(struct {
		LiquidityGross                 *big.Int
		LiquidityNet                   *big.Int
		FeeGrowthOutside0X128          *big.Int
		FeeGrowthOutside1X128          *big.Int
		TickCumulativeOutside          *big.Int
		SecondsPerLiquidityOutsideX128 *big.Int
		SecondsOutside                 uint32
		Initialized                    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidityGross = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidityNet = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthOutside0X128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthOutside1X128 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TickCumulativeOutside = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityOutsideX128 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.SecondsOutside = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.Initialized = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateSession) Ticks(tick *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	return _Iuniswapv3poolstate.Contract.Ticks(&_Iuniswapv3poolstate.CallOpts, tick)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_Iuniswapv3poolstate *Iuniswapv3poolstateCallerSession) Ticks(tick *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	return _Iuniswapv3poolstate.Contract.Ticks(&_Iuniswapv3poolstate.CallOpts, tick)
}
