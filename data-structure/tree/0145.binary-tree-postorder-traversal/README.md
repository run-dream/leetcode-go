[Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/)

### 思路
后序遍历
非递归实现
```go
func postorder(root *TreeNode){
    if root == nil{
        return
    } 
    stack := []*TreeNode{}
    cur := root
    prev := nil
    for len(stack) > 0 || cur != nil {
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        }
        if len(stack) > 0 {
            cur = stack[:len(stack)-1]
            stack = stack[:len(stack)-1]
            if cur.Right == nil || cur.Right == prev{
                fmt.Println(cur)
                prev = cur
                cur = nil
            }else{
                stack = append(stack, cur)
                cur = cur.Right
            }
        }
    }
}
```