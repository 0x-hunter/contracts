
import * as ERC from "./OperatorControllerForRewardsV2";
import Web3 from "web3"

export default function(w3: Web3, addr: string): ERC.OperatorControllerForRewardsV2 {
    let contract = new w3.eth.Contract([{"inputs":[{"internalType":"address","name":"_feeSharingSystem","type":"address"},{"internalType":"address","name":"_feeSharingSetter","type":"address"},{"internalType":"address","name":"_tokenSplitter","type":"address"},{"internalType":"address","name":"_teamVesting","type":"address"},{"internalType":"address","name":"_treasuryVesting","type":"address"},{"internalType":"address","name":"_tradingRewardsDistributor","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"inputs":[],"name":"canRelease","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"releaseTokensAndUpdateRewards","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"}] as any, addr)
    return contract.methods as any as ERC.OperatorControllerForRewardsV2;
}
