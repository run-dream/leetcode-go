[Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/)

### 思路
- 方案一
用dp[i][j]=true表示从i到j是回文串，则有  
dp[i-1][j+1] = dp[i][j] && s[i-1] == s[j+1]  
且有dp[i][i] = true, dp[i][i+1] = s[i] == s[i+1]  
从len i = 3 开始填表  
- 方案二
从字符串的某一位开始，尝试着去扩展子字符串。