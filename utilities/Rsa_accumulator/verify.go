package Rsa_accumulator

import (
	"math/big"
)

/*
Verification is simply
W^e (mod N) == Acc
u and W are coming from the prover
Accumulator and N are stored on chain
Therefore, args[] should be constructed on chain
*/

func Verify(args []big.Int) bool {

	u, W, Accumulator, N := args[0], args[1], args[2], args[3]
	e := Hprime(u)
	Acc_dash := new(big.Int).Exp(&W, &e, &N)
	return Acc_dash.Cmp(&Accumulator) == 0

}
