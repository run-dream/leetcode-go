### 描述
[最长公共前缀](https://leetcode-cn.com/explore/featured/card/bytedance/242/string/1014/)
编写一个函数来查找字符串数组中的最长公共前缀

### 思路
1. 取 strs[0] 为 sample, 从0到len(sample)遍历 sample和其他字符串,出现不匹配或者超出长度就可以停下来了
