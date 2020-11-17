func longestMountain(A []int) int {
	var max int
	var i int
	for i < len(A) {
		base := i
		for i + 1 < len(A) && A[i] < A[i+1] {
			i++
		}
		if base == i {
			i++
			continue
		}
		
		peak := i
		
		for i + 1 < len(A) && A[i] > A[i+1] {
			i++
		}
		if peak == i {
			continue
		}
		
		if i - base + 1 > max {
			max = i - base + 1
		}
	}
	return max
}