func twoSumLessThanK(A []int, K int) int {
	if len(A) == 1 {
		return -1
	}
	sort.Ints(A)
	
	result, i, j := -1, 0, len(A) - 1
	for i < j {
		sum := A[i] + A[j]
		if sum >= K {
			j -= 1
		} else {
			result = max(result, sum);
			i += 1
		}
	}
	
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}