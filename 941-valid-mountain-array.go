func validMountainArray(arr []int) bool {
	var idx int
	for idx + 1 < len(arr) && arr[idx] < arr[idx+1] {
		idx += 1
	}
	
	if idx == 0 || idx == len(arr) - 1 {
		return false
	}
	
	for idx + 1 < len(arr) && arr[idx] > arr[idx+1] {
		idx += 1
	}
	
	return idx == len(arr) - 1
}

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	
	index := 1
	for index < len(arr) - 1 {
		if arr[index] > arr[index-1] && arr[index] > arr[index+1] {
			break
		}
		index++
	}
	
	if index == len(arr) - 1 {
		return false
	}

	for i := 1; i <= index; i++ {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	for i := index; i < len(arr) - 1; i++ {
		if arr[i] <= arr[i+1] {
			return false
		}
	}
	return true
}