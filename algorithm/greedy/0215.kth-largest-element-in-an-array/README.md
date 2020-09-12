[215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)

解题思路：
1. 全部排序 取倒数第k个元素
2. 选择排序 选出第k个
3. 快速排序 直到分区的值为k
选择数组区间 A[0…n-1]的最后一个元素 A[n-1]作为 pivot，对数组 A[0…n-1]原地分区，这样数组就分成了三部分，A[0…p-1]、A[p]、A[p+1…n-1]。
如果 p+1 = K，那 A[p]就是要求解的元素；
如果 K > p+1, 说明第 K 大元素出现在 A[p+1…n-1]区间，我们再按照上面的思路递归地在 A[p+1…n-1]这个区间内查找。
如果 K < p+1，那我们就在 A[0…p-1]区间查找。
4. 建立最大堆