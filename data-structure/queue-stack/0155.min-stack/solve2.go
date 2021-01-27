package problem0155

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
func (stack *MyStack) Pop() int {
	stack.size--
	return stack.data[stack.size]
}

// Top 栈顶
func (stack *MyStack) Top() int {
	return stack.data[stack.size-1]
}

// Empty 为空
func (stack *MyStack) Empty() bool {
	return stack.size == 0
}

// MinStack2 用两个栈来处理最小值
type MinStack2 struct {
	valueStack MyStack
	minStack   MyStack
	min        int
}

// Constructor2 构造函数
func Constructor2() MinStack2 {
	return MinStack2{
		valueStack: InitStack(),
		minStack:   InitStack(),
		min:        2<<32 - 1,
	}
}

// Push 入栈
func (stack *MinStack2) Push(x int) {
	stack.valueStack.Push(x)
	if x <= stack.min {
		stack.min = x
	}
	stack.minStack.Push(x)
}

// Pop 出栈
func (stack *MinStack2) Pop() {
	top := stack.valueStack.Pop()
	if top <= stack.min {
		if !stack.minStack.Empty() {
			stack.min = stack.minStack.Pop()
		} else {
			stack.min = 2<<32 - 1
		}
	}
}

// Top 栈顶
func (stack *MinStack2) Top() int {
	return stack.valueStack.Top()
}

// GetMin 获取栈中最小元素
func (stack *MinStack2) GetMin() int {
	return stack.min
}
