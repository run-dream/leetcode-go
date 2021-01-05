[https://leetcode.com/problems/perfect-squares/](Perfect Squares)

### 思路
1. 回溯法
首先找出n以内的完全平方数
``` go
func backtrack(nums int, target int,i int, count int, result int){
    if target == 0 {
        if result == -1 {
            return count
        }
        if result < min{
            return result
        }
        return min
     }
     if target < 0 {
        return result
     }
     for j := i ; j < len(nums); j ++{
         tmp = backtrack(nums, target - nums[j], j, count + 1, result)
         if tmp < result{
             result = tmp
         }
     }
     return result
}
```
2. 动态规划
动态转移方程
f(n) = min(f(n-1), f(n-4), f(n-9)) + 1