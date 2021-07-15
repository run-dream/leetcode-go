[638. Shopping Offers](https://leetcode.com/problems/shopping-offers/)

### 思路
问题有点像是背包问题，只是容量变成了数组 
我们将special 进行扩展,将price也作为special的一部分进行统一处理 
以 Input: price = [2,5], special = [[3,0,5],[1,2,10]], needs = [3,2] 为例 
offers = [[1, 0, 2], [0, 1, 5], [3, 0, 5], [1, 2, 10]] 
问题变成了在上者些组合中去选择
状态有 在第i个offer进行选择 剩余的需求
定义 dp `dp(i, needs)` 表示还剩余 needs 未满足时的成本
动态转移方程为
```
dp(i, needs)= min(dp[i-1][needs - offer[i]*n])
```
最后返回 `dp(i, needs)` 的值