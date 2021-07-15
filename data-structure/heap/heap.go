package heap

type Heap struct {
	size   int
	values []int
}

func NewHeap() *Heap {
	return &Heap{
		size:   0,
		values: []int{},
	}
}

func (this *Heap) Top() int {
	return this.values[0]
}

func (this *Heap) Pop() int {
	this.values[0], this.values[this.size-1] = this.values[this.size-1], this.values[0]
	result := this.values[this.size-1]
	this.values = this.values[0:this.size]
	this.size--
	this.sunk(0)
	return result
}

func (this *Heap) Push(val int) {
	this.values = append(this.values, val)
	this.size++
	this.swim(this.size - 1)
}

func (this *Heap) swim(node int) {
	if node == 0 {
		return
	}
	root := (node - 1) / 2
	if this.values[root] < this.values[node] {
		this.values[root], this.values[node] = this.values[node], this.values[root]
		this.swim(root)
	}
}

func (this *Heap) sunk(root int) {
	left := root*2 + 1
	right := root*2 + 2
	largest := root
	if left < this.size && this.values[left] > this.values[largest] {
		largest = left
	}
	if right < this.size && this.values[right] > this.values[largest] {
		largest = right
	}
	if largest == root {
		return
	}
	this.values[largest], this.values[root] = this.values[root], this.values[largest]
	this.sunk(largest)
}
