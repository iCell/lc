package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	lr, rr := 0, len(matrix)-1
	for lr <= rr {
		mr := lr + (rr-lr)>>1

		if target > matrix[mr][len(matrix[mr])-1] {
			lr = mr + 1
		} else if target < matrix[mr][0] {
			rr = mr - 1
		} else {
			row := matrix[mr]
			lr, rr = 0, len(row)-1
			for lr <= rr {
				mr := lr + (rr-lr)>>1
				if row[mr] == target {
					return true
				} else if row[mr] > target {
					rr = mr - 1
				} else {
					lr = mr + 1
				}
			}
			return false
		}
	}

	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])

	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) >> 1

		v := matrix[mid/n][mid%n]

		if v == target {
			return true
		}
		if v > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
