[Valid Perfect Square](https://leetcode.com/problems/valid-perfect-square/)

### 思路
1. 从 1...num 进行二分来求解
2. 平方序列：1,4,9,16,.. 间隔：3,5,7,...

间隔为等差数列，使用这个特性可以得到从 1 开始的平方序列。