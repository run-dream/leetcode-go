package str2int

import "testing"

func TestSolve(t *testing.T) {
	result, err := Atoi("01")
	if err != nil {
		t.Log(err)
	}
	t.Log(result)
}
