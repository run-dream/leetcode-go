package problem0155

// MinStack 最小栈
type MinStack struct {
	data     []int
	size     int
	capacity int
	min      int
}

// Constructor  构造器
func Constructor() MinStack {
	return MinStack{
		data:     make([]int, 10),
		capacity: 10,
		size:     0,
		min:      0,
	}
}

// Push 入栈
func (stack *MinStack) Push(x int) {
	if stack.size >= stack.capacity {
		stack.capacity *= 2
		tmp := make([]int, stack.capacity)
		copy(tmp, stack.data)
		stack.data = tmp
	}
	if stack.size == 0 || stack.min > x {
		stack.min = x
	}
	stack.data[stack.size] = x
	stack.size++
}

// Pop 出栈
func (stack *MinStack) Pop() {
	top := stack.Top()
	stack.size--
	if top <= stack.min {
		min := stack.data[0]
		for i := 1; i < stack.size; i++ {
			if min > stack.data[i] {
				min = stack.data[i]
			}
		}
		stack.min = min
	}
}

// Top 栈顶
func (stack *MinStack) Top() int {
	return stack.data[stack.size-1]
}

// GetMin 获取栈中最小元素
func (stack *MinStack) GetMin() int {
	return stack.min
}
