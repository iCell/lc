package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	index := -1
	left, right := 0, len(matrix)-1
	for left <= right {
		mid := left + (right-left)>>1

		mr := matrix[mid]

		if target >= mr[0] && target <= mr[len(mr)-1] {
			index = mid
		}
		if target > mr[len(mr)-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if index == -1 {
		return false
	}

	row := matrix[index]
	left, right = 0, len(row)-1
	for left <= right {
		mid := (left + right) / 2
		if target == row[mid] {
			return true
		}
		if target > row[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
