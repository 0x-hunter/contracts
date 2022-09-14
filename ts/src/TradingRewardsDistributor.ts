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
  TradingRewardsDistributor,
  TradingRewardsDistributorMethodNames,
  TradingRewardsDistributorEventsContext,
  TradingRewardsDistributorEvents
>;
export type TradingRewardsDistributorEvents =
  | 'OwnershipTransferred'
  | 'Paused'
  | 'RewardsClaim'
  | 'TokenWithdrawnOwner'
  | 'Unpaused'
  | 'UpdateTradingRewards';
export interface TradingRewardsDistributorEventsContext {
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
  Paused(
    parameters: {
      filter?: {};
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  RewardsClaim(
    parameters: {
      filter?: { user?: string | string[]; rewardRound?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  TokenWithdrawnOwner(
    parameters: {
      filter?: {};
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  Unpaused(
    parameters: {
      filter?: {};
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  UpdateTradingRewards(
    parameters: {
      filter?: { rewardRound?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
}
export type TradingRewardsDistributorMethodNames =
  | 'new'
  | 'BUFFER_ADMIN_WITHDRAW'
  | 'amountClaimedByUser'
  | 'canClaim'
  | 'claim'
  | 'currentRewardRound'
  | 'hasUserClaimedForRewardRound'
  | 'hunterGovernanceToken'
  | 'lastPausedTimestamp'
  | 'maximumAmountPerUserInCurrentTree'
  | 'merkleRootOfRewardRound'
  | 'merkleRootUsed'
  | 'owner'
  | 'pauseDistribution'
  | 'paused'
  | 'renounceOwnership'
  | 'transferOwnership'
  | 'unpauseDistribution'
  | 'updateTradingRewards'
  | 'withdrawTokenRewards';
export interface OwnershipTransferredEventEmittedResponse {
  previousOwner: string;
  newOwner: string;
}
export interface PausedEventEmittedResponse {
  account: string;
}
export interface RewardsClaimEventEmittedResponse {
  user: string;
  rewardRound: string;
  amount: string;
}
export interface TokenWithdrawnOwnerEventEmittedResponse {
  amount: string;
}
export interface UnpausedEventEmittedResponse {
  account: string;
}
export interface UpdateTradingRewardsEventEmittedResponse {
  rewardRound: string;
}
export interface CanClaimResponse {
  result0: boolean;
  result1: string;
}
export interface TradingRewardsDistributor {
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: constructor
   * @param hunterGovernanceToken_ Type: address, Indexed: false
   */
  'new'(hunterGovernanceToken_: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  BUFFER_ADMIN_WITHDRAW(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param parameter0 Type: address, Indexed: false
   */
  amountClaimedByUser(parameter0: string): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param user Type: address, Indexed: false
   * @param amount Type: uint256, Indexed: false
   * @param merkleProof Type: bytes32[], Indexed: false
   */
  canClaim(
    user: string,
    amount: string,
    merkleProof: string | number[][]
  ): MethodConstantReturnContext<CanClaimResponse>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param amount Type: uint256, Indexed: false
   * @param merkleProof Type: bytes32[], Indexed: false
   */
  claim(amount: string, merkleProof: string | number[][]): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  currentRewardRound(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param parameter0 Type: uint256, Indexed: false
   * @param parameter1 Type: address, Indexed: false
   */
  hasUserClaimedForRewardRound(
    parameter0: string,
    parameter1: string
  ): MethodConstantReturnContext<boolean>;
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
  lastPausedTimestamp(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  maximumAmountPerUserInCurrentTree(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param parameter0 Type: uint256, Indexed: false
   */
  merkleRootOfRewardRound(
    parameter0: string
  ): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param parameter0 Type: bytes32, Indexed: false
   */
  merkleRootUsed(
    parameter0: string | number[]
  ): MethodConstantReturnContext<boolean>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  owner(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   */
  pauseDistribution(): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  paused(): MethodConstantReturnContext<boolean>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   */
  renounceOwnership(): MethodReturnContext;
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
   */
  unpauseDistribution(): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param merkleRoot Type: bytes32, Indexed: false
   * @param newMaximumAmountPerUser Type: uint256, Indexed: false
   */
  updateTradingRewards(
    merkleRoot: string | number[],
    newMaximumAmountPerUser: string
  ): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param amount Type: uint256, Indexed: false
   */
  withdrawTokenRewards(amount: string): MethodReturnContext;
}
