[315. Count of Smaller Numbers After Self](https://leetcode.com/problems/count-of-smaller-numbers-after-self/)

### 解题思路
这道题目实际上是在求数组的有序度
1. 暴力求解 O(n^2)
直接遍历数组，去找每个数字后出现的次数

2. 归并排序 O(n logn)
关键在于
1. 合并时如何对count进行更新
![merge](https://img-blog.csdnimg.cn/20181221171333773.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3l5c2F2ZQ==,size_16,color_FFFFFF,t_70)
> 看上面的图，本题的主要思路：
> 假设两个有序（升序）数组进行归并时候，怎么计算count[i]呢（count表示后面有多少个数比nums[i]小）
>（1）首先nums[ i_0]=-7,nums[ j_0]=-2,所以-7后面没有比它小的数,count[i_0]不变，i到达i_1的位置, j位置不变
>（2）此时nums[i_1]=1,nums[j_0]=-2,所以count[i_1]++, j到达j_1位置（注明一点，开头已经说了有序，所以在一个子数组里面不可能有数比前面的数小）nums[j_1]=1
>此时相等怎么办呢，count不用加就好了
>……然后后面依次是这么计算count数组的

2. 如何绑定count和nums 
> 我们定义数组pos[i] = j，表示第排名i个数据的元素下标是j。i就是排序完之后的数组的顺序，pos[i]记录的是这个数在nums数组的下标，这样最后输出就能按照题目的格式要求了。

### 参考资料
[leetcode 315](https://blog.csdn.net/yysave/article/details/85165099)