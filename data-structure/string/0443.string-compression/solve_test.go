package problem0443

import "testing"

func TestSolve(t *testing.T) {
	bytes := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	t.Log(compress(bytes))
	t.Log(bytes)
}
