package problem0078

import "testing"

func TestSubset(t *testing.T) {
	t.Log(subsets([]int{}))
	t.Log(subsets([]int{1, 2, 3}))
}
