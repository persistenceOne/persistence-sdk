package Rsa_accumulator

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func TestGenerate_acc(t *testing.T) {
	//hash2prime-----------------------------------------------------------------------
	H, _ := rand.Int(rand.Reader, big.NewInt(512))
	E := Hprime(*H)

	if !E.ProbablyPrime(10) {
		t.Fatalf("Not a prime")
	}
	//generate_acc---------------------------------------------------------------------
	key := Rsa_keygen(int(32))
	if key.G.Cmp(&key.N) == 1 && key.N.BitLen() == 32 {
		t.Fatalf("Incorrect generator")
	}

	U := []big.Int{*big.NewInt(123), *big.NewInt(124), *big.NewInt(125), *big.NewInt(126)}
	Accumulator := Generate_Acc(key, U)
	G := key.G
	for _, u := range U {
		x := Hprime(u)
		G.Exp(&G, &x, &key.N)
	}
	if G.Cmp(&Accumulator.Acc) != 0 {
		t.Fatalf("Incorrect accumulator")
	}

	//witness---------------------------------------------------------------------------
	list := make(map[string]big.Int, len(Accumulator.U))
	w := &Witness_list{Acc: Accumulator.Acc, List: list}
	w.Precompute_witness(Accumulator.G, Accumulator.U, Accumulator)

	witnesses := make([]big.Int, len(Accumulator.U))
	for i, u := range Accumulator.U {
		witnesses[i] = generate_witness(u, key, Accumulator.U)
	}

	for i, u := range U {
		witness := w.List[u.String()]
		if witnesses[i].Cmp(&witness) != 0 {
			t.Fatalf("Incorrect witness")
		}
	}

	//Update(Add)------------------------------------------------------------------------------
	Accumulator_prev1 := Accumulator.Acc
	Accumulator.Add_member(*big.NewInt(127), w)
	e := Hprime(*big.NewInt(127))
	Accumulator_prev1.Exp(&Accumulator_prev1, &e, &key.N)
	if Accumulator_prev1.Cmp(&Accumulator.Acc) != 0 {
		t.Fatalf("Incorrect value of Accumulator after update")
	}

	U = append(U[:], *big.NewInt(127))
	for i, u := range U {
		if u.Cmp(&Accumulator.U[i]) != 0 {
			t.Fatalf("Incorrect update of set")
		}
	}
	u := big.NewInt(127)
	W1 := generate_witness(*big.NewInt(127), key, U)
	W2 := w.List[u.String()]
	if W1.Cmp(&W2) != 0 {
		t.Fatalf("Incorrect witness computation after update")
	}
	//Update(Delete)------------------------------------------------------------------------------
	Accumulator.Delete_member(*big.NewInt(126), w)
	Accumulator_prev2 := Generate_Acc(key, Accumulator.U)
	if Accumulator_prev2.Acc.Cmp(&Accumulator.Acc) != 0 {
		t.Fatalf("Incorrect value of Accumulator after delete")
	}

	W3 := generate_witness(*big.NewInt(127), key, Accumulator.U)
	W4 := w.List[u.String()]
	if W4.Cmp(&W3) != 0 {
		t.Fatalf("Incorrect witness computation after update")
	}

}

//Generation of witness is multiplication of all primes mapped from members except the one we
//are proving,prod(say) then,
// Witness = G^prod(mod N)
func generate_witness(u big.Int, key Rsa_key, U []big.Int) big.Int {

	N := key.N

	Primes := make([]big.Int, len(U))
	G := key.G
	for i, u_dash := range U {
		Primes[i] = Hprime(u_dash)
		if u_dash.Cmp(&u) != 0 {

			G.Exp(&G, &Primes[i], &N)
		}

	}

	return G
}
