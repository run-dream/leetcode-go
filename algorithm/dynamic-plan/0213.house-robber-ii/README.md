[House Robber II](https://leetcode.com/problems/house-robber-ii/)

1. 思路
动态规划
环状的街道导致问题变成抢了第0个就不能抢最后一个
问题变成求
max(rob(0, len(nums)-2), rob(1, len(nums)-1))
