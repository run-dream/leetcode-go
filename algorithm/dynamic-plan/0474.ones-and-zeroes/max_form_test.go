package problem0474

import "testing"

func TestMaxForm(t *testing.T) {
	t.Log(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	t.Log(findMaxForm([]string{"10", "1", "0"}, 1, 1))
}
