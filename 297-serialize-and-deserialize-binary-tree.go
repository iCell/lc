package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	root.Left = node2
	root.Right = node3
	node3.Left = node4
	node3.Right = node5

	c := Constructor()
	fmt.Println(c.serialize(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	Depth int
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var result [][]*TreeNode
	queue := &Queue{}
	queue.InQueue(root)

	for !queue.IsEmpty() {
		size := queue.Size()
		values := make([]*TreeNode, 0, size)
		hasValue := false
		for size > 0 {
			node := queue.DeQueue()
			// fmt.Println(node.Val)
			if node == nil {
				values = append(values, nil)
			} else {
				hasValue = true
				values = append(values, node)
				queue.InQueue(node.Left)
				queue.InQueue(node.Right)
			}
			size--
		}
		if hasValue {
			result = append(result, values)
		}
	}

	var r []string
	for _, values := range result {
		for _, v := range values {
			if v == nil {
				r = append(r, "null")
			} else {
				r = append(r, fmt.Sprintf("%d", v.Val))
			}
		}
	}

	this.Depth = len(result)
	return "[" + strings.Join(r, ",") + "]"
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	values := strings.Split(data, ",")
	depth := 1
	for depth <= this.Depth {
		int(math.Pow(2, float64(depth)))
	}
	return nil
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
type Queue struct {
	values []*TreeNode
}

func (q *Queue) InQueue(n *TreeNode) {
	q.values = append(q.values, n)
}

func (q *Queue) DeQueue() *TreeNode {
	if q.IsEmpty() {
		return nil
	}
	first, values := q.values[0], q.values[1:]
	q.values = values
	return first
}

func (q *Queue) IsEmpty() bool {
	return len(q.values) == 0
}

func (q *Queue) Size() int {
	return len(q.values)
}
