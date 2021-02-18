type MyCircularQueue struct {
	head int
	count int
	capacity int
	
	queue []int
}


func Constructor(k int) MyCircularQueue {
	return MyCircularQueue {
		capacity: k,
		queue: make([]int, k),
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.count == this.capacity {
		return false
	}
	this.queue[(this.head + this.count) % this.capacity] = value
	this.count += 1
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.count == 0 {
		return false
	}
	this.head = (this.head + 1) % this.capacity
	this.count -= 1
	return true
}


func (this *MyCircularQueue) Front() int {
	if this.count == 0 {
		return -1
	}
	return this.queue[this.head]
}


func (this *MyCircularQueue) Rear() int {
	if this.count == 0 {
		return -1
	}
	return this.queue[(this.head + this.count - 1) % this.capacity]
}


func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}


func (this *MyCircularQueue) IsFull() bool {
	return this.capacity == this.count
}


-----------------

type Node struct {
	Value int
	Next *Node
}

func NewNode(v int) *Node {
	return &Node{Value: v}
}

type MyCircularQueue struct {
	head *Node
	tail *Node
	
	count int
	capacity int
}


func Constructor(k int) MyCircularQueue {
	return MyCircularQueue {capacity: k}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.count == this.capacity {
		return false
	}
	
	node := NewNode(value)
	if this.count == 0 {
		this.head = node
		this.tail = node    
	} else {
		this.tail.Next = node
		this.tail = this.tail.Next
	}
	this.count += 1
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.count == 0 {
		return false
	}
	this.head = this.head.Next
	this.count -= 1
	return true
}


func (this *MyCircularQueue) Front() int {
	if this.count == 0 {
		return -1
	}
	return this.head.Value
}


func (this *MyCircularQueue) Rear() int {
	if this.count == 0 {
		return -1
	}
	return this.tail.Value
}


func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}


func (this *MyCircularQueue) IsFull() bool {
	return this.capacity == this.count
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */