package main

import (
	"fmt"
)

type Heap struct {
	values []int
}

func main() {
	h := New()
	h.Insert(1)
	h.Insert(2)
	h.Insert(3)
	h.Insert(9)
	h.Insert(10)

	fmt.Println(h.values)
	fmt.Println(h.Delete())
	fmt.Println(h.values)
}

func New() *Heap {
	return &Heap{}
}

func (h *Heap) Insert(v int) {
	h.values = append(h.values, v)
	h.sinkUp()
}

func (h *Heap) Delete() int {
	h.values[0], h.values[len(h.values)-1] = h.values[len(h.values)-1], h.values[0]
	x, v := h.values[len(h.values)-1], h.values[:len(h.values)-1]
	h.values = v
	h.sinkDown()
	return x
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
	for parentIndex >= 0 && current > parent {
		h.values[currentIndex], h.values[parentIndex] = parent, current
		currentIndex = parentIndex
		parentIndex = h.parent(currentIndex)
		current, parent = h.values[currentIndex], h.values[parentIndex]
	}
}

func (h *Heap) sinkDown() {

}
