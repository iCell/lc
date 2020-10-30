import "math"

func maxDistToClosest(seats []int) int {
    var result int
    for i, v := range seats {
        if v == 1 {
            continue
        }
        left := math.MaxInt64
        for j := i - 1; j >= 0; j-- {
            if seats[j] == 1 {
                left = i - j
                break
            }
        }
        right := math.MaxInt64
        for j := i + 1; j < len(seats); j++ {
            if seats[j] == 1 {
                right = j - i
                break
            }
        }
        m := min(left, right)
        result = max(m, result)
    }
    return result
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}