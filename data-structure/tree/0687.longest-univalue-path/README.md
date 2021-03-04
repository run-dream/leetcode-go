[687. Longest Univalue Path](https://leetcode.com/problems/longest-univalue-path/)

### 思路
分情况讨论,以root为拐点的path的最大值，分别计算出每个中间节点的最大值，然后求树的最大值,注意对于上级节点来说只能选择一条最大的分支