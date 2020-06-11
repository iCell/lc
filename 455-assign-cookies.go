package main

import "sort"

func findContentChildren(children []int, cookies []int) int {
	sort.Ints(children)
	sort.Ints(cookies)
	num := 0
	i, j := 0, 0
	for {
		if j >= len(cookies) || i >= len(children) {
			break
		}
		child := children[i]
		cookie := cookies[j]

		if cookie >= child {
			num++
			j++
			i++
		} else {
			j++
		}
	}
	return num
}
