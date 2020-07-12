func countBits(num int) []int {
    var result []int
    for i := 0; i <= num; i++ {
        result = append(result, bitCount(i))
    }
    return result
}

func bitCount(num int) int {
    var count int
    for num != 0 {
        num = num & (num - 1)
        count++
    }
    return count
}

// giving a num,
// if is not odd number, the numbers of 1 is equal to num/2
func countBits2(num int) []int {
    dp := make([]int, num+1, num+1)
    for i := 1; i < len(dp); i++ {
        dp[i] = dp[i>>1] + i&1
    }
    return dp
}

