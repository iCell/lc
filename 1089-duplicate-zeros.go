func duplicateZeros(arr []int)  {
	length := len(arr)
	for i := 0; i < length - 1; i++ {
		if arr[i] != 0 {
			continue
		}
		for j := len(arr) - 1; j > i; j-- {
			arr[j] = arr[j-1]
		}
		i += 1
		arr[i] = 0
	}
}

