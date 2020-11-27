# 定义
Bit-map就是用一个bit位来标记某个元素对应的Value， 而Key即是该元素。由于采用了Bit为单位来存储数据，因此在存储空间方面，可以大大节省。
# 用途
使用bit数组来表示某些元素是否存在
# Tips
```
 n = byte * sizeof(unit) + bit
```