package trie

import "testing"

func TestTrie(t *testing.T) {
	trie := Trie{NewTrieNode(byte(' '), false)}
	strs := []string{"hello", "hi", "hey", "how", "so", "see"}
	for _, str := range strs {
		trie.Insert(str)
	}
	trie.Insert("hell")
	t.Log(trie.Match("hi"))
	t.Log(trie.Match("hello"))
	t.Log(trie.Match("hell"))
	t.Log(trie.Match("heaven"))
	trie.Remove("hello")
	t.Log(trie.Match("hello"))
	t.Log(trie.Match("hell"))
}
