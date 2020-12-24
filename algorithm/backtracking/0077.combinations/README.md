[Combinations](https://leetcode.com/problems/combinations/)
1. 基本思路
注意组合的定义是从n个数字中选择k个数，并不关注选取的顺序，即[1,2]和[2,1]是相同的.
问题分解为 遍历1...n,每次可做出选取或者不选,直到选够k个数字。
迭代函数为
``` golang
func backtract(start int, n int, k int, combination []int, result *[][]int){
    // start 记录已经遍历到哪个数字了
    // n 总数
    // k 剩余可选数量
}
```
