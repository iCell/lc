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

	if len(this.minValues) == 0 {
		this.minValues = append(this.minValues, x)
		return
	}

	if this.minValues[len(this.minValues)-1] >= x {
		this.minValues = append(this.minValues, x)
	}
}

func (this *MinStack) Pop() {
	x, values := this.values[len(this.values)-1], this.values[:len(this.values)-1]
	min := this.minValues[len(this.minValues)-1]
	if min == x {
		this.minValues = this.minValues[:len(this.minValues)-1]
	}
	this.values = values
}

func (this *MinStack) Top() int {
	return this.values[len(this.values)-1]
}

func (this *MinStack) GetMin() int {
	return this.minValues[len(this.minValues)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
