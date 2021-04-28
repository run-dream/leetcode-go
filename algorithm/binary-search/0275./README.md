[275. H-Index II](https://leetcode.com/problems/h-index-ii/)

### 思路
二分查找
查出第一个 i 使得 citiations[i] <= n - i
则 n - i 则为所求h-index