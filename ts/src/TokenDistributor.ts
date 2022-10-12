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
  TokenDistributor,
  TokenDistributorMethodNames,
  TokenDistributorEventsContext,
  TokenDistributorEvents
>;
export type TokenDistributorEvents =
  | 'Compound'
  | 'Deposit'
  | 'NewRewardsPerBlock'
  | 'Withdraw';
export interface TokenDistributorEventsContext {
  Compound(
    parameters: {
      filter?: { user?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  Deposit(
    parameters: {
      filter?: { user?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  NewRewardsPerBlock(
    parameters: {
      filter?: { currentPhase?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  Withdraw(
    parameters: {
      filter?: { user?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
}
export type TokenDistributorMethodNames =
  | 'new'
  | 'NUMBER_PERIODS'
  | 'PRECISION_FACTOR'
  | 'START_BLOCK'
  | 'accTokenPerShare'
  | 'calculatePendingRewards'
  | 'currentPhase'
  | 'deposit'
  | 'endBlock'
  | 'harvestAndCompound'
  | 'lastRewardBlock'
  | 'looksRareToken'
  | 'rewardPerBlockForOthers'
  | 'rewardPerBlockForStaking'
  | 'stakingPeriod'
  | 'tokenSplitter'
  | 'totalAmountStaked'
  | 'updatePool'
  | 'userInfo'
  | 'withdraw'
  | 'withdrawAll';
export interface CompoundEventEmittedResponse {
  user: string;
  harvestedAmount: string;
}
export interface DepositEventEmittedResponse {
  user: string;
  amount: string;
  harvestedAmount: string;
}
export interface NewRewardsPerBlockEventEmittedResponse {
  currentPhase: string;
  startBlock: string;
  rewardPerBlockForStaking: string;
  rewardPerBlockForOthers: string;
}
export interface WithdrawEventEmittedResponse {
  user: string;
  amount: string;
  harvestedAmount: string;
}
export interface StakingPeriodResponse {
  rewardPerBlockForStaking: string;
  rewardPerBlockForOthers: string;
  periodLengthInBlock: string;
}
export interface UserInfoResponse {
  amount: string;
  rewardDebt: string;
}
export interface TokenDistributor {
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: constructor
   * @param _looksRareToken Type: address, Indexed: false
   * @param _tokenSplitter Type: address, Indexed: false
   * @param _startBlock Type: uint256, Indexed: false
   * @param _rewardsPerBlockForStaking Type: uint256[], Indexed: false
   * @param _rewardsPerBlockForOthers Type: uint256[], Indexed: false
   * @param _periodLengthesInBlocks Type: uint256[], Indexed: false
   * @param _numberPeriods Type: uint256, Indexed: false
   */
  'new'(
    _looksRareToken: string,
    _tokenSplitter: string,
    _startBlock: string,
    _rewardsPerBlockForStaking: string[],
    _rewardsPerBlockForOthers: string[],
    _periodLengthesInBlocks: string[],
    _numberPeriods: string
  ): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  NUMBER_PERIODS(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  PRECISION_FACTOR(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  START_BLOCK(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  accTokenPerShare(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param user Type: address, Indexed: false
   */
  calculatePendingRewards(user: string): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  currentPhase(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param amount Type: uint256, Indexed: false
   */
  deposit(amount: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  endBlock(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   */
  harvestAndCompound(): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  lastRewardBlock(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  looksRareToken(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  rewardPerBlockForOthers(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  rewardPerBlockForStaking(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param parameter0 Type: uint256, Indexed: false
   */
  stakingPeriod(
    parameter0: string
  ): MethodConstantReturnContext<StakingPeriodResponse>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  tokenSplitter(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  totalAmountStaked(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   */
  updatePool(): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param parameter0 Type: address, Indexed: false
   */
  userInfo(parameter0: string): MethodConstantReturnContext<UserInfoResponse>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param amount Type: uint256, Indexed: false
   */
  withdraw(amount: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   */
  withdrawAll(): MethodReturnContext;
}
