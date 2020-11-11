func flipAndInvertImage(A [][]int) [][]int {
	for i := range A {
		a := A[i]
		i, j := 0, len(a) - 1
		for i < j {
			a[i], a[j] = a[j], a[i]
			a[i] = a[i] ^ 1
			a[j] = a[j] ^ 1
			i++
			j--
		}
		if len(a) % 2 != 0 {
			a[i] = a[i] ^ 1
		}
	}
	return A
}
