[2 Keys Keyboard](https://leetcode.com/problems/2-keys-keyboard/)

### 思路
1. 动态规划
每一步骤有两个可选操作 复制 或者 粘贴
复制的结果是 粘贴板的长度变为 m 实际长度仍为 m
粘贴的结果是 粘贴板的长度不变 为p 实际长度为m+p
注意到获取粘贴板的次数获取也遵从同样的逻辑
所以
if i % j == 0 dp[i] = min(dp[j] + dp[i/j])  0 < j < i

2. 质因素分解