[House Robber III](https://leetcode.com/problems/house-robber-iii/)


### 思路
对于任意一个节点，其可以获取的最大值是
rob(root) = max(root.Val + rob(root.Left.Left) + rob(root.Left.Right) + rob(root.Right.Left)+ rob(root.Right.Right), rob(root.Left) + rob(root.Right))

### 优化
上面的解决方式中 rob 重复计算的部分比较多， 实际在我们在访问任何一个节点的时候就可以返回两种方式的值

### 参考
https://leetcode.com/problems/house-robber-iii/discuss/79330/Step-by-step-tackling-of-the-problem