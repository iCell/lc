func subarraySum(nums []int, k int) int {
    var count int
    // left bound
    for left := 0; left < len(nums); left++ {
        // right bound
        for right := left; right < len(nums); right++ {
            var sum int
            // check all possibilities within left and right bound
            for i := left; i <= right; i++ {
                sum += nums[i]
            }
            if sum == k {
                count++
            }
        }
    }
    return count
}

// sum[0-j] = sum[0-(j-1)] + nums[j]
func subarraySum(nums []int, k int) int {
    var count int
    for i := 0; i < len(nums); i++ {
        var sum int
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            if sum == k {
                count++
            }
        }
    }
    return count
}

// sum[i, j] = sum[0, j] - sum[0, i-1]
// k = sum[j] - sum[i-1]
// sum[j] - k = sum[i-1]  j 减去 k 的和正好为 i-1的，那么意味着存在一个前缀和正好为 k
// 当 i = 0 时，sum[i-1] = sum[-1] = 0,
func subarraySum(nums []int, k int) int {
    memo := make(map[int]int)
    memo[0] = 1

    var sum int
    var count int
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        count += memo[sum-k]
        memo[sum] += 1
    }
    return count
}

