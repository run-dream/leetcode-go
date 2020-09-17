package trie

type TrieNode struct {
	data      byte
	children  [26]*TrieNode
	isWordEnd bool
}

type Trie struct {
	root *TrieNode
}

func byteToInt(b byte) int {
	return int(b - byte('a'))
}

func (t *Trie) Insert(str string) *Trie {
	cursor := t.root
	for i := 0; i < len(str); i++ {
		cursor = cursor.children[byteToInt(str[i])]
		if cursor == nil {
			cursor = &TrieNode{str[i], [26]*TrieNode{}, i == len(str)-1}
		}
	}
	return t
}

func (t *Trie) Remove(str string) *Trie {
	cursor := t.root
	var prev *TrieNode
	var char byte
	for i := 0; i < len(str); i++ {
		current := cursor
		cursor = cursor.children[byteToInt(str[i])]
		if cursor.isWordEnd && i != len(str)-1 {
			prev = current
			char = str[i]
		}
	}
	hasNext := false
	for i := 0; i < len(cursor.children); i++ {
		if cursor.children[i] != nil {
			hasNext = true
			break
		}
	}
	if !hasNext {
		prev.children[byteToInt(char)] = nil
	}
	return t
}

func (t *Trie) Match(str string) bool {
	cursor := t.root
	for i := 0; i < len(str); i++ {
		cursor = cursor.children[byteToInt(str[i])]
		if cursor == nil {
			return false
		}
	}
	return true
}
