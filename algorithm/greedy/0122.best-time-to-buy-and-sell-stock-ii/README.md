[Best Time to Buy and Sell Stock II](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description/)

### 基本思路
次交易包含买入和卖出，多个交易之间不能交叉进行。
对于 [a, b, c, d]，如果有 a <= b <= c <= d ，那么最大收益为 d - a。而 d - a = (d - c) + (c - b) + (b - a) ，因此当访问到一个 prices[i] 且 prices[i] - prices[i-1] > 0，那么就把 prices[i] - prices[i-1] 添加加到收益中，从而在局部最优的情况下也保证全局最优。
