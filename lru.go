package main

import "fmt"

type Node struct {
	Key   int
	Value int
	Pre   *Node
	Next  *Node
}

type LRUCache struct {
	Cap   int
	Head  *Node
	Tail  *Node
	nodes map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Pre = head
	return LRUCache{
		Cap:   capacity,
		Head:  head,
		Tail:  tail,
		nodes: make(map[int]*Node, capacity),
	}
}

func (this *LRUCache) add(node *Node) {
	node.Next = this.Head.Next
	node.Pre = this.Head

	this.Head.Next.Pre = node
	this.Head.Next = node
}

func (this *LRUCache) remove(node *Node) {
	pre := node.Pre
	next := node.Next

	next.Pre = pre
	pre.Next = next
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
		if len(this.nodes) == this.Cap {
			delete(this.nodes, this.Tail.Pre.Key)
			this.remove(this.Tail.Pre)
		}
		newNode := &Node{Key: key, Value: value}
		this.add(newNode)
		this.nodes[key] = newNode
		return
	}

	node.Value = value
	this.nodes[key] = node
	this.move(node)
}

func (this *LRUCache) Print() {
	fmt.Println("......")
	node := this.Head
	for {
		if node == this.Tail {
			return
		}
		fmt.Println(node.Value)
		node = node.Next
	}
}
