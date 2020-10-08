import "sort"

func findPairs(nums []int, k int) int {
    sort.Ints(nums)

    var result int
    left, right := 0, 1
    for left < len(nums) && right < len(nums) {
        if left == right || nums[right]-nums[left] < k {
            right++
        } else if nums[right]-nums[left] > k {
            left++
        } else {
            result++
            left++
            for left < len(nums) && nums[left] == nums[left-1] {
                left++
            }
        }
    }

    return result
}

func findPairs(nums []int, k int) int {
    sort.Ints(nums)

    var result int
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for j := i + 1; j < len(nums); j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            if abs(nums[j], nums[i]) == k {
                result++
            }
        }
    }
    return result
}

func abs(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}

func findPairs(nums []int, k int) int {
    memo := make(map[int]int, len(nums))
    for _, num := range nums {
        memo[num] += 1
    }

    var result int
    for num, count := range memo {
        if k == 0 && count > 1 {
            result++
        } else if k > 0 {
            _, exist := memo[num+k]
            if exist {
                result++
            }
        }
    }
    return result
}
