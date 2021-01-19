func minOperations(nums []int, x int) int {
	var sum int
	for _, num := range nums {
		sum += num
	} 
	
	ans := -1
	
	target, left, current := sum - x, 0, 0
	for right := 0; right < len(nums); right++ {
		current += nums[right]
		for left <= right && current > target {
			current -= nums[left]
			left += 1
		}
		if current == target {
			ans = max(ans, right - left + 1)
		}
	}
	
	if ans == -1 {
		return -1
	}
	return len(nums) - ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}