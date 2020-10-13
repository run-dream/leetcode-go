package problem0347

import (
	"container/heap"
)

func topKFrequentHeap(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, n := range nums {
		counts[n]++
	}

	pq := &priorityQueue{}
	heap.Init(pq)
	for k, v := range counts {
		heap.Push(pq, num{
			value: k,
			count: v,
		})
	}
	ans := []int{}
	for i := 0; i < k; i++ {
		n := heap.Pop(pq).(num)
		ans = append(ans, n.value)
	}
	return ans
}

type num struct {
	value int
	count int
}

type priorityQueue []num

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	num := x.(num)
	*pq = append(*pq, num)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	num := old[n-1]
	*pq = old[0 : n-1]
	return num
}
