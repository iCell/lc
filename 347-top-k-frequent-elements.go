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
	i := len(h.values) - 1
	for i/2 != 0 && h.values[i/2].Frenquent < v.Frenquent {
		h.values[i] = h.values[i/2]
		i = i / 2
	}
	h.values[i] = v
}

func (h *Heap) Delete() *Element {
	if len(h.values) == 1 {
		panic("heap is empty")
	}
	max := h.values[1]
	last := h.values[len(h.values)-1]

	var i, child int
	for i = 1; i*2 < len(h.values); i = child {
		child = i * 2
		if child < len(h.values)-1 && h.values[child+1].Frenquent > h.values[child].Frenquent {
			child += 1
		}
		if last.Frenquent < h.values[child].Frenquent {
			h.values[i] = h.values[child]
		} else {
			break
		}
	}
	h.values[i] = last
	h.values = h.values[:len(h.values)-1]
	return max
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
