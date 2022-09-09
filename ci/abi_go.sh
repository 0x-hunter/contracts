#!/bin/bash

cd abi

for ITEM in $(ls *.json)
do
    NAME=${ITEM%.*}
    ../ci/abi_abi.js $ITEM ./$NAME.abi.json
    PKG=$(echo $NAME | tr A-Z a-z)
    mkdir -p ../$PKG
    abigen --abi=./$NAME.abi.json --pkg=$PKG --out=../$PKG/$NAME.go
    rm -f ./$NAME.abi.json
done
