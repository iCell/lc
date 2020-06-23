package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num > 0 {
			break
		}
		if i > 0 && num == nums[i-1] {
			continue
		}

		target := -num
		l, r := i+1, len(nums)-1
		for l < r {
			lv, rv := nums[l], nums[r]

			if lv+rv > target {
				r--
			} else if lv+rv < target {
				l++
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			}
		}
	}

	return result
}
