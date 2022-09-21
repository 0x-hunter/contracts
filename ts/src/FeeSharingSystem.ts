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
  FeeSharingSystem,
  FeeSharingSystemMethodNames,
  FeeSharingSystemEventsContext,
  FeeSharingSystemEvents
>;
export type FeeSharingSystemEvents =
  | 'Deposit'
  | 'Harvest'
  | 'NewRewardPeriod'
  | 'OwnershipTransferred'
  | 'Withdraw';
export interface FeeSharingSystemEventsContext {
  Deposit(
    parameters: {
      filter?: { user?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  Harvest(
    parameters: {
      filter?: { user?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  NewRewardPeriod(
    parameters: {
      filter?: {};
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  OwnershipTransferred(
    parameters: {
      filter?: {
        previousOwner?: string | string[];
        newOwner?: string | string[];
      };
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
export type FeeSharingSystemMethodNames =
  | 'new'
  | 'PRECISION_FACTOR'
  | 'calculatePendingRewards'
  | 'calculateSharePriceInLOOKS'
  | 'calculateSharesValueInLOOKS'
  | 'currentRewardPerBlock'
  | 'deposit'
  | 'harvest'
  | 'hunterGovernanceToken'
  | 'lastRewardAdjustment'
  | 'lastRewardBlock'
  | 'lastUpdateBlock'
  | 'owner'
  | 'periodEndBlock'
  | 'renounceOwnership'
  | 'rewardPerTokenStored'
  | 'rewardToken'
  | 'tokenDistributor'
  | 'totalShares'
  | 'transferOwnership'
  | 'updateRewards'
  | 'userInfo'
  | 'withdraw'
  | 'withdrawAll';
export interface DepositEventEmittedResponse {
  user: string;
  amount: string;
  harvestedAmount: string;
}
export interface HarvestEventEmittedResponse {
  user: string;
  harvestedAmount: string;
}
export interface NewRewardPeriodEventEmittedResponse {
  numberBlocks: string;
  rewardPerBlock: string;
  reward: string;
}
export interface OwnershipTransferredEventEmittedResponse {
  previousOwner: string;
  newOwner: string;
}
export interface WithdrawEventEmittedResponse {
  user: string;
  amount: string;
  harvestedAmount: string;
}
export interface UserInfoResponse {
  shares: string;
  userRewardPerTokenPaid: string;
  rewards: string;
}
export interface FeeSharingSystem {
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: constructor
   * @param _looksRareToken Type: address, Indexed: false
   * @param _rewardToken Type: address, Indexed: false
   * @param _tokenDistributor Type: address, Indexed: false
   */
  'new'(
    _looksRareToken: string,
    _rewardToken: string,
    _tokenDistributor: string
  ): MethodReturnContext;
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
   * @param user Type: address, Indexed: false
   */
  calculatePendingRewards(user: string): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  calculateSharePriceInLOOKS(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param user Type: address, Indexed: false
   */
  calculateSharesValueInLOOKS(
    user: string
  ): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  currentRewardPerBlock(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param amount Type: uint256, Indexed: false
   * @param claimRewardToken Type: bool, Indexed: false
   */
  deposit(amount: string, claimRewardToken: boolean): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   */
  harvest(): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  hunterGovernanceToken(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  lastRewardAdjustment(): MethodConstantReturnContext<string>;
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
  lastUpdateBlock(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  owner(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  periodEndBlock(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   */
  renounceOwnership(): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  rewardPerTokenStored(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  rewardToken(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  tokenDistributor(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  totalShares(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param newOwner Type: address, Indexed: false
   */
  transferOwnership(newOwner: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param reward Type: uint256, Indexed: false
   * @param rewardDurationInBlocks Type: uint256, Indexed: false
   */
  updateRewards(
    reward: string,
    rewardDurationInBlocks: string
  ): MethodReturnContext;
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
   * @param shares Type: uint256, Indexed: false
   * @param claimRewardToken Type: bool, Indexed: false
   */
  withdraw(shares: string, claimRewardToken: boolean): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param claimRewardToken Type: bool, Indexed: false
   */
  withdrawAll(claimRewardToken: boolean): MethodReturnContext;
}
