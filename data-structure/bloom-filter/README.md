# 设计思想

可以认为是 BitMap 的扩展，但 BitMap 当数据范围比较大时,会消耗大量内存，所以通过 k 个 hash 算法将输入映射为 k 位进行压缩。
Bloom Filter 是由一个长度为 m 的比特位数组（bit array）与 k 个哈希函数（hash function）组成的数据结构。位数组均初始化为 0，所有哈希函数都可以分别把输入数据尽量均匀地散列。
当要插入一个元素时，将其数据分别输入 k 个哈希函数，产生 k 个哈希值。以哈希值作为位数组中的下标，将所有 k 个对应的比特置为 1。
当要查询（即判断是否存在）一个元素时，同样将其数据输入哈希函数，然后检查对应的 k 个比特。如果有任意一个比特为 0，表明该元素一定不在集合中。如果所有比特均为 1，表明该元素有（较大的）可能性在集合中。为什么不是一定在集合中呢？因为一个比特被置为 1 有可能会受到其他元素的影响，这就是所谓“假阳性”（false positive）。相对地，“假阴性”（false negative）在 Bloom Filter 中是绝不会出现的。
![avatar](https://upload-images.jianshu.io/upload_images/4360245-015a6e823496b8b3.png?imageMogr2/auto-orient/strip|imageView2/2/w/952/format/webp)

# 优缺点与用途

## 优点

- 空间占用极低
- 能够保密数据
- 时间效率也较高 O(k)
- 哈希函数之间相互独立，可以在硬件指令层面并行计算

## 缺点

- 存在假阳性的概率
- 只能插入和查询元素，不能删除元素

# Redis BloomFilter

Redis 中使用 BloomFilter 需要安装 RedisBloom 插件，下载源码编译后生成一个 rebloom.so 文件，然后需要在在 Redis 的配置文件 redis.conf 中加入该模块即可:

```
loadmodule /${path}/rebloom.so
```

添加 items

```
BF.ADD newFilter foo
BF.MADD myFilter foo bar baz
```

检查是否存在

```
BF.EXISTS newFilter foo
BF.MEXISTS myFilter foo nonexist bar
```

[redisbloom](https://oss.redislabs.com/redisbloom/Quick_Start/)

# 参考

[Bloom Filter](https://dongzl.github.io/2020/03/21/17-Bloom-Filter-Summary/index.html)
