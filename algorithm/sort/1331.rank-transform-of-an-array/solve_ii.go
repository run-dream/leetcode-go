package problem1331

import "sort"

func arrayRankTransformII(arr []int) []int {
	g := make([]int, len(arr))
	copy(g, arr)
	m := make(map[int]int)
	sort.Ints(g)
	counter := 1
	for i := range g {
		if i >= 1 && g[i] == g[i-1] {
			counter -= 1
		}
		m[g[i]] = counter
		counter++
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = m[arr[i]]
	}
	return arr
}
