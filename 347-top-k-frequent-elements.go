package main

import (
	"fmt"
)

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
	return &Heap{values: []*Element{&Element{}}}
}

func (h *Heap) Insert(v *Element) {
	h.values = append(h.values, v)
	h.sinkUp(len(h.values) - 1)
}

func (h *Heap) Delete() *Element {
	if len(h.values) == 1 {
		panic("heap is empty")
	}
	max, last := h.values[1], h.values[len(h.values)-1]
	h.values[1], h.values = last, h.values[:len(h.values)-1]
	h.sinkDown(1)
	return max
}

func (h *Heap) sinkUp(index int) {
	for index/2 > 0 && h.values[index].Frenquent > h.values[index/2].Frenquent {
		h.values[index], h.values[index/2] = h.values[index/2], h.values[index]
		index = index / 2
	}
}

func (h *Heap) sinkDown(index int) {
	child, length := index*2, len(h.values)
	for child < length {
		if child+1 < length && h.values[child].Frenquent < h.values[child+1].Frenquent {
			child++
		}
		if h.values[index].Frenquent < h.values[child].Frenquent {
			break
		}
		h.values[index], h.values[child] = h.values[child], h.values[index]
		child = child * 2
	}
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num] += 1
	}

	heap := New()
	for k, v := range m {
		heap.Insert(&Element{
			Value:     k,
			Frenquent: v,
		})
	}

	var result []int
	for i := 0; i < k; i++ {
		result = append(result, heap.Delete().Value)
	}
	return result
}
