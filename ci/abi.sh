#!/bin/sh

abi-types-generator abi/Raffle.json --output ts/src
abi-types-generator abi/MockERC20.json --output ts/src
abi-types-generator abi/MockERC721.json --output ts/src
