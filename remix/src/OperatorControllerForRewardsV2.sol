// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";


/**
 * @title OperatorControllerForRewardsV2
 * @notice It splits pending HTGs and updates trading rewards.
 */
contract OperatorControllerForRewardsV2 is Ownable {


    /**
     * @notice Constructor
     * @param _feeSharingSystem address of the fee sharing system contract
     * @param _feeSharingSetter address of the fee sharing setter contract
     * @param _tokenSplitter address of the token splitter contract
     * @param _teamVesting address of the team vesting contract
     * @param _treasuryVesting address of the treasury vesting contract
     * @param _tradingRewardsDistributor address of the trading rewards distributor contract
     */
    constructor(
        address _feeSharingSystem,
        address _feeSharingSetter,
        address _tokenSplitter,
        address _teamVesting,
        address _treasuryVesting,
        address _tradingRewardsDistributor
    ) {
        
    }

    /**
     * @notice Release HTGs tokens from the TokenSplitter and update fee-sharing rewards
     */
    function releaseTokensAndUpdateRewards() external {// onlyOwner {
        
    }

    /**
     * @notice It verifies that the lastUpdateBlock is greater than endBlock
     */
    function canRelease() public view returns (bool) {
    
    }
}
