package problem0405

import "testing"

func TestToHex(t *testing.T) {
	t.Log(toHex(0))
	t.Log(toHex(-1))
	t.Log(toHex(26))
}
