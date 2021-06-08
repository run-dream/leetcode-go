[Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/)

### 思路
中序遍历,先访问左节点，再访问中间节点，最后访问右节点
利用栈的性质进行非递归实现
```go
func inorder(root *TreeNode){
    if root == nil{
        return
    }
	stack := []*TreeNode{}
	cur := root
    for len(stack) > 0 || cur != nil {
        // 遍历左子树
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
        fmt.Println(cur.Val)
        
        cur = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        cur = cur.Right
    }
}
```