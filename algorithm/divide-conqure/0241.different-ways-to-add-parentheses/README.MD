[https://leetcode.com/problems/different-ways-to-add-parentheses/](Different Ways to Add Parentheses)

### 基本思路
分治算法
往算式里加括号 实际上是在调整计算符号的计算顺序
问题可以拆分成 left operation right 的形式
只有当 left 和 right 都是确定的数字时才可以计算直接求解
可以利用换成进行加速