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
  TokenSplitter,
  TokenSplitterMethodNames,
  TokenSplitterEventsContext,
  TokenSplitterEvents
>;
export type TokenSplitterEvents =
  | 'NewSharesOwner'
  | 'OwnershipTransferred'
  | 'TokensTransferred';
export interface TokenSplitterEventsContext {
  NewSharesOwner(
    parameters: {
      filter?: {
        oldRecipient?: string | string[];
        newRecipient?: string | string[];
      };
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
  TokensTransferred(
    parameters: {
      filter?: { account?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
}
export type TokenSplitterMethodNames =
  | 'new'
  | 'TOTAL_SHARES'
  | 'accountInfo'
  | 'calculatePendingRewards'
  | 'hunterGovernanceToken'
  | 'owner'
  | 'releaseTokens'
  | 'renounceOwnership'
  | 'totalTokensDistributed'
  | 'transferOwnership'
  | 'updateSharesOwner';
export interface NewSharesOwnerEventEmittedResponse {
  oldRecipient: string;
  newRecipient: string;
}
export interface OwnershipTransferredEventEmittedResponse {
  previousOwner: string;
  newOwner: string;
}
export interface TokensTransferredEventEmittedResponse {
  account: string;
  amount: string;
}
export interface AccountInfoResponse {
  shares: string;
  tokensDistributedToAccount: string;
}
export interface TokenSplitter {
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: constructor
   * @param _accounts Type: address[], Indexed: false
   * @param _shares Type: uint256[], Indexed: false
   * @param _hunterGovernanceToken Type: address, Indexed: false
   */
  'new'(
    _accounts: string[],
    _shares: string[],
    _hunterGovernanceToken: string
  ): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  TOTAL_SHARES(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param parameter0 Type: address, Indexed: false
   */
  accountInfo(
    parameter0: string
  ): MethodConstantReturnContext<AccountInfoResponse>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param account Type: address, Indexed: false
   */
  calculatePendingRewards(account: string): MethodConstantReturnContext<string>;
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
  owner(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param account Type: address, Indexed: false
   */
  releaseTokens(account: string): MethodReturnContext;
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
  totalTokensDistributed(): MethodConstantReturnContext<string>;
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
   * @param _newRecipient Type: address, Indexed: false
   * @param _currentRecipient Type: address, Indexed: false
   */
  updateSharesOwner(
    _newRecipient: string,
    _currentRecipient: string
  ): MethodReturnContext;
}
