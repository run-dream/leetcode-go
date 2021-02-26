[669. Trim a Binary Search Tree](https://leetcode.com/problems/trim-a-binary-search-tree/)


### 思路
二叉查找树满足
root.Left.Val < root.Val < root.Right.Val
要筛选出[low, high]之间节点，dfs查找树
对于每个节点，如果node.val ~ [low, high],则保留其左右节点，继续遍历
如果 node.val > high 放弃其右节点，遍历左子树
如果 node.val < low 放弃其左节点，遍历右子树