[1680. Concatenation of Consecutive Binary Numbers](https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/)


### 思路
 N % m = sum(Ni % M) % M 其中 Ni 之和为 N

 n = 1, 1 => 1
 n = 2, 1 10 => 6
 n = 3, 11 110 => 30

 n(k) = n(k-1) << len(k) + k

考虑溢出的问题 最大的长度不超过64位