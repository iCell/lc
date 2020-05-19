package main

func moveZeroes(nums []int) {
	index := 0
	for i, num := range nums {
		if num != 0 {
			// swap num and nums[index]
			var temp = nums[index]
			nums[index] = num
			nums[i] = temp

			index += 1
		}
	}
}
