type BSTIterator struct {
    Stack   []*TreeNode
    Current *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
    return BSTIterator{
        Current: root,
        Stack:   make([]*TreeNode, 0),
    }
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    for this.Current != nil {
        this.Stack = append(this.Stack, this.Current)
        this.Current = this.Current.Left
    }

    pop, temp := this.Stack[len(this.Stack)-1], this.Stack[:len(this.Stack)-1]
    this.Stack = temp
    this.Current = pop.Right

    return pop.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return len(this.Stack) != 0 || this.Current != nil
}

