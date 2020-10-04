func findPairs(nums []int, k int) int {
    visited, result := make(map[int]bool), make(map[int]int)
    for _, num := range nums {
        target1, target2 := num - k, num + k
        if visited[target1] {
            result[target1] = num
        }
        if visited[target2] {
            result[num] = target2
        }
        visited[num] = true
    }
    return len(result)
}