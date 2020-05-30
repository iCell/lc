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
	for i, str := range strs {
		converted := sortStr(str)
		temp[converted] = append(temp[converted], i)
	}
	var r [][]string
	for _, v := range temp {
		var sub []string
		for _, i := range v {
			sub = append(sub, strs[i])
		}
		r = append(r, sub)
	}
	return r
}

func groupAnagrams2(strs []string) [][]string {
	var r [][]string

	temp := make(map[string]int)
	for _, str := range strs {
		converted := sortStr(str)
		index, ok := temp[converted]
		if ok {
			r[index] = append(r[index], str)
		} else {
			r = append(r, []string{str})
			temp[converted] = len(r) - 1
		}
	}

	return r
}

func groupAnagrams3(strs []string) [][]string {
	temp := make(map[string][]string)
	for _, str := range strs {
		r := make([]rune, 26, 26)
		for _, rune := range str {
			r[int(rune-'a')]++
		}
		temp[string(r)] = append(temp[string(r)], str)
	}
	return r
}
