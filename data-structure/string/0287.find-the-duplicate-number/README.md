[Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/)

###  要求
空间 O(1)
时间 < O(n*n)
不修改原数组内容

### 思路
二分法确定所在区间
如果k重复出现了，那么对于两个区间 [low, mid],[mid, high] 的元素个数较多的哪个就是k所在的区间