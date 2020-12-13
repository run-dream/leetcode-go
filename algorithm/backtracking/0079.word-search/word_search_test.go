package problem0079

import "testing"

func TestWordSearch(t *testing.T) {
	/*board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	t.Log(exist(board, "ABC"))
	t.Log(exist(board, "ASAD"))
	t.Log(exist(board, "FCS"))
	t.Log(exist(board, "ABES"))
	*/
	//t.Log(exist([][]byte{[]byte{'a'}}, "a"))
	t.Log(exist(
		[][]byte{
			[]byte{'A', 'B', 'C', 'E'},
			[]byte{'S', 'F', 'E', 'S'},
			[]byte{'A', 'D', 'E', 'E'},
		},
		"ABCESEEEFS",
	))
}
