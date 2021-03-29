package heap_sort

type Heap struct {
	size int
	data []int
}

// NewHeap 构造一个堆
func NewHeap(data []int) Heap {
	heap := Heap{
		size: len(data),
		data: data,
	}
	for i := heap.size / 2; i >= 0; i-- {
		heap.adjust(i)
	}
	return heap
}

// adjust 用于调整堆，保持堆的固有性质
func (heap *Heap) adjust(i int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < heap.size && heap.data[left] > heap.data[largest] {
		largest = left
	}
	if right < heap.size && heap.data[right] > heap.data[largest] {
		largest = right
	}
	if largest != i {
		heap.data[i], heap.data[largest] = heap.data[largest], heap.data[i]
		heap.adjust(largest)
	}
}

// Sort 堆排序
func Sort(data []int) {
	heap := NewHeap(data)
	for heap.size > 0 {
		heap.data[0], heap.data[heap.size-1] = heap.data[heap.size-1], heap.data[0]
		heap.size--
		heap.adjust(0)
	}
}
