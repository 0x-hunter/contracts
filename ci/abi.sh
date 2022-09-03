#!/bin/sh

cd abi
for ITEM in $(ls *.json)
do
    abi-types-generator $ITEM --output ../ts/src --provider=web3
    ../ci/abi_copy.js $ITEM ../ts/src
done

exit 0

abi-types-generator abi/Raffle.json --output ts/src --provider=web3
abi-types-generator abi/MockERC20.json --output ts/src --provider=web3
abi-types-generator abi/MockERC721.json --output ts/src --provider=web3
