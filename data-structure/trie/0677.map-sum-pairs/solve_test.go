package problem0677

import "testing"

func TestSolve(t *testing.T) {
	mapSum := Constructor()
	ops := []string{"insert", "sum", "insert", "sum"}
	strs := []string{"apple", "ap", "app", "ap"}
	vals := []int{3, 0, 2}
	for i := 0; i < len(ops); i++ {
		op := ops[i]
		if op == "insert" {
			mapSum.Insert(strs[i], vals[i])
		} else if op == "sum" {
			t.Log(mapSum.Sum(strs[i]))
		}
	}
}
