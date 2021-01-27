[Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/)

### 思路
注意要求  
1. 不使用除法
2. O(1)的时间复杂度

将数组分为left和right两部分，结果result[i]  = left * right