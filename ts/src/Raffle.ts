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
  Raffle,
  RaffleMethodNames,
  RaffleEventsContext,
  RaffleEvents
>;
export type RaffleEvents =
  | 'FundsClaimed'
  | 'OwnershipTransferred'
  | 'PrizeClaimed'
  | 'RaffleCancelled'
  | 'RaffleClosed'
  | 'RaffleCreated'
  | 'RaffleEntered'
  | 'WinningTicketPicked';
export interface RaffleEventsContext {
  FundsClaimed(
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
  PrizeClaimed(
    parameters: {
      filter?: {};
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  RaffleCancelled(
    parameters: {
      filter?: {};
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  RaffleClosed(
    parameters: {
      filter?: {};
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  RaffleCreated(
    parameters: {
      filter?: { taskId?: string | string[]; owner?: string | string[] };
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  RaffleEntered(
    parameters: {
      filter?: {};
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
  WinningTicketPicked(
    parameters: {
      filter?: {};
      fromBlock?: number;
      toBlock?: 'latest' | number;
      topics?: string[];
    },
    callback?: (error: Error, event: EventData) => void
  ): EventResponse;
}
export type RaffleMethodNames =
  | 'new'
  | 'cancel'
  | 'claimFunds'
  | 'claimPrize'
  | 'close'
  | 'create'
  | 'currencyInfo'
  | 'currentTaskId'
  | 'enter'
  | 'feeRate'
  | 'feeSetter'
  | 'getCurrencyPrice'
  | 'onERC721Received'
  | 'owner'
  | 'prizeInfo'
  | 'raffleInfo'
  | 'rawFulfillRandomWords'
  | 'renounceOwnership'
  | 'requestInfo'
  | 'setCurrency'
  | 'transferOwnership'
  | 'userInfo';
export interface FundsClaimedEventEmittedResponse {
  taskId: string;
  fundsAmount: string;
  owner: string;
  fee: string;
}
export interface OwnershipTransferredEventEmittedResponse {
  previousOwner: string;
  newOwner: string;
}
export interface PrizeClaimedEventEmittedResponse {
  taskId: string;
  tokenId: string;
  winner: string;
  fee: string;
}
export interface RaffleCancelledEventEmittedResponse {
  taskId: string;
}
export interface RaffleClosedEventEmittedResponse {
  taskId: string;
}
export interface RaffleCreatedEventEmittedResponse {
  taskId: string;
  owner: string;
  startTime: string;
  duration: string;
}
export interface RaffleEnteredEventEmittedResponse {
  taskId: string;
  user: string;
  lowerBound: string;
  upperBound: string;
}
export interface WinningTicketPickedEventEmittedResponse {
  taskId: string;
  winningTicket: string;
}
export interface CurrencyInfoResponse {
  enabled: boolean;
  feeRate: string;
}
export interface PrizeInfoResponse {
  winningNumber: string;
  prizeClaimed: boolean;
  fundsClaimed: boolean;
}
export interface RaffleInfoResponse {
  owner: string;
  endTime: string;
  ticketSize: string;
  tokenId: string;
  tokenAddress: string;
  ticketPrice: string;
  funds: string;
  soldTickets: string;
  currency: string;
  status: string;
}
export interface Raffle {
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: constructor
   * @param _vrfCoordinator Type: address, Indexed: false
   * @param _feeSetter Type: address, Indexed: false
   * @param _usd Type: address, Indexed: false
   * @param _keyHash Type: bytes32, Indexed: false
   * @param _subscriptionId Type: uint64, Indexed: false
   * @param _uniV3Factory Type: address, Indexed: false
   * @param _currencies Type: address[], Indexed: false
   * @param _feeRates Type: uint24[], Indexed: false
   */
  'new'(
    _vrfCoordinator: string,
    _feeSetter: string,
    _usd: string,
    _keyHash: string | number[],
    _subscriptionId: string,
    _uniV3Factory: string,
    _currencies: string[],
    _feeRates: string | number[]
  ): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param _raffleId Type: uint256, Indexed: false
   */
  cancel(_raffleId: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param _raffleId Type: uint256, Indexed: false
   */
  claimFunds(_raffleId: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param _raffleId Type: uint256, Indexed: false
   */
  claimPrize(_raffleId: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param _raffleId Type: uint256, Indexed: false
   */
  close(_raffleId: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param _tokenAddr Type: address, Indexed: false
   * @param _tokenId Type: uint256, Indexed: false
   * @param _duration Type: uint256, Indexed: false
   * @param _ticketPrice Type: uint64, Indexed: false
   * @param _ticketSize Type: uint32, Indexed: false
   * @param _currency Type: address, Indexed: false
   */
  create(
    _tokenAddr: string,
    _tokenId: string,
    _duration: string,
    _ticketPrice: string,
    _ticketSize: string | number,
    _currency: string
  ): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param parameter0 Type: address, Indexed: false
   */
  currencyInfo(
    parameter0: string
  ): MethodConstantReturnContext<CurrencyInfoResponse>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  currentTaskId(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param _raffleId Type: uint256, Indexed: false
   * @param _amount Type: uint96, Indexed: false
   */
  enter(_raffleId: string, _amount: string): MethodReturnContext;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  feeRate(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   */
  feeSetter(): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param _currency Type: address, Indexed: false
   */
  getCurrencyPrice(_currency: string): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: pure
   * Type: function
   * @param parameter0 Type: address, Indexed: false
   * @param parameter1 Type: address, Indexed: false
   * @param parameter2 Type: uint256, Indexed: false
   * @param parameter3 Type: bytes, Indexed: false
   */
  onERC721Received(
    parameter0: string,
    parameter1: string,
    parameter2: string,
    parameter3: string | number[]
  ): MethodConstantReturnContext<string>;
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
   * @param parameter0 Type: uint256, Indexed: false
   */
  prizeInfo(parameter0: string): MethodConstantReturnContext<PrizeInfoResponse>;
  /**
   * Payable: false
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param parameter0 Type: uint256, Indexed: false
   */
  raffleInfo(
    parameter0: string
  ): MethodConstantReturnContext<RaffleInfoResponse>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param requestId Type: uint256, Indexed: false
   * @param randomWords Type: uint256[], Indexed: false
   */
  rawFulfillRandomWords(
    requestId: string,
    randomWords: string[]
  ): MethodReturnContext;
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
   * @param parameter0 Type: uint256, Indexed: false
   */
  requestInfo(parameter0: string): MethodConstantReturnContext<string>;
  /**
   * Payable: false
   * Constant: false
   * StateMutability: nonpayable
   * Type: function
   * @param _currency Type: address, Indexed: false
   * @param _enable Type: bool, Indexed: false
   * @param _feeRate Type: uint24, Indexed: false
   */
  setCurrency(
    _currency: string,
    _enable: boolean,
    _feeRate: string | number
  ): MethodReturnContext;
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
   * Constant: true
   * StateMutability: view
   * Type: function
   * @param parameter0 Type: uint256, Indexed: false
   * @param parameter1 Type: address, Indexed: false
   */
  userInfo(
    parameter0: string,
    parameter1: string
  ): MethodConstantReturnContext<string>;
}
