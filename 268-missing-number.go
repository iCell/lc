import "sort"

func missingNumber(nums []int) int {
    var sum int
    for i := 0; i <= len(nums); i++ {
        sum += i
    }
    for _, num := range nums {
        sum -= num
    }
    return sum
}

func missingNumber(nums []int) int {
    memo := make(map[int]bool, len(nums)+1)
    for _, num := range nums {
        memo[num] = true
    }
    for i := 0; i <= len(nums); i++ {
        _, exist := memo[i]
        if !exist {
            return i
        }
    }
    return len(nums)
}

func missingNumber(nums []int) int {
    sort.Ints(nums)

    for i := 0; i < len(nums); i++ {
        if i != nums[i] {
            return i
        }
    }
    return len(nums)
}