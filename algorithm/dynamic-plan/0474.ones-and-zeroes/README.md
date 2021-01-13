
[Ones and Zeroes](https://leetcode.com/problems/ones-and-zeroes/)

### 思路
用dp[i][m][n] 来表示 第i次决策时 有 m 个 0,n个1时最大的子集数量
dp[i][m][n] = max(dp[i-1][m-nums0[i]][n-nums1[i]]+1, dp[i-1][m][n])

考虑到 前i个决策的记过的和仅由前 i-1次决策有关系，因此可以将 dp 定义为二维数组  
dp[m][n] = max(dp[m-nums0[i]][n-nums1[i]]+1, dp[m][n])

注意计算时要从后往前算，以免导致覆盖计算