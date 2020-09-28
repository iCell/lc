type Value struct {
    Length int
    Max    int
}

// if current[i] > dp[i-1].Max: dp[i] = dp[i-1].Length + 1
// else dp[i] = 1

// dp[0] = 1

func findLengthOfLCIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    dp := make([]*Value, len(nums), len(nums))
    dp[0] = &Value{Length: 1, Max: nums[0]}

    result := dp[0].Length

    for i := 1; i < len(dp); i++ {
        last := dp[i-1]
        if nums[i] > last.Max {
            dp[i] = &Value{
                Length: dp[i-1].Length + 1,
                Max:    nums[i],
            }
        } else {
            dp[i] = &Value{
                Length: 1,
                Max:    nums[i],
            }
        }
        if dp[i].Length > result {
            result = dp[i].Length
        }
    }

    return result
}