import "sort"

func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    results := make([][]int, 0)
    dfs(candidates, target, 0, []int{}, &results)
    return results
}

func dfs(candidates []int, target int, idx int, temp []int, results *[][]int) {
    if target < 0 || idx == len(candidates) {
        return
    }
    if target == 0 {
        *results = append(*results, append([]int{}, temp...))
        return
    }

    // ignore current
    dfs(candidates, target, idx+1, temp, results)

    // select current
    current := candidates[idx]
    temp = append(temp, current)
    dfs(candidates, target-current, idx, temp, results)
    temp = temp[:len(temp)-1]
}

