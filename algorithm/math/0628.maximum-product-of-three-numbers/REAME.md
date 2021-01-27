[Maximum Product of Three Numbers](https://leetcode.com/problems/maximum-product-of-three-numbers/)

### 思路
1. 将数字排序
考虑如果都是正数的话 取最后三个相乘
如果都是负数的话 取最前三个相乘
如果既有正数又有负数，取 max(nums[0]*nums[1]*nums[-1], nums[-3],nums[-2],nums[-1])

2. 优化其实不需要全部排序 只要获取最小的三个和最大的三个数字就可以了