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

func isAnagram2(s, t string) bool {
	sr, tr := []rune(s), []rune(t)
	sort.Sort(sortRunes(sr))
	sort.Sort(sortRunes(tr))
	return string(sr) == string(tr)
}
