[House Robber](https://leetcode.com/problems/house-robber/)

### 思路
1. 回溯法
```go
func backtrack(nums []int,visited []bool, start int, val int, result int) int{
    if i == len(nums) {
        if val > result{
            result = val
        }
        return result
    }
    for j := start; i < len(nums); j ++ {
        if j > 2 && visited[j-2]{
            continue
        }
        visited[j] = true;
        result = backtract(nums, visited, j, val + nums[j], result)
        visited[j] = false
    }
    return result
}
```

2. 动态规划
动态规划方程
由于不能抢劫邻近住户，因此如果抢劫了第 i 个住户那么只能抢劫 i - 2 和 i - 3 的住户，所以
dp[n] = max(dp[n-2], dp[n-3]) + nums[n]