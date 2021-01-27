[Best Time to Buy and Sell Stock III](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/)


### 思路
- 暴力
问题可以分解为将nums 分解为 nums[0...j], nums[j+1...len(nums)-1]  
求 prices(0,j) + prices(j+1,len) 的最大值  
dp[i][j] = v 来表示 第 i 天买入 第 j 天卖出的最大收益 有  
dp[i][j] = max(dp[i][j-1],dp[i][j])  

- 分析
将买入和卖出都视为一次决策,买入的收益是 - prices[i] 卖出的收益 是 + prices[i]
用dp[i][j] 来表示 在k号位置第i次决策的实时收益
dp[i][j] = max(dp[i-1][k] - prices[k]) ( i % 2 == 0, 0 <= k < j)
      || max(dp[i-1][k] + prices[k] (i %2 == 1, 0 <= k < j)

