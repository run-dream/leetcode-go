[1034. Coloring A Border](https://leetcode.com/problems/coloring-a-border/)

 ### 思路 
bfs, 变色的条件是 grid[x][y] == oldColor && !neignbors.all(neignbor=> neignbor == oldColor)
需要注意的是如何标记已经访问的节点，如何区分边界节点。
这里可以用-1表示内部节点，-2表示边界节点。 