[Valid Palindrome II](https://leetcode.com/problems/valid-palindrome-ii/)

两个指针 left = 0 right = len(s) -1 从前后同时开始遍历数组
若 s[left] == s[right] left++ right--
不等 则分两种情况处理
left++ 判断是否满足回文
right-- 判断是否满足回文