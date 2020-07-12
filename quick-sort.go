package main

import "fmt"

func quickSort(arr []int) {
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, left, right int) {
	if right <= left {
		return
	}

	i, j := left, right
	pivot := right

	for i < j {
		for arr[i] <= arr[pivot] && i < j {
			i++
		}
		for arr[j] >= arr[pivot] && i < j {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
	}

	arr[right], arr[j] = arr[j], arr[right]
	pivot = j

	quickSortHelper(arr, left, pivot-1)
	quickSortHelper(arr, pivot+1, right)
}

func main() {
	arr := []int{6, 5, 2, 7, 3, 9, 8, 4, 10, 1}
	quickSort(arr)
	fmt.Println(arr)
}
