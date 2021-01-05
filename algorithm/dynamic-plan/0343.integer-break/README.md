[Integer Break](https://leetcode.com/problems/integer-break/)

### 基本思路
1. 递归
自顶向下 
f(n) = max(1*f(n-1),2*f(n-2),...,(n-1)*f(1), i * (n-i))

2. 动态规划
dp[i] = max(dp[i], max(j*dp[i-1], j * (i-j)))

3. 数学分析



### 参考
[Leetcode0343](https://www.jianshu.com/p/a9760fa3c686)