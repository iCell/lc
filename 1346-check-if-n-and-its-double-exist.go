func checkIfExist(arr []int) bool {
	memo := make(map[int]int)
	for i, num := range arr {
		memo[num] = i
	}
	for i, num := range arr {
		j, exist := memo[num * 2]
		if exist && i != j {
			return true
		}
	}
	return false
}

func checkIfExist(arr []int) bool {
	memo := make(map[int]bool)
	for _, num := range arr {
		if memo[num * 2] || (num % 2 == 0 && memo[num / 2]) {
			return true
		}
		memo[num] = true
	}
	return false
}