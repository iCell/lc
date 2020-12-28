func checkSubarraySum(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		var sum int
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if j == i {
				continue
			}
			if (k != 0 && sum % k == 0) || (k == 0 && sum == 0) {
				return true
			}
		}
	}   
	return false
}

// sum[i] = c1 * k + r1
// sum[j] = c2 * k + r2
// sum[j] - sum[i-1] = (c2 - c1) * k + (r2 - r1)
// 若 r2 - r1 == 0, 则 sum[j] - sum[i-1] = c * k，必存在
// 因此，若存在两个前缀和对 k 取模的余数相等，则必存在
// 边界 1: i == 0，则 sum[i-1] = sum[-1] 的前缀和为 0

func checkSubarraySum(nums []int, k int) bool {
	memo := make(map[int]int)
	memo[0] = -1
	var sum int
	for i, num := range nums {
		sum += num
		
		var mod int
		if k == 0 {
			mod = sum
		} else {
			mod = sum % k    
		}
		
		j, exist := memo[mod]
		if exist {
			if i - j > 1 {
				return true   
			}
		} else {
			memo[mod] = i   
		}
	}
	return false
}

