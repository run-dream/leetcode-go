[1492. The kth Factor of n](https://leetcode.com/problems/the-kth-factor-of-n/)


### 思路
1. 暴力法
   从 1 遍历 到 n , 如果可能整除 k --, k 为 0 时返回。 时间复杂度为 O(n)

2. 优化
   从 1 遍历到 sqrt(n), 如果 k 为0 ，返回
   否则 遍历 sqrt(n) -> 1, 如果 可以整除， k--。 时间复杂度为 O(log(n))
