package main

import "fmt"

func main() {

	cache := Constructor(2)
	cache.Put(1, 1)

	cache.Put(2, 2)

	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))

	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
