// dp[0] = 1
// dp[i] = max(dp[0...j] if current > num[j] (j range in [0, i-1]), 1 if else) + 1

func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    dp := make([]int, len(nums), len(nums))
    dp[0] = 1

    var maxLength int
    for i, cur := range nums {
        if i == 0 {
            continue
        }

        for j := 0; j < i; j++ {
            var v int
            if cur > nums[j] {
                v = dp[j] + 1
            } else {
                v = 1
            }
            dp[i] = max(v, dp[i])
        }

        maxLength = max(maxLength, dp[i])
    }

    return max(maxLength, 1)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}