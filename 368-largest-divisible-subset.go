import "sort"

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	sort.Ints(nums)

	dp := make([][]int, len(nums), len(nums))
	dp[0] = []int{nums[0]}

	result := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = []int{nums[i]}

		for j := 0; j < i; j++ {
			// 存在整除子集中的最大元素的数字
			if nums[i]%nums[j] != 0 {
				continue
			}
			if len(dp[i]) > len(dp[j]) {
				continue
			}
			copy(dp[i], dp[j])
			dp[i] = append(dp[i], nums[i])
		}
		if len(result) < len(dp[i]) {
			result = dp[i]
		}
	}
	return result
}

