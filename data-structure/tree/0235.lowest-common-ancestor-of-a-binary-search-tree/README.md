[235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)

### 思路
p 和 q 都在二叉查找树中，分情况讨论,设 p < q
1. 如果 root.val >= p and root.val <= q, 则最近公共祖先即为root
2. 如果 root.val >= p and root.val >= q, 则递归其右子树
3. 如果 root.val <= p and root.val <= q, 则递归其左子树