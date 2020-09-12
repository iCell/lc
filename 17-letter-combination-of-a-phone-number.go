var mapping = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	results := make([]string, 0)
	dfs(digits, 0, "", &results)
	return results
}

func dfs(origin string, idx int, digit string, digits *[]string) {
	if idx == len(origin) {
		*digits = append(*digits, digit)
		return
	}

	letter := origin[idx]
	for _, c := range mapping[letter] {
		dfs(origin, idx+1, digit+c, digits)
	}
}