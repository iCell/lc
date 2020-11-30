func sumZero(n int) []int {
	if n == 0 {
		return nil
	}
	
	result, cnt := make([]int, 0, n), 1
	if n % 2 != 0 {
		result = append(result, 0)
		n -= 1
	}
	
	for n != 0 {
		result = append(result, []int{cnt, -cnt}...)
		n -= 2
		cnt += 1
	}
	return result
}