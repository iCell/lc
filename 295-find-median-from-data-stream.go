type MedianFinder struct {
	MinHeap *MinHeap
	MaxHeap *MaxHeap
	count   int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		MinHeap: NewMinHeap(),
		MaxHeap: NewMaxHeap(),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.count += 1
	this.MaxHeap.Insert(num)
	max := this.MaxHeap.Delete()
	this.MinHeap.Insert(max)
	if len(this.MinHeap.values) > len(this.MaxHeap.values) {
		min := this.MinHeap.Delete()
		this.MaxHeap.Insert(min)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.count%2 == 0 {
		return (float64(this.MinHeap.values[0]) + float64(this.MaxHeap.values[0])) / 2
	}
	return float64(this.MaxHeap.values[0])
}

type MaxHeap struct {
	values []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{values: []int{}}
}

func (h *MaxHeap) Insert(v int) {
	h.values = append(h.values, v)
	h.sinkUp()
}

func (h *MaxHeap) Delete() int {
	max, last := h.values[0], h.values[len(h.values)-1]
	h.values[0], h.values = last, h.values[:len(h.values)-1]
	if len(h.values) > 1 {
		h.sinkDown()
	}
	return max
}

func (h *MaxHeap) sinkUp() {
	index := len(h.values) - 1
	for parent(index) >= 0 && h.values[index] > h.values[parent(index)] {
		h.values[index], h.values[parent(index)] = h.values[parent(index)], h.values[index]
		index = parent(index)
	}
}

func (h *MaxHeap) sinkDown() {
	var index int
	for leftChild(index) < len(h.values) {
		left, right := leftChild(index), rightChild(index)
		if right < len(h.values) && h.values[left] < h.values[right] {
			left = right
		}
		if h.values[left] < h.values[index] {
			break
		}
		h.values[left], h.values[index] = h.values[index], h.values[left]
		index = left
	}
}

type MinHeap struct {
	values []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{values: []int{}}
}

func (h *MinHeap) Insert(v int) {
	h.values = append(h.values, v)
	h.sinkUp()
}

func (h *MinHeap) Delete() int {
	max, last := h.values[0], h.values[len(h.values)-1]
	h.values[0], h.values = last, h.values[:len(h.values)-1]
	if len(h.values) > 1 {
		h.sinkDown()
	}
	return max
}

func (h *MinHeap) sinkUp() {
	index := len(h.values) - 1
	for parent(index) >= 0 && h.values[index] < h.values[parent(index)] {
		h.values[index], h.values[parent(index)] = h.values[parent(index)], h.values[index]
		index = parent(index)
	}
}

func (h *MinHeap) sinkDown() {
	var index int
	for leftChild(index) < len(h.values) {
		left, right := leftChild(index), rightChild(index)
		if right < len(h.values) && h.values[left] > h.values[right] {
			left = right
		}
		if h.values[left] > h.values[index] {
			break
		}
		h.values[left], h.values[index] = h.values[index], h.values[left]
		index = left
	}
}

func leftChild(index int) int {
	return index*2 + 1
}

func rightChild(index int) int {
	return index*2 + 2
}

func parent(index int) int {
	return (index - 1) / 2
}