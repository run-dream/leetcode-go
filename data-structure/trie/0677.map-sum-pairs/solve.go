package problem0677

type MapSum struct {
	children [26]*MapSum
	val      int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{[26]*MapSum{}, 0}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this
	for i := 0; i < len(key); i++ {
		index := byteToInt(key[i])
		node := cur.children[index]
		if node == nil {
			newNode := Constructor()
			node = &newNode
			cur.children[index] = node
		}
		if i == len(key)-1 {
			node.val = val
		}
		cur = node
	}
}

func (this *MapSum) Sum(prefix string) int {
	cur := this
	for i := 0; i < len(prefix); i++ {
		index := byteToInt(prefix[i])
		cur = cur.children[index]
		if cur == nil {
			return 0
		}
	}
	queue := []*MapSum{cur}
	sum := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sum += node.val
		for i := 0; i < 26; i++ {
			if node.children[i] != nil {
				queue = append(queue, node.children[i])
			}
		}
	}
	return sum
}

func byteToInt(b byte) int {
	return int(b - byte('a'))
}
