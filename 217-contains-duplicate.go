import "sort"

func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
    for i, num := range nums {
        if i == 0 {
            continue
        }
        if num == nums[i-1] {
            return true
        }
    }
    return false
}

func containsDuplicate2(nums []int) bool {
    hash := make(map[int]bool, len(nums))
    for _, num := range nums {
        _, exist := hash[num]
        if exist {
            return true
        }
        hash[num] = true
    }
    return false
}