[142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)


### 思路
1. 用 map 来存储已经访问过的节点
2. 快慢指针，先找出是否存在环，再将慢指针置为首指针，继续遍历，直到相遇。
   原因是： 相遇的条件是 快指针比慢指针多走了 `n * size + pos`步 