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

// sum[j] - k = sum[i]
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

