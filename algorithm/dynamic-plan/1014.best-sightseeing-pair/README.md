[Best Sightseeing Pai](https://leetcode.com/problems/best-sightseeing-pair/)


### 思路 
遍历数组 O(n*n) 这样会超时

注意到 a[i] + a[j] + j - i = a[j]+ j + a[i] - i = right(i) + left(j)
确定了起点以后，其实终止点的最大值是和起点无关的
left(i) = max(a[i]-i, left(i+1))
考虑从后向前遍历 
``` go
left[n-1] = a[n-1]-(n-1)
result := 0
for i := n - 2; i >= 0;i--{
    left[i] = max(left[i+1], a[i]-i)
    result = max(result, a[i]+i+left[i])
}
```
