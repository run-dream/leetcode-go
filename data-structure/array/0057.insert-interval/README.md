[Insert Interval](https://leetcode.com/problems/insert-interval/)

### 思路
两种情况：
1. 没有可合并的间隔，找出插入的位置，也就是 最后一个 inteval[1] < newInterval[0]的位置
2. 有可合并的间隔，记录合并的起点和长度，更新newInteval