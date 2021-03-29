package problem0871

import "container/heap"

func minRefuelStopsII(target int, startFuel int, stations [][]int) int {
	mh := &maxHeap{}
	heap.Init(mh)
	ans := 0
	tank := startFuel
	prev := 0
	for i := 0; i < len(stations); i++ {
		tank -= (stations[i][0] - prev)
		prev = stations[i][0]

		for mh.Len() > 0 && tank < 0 {
			fuel := heap.Pop(mh).(int)
			tank += fuel
			ans++
		}

		if tank < 0 {
			return -1
		}

		heap.Push(mh, stations[i][1])
	}

	// see if the target can be reached
	{
		tank -= target - prev
		for mh.Len() > 0 && tank < 0 {
			tank += heap.Pop(mh).(int)
			ans++
		}

		if tank < 0 {
			return -1
		}
	}

	return ans
}

type maxHeap []int

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h maxHeap) Len() int {
	return len(h)
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() (v interface{}) {
	v, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return v
}

func (h *maxHeap) Top() int {
	return (*h)[0]
}
