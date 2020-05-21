package main

import "fmt"

type CacheNode struct {
	Key   int
	Value int
	Pre   *CacheNode
	Next  *CacheNode
}

type LRUCache struct {
	Len   int
	Cap   int
	Head  *CacheNode
	Tail  *CacheNode
	nodes map[int]*CacheNode
}

func Constructor(capacity int) LRUCache {
	head := &CacheNode{}
	tail := &CacheNode{}
	head.Next = tail
	tail.Pre = head
	return LRUCache{
		Cap:   capacity,
		Head:  head,
		Tail:  tail,
		nodes: make(map[int]*CacheNode, capacity),
	}
}

func (this *LRUCache) append(node *CacheNode) {
	this.Len++

	this.Tail.Pre.Next = node
	node.Pre = this.Tail.Pre

	node.Next = this.Tail
	this.Tail.Pre = node
}

func (this *LRUCache) remove(node *CacheNode) {
	this.Len--

	pre := node.Pre
	next := node.Next
	pre.Next = next
	next.Pre = pre
}

func (this *LRUCache) move(node *CacheNode) {
	this.remove(node)
	this.append(node)
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
	if ok {
		node.Value = value
		this.move(node)
		return
	}

	if this.Len >= this.Cap {
		toRemove := this.Head.Next
		this.remove(toRemove)
		delete(this.nodes, toRemove.Key)
	}

	node = &CacheNode{
		Key:   key,
		Value: value,
	}
	this.nodes[key] = node
	this.append(node)
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
