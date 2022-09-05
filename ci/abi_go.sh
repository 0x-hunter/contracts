#!/bin/sh

cd abi
for ITEM in $(ls *.json)
do
    ../ci/abi_abi.js $ITEM ../go-contract/$ITEM
    abigen --abi=../go-contract/$ITEM --pkg=contract --out=../go-contract/${ITEM%.*}.go
    rm -rf ../go-contract/$ITEM
done
