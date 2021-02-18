type Node struct {
	Val int
	Pre *Node
	Next *Node
}

func NewNode(val int) *Node {
	return &Node { Val: val }
}

type MyCircularDeque struct {
	count int
	capacity int
	head *Node
	tail *Node
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	head, tail := &Node{}, &Node{}
	head.Next = tail
	tail.Pre = head
	return MyCircularDeque { 
		head: head,
		tail: tail,
		capacity: k,
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.count == this.capacity {
		return false
	}
	
	node := NewNode(value)
	next := this.head.Next
	
	this.head.Next = node
	node.Pre = this.head
	
	node.Next = next
	next.Pre = node
	
	this.count++
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.count == this.capacity {
		return false
	}
	
	node := NewNode(value)
	pre := this.tail.Pre
	
	pre.Next = node
	node.Pre = pre
	
	node.Next = this.tail
	this.tail.Pre = node
	
	this.count++
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.count == 0 {
		return false
	}  
		
	cur := this.head.Next
	next := cur.Next
	next.Pre = this.head
	this.head.Next = next

	this.count--
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.count == 0 {
		return false
	}  
	
	cur := this.tail.Pre
	pre := cur.Pre
	this.tail.Pre = pre
	pre.Next = this.tail
	
	this.count--
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.count == 0 {
		return -1
	}
	return this.head.Next.Val
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.count == 0 {
		return -1
	}
	return this.tail.Pre.Val
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.count == 0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.count == this.capacity
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