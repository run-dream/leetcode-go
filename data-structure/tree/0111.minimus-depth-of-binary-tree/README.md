[Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/)

### 思路
minDepth(root) = min(minDepth(root.Right), minDepth(root.Left)) + 1
初态叶子节点
minDepth(root) = 1 if root.Left == nil && root.Right == nil