package problem0239

func maxSlidingWindow(nums []int, k int) []int {
	deque := &Deque{}
	result := make([]int, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {
		for !deque.IsEmpty() && i-deque.First() >= k {
			deque.Shift()
		}
		for !deque.IsEmpty() && nums[i] > nums[deque.Last()] {
			deque.Pop()
		}
		deque.Push(i)

		if i >= k-1 {
			result[i-k+1] = nums[deque.First()]
		}
	}
	return result
}

type Deque struct {
	data []int
}

func (d *Deque) Push(num int) {
	d.data = append(d.data, num)
}

func (d *Deque) Pop() {
	d.data = d.data[0 : len(d.data)-1]
}

func (d *Deque) Shift() {
	d.data = d.data[1:len(d.data)]
}

func (d *Deque) First() int {
	return d.data[0]
}

func (d *Deque) Last() int {
	return d.data[len(d.data)-1]
}

func (d *Deque) IsEmpty() bool {
	return len(d.data) == 0
}
