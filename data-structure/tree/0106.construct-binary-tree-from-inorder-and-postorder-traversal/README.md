[106. Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)

### 思路
首先用一颗树来做demo，进行先序和后序遍历
      1
    2   3
  4  5  6  7
先序结果为 4 2 5 1 6 3 7
后序结果为 4 5 2 6 7 3 1
观察到后序的最后一个元素为1树的根节点
然后可以根据先序的结果来拆分左右子树的长度
在分别对左右子树进行相同的操作即可求解