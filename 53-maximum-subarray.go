import "math"

func maxSubArray(nums []int) int {
    var current int
    max := math.MinInt64
    for _, num := range nums {
        if current < 0 {
            current = num
        } else {
            current += num
        }
        max = int(math.Max(float64(max), float64(current)))
    }
    return max
}

func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    dp := make([]int, len(nums), len(nums))
    dp[0] = nums[0]

    sum := dp[0]
    for i := 1; i < len(dp); i++ {
        dp[i] = max(dp[i-1]+nums[i], nums[i])
        sum = max(sum, dp[i])
    }

    return sum
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

