func canFormArray(arr []int, pieces [][]int) bool {
	memo := make(map[int]int)
	for i := range pieces {
		if len(pieces[i]) == 0 {
			continue
		}
		memo[pieces[i][0]] = i
	}
	
	var index int
	for index < len(arr) {
		i, exist := memo[arr[index]]
		if !exist {
			return false
		}
		
		target := pieces[i]
		for _, num := range target {
			if num != arr[index] {
				return false
			}
			index++
		}
	}
	
	return true
}