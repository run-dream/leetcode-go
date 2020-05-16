[Non-decreasing Array](https://leetcode.com/problems/non-decreasing-array/)

### 基本思路
关键点是当出现 array[i-1] > array[i] 时修改哪个元素
考虑 array[i-1] 和 array[i+1] 的关系，
1. 优先考虑令 nums[i - 1] = nums[i]，因为如果修改 nums[i] = nums[i - 1] 的话，那么 nums[i] 这个数会变大，那么就有可能比 nums[i + 1] 大，从而影响了后续判断。
2. 还有一个比较特别的情况就是 nums[i] < nums[i - 2]，只修改 nums[i - 1] = nums[i] 不能令数组成为非递减，只能通过修改 nums[i] = nums[i - 1] 才行。