package main

import "sort"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	temp := make(map[rune]int, len(s))
	for _, r := range s {
		times, exist := temp[r]
		if exist {
			temp[r] = times + 1
		} else {
			temp[r] = 1
		}
	}
	for _, r := range t {
		times, exist := temp[r]
		if exist {
			temp[r] = times - 1
		} else {
			return false
		}
	}
	for _, v := range temp {
		if v > 0 {
			return false
		}
	}
	return true
}

func isAnagram2(s, t string) bool {
	sr, tr := []rune(s), []rune(t)
	sort.Sort(sortRunes(sr))
	sort.Sort(sortRunes(tr))
	return string(sr) == string(tr)
}

func isAnagram3(s string, t string) bool {
	sr, tr := []rune(s), []rune(t)
	quickSort(sr)
	quickSort(tr)
	return string(sr) == string(tr)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func quickSort(arr []rune) {
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []rune, left, right int) {
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
