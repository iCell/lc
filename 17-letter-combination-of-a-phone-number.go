package main

var pairs = map[rune][]string{
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
 	result := make([]string, 0)
	helper(&result, 0, len(digits), digits, "")
	return result
}

func helper(result *[]string, level int, max int, digits string, temp string) {
	if level >= max {
		*result = append(*result, temp)
		return
	}

	r := []rune(digits)[level]
	for _, f := range pairs[r] {
		temp = temp + f
		helper(result, level+1, max, digits, temp)
		temp = temp[:len(temp)-1]

		// another approach
		//helper(result, level+1, max, digits, temp + f)
	}
}
