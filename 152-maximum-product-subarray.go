// dp[i][0] min value when i
// dp[i][1] max value when i
// if num[i] >= 0,
//  dp[i][0] = min(nums[i], nums[i] * dp[i-1][0])
//  dp[i][1] = max(nums[i], nums[i] * dp[i-1][1])
// if num[i] < 0,
//  dp[i][0] = min(nums[i], nums[i] * dp[i-1][1])
//  dp[i][1] = max(nums[i], nums[i] * dp[i-1][0])

func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    dp := make([][]int, len(nums), len(nums))
    for i := range dp {
        dp[i] = make([]int, 2, 2)
    }
    dp[0][0] = nums[0]
    dp[0][1] = nums[0]

    for i := 1; i < len(nums); i++ {
        num := nums[i]
        if num >= 0 {
            dp[i][0] = min(num, num*dp[i-1][0])
            dp[i][1] = max(num, num*dp[i-1][1])
        } else {
            dp[i][0] = min(num, num*dp[i-1][1])
            dp[i][1] = max(num, num*dp[i-1][0])
        }
    }

    result := dp[0][1]
    for i := 1; i < len(nums); i++ {
        result = max(result, dp[i][1])
    }

    return result
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}