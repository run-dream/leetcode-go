[671. Second Minimum Node In a Binary Tree](https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/)

### 思路
对于每个节点 有 root.val = min(root.left.val, root.right.val) 
因此分情况讨论 
1. root == nil return -1
2. root.left == nil && root.right == nil return -1
3. return min(leftVal, rightVal)