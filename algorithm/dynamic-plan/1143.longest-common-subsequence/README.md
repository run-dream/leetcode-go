[Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/)

### 思路
动态规划
对于text1[0...i-1]和text2[0...j-1]求最长公共子串，可以得知
如果已知子问题 text1[0...k-1] 0 <= k < i,  text2[0...l-1] 0 <= l < j 的最长公共子串长度 m
可以推出 text1[0...k],text2[0...l]的最长公共子串长度
n = m + 1 if text1[k] == text[l]
否则 n 为 子问题 text1[0...k-1],text2[0...l] 和  text[0...k],text2[0...l-1]的解的最大值

动态转移方程
dp[i][j] = if text1[i] == text2[j] dp[i-1][j-1] + 1
else max(dp[i-1][j],dp[i][j-1])

### 优化
可以用两个数组替换掉前面的二维数组