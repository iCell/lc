func nthUglyNumber(n int) int {
	heap, min := New(), -1
	for n != 0 {
		v := heap.Delete()
		if min == v {
			continue
		}
		min = v
		heap.Insert(min * 2)
		heap.Insert(min * 3)
		heap.Insert(min * 5)
		n--
	}
	return min
}

type Heap struct {
	values []int
}

func New() *Heap {
	return &Heap{
		values: []int{1},
	}
}

func (h *Heap) Insert(v int) {
	h.values = append(h.values, v)
	h.sinkUp()
}

func (h *Heap) Delete() int {
	min, last := h.values[0], h.values[len(h.values)-1]
	h.values[0], h.values = last, h.values[:len(h.values)-1]
	if len(h.values) > 1 {
		h.sinkDown()
	}
	return min
}

func (h *Heap) sinkUp() {
	index := len(h.values) - 1
	for parent(index) >= 0 && h.values[index] < h.values[parent(index)] {
		h.values[index], h.values[parent(index)] = h.values[parent(index)], h.values[index]
		index = parent(index)
	}
}

func (h *Heap) sinkDown() {
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

