import "sort"

func triangleNumber(nums []int) int {
    sort.Ints(nums)

    var result int
    for i := len(nums) - 1; i >= 2; i-- {
        l, r := 0, i-1
        for l < r {
            if nums[l]+nums[r] > nums[i] {
                result += r - l
                r--
            } else {
                l++
            }
        }
    }

    return result
}