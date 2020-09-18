package trie

type TrieNode struct {
	data      byte
	children  [26]*TrieNode
	isWordEnd bool
}

type Trie struct {
	root *TrieNode
}

// NewTrieNode 新增节点
func NewTrieNode(b byte, isWordEnd bool) *TrieNode {
	return &TrieNode{b, [26]*TrieNode{}, isWordEnd}
}

func byteToInt(b byte) int {
	return int(b - byte('a'))
}

func (t *Trie) Insert(str string) *Trie {
	cursor := t.root
	for i := 0; i < len(str); i++ {
		next := cursor.children[byteToInt(str[i])]
		if next == nil {
			cursor.children[byteToInt(str[i])] = NewTrieNode(str[i], false)
			next = cursor.children[byteToInt(str[i])]
		}
		cursor = next
	}
	cursor.isWordEnd = true
	return t
}

func (t *Trie) Remove(str string) *Trie {
	cursor := t.root
	for i := 0; i < len(str); i++ {
		cursor = cursor.children[byteToInt(str[i])]
	}
	cursor.isWordEnd = false
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
	return cursor.isWordEnd
}
