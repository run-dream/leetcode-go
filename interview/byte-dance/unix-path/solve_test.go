package unix_path

import "testing"

func TestSolve(t *testing.T) {
	t.Log(simplifyPath("/home/foo/./.././bar"))
	t.Log(simplifyPathII("/home/foo/./.././bar"))
}
