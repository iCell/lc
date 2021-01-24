func minSubArrayLen(s int, nums []int) int {
	left, sum, ans := 0, 0, math.MaxInt32
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= s {
			ans = min(ans, right - left + 1)
			sum -= nums[left]
			left += 1
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}