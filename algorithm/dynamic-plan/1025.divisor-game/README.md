[DivisorGame](https://leetcode.com/problems/divisor-game/)

### 思路
- 动态规划 
定义dp[i] 为当N为i，且轮到Alice选择时是否会胜利，即Bob走到i-x时，Alice就会赢
dp[i] = true    ## 存在x，满足(dp[i - x] == false) && (N % x == 0) && (1 < x < N)
        false   ## 其他

- 博弈论
奇数必输，偶数必赢