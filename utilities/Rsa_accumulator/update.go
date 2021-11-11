package Rsa_accumulator

import (
	"math/big"
)
/*
	Adding new member to the set which autometically precomputes the all the Witnesses in O(n) time
*/
func (c *Rsa_Acc) Add_member(u big.Int, w *Witness_list) {

	e := Hprime(u)
	preAcc := c.Acc

	newAcc := new(big.Int).Exp(&c.Acc, &e, &c.N)
	newSet := append(c.U[:], u)

	c.Acc = *newAcc
	c.U = newSet

	if len(w.List) == 0 {
		w.Precompute_witness(c.G, c.U, c)
	} else {
		for _, x := range c.U {
			temp := w.List[x.String()]
			w.List[x.String()] = *new(big.Int).Exp(&temp, &e, &c.N)
		}
		w.List[u.String()] = preAcc

	}

}
/*
	Deleting a member from the set in O(nlogn) time
*/
func (c *Rsa_Acc) Delete_member(u big.Int, w *Witness_list) {

	var newSet []big.Int
	var i int
	for i = 0; i < len(c.U); i++ {
		if c.U[i].Cmp(&u) == 0 {
			newSet = append(c.U[:i], c.U[i+1:]...)
			break
		}
	}

	newAcc := w.List[u.String()]
	c.Acc = newAcc
	c.U = newSet
	list := make(map[string]big.Int, len(c.U))
	w.List = list
	w.Precompute_witness(c.G, c.U, c)

}
