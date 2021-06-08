[144. Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/)

### 思路
前序遍历，先访问根结点，再访问左节点，最后访问右节点

### 实现
1. 递归
```go
func preorder(root *TreeNode){
    if root == nil{
        return 
    }
    fmt.Println(root.Val);
    preorder(root.Left)
    preorder(root.Right)
}
```
2. 循环
利用栈的性质
```go
func preorder(root *TreeNode){
    if root == nil{
        return
    }
    stack := []int{root}
    for len(stack) > 0 {
        cur := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        fmt.Prinln(cur.Val)
        if cur.Right != nil{
            stack = append(stack, cur.Right)
        }
        if cur.Left != nil{
            stack = append(stack, cur.Left)
        }
    }
}
```