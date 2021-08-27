[Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)

### 思路
仍然首先去想解空间。可以理解为二叉树的层次遍历。
最后的要求的那个节点要么走一步，要么走两步得到。
动态规划
方程:
dp[n] = dp[n-2] + dp[n-1]

实际上是一个 fibonacci 数列。