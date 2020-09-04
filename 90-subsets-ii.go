import "sort"

func subsetsWithDup(nums []int) [][]int {
    current, result := make([]int, 0), make([][]int, 0)
    sort.Ints(nums)
    helper(nums, 0, current, &result)
    return result
}

func helper(nums []int, index int, current []int, result *[][]int) {
    *result = append(*result, append([]int{}, current...))

    for i := index; i < len(nums); i++ {
        if i > index && nums[i] == nums[i-1] {
            continue
        }
        current = append(current, nums[i])
        helper(nums, i+1, current, result)
        current = current[:len(current)-1]
    }
}
