package problem0204

func countPrimes(n int) int {
	notPrimes := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if notPrimes[i] {
			continue
		}
		count++
		for j := i * i; j < n; j += i {
			notPrimes[j] = true
		}
	}
	return count
}
