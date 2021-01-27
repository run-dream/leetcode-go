package problem0322

import "testing"

func TestCoinChange(t *testing.T) {
	t.Log(coinChange([]int{1, 2, 4}, 11))
	t.Log(coinChange([]int{2}, 3))
	t.Log(coinChange([]int{1}, 0))
	t.Log(coinChange([]int{1}, 1))
}
