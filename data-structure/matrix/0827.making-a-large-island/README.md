[827. Making A Large Island](https://leetcode.com/problems/making-a-large-island/)

### 思路
最暴力的方式是对于每个为0的位置进行尝试，然后计算出的面积，但是这样重复计算的东西太多了。
我们可以先将矩阵染色 并计数 然后再对于每个为0的位置进行尝试