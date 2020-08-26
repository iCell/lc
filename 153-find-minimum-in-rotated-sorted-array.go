func findMin(nums []int) int {
    if len(nums) == 1 || nums[len(nums)-1] > nums[0] {
        return nums[0]
    }

    left, right := 0, len(nums)-1
    for left < right {
        mid := (left + right) / 2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else if nums[mid] < nums[right] {
            right = mid
        }
    }

    return nums[right]
}