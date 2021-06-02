package post_order_traversal

import (
	"fmt"
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
	stack := []*Node{}
	cur := root
	// 需要判断上次访问的节点是位于左子树，还是右子树。若是位于左子树，则需跳过根节点，先进入右子树，再回头访问根节点；若是位于右子树，则直接访问根节点
	var prev *Node
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.left
	}
	for len(stack) > 0 {
		cur =  stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//一个根节点被访问的前提是：无右子树或右子树已被访问过
		if cur.right == nil || cur.right == prev {
			fmt.Println(cur.val)
			// 修改最近被访问的节点
			prev = cur
		} else {
			stack = append(stack, cur)
			cur = cur.right
			for cur != nil {
				stack = append(stack, cur)
				cur = cur.right
			}
		}
	}
	fmt.Println()
}
