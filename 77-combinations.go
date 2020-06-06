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

func combine2(n int, k int) [][]int {
	var result [][]int
	helper(&result, n, k, 1, []int{})
	return result
}

func helper(res *[][]int, n int, k int, begin int, temp []int) {
	if len(temp) == k {
		*res = append(*res, append([]int{}, temp...))
		return
	}

	for i := begin; i <= n; i++ {
		temp = append(temp, i)
		helper(res, n+1, k, i+1, temp)
		temp = temp[:len(temp)-1]
	}
}
