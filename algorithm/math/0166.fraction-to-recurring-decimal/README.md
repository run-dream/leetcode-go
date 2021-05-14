[166. Fraction to Recurring Decimal](https://leetcode.com/problems/fraction-to-recurring-decimal/)


### 思路
1. 处理符号位 
2. 处理整数部分 
3. *** 处理分数部分 ***  
   - 如何判断无限循环 
      以 2 / 3 举例 
      2 *10 = 20 20 / 3 = 6 20 % 3 = 2 
      注意到 2 是之前出现过的除数 之后的结果应该是和之前的一致
      我们用一个map 来 记住 2 出现的结果的开始位置 用 (6) 来代替