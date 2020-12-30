[Sudoku Solver](https://leetcode.com/problems/sudoku-solver/)

1. 思路
回溯法
相当于对每个空的格子尝试填充1-9,不符合数独的要求则回溯
结束条件是找不到空的格子

2. 优化
对每个格子是否可用的情况可以用数组进行优化 不必每次判断是否占用