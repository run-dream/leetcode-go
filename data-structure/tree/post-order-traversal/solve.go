package post_order_traversal

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

type Node struct {
	left  *Node
	right *Node
	val   int
}

func postOrder(root *Node) {
	if root == nil {
		return
	}
	stack1 := stack.New()
	cur := root
	// 需要判断上次访问的节点是位于左子树，还是右子树。若是位于左子树，则需跳过根节点，先进入右子树，再回头访问根节点；若是位于右子树，则直接访问根节点
	var prev *Node
	for cur != nil {
		stack1.Push(cur)
		cur = cur.left
	}
	for stack1.Len() > 0 {
		cur = stack1.Pop().(*Node)
		//一个根节点被访问的前提是：无右子树或右子树已被访问过
		if cur.right == nil || cur.right == prev {
			fmt.Println(cur.val)
			// 修改最近被访问的节点
			prev = cur
		} else {
			stack1.Push(cur)
			cur = cur.right
			for cur != nil {
				stack1.Push(cur)
				cur = cur.right
			}
		}
	}
	fmt.Println()
}
