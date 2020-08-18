func sortArray(nums []int) []int {
    quickSort(nums, 0, len(nums)-1)
    return nums
}

func quickSort(nums []int, left, right int) {
    if left > right {
        return
    }

    i, j, pivot := left, right, right
    for i < j {
        for i < j && nums[i] <= nums[pivot] {
            i++
        }
        for i < j && nums[j] >= nums[pivot] {
            j--
        }

        nums[i], nums[j] = nums[j], nums[i]
    }
    nums[j], nums[pivot] = nums[pivot], nums[j]
    pivot = j

    quickSort(nums, left, pivot-1)
    quickSort(nums, pivot+1, right)
}

