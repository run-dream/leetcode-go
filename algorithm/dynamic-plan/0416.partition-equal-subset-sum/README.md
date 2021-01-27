[Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/)

### 思路
问题等价于从数组 nums[0...n-1]中选取出k个元素，使它们的和恰好为 sum(nums[0...n]) / 2  
问题转化成一个特殊的0-1背包问题  
可以用回溯法来做  
``` go
func backtrack(nums []int, start, value int) bool{
    if value == 0 {
        return true
    }
    if start >= len(nums){
        return false
    }
    return backtrack(nums, start + 1, value - nums[start])
    || backtrack(nums, start + 1, value)
}

```

时间复杂度为 O(pow(2,n))  
这个时间复杂度是无法接受的,考虑用动态规划的方式来用空间换取时间  
用 i 来表示第i次选择, j 来表示剩余空间 ,dp[i][j] 来表示是否可以到达在第i次选择是否可以到达剩余j的状态(bool类型)  
可以得出动态转移方程  
dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]  
将其转化为代码  
这里空间复杂度为 m * n  
考虑到 前i个数字的和仅由前 i-1 件的数字有关系，因此可以将 dp 定义为一维数组  
dp[j] 既可以表示 dp[i-1][j] 也可以表示 dp[i][j]  
动态规划方程为  
dp[j] = dp[j]|| dp[j-nums[i]]  

