[Combination Sum II](https://leetcode.com/problems/combination-sum-ii/description/)
1. 基本思路
回溯法
这道题目类似于0-1背包问题
首先将candidates进行排序处理
问题等同于对于candidates中的每个元素进行一次是否选取的判断.
注意因为一个数字可以出现多次，所以需要去重