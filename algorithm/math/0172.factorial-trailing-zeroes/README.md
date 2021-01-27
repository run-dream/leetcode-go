[Factorial Trailing Zeroes](https://leetcode.com/problems/factorial-trailing-zeroes/)

### 思路
末尾的0 由因式分解后 2 和 5 的个数决定
而从1到n的2的个数 等于 1 到 n-n%2 的2的个数 
这个数量远大于 1到n-n%5中5的个数
也就是0的个数取决于5的个数