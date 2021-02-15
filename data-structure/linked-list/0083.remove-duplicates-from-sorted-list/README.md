[Remove Duplicates from Sorted List](https://leetcode.com/problems/remove-duplicates-from-sorted-list/)

### 思路
1. 遍历列表，记录前一个指针，next指向值不相同的节点， 注意末尾可能会出现重复的元素，所以加上特殊处理
2. 递归处理