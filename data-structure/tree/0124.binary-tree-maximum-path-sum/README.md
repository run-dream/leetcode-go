[124. Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/)

## 思路
1. 首先 确定经过一个节点的最大路径的值 = `node.val + gain(left) + gain(right)`
2. 现在问题变成了求解 `gain(node)`, 容易得到 `gain(node) = max(0, gain(node.left), gain(node.right)) + root.val`
3. 在计算 `gain(node)` 的过程中, 可以通过递归实现，同时注意到遍历的时候其实就可以更新最大值。
