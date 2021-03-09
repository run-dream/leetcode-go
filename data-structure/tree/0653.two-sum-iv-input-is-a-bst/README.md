[653. Two Sum IV - Input is a BST](https://leetcode.com/problems/two-sum-iv-input-is-a-bst/)

### BST
二叉查找树
二叉查找树要求，在树中的任意一个节点，其左子树中的每个节点的值，都要小于这个节点的值，而右子树节点的值都大于这个节点的值。

### 思路
1. 遍历树，对于每个节点node，去找是否存在 k - node.val 时间复杂度 O(n*log(n))
2. 使用中序遍历得到有序数组之后，再利用双指针对数组进行查找。