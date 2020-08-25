func fourSum(nums []int, target int) [][]int {
    if len(nums) < 4 {
        return nil
    }
    
    sort.Ints(nums)
    
    var result [][]int
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        t1 := target - nums[i]
        for j := i + 1; j < len(nums); j++ {
            if j > i+1 && nums[j] == nums[j -a 1] {
                continue
            }
            
            t2 := t1 - nums[j]
            left, right := j + 1, len(nums) - 1
            for left < right {
                sum := nums[left] + nums[right]
                if sum == t2 {
                    result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
                    for left < right && nums[left] == nums[left+1] {
                        left++
                    }
                    for left < right && nums[right] == nums[right-1] {
                        right--
                    }
                    left++
                    right--
                } else if sum > t2 {
                    right--
                } else {
                    left++
                }
            }
        }
    }
    
    return result
}