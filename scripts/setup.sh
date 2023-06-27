#!/bin/bash
scripts/reset.sh

test_mnemonic="wage thunder live sense resemble foil apple course spin horse glass mansion midnight laundry acoustic rhythm loan scale talent push green direct brick please"

simd init test --chain-id test
echo $test_mnemonic | simd keys add test --recover --keyring-backend test
simd genesis add-genesis-account test 100000000000000uxprt,100000000000000stake --keyring-backend test
simd genesis gentx test 10000000stake --chain-id test --keyring-backend test
simd genesis collect-gentxs
