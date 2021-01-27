[ Coin Change](https://leetcode.com/problems/coin-change/)

### 思路
用dp[i][j] 表示用完第i的硬币后到达j的数量的最小值，有动态规划方程
dp[i][j] = min(dp[i-1][j-n*coins[i]]+n, dp[i][j])
考虑到 前i个决策的结果和仅由前 i-1次决策有关系，因此可以将 dp 定义为一维数组 
dp[i] = min(dp[i], dp[j-n*coins[i]] + n |  0 <= n <= n / coins)

这是一个完全背包问题，完全背包问题和 0-1 背包问题在实现上唯一的不同是，第二层循环是从 0 开始的，而不是从尾部开始。