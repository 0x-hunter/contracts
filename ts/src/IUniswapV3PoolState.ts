import BN from 'bn.js';
import BigNumber from 'bignumber.js';
import {
  PromiEvent,
  TransactionReceipt,
  EventResponse,
  EventData,
  Web3ContractContext,
} from 'ethereum-abi-types-generator';

export interface CallOptions {
  from?: string;
  gasPrice?: string;
  gas?: number;
}

export interface SendOptions {
  from: string;
  value?: number | string | BN | BigNumber;
  gasPrice?: string;
  gas?: number;
}

export interface EstimateGasOptions {
  from?: string;
  value?: number | string | BN | BigNumber;
  gas?: number;
}

export interface MethodPayableReturnContext {
  send(options: SendOptions): PromiEvent<TransactionReceipt>;
  send(
    options: SendOptions,
    callback: (error: Error, result: any) => void
  ): PromiEvent<TransactionReceipt>;
  estimateGas(options: EstimateGasOptions): Promise<number>;
  estimateGas(
    options: EstimateGasOptions,
    callback: (error: Error, result: any) => void
  ): Promise<number>;
  encodeABI(): string;
}

export interface MethodConstantReturnContext<TCallReturn> {
  call(): Promise<TCallReturn>;
  call(options: CallOptions): Promise<TCallReturn>;
  call(
    options: CallOptions,
    callback: (error: Error, result: TCallReturn) => void
  ): Promise<TCallReturn>;
  encodeABI(): string;
}

export interface MethodReturnContext extends MethodPayableReturnContext {}

export type ContractContext = Web3ContractContext<
  IUniswapV3PoolState,
  IUniswapV3PoolStateMethodNames,
  IUniswapV3PoolStateEventsContext,
  IUniswapV3PoolStateEvents
>;
export type IUniswapV3PoolStateEvents = undefined;
export interface IUniswapV3PoolStateEventsContext {}
export type IUniswapV3PoolStateMethodNames =
  | 'feeGrowthGlobal0X128'
  | 'feeGrowthGlobal1X128'
  | 'liquidity'
  | 'observations'
  | 'positions'
  | 'protocolFees'
  | 'slot0'
  | 'tickBitmap'
  | 'ticks';
export interface ObservationsResponse {
  blockTimestamp: string;
  tickCumulative: string;
  secondsPerLiquidityCumulativeX128: string;
  initialized: boolean;
}
export interface PositionsResponse {
  _liquidity: string;
  feeGrowthInside0LastX128: string;
  feeGrowthInside1LastX128: string;
  tokensOwed0: string;
  tokensOwed1: string;
}
export interface ProtocolFeesResponse {
  token0: string;
  token1: string;
}
export interface Slot0Response {
  sqrtPriceX96: string;
  tick: string;
  observationIndex: string;
  observationCardinality: string;
  observationCardinalityNext: string;
  feeProtocol: string;
  unlocked: boolean;
}
export interface TicksResponse {
  liquidityGross: string;
  liquidityNet: string;
  feeGrowthOutside0X128: string;
  feeGrowthOutside1X128: string;
  tickCumulativeOutside: string;
  secondsPerLiquidityOutsideX128: string;
  secondsOutside: string;
  initialized: boolean;
}
export interface IUniswapV3PoolState {
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  feeGrowthGlobal0X128(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  feeGrowthGlobal1X128(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  liquidity(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param index Type: uint256, Indexed: false
   */
  observations(
    index: string
  ): MethodConstantReturnContext<ObservationsResponse>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param key Type: bytes32, Indexed: false
   */
  positions(
    key: string | number[]
  ): MethodConstantReturnContext<PositionsResponse>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  protocolFees(): MethodConstantReturnContext<ProtocolFeesResponse>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  slot0(): MethodConstantReturnContext<Slot0Response>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param wordPosition Type: int16, Indexed: false
   */
  tickBitmap(
    wordPosition: string | number
  ): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param tick Type: int24, Indexed: false
   */
  ticks(tick: string | number): MethodConstantReturnContext<TicksResponse>;
}
