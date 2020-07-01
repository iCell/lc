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