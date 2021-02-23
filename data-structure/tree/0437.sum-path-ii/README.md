[Path Sum III](https://leetcode.com/problems/path-sum-iii/)


### 思路
dfs 记录每个节点可达的值和方式数目

### 优化
不必记录每个节点可达的方式数目，只需要记录达到sum-node.Val的值