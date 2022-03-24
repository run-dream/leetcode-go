[278. First Bad Version](https://leetcode.com/problems/first-bad-version/)

#### 思路
二分法,重点在于怎么解决边界情况。
我们要找到 isBadVersion(left) == false, isBadVersion(left + 1) == true 的值
也就是结束时,right=firstBadVersion,left+1=right。
根据二分法, 可以不断缩小 firstBadVersion 的范围
mid = left + (right - left) / 2
如果 isBadVersion(mid) == true 的话，说明 firstBadVersion 在 [left, mid] 中
否则 firstBadVersion 在 [mid+1, right] 中，但是 要求left 不满足，如果+1的话可能就满足了
考虑到我们最后区间长度至少为2,因此 区间改为[mid,right]