func sortColors(nums []int) {
    p0, cur, p2 := 0, 0, len(nums)-1
    for cur <= p2 {
        switch nums[cur] {
        case 0:
            nums[p0], nums[cur] = nums[cur], nums[p0]
            p0++
            cur++
        case 2:
            nums[p2], nums[cur] = nums[cur], nums[p2]
            p2--
        case 1:
            cur++
        }
    }
}