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

