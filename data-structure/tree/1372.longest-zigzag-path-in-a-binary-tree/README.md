[1372. Longest ZigZag Path in a Binary Tree](https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/)

### 思路
对于任一节点 
zigzag(node) = 0 if node.left == nil and node.right == nil
zipzag(node) = zipzag(node.right, left) if node.left == nil
zigzag(node) = max(zigzag(node.left, right), zigzag(node.right, left)) + 1
 
思路 dfs