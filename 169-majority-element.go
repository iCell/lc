func majorityElement(nums []int) int {
	memo := make(map[int]int, len(nums))
	for _, num := range nums {
		memo[num] += 1
		if memo[num] > len(nums) / 2 {
			return num
		}
	}
	return -1
}