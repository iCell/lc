package main

import (
	"fmt"
)

func main() {
	h := New()
	h.Insert(2)
	h.Print()
	h.Insert(1)
	h.Print()
	// h.Insert(9)
	// h.Print()
	// h.Insert(4)
	// h.Print()
	// h.Insert(7)
	// h.Print()
	// h.Insert(10)
	// h.Print()

	fmt.Println(h.Delete())
	h.Print()
	// fmt.Println(h.Delete())
	// h.Print()
	// fmt.Println(h.Delete())
	// h.Print()
	// fmt.Println(h.Delete())
}

type Heap struct {
	values []int
}

func New() *Heap {
	return &Heap{values: []int{}}
}

func (h *Heap) Insert(v int) {
	h.values = append(h.values, v)
	h.sinkUp()
}

func (h *Heap) Delete() int {
	max, last := h.values[0], h.values[len(h.values)-1]
	h.values[0], h.values = last, h.values[:len(h.values)-1]
	if len(h.values) > 1 {
		h.sinkDown()
	}
	return max
}

func (h *Heap) sinkUp() {
	index := len(h.values) - 1
	for parent(index) >= 0 && h.values[index] > h.values[parent(index)] {
		h.values[index], h.values[parent(index)] = h.values[parent(index)], h.values[index]
		index = parent(index)
	}
}

func (h *Heap) sinkDown() {
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

func leftChild(index int) int {
	return index*2 + 1
}

func rightChild(index int) int {
	return index*2 + 2
}

func parent(index int) int {
	return (index - 1) / 2
}
