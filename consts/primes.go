package consts

import (
	"log"
	"math/big"
)

func Primes(n int) []int64 {
	var primes []int64
	one := big.NewInt(1)
	for i := big.NewInt(2); len(primes) < n; i.Add(i, one) {
		if i.ProbablyPrime(0) {
			if !i.IsInt64() {
				log.Fatal("an error has occured")
			}
			primes = append(primes, i.Int64())
		}
	}
	return primes
}
