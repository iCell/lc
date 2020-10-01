import (
    "mat    "sort"
h"
)

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)

    var best = math.MaxInt32

    update := func(cur int) {
        if abs(cur-target) < abs(best-target) {
            best = cur
        }
    }

    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        j, k := i+1, len(nums)-1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if sum == target {
                return target
            }

            update(sum)
            if sum > target {
                k--
            } else {
                j++
            }
        }
    }

    return best
}

func abs(x int) int {
    if x < 0 {
        return -1 * x
    }
    return x
}