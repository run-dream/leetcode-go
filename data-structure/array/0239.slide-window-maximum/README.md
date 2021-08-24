
[239. Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/)

## 思路
1. 模拟这个过程, 暴力
   ``` go
   for i := 0; i < len(nums-k); i ++ {
       ans[i] = max(nums[i:i+k])
   }
   ```
2. 优化，注意到当从 i = j 到 i = j + 1 期间 `max(nums[j:j+k])` 和 `max(nums[j+1:j+1+k]` 中有很大一部分是重复的。
可以想办法利用前面部分重复的信息。由此想到可以使用双端队列。
维护一个双端队列，遍历每个元素 x ，检查队列的头元素，如果已经不在窗口内，则可以去除掉这个元素。
同样，如果队列里的某个元素比 x 小，则也可以去除掉。这样最大的元素一定在队首。
**注意，deque里维护的是 下标**
3. dp
找出其每个元素 left 和 right 边界的最大值。