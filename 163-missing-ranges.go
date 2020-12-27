func findMissingRanges(nums []int, lower int, upper int) []string {
	if len(nums) == 0 {
		return []string{composeRange(lower, upper)}
	}
	
	var ans []string
	if nums[0] > lower {
		ans = append(ans, composeRange(lower, nums[0] - 1))
	}
	for i := 1; i < len(nums); i++ {
		left, right := nums[i-1], nums[i]
		if right - left > 1 {
			ans = append(ans, composeRange(left + 1, right - 1))   
		}
	}
	if nums[len(nums)-1] < upper {
		ans = append(ans, composeRange(nums[len(nums)-1] + 1, upper))
	}
	return ans
}

func composeRange(low, high int) string {
	if low == high {
		return fmt.Sprintf("%d", low)
	}
	return fmt.Sprintf("%d->%d", low, high)
}