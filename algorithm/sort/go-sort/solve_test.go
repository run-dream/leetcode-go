package sort_test

import (
	"sort"
	"testing"
)

type Node struct {
	key1 int
	key2 int
}

func TestSort(t *testing.T) {
	nums := []int{2, 3, 4}
	sort.Ints(nums)
	t.Log(nums)

	strs := []string{"a", "b", "c"}
	sort.Strings(strs)
	t.Log(strs)

	nodes := []Node{
		Node{2, 2},
		Node{1, 3},
		Node{1, 2},
	}
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].key1 == nodes[j].key1 {
			return nodes[i].key2 < nodes[j].key2
		}
		return nodes[i].key1 < nodes[j].key1
	})
	t.Log(nodes)
}
