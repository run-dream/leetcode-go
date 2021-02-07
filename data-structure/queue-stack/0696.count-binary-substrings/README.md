 [Count Binary Substrings](https://leetcode.com/problems/count-binary-substrings/)

### 思路
遍历数组，记录下当前字母和前面数字相同的数量
prev >= cur 则count += 1
如果不同
则 count += 1 prev = cur, cur = 1