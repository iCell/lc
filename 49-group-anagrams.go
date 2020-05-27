package main

import "sort"

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

func sortStr(s string) string {
	runes := sortRunes([]rune(s))
	sort.Sort(runes)
	return string(runes)
}

func groupAnagrams(strs []string) [][]string {
	temp := make(map[string][]int)
	for idx, str := range strs {
		sorted := sortStr(str)
		found, ok := temp[sorted]
		if !ok {
			temp[sorted] = []int{idx}
		} else {
			temp[sorted] = append(found, idx)
		}
	}
	var result [][]string
	for _, v := range temp {
		var sub []string
		for _, idx := range v {
			sub = append(sub, strs[idx])
		}
		result = append(result, sub)
	}
	return result
}
