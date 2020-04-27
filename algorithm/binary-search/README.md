1. 基本思路
对有序的数组 可以利用其有序的性质对半查找
``` go 
func search(array []int, key int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := low + (high-low)/2p
		if array[mid] == key {
			return mid
		} else if array[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
```

2. 相关数据结构
[二叉搜索树](https://baike.baidu.com/item/%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91)