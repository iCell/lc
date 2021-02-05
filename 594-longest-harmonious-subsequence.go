func findLHS(nums []int) int {
	memo := make(map[int]int)
	for _, num := range nums {
		memo[num] += 1
	}
	var ans int
	for _, num := range nums {
		cnt, ok := memo[num + 1]
		if ok {
			ans = max(ans, memo[num] + cnt)
		}
	}
	return ans
}

func findLHS(nums []int) int {
	var ans int
	
	for i := 0; i < len(nums); i++ {
		var count int
		var flag bool
		// 当前的 nums[i] 是最小值         
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			} else if nums[i] + 1 == nums[j] {
				count++
				flag = true
			}
		}
		if flag {
			ans = max(ans, count)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}