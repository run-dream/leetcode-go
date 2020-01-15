[sqrx](https://leetcode.com/problems/sqrtx/)
1. 基本思路
一个数 x 的开方 sqrt 一定在 0 ~ x 之间，并且满足 sqrt == x / sqrt 。可以利用二分查找在 0 ~ x 之间查找 sqrt。

2. 牛顿迭代法
从函数意义上理解：我们是要求函数f(x) = x²，使f(x) = num的近似解，即x² - num = 0的近似解。
