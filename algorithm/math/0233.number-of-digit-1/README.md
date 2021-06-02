[233. Number of Digit One](https://leetcode.com/problems/number-of-digit-one/)


### 思路
计算每一位可能出现1的次数，如
if n = xyzdabc
then
(1) xyz * 1000                     if d == 0
(2) xyz * 1000 + abc + 1           if d == 1
(3) xyz * 1000 + 1000              if d > 1
因此可以从高位到低位一步步的计算1的次数