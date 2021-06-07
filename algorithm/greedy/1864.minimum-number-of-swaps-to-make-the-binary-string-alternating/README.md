[1864. Minimum Number of Swaps to Make the Binary String Alternating](https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-binary-string-alternating/)

### 思路
贪心
解决这个问题的关键是计算出错的位置的数量。 不需要考虑从哪个位置移动哪个位置，因为都是等价的
注意：
- 如果abs(count(1) - count(0)) > 1 问题无解
- 如果 count(1)  > count(0) 最终的字符串 将以1开头，反之则以0开头
- 如果 count(1) == count(0) 取两者的最小值