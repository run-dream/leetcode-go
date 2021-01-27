[Word Break](https://leetcode.com/problems/word-break/)

### 思路
问题等价于对于s[0...n-1], 是否可以有 备选项[]string wordDict 构成  
可以分解成子问题 s[0...i-1]  0 <= i <= n  是否可以由备选项构成
动态规划方程  
dp[i] = max(dp[i] || max(dp[i-len(wordDict[j])] if subseq == word))   
