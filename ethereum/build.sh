#!/bin/bash

rm -rf build
mkdir -p build/locker build/aesnft

# compile
solc --abi --bin ../../aeschain-ethereum/flatten/LockerFlat.sol -o build/locker
solc --abi --bin ../../aeschain-ethereum/flatten/AesNFTFlat.sol -o build/aesnft

# build
abigen --abi=./build/locker/Locker.abi --pkg=locker --out=./contracts/locker/Locker.go
abigen --abi=./build/aesnft/AesNFT.abi --pkg=aesnft --out=./contracts/aesnft/AesNFT.go