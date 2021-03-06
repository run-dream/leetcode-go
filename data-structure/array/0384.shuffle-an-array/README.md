[384. Shuffle an Array](https://leetcode.com/problems/shuffle-an-array/)

### 思路
重点在于如何生成一个随机数组

1. 每次从现存的数组中按照下标随机取一个元素，如果该元素已经取过了，则用取模的运算拿下一个未取到的元素
   WRONG ANSWER 
2. 遍历数组 每次将其和某个随机元素交互
Fisher-Yates shuffle算法 
（1）从数组中随机选取一个数p。
（2）将p与数组中最后（也可以是最前）的元素交换。（如果随机选中的是最后的元素，则相当于没有发生交换）
（3）去掉最后的元素（这里并没有删除操作，而是缩小索引值范围），即选中的p，缩小选取的数组范围。
（4）重复步骤（1）~（3），直到数组的长度为1时结束。

证明
第k个元素在最后一个位置的概率为 1/size
第k个元素在倒数第二个位置的概率为 (size-1)/size * 1/(size-1) = 1/size
第k个元素在倒数第j个位置的概率为 (size-1)/size * (size-2)/(size-1) *... * 1/(size-k+1) = 1/size
