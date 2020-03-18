[arrange-coins](https://leetcode.com/problems/arranging-coins/)
解题思路
1. k行完整的硬币的总数为：
```
f(k) = 1 + 2 + ... + k = (1+k)*k/2 
```
题目等价于求不等式 f(k) <= n 的最大整数解
导数
```
f‘（k） = 1/2 + k
```
可知在k > -1/2 时 函数递增 可以按照二分法求解