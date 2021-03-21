[371. Sum of Two Integers](https://leetcode.com/problems/sum-of-two-integers/description/)

### 思路
不使用加减号实现加法，也就是通过位运算来实现

a ^ b 表示没有考虑进位的情况下两数的和，(a & b) << 1 就是进位。