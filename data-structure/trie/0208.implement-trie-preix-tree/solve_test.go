package problem0208

import "testing"

func TestSolve(t *testing.T) {
	trie := Constructor()
	ops := []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"}
	vals := [][]string{
		[]string{},
		[]string{"apple"},
		[]string{"apple"},
		[]string{"app"},
		[]string{"app"},
		[]string{"app"},
		[]string{"app"},
	}
	for i, op := range ops {
		val := vals[i]
		if op == "insert" {
			trie.Insert(val[0])
		} else if op == "search" {
			t.Log(trie.Search(val[0]))
		} else if op == "startsWith" {
			t.Log(trie.StartsWith(val[0]))
		}
	}
}
