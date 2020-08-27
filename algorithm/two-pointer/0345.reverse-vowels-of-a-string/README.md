[Reverse Vowels of a String](https://leetcode.com/problems/reverse-vowels-of-a-string/)

使用双指针，指向待反转的两个元音字符，一个指针从头向尾遍历，一个指针从尾到头遍历。

可以使用初始化好的bool数组进一步优化速度
``` go
   v := make([]bool, 58)
    v['a' - 'A'] = true
    v['e' - 'A'] = true
    v['i' - 'A'] = true
    v['o' - 'A'] = true
    v['u' - 'A'] = true
    v['A' - 'A'] = true
    v['E' - 'A'] = true
    v['I' - 'A'] = true
    v['O' - 'A'] = true
    v['U' - 'A'] = true
```