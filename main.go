package main

import (
	"fmt"
	"math"
)

func main() {
	h := New()
	h.Insert(1)
	fmt.Println(h.values)
	h.Insert(2)
	fmt.Println(h.values)
	h.Insert(3)
	fmt.Println(h.values)

	h.Insert(10)
	fmt.Println(h.values)
	h.Insert(9)
	fmt.Println(h.values)

	fmt.Println(h.Delete())
	fmt.Println(h.values)
}

type Heap struct {
	values []int
}

func New() *Heap {
	return &Heap{values: []int{math.MinInt64}}
}

func (h *Heap) Insert(v int) {
	h.values = append(h.values, v)
	i := len(h.values) - 1
	for i/2 != 0 && h.values[i/2] < v {
		h.values[i] = h.values[i/2]
		i = i / 2
	}
	h.values[i] = v
}

func (h *Heap) Delete() int {
	if len(h.values) == 1 {
		panic("heap is empty")
	}
	max := h.values[1]
	last := h.values[len(h.values)-1]

	var i, child int
	for i = 1; i*2 < len(h.values); i = child {
		child = i * 2
		if child < len(h.values)-1 && h.values[child+1] > h.values[child] {
			child += 1
		}
		if last < h.values[child] {
			h.values[i] = h.values[child]
		} else {
			break
		}
	}
	h.values[i] = last
	h.values = h.values[:len(h.values)-1]
	return max
}
