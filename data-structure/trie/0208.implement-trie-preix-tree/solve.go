package problem0208

type TrieNode struct {
	children  [26]*TrieNode
	isWordEnd bool
}

type Trie struct {
	root *TrieNode
}

func initTrieNode(isWordEnd bool) *TrieNode {
	return &TrieNode{[26]*TrieNode{}, isWordEnd}
}

func byteToInt(b byte) int {
	return int(b - byte('a'))
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: initTrieNode(false),
	}
}

func (t *Trie) Insert(str string) {
	cursor := t.root
	for i := 0; i < len(str); i++ {
		next := cursor.children[byteToInt(str[i])]
		if next == nil {
			cursor.children[byteToInt(str[i])] = initTrieNode(false)
			next = cursor.children[byteToInt(str[i])]
		}
		cursor = next
	}
	cursor.isWordEnd = true
}

func (t *Trie) Search(str string) bool {
	cursor := t.root
	for i := 0; i < len(str); i++ {
		cursor = cursor.children[byteToInt(str[i])]
		if cursor == nil {
			return false
		}
	}
	return cursor.isWordEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(str string) bool {
	cursor := t.root
	for i := 0; i < len(str); i++ {
		cursor = cursor.children[byteToInt(str[i])]
		if cursor == nil {
			return false
		}
	}
	return cursor != nil
}
