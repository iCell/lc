func firstMissingPositive(nums []int) int {
	max, memo := math.MinInt32, make(map[int]bool)
	for _, num := range nums {
		if num > max {
			max = num
		}
		memo[num] = true
	}
	if max <= 0 {
		return 1
	}

	for i := 1; i <= max + 1; i++ {
		_, exist := memo[i]
		if exist {
			continue
		}
		return i
	}
	return -1
}