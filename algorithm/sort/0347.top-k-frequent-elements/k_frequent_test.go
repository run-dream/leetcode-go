package problem0347

import "testing"

func TestKFrequent(t *testing.T) {
	t.Log(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	t.Log(topKFrequentHeap([]int{1, 1, 1, 2, 2, 3}, 2))
}
