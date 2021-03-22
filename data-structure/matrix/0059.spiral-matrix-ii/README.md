[Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii/)

### 思路
directions := [0,1],[1,0],[0,-1],[-1,0] 
每次按照一个方向走到底，找下一个方向继续走 直到总步数等于n*n
***注意回退***