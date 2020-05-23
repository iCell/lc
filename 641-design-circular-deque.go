package main

import "fmt"

type MyCircularDeque struct {
	Cap    int
	values []int
}

func main() {
	circularDeque := Constructor(3)
	circularDeque.Debug()
	fmt.Println(circularDeque.InsertLast(1)) // return true
	circularDeque.Debug()
	fmt.Println(circularDeque.InsertLast(2)) // return true
	circularDeque.Debug()
	fmt.Println(circularDeque.InsertFront(3)) // return true
	circularDeque.Debug()
	fmt.Println(circularDeque.InsertFront(4)) // return false, the queue is full
	circularDeque.Debug()
	fmt.Println(circularDeque.GetRear())    // return 2
	fmt.Println(circularDeque.IsFull())     // return true
	fmt.Println(circularDeque.DeleteLast()) // return true
	circularDeque.Debug()
	fmt.Println(circularDeque.InsertFront(4)) // return true
	fmt.Println(circularDeque.GetFront())     // return 4
}

func (this *MyCircularDeque) Debug() {
	fmt.Println(this.values)
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		Cap:    k,
		values: make([]int, 0, k),
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.values = append([]int{value}, this.values...)
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.values = append(this.values, value)
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.values = this.values[1:len(this.values)]
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.values = this.values[:len(this.values)-1]
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.values[0]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.values[len(this.values)-1]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.values) == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.Cap <= len(this.values)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
