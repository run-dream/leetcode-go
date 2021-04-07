[Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/description/)


### 基本思路
1. 首先找出最优子结构
2. 写状态转移方程
3. 将状态转移方程写成代码
dp[n] 表示以nums[n]结尾的序列的最长递增子序列长度 
即 如果 已知 dp[1...n] 如何求解 dp[n+1]
如果 nums[j] < nums[n+1] |  1= < j <=n 则 dp[n+1] = max(dp[1...j]) + 1
但是如果找不到符合条件的j,则 dp[n+1] = 1

### 思路二
维护一个上升子序列 对于每个值找到其应该在位置
用二分法去处理