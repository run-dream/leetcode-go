[Longest Palindrome](https://leetcode.com/problems/longest-palindrome/)

### 思路
记录每个字母出现次数，
ans = sum(count[i] | count[i]%2==0) + sum(count[1]/2 | count[i] %2 ==1 && count[i] > 1) + 1 | some(count[i] == 1)