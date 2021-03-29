[Minimum Number of Refueling Stops](https://leetcode.com/problems/minimum-number-of-refueling-stops/)

### 思路
1. 动态规划
背包问题的变形
用dp[i] 来表示 使用i个加油站最多能到达的位置，动态规划方程
dp[i] = max(dp[i], dp[i-1] + max(stations[j].fluel) if dp[j-1] > stations[i].pos and  0 < j <= i)
意义为
如果 经过 j-1个加油站可以越过 
2. 优先队列
有点像回溯的感觉
每次走到最远处 如果没有达到target 从 优先队列中取出经过的加油站中由最多的继续走  