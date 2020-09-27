func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }

    end := binarySearch(nums, target)
    start := binarySearch(nums, target-1) + 1
    if start >= 0 && start <= end && end < len(nums) {
        return []int{start, end}
    }
    return []int{-1, -1}
}

// 用 left　<= right　这种计算方法最终得出来的 index 是最右边符合条件的 index
// 如果需要得到左边界，则可以找 target - 1 的下表，找到后 index + 1 即为左边界
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := (left + right) / 2
        if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return right
}
