[Wiggle Subsequence](https://leetcode.com/problems/wiggle-subsequence/)

### 思路
方法1:
问题即从数组nums中选取两个下标i,j 0 < i < len(nums) , i <= j < len(nums)
我们用dp[n]表示以第n的元素为结尾的子串的wiggle subsequence最大值 
如果第n+1个元素符合前面的趋势，则 dp[n+1] = dp[n] + 1
不符合的话 dp[n+1] = dp[n]
注意处理特殊情况
方法2:
使用两个状态 up 和 down。
