func generate(numRows int) [][]int {
    nums := make([][]int, numRows, numRows)
    for i := range nums {
        nums[i] = make([]int, i+1, i+1)
        if i == 0 {
            nums[0][0] = 1
            continue
        }
        for j := range nums[i] {
            if j == 0 || j == i {
                nums[i][j] = 1
            } else {
                nums[i][j] = nums[i-1][j-1] + nums[i-1][j]
            }
        }
    }
    return nums
}
