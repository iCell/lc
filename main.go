package main

import (
	"fmt"
	"math"
)

func main() {
	h := New()
	h.Insert(6)
	h.Print()
	h.Insert(1)
	h.Print()
	h.Insert(9)
	h.Print()
	h.Insert(4)
	h.Print()
	h.Insert(7)
	h.Print()
	h.Insert(10)
	h.Print()

	fmt.Println(h.Delete())
	h.Print()
	fmt.Println(h.Delete())
	h.Print()
	fmt.Println(h.Delete())
	h.Print()
	fmt.Println(h.Delete())
}

type Heap struct {
	values []int
}

func New() *Heap {
	return &Heap{values: []int{math.MinInt64}}
}

func (h *Heap) Print() {
	fmt.Println(h.values)
}

func (h *Heap) Insert(v int) {
	h.values = append(h.values, v)
	h.sinkUp(len(h.values) - 1)
}

func (h *Heap) Delete() int {
	if len(h.values) == 1 {
		panic("heap is empty")
	}
	max, last := h.values[1], h.values[len(h.values)-1]
	h.values[1], h.values = last, h.values[:len(h.values)-1]
	h.sinkDown(1)
	return max
}

func (h *Heap) sinkUp(index int) {
	for index/2 > 0 && h.values[index] > h.values[index/2] {
		h.values[index], h.values[index/2] = h.values[index/2], h.values[index]
		index = index / 2
	}
}

func (h *Heap) sinkDown(index int) {
	for index*2 < len(h.values) {
		older := index * 2
		right := older + 1
		if right < len(h.values) && h.values[older] < h.values[right] {
			older = right
		}
		if h.values[older] < h.values[index] {
			break
		}
		h.values[older], h.values[index] = h.values[index], h.values[older]
		index = older
	}
}
