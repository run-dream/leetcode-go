[ Find All Duplicates in an Array](https://leetcode.com/problems/find-all-duplicates-in-an-array/)

### 思路
要求不使用额外的空间
1. 考虑交换 nums[i] 和 nums[nums[i]-1] 的位置，
直到 nums[i] == i + 1 || nums[i] == nums[nums[i]-1]
最后找出所有不匹配的即可
2. 直接用负数来标记以出现的值