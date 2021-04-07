package problem0300

import "testing"

func TestLIS(t *testing.T) {
	//t.Log(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	t.Log(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	t.Log(lengthOfLISII([]int{0, 1, 0, 3, 2, 3}))
	//t.Log(lengthOfLIS([]int{7, 7, 7, 7, 7}))
}
