[Majority Elements](https://leetcode.com/problems/majority-element/)

### 思路
题目要求O(n)的时间复杂度，O(1)的空间复杂度。
使用Boyer-Moore Majority Vote Algorithm 来解决这个问题。
使用 cnt 来统计一个元素出现的次数，当遍历到的元素和统计元素不想等时，令 cnt--。如果前面查找了 i 个元素，且 cnt == 0 ，说明前 i 个元素没有 majority，或者有 majority，但是出现的次数少于 i / 2 ，因为如果多于 i / 2 的话 cnt 就一定不会为 0 。此时剩下的 n - i 个元素中，majority 的数目依然多于 (n - i) / 2，因此继续查找就能找出 majority。