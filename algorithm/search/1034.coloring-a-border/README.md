[1034. Coloring A Border](https://leetcode.com/problems/coloring-a-border/)

 ### 思路 
bfs, 变色的条件是 grid[x][y] == oldColor && !neignbors.all(neignbor=> neignbor == oldColor || neignbor == newColor) 