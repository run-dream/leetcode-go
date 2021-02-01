[Longest Harmonious Subsequence](https://leetcode.com/problems/longest-harmonious-subsequence/)

### 思路
1. 用hash表存储每个数字出现的次数
2. 遍历hash表的keys,求max(hash[key]+hash[key-1], hash[key]+hash[key+1])