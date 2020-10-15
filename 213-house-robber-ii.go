// dp[i][0] 代表没有偷第 0 家,dp[i][1] 代表偷了第 0 家
// dp[i][0] 可以偷最后一家，则
// dp[i][1] 不可以偷最后一家，则

// 不偷最后一家：
// dp[i][0] = max(dp[i-2][0] + current, dp[i-1][0])
// dp[i][1] = max(dp[i-2][1] + current, dp[i-1][1])

// 偷最后一家时：
// dp[i][0] = max(dp[i-2][0] + current, dp[i-1][0])
// dp[i][1] = dp[i-1][1]

func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    if len(nums) == 2 {
        return max(nums[0], nums[1])
    }

    dp := make([][]int, len(nums), len(nums))
    for i := range dp {
        dp[i] = make([]int, 2, 2)
    }
    dp[0][0], dp[0][1] = 0, nums[0]
    dp[1][0], dp[1][1] = nums[1], nums[0]

    for i := 2; i < len(nums); i++ {
        if i == len(nums)-1 {
            dp[i][0] = max(dp[i-2][0]+nums[i], dp[i-1][0])
            dp[i][1] = dp[i-1][1]
        } else {
            dp[i][0] = max(dp[i-2][0]+nums[i], dp[i-1][0])
            dp[i][1] = max(dp[i-2][1]+nums[i], dp[i-1][1])
        }
    }

    return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}