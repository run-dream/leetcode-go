[Implement Stack using Queues](https://leetcode.com/problems/implement-stack-using-queues/)

### 思路
- 使用两个栈 push O(1) pop O(n)  
 pop时，用令一个队列存储移除的元素，并交换队列
- 使用一个栈 push O(n) pop O(1) 
 push时，直接逆序存储  