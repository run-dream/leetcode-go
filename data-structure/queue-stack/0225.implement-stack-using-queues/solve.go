package problem0225

// MyQueue 队列实现
type MyQueue struct {
	data     []int
	size     int
	capacity int
}

// InitQueue 初始化队列
func InitQueue() MyQueue {
	return MyQueue{
		data:     make([]int, 10),
		size:     0,
		capacity: 10,
	}
}

// Push 入队列
func (queue *MyQueue) Push(x int) {
	if queue.size >= queue.capacity {
		tmp := make([]int, queue.capacity*2)
		copy(tmp, queue.data)
		queue.data = tmp
		queue.capacity = queue.capacity * 2
	}
	queue.data[queue.size] = x
	queue.size++
}

// Pop 出队列
func (queue *MyQueue) Pop() int {
	tmp := queue.data[0]
	for i := 0; i < queue.size; i++ {
		queue.data[i] = queue.data[i+1]
	}
	queue.size--
	return tmp
}

// Top 栈顶元素
func (queue *MyQueue) Top() int {
	return queue.data[0]
}

// Empty 队列是否为空
func (queue *MyQueue) Empty() bool {
	return queue.size == 0
}

// MyStack 栈
type MyStack struct {
	queue MyQueue
}

// Constructor 构造器
func Constructor() MyStack {
	return MyStack{
		queue: InitQueue(),
	}
}

// Push 入栈
func (stack *MyStack) Push(x int) {
	stack.queue.Push(x)
	for i := 1; i < stack.queue.size; i++ {
		stack.queue.Push(stack.queue.Pop())
	}
}

// Pop 出栈
func (stack *MyStack) Pop() int {
	return stack.queue.Pop()
}

// Top 队列头部元素
func (stack *MyStack) Top() int {
	return stack.queue.Top()
}

// Empty 队列是否为空
func (stack *MyStack) Empty() bool {
	return stack.queue.Empty()
}
