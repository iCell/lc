func nextPermutation(nums []int) {
    if len(nums) == 1 {
        return
    }

    start := len(nums) - 2
    for start >= 0 && nums[start] >= nums[start+1] {
        start--
    }
    if start == -1 {
        reverse(nums, 0)
        return
    }

    end := len(nums) - 1
    for end >= 0 && nums[end] <= nums[start] {
        end--
    }
    nums[start], nums[end] = nums[end], nums[start]
    reverse(nums, start+1)
}

func reverse(nums []int, start int) {
    left, right := start, len(nums)-1
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}