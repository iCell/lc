import "sort"

func findKthLargest(nums []int, k int) int {
    k = len(nums) - k
    return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, left, right int, k int) int {
    if left >= right {
        return nums[left]
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
    nums[pivot], nums[j] = nums[j], nums[pivot]
    pivot = j

    if k == pivot {
        return nums[k]
    } else if k > pivot {
        return quickSelect(nums, pivot+1, right, k)
    } else {
        return quickSelect(nums, left, pivot-1, k)
    }
}

func findKthLargest2(nums []int, k int) int {
    sort.Ints(nums)
    return nums[len(nums)-k]
}