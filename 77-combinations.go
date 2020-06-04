package main

import "fmt"

func main() {
	fmt.Println(combine(2, 2))
}

func combine(n int, k int) [][]int {
	if n == k || k == 0 {
		var r []int
		for i := 1; i <= k; i++ {
			r = append(r, i)
		}
		return [][]int{r}
	}

	var result [][]int
	for _, r := range combine(n-1, k-1) {
		result = append(result, append(r, n))
	}
	result = append(result, combine(n-1, k)...)
	return result
}
