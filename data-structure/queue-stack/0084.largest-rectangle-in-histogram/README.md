[84. Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/)


## 思路
1. 暴力法，对于每个柱子，求以它为高度的最大长方形面积。遍历前后两个方向。时间复杂度 O(n*n)
2. 栈。优化找左右的过程。