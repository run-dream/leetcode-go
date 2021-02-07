### 定义
散列表用的就是数组支持按照下标随机访问的时候，时间复杂度是 O(1) 的特性。我们通过散列函数把元素的键值映射为下标，然后将数据存储在数组中对应下标的位置。当我们按照键值查询元素时，我们用同样的散列函数，将键值转化数组下标，从对应的数组下标的位置取数据。  

### 散列函数
1. 散列函数计算得到的散列值是一个非负整数； 
2. 如果 key1 = key2，那 hash(key1) == hash(key2)； 
3. 如果 key1 ≠ key2，那 hash(key1) ≠ hash(key2)。 

### 散列冲突
1. 开放寻址法  
如果出现了散列冲突，我们就重新探测一个空闲位置，将其插入。 
- 线性探测
- 二次探测
- 双重散列

装载因子:  
散列表的装载因子=填入表中的元素个数/散列表的长度  

2. 连表法  
在散列表中，每个“桶（bucket）”或者“槽（slot）”会对应一条链表，所有散列值相同的元素我们都放到相同槽位对应的链表中。  