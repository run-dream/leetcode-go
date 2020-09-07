# 定义
跳表(Skip List)是一种查询,插入，删除时间复杂度都为O(log(n))的数据结构，用于解决算法中的查找问题。
优势原理简单，容易实现，效率较高。
应用： redis的s有序集合 leveldb

# 算法思路
在连标上加多级索引

# 参考资料
https://juejin.im/post/6844903446475177998
https://cloud.tencent.com/developer/article/1353762
https://github.com/wangzheng0822/algo/blob/master/go/17_skiplist