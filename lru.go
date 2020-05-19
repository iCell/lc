package main

import "fmt"

type Node struct {
	Key   int
	Value int
	Pre   *Node
	Next  *Node
}

type LRUCache struct {
	Len   int
	Cap   int
	Head  *Node
	Tail  *Node
	nodes map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:   capacity,
		Head:  &Node{},
		Tail:  &Node{},
		nodes: make(map[int]*Node, capacity),
	}
}

func (this *LRUCache) add(node *Node) {
	this.nodes[node.Key] = node

	next := this.Head.Next

	if next != nil {
		node.Next = next
		next.Pre = node
	} else {
		this.Tail.Pre = node
	}

	this.Head.Next = node
	this.Len++
}

func (this *LRUCache) remove(node *Node) {
	delete(this.nodes, node.Key)

	pre := node.Pre
	next := node.Next

	if node == this.Tail.Pre {
		this.Tail.Pre = pre
	}
	if node == this.Head.Next {
		this.Head.Next = next
	}

	if pre != nil {
		pre.Next = next
	}
	if next != nil {
		next.Pre = pre
	}
	this.Len--
}

func (this *LRUCache) move(node *Node) {
	this.remove(node)
	this.add(node)
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.nodes[key]
	if !ok {
		return -1
	}
	this.move(node)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.nodes[key]
	if !ok {
		if this.Len == this.Cap {
			this.remove(this.Tail.Pre)
		}
		node := &Node{Key: key, Value: value}
		this.add(node)
		return
	}

	// update
	node.Value = value
	this.nodes[key] = node
	this.move(node)
}

func (this *LRUCache) Print() {
	fmt.Println("......")
	node := this.Head.Next
	for {
		if node == nil {
			return
		}
		fmt.Println(node.Value)
		node = node.Next
	}
}
