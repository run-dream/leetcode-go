[120. Triangle](https://leetcode.com/problems/triangle/)

## 思路
1. 自顶向下，逐步计算每一层的值。可以修改原数组，也可以用额外空间存储。
2. 自底向上， dp 方程为
   `dp[i] = min(dp[i], dp[i+1]) + dp[i][j] for i in xrange(len-2,-1,-1) j in xrange(len(triangle[i])`