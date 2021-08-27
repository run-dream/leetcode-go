[ Coin Change](https://leetcode.com/problems/coin-change/)

### 思路
首先通过暴力的方式来找到思路，题目其实可以理解成对硬币的层次遍历。
发现不满足贪心的全局最优。我们注意到 其实 最终状态的值，只可能由从 coins 选取一个硬币得到。
我们用 dp[i] 表示 凑齐 i 的量需要的银币数目。
也就是 dp[i] = min(dp[i-coins[j]] + 1)
用dp[i][j] 表示用完第i的硬币后到达j的数量的最小值，有动态规划方程
dp[i][j] = min(dp[i-1][j-n*coins[i]]+n, dp[i][j])
考虑到 前i个决策的结果和仅由前 i-1次决策有关系，因此可以将 dp 定义为一维数组 
dp[i] = min(dp[i], dp[j-n*coins[i]] + n |  0 <= n <= n / coins)

这是一个完全背包问题，完全背包问题和 0-1 背包问题在实现上唯一的不同是，第二层循环是从 0 开始的，而不是从尾部开始。