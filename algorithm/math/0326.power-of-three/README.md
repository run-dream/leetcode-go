[ Power of Three](https://leetcode.com/problems/power-of-three/)

### 思路
要求不使用递归或者循环
对于 pow(3^n), 有且只有 pow(3^i), 0 <= i <=n 被其整除
问题等价于 寻找 log(3,(pow(2,31)-1)) = 1162261467