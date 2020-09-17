func topKFrequent(nums []int, k int) []int {
	memo := make(map[int]int)
	for _, num := range nums {
		memo[num] += 1
	}

	heap := New()
	for num, count := range memo {
		heap.Insert(&Frequent{
			Val: num,
			Fre: count,
		})
	}

	var result []int
	for i := 0; i < k; i++ {
		max := heap.Delete()
		result = append(result, max.Val)
	}
	return result
}

type Frequent struct {
	Val int
	Fre int
}

type Heap struct {
	values []*Frequent
}

func New() *Heap {
	return &Heap{values: []*Frequent{}}
}

func (h *Heap) Insert(v *Frequent) {
	h.values = append(h.values, v)
	h.sinkUp()
}

func (h *Heap) Delete() *Frequent {
	max, last := h.values[0], h.values[len(h.values)-1]
	h.values[0], h.values = last, h.values[:len(h.values)-1]
	if len(h.values) > 1 {
		h.sinkDown()
	}
	return max
}

func (h *Heap) sinkUp() {
	index := len(h.values) - 1
	for parent(index) >= 0 && h.values[index].Fre > h.values[parent(index)].Fre {
		h.values[index], h.values[parent(index)] = h.values[parent(index)], h.values[index]
		index = parent(index)
	}
}

func (h *Heap) sinkDown() {
	var index int
	for leftChild(index) < len(h.values) {
		left, right := leftChild(index), rightChild(index)
		if right < len(h.values) && h.values[left].Fre < h.values[right].Fre {
			left = right
		}
		if h.values[left].Fre < h.values[index].Fre {
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
