[538. Convert BST to Greater Tree](https://leetcode.com/problems/convert-bst-to-greater-tree/description/)

### 思路
遍历一遍 计算出总的值 currentSum
再中序遍历，node.Val =  surrentSum - prevNode.val
 
