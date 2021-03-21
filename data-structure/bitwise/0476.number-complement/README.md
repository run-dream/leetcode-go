[476. Number Complement](https://leetcode.com/problems/number-complement/)

### 思路
不考虑二进制表示中的首 0 部分
对于 00000101，要求补码可以将它与 00000111 进行异或操作。那么问题就转换为求掩码 00000111。