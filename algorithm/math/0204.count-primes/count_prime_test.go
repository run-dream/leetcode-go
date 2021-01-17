package problem0204

import "testing"

func TestCountPrime(t *testing.T) {
	t.Log(countPrimes(10))
	t.Log(countPrimes(0))
	t.Log(countPrimes(1))
}
