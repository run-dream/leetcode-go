package problem1492

import "math"

func kthFactor(n int, k int) int {
	bound := int(math.Sqrt(float64(n)))
	for i := 1; i <= bound; i++ {
		if n%i == 0 {
			k--
			if k == 0 {
				return i
			}
		}
	}
	for i := bound; i >= 1; i-- {
		if n%(n/i) == 0 {
			k--
			if k == 0 {
				return n / i
			}
		}
	}
	return -1
}
