package problem0093

import "testing"

func TestRestoreIpAddress(t *testing.T) {
	t.Log(restoreIpAddresses("0000"))
	t.Log(restoreIpAddresses("25525511135"))
	t.Log(restoreIpAddresses("1111"))
	t.Log(restoreIpAddresses("010010"))
	t.Log(restoreIpAddresses("101023"))
}
