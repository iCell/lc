package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge2(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, m)
	copy(temp, nums1[:m])

	index, index1, index2 := 0, 0, 0
	for index1 < m && index2 < n {
		left, right := temp[index1], nums2[index2]
		if left > right {
			nums1[index] = right
			index2++
		} else {
			nums1[index] = left
			index1++
		}
		index++
	}

	for index1 < m {
		nums1[index] = temp[index1]
		index++
		index1++
	}
	for index2 < n {
		nums1[index] = nums2[index2]
		index2++
		index++
	}
}
