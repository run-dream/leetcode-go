[String Without AAA or BBB](https://leetcode.com/problems/string-without-aaa-or-bbb/)

### 思路
1. n * a + m * b 其中最多出现两个连续的a或者b
2. 保证题目有解
3. 题目等价于找到出现次数比较多的那个字符 假设位a 剩余出现次数较少的字符标记位假设为b
   构造 aab * i+ab * j  + a * k (k 属于[0,1,2]）= n  * a + m * b 
   也就是 
   2*i + j + k = n， i + j  = m 
   解出来 i = n - m - k , j = 2m -n + k 现在只需要确定k的值 即可求解 
   若 k = 0 ， i = n -m ,j = 2m -n  看是否大于0  若大于0 则有可能
   若 k = 1，i = n - m - 1,   j = 2m - n + 1  看是否大于0  若大于0 则有可能
   若 k = 2， i = n - m -2, j = 2m - n + 2 看是否大于0  若大于0 则有可能
