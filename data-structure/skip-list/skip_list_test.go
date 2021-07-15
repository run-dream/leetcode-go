package skiplist

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSkipList(t *testing.T) {
	sl := NewSkipList()

	for i := 0; i < 10000; i++ {
		sl.Insert(strconv.Itoa(i), float64(i))
	}
	t.Log(sl.level)
	head, tail := sl.GetRange(1000, 1005)
	Traverse(head, tail)

	sl.Delete("1002", 1002)
	head, tail = sl.GetRange(1000, 1005)
	Traverse(head, tail)
}

func Traverse(head *Node, tail *Node) {
	fmt.Println("BEGIN")
	cur := head
	for cur != tail {
		fmt.Println("Traverse", cur.ele, cur.score)
		cur = cur.levels[0].forward
	}
	if cur != nil && cur == tail {
		fmt.Println("Traverse", cur.ele, cur.score)
	}
	fmt.Println("END")
}
