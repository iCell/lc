func findMin(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    left, right := 0, len(nums)-1
    if nums[right] > nums[left] {
        return nums[left]
    }

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] > nums[mid+1] {
            return nums[mid+1]
        }
        if nums[mid] < nums[mid-1] {
            return nums[mid]
        }
        if nums[mid] > nums[left] {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}

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