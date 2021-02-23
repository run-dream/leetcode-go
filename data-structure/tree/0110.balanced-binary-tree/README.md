[Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/)

### 思路
dfs 记录每个节点的深度，如果存在depth(left) - depth(right) > 1 则返回false