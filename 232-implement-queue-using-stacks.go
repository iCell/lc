type MyQueue struct {
	front int
	stack1 *Stack
	stack2 *Stack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue {
		stack1: NewStack(),
		stack2: NewStack(),
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	if this.stack1.IsEmpty() {
		this.front = x
	}
	this.stack1.Push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.stack2.IsEmpty() {
		for !this.stack1.IsEmpty() {
			v := this.stack1.Pop()
			this.stack2.Push(v)
		}
	}
	return this.stack2.Pop()
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.stack2.IsEmpty() {
		return this.front
	}
	return this.stack2.Peek()
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stack1.IsEmpty() && this.stack2.IsEmpty()
}

type Stack struct {
	values []int
}

func NewStack() *Stack {
	return &Stack{values: make([]int, 0)}
}

func (s *Stack) Push(v int) {
	s.values = append(s.values, v)
}

func (s *Stack) Pop() int {
	pop, temp := s.values[len(s.values) - 1], s.values[:len(s.values) - 1]
	s.values = temp
	return pop
}

func (s *Stack) Peek() int {
	return s.values[len(s.values) - 1]   
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */