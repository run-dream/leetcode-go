1. 基本思路
解决一个回溯问题，实际上就是一个决策树的遍历过程。你只需要思考 3 个问题：
1、路径：也就是已经做出的选择。
2、选择列表：也就是你当前可以做的选择。
3、结束条件：也就是到达决策树底层，无法再做选择的条件。

2. 适用范围
- 全排列
- N皇后

3. 伪代码
``` go
result := []string{}
func backtrack(path, options){
    if meetFinishCondition(){
        result = append(result, path)
        return
    }
    for _, opt := range options{
        deal(opt)
        backtrack(path, options)
        revert(opt)
    }
}
```