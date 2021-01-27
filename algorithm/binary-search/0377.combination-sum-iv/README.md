[Combination Sum IV](https://leetcode.com/problems/combination-sum-iv/)

### 思路
完全背包问题
用dp[i] = v 表示和为 i 的组合有v种
动态方程为：
dp[i] = sum(dp[i-nums[j]*n])
