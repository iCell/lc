// abc def
// cba fed
// def abc
func rotate(nums []int, k int) {
    k = k % len(nums)
    index := len(nums) - 1 - k

    left, right := 0, index
    helper(nums, left, right)

    left, right = index+1, len(nums)-1
    helper(nums, left, right)

    left, right = 0, len(nums)-1
    helper(nums, left, right)
}

func helper(nums []int, left, right int) {
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}