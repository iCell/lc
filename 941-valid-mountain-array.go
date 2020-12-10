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