package main

import (
	"fmt"
	"math/big"
	"strconv"
)

const ModNumber = 1000000007
const MaximumSize = 1001

var factorials = make([]*big.Int, MaximumSize)

func Fact(n uint64) *big.Int {
	if n == 1 {
		return new(big.Int).SetUint64(1)
	}
	return new(big.Int).Mul(new(big.Int).SetUint64(n), factorials[n-1])
}

/*
	Ref: https://en.wikipedia.org/wiki/Lattice_path
	Ref: https://en.wikipedia.org/wiki/Binomial_theorem
	Ref: https://www.wolframalpha.com/input/?i=500%21+mod+%2810%5E9+%2B+7%29
	Binomial coefficients: (n, k) = n! / [k!(n - k)!] = n! / [k! * d!]
	Counting Lattice Paths: (w + h, w)
	D&Q with %(10^9 + 7)
*/
func CountingLatticePaths(m, n uint) uint64 {
	divider := new(big.Int).Mul(factorials[n], factorials[m])
	resultBeforeMod := new(big.Int).Div(factorials[n+m], divider)
	result := new(big.Int).Mod(resultBeforeMod, new(big.Int).SetUint64(ModNumber))
	latticePaths, _ := strconv.ParseUint(result.String(), 10, 64)
	return latticePaths
}

func main() {
	for i := uint64(1); i < MaximumSize; i++ {
		factorials[i] = Fact(i)
	}
	var t, n, m uint
	fmt.Scanf("%d", &t)
	for i := uint(0); i < t; i++ {
		fmt.Scanf("%d %d", &n, &m)
		fmt.Println(CountingLatticePaths(n, m))
	}
}
