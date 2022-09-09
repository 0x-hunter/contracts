#!/bin/bash

cd abi

for ITEM in $(ls *.json)
do
    ../ci/abi_abi.js $ITEM ../go-contract/$ITEM
    NAME=${ITEM%.*}
    PKG=$(echo $NAME | tr A-Z a-z)
    mkdir -p ../go-contract/$PKG
    abigen --abi=../go-contract/$ITEM --pkg=$PKG --out=../go-contract/$PKG/$NAME.go
    rm -rf ../go-contract/$ITEM
done
