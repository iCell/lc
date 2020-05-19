package main

import "fmt"

func main() {
	// node1 := &Node{Key: 1, Value: 1}
	// node2 := &Node{Key: 2, Value: 2}
	// node3 := &Node{Key: 3, Value: 3}

	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)

	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
}
