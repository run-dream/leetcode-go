package problem0232

import "fmt"

// MyStack 栈
type MyStack struct {
	data     []int
	size     int
	capacity int
}

// InitStack  初始栈
func InitStack() MyStack {
	return MyStack{data: make([]int, 10), size: 0, capacity: 10}
}

// Push 入队
func (stack *MyStack) Push(x int) {
	if stack.size >= stack.capacity {
		stack.capacity *= 2
		tmp := make([]int, stack.capacity)
		copy(tmp, stack.data)
		stack.data = tmp
	}
	stack.data[stack.size] = x
	stack.size++
}

// Pop 出队
func (stack *MyStack) Pop() (int, error) {
	if stack.size == 0 {
		return -1, fmt.Errorf("栈为空")
	}
	stack.size--
	return stack.data[stack.size], nil
}

// Peek 巅峰
func (stack *MyStack) Peek() (int, error) {
	if stack.size == 0 {
		return -1, fmt.Errorf("栈为空")
	}
	return stack.data[stack.size-1], nil
}

// Empty 空值
func (stack *MyStack) Empty() bool {
	return stack.size == 0
}

// MyQueue 自定义队列
type MyQueue struct {
	in  MyStack
	out MyStack
}

// Constructor 构造函数
func Constructor() MyQueue {
	return MyQueue{in: InitStack(), out: InitStack()}
}

// Push 入队
func (queue *MyQueue) Push(x int) {
	queue.in.Push(x)
}

// Pop 出队
func (queue *MyQueue) Pop() int {
	queue.InToOut()
	x, _ := queue.out.Pop()
	return x
}

// Peek 队顶
func (queue *MyQueue) Peek() int {
	queue.InToOut()
	x, _ := queue.out.Peek()
	return x
}

// Empty 队列是否为空
func (queue *MyQueue) Empty() bool {
	return queue.in.Empty() && queue.out.Empty()
}

// InToOut 入到出
func (queue *MyQueue) InToOut() {
	if queue.out.Empty() {
		for !queue.in.Empty() {
			x, _ := queue.in.Pop()
			queue.out.Push(x)
		}
	}
}
