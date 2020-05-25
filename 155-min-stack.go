package main

type MinStack struct {
	values    []int
	minValues []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.values = append(this.values, x)

}

func (this *MinStack) Pop() {
	this.values = this.values[:len(this.values)-1]
}

func (this *MinStack) Top() int {
	return this.values[0]
}

func (this *MinStack) GetMin() int {

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
