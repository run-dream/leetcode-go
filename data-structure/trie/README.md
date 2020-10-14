### 定义
Trie 树，也叫“字典树”，“前缀树”。它是一种专门处理字符串匹配的数据结构，用来解决在一组字符串集合中快速查找某个字符串的问题。

### 优点
查询和插入，时间复杂度为O(k)，k为字符串长度

### 缺点
如果大量字符串没有共同前缀时很耗内存。

### 核心思想
它的核心思想就是通过最大限度地减少无谓的字符串比较，使得查询高效率，即「用空间换时间」，再利用共同前缀来提高查询效率。