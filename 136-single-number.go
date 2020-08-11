import "sort"

func singleNumber(nums []int) int {
    sort.Ints(nums)

    if len(nums) > 1 && nums[0] < nums[1] {
        return nums[0]
    }
    if len(nums) > 1 && nums[len(nums)-2] < nums[len(nums)-1] {
        return nums[len(nums)-1]
    }
    i := 1
    for i < len(nums) {
        if nums[i]-nums[i-1] != 0 {
            return nums[i-1]
        }
        i += 2
    }
    return nums[0]
}

func singleNumber2(nums []int) int {
    r := 0
    for _, num := range nums {
        r = r ^ num
    }
    return r
}