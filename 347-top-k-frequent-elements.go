package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

type Element struct {
	Value     int
	Frenquent int
}

type Heap struct {
	values []*Element
}

func New() *Heap {
	return &Heap{}
}

func (h *Heap) Insert(e *Element) {
	h.values = append(h.values, e)
	h.sinkUp()
}

func (h *Heap) left(i int) int {
	r := 2*i + 1
	if r >= len(h.values) {
		return -1
	}
	return r
}

func (h *Heap) right(i int) int {
	r := 2*i + 2
	if r >= len(h.values) {
		return -1
	}
	return r
}

func (h *Heap) parent(i int) int {
	r := (i - 1) / 2
	if r < 0 {
		return -1
	}
	return r
}

func (h *Heap) sinkUp() {
	currentIndex := len(h.values) - 1
	parentIndex := h.parent(currentIndex)

	current, parent := h.values[currentIndex], h.values[parentIndex]
	for current.Frenquent > parent.Frenquent {
		h.values[currentIndex], h.values[parentIndex] = parent, current
		currentIndex = parentIndex
		parentIndex = h.parent(currentIndex)
		current, parent = h.values[currentIndex], h.values[parentIndex]
	}
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num] += 1
	}

	fmt.Println(m)

	heap := New()
	for k, v := range m {
		heap.Insert(&Element{
			Value:     k,
			Frenquent: v,
		})
	}

	var result []int
	for i := 0; i < k; i++ {
		result = append(result, heap.values[i].Value)
	}
	return result
}
